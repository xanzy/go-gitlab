package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiscussionsService_ListIssueDiscussions(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/issues/11/discussions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
				"individual_note": false,
				"notes": [
				  {
					"id": 1126,
					"type": "DiscussionNote",
					"body": "discussion text",
					"attachment": null,
					"author": {
					  "id": 1,
					  "name": "Venkatesh Thalluri",
					  "username": "venky333",
					  "state": "active",
					  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
					  "web_url": "http://localhost:3000/venky333"
					},
					"system": false,
					"noteable_id": 3,
					"noteable_type": "Issue",
					"noteable_iid": null
				  }
				]
			  }
			]
		`)
	})

	want := []*Discussion{{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Issue",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}}

	ds, resp, err := client.Discussions.ListIssueDiscussions(5, 11, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ds)

	ds, resp, err = client.Discussions.ListIssueDiscussions(5.01, 11, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ds)

	ds, resp, err = client.Discussions.ListIssueDiscussions(5, 11, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ds)

	ds, resp, err = client.Discussions.ListIssueDiscussions(3, 11, nil, nil)
	require.Error(t, err)
	require.Nil(t, ds)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_GetIssueDiscussion(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/issues/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
			"individual_note": false,
			"notes": [
			  {
				"id": 1126,
				"type": "DiscussionNote",
				"body": "discussion text",
				"attachment": null,
				"author": {
				  "id": 1,
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				  "web_url": "http://localhost:3000/venky333"
				},
				"system": false,
				"noteable_id": 3,
				"noteable_type": "Issue",
				"noteable_iid": null
			  }
			]
		  }
		`)
	})

	want := &Discussion{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Issue",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}

	d, resp, err := client.Discussions.GetIssueDiscussion(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)

	d, resp, err = client.Discussions.GetIssueDiscussion(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.GetIssueDiscussion(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.GetIssueDiscussion(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.Error(t, err)
	require.Nil(t, d)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_CreateIssueDiscussion(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/issues/11/discussions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
			"individual_note": false,
			"notes": [
			  {
				"id": 1126,
				"type": "DiscussionNote",
				"body": "discussion text",
				"attachment": null,
				"author": {
				  "id": 1,
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				  "web_url": "http://localhost:3000/venky333"
				},
				"system": false,
				"noteable_id": 3,
				"noteable_type": "Issue",
				"noteable_iid": null
			  }
			]
		  }
		`)
	})

	want := &Discussion{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Issue",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}

	d, resp, err := client.Discussions.CreateIssueDiscussion(5, 11, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)

	d, resp, err = client.Discussions.CreateIssueDiscussion(5.01, 11, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.CreateIssueDiscussion(5, 11, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.CreateIssueDiscussion(3, 11, nil, nil)
	require.Error(t, err)
	require.Nil(t, d)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_AddIssueDiscussionNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/issues/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7/notes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": 302,
			"body": "closed",
			"attachment": null,
			"author": {
			  "id": 1,
			  "username": "venky333",
			  "email": "venky333@example.com",
			  "name": "venky333",
			  "state": "active"
			},
			"system": true,
			"noteable_id": 377,
			"noteable_type": "Issue",
			"noteable_iid": 377,
			"resolvable": false,
			"confidential": false
		  }
		`)
	})

	want := &Note{
		ID:         302,
		Type:       "",
		Body:       "closed",
		Attachment: "",
		Title:      "",
		FileName:   "",
		Author: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Username:  "venky333",
			Email:     "venky333@example.com",
			Name:      "venky333",
			State:     "active",
			AvatarURL: "",
			WebURL:    "",
		},
		System:       true,
		ExpiresAt:    nil,
		UpdatedAt:    nil,
		CreatedAt:    nil,
		NoteableID:   377,
		NoteableType: "Issue",
		CommitID:     "",
		Position:     nil,
		Resolvable:   false,
		Resolved:     false,
		ResolvedBy: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{},
		NoteableIID: 377,
	}

	n, resp, err := client.Discussions.AddIssueDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, n)

	n, resp, err = client.Discussions.AddIssueDiscussionNote(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.AddIssueDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.AddIssueDiscussionNote(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.Error(t, err)
	require.Nil(t, n)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_UpdateIssueDiscussionNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/issues/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7/notes/302", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
		  {
			"id": 302,
			"body": "closed",
			"attachment": null,
			"author": {
			  "id": 1,
			  "username": "venky333",
			  "email": "venky333@example.com",
			  "name": "venky333",
			  "state": "active"
			},
			"system": true,
			"noteable_id": 377,
			"noteable_type": "Issue",
			"noteable_iid": 377,
			"resolvable": false,
			"confidential": false
		  }
		`)
	})

	want := &Note{
		ID:         302,
		Type:       "",
		Body:       "closed",
		Attachment: "",
		Title:      "",
		FileName:   "",
		Author: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Username:  "venky333",
			Email:     "venky333@example.com",
			Name:      "venky333",
			State:     "active",
			AvatarURL: "",
			WebURL:    "",
		},
		System:       true,
		ExpiresAt:    nil,
		UpdatedAt:    nil,
		CreatedAt:    nil,
		NoteableID:   377,
		NoteableType: "Issue",
		CommitID:     "",
		Position:     nil,
		Resolvable:   false,
		Resolved:     false,
		ResolvedBy: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{},
		NoteableIID: 377,
	}

	n, resp, err := client.Discussions.UpdateIssueDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, n)

	n, resp, err = client.Discussions.UpdateIssueDiscussionNote(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.UpdateIssueDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.UpdateIssueDiscussionNote(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.Error(t, err)
	require.Nil(t, n)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_DeleteIssueDiscussionNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/issues/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7/notes/302", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Discussions.DeleteIssueDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.Discussions.DeleteIssueDiscussionNote(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.Discussions.DeleteIssueDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.Discussions.DeleteIssueDiscussionNote(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_ListSnippetDiscussions(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/snippets/11/discussions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
				"individual_note": false,
				"notes": [
				  {
					"id": 1126,
					"type": "DiscussionNote",
					"body": "discussion text",
					"attachment": null,
					"author": {
					  "id": 1,
					  "name": "Venkatesh Thalluri",
					  "username": "venky333",
					  "state": "active",
					  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
					  "web_url": "http://localhost:3000/venky333"
					},
					"system": false,
					"noteable_id": 3,
					"noteable_type": "Snippet",
					"noteable_iid": null
				  }
				]
			  }
			]
		`)
	})

	want := []*Discussion{{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Snippet",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}}

	ds, resp, err := client.Discussions.ListSnippetDiscussions(5, 11, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ds)

	ds, resp, err = client.Discussions.ListSnippetDiscussions(5.01, 11, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ds)

	ds, resp, err = client.Discussions.ListSnippetDiscussions(5, 11, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ds)

	ds, resp, err = client.Discussions.ListSnippetDiscussions(3, 11, nil, nil)
	require.Error(t, err)
	require.Nil(t, ds)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_GetSnippetDiscussion(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/snippets/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
			"individual_note": false,
			"notes": [
			  {
				"id": 1126,
				"type": "DiscussionNote",
				"body": "discussion text",
				"attachment": null,
				"author": {
				  "id": 1,
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				  "web_url": "http://localhost:3000/venky333"
				},
				"system": false,
				"noteable_id": 3,
				"noteable_type": "Snippet",
				"noteable_iid": null
			  }
			]
		  }
		`)
	})

	want := &Discussion{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Snippet",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}

	d, resp, err := client.Discussions.GetSnippetDiscussion(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)

	d, resp, err = client.Discussions.GetSnippetDiscussion(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.GetSnippetDiscussion(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.GetSnippetDiscussion(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.Error(t, err)
	require.Nil(t, d)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_CreateSnippetDiscussion(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/snippets/11/discussions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
			"individual_note": false,
			"notes": [
			  {
				"id": 1126,
				"type": "DiscussionNote",
				"body": "discussion text",
				"attachment": null,
				"author": {
				  "id": 1,
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				  "web_url": "http://localhost:3000/venky333"
				},
				"system": false,
				"noteable_id": 3,
				"noteable_type": "Snippet",
				"noteable_iid": null
			  }
			]
		  }
		`)
	})

	want := &Discussion{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Snippet",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}

	d, resp, err := client.Discussions.CreateSnippetDiscussion(5, 11, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)

	d, resp, err = client.Discussions.CreateSnippetDiscussion(5.01, 11, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.CreateSnippetDiscussion(5, 11, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.CreateSnippetDiscussion(3, 11, nil, nil)
	require.Error(t, err)
	require.Nil(t, d)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_AddSnippetDiscussionNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/snippets/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7/notes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": 302,
			"body": "closed",
			"attachment": null,
			"author": {
			  "id": 1,
			  "username": "venky333",
			  "email": "venky333@example.com",
			  "name": "venky333",
			  "state": "active"
			},
			"system": true,
			"noteable_id": 377,
			"noteable_type": "Snippet",
			"noteable_iid": 377,
			"resolvable": false,
			"confidential": false
		  }
		`)
	})

	want := &Note{
		ID:         302,
		Type:       "",
		Body:       "closed",
		Attachment: "",
		Title:      "",
		FileName:   "",
		Author: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Username:  "venky333",
			Email:     "venky333@example.com",
			Name:      "venky333",
			State:     "active",
			AvatarURL: "",
			WebURL:    "",
		},
		System:       true,
		ExpiresAt:    nil,
		UpdatedAt:    nil,
		CreatedAt:    nil,
		NoteableID:   377,
		NoteableType: "Snippet",
		CommitID:     "",
		Position:     nil,
		Resolvable:   false,
		Resolved:     false,
		ResolvedBy: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{},
		NoteableIID: 377,
	}

	n, resp, err := client.Discussions.AddSnippetDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, n)

	n, resp, err = client.Discussions.AddSnippetDiscussionNote(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.AddSnippetDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.AddSnippetDiscussionNote(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.Error(t, err)
	require.Nil(t, n)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_UpdateSnippetDiscussionNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/snippets/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7/notes/302", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
		  {
			"id": 302,
			"body": "closed",
			"attachment": null,
			"author": {
			  "id": 1,
			  "username": "venky333",
			  "email": "venky333@example.com",
			  "name": "venky333",
			  "state": "active"
			},
			"system": true,
			"noteable_id": 377,
			"noteable_type": "Snippet",
			"noteable_iid": 377,
			"resolvable": false,
			"confidential": false
		  }
		`)
	})

	want := &Note{
		ID:         302,
		Type:       "",
		Body:       "closed",
		Attachment: "",
		Title:      "",
		FileName:   "",
		Author: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Username:  "venky333",
			Email:     "venky333@example.com",
			Name:      "venky333",
			State:     "active",
			AvatarURL: "",
			WebURL:    "",
		},
		System:       true,
		ExpiresAt:    nil,
		UpdatedAt:    nil,
		CreatedAt:    nil,
		NoteableID:   377,
		NoteableType: "Snippet",
		CommitID:     "",
		Position:     nil,
		Resolvable:   false,
		Resolved:     false,
		ResolvedBy: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{},
		NoteableIID: 377,
	}

	n, resp, err := client.Discussions.UpdateSnippetDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, n)

	n, resp, err = client.Discussions.UpdateSnippetDiscussionNote(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.UpdateSnippetDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.UpdateSnippetDiscussionNote(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.Error(t, err)
	require.Nil(t, n)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_DeleteSnippetDiscussionNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/snippets/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7/notes/302", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Discussions.DeleteSnippetDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.Discussions.DeleteSnippetDiscussionNote(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.Discussions.DeleteSnippetDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.Discussions.DeleteSnippetDiscussionNote(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_ListGroupEpicDiscussions(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/epics/11/discussions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
				"individual_note": false,
				"notes": [
				  {
					"id": 1126,
					"type": "DiscussionNote",
					"body": "discussion text",
					"attachment": null,
					"author": {
					  "id": 1,
					  "name": "Venkatesh Thalluri",
					  "username": "venky333",
					  "state": "active",
					  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
					  "web_url": "http://localhost:3000/venky333"
					},
					"system": false,
					"noteable_id": 3,
					"noteable_type": "Epic",
					"noteable_iid": null
				  }
				]
			  }
			]
		`)
	})

	want := []*Discussion{{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Epic",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}}

	ds, resp, err := client.Discussions.ListGroupEpicDiscussions(5, 11, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ds)

	ds, resp, err = client.Discussions.ListGroupEpicDiscussions(5.01, 11, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ds)

	ds, resp, err = client.Discussions.ListGroupEpicDiscussions(5, 11, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ds)

	ds, resp, err = client.Discussions.ListGroupEpicDiscussions(3, 11, nil, nil)
	require.Error(t, err)
	require.Nil(t, ds)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_GetEpicDiscussion(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/epics/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
			"individual_note": false,
			"notes": [
			  {
				"id": 1126,
				"type": "DiscussionNote",
				"body": "discussion text",
				"attachment": null,
				"author": {
				  "id": 1,
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				  "web_url": "http://localhost:3000/venky333"
				},
				"system": false,
				"noteable_id": 3,
				"noteable_type": "Epic",
				"noteable_iid": null
			  }
			]
		  }
		`)
	})

	want := &Discussion{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Epic",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}

	d, resp, err := client.Discussions.GetEpicDiscussion(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)

	d, resp, err = client.Discussions.GetEpicDiscussion(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.GetEpicDiscussion(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.GetEpicDiscussion(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.Error(t, err)
	require.Nil(t, d)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_CreateEpicDiscussion(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/epics/11/discussions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
			"individual_note": false,
			"notes": [
			  {
				"id": 1126,
				"type": "DiscussionNote",
				"body": "discussion text",
				"attachment": null,
				"author": {
				  "id": 1,
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				  "web_url": "http://localhost:3000/venky333"
				},
				"system": false,
				"noteable_id": 3,
				"noteable_type": "Epic",
				"noteable_iid": null
			  }
			]
		  }
		`)
	})

	want := &Discussion{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Epic",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}

	d, resp, err := client.Discussions.CreateEpicDiscussion(5, 11, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)

	d, resp, err = client.Discussions.CreateEpicDiscussion(5.01, 11, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.CreateEpicDiscussion(5, 11, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.CreateEpicDiscussion(3, 11, nil, nil)
	require.Error(t, err)
	require.Nil(t, d)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_AddEpicDiscussionNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/epics/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7/notes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": 302,
			"body": "closed",
			"attachment": null,
			"author": {
			  "id": 1,
			  "username": "venky333",
			  "email": "venky333@example.com",
			  "name": "venky333",
			  "state": "active"
			},
			"system": true,
			"noteable_id": 377,
			"noteable_type": "Epic",
			"noteable_iid": 377,
			"resolvable": false,
			"confidential": false
		  }
		`)
	})

	want := &Note{
		ID:         302,
		Type:       "",
		Body:       "closed",
		Attachment: "",
		Title:      "",
		FileName:   "",
		Author: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Username:  "venky333",
			Email:     "venky333@example.com",
			Name:      "venky333",
			State:     "active",
			AvatarURL: "",
			WebURL:    "",
		},
		System:       true,
		ExpiresAt:    nil,
		UpdatedAt:    nil,
		CreatedAt:    nil,
		NoteableID:   377,
		NoteableType: "Epic",
		CommitID:     "",
		Position:     nil,
		Resolvable:   false,
		Resolved:     false,
		ResolvedBy: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{},
		NoteableIID: 377,
	}

	n, resp, err := client.Discussions.AddEpicDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, n)

	n, resp, err = client.Discussions.AddEpicDiscussionNote(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.AddEpicDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.AddEpicDiscussionNote(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.Error(t, err)
	require.Nil(t, n)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_UpdateEpicDiscussionNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/epics/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7/notes/302", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
		  {
			"id": 302,
			"body": "closed",
			"attachment": null,
			"author": {
			  "id": 1,
			  "username": "venky333",
			  "email": "venky333@example.com",
			  "name": "venky333",
			  "state": "active"
			},
			"system": true,
			"noteable_id": 377,
			"noteable_type": "Epic",
			"noteable_iid": 377,
			"resolvable": false,
			"confidential": false
		  }
		`)
	})

	want := &Note{
		ID:         302,
		Type:       "",
		Body:       "closed",
		Attachment: "",
		Title:      "",
		FileName:   "",
		Author: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Username:  "venky333",
			Email:     "venky333@example.com",
			Name:      "venky333",
			State:     "active",
			AvatarURL: "",
			WebURL:    "",
		},
		System:       true,
		ExpiresAt:    nil,
		UpdatedAt:    nil,
		CreatedAt:    nil,
		NoteableID:   377,
		NoteableType: "Epic",
		CommitID:     "",
		Position:     nil,
		Resolvable:   false,
		Resolved:     false,
		ResolvedBy: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{},
		NoteableIID: 377,
	}

	n, resp, err := client.Discussions.UpdateEpicDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, n)

	n, resp, err = client.Discussions.UpdateEpicDiscussionNote(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.UpdateEpicDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.UpdateEpicDiscussionNote(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.Error(t, err)
	require.Nil(t, n)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_DeleteEpicDiscussionNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/epics/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7/notes/302", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Discussions.DeleteEpicDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.Discussions.DeleteEpicDiscussionNote(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.Discussions.DeleteEpicDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.Discussions.DeleteEpicDiscussionNote(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_ListMergeRequestDiscussions(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/merge_requests/11/discussions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
				"individual_note": false,
				"notes": [
				  {
					"id": 1126,
					"type": "DiscussionNote",
					"body": "discussion text",
					"attachment": null,
					"author": {
					  "id": 1,
					  "name": "Venkatesh Thalluri",
					  "username": "venky333",
					  "state": "active",
					  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
					  "web_url": "http://localhost:3000/venky333"
					},
					"system": false,
					"noteable_id": 3,
					"noteable_type": "Merge request",
					"noteable_iid": null
				  }
				]
			  }
			]
		`)
	})

	want := []*Discussion{{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Merge request",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}}

	ds, resp, err := client.Discussions.ListMergeRequestDiscussions(5, 11, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ds)

	ds, resp, err = client.Discussions.ListMergeRequestDiscussions(5.01, 11, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ds)

	ds, resp, err = client.Discussions.ListMergeRequestDiscussions(5, 11, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ds)

	ds, resp, err = client.Discussions.ListMergeRequestDiscussions(3, 11, nil, nil)
	require.Error(t, err)
	require.Nil(t, ds)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_GetMergeRequestDiscussion(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/merge_requests/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
			"individual_note": false,
			"notes": [
			  {
				"id": 1126,
				"type": "DiscussionNote",
				"body": "discussion text",
				"attachment": null,
				"author": {
				  "id": 1,
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				  "web_url": "http://localhost:3000/venky333"
				},
				"system": false,
				"noteable_id": 3,
				"noteable_type": "Merge request",
				"noteable_iid": null
			  }
			]
		  }
		`)
	})

	want := &Discussion{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Merge request",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}

	d, resp, err := client.Discussions.GetMergeRequestDiscussion(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)

	d, resp, err = client.Discussions.GetMergeRequestDiscussion(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.GetMergeRequestDiscussion(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.GetMergeRequestDiscussion(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.Error(t, err)
	require.Nil(t, d)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_CreateMergeRequestDiscussion(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/merge_requests/11/discussions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
			"individual_note": false,
			"notes": [
			  {
				"id": 1126,
				"type": "DiscussionNote",
				"body": "discussion text",
				"attachment": null,
				"author": {
				  "id": 1,
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				  "web_url": "http://localhost:3000/venky333"
				},
				"system": false,
				"noteable_id": 3,
				"noteable_type": "Merge request",
				"noteable_iid": null
			  }
			]
		  }
		`)
	})

	want := &Discussion{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Merge request",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}

	d, resp, err := client.Discussions.CreateMergeRequestDiscussion(5, 11, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)

	d, resp, err = client.Discussions.CreateMergeRequestDiscussion(5.01, 11, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.CreateMergeRequestDiscussion(5, 11, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.CreateMergeRequestDiscussion(3, 11, nil, nil)
	require.Error(t, err)
	require.Nil(t, d)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_ResolveMergeRequestDiscussion(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/merge_requests/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
		  {
			"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
			"individual_note": false,
			"notes": [
			  {
				"id": 1126,
				"type": "DiscussionNote",
				"body": "discussion text",
				"attachment": null,
				"author": {
				  "id": 1,
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				  "web_url": "http://localhost:3000/venky333"
				},
				"system": false,
				"noteable_id": 3,
				"noteable_type": "Merge request",
				"noteable_iid": null
			  }
			]
		  }
		`)
	})

	want := &Discussion{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Merge request",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}

	d, resp, err := client.Discussions.ResolveMergeRequestDiscussion(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)

	d, resp, err = client.Discussions.ResolveMergeRequestDiscussion(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.ResolveMergeRequestDiscussion(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.ResolveMergeRequestDiscussion(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.Error(t, err)
	require.Nil(t, d)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_AddMergeRequestDiscussionNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/merge_requests/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7/notes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": 302,
			"body": "closed",
			"attachment": null,
			"author": {
			  "id": 1,
			  "username": "venky333",
			  "email": "venky333@example.com",
			  "name": "venky333",
			  "state": "active"
			},
			"system": true,
			"noteable_id": 377,
			"noteable_type": "Merge request",
			"noteable_iid": 377,
			"resolvable": false,
			"confidential": false
		  }
		`)
	})

	want := &Note{
		ID:         302,
		Type:       "",
		Body:       "closed",
		Attachment: "",
		Title:      "",
		FileName:   "",
		Author: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Username:  "venky333",
			Email:     "venky333@example.com",
			Name:      "venky333",
			State:     "active",
			AvatarURL: "",
			WebURL:    "",
		},
		System:       true,
		ExpiresAt:    nil,
		UpdatedAt:    nil,
		CreatedAt:    nil,
		NoteableID:   377,
		NoteableType: "Merge request",
		CommitID:     "",
		Position:     nil,
		Resolvable:   false,
		Resolved:     false,
		ResolvedBy: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{},
		NoteableIID: 377,
	}

	n, resp, err := client.Discussions.AddMergeRequestDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, n)

	n, resp, err = client.Discussions.AddMergeRequestDiscussionNote(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.AddMergeRequestDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.AddMergeRequestDiscussionNote(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.Error(t, err)
	require.Nil(t, n)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_UpdateMergeRequestDiscussionNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/merge_requests/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7/notes/302", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
		  {
			"id": 302,
			"body": "closed",
			"attachment": null,
			"author": {
			  "id": 1,
			  "username": "venky333",
			  "email": "venky333@example.com",
			  "name": "venky333",
			  "state": "active"
			},
			"system": true,
			"noteable_id": 377,
			"noteable_type": "Merge request",
			"noteable_iid": 377,
			"resolvable": false,
			"confidential": false
		  }
		`)
	})

	want := &Note{
		ID:         302,
		Type:       "",
		Body:       "closed",
		Attachment: "",
		Title:      "",
		FileName:   "",
		Author: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Username:  "venky333",
			Email:     "venky333@example.com",
			Name:      "venky333",
			State:     "active",
			AvatarURL: "",
			WebURL:    "",
		},
		System:       true,
		ExpiresAt:    nil,
		UpdatedAt:    nil,
		CreatedAt:    nil,
		NoteableID:   377,
		NoteableType: "Merge request",
		CommitID:     "",
		Position:     nil,
		Resolvable:   false,
		Resolved:     false,
		ResolvedBy: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{},
		NoteableIID: 377,
	}

	n, resp, err := client.Discussions.UpdateMergeRequestDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, n)

	n, resp, err = client.Discussions.UpdateMergeRequestDiscussionNote(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.UpdateMergeRequestDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.UpdateMergeRequestDiscussionNote(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.Error(t, err)
	require.Nil(t, n)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_DeleteMergeRequestDiscussionNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/merge_requests/11/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7/notes/302", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Discussions.DeleteMergeRequestDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.Discussions.DeleteMergeRequestDiscussionNote(5.01, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.Discussions.DeleteMergeRequestDiscussionNote(5, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.Discussions.DeleteMergeRequestDiscussionNote(3, 11, "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_ListCommitDiscussions(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/repository/commits/abc123/discussions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
				"individual_note": false,
				"notes": [
				  {
					"id": 1126,
					"type": "DiscussionNote",
					"body": "discussion text",
					"attachment": null,
					"author": {
					  "id": 1,
					  "name": "Venkatesh Thalluri",
					  "username": "venky333",
					  "state": "active",
					  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
					  "web_url": "http://localhost:3000/venky333"
					},
					"system": false,
					"noteable_id": 3,
					"noteable_type": "Commit",
					"noteable_iid": null
				  }
				]
			  }
			]
		`)
	})

	want := []*Discussion{{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Commit",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}}

	ds, resp, err := client.Discussions.ListCommitDiscussions(5, "abc123", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ds)

	ds, resp, err = client.Discussions.ListCommitDiscussions(5.01, "abc123", nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ds)

	ds, resp, err = client.Discussions.ListCommitDiscussions(5, "abc123", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ds)

	ds, resp, err = client.Discussions.ListCommitDiscussions(3, "abc123", nil, nil)
	require.Error(t, err)
	require.Nil(t, ds)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_GetCommitDiscussion(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/repository/commits/abc123/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
			"individual_note": false,
			"notes": [
			  {
				"id": 1126,
				"type": "DiscussionNote",
				"body": "discussion text",
				"attachment": null,
				"author": {
				  "id": 1,
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				  "web_url": "http://localhost:3000/venky333"
				},
				"system": false,
				"noteable_id": 3,
				"noteable_type": "Commit",
				"noteable_iid": null
			  }
			]
		  }
		`)
	})

	want := &Discussion{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Commit",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}

	d, resp, err := client.Discussions.GetCommitDiscussion(5, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)

	d, resp, err = client.Discussions.GetCommitDiscussion(5.01, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.GetCommitDiscussion(5, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.GetCommitDiscussion(3, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.Error(t, err)
	require.Nil(t, d)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_CreateCommitDiscussion(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/repository/commits/abc123/discussions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": "6a9c1750b37d513a43987b574953fceb50b03ce7",
			"individual_note": false,
			"notes": [
			  {
				"id": 1126,
				"type": "DiscussionNote",
				"body": "discussion text",
				"attachment": null,
				"author": {
				  "id": 1,
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				  "web_url": "http://localhost:3000/venky333"
				},
				"system": false,
				"noteable_id": 3,
				"noteable_type": "Commit",
				"noteable_iid": null
			  }
			]
		  }
		`)
	})

	want := &Discussion{
		ID:             "6a9c1750b37d513a43987b574953fceb50b03ce7",
		IndividualNote: false,
		Notes: []*Note{{
			ID:         1126,
			Type:       "DiscussionNote",
			Body:       "discussion text",
			Attachment: "",
			Title:      "",
			FileName:   "",
			Author: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{
				ID:        1,
				Username:  "venky333",
				Email:     "",
				Name:      "Venkatesh Thalluri",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/00afb8fb6ab07c3ee3e9c1f38777e2f4?s=80&d=identicon",
				WebURL:    "http://localhost:3000/venky333",
			},
			System:       false,
			ExpiresAt:    nil,
			UpdatedAt:    nil,
			CreatedAt:    nil,
			NoteableID:   3,
			NoteableType: "Commit",
			CommitID:     "",
			Position:     nil,
			Resolvable:   false,
			Resolved:     false,
			ResolvedBy: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				State     string `json:"state"`
				AvatarURL string `json:"avatar_url"`
				WebURL    string `json:"web_url"`
			}{},
			NoteableIID: 0,
		}},
	}

	d, resp, err := client.Discussions.CreateCommitDiscussion(5, "abc123", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)

	d, resp, err = client.Discussions.CreateCommitDiscussion(5.01, "abc123", nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.CreateCommitDiscussion(5, "abc123", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, d)

	d, resp, err = client.Discussions.CreateCommitDiscussion(3, "abc123", nil, nil)
	require.Error(t, err)
	require.Nil(t, d)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_AddCommitDiscussionNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/repository/commits/abc123/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7/notes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": 302,
			"body": "closed",
			"attachment": null,
			"author": {
			  "id": 1,
			  "username": "venky333",
			  "email": "venky333@example.com",
			  "name": "venky333",
			  "state": "active"
			},
			"system": true,
			"noteable_id": 377,
			"noteable_type": "Commit",
			"noteable_iid": 377,
			"resolvable": false,
			"confidential": false
		  }
		`)
	})

	want := &Note{
		ID:         302,
		Type:       "",
		Body:       "closed",
		Attachment: "",
		Title:      "",
		FileName:   "",
		Author: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Username:  "venky333",
			Email:     "venky333@example.com",
			Name:      "venky333",
			State:     "active",
			AvatarURL: "",
			WebURL:    "",
		},
		System:       true,
		ExpiresAt:    nil,
		UpdatedAt:    nil,
		CreatedAt:    nil,
		NoteableID:   377,
		NoteableType: "Commit",
		CommitID:     "",
		Position:     nil,
		Resolvable:   false,
		Resolved:     false,
		ResolvedBy: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{},
		NoteableIID: 377,
	}

	n, resp, err := client.Discussions.AddCommitDiscussionNote(5, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, n)

	n, resp, err = client.Discussions.AddCommitDiscussionNote(5.01, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.AddCommitDiscussionNote(5, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.AddCommitDiscussionNote(3, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", nil, nil)
	require.Error(t, err)
	require.Nil(t, n)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_UpdateCommitDiscussionNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/repository/commits/abc123/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7/notes/302", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
		  {
			"id": 302,
			"body": "closed",
			"attachment": null,
			"author": {
			  "id": 1,
			  "username": "venky333",
			  "email": "venky333@example.com",
			  "name": "venky333",
			  "state": "active"
			},
			"system": true,
			"noteable_id": 377,
			"noteable_type": "Commit",
			"noteable_iid": 377,
			"resolvable": false,
			"confidential": false
		  }
		`)
	})

	want := &Note{
		ID:         302,
		Type:       "",
		Body:       "closed",
		Attachment: "",
		Title:      "",
		FileName:   "",
		Author: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{
			ID:        1,
			Username:  "venky333",
			Email:     "venky333@example.com",
			Name:      "venky333",
			State:     "active",
			AvatarURL: "",
			WebURL:    "",
		},
		System:       true,
		ExpiresAt:    nil,
		UpdatedAt:    nil,
		CreatedAt:    nil,
		NoteableID:   377,
		NoteableType: "Commit",
		CommitID:     "",
		Position:     nil,
		Resolvable:   false,
		Resolved:     false,
		ResolvedBy: struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		}{},
		NoteableIID: 377,
	}

	n, resp, err := client.Discussions.UpdateCommitDiscussionNote(5, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, n)

	n, resp, err = client.Discussions.UpdateCommitDiscussionNote(5.01, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.UpdateCommitDiscussionNote(5, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, n)

	n, resp, err = client.Discussions.UpdateCommitDiscussionNote(3, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.Error(t, err)
	require.Nil(t, n)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDiscussionsService_DeleteCommitDiscussionNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/repository/commits/abc123/discussions/6a9c1750b37d513a43987b574953fceb50b03ce7/notes/302", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Discussions.DeleteCommitDiscussionNote(5, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.Discussions.DeleteCommitDiscussionNote(5.01, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.Discussions.DeleteCommitDiscussionNote(5, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.Discussions.DeleteCommitDiscussionNote(3, "abc123", "6a9c1750b37d513a43987b574953fceb50b03ce7", 302, nil, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
