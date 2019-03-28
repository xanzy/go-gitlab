//
// Copyright 2017, Sander van Harmelen
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
	"time"
)

// PushEvent represents a push event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#push-events
type PushEvent struct {
	ObjectKind  string `json:"object_kind" yaml:"object_kind"`
	Before      string `json:"before" yaml:"before"`
	After       string `json:"after" yaml:"after"`
	Ref         string `json:"ref" yaml:"ref"`
	CheckoutSHA string `json:"checkout_sha" yaml:"checkout_sha"`
	UserID      int    `json:"user_id" yaml:"user_id"`
	UserName    string `json:"user_name" yaml:"user_name"`
	UserEmail   string `json:"user_email" yaml:"user_email"`
	UserAvatar  string `json:"user_avatar" yaml:"user_avatar"`
	ProjectID   int    `json:"project_id" yaml:"project_id"`
	Project     struct {
		Name              string          `json:"name" yaml:"name"`
		Description       string          `json:"description" yaml:"description"`
		AvatarURL         string          `json:"avatar_url" yaml:"avatar_url"`
		GitSSHURL         string          `json:"git_ssh_url" yaml:"git_ssh_url"`
		GitHTTPURL        string          `json:"git_http_url" yaml:"git_http_url"`
		Namespace         string          `json:"namespace" yaml:"namespace"`
		PathWithNamespace string          `json:"path_with_namespace" yaml:"path_with_namespace"`
		DefaultBranch     string          `json:"default_branch" yaml:"default_branch"`
		Homepage          string          `json:"homepage" yaml:"homepage"`
		URL               string          `json:"url" yaml:"url"`
		SSHURL            string          `json:"ssh_url" yaml:"ssh_url"`
		HTTPURL           string          `json:"http_url" yaml:"http_url"`
		WebURL            string          `json:"web_url" yaml:"web_url"`
		Visibility        VisibilityValue `json:"visibility" yaml:"visibility"`
	} `json:"project" yaml:"project"`
	Repository *Repository `json:"repository" yaml:"repository"`
	Commits    []*struct {
		ID        string     `json:"id" yaml:"id"`
		Message   string     `json:"message" yaml:"message"`
		Timestamp *time.Time `json:"timestamp" yaml:"timestamp"`
		URL       string     `json:"url" yaml:"url"`
		Author    struct {
			Name  string `json:"name" yaml:"name"`
			Email string `json:"email" yaml:"email"`
		} `json:"author" yaml:"author"`
		Added    []string `json:"added" yaml:"added"`
		Modified []string `json:"modified" yaml:"modified"`
		Removed  []string `json:"removed" yaml:"removed"`
	} `json:"commits" yaml:"commits"`
	TotalCommitsCount int `json:"total_commits_count" yaml:"total_commits_count"`
}

// TagEvent represents a tag event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#tag-events
type TagEvent struct {
	ObjectKind  string `json:"object_kind" yaml:"object_kind"`
	Before      string `json:"before" yaml:"before"`
	After       string `json:"after" yaml:"after"`
	Ref         string `json:"ref" yaml:"ref"`
	CheckoutSHA string `json:"checkout_sha" yaml:"checkout_sha"`
	UserID      int    `json:"user_id" yaml:"user_id"`
	UserName    string `json:"user_name" yaml:"user_name"`
	UserAvatar  string `json:"user_avatar" yaml:"user_avatar"`
	ProjectID   int    `json:"project_id" yaml:"project_id"`
	Project     struct {
		Name              string          `json:"name" yaml:"name"`
		Description       string          `json:"description" yaml:"description"`
		AvatarURL         string          `json:"avatar_url" yaml:"avatar_url"`
		GitSSHURL         string          `json:"git_ssh_url" yaml:"git_ssh_url"`
		GitHTTPURL        string          `json:"git_http_url" yaml:"git_http_url"`
		Namespace         string          `json:"namespace" yaml:"namespace"`
		PathWithNamespace string          `json:"path_with_namespace" yaml:"path_with_namespace"`
		DefaultBranch     string          `json:"default_branch" yaml:"default_branch"`
		Homepage          string          `json:"homepage" yaml:"homepage"`
		URL               string          `json:"url" yaml:"url"`
		SSHURL            string          `json:"ssh_url" yaml:"ssh_url"`
		HTTPURL           string          `json:"http_url" yaml:"http_url"`
		WebURL            string          `json:"web_url" yaml:"web_url"`
		Visibility        VisibilityValue `json:"visibility" yaml:"visibility"`
	} `json:"project" yaml:"project"`
	Repository *Repository `json:"repository" yaml:"repository"`
	Commits    []*struct {
		ID        string     `json:"id" yaml:"id"`
		Message   string     `json:"message" yaml:"message"`
		Timestamp *time.Time `json:"timestamp" yaml:"timestamp"`
		URL       string     `json:"url" yaml:"url"`
		Author    struct {
			Name  string `json:"name" yaml:"name"`
			Email string `json:"email" yaml:"email"`
		} `json:"author" yaml:"author"`
		Added    []string `json:"added" yaml:"added"`
		Modified []string `json:"modified" yaml:"modified"`
		Removed  []string `json:"removed" yaml:"removed"`
	} `json:"commits" yaml:"commits"`
	TotalCommitsCount int `json:"total_commits_count" yaml:"total_commits_count"`
}

// IssueEvent represents a issue event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#issues-events
type IssueEvent struct {
	ObjectKind string `json:"object_kind" yaml:"object_kind"`
	User       *User  `json:"user" yaml:"user"`
	Project    struct {
		Name              string          `json:"name" yaml:"name"`
		Description       string          `json:"description" yaml:"description"`
		AvatarURL         string          `json:"avatar_url" yaml:"avatar_url"`
		GitSSHURL         string          `json:"git_ssh_url" yaml:"git_ssh_url"`
		GitHTTPURL        string          `json:"git_http_url" yaml:"git_http_url"`
		Namespace         string          `json:"namespace" yaml:"namespace"`
		PathWithNamespace string          `json:"path_with_namespace" yaml:"path_with_namespace"`
		DefaultBranch     string          `json:"default_branch" yaml:"default_branch"`
		Homepage          string          `json:"homepage" yaml:"homepage"`
		URL               string          `json:"url" yaml:"url"`
		SSHURL            string          `json:"ssh_url" yaml:"ssh_url"`
		HTTPURL           string          `json:"http_url" yaml:"http_url"`
		WebURL            string          `json:"web_url" yaml:"web_url"`
		Visibility        VisibilityValue `json:"visibility" yaml:"visibility"`
	} `json:"project" yaml:"project"`
	Repository       *Repository `json:"repository" yaml:"repository"`
	ObjectAttributes struct {
		ID          int    `json:"id" yaml:"id"`
		Title       string `json:"title" yaml:"title"`
		AssigneeID  int    `json:"assignee_id" yaml:"assignee_id"`
		AuthorID    int    `json:"author_id" yaml:"author_id"`
		ProjectID   int    `json:"project_id" yaml:"project_id"`
		CreatedAt   string `json:"created_at" yaml:"created_at"` // Should be *time.Time (see Gitlab issue #21468)
		UpdatedAt   string `json:"updated_at" yaml:"updated_at"` // Should be *time.Time (see Gitlab issue #21468)
		Position    int    `json:"position" yaml:"position"`
		BranchName  string `json:"branch_name" yaml:"branch_name"`
		Description string `json:"description" yaml:"description"`
		MilestoneID int    `json:"milestone_id" yaml:"milestone_id"`
		State       string `json:"state" yaml:"state"`
		IID         int    `json:"iid" yaml:"iid"`
		URL         string `json:"url" yaml:"url"`
		Action      string `json:"action" yaml:"action"`
	} `json:"object_attributes" yaml:"object_attributes"`
	Assignee struct {
		Name      string `json:"name" yaml:"name"`
		Username  string `json:"username" yaml:"username"`
		AvatarURL string `json:"avatar_url" yaml:"avatar_url"`
	} `json:"assignee" yaml:"assignee"`
	Labels []Label `json:"labels" yaml:"labels"`
	Changes struct {
		Labels struct {
			Previous []Label `json:"previous" yaml:"previous"`
			Current  []Label `json:"current" yaml:"current"`
		} `json:"labels" yaml:"labels"`
		UpdatedByID []int `json:"updated_by_id" yaml:"updated_by_id"`
	} `json:"changes" yaml:"changes"`
}

// JobEvent represents a job event.
//
// GitLab API docs:
// TODO: link to docs instead of src once they are published.
// https://gitlab.com/gitlab-org/gitlab-ce/blob/master/lib/gitlab/data_builder/build.rb
type JobEvent struct {
	ObjectKind        string  `json:"object_kind" yaml:"object_kind"`
	Ref               string  `json:"ref" yaml:"ref"`
	Tag               bool    `json:"tag" yaml:"tag"`
	BeforeSHA         string  `json:"before_sha" yaml:"before_sha"`
	SHA               string  `json:"sha" yaml:"sha"`
	BuildID           int     `json:"build_id" yaml:"build_id"`
	BuildName         string  `json:"build_name" yaml:"build_name"`
	BuildStage        string  `json:"build_stage" yaml:"build_stage"`
	BuildStatus       string  `json:"build_status" yaml:"build_status"`
	BuildStartedAt    string  `json:"build_started_at" yaml:"build_started_at"`
	BuildFinishedAt   string  `json:"build_finished_at" yaml:"build_finished_at"`
	BuildDuration     float64 `json:"build_duration" yaml:"build_duration"`
	BuildAllowFailure bool    `json:"build_allow_failure" yaml:"build_allow_failure"`
	ProjectID         int     `json:"project_id" yaml:"project_id"`
	ProjectName       string  `json:"project_name" yaml:"project_name"`
	User              struct {
		ID    int    `json:"id" yaml:"id"`
		Name  string `json:"name" yaml:"name"`
		Email string `json:"email" yaml:"email"`
	} `json:"user" yaml:"user"`
	Commit struct {
		ID          int    `json:"id" yaml:"id"`
		SHA         string `json:"sha" yaml:"sha"`
		Message     string `json:"message" yaml:"message"`
		AuthorName  string `json:"author_name" yaml:"author_name"`
		AuthorEmail string `json:"author_email" yaml:"author_email"`
		AuthorURL   string `json:"author_url" yaml:"author_url"`
		Status      string `json:"status" yaml:"status"`
		Duration    int    `json:"duration" yaml:"duration"`
		StartedAt   string `json:"started_at" yaml:"started_at"`
		FinishedAt  string `json:"finished_at" yaml:"finished_at"`
	} `json:"commit" yaml:"commit"`
	Repository *Repository `json:"repository" yaml:"repository"`
}

// CommitCommentEvent represents a comment on a commit event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#comment-on-commit
type CommitCommentEvent struct {
	ObjectKind string `json:"object_kind" yaml:"object_kind"`
	User       *User  `json:"user" yaml:"user"`
	ProjectID  int    `json:"project_id" yaml:"project_id"`
	Project    struct {
		Name              string          `json:"name" yaml:"name"`
		Description       string          `json:"description" yaml:"description"`
		AvatarURL         string          `json:"avatar_url" yaml:"avatar_url"`
		GitSSHURL         string          `json:"git_ssh_url" yaml:"git_ssh_url"`
		GitHTTPURL        string          `json:"git_http_url" yaml:"git_http_url"`
		Namespace         string          `json:"namespace" yaml:"namespace"`
		PathWithNamespace string          `json:"path_with_namespace" yaml:"path_with_namespace"`
		DefaultBranch     string          `json:"default_branch" yaml:"default_branch"`
		Homepage          string          `json:"homepage" yaml:"homepage"`
		URL               string          `json:"url" yaml:"url"`
		SSHURL            string          `json:"ssh_url" yaml:"ssh_url"`
		HTTPURL           string          `json:"http_url" yaml:"http_url"`
		WebURL            string          `json:"web_url" yaml:"web_url"`
		Visibility        VisibilityValue `json:"visibility" yaml:"visibility"`
	} `json:"project" yaml:"project"`
	Repository       *Repository `json:"repository" yaml:"repository"`
	ObjectAttributes struct {
		ID           int    `json:"id" yaml:"id"`
		Note         string `json:"note" yaml:"note"`
		NoteableType string `json:"noteable_type" yaml:"noteable_type"`
		AuthorID     int    `json:"author_id" yaml:"author_id"`
		CreatedAt    string `json:"created_at" yaml:"created_at"`
		UpdatedAt    string `json:"updated_at" yaml:"updated_at"`
		ProjectID    int    `json:"project_id" yaml:"project_id"`
		Attachment   string `json:"attachment" yaml:"attachment"`
		LineCode     string `json:"line_code" yaml:"line_code"`
		CommitID     string `json:"commit_id" yaml:"commit_id"`
		NoteableID   int    `json:"noteable_id" yaml:"noteable_id"`
		System       bool   `json:"system" yaml:"system"`
		StDiff       struct {
			Diff        string `json:"diff" yaml:"diff"`
			NewPath     string `json:"new_path" yaml:"new_path"`
			OldPath     string `json:"old_path" yaml:"old_path"`
			AMode       string `json:"a_mode" yaml:"a_mode"`
			BMode       string `json:"b_mode" yaml:"b_mode"`
			NewFile     bool   `json:"new_file" yaml:"new_file"`
			RenamedFile bool   `json:"renamed_file" yaml:"renamed_file"`
			DeletedFile bool   `json:"deleted_file" yaml:"deleted_file"`
		} `json:"st_diff" yaml:"st_diff"`
	} `json:"object_attributes" yaml:"object_attributes"`
	Commit *struct {
		ID        string     `json:"id" yaml:"id"`
		Message   string     `json:"message" yaml:"message"`
		Timestamp *time.Time `json:"timestamp" yaml:"timestamp"`
		URL       string     `json:"url" yaml:"url"`
		Author    struct {
			Name  string `json:"name" yaml:"name"`
			Email string `json:"email" yaml:"email"`
		} `json:"author" yaml:"author"`
	} `json:"commit" yaml:"commit"`
}

// MergeCommentEvent represents a comment on a merge event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#comment-on-merge-request
type MergeCommentEvent struct {
	ObjectKind string `json:"object_kind" yaml:"object_kind"`
	User       *User  `json:"user" yaml:"user"`
	ProjectID  int    `json:"project_id" yaml:"project_id"`
	Project    struct {
		Name              string          `json:"name" yaml:"name"`
		Description       string          `json:"description" yaml:"description"`
		AvatarURL         string          `json:"avatar_url" yaml:"avatar_url"`
		GitSSHURL         string          `json:"git_ssh_url" yaml:"git_ssh_url"`
		GitHTTPURL        string          `json:"git_http_url" yaml:"git_http_url"`
		Namespace         string          `json:"namespace" yaml:"namespace"`
		PathWithNamespace string          `json:"path_with_namespace" yaml:"path_with_namespace"`
		DefaultBranch     string          `json:"default_branch" yaml:"default_branch"`
		Homepage          string          `json:"homepage" yaml:"homepage"`
		URL               string          `json:"url" yaml:"url"`
		SSHURL            string          `json:"ssh_url" yaml:"ssh_url"`
		HTTPURL           string          `json:"http_url" yaml:"http_url"`
		WebURL            string          `json:"web_url" yaml:"web_url"`
		Visibility        VisibilityValue `json:"visibility" yaml:"visibility"`
	} `json:"project" yaml:"project"`
	ObjectAttributes struct {
		ID           int    `json:"id" yaml:"id"`
		Note         string `json:"note" yaml:"note"`
		NoteableType string `json:"noteable_type" yaml:"noteable_type"`
		AuthorID     int    `json:"author_id" yaml:"author_id"`
		CreatedAt    string `json:"created_at" yaml:"created_at"`
		UpdatedAt    string `json:"updated_at" yaml:"updated_at"`
		ProjectID    int    `json:"project_id" yaml:"project_id"`
		Attachment   string `json:"attachment" yaml:"attachment"`
		LineCode     string `json:"line_code" yaml:"line_code"`
		CommitID     string `json:"commit_id" yaml:"commit_id"`
		NoteableID   int    `json:"noteable_id" yaml:"noteable_id"`
		System       bool   `json:"system" yaml:"system"`
		StDiff       *Diff  `json:"st_diff" yaml:"st_diff"`
		URL          string `json:"url" yaml:"url"`
	} `json:"object_attributes" yaml:"object_attributes"`
	Repository   *Repository `json:"repository" yaml:"repository"`
	MergeRequest struct {
		ID              int    `json:"id" yaml:"id"`
		TargetBranch    string `json:"target_branch" yaml:"target_branch"`
		SourceBranch    string `json:"source_branch" yaml:"source_branch"`
		SourceProjectID int    `json:"source_project_id" yaml:"source_project_id"`
		AuthorID        int    `json:"author_id" yaml:"author_id"`
		AssigneeID      int    `json:"assignee_id" yaml:"assignee_id"`
		Title           string `json:"title" yaml:"title"`
		CreatedAt       string `json:"created_at" yaml:"created_at"`
		UpdatedAt       string `json:"updated_at" yaml:"updated_at"`
		MilestoneID     int    `json:"milestone_id" yaml:"milestone_id"`
		State           string `json:"state" yaml:"state"`
		MergeStatus     string `json:"merge_status" yaml:"merge_status"`
		TargetProjectID int    `json:"target_project_id" yaml:"target_project_id"`
		IID             int    `json:"iid" yaml:"iid"`
		Description     string `json:"description" yaml:"description"`
		Position        int    `json:"position" yaml:"position"`
		LockedAt        string `json:"locked_at" yaml:"locked_at"`
		UpdatedByID     int    `json:"updated_by_id" yaml:"updated_by_id"`
		MergeError      string `json:"merge_error" yaml:"merge_error"`
		MergeParams     struct {
			ForceRemoveSourceBranch string `json:"force_remove_source_branch" yaml:"force_remove_source_branch"`
		} `json:"merge_params" yaml:"merge_params"`
		MergeWhenPipelineSucceeds bool        `json:"merge_when_pipeline_succeeds" yaml:"merge_when_pipeline_succeeds"`
		MergeUserID               int         `json:"merge_user_id" yaml:"merge_user_id"`
		MergeCommitSHA            string      `json:"merge_commit_sha" yaml:"merge_commit_sha"`
		DeletedAt                 string      `json:"deleted_at" yaml:"deleted_at"`
		InProgressMergeCommitSHA  string      `json:"in_progress_merge_commit_sha" yaml:"in_progress_merge_commit_sha"`
		LockVersion               int         `json:"lock_version" yaml:"lock_version"`
		ApprovalsBeforeMerge      string      `json:"approvals_before_merge" yaml:"approvals_before_merge"`
		RebaseCommitSHA           string      `json:"rebase_commit_sha" yaml:"rebase_commit_sha"`
		TimeEstimate              int         `json:"time_estimate" yaml:"time_estimate"`
		Squash                    bool        `json:"squash" yaml:"squash"`
		LastEditedAt              string      `json:"last_edited_at" yaml:"last_edited_at"`
		LastEditedByID            int         `json:"last_edited_by_id" yaml:"last_edited_by_id"`
		Source                    *Repository `json:"source" yaml:"source"`
		Target                    *Repository `json:"target" yaml:"target"`
		LastCommit                struct {
			ID        string     `json:"id" yaml:"id"`
			Message   string     `json:"message" yaml:"message"`
			Timestamp *time.Time `json:"timestamp" yaml:"timestamp"`
			URL       string     `json:"url" yaml:"url"`
			Author    struct {
				Name  string `json:"name" yaml:"name"`
				Email string `json:"email" yaml:"email"`
			} `json:"author" yaml:"author"`
		} `json:"last_commit" yaml:"last_commit"`
		WorkInProgress bool `json:"work_in_progress" yaml:"work_in_progress"`
		TotalTimeSpent int  `json:"total_time_spent" yaml:"total_time_spent"`
	} `json:"merge_request" yaml:"merge_request"`
}

// IssueCommentEvent represents a comment on an issue event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#comment-on-issue
type IssueCommentEvent struct {
	ObjectKind string `json:"object_kind" yaml:"object_kind"`
	User       *User  `json:"user" yaml:"user"`
	ProjectID  int    `json:"project_id" yaml:"project_id"`
	Project    struct {
		Name              string          `json:"name" yaml:"name"`
		Description       string          `json:"description" yaml:"description"`
		AvatarURL         string          `json:"avatar_url" yaml:"avatar_url"`
		GitSSHURL         string          `json:"git_ssh_url" yaml:"git_ssh_url"`
		GitHTTPURL        string          `json:"git_http_url" yaml:"git_http_url"`
		Namespace         string          `json:"namespace" yaml:"namespace"`
		PathWithNamespace string          `json:"path_with_namespace" yaml:"path_with_namespace"`
		DefaultBranch     string          `json:"default_branch" yaml:"default_branch"`
		Homepage          string          `json:"homepage" yaml:"homepage"`
		URL               string          `json:"url" yaml:"url"`
		SSHURL            string          `json:"ssh_url" yaml:"ssh_url"`
		HTTPURL           string          `json:"http_url" yaml:"http_url"`
		WebURL            string          `json:"web_url" yaml:"web_url"`
		Visibility        VisibilityValue `json:"visibility" yaml:"visibility"`
	} `json:"project" yaml:"project"`
	Repository       *Repository `json:"repository" yaml:"repository"`
	ObjectAttributes struct {
		ID           int     `json:"id" yaml:"id"`
		Note         string  `json:"note" yaml:"note"`
		NoteableType string  `json:"noteable_type" yaml:"noteable_type"`
		AuthorID     int     `json:"author_id" yaml:"author_id"`
		CreatedAt    string  `json:"created_at" yaml:"created_at"`
		UpdatedAt    string  `json:"updated_at" yaml:"updated_at"`
		ProjectID    int     `json:"project_id" yaml:"project_id"`
		Attachment   string  `json:"attachment" yaml:"attachment"`
		LineCode     string  `json:"line_code" yaml:"line_code"`
		CommitID     string  `json:"commit_id" yaml:"commit_id"`
		NoteableID   int     `json:"noteable_id" yaml:"noteable_id"`
		System       bool    `json:"system" yaml:"system"`
		StDiff       []*Diff `json:"st_diff" yaml:"st_diff"`
		URL          string  `json:"url" yaml:"url"`
	} `json:"object_attributes" yaml:"object_attributes"`
	Issue *Issue `json:"issue" yaml:"issue"`
}

// SnippetCommentEvent represents a comment on a snippet event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#comment-on-code-snippet
type SnippetCommentEvent struct {
	ObjectKind string `json:"object_kind" yaml:"object_kind"`
	User       *User  `json:"user" yaml:"user"`
	ProjectID  int    `json:"project_id" yaml:"project_id"`
	Project    struct {
		Name              string          `json:"name" yaml:"name"`
		Description       string          `json:"description" yaml:"description"`
		AvatarURL         string          `json:"avatar_url" yaml:"avatar_url"`
		GitSSHURL         string          `json:"git_ssh_url" yaml:"git_ssh_url"`
		GitHTTPURL        string          `json:"git_http_url" yaml:"git_http_url"`
		Namespace         string          `json:"namespace" yaml:"namespace"`
		PathWithNamespace string          `json:"path_with_namespace" yaml:"path_with_namespace"`
		DefaultBranch     string          `json:"default_branch" yaml:"default_branch"`
		Homepage          string          `json:"homepage" yaml:"homepage"`
		URL               string          `json:"url" yaml:"url"`
		SSHURL            string          `json:"ssh_url" yaml:"ssh_url"`
		HTTPURL           string          `json:"http_url" yaml:"http_url"`
		WebURL            string          `json:"web_url" yaml:"web_url"`
		Visibility        VisibilityValue `json:"visibility" yaml:"visibility"`
	} `json:"project" yaml:"project"`
	Repository       *Repository `json:"repository" yaml:"repository"`
	ObjectAttributes struct {
		ID           int    `json:"id" yaml:"id"`
		Note         string `json:"note" yaml:"note"`
		NoteableType string `json:"noteable_type" yaml:"noteable_type"`
		AuthorID     int    `json:"author_id" yaml:"author_id"`
		CreatedAt    string `json:"created_at" yaml:"created_at"`
		UpdatedAt    string `json:"updated_at" yaml:"updated_at"`
		ProjectID    int    `json:"project_id" yaml:"project_id"`
		Attachment   string `json:"attachment" yaml:"attachment"`
		LineCode     string `json:"line_code" yaml:"line_code"`
		CommitID     string `json:"commit_id" yaml:"commit_id"`
		NoteableID   int    `json:"noteable_id" yaml:"noteable_id"`
		System       bool   `json:"system" yaml:"system"`
		StDiff       *Diff  `json:"st_diff" yaml:"st_diff"`
		URL          string `json:"url" yaml:"url"`
	} `json:"object_attributes" yaml:"object_attributes"`
	Snippet *Snippet `json:"snippet" yaml:"snippet"`
}

// MergeEvent represents a merge event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#merge-request-events
type MergeEvent struct {
	ObjectKind string `json:"object_kind" yaml:"object_kind"`
	User       *User  `json:"user" yaml:"user"`
	Project    struct {
		ID                int             `json:"id" yaml:"id"`
		Name              string          `json:"name" yaml:"name"`
		Description       string          `json:"description" yaml:"description"`
		AvatarURL         string          `json:"avatar_url" yaml:"avatar_url"`
		GitSSHURL         string          `json:"git_ssh_url" yaml:"git_ssh_url"`
		GitHTTPURL        string          `json:"git_http_url" yaml:"git_http_url"`
		Namespace         string          `json:"namespace" yaml:"namespace"`
		PathWithNamespace string          `json:"path_with_namespace" yaml:"path_with_namespace"`
		DefaultBranch     string          `json:"default_branch" yaml:"default_branch"`
		Homepage          string          `json:"homepage" yaml:"homepage"`
		URL               string          `json:"url" yaml:"url"`
		SSHURL            string          `json:"ssh_url" yaml:"ssh_url"`
		HTTPURL           string          `json:"http_url" yaml:"http_url"`
		WebURL            string          `json:"web_url" yaml:"web_url"`
		Visibility        VisibilityValue `json:"visibility" yaml:"visibility"`
	} `json:"project" yaml:"project"`
	ObjectAttributes struct {
		ID              int       `json:"id" yaml:"id"`
		TargetBranch    string    `json:"target_branch" yaml:"target_branch"`
		SourceBranch    string    `json:"source_branch" yaml:"source_branch"`
		SourceProjectID int       `json:"source_project_id" yaml:"source_project_id"`
		AuthorID        int       `json:"author_id" yaml:"author_id"`
		AssigneeID      int       `json:"assignee_id" yaml:"assignee_id"`
		Title           string    `json:"title" yaml:"title"`
		CreatedAt       string    `json:"created_at" yaml:"created_at"` // Should be *time.Time (see Gitlab issue #21468)
		UpdatedAt       string    `json:"updated_at" yaml:"updated_at"` // Should be *time.Time (see Gitlab issue #21468)
		StCommits       []*Commit `json:"st_commits" yaml:"st_commits"`
		StDiffs         []*Diff   `json:"st_diffs" yaml:"st_diffs"`
		MilestoneID     int       `json:"milestone_id" yaml:"milestone_id"`
		State           string    `json:"state" yaml:"state"`
		MergeStatus     string    `json:"merge_status" yaml:"merge_status"`
		TargetProjectID int       `json:"target_project_id" yaml:"target_project_id"`
		IID             int       `json:"iid" yaml:"iid"`
		Description     string    `json:"description" yaml:"description"`
		Position        int       `json:"position" yaml:"position"`
		LockedAt        string    `json:"locked_at" yaml:"locked_at"`
		UpdatedByID     int       `json:"updated_by_id" yaml:"updated_by_id"`
		MergeError      string    `json:"merge_error" yaml:"merge_error"`
		MergeParams     struct {
			ForceRemoveSourceBranch string `json:"force_remove_source_branch" yaml:"force_remove_source_branch"`
		} `json:"merge_params" yaml:"merge_params"`
		MergeWhenBuildSucceeds   bool        `json:"merge_when_build_succeeds" yaml:"merge_when_build_succeeds"`
		MergeUserID              int         `json:"merge_user_id" yaml:"merge_user_id"`
		MergeCommitSHA           string      `json:"merge_commit_sha" yaml:"merge_commit_sha"`
		DeletedAt                string      `json:"deleted_at" yaml:"deleted_at"`
		ApprovalsBeforeMerge     string      `json:"approvals_before_merge" yaml:"approvals_before_merge"`
		RebaseCommitSHA          string      `json:"rebase_commit_sha" yaml:"rebase_commit_sha"`
		InProgressMergeCommitSHA string      `json:"in_progress_merge_commit_sha" yaml:"in_progress_merge_commit_sha"`
		LockVersion              int         `json:"lock_version" yaml:"lock_version"`
		TimeEstimate             int         `json:"time_estimate" yaml:"time_estimate"`
		Source                   *Repository `json:"source" yaml:"source"`
		Target                   *Repository `json:"target" yaml:"target"`
		LastCommit               struct {
			ID        string     `json:"id" yaml:"id"`
			Message   string     `json:"message" yaml:"message"`
			Timestamp *time.Time `json:"timestamp" yaml:"timestamp"`
			URL       string     `json:"url" yaml:"url"`
			Author    struct {
				Name  string `json:"name" yaml:"name"`
				Email string `json:"email" yaml:"email"`
			} `json:"author" yaml:"author"`
		} `json:"last_commit" yaml:"last_commit"`
		WorkInProgress bool   `json:"work_in_progress" yaml:"work_in_progress"`
		URL            string `json:"url" yaml:"url"`
		Action         string `json:"action" yaml:"action"`
		OldRev         string `json:"oldrev" yaml:"oldrev"`
		Assignee       struct {
			Name      string `json:"name" yaml:"name"`
			Username  string `json:"username" yaml:"username"`
			AvatarURL string `json:"avatar_url" yaml:"avatar_url"`
		} `json:"assignee" yaml:"assignee"`
	} `json:"object_attributes" yaml:"object_attributes"`
	Repository *Repository `json:"repository" yaml:"repository"`
	Assignee   struct {
		Name      string `json:"name" yaml:"name"`
		Username  string `json:"username" yaml:"username"`
		AvatarURL string `json:"avatar_url" yaml:"avatar_url"`
	} `json:"assignee" yaml:"assignee"`
	Labels []Label `json:"labels" yaml:"labels"`
	Changes struct {
		AssigneeID struct {
			Previous int `json:"previous" yaml:"previous"`
			Current  int `json:"current" yaml:"current"`
		} `json:"assignee_id" yaml:"assignee_id"`
		Description struct {
			Previous string `json:"previous" yaml:"previous"`
			Current  string `json:"current" yaml:"current"`
		} `json:"description" yaml:"description"`
		Labels struct {
			Previous []Label `json:"previous" yaml:"previous"`
			Current  []Label `json:"current" yaml:"current"`
		} `json:"labels" yaml:"labels"`
		UpdatedByID []int `json:"updated_by_id" yaml:"updated_by_id"`
	} `json:"changes" yaml:"changes"`
}

// WikiPageEvent represents a wiki page event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#wiki-page-events
type WikiPageEvent struct {
	ObjectKind string `json:"object_kind" yaml:"object_kind"`
	User       *User  `json:"user" yaml:"user"`
	Project    struct {
		Name              string          `json:"name" yaml:"name"`
		Description       string          `json:"description" yaml:"description"`
		AvatarURL         string          `json:"avatar_url" yaml:"avatar_url"`
		GitSSHURL         string          `json:"git_ssh_url" yaml:"git_ssh_url"`
		GitHTTPURL        string          `json:"git_http_url" yaml:"git_http_url"`
		Namespace         string          `json:"namespace" yaml:"namespace"`
		PathWithNamespace string          `json:"path_with_namespace" yaml:"path_with_namespace"`
		DefaultBranch     string          `json:"default_branch" yaml:"default_branch"`
		Homepage          string          `json:"homepage" yaml:"homepage"`
		URL               string          `json:"url" yaml:"url"`
		SSHURL            string          `json:"ssh_url" yaml:"ssh_url"`
		HTTPURL           string          `json:"http_url" yaml:"http_url"`
		WebURL            string          `json:"web_url" yaml:"web_url"`
		Visibility        VisibilityValue `json:"visibility" yaml:"visibility"`
	} `json:"project" yaml:"project"`
	Wiki struct {
		WebURL            string `json:"web_url" yaml:"web_url"`
		GitSSHURL         string `json:"git_ssh_url" yaml:"git_ssh_url"`
		GitHTTPURL        string `json:"git_http_url" yaml:"git_http_url"`
		PathWithNamespace string `json:"path_with_namespace" yaml:"path_with_namespace"`
		DefaultBranch     string `json:"default_branch" yaml:"default_branch"`
	} `json:"wiki" yaml:"wiki"`
	ObjectAttributes struct {
		Title   string `json:"title" yaml:"title"`
		Content string `json:"content" yaml:"content"`
		Format  string `json:"format" yaml:"format"`
		Message string `json:"message" yaml:"message"`
		Slug    string `json:"slug" yaml:"slug"`
		URL     string `json:"url" yaml:"url"`
		Action  string `json:"action" yaml:"action"`
	} `json:"object_attributes" yaml:"object_attributes"`
}

// PipelineEvent represents a pipeline event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#pipeline-events
type PipelineEvent struct {
	ObjectKind       string `json:"object_kind" yaml:"object_kind"`
	ObjectAttributes struct {
		ID         int      `json:"id" yaml:"id"`
		Ref        string   `json:"ref" yaml:"ref"`
		Tag        bool     `json:"tag" yaml:"tag"`
		SHA        string   `json:"sha" yaml:"sha"`
		BeforeSHA  string   `json:"before_sha" yaml:"before_sha"`
		Status     string   `json:"status" yaml:"status"`
		Stages     []string `json:"stages" yaml:"stages"`
		CreatedAt  string   `json:"created_at" yaml:"created_at"`
		FinishedAt string   `json:"finished_at" yaml:"finished_at"`
		Duration   int      `json:"duration" yaml:"duration"`
	} `json:"object_attributes" yaml:"object_attributes"`
	User struct {
		Name      string `json:"name" yaml:"name"`
		Username  string `json:"username" yaml:"username"`
		AvatarURL string `json:"avatar_url" yaml:"avatar_url"`
	} `json:"user" yaml:"user"`
	Project struct {
		ID                int             `json:"id" yaml:"id"`
		Name              string          `json:"name" yaml:"name"`
		Description       string          `json:"description" yaml:"description"`
		AvatarURL         string          `json:"avatar_url" yaml:"avatar_url"`
		GitSSHURL         string          `json:"git_ssh_url" yaml:"git_ssh_url"`
		GitHTTPURL        string          `json:"git_http_url" yaml:"git_http_url"`
		Namespace         string          `json:"namespace" yaml:"namespace"`
		PathWithNamespace string          `json:"path_with_namespace" yaml:"path_with_namespace"`
		DefaultBranch     string          `json:"default_branch" yaml:"default_branch"`
		Homepage          string          `json:"homepage" yaml:"homepage"`
		URL               string          `json:"url" yaml:"url"`
		SSHURL            string          `json:"ssh_url" yaml:"ssh_url"`
		HTTPURL           string          `json:"http_url" yaml:"http_url"`
		WebURL            string          `json:"web_url" yaml:"web_url"`
		Visibility        VisibilityValue `json:"visibility" yaml:"visibility"`
	} `json:"project" yaml:"project"`
	Commit struct {
		ID        string     `json:"id" yaml:"id"`
		Message   string     `json:"message" yaml:"message"`
		Timestamp *time.Time `json:"timestamp" yaml:"timestamp"`
		URL       string     `json:"url" yaml:"url"`
		Author    struct {
			Name  string `json:"name" yaml:"name"`
			Email string `json:"email" yaml:"email"`
		} `json:"author" yaml:"author"`
	} `json:"commit" yaml:"commit"`
	Builds []struct {
		ID         int    `json:"id" yaml:"id"`
		Stage      string `json:"stage" yaml:"stage"`
		Name       string `json:"name" yaml:"name"`
		Status     string `json:"status" yaml:"status"`
		CreatedAt  string `json:"created_at" yaml:"created_at"`
		StartedAt  string `json:"started_at" yaml:"started_at"`
		FinishedAt string `json:"finished_at" yaml:"finished_at"`
		When       string `json:"when" yaml:"when"`
		Manual     bool   `json:"manual" yaml:"manual"`
		User       struct {
			Name      string `json:"name" yaml:"name"`
			Username  string `json:"username" yaml:"username"`
			AvatarURL string `json:"avatar_url" yaml:"avatar_url"`
		} `json:"user" yaml:"user"`
		Runner struct {
			ID          int    `json:"id" yaml:"id"`
			Description string `json:"description" yaml:"description"`
			Active      bool   `json:"active" yaml:"active"`
			IsShared    bool   `json:"is_shared" yaml:"is_shared"`
		} `json:"runner" yaml:"runner"`
		ArtifactsFile struct {
			Filename string `json:"filename" yaml:"filename"`
			Size     int    `json:"size" yaml:"size"`
		} `json:"artifacts_file" yaml:"artifacts_file"`
	} `json:"builds" yaml:"builds"`
}

//BuildEvent represents a build event
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#build-events
type BuildEvent struct {
	ObjectKind        string  `json:"object_kind" yaml:"object_kind"`
	Ref               string  `json:"ref" yaml:"ref"`
	Tag               bool    `json:"tag" yaml:"tag"`
	BeforeSHA         string  `json:"before_sha" yaml:"before_sha"`
	SHA               string  `json:"sha" yaml:"sha"`
	BuildID           int     `json:"build_id" yaml:"build_id"`
	BuildName         string  `json:"build_name" yaml:"build_name"`
	BuildStage        string  `json:"build_stage" yaml:"build_stage"`
	BuildStatus       string  `json:"build_status" yaml:"build_status"`
	BuildStartedAt    string  `json:"build_started_at" yaml:"build_started_at"`
	BuildFinishedAt   string  `json:"build_finished_at" yaml:"build_finished_at"`
	BuildDuration     float64 `json:"build_duration" yaml:"build_duration"`
	BuildAllowFailure bool    `json:"build_allow_failure" yaml:"build_allow_failure"`
	ProjectID         int     `json:"project_id" yaml:"project_id"`
	ProjectName       string  `json:"project_name" yaml:"project_name"`
	User              struct {
		ID    int    `json:"id" yaml:"id"`
		Name  string `json:"name" yaml:"name"`
		Email string `json:"email" yaml:"email"`
	} `json:"user" yaml:"user"`
	Commit struct {
		ID          int    `json:"id" yaml:"id"`
		SHA         string `json:"sha" yaml:"sha"`
		Message     string `json:"message" yaml:"message"`
		AuthorName  string `json:"author_name" yaml:"author_name"`
		AuthorEmail string `json:"author_email" yaml:"author_email"`
		Status      string `json:"status" yaml:"status"`
		Duration    int    `json:"duration" yaml:"duration"`
		StartedAt   string `json:"started_at" yaml:"started_at"`
		FinishedAt  string `json:"finished_at" yaml:"finished_at"`
	} `json:"commit" yaml:"commit"`
	Repository *Repository `json:"repository" yaml:"repository"`
}
