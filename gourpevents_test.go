package gitlab

import (
	"encoding/json"
	"testing"
)

func TestGroupMergeEventUnmarshal(t *testing.T) {

	jsonObject := `{
	"object_kind": "merge_request",
	"user": {
		"name": "Administrator",
		"username": "root",
		"avatar_url": "http://www.gravatar.com/avatar/d22738dc40839e3d95fca77ca3eac067?s=80\u0026d=identicon"
	},
	"project": {
		"name": "example-project",
		"description": "",
		"web_url": "http://example.com/exm-namespace/example-project",
		"avatar_url": null,
		"git_ssh_url": "git@example.com:exm-namespace/example-project.git",
		"git_http_url": "http://example.com/exm-namespace/example-project.git",
		"namespace": "exm-namespace",
		"visibility_level": 0,
		"path_with_namespace": "exm-namespace/example-project",
		"default_branch": "master",
		"homepage": "http://example.com/exm-namespace/example-project",
		"url": "git@example.com:exm-namespace/example-project.git",
		"ssh_url": "git@example.com:exm-namespace/example-project.git",
		"http_url": "http://example.com/exm-namespace/example-project.git"
	},
	"object_attributes": {
		"id": 15917,
		"target_branch ": "master ",
		"source_branch ": "source-branch-test ",
		"source_project_id ": 87,
		"author_id ": 15,
		"assignee_id ": 29,
		"title ": "source-branch-test ",
		"created_at ": "2016 - 12 - 01 13: 11: 10 UTC ",
		"updated_at ": "2016 - 12 - 01 13: 21: 20 UTC ",
		"milestone_id ": null,
		"state ": "merged ",
		"merge_status ": "can_be_merged ",
		"target_project_id ": 87,
		"iid ": 1402,
		"description ": "word doc support for e - ticket ",
		"position ": 0,
		"locked_at ": null,
		"updated_by_id ": null,
		"merge_error ": null,
		"merge_params": {
			"force_remove_source_branch": "0"
		},
		"merge_when_build_succeeds": false,
		"merge_user_id": null,
		"merge_commit_sha": "ac3ca1559bc39abf963586372eff7f8fdded646e",
		"deleted_at": null,
		"approvals_before_merge": null,
		"rebase_commit_sha": null,
		"in_progress_merge_commit_sha": null,
		"lock_version": 0,
		"time_estimate": 0,
		"source": {
			"name": "example-project",
			"description": "",
			"web_url": "http://example.com/exm-namespace/example-project",
			"avatar_url": null,
			"git_ssh_url": "git@example.com:exm-namespace/example-project.git",
			"git_http_url": "http://example.com/exm-namespace/example-project.git",
			"namespace": "exm-namespace",
			"visibility_level": 0,
			"path_with_namespace": "exm-namespace/example-project",
			"default_branch": "master",
			"homepage": "http://example.com/exm-namespace/example-project",
			"url": "git@example.com:exm-namespace/example-project.git",
			"ssh_url": "git@example.com:exm-namespace/example-project.git",
			"http_url": "http://example.com/exm-namespace/example-project.git"
		},
		"target": {
			"name": "example-project",
			"description": "",
			"web_url": "http://example.com/exm-namespace/example-project",
			"avatar_url": null,
			"git_ssh_url": "git@example.com:exm-namespace/example-project.git",
			"git_http_url": "http://example.com/exm-namespace/example-project.git",
			"namespace": "exm-namespace",
			"visibility_level": 0,
			"path_with_namespace": "exm-namespace/example-project",
			"default_branch": "master",
			"homepage": "http://example.com/exm-namespace/example-project",
			"url": "git@example.com:exm-namespace/example-project.git",
			"ssh_url": "git@example.com:exm-namespace/example-project.git",
			"http_url": "http://example.com/exm-namespace/example-project.git"
		},
		"last_commit": {
			"id": "61b6a0d35dbaf915760233b637622e383d3cc9ec",
			"message": "commit message",
			"timestamp": "2016-12-01T15:07:53+02:00",
			"url": "http://example.com/exm-namespace/example-project/commit/61b6a0d35dbaf915760233b637622e383d3cc9ec",
			"author": {
				"name": "Test User",
				"email": "test.user@mail.com"
			}
		},
		"work_in_progress": false,
		"url": "http://example.com/exm-namespace/example-project/merge_requests/1402",
		"action": "merge"
	},
	"repository": {
		"name": "example-project",
		"url": "git@example.com:exm-namespace/example-project.git",
		"description": "",
		"homepage": "http://example.com/exm-namespace/example-project"
	},
	"assignee": {
		"name": "Administrator",
		"username": "root",
		"avatar_url": "http://www.gravatar.com/avatar/d22738dc40839e3d95fca77ca3eac067?s=80\u0026d=identicon"
	}
}`

	var event *GroupMergeEvent
	err := json.Unmarshal([]byte(jsonObject), &event)

	if err != nil {
		t.Errorf("Group Merge Event can not unmarshaled: %v\n ", err.Error())
	}

	if event == nil {
		t.Errorf("Group Merge Event is null")
	}

	if event.ObjectKind != "merge_request" {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, "merge_request")
	}

	if event.User.Username != "root" {
		t.Errorf("User.Username is %v, want %v", event.User.Username, "root")
	}

	if event.Project.Name != "example-project" {
		t.Errorf("Project.Name is %v, want %v", event.Project.Name, "example-project")
	}

	if event.ObjectAttributes.ID != 15917 {
		t.Errorf("ObjectAttributes.ID is %v, want %v", event.ObjectAttributes.ID, 15917)
	}

	if event.ObjectAttributes.Source.Name != "example-project" {
		t.Errorf("ObjectAttributes.Source.Name is %v, want %v", event.ObjectAttributes.Source.Name, "example-project")
	}

	if event.ObjectAttributes.LastCommit.Author.Email != "test.user@mail.com" {
		t.Errorf("ObjectAttributes.LastCommit.Author.Email is %v, want %v", event.ObjectAttributes.LastCommit.Author.Email, "test.user@mail.com")
	}

	if event.Repository.Name != "example-project" {
		t.Errorf("Repository.Name is %v, want %v", event.Repository.Name, "example-project")
	}

	if event.Assignee.Username != "root" {
		t.Errorf("Assignee.Username is %v, want %v", event.Assignee, "root")
	}
}
