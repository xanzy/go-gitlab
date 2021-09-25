package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestProjectSnippetsService_ListSnippets(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/snippets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
				{
				  "id": 1,
				  "title": "test",
				  "file_name": "add.rb",
				  "description": "Ruby test snippet",
				  "author": {
					"id": 1,
					"username": "venkatesh_thalluri",
					"email": "venky@example.com",
					"name": "Venkatesh Thalluri",
					"state": "active"
				  },
				  "project_id": 1,
				  "web_url": "http://example.com/example/example/snippets/1",
				  "raw_url": "http://example.com/example/example/snippets/1/raw"
				}
			]
		`)
	})

	want := []*Snippet{{
		ID:          1,
		Title:       "test",
		FileName:    "add.rb",
		Description: "Ruby test snippet",
		Author: struct {
			ID        int        `json:"id"`
			Username  string     `json:"username"`
			Email     string     `json:"email"`
			Name      string     `json:"name"`
			State     string     `json:"state"`
			CreatedAt *time.Time `json:"created_at"`
		}{
			ID:       1,
			Username: "venkatesh_thalluri",
			Email:    "venky@example.com",
			Name:     "Venkatesh Thalluri",
			State:    "active",
		},
		WebURL: "http://example.com/example/example/snippets/1",
		RawURL: "http://example.com/example/example/snippets/1/raw",
	}}

	ss, resp, err := client.ProjectSnippets.ListSnippets(1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ss)

	ss, resp, err = client.ProjectSnippets.ListSnippets(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ss)

	ss, resp, err = client.ProjectSnippets.ListSnippets(1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ss)

	ss, resp, err = client.ProjectSnippets.ListSnippets(2, nil, nil)
	require.Error(t, err)
	require.Nil(t, ss)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectSnippetsService_GetSnippet(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/snippets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
			  "id": 1,
			  "title": "test",
			  "file_name": "add.rb",
			  "description": "Ruby test snippet",
			  "author": {
				"id": 1,
				"username": "venkatesh_thalluri",
				"email": "venky@example.com",
				"name": "Venkatesh Thalluri",
				"state": "active"
			  },
			  "project_id": 1,
			  "web_url": "http://example.com/example/example/snippets/1",
			  "raw_url": "http://example.com/example/example/snippets/1/raw"
			}
		`)
	})

	want := &Snippet{
		ID:          1,
		Title:       "test",
		FileName:    "add.rb",
		Description: "Ruby test snippet",
		Author: struct {
			ID        int        `json:"id"`
			Username  string     `json:"username"`
			Email     string     `json:"email"`
			Name      string     `json:"name"`
			State     string     `json:"state"`
			CreatedAt *time.Time `json:"created_at"`
		}{
			ID:       1,
			Username: "venkatesh_thalluri",
			Email:    "venky@example.com",
			Name:     "Venkatesh Thalluri",
			State:    "active",
		},
		WebURL: "http://example.com/example/example/snippets/1",
		RawURL: "http://example.com/example/example/snippets/1/raw",
	}

	s, resp, err := client.ProjectSnippets.GetSnippet(1, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, s)

	s, resp, err = client.ProjectSnippets.GetSnippet(1.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, s)

	s, resp, err = client.ProjectSnippets.GetSnippet(1, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, s)

	s, resp, err = client.ProjectSnippets.GetSnippet(2, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, s)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
