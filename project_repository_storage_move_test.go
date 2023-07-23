package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProjectRepositoryStorageMove_RetrieveAllProjectStorageMoves(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/project_repository_storage_moves", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
		{
		  "id":123,
		  "state":"scheduled",
		  "project":{
			"id":1,
			"name":"Test Project"
		  }
		},
		{
		  "id":122,
		  "state":"finished",
		  "project":{
			"id":2,
			"name":"Test Project 2"
		  }
		}]`)
	})

	opts := RetrieveAllProjectStorageMovesOptions{Page: 1, PerPage: 2}

	ssms, _, err := client.ProjectRepositoryStorageMove.RetrieveAllStorageMoves(opts)
	require.NoError(t, err)

	want := []*ProjectRepositoryStorageMove{
		{
			ID:    123,
			State: "scheduled",
			Project: &RepositoryProject{
				ID:   1,
				Name: "Test Project",
			},
		},
		{
			ID:    122,
			State: "finished",
			Project: &RepositoryProject{
				ID:   2,
				Name: "Test Project 2",
			},
		},
	}
	require.Equal(t, want, ssms)
}

func TestProjectRepositoryStorageMove_RetrieveAllStorageMovesForProject(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository_storage_moves", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
		{
		  "id":123,
		  "state":"scheduled",
		  "project":{
			"id":1,
			"name":"Test Project"
		  }
		},
		{
		  "id":122,
		  "state":"finished",
		  "project":{
			"id":1,
			"name":"Test Project"
		  }
		}]`)
	})

	opts := RetrieveAllProjectStorageMovesOptions{Page: 1, PerPage: 2}

	ssms, _, err := client.ProjectRepositoryStorageMove.RetrieveAllStorageMovesForProject(1, opts)
	require.NoError(t, err)

	want := []*ProjectRepositoryStorageMove{
		{
			ID:    123,
			State: "scheduled",
			Project: &RepositoryProject{
				ID:   1,
				Name: "Test Project",
			},
		},
		{
			ID:    122,
			State: "finished",
			Project: &RepositoryProject{
				ID:   1,
				Name: "Test Project",
			},
		},
	}
	require.Equal(t, want, ssms)
}

func TestProjectRepositoryStorageMove_GetStorageMove(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/project_repository_storage_moves/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
		{
		  "id":123,
		  "state":"scheduled",
		  "project":{
			"id":1,
			"name":"Test Project"
		  }
		}`)
	})

	ssm, _, err := client.ProjectRepositoryStorageMove.GetStorageMove(123)
	require.NoError(t, err)

	want := &ProjectRepositoryStorageMove{
		ID:    123,
		State: "scheduled",
		Project: &RepositoryProject{
			ID:   1,
			Name: "Test Project",
		},
	}
	require.Equal(t, want, ssm)
}

func TestProjectRepositoryStorageMove_GetStorageMoveForProject(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository_storage_moves/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
		{
		  "id":123,
		  "state":"scheduled",
		  "project":{
			"id":1,
			"name":"Test Project"
		  }
		}`)
	})

	ssm, _, err := client.ProjectRepositoryStorageMove.GetStorageMoveForProject(1, 123)
	require.NoError(t, err)

	want := &ProjectRepositoryStorageMove{
		ID:    123,
		State: "scheduled",
		Project: &RepositoryProject{
			ID:   1,
			Name: "Test Project",
		},
	}
	require.Equal(t, want, ssm)
}

func TestProjectRepositoryStorageMove_ScheduleStorageMoveForProject(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository_storage_moves", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `
		{
		  "id":124,
		  "state":"scheduled",
		  "project":{
			"id":1,
			"name":"Test Project"
		  }
		}`)
	})

	ssm, _, err := client.ProjectRepositoryStorageMove.ScheduleStorageMoveForProject(1, ScheduleStorageMoveForProjectOptions{})
	require.NoError(t, err)

	want := &ProjectRepositoryStorageMove{
		ID:    124,
		State: "scheduled",
		Project: &RepositoryProject{
			ID:   1,
			Name: "Test Project",
		},
	}
	require.Equal(t, want, ssm)
}

func TestProjectRepositoryStorageMove_ScheduleAllStorageMoves(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/project_repository_storage_moves", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"message": "202 Accepted"}`)
	})

	_, err := client.ProjectRepositoryStorageMove.ScheduleAllStorageMoves(
		ScheduleAllProjectStorageMovesOptions{
			SourceStorageName: String("default"),
		},
	)
	require.NoError(t, err)
}
