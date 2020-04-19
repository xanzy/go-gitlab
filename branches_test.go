package gitlab

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetBranch(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/repository/branches/master", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		mustWriteHTTPResponse(t, w, "testdata/get_branch.json")
	})

	branch, resp, err := client.Branches.GetBranch(1, "master")
	if err != nil {
		t.Fatalf("Branches.GetBranch returned error: %v, response %v", err, resp)
	}

	authoredDate := time.Date(2012, 6, 27, 5, 51, 39, 0, time.UTC)
	committedDate := time.Date(2012, 6, 28, 3, 44, 20, 0, time.UTC)
	want := &Branch{
		Name:               "master",
		Merged:             false,
		Protected:          true,
		Default:            true,
		DevelopersCanPush:  false,
		DevelopersCanMerge: false,
		CanPush:            true,
		Commit: &Commit{
			AuthorEmail:    "john@example.com",
			AuthorName:     "John Smith",
			AuthoredDate:   &authoredDate,
			CommittedDate:  &committedDate,
			CommitterEmail: "john@example.com",
			CommitterName:  "John Smith",
			ID:             "7b5c3cc8be40ee161ae89a06bba6229da1032a0c",
			ShortID:        "7b5c3cc",
			Title:          "add projects API",
			Message:        "add projects API",
			ParentIDs:      []string{"4ad91d3c1144c406e50c7b33bae684bd6837faf8"},
		},
	}

	assert.Equal(t, want, branch)
}
