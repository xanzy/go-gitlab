package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSystemHooksService_ListHooks(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/hooks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
		[
			{
			  "id":1,
			  "url":"https://gitlab.example.com/hook",
			  "created_at":"2016-10-31T12:32:15.192Z",
			  "push_events":true,
			  "tag_push_events":false,
			  "merge_requests_events": true,
			  "repository_update_events": true,
			  "enable_ssl_verification":true
			}
		]`)
	})

	hooks, _, err := client.SystemHooks.ListHooks()
	require.NoError(t, err)

	createdAt := time.Date(2016, 10, 31, 12, 32, 15, 192000000, time.UTC)
	want := []*Hook{{
		ID:                     1,
		URL:                    "https://gitlab.example.com/hook",
		CreatedAt:              &createdAt,
		PushEvents:             true,
		TagPushEvents:          false,
		MergeRequestsEvents:    true,
		RepositoryUpdateEvents: true,
		EnableSSLVerification:  true,
	}}
	require.Equal(t, want, hooks)
}

func TestSystemHooksService_GetHook(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
		{
			"id":1,
			"url":"https://gitlab.example.com/hook",
			"created_at":"2016-10-31T12:32:15.192Z",
			"push_events":true,
			"tag_push_events":false,
			"merge_requests_events": true,
			"repository_update_events": true,
			"enable_ssl_verification":true
		}`)
	})

	hooks, _, err := client.SystemHooks.GetHook(1)
	require.NoError(t, err)

	createdAt := time.Date(2016, 10, 31, 12, 32, 15, 192000000, time.UTC)
	want := &Hook{
		ID:                     1,
		URL:                    "https://gitlab.example.com/hook",
		CreatedAt:              &createdAt,
		PushEvents:             true,
		TagPushEvents:          false,
		MergeRequestsEvents:    true,
		RepositoryUpdateEvents: true,
		EnableSSLVerification:  true,
	}
	require.Equal(t, want, hooks)
}

func TestSystemHooksService_AddHook(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/hooks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id": 1, "url": "https://gitlab.example.com/hook"}`)
	})

	opt := &AddHookOptions{
		URL: String("https://gitlab.example.com/hook"),
	}

	hook, _, err := client.SystemHooks.AddHook(opt)
	require.NoError(t, err)

	want := &Hook{ID: 1, URL: "https://gitlab.example.com/hook", CreatedAt: (*time.Time)(nil)}
	require.Equal(t, want, hook)
}

func TestSystemHooksService_TestHook(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"project_id" : 1,"owner_email" : "example@gitlabhq.com","owner_name" : "Someone",
				"name" : "Ruby","path" : "ruby","event_name" : "project_create"}`)
	})

	hook, _, err := client.SystemHooks.TestHook(1)
	require.NoError(t, err)

	want := &HookEvent{
		EventName:  "project_create",
		Name:       "Ruby",
		Path:       "ruby",
		ProjectID:  1,
		OwnerName:  "Someone",
		OwnerEmail: "example@gitlabhq.com",
	}
	require.Equal(t, want, hook)
}

func TestSystemHooksService_DeleteHook(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.SystemHooks.DeleteHook(1)
	require.NoError(t, err)
}
