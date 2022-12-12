package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResourceLabelEventsService_ListIssueLabelEvents(t *testing.T) {
	mux, client := setup(t)

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

	les, resp, err := client.ResourceLabelEvents.ListIssueLabelEvents(5, 11, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, les)

	les, resp, err = client.ResourceLabelEvents.ListIssueLabelEvents(1.5, 11, nil)
	require.EqualError(t, err, "invalid ID type 1.5, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, les)

	les, resp, err = client.ResourceLabelEvents.ListIssueLabelEvents(5, 11, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, les)

	les, resp, err = client.ResourceLabelEvents.ListIssueLabelEvents(6, 11, nil)
	require.Error(t, err)
	require.Nil(t, les)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestResourceLabelEventsService_GetIssueLabelEvent(t *testing.T) {
	mux, client := setup(t)

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

	le, resp, err := client.ResourceLabelEvents.GetIssueLabelEvent(5, 11, 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, le)

	le, resp, err = client.ResourceLabelEvents.GetIssueLabelEvent(1.5, 11, 1)
	require.EqualError(t, err, "invalid ID type 1.5, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, le)

	le, resp, err = client.ResourceLabelEvents.GetIssueLabelEvent(5, 11, 1, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, le)

	le, resp, err = client.ResourceLabelEvents.GetIssueLabelEvent(6, 11, 1)
	require.Error(t, err)
	require.Nil(t, le)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestResourceLabelEventsService_ListGroupEpicLabelEvents(t *testing.T) {
	mux, client := setup(t)

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

	les, resp, err := client.ResourceLabelEvents.ListGroupEpicLabelEvents(1, 11, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, les)

	les, resp, err = client.ResourceLabelEvents.ListGroupEpicLabelEvents(1.5, 11, nil)
	require.EqualError(t, err, "invalid ID type 1.5, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, les)

	les, resp, err = client.ResourceLabelEvents.ListGroupEpicLabelEvents(1, 11, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, les)

	les, resp, err = client.ResourceLabelEvents.ListGroupEpicLabelEvents(6, 11, nil)
	require.Error(t, err)
	require.Nil(t, les)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestResourceLabelEventsService_GetGroupEpicLabelEvent(t *testing.T) {
	mux, client := setup(t)

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

	le, resp, err := client.ResourceLabelEvents.GetGroupEpicLabelEvent(1, 11, 107)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, le)

	le, resp, err = client.ResourceLabelEvents.GetGroupEpicLabelEvent(1.5, 11, 107)
	require.EqualError(t, err, "invalid ID type 1.5, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, le)

	le, resp, err = client.ResourceLabelEvents.GetGroupEpicLabelEvent(1, 11, 107, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, le)

	le, resp, err = client.ResourceLabelEvents.GetGroupEpicLabelEvent(6, 11, 107)
	require.Error(t, err)
	require.Nil(t, le)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestResourceLabelEventsService_ListMergeRequestsLabelEvents(t *testing.T) {
	mux, client := setup(t)

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

	les, resp, err := client.ResourceLabelEvents.ListMergeRequestsLabelEvents(5, 11, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, les)

	les, resp, err = client.ResourceLabelEvents.ListMergeRequestsLabelEvents(1.5, 11, nil)
	require.EqualError(t, err, "invalid ID type 1.5, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, les)

	les, resp, err = client.ResourceLabelEvents.ListMergeRequestsLabelEvents(5, 11, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, les)

	les, resp, err = client.ResourceLabelEvents.ListMergeRequestsLabelEvents(6, 11, nil)
	require.Error(t, err)
	require.Nil(t, les)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestResourceLabelEventsService_GetMergeRequestLabelEvent(t *testing.T) {
	mux, client := setup(t)

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

	le, resp, err := client.ResourceLabelEvents.GetMergeRequestLabelEvent(5, 11, 120)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, le)

	le, resp, err = client.ResourceLabelEvents.GetMergeRequestLabelEvent(1.5, 11, 120)
	require.EqualError(t, err, "invalid ID type 1.5, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, le)

	le, resp, err = client.ResourceLabelEvents.GetMergeRequestLabelEvent(5, 11, 120, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, le)

	le, resp, err = client.ResourceLabelEvents.GetMergeRequestLabelEvent(6, 11, 120)
	require.Error(t, err)
	require.Nil(t, le)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
