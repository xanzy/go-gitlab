//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBranch(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/branches/master", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
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
			AuthorName:     exampleEventUserName,
			AuthoredDate:   &authoredDate,
			CommittedDate:  &committedDate,
			CommitterEmail: "john@example.com",
			CommitterName:  exampleEventUserName,
			ID:             "7b5c3cc8be40ee161ae89a06bba6229da1032a0c",
			ShortID:        "7b5c3cc",
			Title:          "add projects API",
			Message:        "add projects API",
			ParentIDs:      []string{"4ad91d3c1144c406e50c7b33bae684bd6837faf8"},
		},
	}

	assert.Equal(t, want, branch)
}

func TestBranchesService_ListBranches(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/repository/branches", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_branches.json")
	})

	authoredDate := time.Date(2012, 6, 27, 5, 51, 39, 0, time.UTC)
	committedDate := time.Date(2012, 6, 28, 3, 44, 20, 0, time.UTC)
	want := []*Branch{{
		Name:               "master",
		Merged:             false,
		Protected:          true,
		Default:            true,
		DevelopersCanPush:  false,
		DevelopersCanMerge: false,
		CanPush:            true,
		WebURL:             "https://gitlab.example.com/my-group/my-project/-/tree/master",
		Commit: &Commit{
			AuthorEmail:    "john@example.com",
			AuthorName:     exampleEventUserName,
			AuthoredDate:   &authoredDate,
			CommittedDate:  &committedDate,
			CommitterEmail: "john@example.com",
			CommitterName:  exampleEventUserName,
			ID:             "7b5c3cc8be40ee161ae89a06bba6229da1032a0c",
			ShortID:        "7b5c3cc",
			Title:          "add projects API",
			Message:        "add projects API",
			ParentIDs:      []string{"4ad91d3c1144c406e50c7b33bae684bd6837faf8"},
		},
	}}

	b, resp, err := client.Branches.ListBranches(5, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, b)

	b, resp, err = client.Branches.ListBranches(5.01, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.Branches.ListBranches(5, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.Branches.ListBranches(3, nil)
	require.Error(t, err)
	require.Nil(t, b)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestBranchesService_ProtectBranch(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/branches/master/protect", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		mustWriteHTTPResponse(t, w, "testdata/get_branch.json")
	})

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
			AuthorName:     exampleEventUserName,
			AuthoredDate:   &authoredDate,
			CommittedDate:  &committedDate,
			CommitterEmail: "john@example.com",
			CommitterName:  exampleEventUserName,
			ID:             "7b5c3cc8be40ee161ae89a06bba6229da1032a0c",
			ShortID:        "7b5c3cc",
			Title:          "add projects API",
			Message:        "add projects API",
			ParentIDs:      []string{"4ad91d3c1144c406e50c7b33bae684bd6837faf8"},
		},
	}

	b, resp, err := client.Branches.ProtectBranch(1, "master", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, b)

	b, resp, err = client.Branches.ProtectBranch(1.01, "master", nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.Branches.ProtectBranch(1, "master", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.Branches.ProtectBranch(3, "master", nil)
	require.Error(t, err)
	require.Nil(t, b)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestBranchesService_UnprotectBranch(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/branches/master/unprotect", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		mustWriteHTTPResponse(t, w, "testdata/get_branch.json")
	})

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
			AuthorName:     exampleEventUserName,
			AuthoredDate:   &authoredDate,
			CommittedDate:  &committedDate,
			CommitterEmail: "john@example.com",
			CommitterName:  exampleEventUserName,
			ID:             "7b5c3cc8be40ee161ae89a06bba6229da1032a0c",
			ShortID:        "7b5c3cc",
			Title:          "add projects API",
			Message:        "add projects API",
			ParentIDs:      []string{"4ad91d3c1144c406e50c7b33bae684bd6837faf8"},
		},
	}

	b, resp, err := client.Branches.UnprotectBranch(1, "master", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, b)

	b, resp, err = client.Branches.UnprotectBranch(1.01, "master", nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.Branches.UnprotectBranch(1, "master", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.Branches.UnprotectBranch(3, "master", nil)
	require.Error(t, err)
	require.Nil(t, b)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestBranchesService_CreateBranch(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/branches", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		mustWriteHTTPResponse(t, w, "testdata/get_branch.json")
	})

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
			AuthorName:     exampleEventUserName,
			AuthoredDate:   &authoredDate,
			CommittedDate:  &committedDate,
			CommitterEmail: "john@example.com",
			CommitterName:  exampleEventUserName,
			ID:             "7b5c3cc8be40ee161ae89a06bba6229da1032a0c",
			ShortID:        "7b5c3cc",
			Title:          "add projects API",
			Message:        "add projects API",
			ParentIDs:      []string{"4ad91d3c1144c406e50c7b33bae684bd6837faf8"},
		},
	}

	b, resp, err := client.Branches.CreateBranch(1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, b)

	b, resp, err = client.Branches.CreateBranch(1.01, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.Branches.CreateBranch(1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.Branches.CreateBranch(3, nil)
	require.Error(t, err)
	require.Nil(t, b)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestBranchesService_DeleteBranch(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/branches/master", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Branches.DeleteBranch(1, "master", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.Branches.DeleteBranch(1.01, "master", nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.Branches.DeleteBranch(1, "master", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.Branches.DeleteBranch(3, "master", nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestBranchesService_DeleteMergedBranches(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/merged_branches", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Branches.DeleteMergedBranches(1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.Branches.DeleteMergedBranches(1.01, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.Branches.DeleteMergedBranches(1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.Branches.DeleteMergedBranches(3, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
