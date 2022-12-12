package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResourceStateEventsService_ListIssueStateEvents(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/issues/11/resource_state_events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
		  {
			"id": 142,
			"user": {
			  "id": 1,
			  "name": "Administrator",
			  "username": "root",
			  "state": "active",
			  "avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "http://gitlab.example.com/root"
			},
			"resource_type": "Issue",
			"resource_id": 11,
			"state": "opened"
		  }
		]`)
	})

	opt := &ListStateEventsOptions{ListOptions{Page: 1, PerPage: 10}}

	ses, _, err := client.ResourceStateEvents.ListIssueStateEvents(5, 11, opt)
	require.NoError(t, err)

	want := []*StateEvent{{
		ID: 142,
		User: &BasicUser{
			ID:        1,
			Username:  "root",
			Name:      "Administrator",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/root",
		},
		ResourceType: "Issue",
		ResourceID:   11,
		State:        "opened",
	}}
	require.Equal(t, want, ses)
}

func TestResourceStateEventsService_GetIssueStateEvent(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/issues/11/resource_state_events/143", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
			  "id": 143,
			  "user": {
				"id": 1,
				"name": "Administrator",
				"username": "root",
				"state": "active",
				"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				"web_url": "http://gitlab.example.com/root"
			  },
			  "resource_type": "Issue",
			  "resource_id": 11,
			  "state": "closed"
			}`,
		)
	})

	se, _, err := client.ResourceStateEvents.GetIssueStateEvent(5, 11, 143)
	require.NoError(t, err)

	want := &StateEvent{
		ID: 143,
		User: &BasicUser{
			ID:        1,
			Username:  "root",
			Name:      "Administrator",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/root",
		},
		ResourceType: "Issue",
		ResourceID:   11,
		State:        "closed",
	}
	require.Equal(t, want, se)
}

func TestResourceStateEventsService_ListMergeStateEvents(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/merge_requests/11/resource_state_events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
		  {
			"id": 142,
			"user": {
			  "id": 1,
			  "name": "Administrator",
			  "username": "root",
			  "state": "active",
			  "avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "http://gitlab.example.com/root"
			},
			"resource_type": "MergeRequest",
			"resource_id": 11,
			"state": "opened"
		  }]`,
		)
	})

	opt := &ListStateEventsOptions{ListOptions{Page: 1, PerPage: 10}}

	ses, _, err := client.ResourceStateEvents.ListMergeStateEvents(5, 11, opt)
	require.NoError(t, err)

	want := []*StateEvent{{
		ID: 142,
		User: &BasicUser{
			ID:        1,
			Username:  "root",
			Name:      "Administrator",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/root",
		},
		ResourceType: "MergeRequest",
		ResourceID:   11,
		State:        "opened",
	}}
	require.Equal(t, want, ses)
}

func TestResourceStateEventsService_GetMergeRequestStateEvent(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/merge_requests/11/resource_state_events/120", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
		  "id": 120,
		  "user": {
			"id": 1,
			"name": "Administrator",
			"username": "root",
			"state": "active",
			"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			"web_url": "http://gitlab.example.com/root"
		  },
		  "resource_type": "MergeRequest",
		  "resource_id": 11,
		  "state": "closed"
		}`)
	})

	se, _, err := client.ResourceStateEvents.GetMergeRequestStateEvent(5, 11, 120)
	require.NoError(t, err)

	want := &StateEvent{
		ID: 120,
		User: &BasicUser{
			ID:        1,
			Username:  "root",
			Name:      "Administrator",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/root",
		},
		ResourceType: "MergeRequest",
		ResourceID:   11,
		State:        "closed",
	}
	require.Equal(t, want, se)
}
