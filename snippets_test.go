package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSnippetsService_ListSnippets(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/snippets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":42,"title":"test"}]`)
	})

	opt := &ListSnippetsOptions{Page: 1, PerPage: 10}

	ss, _, err := client.Snippets.ListSnippets(opt)
	require.NoError(t, err)

	want := []*Snippet{{ID: 42, Title: "test"}}
	require.Equal(t, want, ss)
}

func TestSnippetsService_GetSnippet(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/snippets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1, "title":"test"}`)
	})

	s, _, err := client.Snippets.GetSnippet(1)
	require.NoError(t, err)

	want := &Snippet{ID: 1, Title: "test"}
	require.Equal(t, want, s)
}

func TestSnippetsService_CreateSnippet(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/snippets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1, "title":"test"}`)
	})

	opt := &CreateSnippetOptions{
		Title:       String("test"),
		FileName:    String("file"),
		Description: String("description"),
		Content:     String("content"),
		Visibility:  Visibility(PublicVisibility),
	}

	s, _, err := client.Snippets.CreateSnippet(opt)
	require.NoError(t, err)

	want := &Snippet{ID: 1, Title: "test"}
	require.Equal(t, want, s)
}

func TestSnippetsService_UpdateSnippet(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/snippets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{"id":1, "title":"test"}`)
	})

	opt := &UpdateSnippetOptions{
		Title:       String("test"),
		FileName:    String("file"),
		Description: String("description"),
		Content:     String("content"),
		Visibility:  Visibility(PublicVisibility),
	}

	s, _, err := client.Snippets.UpdateSnippet(1, opt)
	require.NoError(t, err)

	want := &Snippet{ID: 1, Title: "test"}
	require.Equal(t, want, s)
}

func TestSnippetsService_DeleteSnippet(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/snippets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.Snippets.DeleteSnippet(1)
	require.NoError(t, err)
}

func TestSnippetsService_SnippetContent(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/snippets/1/raw", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, "Hello World")
	})

	b, _, err := client.Snippets.SnippetContent(1)
	require.NoError(t, err)

	want := []byte("Hello World")
	require.Equal(t, want, b)
}

func TestSnippetsService_ExploreSnippets(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/snippets/public", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":42,"title":"test"}]`)
	})

	opt := &ExploreSnippetsOptions{Page: 1, PerPage: 10}

	ss, _, err := client.Snippets.ExploreSnippets(opt)
	require.NoError(t, err)

	want := []*Snippet{{ID: 42, Title: "test"}}
	require.Equal(t, want, ss)
}
