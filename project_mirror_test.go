package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProjectMirrorService_ListProjectMirror(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/42/remote_mirrors", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"enabled": true,
				"id": 101486,
				"last_error": null,
				"only_protected_branches": true,
				"keep_divergent_refs": true,
				"update_status": "finished",
				"url": "https://*****:*****@gitlab.com/gitlab-org/security/gitlab.git"
			  }
			]
		`)
	})

	want := []*ProjectMirror{{
		Enabled:               true,
		ID:                    101486,
		LastError:             "",
		OnlyProtectedBranches: true,
		KeepDivergentRefs:     true,
		UpdateStatus:          "finished",
		URL:                   "https://*****:*****@gitlab.com/gitlab-org/security/gitlab.git",
	}}

	pms, resp, err := client.ProjectMirrors.ListProjectMirror(42, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pms)

	pms, resp, err = client.ProjectMirrors.ListProjectMirror(42.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 42.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pms)

	pms, resp, err = client.ProjectMirrors.ListProjectMirror(42, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pms)

	pms, resp, err = client.ProjectMirrors.ListProjectMirror(43, nil, nil)
	require.Error(t, err)
	require.Nil(t, pms)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectMirrorService_GetProjectMirror(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/42/remote_mirrors/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
				"enabled": true,
				"id": 101486,
				"last_error": null,
				"only_protected_branches": true,
				"keep_divergent_refs": true,
				"update_status": "finished",
				"url": "https://*****:*****@gitlab.com/gitlab-org/security/gitlab.git"
			}
		`)
	})

	want := &ProjectMirror{
		Enabled:               true,
		ID:                    101486,
		LastError:             "",
		OnlyProtectedBranches: true,
		KeepDivergentRefs:     true,
		UpdateStatus:          "finished",
		URL:                   "https://*****:*****@gitlab.com/gitlab-org/security/gitlab.git",
	}

	pm, resp, err := client.ProjectMirrors.GetProjectMirror(42, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pm)
}

func TestProjectMirrorService_AddProjectMirror(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/42/remote_mirrors", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
				"enabled": false,
				"id": 101486,
				"last_error": null,
				"last_successful_update_at": null,
				"last_update_at": null,
				"last_update_started_at": null,
				"only_protected_branches": false,
				"keep_divergent_refs": false,
				"update_status": "none",
				"url": "https://*****:*****@example.com/gitlab/example.git"
			}
		`)
	})

	want := &ProjectMirror{
		Enabled:                false,
		ID:                     101486,
		LastError:              "",
		LastSuccessfulUpdateAt: nil,
		LastUpdateAt:           nil,
		LastUpdateStartedAt:    nil,
		OnlyProtectedBranches:  false,
		KeepDivergentRefs:      false,
		UpdateStatus:           "none",
		URL:                    "https://*****:*****@example.com/gitlab/example.git",
	}

	pm, resp, err := client.ProjectMirrors.AddProjectMirror(42, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pm)

	pm, resp, err = client.ProjectMirrors.AddProjectMirror(42.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 42.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pm)

	pm, resp, err = client.ProjectMirrors.AddProjectMirror(42, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pm)

	pm, resp, err = client.ProjectMirrors.AddProjectMirror(43, nil, nil)
	require.Error(t, err)
	require.Nil(t, pm)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectMirrorService_EditProjectMirror(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/42/remote_mirrors/101486", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
			{
				"enabled": false,
				"id": 101486,
				"last_error": null,
				"only_protected_branches": true,
				"keep_divergent_refs": true,
				"update_status": "finished",
				"url": "https://*****:*****@gitlab.com/gitlab-org/security/gitlab.git"
			}
		`)
	})

	want := &ProjectMirror{
		Enabled:               false,
		ID:                    101486,
		LastError:             "",
		OnlyProtectedBranches: true,
		KeepDivergentRefs:     true,
		UpdateStatus:          "finished",
		URL:                   "https://*****:*****@gitlab.com/gitlab-org/security/gitlab.git",
	}

	pm, resp, err := client.ProjectMirrors.EditProjectMirror(42, 101486, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pm)

	pm, resp, err = client.ProjectMirrors.EditProjectMirror(42.01, 101486, nil, nil)
	require.EqualError(t, err, "invalid ID type 42.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pm)

	pm, resp, err = client.ProjectMirrors.EditProjectMirror(42, 101486, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pm)

	pm, resp, err = client.ProjectMirrors.EditProjectMirror(43, 101486, nil, nil)
	require.Error(t, err)
	require.Nil(t, pm)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
