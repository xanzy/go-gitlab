package gitlab

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

var (
	resp     = `{"id":1, "title":"test"}`
	respList = `[{"id":42,"title":"Voluptatem iure ut qui aut et consequatur quaerat."}]`
	want     = &Snippet{ID: 1, Title: "test"}
)

func TestSnippetsService_ListSnippets(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/snippets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, respList)
	})

	opt := &ListSnippetsOptions{Page: 1, PerPage: 10}

	ss, _, err := client.Snippets.ListSnippets(opt)
	require.NoError(t, err)

	want := []*Snippet{{ID: 42, Title: "Voluptatem iure ut qui aut et consequatur quaerat."}}
	require.Equal(t, want, ss)
}

func TestSnippetsService_GetSnippet(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/snippets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, resp)
	})

	s, _, err := client.Snippets.GetSnippet(1)
	require.NoError(t, err)

	require.Equal(t, want, s)
}

func TestSnippetsService_CreateSnippet(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/snippets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, resp)
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

	require.Equal(t, want, s)
}

func TestSnippetsService_UpdateSnippet(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/snippets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, resp)
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

	require.Equal(t, want, s)
}

func TestSnippetsService_DeleteSnippet(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/snippets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.Snippets.DeleteSnippet(1)
	require.NoError(t, err)
}

func TestSnippetsService_SnippetContent(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

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
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/snippets/public", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, respList)
	})

	opt := &ExploreSnippetsOptions{Page: 1, PerPage: 10}

	ss, _, err := client.Snippets.ExploreSnippets(opt)
	require.NoError(t, err)

	want := []*Snippet{{ID: 42, Title: "Voluptatem iure ut qui aut et consequatur quaerat."}}
	require.Equal(t, want, ss)
}
