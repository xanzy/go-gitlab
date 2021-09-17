package gitlab

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
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
	h, _, err := client.SystemHooks.AddHook(opt)

	require.NoError(t, err)
	want := &Hook{ID: 1, URL: "https://gitlab.example.com/hook", CreatedAt: (*time.Time)(nil)}
	require.Equal(t, want, h)
}

func TestSystemHooksService_TestHook(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"project_id" : 1,"owner_email" : "example@gitlabhq.com","owner_name" : "Someone",
				"name" : "Ruby","path" : "ruby","event_name" : "project_create"}`)
	})

	he, _, err := client.SystemHooks.TestHook(1)

	require.NoError(t, err)
	want := &HookEvent{
		EventName:  "project_create",
		Name:       "Ruby",
		Path:       "ruby",
		ProjectID:  1,
		OwnerName:  "Someone",
		OwnerEmail: "example@gitlabhq.com",
	}
	require.Equal(t, want, he)
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

func TestHook_String(t *testing.T) {
	h := &Hook{
		ID:        1,
		URL:       "https://github.com/",
		CreatedAt: nil,
	}
	hs := h.String()
	want := "gitlab.Hook{ID:1, URL:\"https://github.com/\"}"
	require.Equal(t,want , hs)
}

func TestHookEvent_String(t *testing.T) {
	he := &HookEvent{
		EventName:  "event",
		Name:       "name",
		Path:       "path",
		ProjectID:  1,
		OwnerName:  "owner",
		OwnerEmail: "owner_email",
	}

	hes := he.String()
	want := "gitlab.HookEvent{EventName:\"event\", Name:\"name\", Path:\"path\", ProjectID:1, OwnerName:\"owner\", OwnerEmail:\"owner_email\"}"
	require.Equal(t, want, hes)
}
