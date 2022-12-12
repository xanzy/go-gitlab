package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAwardEmojiService_ListMergeRequestAwardEmoji(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/80/award_emoji", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 4,
				"name": "1234",
				"user": {
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "id": 1,
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				  "web_url": "http://gitlab.example.com/venky333"
				},
				"awardable_id": 80,
				"awardable_type": "Merge request"
			  }
			]
		`)
	})

	want := []*AwardEmoji{{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Merge request",
	}}

	aes, resp, err := client.AwardEmoji.ListMergeRequestAwardEmoji(1, 80, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, aes)

	aes, resp, err = client.AwardEmoji.ListMergeRequestAwardEmoji(1.01, 80, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AwardEmoji.ListMergeRequestAwardEmoji(1, 80, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AwardEmoji.ListMergeRequestAwardEmoji(3, 80, nil)
	require.Error(t, err)
	require.Nil(t, aes)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_ListIssueAwardEmoji(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/80/award_emoji", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 4,
				"name": "1234",
				"user": {
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "id": 1,
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				  "web_url": "http://gitlab.example.com/venky333"
				},
				"awardable_id": 80,
				"awardable_type": "Issue"
			  }
			]
		`)
	})

	want := []*AwardEmoji{{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Issue",
	}}

	aes, resp, err := client.AwardEmoji.ListIssueAwardEmoji(1, 80, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, aes)

	aes, resp, err = client.AwardEmoji.ListIssueAwardEmoji(1.01, 80, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AwardEmoji.ListIssueAwardEmoji(1, 80, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AwardEmoji.ListIssueAwardEmoji(3, 80, nil)
	require.Error(t, err)
	require.Nil(t, aes)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_ListSnippetAwardEmoji(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/snippets/80/award_emoji", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 4,
				"name": "1234",
				"user": {
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "id": 1,
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				  "web_url": "http://gitlab.example.com/venky333"
				},
				"awardable_id": 80,
				"awardable_type": "Snippet"
			  }
			]
		`)
	})

	want := []*AwardEmoji{{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Snippet",
	}}

	aes, resp, err := client.AwardEmoji.ListSnippetAwardEmoji(1, 80, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, aes)

	aes, resp, err = client.AwardEmoji.ListSnippetAwardEmoji(1.01, 80, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AwardEmoji.ListSnippetAwardEmoji(1, 80, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AwardEmoji.ListSnippetAwardEmoji(3, 80, nil)
	require.Error(t, err)
	require.Nil(t, aes)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_GetMergeRequestAwardEmoji(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/80/award_emoji/4", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": 4,
			"name": "1234",
			"user": {
			  "name": "Venkatesh Thalluri",
			  "username": "venky333",
			  "id": 1,
			  "state": "active",
			  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "http://gitlab.example.com/venky333"
			},
			"awardable_id": 80,
			"awardable_type": "Merge request"
		  }
		`)
	})

	want := &AwardEmoji{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Merge request",
	}

	ae, resp, err := client.AwardEmoji.GetMergeRequestAwardEmoji(1, 80, 4, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ae)

	ae, resp, err = client.AwardEmoji.GetMergeRequestAwardEmoji(1.01, 80, 4, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.GetMergeRequestAwardEmoji(1, 80, 4, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.GetMergeRequestAwardEmoji(3, 80, 4, nil)
	require.Error(t, err)
	require.Nil(t, ae)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_GetIssueAwardEmoji(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/80/award_emoji/4", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": 4,
			"name": "1234",
			"user": {
			  "name": "Venkatesh Thalluri",
			  "username": "venky333",
			  "id": 1,
			  "state": "active",
			  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "http://gitlab.example.com/venky333"
			},
			"awardable_id": 80,
			"awardable_type": "Issue"
		  }
		`)
	})

	want := &AwardEmoji{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Issue",
	}

	ae, resp, err := client.AwardEmoji.GetIssueAwardEmoji(1, 80, 4, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ae)

	ae, resp, err = client.AwardEmoji.GetIssueAwardEmoji(1.01, 80, 4, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.GetIssueAwardEmoji(1, 80, 4, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.GetIssueAwardEmoji(3, 80, 4, nil)
	require.Error(t, err)
	require.Nil(t, ae)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_GetSnippetAwardEmoji(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/snippets/80/award_emoji/4", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": 4,
			"name": "1234",
			"user": {
			  "name": "Venkatesh Thalluri",
			  "username": "venky333",
			  "id": 1,
			  "state": "active",
			  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "http://gitlab.example.com/venky333"
			},
			"awardable_id": 80,
			"awardable_type": "Snippet"
		  }
		`)
	})

	want := &AwardEmoji{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Snippet",
	}

	ae, resp, err := client.AwardEmoji.GetSnippetAwardEmoji(1, 80, 4, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ae)

	ae, resp, err = client.AwardEmoji.GetSnippetAwardEmoji(1.01, 80, 4, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.GetSnippetAwardEmoji(1, 80, 4, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.GetSnippetAwardEmoji(3, 80, 4, nil)
	require.Error(t, err)
	require.Nil(t, ae)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_CreateMergeRequestAwardEmoji(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/80/award_emoji", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": 4,
			"name": "1234",
			"user": {
			  "name": "Venkatesh Thalluri",
			  "username": "venky333",
			  "id": 1,
			  "state": "active",
			  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "http://gitlab.example.com/venky333"
			},
			"awardable_id": 80,
			"awardable_type": "Merge request"
		  }
		`)
	})

	want := &AwardEmoji{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Merge request",
	}

	ae, resp, err := client.AwardEmoji.CreateMergeRequestAwardEmoji(1, 80, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ae)

	ae, resp, err = client.AwardEmoji.CreateMergeRequestAwardEmoji(1.01, 80, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.CreateMergeRequestAwardEmoji(1, 80, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.CreateMergeRequestAwardEmoji(3, 80, nil)
	require.Error(t, err)
	require.Nil(t, ae)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_CreateIssueAwardEmoji(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/80/award_emoji", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": 4,
			"name": "1234",
			"user": {
			  "name": "Venkatesh Thalluri",
			  "username": "venky333",
			  "id": 1,
			  "state": "active",
			  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "http://gitlab.example.com/venky333"
			},
			"awardable_id": 80,
			"awardable_type": "Issue"
		  }
		`)
	})

	want := &AwardEmoji{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Issue",
	}

	ae, resp, err := client.AwardEmoji.CreateIssueAwardEmoji(1, 80, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ae)

	ae, resp, err = client.AwardEmoji.CreateIssueAwardEmoji(1.01, 80, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.CreateIssueAwardEmoji(1, 80, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.CreateIssueAwardEmoji(3, 80, nil)
	require.Error(t, err)
	require.Nil(t, ae)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_CreateSnippetAwardEmoji(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/snippets/80/award_emoji", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": 4,
			"name": "1234",
			"user": {
			  "name": "Venkatesh Thalluri",
			  "username": "venky333",
			  "id": 1,
			  "state": "active",
			  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "http://gitlab.example.com/venky333"
			},
			"awardable_id": 80,
			"awardable_type": "Snippet"
		  }
		`)
	})

	want := &AwardEmoji{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Snippet",
	}

	ae, resp, err := client.AwardEmoji.CreateSnippetAwardEmoji(1, 80, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ae)

	ae, resp, err = client.AwardEmoji.CreateSnippetAwardEmoji(1.01, 80, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.CreateSnippetAwardEmoji(1, 80, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.CreateSnippetAwardEmoji(3, 80, nil)
	require.Error(t, err)
	require.Nil(t, ae)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_DeleteMergeRequestAwardEmoji(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/80/award_emoji/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.AwardEmoji.DeleteMergeRequestAwardEmoji(1, 80, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.AwardEmoji.DeleteMergeRequestAwardEmoji(1.01, 80, 1, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.AwardEmoji.DeleteMergeRequestAwardEmoji(1, 80, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.AwardEmoji.DeleteMergeRequestAwardEmoji(3, 80, 1, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_DeleteIssueAwardEmoji(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/80/award_emoji/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.AwardEmoji.DeleteIssueAwardEmoji(1, 80, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.AwardEmoji.DeleteIssueAwardEmoji(1.01, 80, 1, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.AwardEmoji.DeleteIssueAwardEmoji(1, 80, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.AwardEmoji.DeleteIssueAwardEmoji(3, 80, 1, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_DeleteSnippetAwardEmoji(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/snippets/80/award_emoji/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.AwardEmoji.DeleteSnippetAwardEmoji(1, 80, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.AwardEmoji.DeleteSnippetAwardEmoji(1.01, 80, 1, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.AwardEmoji.DeleteSnippetAwardEmoji(1, 80, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.AwardEmoji.DeleteSnippetAwardEmoji(3, 80, 1, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_ListMergeRequestAwardEmojiOnNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/80/notes/1/award_emoji", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 4,
				"name": "1234",
				"user": {
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "id": 1,
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				  "web_url": "http://gitlab.example.com/venky333"
				},
				"awardable_id": 80,
				"awardable_type": "Merge request"
			  }
			]
		`)
	})

	want := []*AwardEmoji{{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Merge request",
	}}

	aes, resp, err := client.AwardEmoji.ListMergeRequestAwardEmojiOnNote(1, 80, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, aes)

	aes, resp, err = client.AwardEmoji.ListMergeRequestAwardEmojiOnNote(1.01, 80, 1, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AwardEmoji.ListMergeRequestAwardEmojiOnNote(1, 80, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AwardEmoji.ListMergeRequestAwardEmojiOnNote(3, 80, 1, nil)
	require.Error(t, err)
	require.Nil(t, aes)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_ListIssuesAwardEmojiOnNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/80/notes/1/award_emoji", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 4,
				"name": "1234",
				"user": {
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "id": 1,
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				  "web_url": "http://gitlab.example.com/venky333"
				},
				"awardable_id": 80,
				"awardable_type": "Issue"
			  }
			]
		`)
	})

	want := []*AwardEmoji{{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Issue",
	}}

	aes, resp, err := client.AwardEmoji.ListIssuesAwardEmojiOnNote(1, 80, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, aes)

	aes, resp, err = client.AwardEmoji.ListIssuesAwardEmojiOnNote(1.01, 80, 1, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AwardEmoji.ListIssuesAwardEmojiOnNote(1, 80, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AwardEmoji.ListIssuesAwardEmojiOnNote(3, 80, 1, nil)
	require.Error(t, err)
	require.Nil(t, aes)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_ListSnippetAwardEmojiOnNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/snippets/80/notes/1/award_emoji", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 4,
				"name": "1234",
				"user": {
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "id": 1,
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				  "web_url": "http://gitlab.example.com/venky333"
				},
				"awardable_id": 80,
				"awardable_type": "Snippet"
			  }
			]
		`)
	})

	want := []*AwardEmoji{{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Snippet",
	}}

	aes, resp, err := client.AwardEmoji.ListSnippetAwardEmojiOnNote(1, 80, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, aes)

	aes, resp, err = client.AwardEmoji.ListSnippetAwardEmojiOnNote(1.01, 80, 1, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AwardEmoji.ListSnippetAwardEmojiOnNote(1, 80, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, aes)

	aes, resp, err = client.AwardEmoji.ListSnippetAwardEmojiOnNote(3, 80, 1, nil)
	require.Error(t, err)
	require.Nil(t, aes)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_GetMergeRequestAwardEmojiOnNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/80/notes/1/award_emoji/4", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": 4,
			"name": "1234",
			"user": {
			  "name": "Venkatesh Thalluri",
			  "username": "venky333",
			  "id": 1,
			  "state": "active",
			  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "http://gitlab.example.com/venky333"
			},
			"awardable_id": 80,
			"awardable_type": "Merge request"
		  }
		`)
	})

	want := &AwardEmoji{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Merge request",
	}

	ae, resp, err := client.AwardEmoji.GetMergeRequestAwardEmojiOnNote(1, 80, 1, 4, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ae)

	ae, resp, err = client.AwardEmoji.GetMergeRequestAwardEmojiOnNote(1.01, 80, 1, 4, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.GetMergeRequestAwardEmojiOnNote(1, 80, 1, 4, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.GetMergeRequestAwardEmojiOnNote(3, 80, 1, 4, nil)
	require.Error(t, err)
	require.Nil(t, ae)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_GetIssuesAwardEmojiOnNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/80/notes/1/award_emoji/4", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": 4,
			"name": "1234",
			"user": {
			  "name": "Venkatesh Thalluri",
			  "username": "venky333",
			  "id": 1,
			  "state": "active",
			  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "http://gitlab.example.com/venky333"
			},
			"awardable_id": 80,
			"awardable_type": "Issue"
		  }
		`)
	})

	want := &AwardEmoji{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Issue",
	}

	ae, resp, err := client.AwardEmoji.GetIssuesAwardEmojiOnNote(1, 80, 1, 4, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ae)

	ae, resp, err = client.AwardEmoji.GetIssuesAwardEmojiOnNote(1.01, 80, 1, 4, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.GetIssuesAwardEmojiOnNote(1, 80, 1, 4, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.GetIssuesAwardEmojiOnNote(3, 80, 1, 4, nil)
	require.Error(t, err)
	require.Nil(t, ae)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_GetSnippetAwardEmojiOnNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/snippets/80/notes/1/award_emoji/4", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": 4,
			"name": "1234",
			"user": {
			  "name": "Venkatesh Thalluri",
			  "username": "venky333",
			  "id": 1,
			  "state": "active",
			  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "http://gitlab.example.com/venky333"
			},
			"awardable_id": 80,
			"awardable_type": "Snippet"
		  }
		`)
	})

	want := &AwardEmoji{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Snippet",
	}

	ae, resp, err := client.AwardEmoji.GetSnippetAwardEmojiOnNote(1, 80, 1, 4, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ae)

	ae, resp, err = client.AwardEmoji.GetSnippetAwardEmojiOnNote(1.01, 80, 1, 4, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.GetSnippetAwardEmojiOnNote(1, 80, 1, 4, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.GetSnippetAwardEmojiOnNote(3, 80, 1, 4, nil)
	require.Error(t, err)
	require.Nil(t, ae)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_CCreateMergeRequestAwardEmojiOnNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/80/notes/1/award_emoji", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": 4,
			"name": "1234",
			"user": {
			  "name": "Venkatesh Thalluri",
			  "username": "venky333",
			  "id": 1,
			  "state": "active",
			  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "http://gitlab.example.com/venky333"
			},
			"awardable_id": 80,
			"awardable_type": "Merge request"
		  }
		`)
	})

	want := &AwardEmoji{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Merge request",
	}

	ae, resp, err := client.AwardEmoji.CreateMergeRequestAwardEmojiOnNote(1, 80, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ae)

	ae, resp, err = client.AwardEmoji.CreateMergeRequestAwardEmojiOnNote(1.01, 80, 1, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.CreateMergeRequestAwardEmojiOnNote(1, 80, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.CreateMergeRequestAwardEmojiOnNote(3, 80, 1, nil)
	require.Error(t, err)
	require.Nil(t, ae)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_CreateIssuesAwardEmojiOnNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/80/notes/1/award_emoji", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": 4,
			"name": "1234",
			"user": {
			  "name": "Venkatesh Thalluri",
			  "username": "venky333",
			  "id": 1,
			  "state": "active",
			  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "http://gitlab.example.com/venky333"
			},
			"awardable_id": 80,
			"awardable_type": "Issue"
		  }
		`)
	})

	want := &AwardEmoji{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Issue",
	}

	ae, resp, err := client.AwardEmoji.CreateIssuesAwardEmojiOnNote(1, 80, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ae)

	ae, resp, err = client.AwardEmoji.CreateIssuesAwardEmojiOnNote(1.01, 80, 1, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.CreateIssuesAwardEmojiOnNote(1, 80, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.CreateIssuesAwardEmojiOnNote(3, 80, 1, nil)
	require.Error(t, err)
	require.Nil(t, ae)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_CreateSnippetAwardEmojiOnNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/snippets/80/notes/1/award_emoji", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": 4,
			"name": "1234",
			"user": {
			  "name": "Venkatesh Thalluri",
			  "username": "venky333",
			  "id": 1,
			  "state": "active",
			  "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "http://gitlab.example.com/venky333"
			},
			"awardable_id": 80,
			"awardable_type": "Snippet"
		  }
		`)
	})

	want := &AwardEmoji{
		ID:   4,
		Name: "1234",
		User: struct {
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
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/venky333",
		},
		CreatedAt:     nil,
		UpdatedAt:     nil,
		AwardableID:   80,
		AwardableType: "Snippet",
	}

	ae, resp, err := client.AwardEmoji.CreateSnippetAwardEmojiOnNote(1, 80, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ae)

	ae, resp, err = client.AwardEmoji.CreateSnippetAwardEmojiOnNote(1.01, 80, 1, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.CreateSnippetAwardEmojiOnNote(1, 80, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ae)

	ae, resp, err = client.AwardEmoji.CreateSnippetAwardEmojiOnNote(3, 80, 1, nil)
	require.Error(t, err)
	require.Nil(t, ae)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_DeleteMergeRequestAwardEmojiOnNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/80/notes/1/award_emoji/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.AwardEmoji.DeleteMergeRequestAwardEmojiOnNote(1, 80, 1, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.AwardEmoji.DeleteMergeRequestAwardEmojiOnNote(1.01, 80, 1, 1, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.AwardEmoji.DeleteMergeRequestAwardEmojiOnNote(1, 80, 1, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.AwardEmoji.DeleteMergeRequestAwardEmojiOnNote(3, 80, 1, 1, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_DeleteIssuesAwardEmojiOnNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/80/notes/1/award_emoji/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.AwardEmoji.DeleteIssuesAwardEmojiOnNote(1, 80, 1, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.AwardEmoji.DeleteIssuesAwardEmojiOnNote(1.01, 80, 1, 1, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.AwardEmoji.DeleteIssuesAwardEmojiOnNote(1, 80, 1, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.AwardEmoji.DeleteIssuesAwardEmojiOnNote(3, 80, 1, 1, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAwardEmojiService_DeleteSnippetAwardEmojiOnNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/snippets/80/notes/1/award_emoji/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.AwardEmoji.DeleteSnippetAwardEmojiOnNote(1, 80, 1, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.AwardEmoji.DeleteSnippetAwardEmojiOnNote(1.01, 80, 1, 1, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.AwardEmoji.DeleteSnippetAwardEmojiOnNote(1, 80, 1, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.AwardEmoji.DeleteSnippetAwardEmojiOnNote(3, 80, 1, 1, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
