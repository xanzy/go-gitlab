package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMilestonesService_ListMilestones(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/5/milestones", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 12,
				"iid": 3,
				"project_id": 16,
				"title": "10.0",
				"description": "Version",
				"state": "active",
				"expired": false
			  }
			]
		`)
	})

	want := []*Milestone{{
		ID:          12,
		IID:         3,
		ProjectID:   16,
		Title:       "10.0",
		Description: "Version",
		State:       "active",
		WebURL:      "",
		Expired:     Bool(false),
	}}

	ms, resp, err := client.Milestones.ListMilestones(5, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ms)

	ms, resp, err = client.Milestones.ListMilestones(5.01, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ms)

	ms, resp, err = client.Milestones.ListMilestones(5, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ms)

	ms, resp, err = client.Milestones.ListMilestones(3, nil)
	require.Error(t, err)
	require.Nil(t, ms)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestMilestonesService_GetMilestone(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/5/milestones/12", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": 12,
			"iid": 3,
			"project_id": 16,
			"title": "10.0",
			"description": "Version",
			"state": "active",
			"expired": false
		  }
		`)
	})

	want := &Milestone{
		ID:          12,
		IID:         3,
		ProjectID:   16,
		Title:       "10.0",
		Description: "Version",
		State:       "active",
		WebURL:      "",
		Expired:     Bool(false),
	}

	ms, resp, err := client.Milestones.GetMilestone(5, 12, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ms)

	ms, resp, err = client.Milestones.GetMilestone(5.01, 12, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ms)

	ms, resp, err = client.Milestones.GetMilestone(5, 12, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ms)

	ms, resp, err = client.Milestones.GetMilestone(3, 12, nil)
	require.Error(t, err)
	require.Nil(t, ms)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
