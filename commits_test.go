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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

var testRevertCommitTargetBranch = "release"

func TestGetCommit(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/commits/b0b3a907f41409829b307a28b82fdbd552ee5a27", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_commit.json")
	})

	commit, resp, err := client.Commits.GetCommit("1", "b0b3a907f41409829b307a28b82fdbd552ee5a27", nil)
	if err != nil {
		t.Fatalf("Commits.GetCommit returned error: %v, response: %v", err, resp)
	}

	updatedAt := time.Date(2019, 11, 4, 15, 39, 0o3, 935000000, time.UTC)
	createdAt := time.Date(2019, 11, 4, 15, 38, 53, 154000000, time.UTC)
	want := &Commit{
		ID:             "6104942438c14ec7bd21c6cd5bd995272b3faff6",
		ShortID:        "6104942438c",
		Title:          "Sanitize for network graph",
		AuthorName:     "randx",
		AuthorEmail:    "dmitriy.zaporozhets@gmail.com",
		CommitterName:  "Dmitriy",
		CommitterEmail: "dmitriy.zaporozhets@gmail.com",
		Message:        "Sanitize for network graph",
		ParentIDs:      []string{"ae1d9fb46aa2b07ee9836d49862ec4e2c46fbbba"},
		Stats:          &CommitStats{Additions: 15, Deletions: 10, Total: 25},
		Status:         BuildState(Running),
		LastPipeline: &PipelineInfo{
			ID:        8,
			Ref:       "master",
			SHA:       "2dc6aa325a317eda67812f05600bdf0fcdc70ab0",
			Status:    "created",
			WebURL:    "https://gitlab.com/gitlab-org/gitlab-ce/pipelines/54268416",
			UpdatedAt: &updatedAt,
			CreatedAt: &createdAt,
		},
		ProjectID: 13083,
	}

	assert.Equal(t, want, commit)
}

func TestGetCommitStatuses(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/commits/b0b3a907f41409829b307a28b82fdbd552ee5a27/statuses", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1}]`)
	})

	opt := &GetCommitStatusesOptions{
		Ref:   String("master"),
		Stage: String("test"),
		Name:  String("ci/jenkins"),
		All:   Bool(true),
	}
	statuses, _, err := client.Commits.GetCommitStatuses("1", "b0b3a907f41409829b307a28b82fdbd552ee5a27", opt)
	if err != nil {
		t.Errorf("Commits.GetCommitStatuses returned error: %v", err)
	}

	want := []*CommitStatus{{ID: 1}}
	if !reflect.DeepEqual(want, statuses) {
		t.Errorf("Commits.GetCommitStatuses returned %+v, want %+v", statuses, want)
	}
}

func TestSetCommitStatus(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/statuses/b0b3a907f41409829b307a28b82fdbd552ee5a27", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)

		var content SetCommitStatusOptions
		err = json.Unmarshal(body, &content)
		require.NoError(t, err)

		assert.Equal(t, "ci/jenkins", *content.Name)
		assert.Equal(t, 99.9, *content.Coverage)
		fmt.Fprint(w, `{"id":1}`)
	})

	cov := 99.9
	opt := &SetCommitStatusOptions{
		State:       Running,
		Ref:         String("master"),
		Name:        String("ci/jenkins"),
		Context:     String(""),
		TargetURL:   String("http://abc"),
		Description: String("build"),
		Coverage:    &cov,
	}
	status, _, err := client.Commits.SetCommitStatus("1", "b0b3a907f41409829b307a28b82fdbd552ee5a27", opt)
	if err != nil {
		t.Errorf("Commits.SetCommitStatus returned error: %v", err)
	}

	want := &CommitStatus{ID: 1}
	if !reflect.DeepEqual(want, status) {
		t.Errorf("Commits.SetCommitStatus returned %+v, want %+v", status, want)
	}
}

func TestRevertCommit_NoOptions(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/commits/b0b3a907f41409829b307a28b82fdbd552ee5a27/revert", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		mustWriteHTTPResponse(t, w, "testdata/get_commit.json")
	})

	commit, resp, err := client.Commits.RevertCommit("1", "b0b3a907f41409829b307a28b82fdbd552ee5a27", nil)
	if err != nil {
		t.Fatalf("Commits.RevertCommit returned error: %v, response: %v", err, resp)
	}

	updatedAt := time.Date(2019, 11, 4, 15, 39, 0o3, 935000000, time.UTC)
	createdAt := time.Date(2019, 11, 4, 15, 38, 53, 154000000, time.UTC)
	want := &Commit{
		ID:             "6104942438c14ec7bd21c6cd5bd995272b3faff6",
		ShortID:        "6104942438c",
		Title:          "Sanitize for network graph",
		AuthorName:     "randx",
		AuthorEmail:    "dmitriy.zaporozhets@gmail.com",
		CommitterName:  "Dmitriy",
		CommitterEmail: "dmitriy.zaporozhets@gmail.com",
		Message:        "Sanitize for network graph",
		ParentIDs:      []string{"ae1d9fb46aa2b07ee9836d49862ec4e2c46fbbba"},
		Stats:          &CommitStats{Additions: 15, Deletions: 10, Total: 25},
		Status:         BuildState(Running),
		LastPipeline: &PipelineInfo{
			ID:        8,
			Ref:       "master",
			SHA:       "2dc6aa325a317eda67812f05600bdf0fcdc70ab0",
			Status:    "created",
			WebURL:    "https://gitlab.com/gitlab-org/gitlab-ce/pipelines/54268416",
			UpdatedAt: &updatedAt,
			CreatedAt: &createdAt,
		},
		ProjectID: 13083,
	}

	assert.Equal(t, want, commit)
}

func TestRevertCommit_WithOptions(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/commits/b0b3a907f41409829b307a28b82fdbd552ee5a27/revert", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"branch":"release"}`)
		mustWriteHTTPResponse(t, w, "testdata/get_commit.json")
	})

	commit, resp, err := client.Commits.RevertCommit("1", "b0b3a907f41409829b307a28b82fdbd552ee5a27", &RevertCommitOptions{
		Branch: &testRevertCommitTargetBranch,
	})
	if err != nil {
		t.Fatalf("Commits.RevertCommit returned error: %v, response: %v", err, resp)
	}

	updatedAt := time.Date(2019, 11, 4, 15, 39, 0o3, 935000000, time.UTC)
	createdAt := time.Date(2019, 11, 4, 15, 38, 53, 154000000, time.UTC)
	want := &Commit{
		ID:             "6104942438c14ec7bd21c6cd5bd995272b3faff6",
		ShortID:        "6104942438c",
		Title:          "Sanitize for network graph",
		AuthorName:     "randx",
		AuthorEmail:    "dmitriy.zaporozhets@gmail.com",
		CommitterName:  "Dmitriy",
		CommitterEmail: "dmitriy.zaporozhets@gmail.com",
		Message:        "Sanitize for network graph",
		ParentIDs:      []string{"ae1d9fb46aa2b07ee9836d49862ec4e2c46fbbba"},
		Stats:          &CommitStats{Additions: 15, Deletions: 10, Total: 25},
		Status:         BuildState(Running),
		LastPipeline: &PipelineInfo{
			ID:        8,
			Ref:       "master",
			SHA:       "2dc6aa325a317eda67812f05600bdf0fcdc70ab0",
			Status:    "created",
			WebURL:    "https://gitlab.com/gitlab-org/gitlab-ce/pipelines/54268416",
			UpdatedAt: &updatedAt,
			CreatedAt: &createdAt,
		},
		ProjectID: 13083,
	}

	assert.Equal(t, want, commit)
}

func TestGetGPGSignature(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/commits/b0b3a907f41409829b307a28b82fdbd552ee5a27/signature", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_signature.json")
	})

	sig, resp, err := client.Commits.GetGPGSiganature("1", "b0b3a907f41409829b307a28b82fdbd552ee5a27", nil)
	if err != nil {
		t.Fatalf("Commits.GetGPGSignature returned error: %v, response: %v", err, resp)
	}

	want := &GPGSignature{
		KeyID:              7977,
		KeyPrimaryKeyID:    "627C5F589F467F17",
		KeyUserName:        "Dmitriy Zaporozhets",
		KeyUserEmail:       "dmitriy.zaporozhets@gmail.com",
		VerificationStatus: "verified",
		KeySubkeyID:        0,
	}

	assert.Equal(t, want, sig)
}

func TestCommitsService_ListCommits(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/commits", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
				{
					"id": "6104942438c14ec7bd21c6cd5bd995272b3faff6",
					"short_id": "6104942438c",
					"title": "Sanitize for network graph",
					"author_name": "randx",
					"author_email": "venkateshthalluri123@gmail.com",
					"committer_name": "Venkatesh",
					"committer_email": "venkateshthalluri123@gmail.com",
					"message": "Sanitize for network graph",
					"parent_ids": [
						"ae1d9fb46aa2b07ee9836d49862ec4e2c46fbbba"
					],
					"last_pipeline": {
						"id": 8,
						"ref": "master",
						"sha": "2dc6aa325a317eda67812f05600bdf0fcdc70ab0",
						"status": "created",
						"web_url": "https://gitlab.com/gitlab-org/gitlab-ce/pipelines/54268416",
						"created_at": "2019-11-04T15:38:53.154Z",
						"updated_at": "2019-11-04T15:39:03.935Z"
					},
					"stats": {
						"additions": 15,
						"deletions": 10,
						"total": 25
					},
					"status": "running",
					"project_id": 13083
				}
			]
		`)
	})

	updatedAt := time.Date(2019, 11, 4, 15, 39, 0o3, 935000000, time.UTC)
	createdAt := time.Date(2019, 11, 4, 15, 38, 53, 154000000, time.UTC)
	want := []*Commit{{
		ID:             "6104942438c14ec7bd21c6cd5bd995272b3faff6",
		ShortID:        "6104942438c",
		Title:          "Sanitize for network graph",
		AuthorName:     "randx",
		AuthorEmail:    "venkateshthalluri123@gmail.com",
		CommitterName:  "Venkatesh",
		CommitterEmail: "venkateshthalluri123@gmail.com",
		Message:        "Sanitize for network graph",
		ParentIDs:      []string{"ae1d9fb46aa2b07ee9836d49862ec4e2c46fbbba"},
		Stats:          &CommitStats{Additions: 15, Deletions: 10, Total: 25},
		Status:         BuildState(Running),
		LastPipeline: &PipelineInfo{
			ID:        8,
			Ref:       "master",
			SHA:       "2dc6aa325a317eda67812f05600bdf0fcdc70ab0",
			Status:    "created",
			WebURL:    "https://gitlab.com/gitlab-org/gitlab-ce/pipelines/54268416",
			UpdatedAt: &updatedAt,
			CreatedAt: &createdAt,
		},
		ProjectID: 13083,
	}}

	cs, resp, err := client.Commits.ListCommits(1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, cs)

	cs, resp, err = client.Commits.ListCommits(1.01, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, cs)

	cs, resp, err = client.Commits.ListCommits(1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, cs)

	cs, resp, err = client.Commits.ListCommits(3, nil)
	require.Error(t, err)
	require.Nil(t, cs)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestCommitsService_GetCommitRefs(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/commits/5937ac0a7beb003549fc5fd26fc247adbce4a52e/refs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {"type": "branch", "name": "test"},
			  {"type": "branch", "name": "add-balsamiq-file"},
			  {"type": "branch", "name": "wip"},
			  {"type": "tag", "name": "v1.1.0"}
			 ]
		`)
	})

	want := []*CommitRef{
		{
			Type: "branch",
			Name: "test",
		},
		{
			Type: "branch",
			Name: "add-balsamiq-file",
		},
		{
			Type: "branch",
			Name: "wip",
		},
		{
			Type: "tag",
			Name: "v1.1.0",
		},
	}

	crs, resp, err := client.Commits.GetCommitRefs(1, "5937ac0a7beb003549fc5fd26fc247adbce4a52e", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, crs)

	crs, resp, err = client.Commits.GetCommitRefs(1.01, "5937ac0a7beb003549fc5fd26fc247adbce4a52e", nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, crs)

	crs, resp, err = client.Commits.GetCommitRefs(1, "5937ac0a7beb003549fc5fd26fc247adbce4a52e", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, crs)

	crs, resp, err = client.Commits.GetCommitRefs(3, "5937ac0a7beb003549fc5fd26fc247adbce4a52e", nil)
	require.Error(t, err)
	require.Nil(t, crs)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestCommitsService_CreateCommit(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/commits", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
				"id": "6104942438c14ec7bd21c6cd5bd995272b3faff6",
				"short_id": "6104942438c",
				"title": "Sanitize for network graph",
				"author_name": "randx",
				"author_email": "venkateshthalluri123@gmail.com",
				"committer_name": "Venkatesh",
				"committer_email": "venkateshthalluri123@gmail.com",
				"message": "Sanitize for network graph",
				"parent_ids": [
					"ae1d9fb46aa2b07ee9836d49862ec4e2c46fbbba"
				],
				"last_pipeline": {
					"id": 8,
					"ref": "master",
					"sha": "2dc6aa325a317eda67812f05600bdf0fcdc70ab0",
					"status": "created",
					"web_url": "https://gitlab.com/gitlab-org/gitlab-ce/pipelines/54268416",
					"created_at": "2019-11-04T15:38:53.154Z",
					"updated_at": "2019-11-04T15:39:03.935Z"
				},
				"stats": {
					"additions": 15,
					"deletions": 10,
					"total": 25
				},
				"status": "running",
				"project_id": 13083
			}
		`)
	})

	updatedAt := time.Date(2019, 11, 4, 15, 39, 0o3, 935000000, time.UTC)
	createdAt := time.Date(2019, 11, 4, 15, 38, 53, 154000000, time.UTC)
	want := &Commit{
		ID:             "6104942438c14ec7bd21c6cd5bd995272b3faff6",
		ShortID:        "6104942438c",
		Title:          "Sanitize for network graph",
		AuthorName:     "randx",
		AuthorEmail:    "venkateshthalluri123@gmail.com",
		CommitterName:  "Venkatesh",
		CommitterEmail: "venkateshthalluri123@gmail.com",
		Message:        "Sanitize for network graph",
		ParentIDs:      []string{"ae1d9fb46aa2b07ee9836d49862ec4e2c46fbbba"},
		Stats:          &CommitStats{Additions: 15, Deletions: 10, Total: 25},
		Status:         BuildState(Running),
		LastPipeline: &PipelineInfo{
			ID:        8,
			Ref:       "master",
			SHA:       "2dc6aa325a317eda67812f05600bdf0fcdc70ab0",
			Status:    "created",
			WebURL:    "https://gitlab.com/gitlab-org/gitlab-ce/pipelines/54268416",
			UpdatedAt: &updatedAt,
			CreatedAt: &createdAt,
		},
		ProjectID: 13083,
	}

	c, resp, err := client.Commits.CreateCommit(1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, c)

	c, resp, err = client.Commits.CreateCommit(1.01, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, c)

	c, resp, err = client.Commits.CreateCommit(1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, c)

	c, resp, err = client.Commits.CreateCommit(3, nil)
	require.Error(t, err)
	require.Nil(t, c)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestCommitsService_GetCommitDiff(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/commits/master/diff", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"diff": "--- a/doc/update/5.4-to-6.0.md\n+++ b/doc/update/5.4-to-6.0.md\n@@ -71,6 +71,8 @@\n sudo -u git -H bundle exec rake migrate_keys RAILS_ENV=production\n sudo -u git -H bundle exec rake migrate_inline_notes RAILS_ENV=production\n \n+sudo -u git -H bundle exec rake gitlab:assets:compile RAILS_ENV=production\n+\n \n \n ### 6. Update config files",
				"new_path": "doc/update/5.4-to-6.0.md",
				"old_path": "doc/update/5.4-to-6.0.md",
				"a_mode": null,
				"b_mode": "100644",
				"new_file": false,
				"renamed_file": false,
				"deleted_file": false
				}
			]
		`)
	})

	want := []*Diff{{
		Diff:    "--- a/doc/update/5.4-to-6.0.md\n+++ b/doc/update/5.4-to-6.0.md\n@@ -71,6 +71,8 @@\n sudo -u git -H bundle exec rake migrate_keys RAILS_ENV=production\n sudo -u git -H bundle exec rake migrate_inline_notes RAILS_ENV=production\n \n+sudo -u git -H bundle exec rake gitlab:assets:compile RAILS_ENV=production\n+\n \n \n ### 6. Update config files",
		NewPath: "doc/update/5.4-to-6.0.md",
		OldPath: "doc/update/5.4-to-6.0.md",
		BMode:   "100644",
	}}

	ds, resp, err := client.Commits.GetCommitDiff(1, "master", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ds)

	ds, resp, err = client.Commits.GetCommitDiff(1.01, "master", nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ds)

	ds, resp, err = client.Commits.GetCommitDiff(1, "master", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ds)

	ds, resp, err = client.Commits.GetCommitDiff(3, "master", nil)
	require.Error(t, err)
	require.Nil(t, ds)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestCommitsService_GetCommitComments(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/commits/master/comments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"note": "this code is really nice",
				"author": {
				  "id": 11,
				  "username": "venky333",
				  "email": "venkateshthalluri123@gmail.com",
				  "name": "Venkatesh Thalluri",
				  "state": "active"
				}
			  }
			]
		`)
	})

	want := []*CommitComment{{
		Note: "this code is really nice",
		Author: Author{
			ID:       11,
			Username: "venky333",
			Email:    "venkateshthalluri123@gmail.com",
			Name:     "Venkatesh Thalluri",
			State:    "active",
		},
	}}

	ccs, resp, err := client.Commits.GetCommitComments(1, "master", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ccs)

	ccs, resp, err = client.Commits.GetCommitComments(1.01, "master", nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ccs)

	ccs, resp, err = client.Commits.GetCommitComments(1, "master", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ccs)

	ccs, resp, err = client.Commits.GetCommitComments(3, "master", nil)
	require.Error(t, err)
	require.Nil(t, ccs)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestCommitsService_PostCommitComment(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/commits/master/comments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"note": "this code is really nice",
			"author": {
			  "id": 11,
			  "username": "venky333",
			  "email": "venkateshthalluri123@gmail.com",
			  "name": "Venkatesh Thalluri",
			  "state": "active"
			}
		  }
		`)
	})

	want := &CommitComment{
		Note: "this code is really nice",
		Author: Author{
			ID:       11,
			Username: "venky333",
			Email:    "venkateshthalluri123@gmail.com",
			Name:     "Venkatesh Thalluri",
			State:    "active",
		},
	}

	cc, resp, err := client.Commits.PostCommitComment(1, "master", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, cc)

	cc, resp, err = client.Commits.PostCommitComment(1.01, "master", nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, cc)

	cc, resp, err = client.Commits.PostCommitComment(1, "master", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, cc)

	cc, resp, err = client.Commits.PostCommitComment(3, "master", nil)
	require.Error(t, err)
	require.Nil(t, cc)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestCommitsService_ListMergeRequestsByCommit(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/commits/master/merge_requests", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 1,
				"iid": 1,
				"project_id": 3,
				"title": "test1",
				"description": "fixed login page css paddings",
				"state": "merged",
				"merged_by": {
				  "id": 87854,
				  "name": "Douwe Maan",
				  "username": "DouweM",
				  "state": "active",
				  "avatar_url": "https://gitlab.example.com/uploads/-/system/user/avatar/87854/avatar.png",
				  "web_url": "https://gitlab.com/DouweM"
				},
				"closed_by": null,
				"closed_at": null,
				"target_branch": "master",
				"source_branch": "test1",
				"upvotes": 0,
				"downvotes": 0,
				"author": {
				  "id": 1,
				  "name": "Administrator",
				  "username": "admin",
				  "state": "active",
				  "avatar_url": null,
				  "web_url" : "https://gitlab.example.com/admin"
				},
				"assignee": {
				  "id": 1,
				  "name": "Administrator",
				  "username": "admin",
				  "state": "active",
				  "avatar_url": null,
				  "web_url" : "https://gitlab.example.com/admin"
				},
				"assignees": [{
				  "name": "Venkatesh Thalluri",
				  "username": "venkatesh.thalluri",
				  "id": 12,
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/46f6f7dc858ada7be1853f7fb96e81da?s=80&d=identicon",
				  "web_url": "https://gitlab.example.com/axel.block"
				}],
				"reviewers": [{
				  "id": 2,
				  "name": "Sam Bauch",
				  "username": "kenyatta_oconnell",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/956c92487c6f6f7616b536927e22c9a0?s=80&d=identicon",
				  "web_url": "http://gitlab.example.com//kenyatta_oconnell"
				}],
				"source_project_id": 2,
				"target_project_id": 3,
				"draft": false,
				"work_in_progress": false,
				"milestone": {
				  "id": 5,
				  "iid": 1,
				  "project_id": 3,
				  "title": "v2.0",
				  "description": "Assumenda aut placeat expedita exercitationem labore sunt enim earum.",
				  "state": "closed",
				  "web_url": "https://gitlab.example.com/my-group/my-project/milestones/1"
				},
				"merge_when_pipeline_succeeds": true,
				"detailed_merge_status": "mergeable",
				"sha": "8888888888888888888888888888888888888888",
				"merge_commit_sha": null,
				"squash_commit_sha": null,
				"user_notes_count": 1,
				"discussion_locked": null,
				"should_remove_source_branch": true,
				"force_remove_source_branch": false,
				"allow_collaboration": false,
				"allow_maintainer_to_push": false,
				"web_url": "http://gitlab.example.com/my-group/my-project/merge_requests/1",
				"references": {
				  "short": "!1",
				  "relative": "my-group/my-project!1",
				  "full": "my-group/my-project!1"
				},
				"squash": false,
				"task_completion_status":{
				  "count":0,
				  "completed_count":0
				}
			  }
			]
		`)
	})

	want := []*MergeRequest{{
		ID:           1,
		IID:          1,
		TargetBranch: "master",
		SourceBranch: "test1",
		ProjectID:    3,
		Title:        "test1",
		State:        "merged",
		Upvotes:      0,
		Downvotes:    0,
		Author: &BasicUser{
			ID:        1,
			Username:  "admin",
			Name:      "Administrator",
			State:     "active",
			CreatedAt: nil,
			AvatarURL: "",
			WebURL:    "https://gitlab.example.com/admin",
		},
		Assignee: &BasicUser{
			ID: 1, Username: "admin",
			Name:      "Administrator",
			State:     "active",
			AvatarURL: "",
			WebURL:    "https://gitlab.example.com/admin",
		},
		Assignees: []*BasicUser{{
			ID:        12,
			Username:  "venkatesh.thalluri",
			Name:      "Venkatesh Thalluri",
			State:     "active",
			AvatarURL: "http://www.gravatar.com/avatar/46f6f7dc858ada7be1853f7fb96e81da?s=80&d=identicon", WebURL: "https://gitlab.example.com/axel.block",
		}},
		Reviewers: []*BasicUser{{
			ID:        2,
			Username:  "kenyatta_oconnell",
			Name:      "Sam Bauch",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/956c92487c6f6f7616b536927e22c9a0?s=80&d=identicon", WebURL: "http://gitlab.example.com//kenyatta_oconnell",
		}},
		SourceProjectID: 2,
		TargetProjectID: 3,
		Description:     "fixed login page css paddings",
		WorkInProgress:  false,
		Milestone: &Milestone{
			ID:          5,
			IID:         1,
			ProjectID:   3,
			Title:       "v2.0",
			Description: "Assumenda aut placeat expedita exercitationem labore sunt enim earum.",
			State:       "closed",
			WebURL:      "https://gitlab.example.com/my-group/my-project/milestones/1",
		},
		MergeWhenPipelineSucceeds: true,
		DetailedMergeStatus:       "mergeable",
		MergeError:                "",
		MergedBy: &BasicUser{
			ID:        87854,
			Username:  "DouweM",
			Name:      "Douwe Maan",
			State:     "active",
			AvatarURL: "https://gitlab.example.com/uploads/-/system/user/avatar/87854/avatar.png",
			WebURL:    "https://gitlab.com/DouweM",
		},
		Subscribed:               false,
		SHA:                      "8888888888888888888888888888888888888888",
		MergeCommitSHA:           "",
		SquashCommitSHA:          "",
		UserNotesCount:           1,
		ChangesCount:             "",
		ShouldRemoveSourceBranch: true,
		ForceRemoveSourceBranch:  false,
		AllowCollaboration:       false,
		WebURL:                   "http://gitlab.example.com/my-group/my-project/merge_requests/1",
		References: &IssueReferences{
			Short:    "!1",
			Relative: "my-group/my-project!1",
			Full:     "my-group/my-project!1",
		},
		DiscussionLocked:     false,
		Squash:               false,
		DivergedCommitsCount: 0,
		RebaseInProgress:     false,
		ApprovalsBeforeMerge: 0,
		Reference:            "",
		FirstContribution:    false,
		TaskCompletionStatus: &TasksCompletionStatus{
			Count:          0,
			CompletedCount: 0,
		},
		HasConflicts:                false,
		BlockingDiscussionsResolved: false,
		Overflow:                    false,
	}}

	mrs, resp, err := client.Commits.ListMergeRequestsByCommit(1, "master", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, mrs)

	mrs, resp, err = client.Commits.ListMergeRequestsByCommit(1.01, "master", nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, mrs)

	mrs, resp, err = client.Commits.ListMergeRequestsByCommit(1, "master", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, mrs)

	mrs, resp, err = client.Commits.ListMergeRequestsByCommit(3, "master", nil)
	require.Error(t, err)
	require.Nil(t, mrs)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestCommitsService_CherryPickCommit(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/commits/master/cherry_pick", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
				"id": "6104942438c14ec7bd21c6cd5bd995272b3faff6",
				"short_id": "6104942438c",
				"title": "Sanitize for network graph",
				"author_name": "randx",
				"author_email": "venkateshthalluri123@gmail.com",
				"committer_name": "Venkatesh",
				"committer_email": "venkateshthalluri123@gmail.com",
				"message": "Sanitize for network graph",
				"parent_ids": [
					"ae1d9fb46aa2b07ee9836d49862ec4e2c46fbbba"
				],
				"last_pipeline": {
					"id": 8,
					"ref": "master",
					"sha": "2dc6aa325a317eda67812f05600bdf0fcdc70ab0",
					"status": "created",
					"web_url": "https://gitlab.com/gitlab-org/gitlab-ce/pipelines/54268416",
					"created_at": "2019-11-04T15:38:53.154Z",
					"updated_at": "2019-11-04T15:39:03.935Z"
				},
				"stats": {
					"additions": 15,
					"deletions": 10,
					"total": 25
				},
				"status": "running",
				"project_id": 13083
			}
		`)
	})

	updatedAt := time.Date(2019, 11, 4, 15, 39, 0o3, 935000000, time.UTC)
	createdAt := time.Date(2019, 11, 4, 15, 38, 53, 154000000, time.UTC)
	want := &Commit{
		ID:             "6104942438c14ec7bd21c6cd5bd995272b3faff6",
		ShortID:        "6104942438c",
		Title:          "Sanitize for network graph",
		AuthorName:     "randx",
		AuthorEmail:    "venkateshthalluri123@gmail.com",
		CommitterName:  "Venkatesh",
		CommitterEmail: "venkateshthalluri123@gmail.com",
		Message:        "Sanitize for network graph",
		ParentIDs:      []string{"ae1d9fb46aa2b07ee9836d49862ec4e2c46fbbba"},
		Stats:          &CommitStats{Additions: 15, Deletions: 10, Total: 25},
		Status:         BuildState(Running),
		LastPipeline: &PipelineInfo{
			ID:        8,
			Ref:       "master",
			SHA:       "2dc6aa325a317eda67812f05600bdf0fcdc70ab0",
			Status:    "created",
			WebURL:    "https://gitlab.com/gitlab-org/gitlab-ce/pipelines/54268416",
			UpdatedAt: &updatedAt,
			CreatedAt: &createdAt,
		},
		ProjectID: 13083,
	}

	c, resp, err := client.Commits.CherryPickCommit(1, "master", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, c)

	c, resp, err = client.Commits.CherryPickCommit(1.01, "master", nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, c)

	c, resp, err = client.Commits.CherryPickCommit(1, "master", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, c)

	c, resp, err = client.Commits.CherryPickCommit(3, "master", nil)
	require.Error(t, err)
	require.Nil(t, c)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
