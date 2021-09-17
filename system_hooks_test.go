package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSystemHooksService_ListHooks(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/hooks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1,"url":"https://gitlab.example.com/hook"}]`)
	})

	hooks, _, err := client.SystemHooks.ListHooks()
	require.NoError(t, err)

	want := []*Hook{{ID: 1, URL: "https://gitlab.example.com/hook"}}
	require.Equal(t, want, hooks)
}

func TestSystemHooksService_AddHook(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

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
	mux, server, client := setup(t)
	defer teardown(server)

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
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.SystemHooks.DeleteHook(1)
	require.NoError(t, err)
}
