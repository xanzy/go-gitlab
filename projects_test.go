package gitlab

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestListProjects(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{
		ListOptions: ListOptions{2, 3},
		Archived:    Bool(true),
		OrderBy:     String("name"),
		Sort:        String("asc"),
		Search:      String("query"),
		Simple:      Bool(true),
		Visibility:  Visibility(PublicVisibility),
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
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/users/1/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{
		ListOptions: ListOptions{2, 3},
		Archived:    Bool(true),
		OrderBy:     String("name"),
		Sort:        String("asc"),
		Search:      String("query"),
		Simple:      Bool(true),
		Visibility:  Visibility(PublicVisibility),
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

func TestListProjectsUsersByID(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/", func(w http.ResponseWriter, r *http.Request) {
		testURL(t, r, "/api/v4/projects/1/users?page=2&per_page=3&search=query")
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectUserOptions{
		ListOptions: ListOptions{2, 3},
		Search:      String("query"),
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
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/", func(w http.ResponseWriter, r *http.Request) {
		testURL(t, r, "/api/v4/projects/namespace%2Fname/users?page=2&per_page=3&search=query")
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectUserOptions{
		ListOptions: ListOptions{2, 3},
		Search:      String("query"),
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

func TestListOwnedProjects(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{
		ListOptions: ListOptions{2, 3},
		Archived:    Bool(true),
		OrderBy:     String("name"),
		Sort:        String("asc"),
		Search:      String("query"),
		Simple:      Bool(true),
		Owned:       Bool(true),
		Visibility:  Visibility(PublicVisibility),
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

func TestListStarredProjects(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{
		ListOptions: ListOptions{2, 3},
		Archived:    Bool(true),
		OrderBy:     String("name"),
		Sort:        String("asc"),
		Search:      String("query"),
		Simple:      Bool(true),
		Starred:     Bool(true),
		Visibility:  Visibility(PublicVisibility),
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
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &Project{ID: 1}

	project, _, err := client.Projects.GetProject(1, nil)
	if err != nil {
		t.Fatalf("Projects.GetProject returns an error: %v", err)
	}

	if !reflect.DeepEqual(want, project) {
		t.Errorf("Projects.GetProject returned %+v, want %+v", project, want)
	}
}

func TestGetProjectByName(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/", func(w http.ResponseWriter, r *http.Request) {
		testURL(t, r, "/api/v4/projects/namespace%2Fname")
		testMethod(t, r, "GET")
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
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
			"id":1,
			"statistics": {
				"commit_count": 37,
				"storage_size": 1038090,
				"repository_size": 1038090,
				"lfs_objects_size": 0,
				"job_artifacts_size": 0
			}}`)
	})
	want := &Project{ID: 1, Statistics: &ProjectStatistics{
		CommitCount: 37,
		StorageStatistics: StorageStatistics{
			StorageSize:      1038090,
			RepositorySize:   1038090,
			LfsObjectsSize:   0,
			JobArtifactsSize: 0,
		},
	}}

	project, _, err := client.Projects.GetProject(1, &GetProjectOptions{Statistics: Bool(true)})
	if err != nil {
		t.Fatalf("Projects.GetProject returns an error: %v", err)
	}

	if !reflect.DeepEqual(want, project) {
		t.Errorf("Projects.GetProject returned %+v, want %+v", project, want)
	}
}

func TestCreateProject(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"id":1}`)
	})

	opt := &CreateProjectOptions{
		Name:        String("n"),
		MergeMethod: MergeMethod(RebaseMerge),
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
	mux, server, client := setup()
	defer teardown(server)

	tf, _ := ioutil.TempFile(os.TempDir(), "test")
	defer os.Remove(tf.Name())

	mux.HandleFunc("/api/v4/projects/1/uploads", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		if false == strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data;") {
			t.Fatalf("Prokects.UploadFile request content-type %+v want multipart/form-data;", r.Header.Get("Content-Type"))
		}
		if r.ContentLength == -1 {
			t.Fatalf("Prokects.UploadFile request content-length is -1")
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

	file, _, err := client.Projects.UploadFile(1, tf.Name())

	if err != nil {
		t.Fatalf("Prokects.UploadFile returns an error: %v", err)
	}

	if !reflect.DeepEqual(want, file) {
		t.Errorf("Prokects.UploadFile returned %+v, want %+v", file, want)
	}
}

func TestListProjectForks(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/", func(w http.ResponseWriter, r *http.Request) {
		testURL(t, r, "/api/v4/projects/namespace%2Fname/forks?archived=true&order_by=name&page=2&per_page=3&search=query&simple=true&sort=asc&visibility=public")
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{}
	opt.ListOptions = ListOptions{2, 3}
	opt.Archived = Bool(true)
	opt.OrderBy = String("name")
	opt.Sort = String("asc")
	opt.Search = String("query")
	opt.Simple = Bool(true)
	opt.Visibility = Visibility(PublicVisibility)

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
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/share", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
	})

	opt := &ShareWithGroupOptions{
		GroupID:     Int(1),
		GroupAccess: AccessLevel(AccessLevelValue(50)),
	}

	_, err := client.Projects.ShareProjectWithGroup(1, opt)
	if err != nil {
		t.Errorf("Projects.ShareProjectWithGroup returned error: %v", err)
	}
}

func TestDeleteSharedProjectFromGroup(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/share/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Projects.DeleteSharedProjectFromGroup(1, 2)
	if err != nil {
		t.Errorf("Projects.DeleteSharedProjectFromGroup returned error: %v", err)
	}
}

func TestGetApprovalConfiguration(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/approvals", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
			"approvers": [],
			"approver_groups": [],
			"approvals_before_merge": 3,
			"reset_approvals_on_push": false,
			"disable_overriding_approvers_per_merge_request": false,
			"merge_requests_author_approval": true
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
	}

	if !reflect.DeepEqual(want, approvals) {
		t.Errorf("Projects.GetApprovalConfiguration returned %+v, want %+v", approvals, want)
	}
}

func TestChangeApprovalConfiguration(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/approvals", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"approvals_before_merge":3}`)
		fmt.Fprint(w, `{
			"approvers": [],
			"approver_groups": [],
			"approvals_before_merge": 3,
			"reset_approvals_on_push": false,
			"disable_overriding_approvers_per_merge_request": false,
			"merge_requests_author_approval": true
		}`)
	})

	opt := &ChangeApprovalConfigurationOptions{
		ApprovalsBeforeMerge: Int(3),
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
	}

	if !reflect.DeepEqual(want, approvals) {
		t.Errorf("Projects.ChangeApprovalConfigurationOptions  returned %+v, want %+v", approvals, want)
	}
}

func TestChangeAllowedApprovers(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/approvers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testBody(t, r, `{"approver_ids":[1],"approver_group_ids":[2]}`)
		fmt.Fprint(w, `{
			"approvers": [{"user":{"id":1}}],
			"approver_groups": [{"group":{"id":2}}]
		}`)
	})

	opt := &ChangeAllowedApproversOptions{
		ApproverIDs:      []*int{Int(1)},
		ApproverGroupIDs: []*int{Int(2)},
	}

	approvals, _, err := client.Projects.ChangeAllowedApprovers(1, opt)
	if err != nil {
		t.Errorf("Projects.ChangeApproversConfigurationOptions returned error: %v", err)
	}

	want := &ProjectApprovals{
		Approvers: []*MergeRequestApproverUser{
			&MergeRequestApproverUser{
				User: &BasicUser{
					ID: 1,
				},
			},
		},
		ApproverGroups: []*MergeRequestApproverGroup{
			&MergeRequestApproverGroup{
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
	mux, server, client := setup()
	defer teardown(server)

	namespace := "mynamespace"
	name := "myreponame"
	path := "myrepopath"

	mux.HandleFunc("/api/v4/projects/1/fork", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, fmt.Sprintf(`{"namespace":"%s","name":"%s","path":"%s"}`, namespace, name, path))
		fmt.Fprint(w, `{"id":2}`)
	})

	project, _, err := client.Projects.ForkProject(1, &ForkProjectOptions{
		Namespace: String(namespace),
		Name:      String(name),
		Path:      String(path),
	})
	if err != nil {
		t.Errorf("Projects.ForkProject returned error: %v", err)
	}

	want := &Project{ID: 2}
	if !reflect.DeepEqual(want, project) {
		t.Errorf("Projects.ForProject returned %+v, want %+v", project, want)
	}
}
