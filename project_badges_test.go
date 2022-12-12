package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProjectBadgesService_ListProjectBadges(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/badges", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"name": "Coverage",
				"id": 1,
				"link_url": "http://example.com/ci_status.svg?project={project_path}&ref={default_branch}",
				"image_url": "https://shields.io/my/badge",
				"rendered_link_url": "http://example.com/ci_status.svg?project=example-org/example-project&ref=master",
				"rendered_image_url": "https://shields.io/my/badge",
				"kind": "project"
			  }
			]
		`)
	})

	want := []*ProjectBadge{{
		ID:               1,
		Name:             "Coverage",
		LinkURL:          "http://example.com/ci_status.svg?project={project_path}&ref={default_branch}",
		ImageURL:         "https://shields.io/my/badge",
		RenderedLinkURL:  "http://example.com/ci_status.svg?project=example-org/example-project&ref=master",
		RenderedImageURL: "https://shields.io/my/badge",
		Kind:             "project",
	}}

	pbs, resp, err := client.ProjectBadges.ListProjectBadges(1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pbs)

	pbs, resp, err = client.ProjectBadges.ListProjectBadges(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pbs)

	pbs, resp, err = client.ProjectBadges.ListProjectBadges(1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pbs)

	pbs, resp, err = client.ProjectBadges.ListProjectBadges(2, nil, nil)
	require.Error(t, err)
	require.Nil(t, pbs)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectBadgesService_GetProjectBadge(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/badges/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"name": "Coverage",
			"id": 1,
			"link_url": "http://example.com/ci_status.svg?project={project_path}&ref={default_branch}",
			"image_url": "https://shields.io/my/badge",
			"rendered_link_url": "http://example.com/ci_status.svg?project=example-org/example-project&ref=master",
			"rendered_image_url": "https://shields.io/my/badge",
			"kind": "project"
		  }
		`)
	})

	want := &ProjectBadge{
		ID:               1,
		Name:             "Coverage",
		LinkURL:          "http://example.com/ci_status.svg?project={project_path}&ref={default_branch}",
		ImageURL:         "https://shields.io/my/badge",
		RenderedLinkURL:  "http://example.com/ci_status.svg?project=example-org/example-project&ref=master",
		RenderedImageURL: "https://shields.io/my/badge",
		Kind:             "project",
	}

	pb, resp, err := client.ProjectBadges.GetProjectBadge(1, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pb)

	pb, resp, err = client.ProjectBadges.GetProjectBadge(1.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pb)

	pb, resp, err = client.ProjectBadges.GetProjectBadge(1, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pb)

	pb, resp, err = client.ProjectBadges.GetProjectBadge(2, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, pb)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectBadgesService_AddProjectBadge(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/badges", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"name": "mybadge",
			"id": 1,
			"link_url": "http://example.com/ci_status.svg?project={project_path}&ref={default_branch}",
			"image_url": "https://shields.io/my/badge",
			"rendered_link_url": "http://example.com/ci_status.svg?project=example-org/example-project&ref=master",
			"rendered_image_url": "https://shields.io/my/badge",
			"kind": "project"
		  }
		`)
	})

	want := &ProjectBadge{
		ID:               1,
		Name:             "mybadge",
		LinkURL:          "http://example.com/ci_status.svg?project={project_path}&ref={default_branch}",
		ImageURL:         "https://shields.io/my/badge",
		RenderedLinkURL:  "http://example.com/ci_status.svg?project=example-org/example-project&ref=master",
		RenderedImageURL: "https://shields.io/my/badge",
		Kind:             "project",
	}

	pb, resp, err := client.ProjectBadges.AddProjectBadge(1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pb)

	pb, resp, err = client.ProjectBadges.AddProjectBadge(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pb)

	pb, resp, err = client.ProjectBadges.AddProjectBadge(1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pb)

	pb, resp, err = client.ProjectBadges.AddProjectBadge(2, nil, nil)
	require.Error(t, err)
	require.Nil(t, pb)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectBadgesService_EditProjectBadge(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/badges/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
		  {
			"name": "mybadge",
			"id": 1,
			"link_url": "http://example.com/ci_status.svg?project={project_path}&ref={default_branch}",
			"image_url": "https://shields.io/my/badge",
			"rendered_link_url": "http://example.com/ci_status.svg?project=example-org/example-project&ref=master",
			"rendered_image_url": "https://shields.io/my/badge",
			"kind": "project"
		  }
		`)
	})

	want := &ProjectBadge{
		ID:               1,
		Name:             "mybadge",
		LinkURL:          "http://example.com/ci_status.svg?project={project_path}&ref={default_branch}",
		ImageURL:         "https://shields.io/my/badge",
		RenderedLinkURL:  "http://example.com/ci_status.svg?project=example-org/example-project&ref=master",
		RenderedImageURL: "https://shields.io/my/badge",
		Kind:             "project",
	}

	pb, resp, err := client.ProjectBadges.EditProjectBadge(1, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pb)

	pb, resp, err = client.ProjectBadges.EditProjectBadge(1.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pb)

	pb, resp, err = client.ProjectBadges.EditProjectBadge(1, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pb)

	pb, resp, err = client.ProjectBadges.EditProjectBadge(2, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, pb)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectBadgesService_DeleteProjectBadge(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/badges/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.ProjectBadges.DeleteProjectBadge(1, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.ProjectBadges.DeleteProjectBadge(1.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.ProjectBadges.DeleteProjectBadge(1, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.ProjectBadges.DeleteProjectBadge(2, 1, nil, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectBadgesService_PreviewProjectBadge(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/badges/render", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"link_url": "http://example.com/ci_status.svg?project={project_path}&ref={default_branch}",
			"image_url": "https://shields.io/my/badge",
			"rendered_link_url": "http://example.com/ci_status.svg?project=example-org/example-project&ref=master",
			"rendered_image_url": "https://shields.io/my/badge"
		  }
		`)
	})

	want := &ProjectBadge{
		LinkURL:          "http://example.com/ci_status.svg?project={project_path}&ref={default_branch}",
		ImageURL:         "https://shields.io/my/badge",
		RenderedLinkURL:  "http://example.com/ci_status.svg?project=example-org/example-project&ref=master",
		RenderedImageURL: "https://shields.io/my/badge",
	}

	pb, resp, err := client.ProjectBadges.PreviewProjectBadge(1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pb)

	pb, resp, err = client.ProjectBadges.PreviewProjectBadge(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pb)

	pb, resp, err = client.ProjectBadges.PreviewProjectBadge(1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pb)

	pb, resp, err = client.ProjectBadges.PreviewProjectBadge(2, nil, nil)
	require.Error(t, err)
	require.Nil(t, pb)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
