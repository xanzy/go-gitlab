package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLicenseService_GetLicense(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/license", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
			  "id": 2,
			  "plan": "gold",
			  "historical_max": 300,
			  "maximum_user_count": 300,
			  "expired": false,
			  "overage": 200,
			  "user_limit": 100,
			  "active_users": 300,
			  "licensee": {
				"Name": "Venkatesh Thalluri"
			  },
			  "add_ons": {
				"GitLab_FileLocks": 1,
				"GitLab_Auditor_User": 1
			  }
			}
		`)
	})

	want := &License{
		ID:               2,
		Plan:             "gold",
		HistoricalMax:    300,
		MaximumUserCount: 300,
		Expired:          false,
		Overage:          200,
		UserLimit:        100,
		ActiveUsers:      300,
		Licensee: struct {
			Name    string `json:"Name"`
			Company string `json:"Company"`
			Email   string `json:"Email"`
		}{
			Name:    "Venkatesh Thalluri",
			Company: "",
			Email:   "",
		},
		AddOns: struct {
			GitLabAuditorUser int `json:"GitLab_Auditor_User"`
			GitLabDeployBoard int `json:"GitLab_DeployBoard"`
			GitLabFileLocks   int `json:"GitLab_FileLocks"`
			GitLabGeo         int `json:"GitLab_Geo"`
			GitLabServiceDesk int `json:"GitLab_ServiceDesk"`
		}{
			GitLabAuditorUser: 1,
			GitLabDeployBoard: 0,
			GitLabFileLocks:   1,
			GitLabGeo:         0,
			GitLabServiceDesk: 0,
		},
	}

	l, resp, err := client.License.GetLicense()
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, l)
}

func TestLicenseService_GetLicense_StatusNotFound(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/license", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusNotFound)
	})

	l, resp, err := client.License.GetLicense()
	require.Error(t, err)
	require.Nil(t, l)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestLicenseService_AddLicense(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/license", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
			  "id": 2,
			  "plan": "gold",
			  "historical_max": 300,
			  "maximum_user_count": 300,
			  "expired": false,
			  "overage": 200,
			  "user_limit": 100,
			  "active_users": 300,
			  "licensee": {
				"Name": "Venkatesh Thalluri"
			  },
			  "add_ons": {
				"GitLab_FileLocks": 1,
				"GitLab_Auditor_User": 1
			  }
			}
		`)
	})

	want := &License{
		ID:               2,
		Plan:             "gold",
		HistoricalMax:    300,
		MaximumUserCount: 300,
		Expired:          false,
		Overage:          200,
		UserLimit:        100,
		ActiveUsers:      300,
		Licensee: struct {
			Name    string `json:"Name"`
			Company string `json:"Company"`
			Email   string `json:"Email"`
		}{
			Name:    "Venkatesh Thalluri",
			Company: "",
			Email:   "",
		},
		AddOns: struct {
			GitLabAuditorUser int `json:"GitLab_Auditor_User"`
			GitLabDeployBoard int `json:"GitLab_DeployBoard"`
			GitLabFileLocks   int `json:"GitLab_FileLocks"`
			GitLabGeo         int `json:"GitLab_Geo"`
			GitLabServiceDesk int `json:"GitLab_ServiceDesk"`
		}{
			GitLabAuditorUser: 1,
			GitLabDeployBoard: 0,
			GitLabFileLocks:   1,
			GitLabGeo:         0,
			GitLabServiceDesk: 0,
		},
	}

	l, resp, err := client.License.AddLicense(nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, l)

	l, resp, err = client.License.AddLicense(nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, l)
}

func TestLicenseService_AddLicense_StatusNotFound(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/license", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	l, resp, err := client.License.AddLicense(nil)
	require.Error(t, err)
	require.Nil(t, l)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
