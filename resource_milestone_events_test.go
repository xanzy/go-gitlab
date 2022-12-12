package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func ListIssueMilestoneEvents(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/issues/11/resource_milestone_events", func(w http.ResponseWriter, r *http.Request) {
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
			"milestone": {
			  "id": 61,
			  "iid": 9,
			  "project_id": 7,
			  "title": "v1.2",
			  "description": "Ipsum Lorem",
			  "state": "active",
			  "web_url": "http://gitlab.example.com:3000/group/project/-/milestones/9"
			},
			"action": "add"
		  }
		]`)
	})

	opt := &ListMilestoneEventsOptions{ListOptions{Page: 1, PerPage: 10}}

	mes, _, err := client.ResourceMilestoneEvents.ListIssueMilestoneEvents(5, 11, opt)
	require.NoError(t, err)

	want := []*MilestoneEvent{{
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
		Milestone: &Milestone{
			ID:          61,
			IID:         9,
			ProjectID:   7,
			Title:       "v1.2",
			Description: "Ipsum Lorem",
			State:       "active",
			WebURL:      "http://gitlab.example.com:3000/group/project/-/milestones/9",
		},
		Action: "add",
	}}
	require.Equal(t, want, mes)
}

func GetIssueMilestoneEvent(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/issues/11/resource_milestone_events/143", func(w http.ResponseWriter, r *http.Request) {
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
			  "milestone": {
			    "id": 61,
			    "iid": 9,
			    "project_id": 7,
			    "title": "v1.2",
			    "description": "Ipsum Lorem",
			    "state": "active",
			    "web_url": "http://gitlab.example.com:3000/group/project/-/milestones/9"
			  },
			  "action": "remove"
			}`,
		)
	})

	me, _, err := client.ResourceMilestoneEvents.GetIssueMilestoneEvent(5, 11, 143)
	require.NoError(t, err)

	want := &MilestoneEvent{
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
		Milestone: &Milestone{
			ID:          61,
			IID:         9,
			ProjectID:   7,
			Title:       "v1.2",
			Description: "Ipsum Lorem",
			State:       "active",
			WebURL:      "http://gitlab.example.com:3000/group/project/-/milestones/9",
		},
		Action: "remove",
	}
	require.Equal(t, want, me)
}

func ListMergeMilestoneEvents(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/merge_requests/11/resource_milestone_events", func(w http.ResponseWriter, r *http.Request) {
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
			"milestone": {
			  "id": 61,
			  "iid": 9,
			  "project_id": 7,
			  "title": "v1.2",
			  "description": "Ipsum Lorem",
			  "state": "active",
			  "web_url": "http://gitlab.example.com:3000/group/project/-/milestones/9"
			},
			"action": "add"
		  }]`,
		)
	})

	opt := &ListMilestoneEventsOptions{ListOptions{Page: 1, PerPage: 10}}

	ses, _, err := client.ResourceMilestoneEvents.ListMergeMilestoneEvents(5, 11, opt)
	require.NoError(t, err)

	want := []*MilestoneEvent{{
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
		Milestone: &Milestone{
			ID:          61,
			IID:         9,
			ProjectID:   7,
			Title:       "v1.2",
			Description: "Ipsum Lorem",
			State:       "active",
			WebURL:      "http://gitlab.example.com:3000/group/project/-/milestones/9",
		},
		Action: "add",
	}}
	require.Equal(t, want, ses)
}

func GetMergeRequestMilestoneEvent(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/merge_requests/11/resource_milestone_events/120", func(w http.ResponseWriter, r *http.Request) {
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
          "milestone": {
            "id": 61,
            "iid": 9,
            "project_id": 7,
            "title": "v1.2",
            "description": "Ipsum Lorem",
            "state": "active",
            "web_url": "http://gitlab.example.com:3000/group/project/-/milestones/9"
          },
          "action": "remove"
		}`)
	})

	me, _, err := client.ResourceMilestoneEvents.GetMergeRequestMilestoneEvent(5, 11, 120)
	require.NoError(t, err)

	want := &MilestoneEvent{
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
		Milestone: &Milestone{
			ID:          61,
			IID:         9,
			ProjectID:   7,
			Title:       "v1.2",
			Description: "Ipsum Lorem",
			State:       "active",
			WebURL:      "http://gitlab.example.com:3000/group/project/-/milestones/9",
		},
		Action: "remove",
	}
	require.Equal(t, want, me)
}
