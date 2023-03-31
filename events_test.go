package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUsersService_ListUserContributionEvents(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/users/1/events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 3,
				"title": null,
				"project_id": 15,
				"action_name": "closed",
				"target_id": 830,
				"target_type": "Issue",
				"author_id": 1,
				"target_title": "Public project search field",
				"author": {
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "id": 1,
				  "state": "active",
				  "avatar_url": "http://localhost:3000/uploads/user/avatar/1/fox_avatar.png",
				  "web_url": "http://localhost:3000/venky333"
				},
				"author_username": "venky333"
			  }
			]
		`)
	})

	want := []*ContributionEvent{
		{
			ID:          3,
			Title:       "",
			ProjectID:   15,
			ActionName:  "closed",
			TargetID:    830,
			TargetIID:   0,
			TargetType:  "Issue",
			AuthorID:    1,
			TargetTitle: "Public project search field",
			Note:        nil,
			Author: struct {
				Name      string `json:"name"`
				Username  string `json:"username"`
				ID        int    `json:"id"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				Name:      "Venkatesh Thalluri",
				Username:  "venky333",
				ID:        1,
				State:     "active",
				AvatarURL: "http://localhost:3000/uploads/user/avatar/1/fox_avatar.png",
				WebURL:    "http://localhost:3000/venky333",
			},
			AuthorUsername: "venky333",
		},
	}

	ces, resp, err := client.Users.ListUserContributionEvents(1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ces)

	ces, resp, err = client.Users.ListUserContributionEvents(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ces)

	ces, resp, err = client.Users.ListUserContributionEvents(1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ces)

	ces, resp, err = client.Users.ListUserContributionEvents(3, nil, nil)
	require.Error(t, err)
	require.Nil(t, ces)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestEventsService_ListCurrentUserContributionEvents(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 1,
				"title":null,
				"project_id":1,
				"action_name":"opened",
				"target_id":160,
				"target_type":"Issue",
				"author_id":25,
				"target_title":"Qui natus eos odio tempore et quaerat consequuntur ducimus cupiditate quis.",
				"author":{
				  "name":"Venkatesh Thalluri",
				  "username":"venky333",
				  "id":25,
				  "state":"active",
				  "avatar_url":"http://www.gravatar.com/avatar/97d6d9441ff85fdc730e02a6068d267b?s=80u0026d=identicon",
				  "web_url":"https://gitlab.example.com/venky333"
				},
				"author_username":"venky333"
			  }
			]
		`)
	})

	want := []*ContributionEvent{
		{
			ID:          1,
			Title:       "",
			ProjectID:   1,
			ActionName:  "opened",
			TargetID:    160,
			TargetIID:   0,
			TargetType:  "Issue",
			AuthorID:    25,
			TargetTitle: "Qui natus eos odio tempore et quaerat consequuntur ducimus cupiditate quis.",
			Note:        nil,
			Author: struct {
				Name      string `json:"name"`
				Username  string `json:"username"`
				ID        int    `json:"id"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				Name:      "Venkatesh Thalluri",
				Username:  "venky333",
				ID:        25,
				State:     "active",
				AvatarURL: "http://www.gravatar.com/avatar/97d6d9441ff85fdc730e02a6068d267b?s=80u0026d=identicon",
				WebURL:    "https://gitlab.example.com/venky333",
			},
			AuthorUsername: "venky333",
		},
	}

	ces, resp, err := client.Events.ListCurrentUserContributionEvents(nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ces)

	ces, resp, err = client.Events.ListCurrentUserContributionEvents(nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ces)
}

func TestEventsService_ListCurrentUserContributionEvents_StatusNotFound(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusNotFound)
	})

	ces, resp, err := client.Events.ListCurrentUserContributionEvents(nil, nil)
	require.Error(t, err)
	require.Nil(t, ces)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestEventsService_ListProjectVisibleEvents(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/15/events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 3,
				"title": null,
				"project_id": 15,
				"action_name": "closed",
				"target_id": 830,
				"target_type": "Issue",
				"author_id": 1,
				"target_title": "Public project search field",
				"author": {
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "id": 1,
				  "state": "active",
				  "avatar_url": "http://localhost:3000/uploads/user/avatar/1/fox_avatar.png",
				  "web_url": "http://localhost:3000/venky333"
				},
				"author_username": "venky333"
			},
			{
				"id": 4,
				"title": null,
				"project_id": 15,
				"action_name": "pushed",
				"target_id": null,
				"target_type": null,
				"author_id": 1,
				"author": {
				"name": "Dmitriy Zaporozhets",
				"username": "root",
				"id": 1,
				"state": "active",
				"avatar_url": "http://localhost:3000/uploads/user/avatar/1/fox_avatar.png",
				"web_url": "http://localhost:3000/root"
				},
				"author_username": "john",
				"push_data": {
				"commit_count": 1,
				"action": "pushed",
				"ref_type": "branch",
				"commit_from": "50d4420237a9de7be1304607147aec22e4a14af7",
				"commit_to": "c5feabde2d8cd023215af4d2ceeb7a64839fc428",
				"ref": "master",
				"commit_title": "Add simple search to projects in public area"
				},
				"target_title": null
			  }
			]
		`)
	})

	want := []*ProjectEvent{
		{
			ID:          3,
			Title:       "",
			ProjectID:   15,
			ActionName:  "closed",
			TargetID:    830,
			TargetIID:   0,
			TargetType:  "Issue",
			AuthorID:    1,
			TargetTitle: "Public project search field",
			Author: struct {
				Name      string `json:"name"`
				Username  string `json:"username"`
				ID        int    `json:"id"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				Name:      "Venkatesh Thalluri",
				Username:  "venky333",
				ID:        1,
				State:     "active",
				AvatarURL: "http://localhost:3000/uploads/user/avatar/1/fox_avatar.png",
				WebURL:    "http://localhost:3000/venky333",
			},
			AuthorUsername: "venky333",
		},
		// example from https://docs.gitlab.com/ee/api/events.html#get-user-contribution-events
		{
			ID:          4,
			Title:       "",
			ProjectID:   15,
			ActionName:  "pushed",
			TargetID:    0,
			TargetIID:   0,
			TargetType:  "",
			AuthorID:    1,
			TargetTitle: "",
			CreatedAt:   "",
			Author: struct {
				Name      string `json:"name"`
				Username  string `json:"username"`
				ID        int    `json:"id"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				Name:      "Dmitriy Zaporozhets",
				Username:  "root",
				ID:        1,
				State:     "active",
				AvatarURL: "http://localhost:3000/uploads/user/avatar/1/fox_avatar.png",
				WebURL:    "http://localhost:3000/root",
			},
			AuthorUsername: "john",
			PushData: struct {
				CommitCount int    `json:"commit_count"`
				Action      string `json:"action"`
				RefType     string `json:"ref_type"`
				CommitFrom  string `json:"commit_from"`
				CommitTo    string `json:"commit_to"`
				Ref         string `json:"ref"`
				CommitTitle string `json:"commit_title"`
			}{
				CommitCount: 1,
				Action:      "pushed",
				RefType:     "branch",
				CommitFrom:  "50d4420237a9de7be1304607147aec22e4a14af7",
				CommitTo:    "c5feabde2d8cd023215af4d2ceeb7a64839fc428",
				Ref:         "master",
				CommitTitle: "Add simple search to projects in public area",
			},
		},
	}

	ces, resp, err := client.Events.ListProjectVisibleEvents(15, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ces)

	ces, resp, err = client.Events.ListProjectVisibleEvents(15.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 15.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ces)

	ces, resp, err = client.Events.ListProjectVisibleEvents(15, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ces)

	ces, resp, err = client.Events.ListProjectVisibleEvents(3, nil, nil)
	require.Error(t, err)
	require.Nil(t, ces)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
