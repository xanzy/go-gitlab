package gitlab

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestListGroupSSHCertificates(t *testing.T) {
	mux, client := setup(t)

	path := "/api/v4/groups/1/ssh_certificates"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_group_ssh_certificates.json")
	})

	certificates, _, err := client.GroupSSHCertificates.ListGroupSSHCertificates(1)
	require.NoError(t, err)

	want := []*GroupSSHCertificate{
		{
			ID:        1876,
			Title:     "SSH Certificate",
			Key:       "ssh-rsa FAKE-KEY example@gitlab.com",
			CreatedAt: Ptr(time.Date(2022, time.March, 20, 20, 42, 40, 221000000, time.UTC)),
		},
	}

	require.Equal(t, want, certificates)
}

func TestCreateGroupSSHCertificate(t *testing.T) {
	mux, client := setup(t)

	path := "/api/v4/groups/84/ssh_certificates"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		mustWriteHTTPResponse(t, w, "testdata/create_group_ssh_certificates.json")
	})

	cert, _, err := client.GroupSSHCertificates.CreateGroupSSHCertificate(84, &CreateGroupSSHCertificateOptions{
		Key:   Ptr("ssh-rsa FAKE-KEY example@gitlab.com"),
		Title: Ptr("SSH Certificate"),
	})
	require.NoError(t, err)

	want := &GroupSSHCertificate{
		ID:        1876,
		Title:     "SSH Certificate",
		Key:       "ssh-rsa FAKE-KEY example@gitlab.com",
		CreatedAt: Ptr(time.Date(2022, time.March, 20, 20, 42, 40, 221000000, time.UTC)),
	}

	require.Equal(t, want, cert)
}

func TestDeleteGroupSSHCertificate(t *testing.T) {
	mux, client := setup(t)

	path := "/api/v4/groups/1/ssh_certificates/1876"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.GroupSSHCertificates.DeleteGroupSSHCertificate(1, 1876)
	require.NoError(t, err)
}
