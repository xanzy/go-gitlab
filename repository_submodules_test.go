package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepositorySubmodulesService_UpdateSubmodule(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/13083/repository/submodules/app%2Fproject", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
			{
			  "id": "6104942438c14ec7bd21c6cd5bd995272b3faff6",
			  "short_id": "6104942438c",
			  "title": "Update my submodule",
			  "author_name": "popdabubbles",
			  "author_email": "noty@gmail.com",
			  "committer_name": "Will",
			  "committer_email": "noty@gmail.com",
			  "message": "Update my submodule",
			  "parent_ids": [
			    "ae1d9fb46aa2b07ee9836d49862ec4e2c46fbbba"
			  ],
			  "status": "running"
			}
		`)
	})

	want := &SubmoduleCommit{
		ID:             "6104942438c14ec7bd21c6cd5bd995272b3faff6",
		ShortID:        "6104942438c",
		Title:          "Update my submodule",
		AuthorName:     "popdabubbles",
		AuthorEmail:    "noty@gmail.com",
		CommitterName:  "Will",
		CommitterEmail: "noty@gmail.com",
		Message:        "Update my submodule",
		ParentIDs:      []string{"ae1d9fb46aa2b07ee9836d49862ec4e2c46fbbba"},
		Status:         BuildState(Running),
	}

	sc, resp, err := client.RepositorySubmodules.UpdateSubmodule(13083, "app%2Fproject", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, sc)
}
