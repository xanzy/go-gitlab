package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSnippetRepositoryStorageMove_RetrieveAllSnippetStorageMoves(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/snippet_repository_storage_moves", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
		{
		  "id":123,
		  "state":"scheduled",
		  "snippet":{
			"id":65,
			"title":"Test Snippet"
		  }
		},
		{
		  "id":122,
		  "state":"finished",
		  "snippet":{
			"id":64,
			"title":"Test Snippet 2"
		  }
		}]`)
	})

	opts := RetrieveAllSnippetStorageMovesOptions{Page: 1, PerPage: 2}

	ssms, _, err := client.SnippetRepositoryStorageMove.RetrieveAllStorageMoves(opts)
	require.NoError(t, err)

	want := []*SnippetRepositoryStorageMove{
		{
			ID:    123,
			State: "scheduled",
			Snippet: &RepositorySnippet{
				ID:    65,
				Title: "Test Snippet",
			},
		},
		{
			ID:    122,
			State: "finished",
			Snippet: &RepositorySnippet{
				ID:    64,
				Title: "Test Snippet 2",
			},
		},
	}
	require.Equal(t, want, ssms)
}

func TestSnippetRepositoryStorageMove_RetrieveAllStorageMovesForSnippet(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/snippets/65/repository_storage_moves", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
		{
		  "id":123,
		  "state":"scheduled",
		  "snippet":{
			"id":65,
			"title":"Test Snippet"
		  }
		},
		{
		  "id":122,
		  "state":"finished",
		  "snippet":{
			"id":65,
			"title":"Test Snippet"
		  }
		}]`)
	})

	opts := RetrieveAllSnippetStorageMovesOptions{Page: 1, PerPage: 2}

	ssms, _, err := client.SnippetRepositoryStorageMove.RetrieveAllStorageMovesForSnippet(65, opts)
	require.NoError(t, err)

	want := []*SnippetRepositoryStorageMove{
		{
			ID:    123,
			State: "scheduled",
			Snippet: &RepositorySnippet{
				ID:    65,
				Title: "Test Snippet",
			},
		},
		{
			ID:    122,
			State: "finished",
			Snippet: &RepositorySnippet{
				ID:    65,
				Title: "Test Snippet",
			},
		},
	}
	require.Equal(t, want, ssms)
}

func TestSnippetRepositoryStorageMove_GetStorageMove(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/snippet_repository_storage_moves/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
		{
		  "id":123,
		  "state":"scheduled",
		  "snippet":{
			"id":65,
			"title":"Test Snippet"
		  }
		}`)
	})

	ssm, _, err := client.SnippetRepositoryStorageMove.GetStorageMove(123)
	require.NoError(t, err)

	want := &SnippetRepositoryStorageMove{
		ID:    123,
		State: "scheduled",
		Snippet: &RepositorySnippet{
			ID:    65,
			Title: "Test Snippet",
		},
	}
	require.Equal(t, want, ssm)
}

func TestSnippetRepositoryStorageMove_GetStorageMoveForSnippet(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/snippets/65/repository_storage_moves/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
		{
		  "id":123,
		  "state":"scheduled",
		  "snippet":{
			"id":65,
			"title":"Test Snippet"
		  }
		}`)
	})

	ssm, _, err := client.SnippetRepositoryStorageMove.GetStorageMoveForSnippet(65, 123)
	require.NoError(t, err)

	want := &SnippetRepositoryStorageMove{
		ID:    123,
		State: "scheduled",
		Snippet: &RepositorySnippet{
			ID:    65,
			Title: "Test Snippet",
		},
	}
	require.Equal(t, want, ssm)
}

func TestSnippetRepositoryStorageMove_ScheduleStorageMoveForSnippet(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/snippets/65/repository_storage_moves", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `
		{
		  "id":124,
		  "state":"scheduled",
		  "snippet":{
			"id":65,
			"title":"Test Snippet"
		  }
		}`)
	})

	ssm, _, err := client.SnippetRepositoryStorageMove.ScheduleStorageMoveForSnippet(65, ScheduleStorageMoveForSnippetOptions{})
	require.NoError(t, err)

	want := &SnippetRepositoryStorageMove{
		ID:    124,
		State: "scheduled",
		Snippet: &RepositorySnippet{
			ID:    65,
			Title: "Test Snippet",
		},
	}
	require.Equal(t, want, ssm)
}

func TestSnippetRepositoryStorageMove_ScheduleAllStorageMoves(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/snippet_repository_storage_moves", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"message": "202 Accepted"}`)
	})

	_, err := client.SnippetRepositoryStorageMove.ScheduleAllStorageMoves(
		ScheduleAllSnippetStorageMovesOptions{
			SourceStorageName: Ptr("default"),
		},
	)
	require.NoError(t, err)
}
