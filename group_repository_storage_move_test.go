package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGroupRepositoryStorageMove_RetrieveAllGroupStorageMoves(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/group_repository_storage_moves", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
		{
		  "id":123,
		  "state":"scheduled",
		  "group":{
			"id":283,
			"Name":"Test Group",
			"web_url": "https://gitlab.example.com/groups/test_group"
		  }
		},
		{
		  "id":122,
		  "state":"finished",
		  "group":{
			"id":284,
			"Name":"Test Group 2",
			"web_url": "https://gitlab.example.com/groups/test_group_2"
		  }
		}]`)
	})

	opts := RetrieveAllGroupStorageMovesOptions{Page: 1, PerPage: 2}

	gsms, _, err := client.GroupRepositoryStorageMove.RetrieveAllStorageMoves(opts)
	require.NoError(t, err)

	want := []*GroupRepositoryStorageMove{
		{
			ID:    123,
			State: "scheduled",
			Group: &RepositoryGroup{
				ID:     283,
				Name:   "Test Group",
				WebURL: "https://gitlab.example.com/groups/test_group",
			},
		},
		{
			ID:    122,
			State: "finished",
			Group: &RepositoryGroup{
				ID:     284,
				Name:   "Test Group 2",
				WebURL: "https://gitlab.example.com/groups/test_group_2",
			},
		},
	}
	require.Equal(t, want, gsms)
}

func TestGroupRepositoryStorageMove_RetrieveAllStorageMovesForGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/283/repository_storage_moves", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
		{
		  "id":123,
		  "state":"scheduled",
		  "group":{
			"id":283,
			"Name":"Test Group",
			"web_url": "https://gitlab.example.com/groups/test_group"
		  }
		},
		{
		  "id":122,
		  "state":"finished",
		  "group":{
			"id":283,
			"Name":"Test Group",
			"web_url": "https://gitlab.example.com/groups/test_group"
		  }
		}]`)
	})

	opts := RetrieveAllGroupStorageMovesOptions{Page: 1, PerPage: 2}

	gsms, _, err := client.GroupRepositoryStorageMove.RetrieveAllStorageMovesForGroup(283, opts)
	require.NoError(t, err)

	want := []*GroupRepositoryStorageMove{
		{
			ID:    123,
			State: "scheduled",
			Group: &RepositoryGroup{
				ID:     283,
				Name:   "Test Group",
				WebURL: "https://gitlab.example.com/groups/test_group",
			},
		},
		{
			ID:    122,
			State: "finished",
			Group: &RepositoryGroup{
				ID:     283,
				Name:   "Test Group",
				WebURL: "https://gitlab.example.com/groups/test_group",
			},
		},
	}
	require.Equal(t, want, gsms)
}

func TestGroupRepositoryStorageMove_GetStorageMove(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/group_repository_storage_moves/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
		{
		  "id":123,
		  "state":"scheduled",
		  "group":{
			"id":283,
			"Name":"Test Group",
			"web_url": "https://gitlab.example.com/groups/test_group"
		  }
		}`)
	})

	gsm, _, err := client.GroupRepositoryStorageMove.GetStorageMove(123)
	require.NoError(t, err)

	want := &GroupRepositoryStorageMove{
		ID:    123,
		State: "scheduled",
		Group: &RepositoryGroup{
			ID:     283,
			Name:   "Test Group",
			WebURL: "https://gitlab.example.com/groups/test_group",
		},
	}
	require.Equal(t, want, gsm)
}

func TestGroupRepositoryStorageMove_GetStorageMoveForGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/283/repository_storage_moves/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
		{
		  "id":123,
		  "state":"scheduled",
		  "group":{
			"id":283,
			"Name":"Test Group",
			"web_url": "https://gitlab.example.com/groups/test_group"
		  }
		}`)
	})

	gsm, _, err := client.GroupRepositoryStorageMove.GetStorageMoveForGroup(283, 123)
	require.NoError(t, err)

	want := &GroupRepositoryStorageMove{
		ID:    123,
		State: "scheduled",
		Group: &RepositoryGroup{
			ID:     283,
			Name:   "Test Group",
			WebURL: "https://gitlab.example.com/groups/test_group",
		},
	}
	require.Equal(t, want, gsm)
}

func TestGroupRepositoryStorageMove_ScheduleStorageMoveForGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/283/repository_storage_moves", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `
		{
		  "id":123,
		  "state":"scheduled",
		  "group":{
			"id":283,
			"Name":"Test Group",
			"web_url": "https://gitlab.example.com/groups/test_group"
		  }
		}`)
	})

	ssm, _, err := client.GroupRepositoryStorageMove.ScheduleStorageMoveForGroup(283, ScheduleStorageMoveForGroupOptions{})
	require.NoError(t, err)

	want := &GroupRepositoryStorageMove{
		ID:    123,
		State: "scheduled",
		Group: &RepositoryGroup{
			ID:     283,
			Name:   "Test Group",
			WebURL: "https://gitlab.example.com/groups/test_group",
		},
	}
	require.Equal(t, want, ssm)
}

func TestGroupRepositoryStorageMove_ScheduleAllStorageMoves(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/group_repository_storage_moves", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"message": "202 Accepted"}`)
	})

	_, err := client.GroupRepositoryStorageMove.ScheduleAllStorageMoves(
		ScheduleAllGroupStorageMovesOptions{
			SourceStorageName: Ptr("default"),
		},
	)
	require.NoError(t, err)
}
