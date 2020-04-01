package gitlab

import (
	"crypto/md5"
	"encoding/base64"
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerraformService_LockUnlockState(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	lockInfo := lockInfoTest()
	mux.HandleFunc("/api/v4/projects/1/terraform/state/test1/lock", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost: // lock
			body := testBodyJSON(t, r, &lockInfo)
			assertMD5HeaderMatchesBody(t, r, body)
			// TODO Does it have to reply with something? Not clear.
		case http.MethodDelete: // unlock
			body := testBodyJSON(t, r, &lockInfo)
			assertMD5HeaderMatchesBody(t, r, body)
		default:
			assert.Failf(t, "unexpected HTTP method: %s", r.Method)
		}
	})

	lockData, err := client.Terraform.LockState(1, "test1", &lockInfo)
	require.NoError(t, err)

	assert.Equal(t, lockInfo, lockData.RequestedLockInfo)
	assert.Equal(t, "1", lockData.Project)
	assert.Equal(t, "test1", lockData.StateName)

	err = client.Terraform.UnlockState(lockData)
	require.NoError(t, err)
}

func TestTerraformService_LockState_Failures(t *testing.T) {
	tests := []struct {
		statusCode int
		msg        string
	}{
		{
			statusCode: http.StatusForbidden,
			msg:        "remote state endpoint invalid auth",
		},
		{
			statusCode: http.StatusConflict,
			msg:        "remote state already locked: ID=id1",
		},
		{
			statusCode: http.StatusLocked,
			msg:        "remote state already locked: ID=id1",
		},
	}
	for _, test := range tests {
		t.Run(strconv.Itoa(test.statusCode), func(t *testing.T) {
			mux, server, client := setup(t)
			defer teardown(server)

			lockInfo := lockInfoTest()
			mux.HandleFunc("/api/v4/projects/1/terraform/state/test1/lock", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(test.statusCode)
				_, err := io.Copy(w, r.Body) // just respond with the same ID
				assert.NoError(t, err)
			})
			_, err := client.Terraform.LockState(1, "test1", &lockInfo)
			assert.EqualError(t, err, test.msg)
		})
	}
}

func TestTerraformService_GetState_NoMd5(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	body := []byte(`{"state":42}`)
	mux.HandleFunc("/api/v4/projects/1/terraform/state/test1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeader(t, r, "Accept", "application/json")
		_, err := w.Write(body)
		assert.NoError(t, err)
	})

	payload, err := client.Terraform.GetState(1, "test1")
	require.NoError(t, err)
	assertMD5(t, base64.StdEncoding.EncodeToString(payload.MD5), payload.Data)
	assert.Equal(t, body, payload.Data)
}

func TestTerraformService_GetState_Md5(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	body := []byte(`{"state":42}`)
	mux.HandleFunc("/api/v4/projects/1/terraform/state/test1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Md5", md5Str(body))
		_, err := w.Write(body)
		assert.NoError(t, err)
	})

	payload, err := client.Terraform.GetState(1, "test1")
	require.NoError(t, err)
	assertMD5(t, base64.StdEncoding.EncodeToString(payload.MD5), payload.Data)
	assert.Equal(t, body, payload.Data)
}

func TestTerraformService_GetState_EmptyResponse(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/terraform/state/test1", func(w http.ResponseWriter, r *http.Request) {
	})

	payload, err := client.Terraform.GetState(1, "test1")
	require.NoError(t, err)
	assert.Nil(t, payload)
}

func lockInfoTest() LockInfo {
	return LockInfo{
		ID:        "id1",
		Operation: "op",
		Info:      "info",
		Who:       "who",
		Version:   "version",
		Created:   time.Now().UTC(),
		Path:      "path",
	}
}

func assertMD5HeaderMatchesBody(t *testing.T, r *http.Request, body []byte) {
	testHeader(t, r, "Content-Md5", md5Str(body))
}

func assertMD5(t *testing.T, md5expected string, data []byte) {
	assert.Equal(t, md5expected, md5Str(data))
}

func md5Str(data []byte) string {
	hash := md5.Sum(data)
	return base64.StdEncoding.EncodeToString(hash[:])
}
