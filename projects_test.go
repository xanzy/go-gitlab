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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListProjects(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{
		ListOptions: ListOptions{Page: 2, PerPage: 3},
		Archived:    Ptr(true),
		OrderBy:     Ptr("name"),
		Sort:        Ptr("asc"),
		Search:      Ptr("query"),
		Simple:      Ptr(true),
		Visibility:  Ptr(PublicVisibility),
	}

	projects, _, err := client.Projects.ListProjects(opt)
	if err != nil {
		t.Errorf("Projects.ListProjects returned error: %v", err)
	}

	want := []*Project{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Projects.ListProjects returned %+v, want %+v", projects, want)
	}
}

func TestListUserProjects(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/users/1/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{
		ListOptions: ListOptions{Page: 2, PerPage: 3},
		Archived:    Ptr(true),
		OrderBy:     Ptr("name"),
		Sort:        Ptr("asc"),
		Search:      Ptr("query"),
		Simple:      Ptr(true),
		Visibility:  Ptr(PublicVisibility),
	}

	projects, _, err := client.Projects.ListUserProjects(1, opt)
	if err != nil {
		t.Errorf("Projects.ListUserProjects returned error: %v", err)
	}

	want := []*Project{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Projects.ListUserProjects returned %+v, want %+v", projects, want)
	}
}

func TestListUserContributedProjects(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/users/1/contributed_projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{
		ListOptions: ListOptions{Page: 2, PerPage: 3},
		Archived:    Bool(true),
		OrderBy:     String("name"),
		Sort:        String("asc"),
		Search:      String("query"),
		Simple:      Bool(true),
		Visibility:  Visibility(PublicVisibility),
	}

	projects, _, err := client.Projects.ListUserContributedProjects(1, opt)
	if err != nil {
		t.Errorf("Projects.ListUserContributedProjects returned error: %v", err)
	}

	want := []*Project{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Projects.ListUserContributedProjects returned %+v, want %+v", projects, want)
	}
}

func TestListUserStarredProjects(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/users/1/starred_projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{
		ListOptions: ListOptions{Page: 2, PerPage: 3},
		Archived:    Ptr(true),
		OrderBy:     Ptr("name"),
		Sort:        Ptr("asc"),
		Search:      Ptr("query"),
		Simple:      Ptr(true),
		Visibility:  Ptr(PublicVisibility),
	}

	projects, _, err := client.Projects.ListUserStarredProjects(1, opt)
	if err != nil {
		t.Errorf("Projects.ListUserStarredProjects returned error: %v", err)
	}

	want := []*Project{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Projects.ListUserStarredProjects returned %+v, want %+v", projects, want)
	}
}

func TestListProjectsUsersByID(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/", func(w http.ResponseWriter, r *http.Request) {
		testURL(t, r, "/api/v4/projects/1/users?page=2&per_page=3&search=query")
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectUserOptions{
		ListOptions: ListOptions{Page: 2, PerPage: 3},
		Search:      Ptr("query"),
	}

	projects, _, err := client.Projects.ListProjectsUsers(1, opt)
	if err != nil {
		t.Errorf("Projects.ListProjectsUsers returned error: %v", err)
	}

	want := []*ProjectUser{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Projects.ListProjectsUsers returned %+v, want %+v", projects, want)
	}
}

func TestListProjectsUsersByName(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/", func(w http.ResponseWriter, r *http.Request) {
		testURL(t, r, "/api/v4/projects/namespace%2Fname/users?page=2&per_page=3&search=query")
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectUserOptions{
		ListOptions: ListOptions{Page: 2, PerPage: 3},
		Search:      Ptr("query"),
	}

	projects, _, err := client.Projects.ListProjectsUsers("namespace/name", opt)
	if err != nil {
		t.Errorf("Projects.ListProjectsUsers returned error: %v", err)
	}

	want := []*ProjectUser{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Projects.ListProjectsUsers returned %+v, want %+v", projects, want)
	}
}

func TestListProjectsGroupsByID(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/", func(w http.ResponseWriter, r *http.Request) {
		testURL(t, r, "/api/v4/projects/1/groups?page=2&per_page=3&search=query")
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectGroupOptions{
		ListOptions: ListOptions{Page: 2, PerPage: 3},
		Search:      Ptr("query"),
	}

	groups, _, err := client.Projects.ListProjectsGroups(1, opt)
	if err != nil {
		t.Errorf("Projects.ListProjectsGroups returned error: %v", err)
	}

	want := []*ProjectGroup{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, groups) {
		t.Errorf("Projects.ListProjectsGroups returned %+v, want %+v", groups, want)
	}
}

func TestListProjectsGroupsByName(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/", func(w http.ResponseWriter, r *http.Request) {
		testURL(t, r, "/api/v4/projects/namespace%2Fname/groups?page=2&per_page=3&search=query")
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectGroupOptions{
		ListOptions: ListOptions{Page: 2, PerPage: 3},
		Search:      Ptr("query"),
	}

	groups, _, err := client.Projects.ListProjectsGroups("namespace/name", opt)
	if err != nil {
		t.Errorf("Projects.ListProjectsGroups returned error: %v", err)
	}

	want := []*ProjectGroup{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, groups) {
		t.Errorf("Projects.ListProjectsGroups returned %+v, want %+v", groups, want)
	}
}

func TestListOwnedProjects(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{
		ListOptions: ListOptions{Page: 2, PerPage: 3},
		Archived:    Ptr(true),
		OrderBy:     Ptr("name"),
		Sort:        Ptr("asc"),
		Search:      Ptr("query"),
		Simple:      Ptr(true),
		Owned:       Ptr(true),
		Visibility:  Ptr(PublicVisibility),
	}

	projects, _, err := client.Projects.ListProjects(opt)
	if err != nil {
		t.Errorf("Projects.ListOwnedProjects returned error: %v", err)
	}

	want := []*Project{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Projects.ListOwnedProjects returned %+v, want %+v", projects, want)
	}
}

func TestEditProject(t *testing.T) {
	mux, client := setup(t)

	var developerAccessLevel AccessControlValue = "developer"
	opt := &EditProjectOptions{
		CIRestrictPipelineCancellationRole: Ptr(developerAccessLevel),
	}

	// Store whether we've set the restrict value in our edit properly
	restrictValueSet := false

	mux.HandleFunc("/api/v4/projects/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)

		// Check that our request properly included ci_restrict_pipeline_cancellation_role
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Unable to read body properly. Error: %v", err)
		}

		// Set the value to check if our value is included
		restrictValueSet = strings.Contains(string(body), "ci_restrict_pipeline_cancellation_role")

		// Print the start of the mock example from https://docs.gitlab.com/ee/api/projects.html#edit-project
		// including the attribute we edited
		fmt.Fprint(w, `
		{
			"id": 1,
			"description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			"description_html": "<p data-sourcepos=\"1:1-1:56\" dir=\"auto\">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>",
			"default_branch": "main",
			"visibility": "private",
			"ssh_url_to_repo": "git@example.com:diaspora/diaspora-project-site.git",
			"http_url_to_repo": "http://example.com/diaspora/diaspora-project-site.git",
			"web_url": "http://example.com/diaspora/diaspora-project-site",
			"readme_url": "http://example.com/diaspora/diaspora-project-site/blob/main/README.md",
			"ci_restrict_pipeline_cancellation_role": "developer"
		}`)
	})

	project, resp, err := client.Projects.EditProject(1, opt)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, true, restrictValueSet)
	assert.Equal(t, developerAccessLevel, project.CIRestrictPipelineCancellationRole)
}

func TestListStarredProjects(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{
		ListOptions: ListOptions{Page: 2, PerPage: 3},
		Archived:    Ptr(true),
		OrderBy:     Ptr("name"),
		Sort:        Ptr("asc"),
		Search:      Ptr("query"),
		Simple:      Ptr(true),
		Starred:     Ptr(true),
		Visibility:  Ptr(PublicVisibility),
	}

	projects, _, err := client.Projects.ListProjects(opt)
	if err != nil {
		t.Errorf("Projects.ListStarredProjects returned error: %v", err)
	}

	want := []*Project{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Projects.ListStarredProjects returned %+v, want %+v", projects, want)
	}
}

func TestGetProjectByID(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": 1,
			"container_registry_enabled": true,
			"container_expiration_policy": {
			  "cadence": "7d",
			  "enabled": false,
			  "keep_n": null,
			  "older_than": null,
			  "name_regex_delete": null,
			  "name_regex_keep": null,
			  "next_run_at": "2020-01-07T21:42:58.658Z"
			},
			"ci_forward_deployment_enabled": true,
			"ci_forward_deployment_rollback_allowed": true,
			"ci_restrict_pipeline_cancellation_role": "developer",
			"packages_enabled": false,
			"build_coverage_regex": "Total.*([0-9]{1,3})%"
		  }`)
	})

	wantTimestamp := time.Date(2020, 0o1, 0o7, 21, 42, 58, 658000000, time.UTC)
	want := &Project{
		ID:                       1,
		ContainerRegistryEnabled: true,
		ContainerExpirationPolicy: &ContainerExpirationPolicy{
			Cadence:   "7d",
			NextRunAt: &wantTimestamp,
		},
		PackagesEnabled:                    false,
		BuildCoverageRegex:                 `Total.*([0-9]{1,3})%`,
		CIForwardDeploymentEnabled:         true,
		CIForwardDeploymentRollbackAllowed: true,
		CIRestrictPipelineCancellationRole: "developer",
	}

	project, _, err := client.Projects.GetProject(1, nil)
	if err != nil {
		t.Fatalf("Projects.GetProject returns an error: %v", err)
	}

	if !reflect.DeepEqual(want, project) {
		t.Errorf("Projects.GetProject returned %+v, want %+v", project, want)
	}
}

func TestGetProjectByName(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/", func(w http.ResponseWriter, r *http.Request) {
		testURL(t, r, "/api/v4/projects/namespace%2Fname")
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &Project{ID: 1}

	project, _, err := client.Projects.GetProject("namespace/name", nil)
	if err != nil {
		t.Fatalf("Projects.GetProject returns an error: %v", err)
	}

	if !reflect.DeepEqual(want, project) {
		t.Errorf("Projects.GetProject returned %+v, want %+v", project, want)
	}
}

func TestGetProjectWithOptions(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id":1,
			"statistics": {
				"commit_count": 37,
				"storage_size": 1038090,
				"repository_size": 1038090,
				"wiki_size": 10,
				"lfs_objects_size": 0,
				"job_artifacts_size": 0,
				"pipeline_artifacts_size": 0,
				"packages_size": 238906167,
				"snippets_size": 146800,
				"uploads_size": 6523619,
				"container_registry_size": 284453
			}}`)
	})
	want := &Project{ID: 1, Statistics: &Statistics{
		CommitCount:           37,
		StorageSize:           1038090,
		RepositorySize:        1038090,
		WikiSize:              10,
		LFSObjectsSize:        0,
		JobArtifactsSize:      0,
		PipelineArtifactsSize: 0,
		PackagesSize:          238906167,
		SnippetsSize:          146800,
		UploadsSize:           6523619,
		ContainerRegistrySize: 284453,
	}}

	project, _, err := client.Projects.GetProject(1, &GetProjectOptions{Statistics: Ptr(true)})
	if err != nil {
		t.Fatalf("Projects.GetProject returns an error: %v", err)
	}

	if !reflect.DeepEqual(want, project) {
		t.Errorf("Projects.GetProject returned %+v, want %+v", project, want)
	}
}

func TestCreateProject(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1}`)
	})

	opt := &CreateProjectOptions{
		Name:        Ptr("n"),
		MergeMethod: Ptr(RebaseMerge),
	}

	project, _, err := client.Projects.CreateProject(opt)
	if err != nil {
		t.Errorf("Projects.CreateProject returned error: %v", err)
	}

	want := &Project{ID: 1}
	if !reflect.DeepEqual(want, project) {
		t.Errorf("Projects.CreateProject returned %+v, want %+v", project, want)
	}
}

func TestUploadFile(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/uploads", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		if !strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data;") {
			t.Fatalf("Projects.UploadFile request content-type %+v want multipart/form-data;", r.Header.Get("Content-Type"))
		}
		if r.ContentLength == -1 {
			t.Fatalf("Projects.UploadFile request content-length is -1")
		}
		fmt.Fprint(w, `{
		  "alt": "dk",
			"url": "/uploads/66dbcd21ec5d24ed6ea225176098d52b/dk.md",
			"markdown": "![dk](/uploads/66dbcd21ec5d24ed6ea225176098d52b/dk.png)"
		}`)
	})

	want := &ProjectFile{
		Alt:      "dk",
		URL:      "/uploads/66dbcd21ec5d24ed6ea225176098d52b/dk.md",
		Markdown: "![dk](/uploads/66dbcd21ec5d24ed6ea225176098d52b/dk.png)",
	}

	file := bytes.NewBufferString("dummy")
	projectFile, _, err := client.Projects.UploadFile(1, file, "test.txt")
	if err != nil {
		t.Fatalf("Projects.UploadFile returns an error: %v", err)
	}

	if !reflect.DeepEqual(want, projectFile) {
		t.Errorf("Projects.UploadFile returned %+v, want %+v", projectFile, want)
	}
}

func TestUploadFile_Retry(t *testing.T) {
	mux, client := setup(t)

	tf, _ := os.CreateTemp(os.TempDir(), "test")
	defer os.Remove(tf.Name())

	isFirstRequest := true
	mux.HandleFunc("/api/v4/projects/1/uploads", func(w http.ResponseWriter, r *http.Request) {
		if isFirstRequest {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			isFirstRequest = false
			return
		}
		if !strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data;") {
			t.Fatalf("Projects.UploadFile request content-type %+v want multipart/form-data;", r.Header.Get("Content-Type"))
		}
		if r.ContentLength == -1 {
			t.Fatalf("Projects.UploadFile request content-length is -1")
		}
		fmt.Fprint(w, `{
                  "alt": "dk",
                    "url": "/uploads/66dbcd21ec5d24ed6ea225176098d52b/dk.md",
                    "markdown": "![dk](/uploads/66dbcd21ec5d24ed6ea225176098d52b/dk.png)"
                }`)
	})

	want := &ProjectFile{
		Alt:      "dk",
		URL:      "/uploads/66dbcd21ec5d24ed6ea225176098d52b/dk.md",
		Markdown: "![dk](/uploads/66dbcd21ec5d24ed6ea225176098d52b/dk.png)",
	}

	file := bytes.NewBufferString("dummy")
	projectFile, _, err := client.Projects.UploadFile(1, file, "test.txt")
	if err != nil {
		t.Fatalf("Projects.UploadFile returns an error: %v", err)
	}

	if !reflect.DeepEqual(want, projectFile) {
		t.Errorf("Projects.UploadFile returned %+v, want %+v", projectFile, want)
	}
}

func TestUploadAvatar(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		if !strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data;") {
			t.Fatalf("Projects.UploadAvatar request content-type %+v want multipart/form-data;", r.Header.Get("Content-Type"))
		}
		if r.ContentLength == -1 {
			t.Fatalf("Projects.UploadAvatar request content-length is -1")
		}
		fmt.Fprint(w, `{}`)
	})

	avatar := new(bytes.Buffer)
	_, _, err := client.Projects.UploadAvatar(1, avatar, "avatar.png")
	if err != nil {
		t.Fatalf("Projects.UploadAvatar returns an error: %v", err)
	}
}

func TestUploadAvatar_Retry(t *testing.T) {
	mux, client := setup(t)

	isFirstRequest := true
	mux.HandleFunc("/api/v4/projects/1", func(w http.ResponseWriter, r *http.Request) {
		if isFirstRequest {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			isFirstRequest = false
			return
		}
		if !strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data;") {
			t.Fatalf("Projects.UploadAvatar request content-type %+v want multipart/form-data;", r.Header.Get("Content-Type"))
		}
		if r.ContentLength == -1 {
			t.Fatalf("Projects.UploadAvatar request content-length is -1")
		}
		fmt.Fprint(w, `{}`)
	})

	avatar := new(bytes.Buffer)
	_, _, err := client.Projects.UploadAvatar(1, avatar, "avatar.png")
	if err != nil {
		t.Fatalf("Projects.UploadAvatar returns an error: %v", err)
	}
}

func TestListProjectForks(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/", func(w http.ResponseWriter, r *http.Request) {
		testURL(t, r, "/api/v4/projects/namespace%2Fname/forks?archived=true&order_by=name&page=2&per_page=3&search=query&simple=true&sort=asc&visibility=public")
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{}
	opt.ListOptions = ListOptions{Page: 2, PerPage: 3}
	opt.Archived = Ptr(true)
	opt.OrderBy = Ptr("name")
	opt.Sort = Ptr("asc")
	opt.Search = Ptr("query")
	opt.Simple = Ptr(true)
	opt.Visibility = Ptr(PublicVisibility)

	projects, _, err := client.Projects.ListProjectForks("namespace/name", opt)
	if err != nil {
		t.Errorf("Projects.ListProjectForks returned error: %v", err)
	}

	want := []*Project{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Projects.ListProjects returned %+v, want %+v", projects, want)
	}
}

func TestShareProjectWithGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/share", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	opt := &ShareWithGroupOptions{
		GroupID:     Ptr(1),
		GroupAccess: Ptr(AccessLevelValue(50)),
	}

	_, err := client.Projects.ShareProjectWithGroup(1, opt)
	if err != nil {
		t.Errorf("Projects.ShareProjectWithGroup returned error: %v", err)
	}
}

func TestDeleteSharedProjectFromGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/share/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Projects.DeleteSharedProjectFromGroup(1, 2)
	if err != nil {
		t.Errorf("Projects.DeleteSharedProjectFromGroup returned error: %v", err)
	}
}

func TestGetApprovalConfiguration(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/approvals", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"approvers": [],
			"approver_groups": [],
			"approvals_before_merge": 3,
			"reset_approvals_on_push": false,
			"disable_overriding_approvers_per_merge_request": false,
			"merge_requests_author_approval": true,
			"merge_requests_disable_committers_approval": true,
			"require_password_to_approve": true
		}`)
	})

	approvals, _, err := client.Projects.GetApprovalConfiguration(1)
	if err != nil {
		t.Errorf("Projects.GetApprovalConfiguration returned error: %v", err)
	}

	want := &ProjectApprovals{
		Approvers:            []*MergeRequestApproverUser{},
		ApproverGroups:       []*MergeRequestApproverGroup{},
		ApprovalsBeforeMerge: 3,
		ResetApprovalsOnPush: false,
		DisableOverridingApproversPerMergeRequest: false,
		MergeRequestsAuthorApproval:               true,
		MergeRequestsDisableCommittersApproval:    true,
		RequirePasswordToApprove:                  true,
	}

	if !reflect.DeepEqual(want, approvals) {
		t.Errorf("Projects.GetApprovalConfiguration returned %+v, want %+v", approvals, want)
	}
}

func TestChangeApprovalConfiguration(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/approvals", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"approvals_before_merge":3}`)
		fmt.Fprint(w, `{
			"approvers": [],
			"approver_groups": [],
			"approvals_before_merge": 3,
			"reset_approvals_on_push": false,
			"disable_overriding_approvers_per_merge_request": false,
			"merge_requests_author_approval": true,
			"merge_requests_disable_committers_approval": true,
			"require_password_to_approve": true
		}`)
	})

	opt := &ChangeApprovalConfigurationOptions{
		ApprovalsBeforeMerge: Ptr(3),
	}

	approvals, _, err := client.Projects.ChangeApprovalConfiguration(1, opt)
	if err != nil {
		t.Errorf("Projects.ChangeApprovalConfigurationOptions returned error: %v", err)
	}

	want := &ProjectApprovals{
		Approvers:            []*MergeRequestApproverUser{},
		ApproverGroups:       []*MergeRequestApproverGroup{},
		ApprovalsBeforeMerge: 3,
		ResetApprovalsOnPush: false,
		DisableOverridingApproversPerMergeRequest: false,
		MergeRequestsAuthorApproval:               true,
		MergeRequestsDisableCommittersApproval:    true,
		RequirePasswordToApprove:                  true,
	}

	if !reflect.DeepEqual(want, approvals) {
		t.Errorf("Projects.ChangeApprovalConfigurationOptions  returned %+v, want %+v", approvals, want)
	}
}

func TestChangeAllowedApprovers(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/approvers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testBody(t, r, `{"approver_group_ids":[1],"approver_ids":[2]}`)
		fmt.Fprint(w, `{
			"approver_groups": [{"group":{"id":1}}],
			"approvers": [{"user":{"id":2}}]
		}`)
	})

	opt := &ChangeAllowedApproversOptions{
		ApproverGroupIDs: &[]int{1},
		ApproverIDs:      &[]int{2},
	}

	approvals, _, err := client.Projects.ChangeAllowedApprovers(1, opt)
	if err != nil {
		t.Errorf("Projects.ChangeApproversConfigurationOptions returned error: %v", err)
	}

	want := &ProjectApprovals{
		ApproverGroups: []*MergeRequestApproverGroup{
			{
				Group: struct {
					ID                   int    `json:"id"`
					Name                 string `json:"name"`
					Path                 string `json:"path"`
					Description          string `json:"description"`
					Visibility           string `json:"visibility"`
					AvatarURL            string `json:"avatar_url"`
					WebURL               string `json:"web_url"`
					FullName             string `json:"full_name"`
					FullPath             string `json:"full_path"`
					LFSEnabled           bool   `json:"lfs_enabled"`
					RequestAccessEnabled bool   `json:"request_access_enabled"`
				}{
					ID: 1,
				},
			},
		},
		Approvers: []*MergeRequestApproverUser{
			{
				User: &BasicUser{
					ID: 2,
				},
			},
		},
	}

	if !reflect.DeepEqual(want, approvals) {
		t.Errorf("Projects.ChangeAllowedApprovers returned %+v, want %+v", approvals, want)
	}
}

func TestForkProject(t *testing.T) {
	mux, client := setup(t)

	namespaceID := 42
	name := "myreponame"
	path := "myrepopath"

	mux.HandleFunc("/api/v4/projects/1/fork", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, fmt.Sprintf(`{"name":"%s","namespace_id":%d,"path":"%s"}`, name, namespaceID, path))
		fmt.Fprint(w, `{"id":2}`)
	})

	project, _, err := client.Projects.ForkProject(1, &ForkProjectOptions{
		NamespaceID: Ptr(namespaceID),
		Name:        Ptr(name),
		Path:        Ptr(path),
	})
	if err != nil {
		t.Errorf("Projects.ForkProject returned error: %v", err)
	}

	want := &Project{ID: 2}
	if !reflect.DeepEqual(want, project) {
		t.Errorf("Projects.ForProject returned %+v, want %+v", project, want)
	}
}

func TestGetProjectApprovalRules(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/approval_rules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "security",
				"rule_type": "regular",
				"eligible_approvers": [
					{
						"id": 5,
						"name": "John Doe",
						"username": "jdoe",
						"state": "active",
						"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
						"web_url": "http://localhost/jdoe"
					},
					{
						"id": 50,
						"name": "Group Member 1",
						"username": "group_member_1",
						"state": "active",
						"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
						"web_url": "http://localhost/group_member_1"
					}
				],
				"approvals_required": 3,
				"users": [
					{
						"id": 5,
						"name": "John Doe",
						"username": "jdoe",
						"state": "active",
						"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
						"web_url": "http://localhost/jdoe"
					}
				],
				"groups": [
					{
						"id": 5,
						"name": "group1",
						"path": "group1",
						"description": "",
						"visibility": "public",
						"lfs_enabled": false,
						"avatar_url": null,
						"web_url": "http://localhost/groups/group1",
						"request_access_enabled": false,
						"full_name": "group1",
						"full_path": "group1",
						"parent_id": null,
						"ldap_cn": null,
						"ldap_access": null
					}
				],
				"protected_branches": [
					  {
						"id": 1,
						"name": "master",
						"push_access_levels": [
						  {
							"access_level": 30,
							"access_level_description": "Developers + Maintainers"
						  }
						],
						"merge_access_levels": [
						  {
							"access_level": 30,
							"access_level_description": "Developers + Maintainers"
						  }
						],
						"unprotect_access_levels": [
						  {
							"access_level": 40,
							"access_level_description": "Maintainers"
						  }
						],
						"code_owner_approval_required": false
					  }
                ],
				"contains_hidden_groups": false
			}
		]`)
	})

	approvals, _, err := client.Projects.GetProjectApprovalRules(1, nil)
	if err != nil {
		t.Errorf("Projects.GetProjectApprovalRules returned error: %v", err)
	}

	want := []*ProjectApprovalRule{
		{
			ID:       1,
			Name:     "security",
			RuleType: "regular",
			EligibleApprovers: []*BasicUser{
				{
					ID:        5,
					Name:      "John Doe",
					Username:  "jdoe",
					State:     "active",
					AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
					WebURL:    "http://localhost/jdoe",
				},
				{
					ID:        50,
					Name:      "Group Member 1",
					Username:  "group_member_1",
					State:     "active",
					AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
					WebURL:    "http://localhost/group_member_1",
				},
			},
			ApprovalsRequired: 3,
			Users: []*BasicUser{
				{
					ID:        5,
					Name:      "John Doe",
					Username:  "jdoe",
					State:     "active",
					AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
					WebURL:    "http://localhost/jdoe",
				},
			},
			Groups: []*Group{
				{
					ID:                   5,
					Name:                 "group1",
					Path:                 "group1",
					Description:          "",
					Visibility:           PublicVisibility,
					LFSEnabled:           false,
					AvatarURL:            "",
					WebURL:               "http://localhost/groups/group1",
					RequestAccessEnabled: false,
					FullName:             "group1",
					FullPath:             "group1",
				},
			},
			ProtectedBranches: []*ProtectedBranch{
				{
					ID:   1,
					Name: "master",
					PushAccessLevels: []*BranchAccessDescription{
						{
							AccessLevel:            30,
							AccessLevelDescription: "Developers + Maintainers",
						},
					},
					MergeAccessLevels: []*BranchAccessDescription{
						{
							AccessLevel:            30,
							AccessLevelDescription: "Developers + Maintainers",
						},
					},
					UnprotectAccessLevels: []*BranchAccessDescription{
						{
							AccessLevel:            40,
							AccessLevelDescription: "Maintainers",
						},
					},
					AllowForcePush:            false,
					CodeOwnerApprovalRequired: false,
				},
			},
		},
	}

	if !reflect.DeepEqual(want, approvals) {
		t.Errorf("Projects.GetProjectApprovalRules returned %+v, want %+v", approvals, want)
	}
}

func TestGetProjectApprovalRule(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/approval_rules/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "security",
			"rule_type": "regular",
			"eligible_approvers": [
				{
					"id": 5,
					"name": "John Doe",
					"username": "jdoe",
					"state": "active",
					"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
					"web_url": "http://localhost/jdoe"
				},
				{
					"id": 50,
					"name": "Group Member 1",
					"username": "group_member_1",
					"state": "active",
					"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
					"web_url": "http://localhost/group_member_1"
				}
			],
			"approvals_required": 3,
			"users": [
				{
					"id": 5,
					"name": "John Doe",
					"username": "jdoe",
					"state": "active",
					"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
					"web_url": "http://localhost/jdoe"
				}
			],
			"groups": [
				{
					"id": 5,
					"name": "group1",
					"path": "group1",
					"description": "",
					"visibility": "public",
					"lfs_enabled": false,
					"avatar_url": null,
					"web_url": "http://localhost/groups/group1",
					"request_access_enabled": false,
					"full_name": "group1",
					"full_path": "group1",
					"parent_id": null,
					"ldap_cn": null,
					"ldap_access": null
				}
			],
			"protected_branches": [
					{
					"id": 1,
					"name": "master",
					"push_access_levels": [
						{
						"access_level": 30,
						"access_level_description": "Developers + Maintainers"
						}
					],
					"merge_access_levels": [
						{
						"access_level": 30,
						"access_level_description": "Developers + Maintainers"
						}
					],
					"unprotect_access_levels": [
						{
						"access_level": 40,
						"access_level_description": "Maintainers"
						}
					],
					"code_owner_approval_required": false
					}
			],
			"contains_hidden_groups": false
		}`)
	})

	approvals, _, err := client.Projects.GetProjectApprovalRule(1, 1)
	if err != nil {
		t.Errorf("Projects.GetProjectApprovalRule returned error: %v", err)
	}

	want := &ProjectApprovalRule{
		ID:       1,
		Name:     "security",
		RuleType: "regular",
		EligibleApprovers: []*BasicUser{
			{
				ID:        5,
				Name:      "John Doe",
				Username:  "jdoe",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
				WebURL:    "http://localhost/jdoe",
			},
			{
				ID:        50,
				Name:      "Group Member 1",
				Username:  "group_member_1",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
				WebURL:    "http://localhost/group_member_1",
			},
		},
		ApprovalsRequired: 3,
		Users: []*BasicUser{
			{
				ID:        5,
				Name:      "John Doe",
				Username:  "jdoe",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
				WebURL:    "http://localhost/jdoe",
			},
		},
		Groups: []*Group{
			{
				ID:                   5,
				Name:                 "group1",
				Path:                 "group1",
				Description:          "",
				Visibility:           PublicVisibility,
				LFSEnabled:           false,
				AvatarURL:            "",
				WebURL:               "http://localhost/groups/group1",
				RequestAccessEnabled: false,
				FullName:             "group1",
				FullPath:             "group1",
			},
		},
		ProtectedBranches: []*ProtectedBranch{
			{
				ID:   1,
				Name: "master",
				PushAccessLevels: []*BranchAccessDescription{
					{
						AccessLevel:            30,
						AccessLevelDescription: "Developers + Maintainers",
					},
				},
				MergeAccessLevels: []*BranchAccessDescription{
					{
						AccessLevel:            30,
						AccessLevelDescription: "Developers + Maintainers",
					},
				},
				UnprotectAccessLevels: []*BranchAccessDescription{
					{
						AccessLevel:            40,
						AccessLevelDescription: "Maintainers",
					},
				},
				AllowForcePush:            false,
				CodeOwnerApprovalRequired: false,
			},
		},
	}

	if !reflect.DeepEqual(want, approvals) {
		t.Errorf("Projects.GetProjectApprovalRule returned %+v, want %+v", approvals, want)
	}
}

func TestCreateProjectApprovalRule(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/approval_rules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "security",
			"rule_type": "regular",
			"eligible_approvers": [
				{
					"id": 5,
					"name": "John Doe",
					"username": "jdoe",
					"state": "active",
					"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
					"web_url": "http://localhost/jdoe"
				},
				{
					"id": 50,
					"name": "Group Member 1",
					"username": "group_member_1",
					"state": "active",
					"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
					"web_url": "http://localhost/group_member_1"
				}
			],
			"approvals_required": 3,
			"users": [
				{
					"id": 5,
					"name": "John Doe",
					"username": "jdoe",
					"state": "active",
					"avatar_url": "https://www.gravatar.com/avatar/0?s=80&d=identicon",
					"web_url": "http://localhost/jdoe"
				}
			],
			"groups": [
				{
					"id": 5,
					"name": "group1",
					"path": "group1",
					"description": "",
					"visibility": "public",
					"lfs_enabled": false,
					"avatar_url": null,
					"web_url": "http://localhost/groups/group1",
					"request_access_enabled": false,
					"full_name": "group1",
					"full_path": "group1",
					"parent_id": null,
					"ldap_cn": null,
					"ldap_access": null
				}
			],
			"protected_branches": [
				{
				  "id": 1,
				  "name": "master",
				  "push_access_levels": [
					{
					  "access_level": 30,
					  "access_level_description": "Developers + Maintainers"
					}
				  ],
				  "merge_access_levels": [
					{
					  "access_level": 30,
					  "access_level_description": "Developers + Maintainers"
					}
				  ],
				  "unprotect_access_levels": [
					{
					  "access_level": 40,
					  "access_level_description": "Maintainers"
					}
				  ],
				  "code_owner_approval_required": false
				}
			],
			"contains_hidden_groups": false
		}`)
	})

	opt := &CreateProjectLevelRuleOptions{
		Name:              Ptr("security"),
		ApprovalsRequired: Ptr(3),
		UserIDs:           &[]int{5, 50},
		GroupIDs:          &[]int{5},
		ReportType:        String("code_coverage"),
	}

	rule, _, err := client.Projects.CreateProjectApprovalRule(1, opt)
	if err != nil {
		t.Errorf("Projects.CreateProjectApprovalRule returned error: %v", err)
	}

	want := &ProjectApprovalRule{
		ID:       1,
		Name:     "security",
		RuleType: "regular",
		EligibleApprovers: []*BasicUser{
			{
				ID:        5,
				Name:      "John Doe",
				Username:  "jdoe",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
				WebURL:    "http://localhost/jdoe",
			},
			{
				ID:        50,
				Name:      "Group Member 1",
				Username:  "group_member_1",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
				WebURL:    "http://localhost/group_member_1",
			},
		},
		ApprovalsRequired: 3,
		Users: []*BasicUser{
			{
				ID:        5,
				Name:      "John Doe",
				Username:  "jdoe",
				State:     "active",
				AvatarURL: "https://www.gravatar.com/avatar/0?s=80&d=identicon",
				WebURL:    "http://localhost/jdoe",
			},
		},
		Groups: []*Group{
			{
				ID:                   5,
				Name:                 "group1",
				Path:                 "group1",
				Description:          "",
				Visibility:           PublicVisibility,
				LFSEnabled:           false,
				AvatarURL:            "",
				WebURL:               "http://localhost/groups/group1",
				RequestAccessEnabled: false,
				FullName:             "group1",
				FullPath:             "group1",
			},
		},
		ProtectedBranches: []*ProtectedBranch{
			{
				ID:   1,
				Name: "master",
				PushAccessLevels: []*BranchAccessDescription{
					{
						AccessLevel:            30,
						AccessLevelDescription: "Developers + Maintainers",
					},
				},
				MergeAccessLevels: []*BranchAccessDescription{
					{
						AccessLevel:            30,
						AccessLevelDescription: "Developers + Maintainers",
					},
				},
				UnprotectAccessLevels: []*BranchAccessDescription{
					{
						AccessLevel:            40,
						AccessLevelDescription: "Maintainers",
					},
				},
				AllowForcePush:            false,
				CodeOwnerApprovalRequired: false,
			},
		},
	}

	if !reflect.DeepEqual(want, rule) {
		t.Errorf("Projects.CreateProjectApprovalRule returned %+v, want %+v", rule, want)
	}
}

func TestGetProjectPullMirrorDetails(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/mirror/pull", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
		  "id": 101486,
		  "last_error": null,
		  "last_successful_update_at": "2020-01-06T17:32:02.823Z",
		  "last_update_at": "2020-01-06T17:32:02.823Z",
		  "last_update_started_at": "2020-01-06T17:31:55.864Z",
		  "update_status": "finished",
		  "url": "https://*****:*****@gitlab.com/gitlab-org/security/gitlab.git"
		}`)
	})

	pullMirror, _, err := client.Projects.GetProjectPullMirrorDetails(1)
	if err != nil {
		t.Errorf("Projects.GetProjectPullMirrorDetails returned error: %v", err)
	}

	wantLastSuccessfulUpdateAtTimestamp := time.Date(2020, 0o1, 0o6, 17, 32, 0o2, 823000000, time.UTC)
	wantLastUpdateAtTimestamp := time.Date(2020, 0o1, 0o6, 17, 32, 0o2, 823000000, time.UTC)
	wantLastUpdateStartedAtTimestamp := time.Date(2020, 0o1, 0o6, 17, 31, 55, 864000000, time.UTC)
	want := &ProjectPullMirrorDetails{
		ID:                     101486,
		LastError:              "",
		LastSuccessfulUpdateAt: &wantLastSuccessfulUpdateAtTimestamp,
		LastUpdateAt:           &wantLastUpdateAtTimestamp,
		LastUpdateStartedAt:    &wantLastUpdateStartedAtTimestamp,
		UpdateStatus:           "finished",
		URL:                    "https://*****:*****@gitlab.com/gitlab-org/security/gitlab.git",
	}

	if !reflect.DeepEqual(want, pullMirror) {
		t.Errorf("Projects.GetProjectPullMirrorDetails returned %+v, want %+v", pullMirror, want)
	}
}

func TestCreateProjectApprovalRuleEligibleApprovers(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/approval_rules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "Any name",
			"rule_type": "any_approver",
			"eligible_approvers": [],
			"approvals_required": 1,
			"users": [],
			"groups": [],
			"contains_hidden_groups": false,
			"protected_branches": []
		}`)
	})

	opt := &CreateProjectLevelRuleOptions{
		Name:              Ptr("Any name"),
		ApprovalsRequired: Ptr(1),
	}

	rule, _, err := client.Projects.CreateProjectApprovalRule(1, opt)
	if err != nil {
		t.Errorf("Projects.CreateProjectApprovalRule returned error: %v", err)
	}

	want := &ProjectApprovalRule{
		ID:                1,
		Name:              "Any name",
		RuleType:          "any_approver",
		EligibleApprovers: []*BasicUser{},
		ApprovalsRequired: 1,
		Users:             []*BasicUser{},
		Groups:            []*Group{},
		ProtectedBranches: []*ProtectedBranch{},
	}

	if !reflect.DeepEqual(want, rule) {
		t.Errorf("Projects.CreateProjectApprovalRule returned %+v, want %+v", rule, want)
	}
}

func TestProjectModelsOptionalMergeAttribute(t *testing.T) {
	// Create a `CreateProjectOptions` struct, ensure that merge attribute doesn't serialize
	jsonString, err := json.Marshal(&CreateProjectOptions{
		Name: Ptr("testProject"),
	})
	if err != nil {
		t.Fatal("Failed to marshal object", err)
	}
	assert.False(t, strings.Contains(string(jsonString), "only_allow_merge_if_all_status_checks_passed"))

	// Test the same thing but for `EditProjectOptions` struct
	jsonString, err = json.Marshal(&EditProjectOptions{
		Name: Ptr("testProject"),
	})
	if err != nil {
		t.Fatal("Failed to marshal object", err)
	}
	assert.False(t, strings.Contains(string(jsonString), "only_allow_merge_if_all_status_checks_passed"))
}

// Test that the "CustomWebhookTemplate" serializes properly
func TestProjectAddWebhook_CustomTemplateStuff(t *testing.T) {
	mux, client := setup(t)
	customWebhookSet := false
	authValueSet := false

	mux.HandleFunc("/api/v4/projects/1/hooks",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			w.WriteHeader(http.StatusCreated)

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("Unable to read body properly. Error: %v", err)
			}
			customWebhookSet = strings.Contains(string(body), "custom_webhook_template")
			authValueSet = strings.Contains(string(body), `"value":"stuff"`)

			fmt.Fprint(w, `{
				"custom_webhook_template": "testValue",
				"custom_headers": [
					{
						"key": "Authorization"
					},
					{
						"key": "Favorite-Pet"
					}
				]
			}`)
		},
	)

	hook, resp, err := client.Projects.AddProjectHook(1, &AddProjectHookOptions{
		CustomWebhookTemplate: Ptr(`{"example":"{{object_kind}}"}`),
		CustomHeaders: &[]*HookCustomHeader{
			{
				Key:   "Authorization",
				Value: "stuff",
			},
			{
				Key:   "Favorite-Pet",
				Value: "Cats",
			},
		},
	})

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, true, customWebhookSet)
	assert.Equal(t, true, authValueSet)
	assert.Equal(t, "testValue", hook.CustomWebhookTemplate)
	assert.Equal(t, 2, len(hook.CustomHeaders))
}

// Test that the "CustomWebhookTemplate" serializes properly when editing
func TestProjectEditWebhook_CustomTemplateStuff(t *testing.T) {
	mux, client := setup(t)
	customWebhookSet := false
	authValueSet := false

	mux.HandleFunc("/api/v4/projects/1/hooks/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			w.WriteHeader(http.StatusOK)

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("Unable to read body properly. Error: %v", err)
			}
			customWebhookSet = strings.Contains(string(body), "custom_webhook_template")
			authValueSet = strings.Contains(string(body), `"value":"stuff"`)

			fmt.Fprint(w, `{
				"custom_webhook_template": "testValue",
				"custom_headers": [
					{
						"key": "Authorization"
					},
					{
						"key": "Favorite-Pet"
					}
				]}`)
		},
	)

	hook, resp, err := client.Projects.EditProjectHook(1, 1, &EditProjectHookOptions{
		CustomWebhookTemplate: Ptr(`{"example":"{{object_kind}}"}`),
		CustomHeaders: &[]*HookCustomHeader{
			{
				Key:   "Authorization",
				Value: "stuff",
			},
			{
				Key:   "Favorite-Pet",
				Value: "Cats",
			},
		},
	})

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, true, customWebhookSet)
	assert.Equal(t, true, authValueSet)
	assert.Equal(t, "testValue", hook.CustomWebhookTemplate)
	assert.Equal(t, 2, len(hook.CustomHeaders))
}

func TestGetProjectPushRules(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/push_rule", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": 1,
			"commit_message_regex": "Fixes \\d+\\..*",
			"commit_message_negative_regex": "ssh\\:\\/\\/",
			"branch_name_regex": "(feat|fix)\\/*",
			"deny_delete_tag": false,
			"member_check": false,
			"prevent_secrets": false,
			"author_email_regex": "@company.com$",
			"file_name_regex": "(jar|exe)$",
			"max_file_size": 5,
			"commit_committer_check": false,
			"commit_committer_name_check": false,
			"reject_unsigned_commits": false
		  }`)
	})

	rule, _, err := client.Projects.GetProjectPushRules(1)
	if err != nil {
		t.Errorf("Projects.GetProjectPushRules returned error: %v", err)
	}

	want := &ProjectPushRules{
		ID:                         1,
		CommitMessageRegex:         "Fixes \\d+\\..*",
		CommitMessageNegativeRegex: "ssh\\:\\/\\/",
		BranchNameRegex:            "(feat|fix)\\/*",
		DenyDeleteTag:              false,
		MemberCheck:                false,
		PreventSecrets:             false,
		AuthorEmailRegex:           "@company.com$",
		FileNameRegex:              "(jar|exe)$",
		MaxFileSize:                5,
		CommitCommitterCheck:       false,
		CommitCommitterNameCheck:   false,
		RejectUnsignedCommits:      false,
	}

	if !reflect.DeepEqual(want, rule) {
		t.Errorf("Projects.GetProjectPushRules returned %+v, want %+v", rule, want)
	}
}

func TestAddProjectPushRules(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/push_rule", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"id": 1,
			"commit_message_regex": "Fixes \\d+\\..*",
			"commit_message_negative_regex": "ssh\\:\\/\\/",
			"branch_name_regex": "(feat|fix)\\/*",
			"deny_delete_tag": false,
			"member_check": false,
			"prevent_secrets": false,
			"author_email_regex": "@company.com$",
			"file_name_regex": "(jar|exe)$",
			"max_file_size": 5,
			"commit_committer_check": false,
			"commit_committer_name_check": false,
			"reject_unsigned_commits": false
		  }`)
	})

	opt := &AddProjectPushRuleOptions{
		CommitMessageRegex:         Ptr("Fixes \\d+\\..*"),
		CommitMessageNegativeRegex: Ptr("ssh\\:\\/\\/"),
		BranchNameRegex:            Ptr("(feat|fix)\\/*"),
		DenyDeleteTag:              Ptr(false),
		MemberCheck:                Ptr(false),
		PreventSecrets:             Ptr(false),
		AuthorEmailRegex:           Ptr("@company.com$"),
		FileNameRegex:              Ptr("(jar|exe)$"),
		MaxFileSize:                Ptr(5),
		CommitCommitterCheck:       Ptr(false),
		CommitCommitterNameCheck:   Ptr(false),
		RejectUnsignedCommits:      Ptr(false),
	}

	rule, _, err := client.Projects.AddProjectPushRule(1, opt)
	if err != nil {
		t.Errorf("Projects.AddProjectPushRule returned error: %v", err)
	}

	want := &ProjectPushRules{
		ID:                         1,
		CommitMessageRegex:         "Fixes \\d+\\..*",
		CommitMessageNegativeRegex: "ssh\\:\\/\\/",
		BranchNameRegex:            "(feat|fix)\\/*",
		DenyDeleteTag:              false,
		MemberCheck:                false,
		PreventSecrets:             false,
		AuthorEmailRegex:           "@company.com$",
		FileNameRegex:              "(jar|exe)$",
		MaxFileSize:                5,
		CommitCommitterCheck:       false,
		CommitCommitterNameCheck:   false,
		RejectUnsignedCommits:      false,
	}

	if !reflect.DeepEqual(want, rule) {
		t.Errorf("Projects.AddProjectPushRule returned %+v, want %+v", rule, want)
	}
}

func TestEditProjectPushRules(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/push_rule", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{
			"id": 1,
			"commit_message_regex": "Fixes \\d+\\..*",
			"commit_message_negative_regex": "ssh\\:\\/\\/",
			"branch_name_regex": "(feat|fix)\\/*",
			"deny_delete_tag": false,
			"member_check": false,
			"prevent_secrets": false,
			"author_email_regex": "@company.com$",
			"file_name_regex": "(jar|exe)$",
			"max_file_size": 5,
			"commit_committer_check": false,
			"commit_committer_name_check": false,
			"reject_unsigned_commits": false
		  }`)
	})

	opt := &EditProjectPushRuleOptions{
		CommitMessageRegex:         Ptr("Fixes \\d+\\..*"),
		CommitMessageNegativeRegex: Ptr("ssh\\:\\/\\/"),
		BranchNameRegex:            Ptr("(feat|fix)\\/*"),
		DenyDeleteTag:              Ptr(false),
		MemberCheck:                Ptr(false),
		PreventSecrets:             Ptr(false),
		AuthorEmailRegex:           Ptr("@company.com$"),
		FileNameRegex:              Ptr("(jar|exe)$"),
		MaxFileSize:                Ptr(5),
		CommitCommitterCheck:       Ptr(false),
		CommitCommitterNameCheck:   Ptr(false),
		RejectUnsignedCommits:      Ptr(false),
	}

	rule, _, err := client.Projects.EditProjectPushRule(1, opt)
	if err != nil {
		t.Errorf("Projects.EditProjectPushRule returned error: %v", err)
	}

	want := &ProjectPushRules{
		ID:                         1,
		CommitMessageRegex:         "Fixes \\d+\\..*",
		CommitMessageNegativeRegex: "ssh\\:\\/\\/",
		BranchNameRegex:            "(feat|fix)\\/*",
		DenyDeleteTag:              false,
		MemberCheck:                false,
		PreventSecrets:             false,
		AuthorEmailRegex:           "@company.com$",
		FileNameRegex:              "(jar|exe)$",
		MaxFileSize:                5,
		CommitCommitterCheck:       false,
		CommitCommitterNameCheck:   false,
		RejectUnsignedCommits:      false,
	}

	if !reflect.DeepEqual(want, rule) {
		t.Errorf("Projects.EditProjectPushRule returned %+v, want %+v", rule, want)
	}
}

func TestGetProjectWebhookHeader(t *testing.T) {
	mux, client := setup(t)

	// Removed most of the arguments to keep test slim
	mux.HandleFunc("/api/v4/projects/1/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": 1,
			"custom_webhook_template": "{\"event\":\"{{object_kind}}\"}",
			"custom_headers": [
			  {
				"key": "Authorization"
			  },
			  {
				"key": "OtherKey"
			  }
			]
		  }`)
	})

	hook, _, err := client.Projects.GetProjectHook(1, 1)
	if err != nil {
		t.Errorf("Projects.GetProjectHook returned error: %v", err)
	}

	want := &ProjectHook{
		ID:                    1,
		CustomWebhookTemplate: "{\"event\":\"{{object_kind}}\"}",
		CustomHeaders: []*HookCustomHeader{
			{
				Key: "Authorization",
			},
			{
				Key: "OtherKey",
			},
		},
	}

	if !reflect.DeepEqual(want, hook) {
		t.Errorf("Projects.GetProjectHook returned %+v, want %+v", hook, want)
	}
}

func TestSetProjectWebhookHeader(t *testing.T) {
	mux, client := setup(t)
	var bodyJson map[string]interface{}

	// Removed most of the arguments to keep test slim
	mux.HandleFunc("/api/v4/projects/1/hooks/1/custom_headers/Authorization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		w.WriteHeader(http.StatusNoContent)

		// validate that the `value` body is sent properly
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Unable to read body properly. Error: %v", err)
		}

		// Unmarshal the body into JSON so we can check it
		_ = json.Unmarshal(body, &bodyJson)

		fmt.Fprint(w, ``)
	})

	req, err := client.Projects.SetProjectCustomHeader(1, 1, "Authorization", &SetHookCustomHeaderOptions{Value: Ptr("testValue")})
	if err != nil {
		t.Errorf("Projects.SetProjectCustomHeader returned error: %v", err)
	}

	assert.Equal(t, bodyJson["value"], "testValue")
	assert.Equal(t, http.StatusNoContent, req.StatusCode)
}

func TestDeleteProjectWebhookHeader(t *testing.T) {
	mux, client := setup(t)

	// Removed most of the arguments to keep test slim
	mux.HandleFunc("/api/v4/projects/1/hooks/1/custom_headers/Authorization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, ``)
	})

	req, err := client.Projects.DeleteProjectCustomHeader(1, 1, "Authorization")
	if err != nil {
		t.Errorf("Projects.DeleteProjectCustomHeader returned error: %v", err)
	}

	assert.Equal(t, http.StatusNoContent, req.StatusCode)
}
