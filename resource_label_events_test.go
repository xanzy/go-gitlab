package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResourceLabelEventsService_ListIssueLabelEvents(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/5/issues/11/resource_label_events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
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
				"resource_id": 253,
				"label": {
				  "id": 73,
				  "name": "a1",
				  "color": "#34495E",
				  "description": ""
				},
				"action": "add"
			  }
			]
		`)
	})

	opt := &ListLabelEventsOptions{ListOptions(struct {
		Page    int
		PerPage int
	}{Page: 1, PerPage: 10})}

	les, _, err := client.ResourceLabelEvents.ListIssueLabelEvents(5, 11, opt)
	require.NoError(t, err)

	want := []*LabelEvent{{
		ID:           142,
		Action:       "add",
		ResourceType: "Issue",
		ResourceID:   253,
		User: struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Username  string `json:"username"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Name:      "Administrator",
			Username:  "root",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/root",
		},
		Label: struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Color       string `json:"color"`
			TextColor   string `json:"text_color"`
			Description string `json:"description"`
		}{
			ID:          73,
			Name:        "a1",
			Color:       "#34495E",
			TextColor:   "",
			Description: "",
		},
	}}
	require.Equal(t, want, les)
}

func TestResourceLabelEventsService_GetIssueLabelEvent(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/5/issues/11/resource_label_events/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			  {
				"id": 1,
				"user": {
				  "id": 1,
				  "name": "Administrator",
				  "username": "root",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				  "web_url": "http://gitlab.example.com/root"
				},
				"resource_type": "Issue",
				"resource_id": 253,
				"label": {
				  "id": 73,
				  "name": "a1",
				  "color": "#34495E",
				  "description": ""
				},
				"action": "add"
			  }`,
		)
	})

	le, _, err := client.ResourceLabelEvents.GetIssueLabelEvent(5, 11, 1)
	require.NoError(t, err)

	want := &LabelEvent{
		ID:           1,
		Action:       "add",
		ResourceType: "Issue",
		ResourceID:   253,
		User: struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Username  string `json:"username"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Name:      "Administrator",
			Username:  "root",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/root",
		},
		Label: struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Color       string `json:"color"`
			TextColor   string `json:"text_color"`
			Description string `json:"description"`
		}{
			ID:          73,
			Name:        "a1",
			Color:       "#34495E",
			TextColor:   "",
			Description: "",
		},
	}
	require.Equal(t, want, le)
}

func TestResourceLabelEventsService_ListGroupEpicLabelEvents(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/1/epics/11/resource_label_events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 106,
				"user": {
				  "id": 1,
				  "name": "Administrator",
				  "username": "root",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				  "web_url": "http://gitlab.example.com/root"
				},
				"resource_type": "Epic",
				"resource_id": 33,
				"label": {
				  "id": 73,
				  "name": "a1",
				  "color": "#34495E",
				  "description": ""
				},
				"action": "add"
			  }
			]
		`)
	})

	opt := &ListLabelEventsOptions{ListOptions(struct {
		Page    int
		PerPage int
	}{Page: 1, PerPage: 10})}

	les, _, err := client.ResourceLabelEvents.ListGroupEpicLabelEvents(1, 11, opt)
	require.NoError(t, err)

	want := []*LabelEvent{{
		ID:           106,
		Action:       "add",
		ResourceType: "Epic",
		ResourceID:   33,
		User: struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Username  string `json:"username"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Name:      "Administrator",
			Username:  "root",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/root",
		},
		Label: struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Color       string `json:"color"`
			TextColor   string `json:"text_color"`
			Description string `json:"description"`
		}{
			ID:          73,
			Name:        "a1",
			Color:       "#34495E",
			TextColor:   "",
			Description: "",
		},
	}}
	require.Equal(t, want, les)
}

func TestResourceLabelEventsService_GetGroupEpicLabelEvent(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/1/epics/11/resource_label_events/107", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		{
			"id": 107,
			"user": {
			"id": 1,
				"name": "Administrator",
				"username": "root",
				"state": "active",
				"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				"web_url": "http://gitlab.example.com/root"
		},
			"resource_type": "Epic",
			"resource_id": 33,
			"label": {
			"id": 73,
				"name": "a1",
				"color": "#34495E",
				"description": ""
		},
			"action": "add"
		}
		`)
	})

	le, _, err := client.ResourceLabelEvents.GetGroupEpicLabelEvent(1, 11, 107)
	require.NoError(t, err)

	want := &LabelEvent{
		ID:           107,
		Action:       "add",
		ResourceType: "Epic",
		ResourceID:   33,
		User: struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Username  string `json:"username"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Name:      "Administrator",
			Username:  "root",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/root",
		},
		Label: struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Color       string `json:"color"`
			TextColor   string `json:"text_color"`
			Description string `json:"description"`
		}{
			ID:          73,
			Name:        "a1",
			Color:       "#34495E",
			TextColor:   "",
			Description: "",
		},
	}
	require.Equal(t, want, le)
}

func TestResourceLabelEventsService_ListMergeRequestsLabelEvents(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/5/merge_requests/11/resource_label_events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 119,
				"user": {
				  "id": 1,
				  "name": "Administrator",
				  "username": "root",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				  "web_url": "http://gitlab.example.com/root"
				},
				"resource_type": "MergeRequest",
				"resource_id": 28,
				"label": {
				  "id": 74,
				  "name": "p1",
				  "color": "#0033CC",
				  "description": ""
				},
				"action": "add"
			  }
			]
		`)
	})

	opt := &ListLabelEventsOptions{ListOptions(struct {
		Page    int
		PerPage int
	}{Page: 1, PerPage: 10})}

	les, _, err := client.ResourceLabelEvents.ListMergeRequestsLabelEvents(5, 11, opt)
	require.NoError(t, err)

	want := []*LabelEvent{{
		ID:           119,
		Action:       "add",
		ResourceType: "MergeRequest",
		ResourceID:   28,
		User: struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Username  string `json:"username"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Name:      "Administrator",
			Username:  "root",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/root",
		},
		Label: struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Color       string `json:"color"`
			TextColor   string `json:"text_color"`
			Description string `json:"description"`
		}{
			ID:          74,
			Name:        "p1",
			Color:       "#0033CC",
			TextColor:   "",
			Description: "",
		},
	}}
	require.Equal(t, want, les)
}

func TestResourceLabelEventsService_GetMergeRequestLabelEvent(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/5/merge_requests/11/resource_label_events/120", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
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
				"resource_id": 28,
				"label": {
				  "id": 74,
				  "name": "p1",
				  "color": "#0033CC",
				  "description": ""
				},
				"action": "add"
			}
		`)
	})

	le, _, err := client.ResourceLabelEvents.GetMergeRequestLabelEvent(5, 11, 120)
	require.NoError(t, err)

	want := &LabelEvent{
		ID:           120,
		Action:       "add",
		ResourceType: "MergeRequest",
		ResourceID:   28,
		User: struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Username  string `json:"username"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Name:      "Administrator",
			Username:  "root",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/root",
		},
		Label: struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Color       string `json:"color"`
			TextColor   string `json:"text_color"`
			Description string `json:"description"`
		}{
			ID:          74,
			Name:        "p1",
			Color:       "#0033CC",
			TextColor:   "",
			Description: "",
		},
	}
	require.Equal(t, want, le)
}
