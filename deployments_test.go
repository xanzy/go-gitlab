package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDeploymentsService_ListProjectDeployments(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/deployments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"status": "created",
				"deployable": {
				  "commit": {
					"author_email": "admin@example.com",
					"author_name": "Administrator",
					"id": "99d03678b90d914dbb1b109132516d71a4a03ea8",
					"message": "Merge branch 'new-title' into 'main'\r\n\r\nUpdate README\r\n\r\n\r\n\r\nSee merge request !1",
					"short_id": "99d03678",
					"title": "Merge branch 'new-title' into 'main'\r"
				  },
				  "coverage": null,
				  "id": 657,
				  "name": "deploy",
				  "ref": "main",
				  "runner": null,
				  "stage": "deploy",
				  "status": "success",
				  "tag": false,
				  "user": {
					"id": 1,
					"name": "Administrator",
					"username": "root",
					"state": "active",
					"avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
					"web_url": "http://gitlab.dev/root",
					"bio": null,
					"location": null,
					"public_email": "",
					"skype": "",
					"linkedin": "",
					"twitter": "",
					"website_url": "",
					"organization": ""
				  },
				  "pipeline": {
					"id": 36,
					"ref": "main",
					"sha": "99d03678b90d914dbb1b109132516d71a4a03ea8",
					"status": "success",
					"web_url": "http://gitlab.dev/root/project/pipelines/12"
				  }
				},
				"environment": {
				  "external_url": "https://about.gitlab.com",
				  "id": 9,
				  "name": "production"
				},
				"id": 41,
				"iid": 1,
				"ref": "main",
				"sha": "99d03678b90d914dbb1b109132516d71a4a03ea8",
				"user": {
				  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				  "id": 1,
				  "name": "Administrator",
				  "state": "active",
				  "username": "root",
				  "web_url": "http://localhost:3000/root"
				}
			  }
			]
		`)
	})

	want := []*Deployment{{
		ID:     41,
		IID:    1,
		Ref:    "main",
		SHA:    "99d03678b90d914dbb1b109132516d71a4a03ea8",
		Status: "created",
		User: &ProjectUser{
			ID:        1,
			Name:      "Administrator",
			Username:  "root",
			State:     "active",
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://localhost:3000/root",
		},
		Environment: &Environment{
			ID:             9,
			Name:           "production",
			Slug:           "",
			State:          "",
			ExternalURL:    "https://about.gitlab.com",
			Project:        nil,
			LastDeployment: nil,
		},
		Deployable: struct {
			ID         int        `json:"id"`
			Status     string     `json:"status"`
			Stage      string     `json:"stage"`
			Name       string     `json:"name"`
			Ref        string     `json:"ref"`
			Tag        bool       `json:"tag"`
			Coverage   float64    `json:"coverage"`
			CreatedAt  *time.Time `json:"created_at"`
			StartedAt  *time.Time `json:"started_at"`
			FinishedAt *time.Time `json:"finished_at"`
			Duration   float64    `json:"duration"`
			User       *User      `json:"user"`
			Commit     *Commit    `json:"commit"`
			Pipeline   struct {
				ID        int        `json:"id"`
				SHA       string     `json:"sha"`
				Ref       string     `json:"ref"`
				Status    string     `json:"status"`
				CreatedAt *time.Time `json:"created_at"`
				UpdatedAt *time.Time `json:"updated_at"`
			} `json:"pipeline"`
			Runner *Runner `json:"runner"`
		}{
			ID:         657,
			Status:     "success",
			Stage:      "deploy",
			Name:       "deploy",
			Ref:        "main",
			Tag:        false,
			Coverage:   0,
			CreatedAt:  nil,
			StartedAt:  nil,
			FinishedAt: nil,
			Duration:   0,
			User: &User{
				ID:        1,
				Name:      "Administrator",
				Username:  "root",
				State:     "active",
				AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				WebURL:    "http://gitlab.dev/root",
			},
			Commit: &Commit{
				ID:             "99d03678b90d914dbb1b109132516d71a4a03ea8",
				ShortID:        "99d03678",
				Title:          "Merge branch 'new-title' into 'main'\r",
				AuthorName:     "Administrator",
				AuthorEmail:    "admin@example.com",
				AuthoredDate:   nil,
				CommitterName:  "",
				CommitterEmail: "",
				CommittedDate:  nil,
				CreatedAt:      nil,
				Message:        "Merge branch 'new-title' into 'main'\r\n\r\nUpdate README\r\n\r\n\r\n\r\nSee merge request !1",
				ParentIDs:      nil,
				Stats:          nil,
				Status:         nil,
				LastPipeline:   nil,
				ProjectID:      0,
				WebURL:         "",
			},
			Pipeline: struct {
				ID        int        `json:"id"`
				SHA       string     `json:"sha"`
				Ref       string     `json:"ref"`
				Status    string     `json:"status"`
				CreatedAt *time.Time `json:"created_at"`
				UpdatedAt *time.Time `json:"updated_at"`
			}{
				ID:        36,
				SHA:       "99d03678b90d914dbb1b109132516d71a4a03ea8",
				Ref:       "main",
				Status:    "success",
				CreatedAt: nil,
				UpdatedAt: nil,
			},
			Runner: nil,
		},
	}}

	ds, resp, err := client.Deployments.ListProjectDeployments(1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ds)

	ds, resp, err = client.Deployments.ListProjectDeployments(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ds)

	ds, resp, err = client.Deployments.ListProjectDeployments(1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ds)

	ds, resp, err = client.Deployments.ListProjectDeployments(3, nil, nil)
	require.Error(t, err)
	require.Nil(t, ds)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDeploymentsService_GetProjectDeployment(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/deployments/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"status": "created",
			"deployable": {
			  "commit": {
				"author_email": "admin@example.com",
				"author_name": "Administrator",
				"id": "99d03678b90d914dbb1b109132516d71a4a03ea8",
				"message": "Merge branch 'new-title' into 'main'\r\n\r\nUpdate README\r\n\r\n\r\n\r\nSee merge request !1",
				"short_id": "99d03678",
				"title": "Merge branch 'new-title' into 'main'\r"
			  },
			  "coverage": null,
			  "id": 657,
			  "name": "deploy",
			  "ref": "main",
			  "runner": null,
			  "stage": "deploy",
			  "status": "success",
			  "tag": false,
			  "user": {
				"id": 1,
				"name": "Administrator",
				"username": "root",
				"state": "active",
				"avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				"web_url": "http://gitlab.dev/root",
				"bio": null,
				"location": null,
				"public_email": "",
				"skype": "",
				"linkedin": "",
				"twitter": "",
				"website_url": "",
				"organization": ""
			  },
			  "pipeline": {
				"id": 36,
				"ref": "main",
				"sha": "99d03678b90d914dbb1b109132516d71a4a03ea8",
				"status": "success",
				"web_url": "http://gitlab.dev/root/project/pipelines/12"
			  }
			},
			"environment": {
			  "external_url": "https://about.gitlab.com",
			  "id": 9,
			  "name": "production"
			},
			"id": 41,
			"iid": 1,
			"ref": "main",
			"sha": "99d03678b90d914dbb1b109132516d71a4a03ea8",
			"user": {
			  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "id": 1,
			  "name": "Administrator",
			  "state": "active",
			  "username": "root",
			  "web_url": "http://localhost:3000/root"
			}
		  }
		`)
	})

	want := &Deployment{
		ID:     41,
		IID:    1,
		Ref:    "main",
		SHA:    "99d03678b90d914dbb1b109132516d71a4a03ea8",
		Status: "created",
		User: &ProjectUser{
			ID:        1,
			Name:      "Administrator",
			Username:  "root",
			State:     "active",
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://localhost:3000/root",
		},
		Environment: &Environment{
			ID:             9,
			Name:           "production",
			Slug:           "",
			State:          "",
			ExternalURL:    "https://about.gitlab.com",
			Project:        nil,
			LastDeployment: nil,
		},
		Deployable: struct {
			ID         int        `json:"id"`
			Status     string     `json:"status"`
			Stage      string     `json:"stage"`
			Name       string     `json:"name"`
			Ref        string     `json:"ref"`
			Tag        bool       `json:"tag"`
			Coverage   float64    `json:"coverage"`
			CreatedAt  *time.Time `json:"created_at"`
			StartedAt  *time.Time `json:"started_at"`
			FinishedAt *time.Time `json:"finished_at"`
			Duration   float64    `json:"duration"`
			User       *User      `json:"user"`
			Commit     *Commit    `json:"commit"`
			Pipeline   struct {
				ID        int        `json:"id"`
				SHA       string     `json:"sha"`
				Ref       string     `json:"ref"`
				Status    string     `json:"status"`
				CreatedAt *time.Time `json:"created_at"`
				UpdatedAt *time.Time `json:"updated_at"`
			} `json:"pipeline"`
			Runner *Runner `json:"runner"`
		}{
			ID:         657,
			Status:     "success",
			Stage:      "deploy",
			Name:       "deploy",
			Ref:        "main",
			Tag:        false,
			Coverage:   0,
			CreatedAt:  nil,
			StartedAt:  nil,
			FinishedAt: nil,
			Duration:   0,
			User: &User{
				ID:        1,
				Name:      "Administrator",
				Username:  "root",
				State:     "active",
				AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				WebURL:    "http://gitlab.dev/root",
			},
			Commit: &Commit{
				ID:             "99d03678b90d914dbb1b109132516d71a4a03ea8",
				ShortID:        "99d03678",
				Title:          "Merge branch 'new-title' into 'main'\r",
				AuthorName:     "Administrator",
				AuthorEmail:    "admin@example.com",
				AuthoredDate:   nil,
				CommitterName:  "",
				CommitterEmail: "",
				CommittedDate:  nil,
				CreatedAt:      nil,
				Message:        "Merge branch 'new-title' into 'main'\r\n\r\nUpdate README\r\n\r\n\r\n\r\nSee merge request !1",
				ParentIDs:      nil,
				Stats:          nil,
				Status:         nil,
				LastPipeline:   nil,
				ProjectID:      0,
				WebURL:         "",
			},
			Pipeline: struct {
				ID        int        `json:"id"`
				SHA       string     `json:"sha"`
				Ref       string     `json:"ref"`
				Status    string     `json:"status"`
				CreatedAt *time.Time `json:"created_at"`
				UpdatedAt *time.Time `json:"updated_at"`
			}{
				ID:        36,
				SHA:       "99d03678b90d914dbb1b109132516d71a4a03ea8",
				Ref:       "main",
				Status:    "success",
				CreatedAt: nil,
				UpdatedAt: nil,
			},
			Runner: nil,
		},
	}

	d, resp, err := client.Deployments.GetProjectDeployment(1, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)

	d, resp, err = client.Deployments.GetProjectDeployment(1.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Deployments.GetProjectDeployment(1, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Deployments.GetProjectDeployment(3, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, d)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDeploymentsService_CreateProjectDeployment(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/deployments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"status": "created",
			"deployable": {
			  "commit": {
				"author_email": "admin@example.com",
				"author_name": "Administrator",
				"id": "99d03678b90d914dbb1b109132516d71a4a03ea8",
				"message": "Merge branch 'new-title' into 'main'\r\n\r\nUpdate README\r\n\r\n\r\n\r\nSee merge request !1",
				"short_id": "99d03678",
				"title": "Merge branch 'new-title' into 'main'\r"
			  },
			  "coverage": null,
			  "id": 657,
			  "name": "deploy",
			  "ref": "main",
			  "runner": null,
			  "stage": "deploy",
			  "status": "success",
			  "tag": false,
			  "user": {
				"id": 1,
				"name": "Administrator",
				"username": "root",
				"state": "active",
				"avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				"web_url": "http://gitlab.dev/root",
				"bio": null,
				"location": null,
				"public_email": "",
				"skype": "",
				"linkedin": "",
				"twitter": "",
				"website_url": "",
				"organization": ""
			  },
			  "pipeline": {
				"id": 36,
				"ref": "main",
				"sha": "99d03678b90d914dbb1b109132516d71a4a03ea8",
				"status": "success",
				"web_url": "http://gitlab.dev/root/project/pipelines/12"
			  }
			},
			"environment": {
			  "external_url": "https://about.gitlab.com",
			  "id": 9,
			  "name": "production"
			},
			"id": 41,
			"iid": 1,
			"ref": "main",
			"sha": "99d03678b90d914dbb1b109132516d71a4a03ea8",
			"user": {
			  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "id": 1,
			  "name": "Administrator",
			  "state": "active",
			  "username": "root",
			  "web_url": "http://localhost:3000/root"
			}
		  }
		`)
	})

	want := &Deployment{
		ID:     41,
		IID:    1,
		Ref:    "main",
		SHA:    "99d03678b90d914dbb1b109132516d71a4a03ea8",
		Status: "created",
		User: &ProjectUser{
			ID:        1,
			Name:      "Administrator",
			Username:  "root",
			State:     "active",
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://localhost:3000/root",
		},
		Environment: &Environment{
			ID:             9,
			Name:           "production",
			Slug:           "",
			State:          "",
			ExternalURL:    "https://about.gitlab.com",
			Project:        nil,
			LastDeployment: nil,
		},
		Deployable: struct {
			ID         int        `json:"id"`
			Status     string     `json:"status"`
			Stage      string     `json:"stage"`
			Name       string     `json:"name"`
			Ref        string     `json:"ref"`
			Tag        bool       `json:"tag"`
			Coverage   float64    `json:"coverage"`
			CreatedAt  *time.Time `json:"created_at"`
			StartedAt  *time.Time `json:"started_at"`
			FinishedAt *time.Time `json:"finished_at"`
			Duration   float64    `json:"duration"`
			User       *User      `json:"user"`
			Commit     *Commit    `json:"commit"`
			Pipeline   struct {
				ID        int        `json:"id"`
				SHA       string     `json:"sha"`
				Ref       string     `json:"ref"`
				Status    string     `json:"status"`
				CreatedAt *time.Time `json:"created_at"`
				UpdatedAt *time.Time `json:"updated_at"`
			} `json:"pipeline"`
			Runner *Runner `json:"runner"`
		}{
			ID:         657,
			Status:     "success",
			Stage:      "deploy",
			Name:       "deploy",
			Ref:        "main",
			Tag:        false,
			Coverage:   0,
			CreatedAt:  nil,
			StartedAt:  nil,
			FinishedAt: nil,
			Duration:   0,
			User: &User{
				ID:        1,
				Name:      "Administrator",
				Username:  "root",
				State:     "active",
				AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				WebURL:    "http://gitlab.dev/root",
			},
			Commit: &Commit{
				ID:             "99d03678b90d914dbb1b109132516d71a4a03ea8",
				ShortID:        "99d03678",
				Title:          "Merge branch 'new-title' into 'main'\r",
				AuthorName:     "Administrator",
				AuthorEmail:    "admin@example.com",
				AuthoredDate:   nil,
				CommitterName:  "",
				CommitterEmail: "",
				CommittedDate:  nil,
				CreatedAt:      nil,
				Message:        "Merge branch 'new-title' into 'main'\r\n\r\nUpdate README\r\n\r\n\r\n\r\nSee merge request !1",
				ParentIDs:      nil,
				Stats:          nil,
				Status:         nil,
				LastPipeline:   nil,
				ProjectID:      0,
				WebURL:         "",
			},
			Pipeline: struct {
				ID        int        `json:"id"`
				SHA       string     `json:"sha"`
				Ref       string     `json:"ref"`
				Status    string     `json:"status"`
				CreatedAt *time.Time `json:"created_at"`
				UpdatedAt *time.Time `json:"updated_at"`
			}{
				ID:        36,
				SHA:       "99d03678b90d914dbb1b109132516d71a4a03ea8",
				Ref:       "main",
				Status:    "success",
				CreatedAt: nil,
				UpdatedAt: nil,
			},
			Runner: nil,
		},
	}

	d, resp, err := client.Deployments.CreateProjectDeployment(1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)

	d, resp, err = client.Deployments.CreateProjectDeployment(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Deployments.CreateProjectDeployment(1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Deployments.CreateProjectDeployment(3, nil, nil)
	require.Error(t, err)
	require.Nil(t, d)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDeploymentsService_UpdateProjectDeployment(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/deployments/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
		  {
			"status": "created",
			"deployable": {
			  "commit": {
				"author_email": "admin@example.com",
				"author_name": "Administrator",
				"id": "99d03678b90d914dbb1b109132516d71a4a03ea8",
				"message": "Merge branch 'new-title' into 'main'\r\n\r\nUpdate README\r\n\r\n\r\n\r\nSee merge request !1",
				"short_id": "99d03678",
				"title": "Merge branch 'new-title' into 'main'\r"
			  },
			  "coverage": null,
			  "id": 657,
			  "name": "deploy",
			  "ref": "main",
			  "runner": null,
			  "stage": "deploy",
			  "status": "success",
			  "tag": false,
			  "user": {
				"id": 1,
				"name": "Administrator",
				"username": "root",
				"state": "active",
				"avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				"web_url": "http://gitlab.dev/root",
				"bio": null,
				"location": null,
				"public_email": "",
				"skype": "",
				"linkedin": "",
				"twitter": "",
				"website_url": "",
				"organization": ""
			  },
			  "pipeline": {
				"id": 36,
				"ref": "main",
				"sha": "99d03678b90d914dbb1b109132516d71a4a03ea8",
				"status": "success",
				"web_url": "http://gitlab.dev/root/project/pipelines/12"
			  }
			},
			"environment": {
			  "external_url": "https://about.gitlab.com",
			  "id": 9,
			  "name": "production"
			},
			"id": 41,
			"iid": 1,
			"ref": "main",
			"sha": "99d03678b90d914dbb1b109132516d71a4a03ea8",
			"user": {
			  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "id": 1,
			  "name": "Administrator",
			  "state": "active",
			  "username": "root",
			  "web_url": "http://localhost:3000/root"
			}
		  }
		`)
	})

	want := &Deployment{
		ID:     41,
		IID:    1,
		Ref:    "main",
		SHA:    "99d03678b90d914dbb1b109132516d71a4a03ea8",
		Status: "created",
		User: &ProjectUser{
			ID:        1,
			Name:      "Administrator",
			Username:  "root",
			State:     "active",
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://localhost:3000/root",
		},
		Environment: &Environment{
			ID:             9,
			Name:           "production",
			Slug:           "",
			State:          "",
			ExternalURL:    "https://about.gitlab.com",
			Project:        nil,
			LastDeployment: nil,
		},
		Deployable: struct {
			ID         int        `json:"id"`
			Status     string     `json:"status"`
			Stage      string     `json:"stage"`
			Name       string     `json:"name"`
			Ref        string     `json:"ref"`
			Tag        bool       `json:"tag"`
			Coverage   float64    `json:"coverage"`
			CreatedAt  *time.Time `json:"created_at"`
			StartedAt  *time.Time `json:"started_at"`
			FinishedAt *time.Time `json:"finished_at"`
			Duration   float64    `json:"duration"`
			User       *User      `json:"user"`
			Commit     *Commit    `json:"commit"`
			Pipeline   struct {
				ID        int        `json:"id"`
				SHA       string     `json:"sha"`
				Ref       string     `json:"ref"`
				Status    string     `json:"status"`
				CreatedAt *time.Time `json:"created_at"`
				UpdatedAt *time.Time `json:"updated_at"`
			} `json:"pipeline"`
			Runner *Runner `json:"runner"`
		}{
			ID:         657,
			Status:     "success",
			Stage:      "deploy",
			Name:       "deploy",
			Ref:        "main",
			Tag:        false,
			Coverage:   0,
			CreatedAt:  nil,
			StartedAt:  nil,
			FinishedAt: nil,
			Duration:   0,
			User: &User{
				ID:        1,
				Name:      "Administrator",
				Username:  "root",
				State:     "active",
				AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				WebURL:    "http://gitlab.dev/root",
			},
			Commit: &Commit{
				ID:             "99d03678b90d914dbb1b109132516d71a4a03ea8",
				ShortID:        "99d03678",
				Title:          "Merge branch 'new-title' into 'main'\r",
				AuthorName:     "Administrator",
				AuthorEmail:    "admin@example.com",
				AuthoredDate:   nil,
				CommitterName:  "",
				CommitterEmail: "",
				CommittedDate:  nil,
				CreatedAt:      nil,
				Message:        "Merge branch 'new-title' into 'main'\r\n\r\nUpdate README\r\n\r\n\r\n\r\nSee merge request !1",
				ParentIDs:      nil,
				Stats:          nil,
				Status:         nil,
				LastPipeline:   nil,
				ProjectID:      0,
				WebURL:         "",
			},
			Pipeline: struct {
				ID        int        `json:"id"`
				SHA       string     `json:"sha"`
				Ref       string     `json:"ref"`
				Status    string     `json:"status"`
				CreatedAt *time.Time `json:"created_at"`
				UpdatedAt *time.Time `json:"updated_at"`
			}{
				ID:        36,
				SHA:       "99d03678b90d914dbb1b109132516d71a4a03ea8",
				Ref:       "main",
				Status:    "success",
				CreatedAt: nil,
				UpdatedAt: nil,
			},
			Runner: nil,
		},
	}

	d, resp, err := client.Deployments.UpdateProjectDeployment(1, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)

	d, resp, err = client.Deployments.UpdateProjectDeployment(1.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Deployments.UpdateProjectDeployment(1, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Deployments.UpdateProjectDeployment(3, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, d)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
