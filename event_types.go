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
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// PushEvent represents a push event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#push-events
type PushEvent struct {
	base
	Before            string      `json:"before"`
	After             string      `json:"after"`
	Ref               string      `json:"ref"`
	CheckoutSHA       string      `json:"checkout_sha"`
	UserID            int         `json:"user_id"`
	UserName          string      `json:"user_name"`
	UserUsername      string      `json:"user_username"`
	UserEmail         string      `json:"user_email"`
	UserAvatar        string      `json:"user_avatar"`
	ProjectID         int         `json:"project_id"`
	Repository        *Repository `json:"repository"`
	Commits           []*commit   `json:"commits"`
	TotalCommitsCount int         `json:"total_commits_count"`
}

// UnmarshalJSON of PushEvent adds field user to maintain common structure
func (e *PushEvent) UnmarshalJSON(b []byte) error {
	type Alias PushEvent
	o := &struct {
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(b, &o); err != nil {
		return err
	}
	e.User = &user{
		ID:        o.UserID,
		Email:     o.UserEmail,
		Name:      o.UserName,
		AvatarURL: o.UserAvatar,
	}
	return nil
}

// TagEvent represents a tag event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#tag-events
type TagEvent struct {
	base
	Before            string      `json:"before"`
	After             string      `json:"after"`
	Ref               string      `json:"ref"`
	CheckoutSHA       string      `json:"checkout_sha"`
	UserID            int         `json:"user_id"`
	UserName          string      `json:"user_name"`
	UserAvatar        string      `json:"user_avatar"`
	ProjectID         int         `json:"project_id"`
	Message           string      `json:"message"`
	Repository        *Repository `json:"repository"`
	Commits           []*commit   `json:"commits"`
	TotalCommitsCount int         `json:"total_commits_count"`
}

// IssueEvent represents a issue event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#issues-events
type IssueEvent struct {
	base
	User             *User       `json:"user"`
	Repository       *Repository `json:"repository"`
	ObjectAttributes *issue      `json:"object_attributes"`
	Assignee         *user       `json:"assignee"`
	Assignees        []*user     `json:"assignees"`
	Labels           []Label     `json:"labels"`
	Changes          struct {
		Labels struct {
			Previous []Label `json:"previous"`
			Current  []Label `json:"current"`
		} `json:"labels"`
		UpdatedByID struct {
			Previous int `json:"previous"`
			Current  int `json:"current"`
		} `json:"updated_by_id"`
	} `json:"changes"`
}

// JobEvent represents a job event.
//
// GitLab API docs:
// TODO: link to docs instead of src once they are published.
// https://gitlab.com/gitlab-org/gitlab-ce/blob/master/lib/gitlab/data_builder/build.rb
type JobEvent struct {
	base
	Ref               string  `json:"ref"`
	Tag               bool    `json:"tag"`
	BeforeSHA         string  `json:"before_sha"`
	SHA               string  `json:"sha"`
	BuildID           int     `json:"build_id"`
	BuildName         string  `json:"build_name"`
	BuildStage        string  `json:"build_stage"`
	BuildStatus       string  `json:"build_status"`
	BuildStartedAt    string  `json:"build_started_at"`
	BuildFinishedAt   string  `json:"build_finished_at"`
	BuildDuration     float64 `json:"build_duration"`
	BuildAllowFailure bool    `json:"build_allow_failure"`
	ProjectID         int     `json:"project_id"`
	ProjectName       string  `json:"project_name"`
	User              *user   `json:"user"`
	Commit            struct {
		ID          int    `json:"id"`
		SHA         string `json:"sha"`
		Message     string `json:"message"`
		AuthorName  string `json:"author_name"`
		AuthorEmail string `json:"author_email"`
		Status      string `json:"status"`
		Duration    int    `json:"duration"`
		StartedAt   string `json:"started_at"`
		FinishedAt  string `json:"finished_at"`
	} `json:"commit"`
	Repository *Repository `json:"repository"`
}

// CommitCommentEvent represents a comment on a commit event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#comment-on-commit
type CommitCommentEvent struct {
	base
	User             *user       `json:"user"`
	ProjectID        int         `json:"project_id"`
	Repository       *Repository `json:"repository"`
	ObjectAttributes *note       `json:"object_attributes"`
	Commit           *commit     `json:"commit"`
}

// MergeCommentEvent represents a comment on a merge event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#comment-on-merge-request
type MergeCommentEvent struct {
	base
	User             *user         `json:"user"`
	ProjectID        int           `json:"project_id"`
	ObjectAttributes *note         `json:"object_attributes"`
	Repository       *Repository   `json:"repository"`
	MergeRequest     *mergeRequest `json:"merge_request"`
}

// IssueCommentEvent represents a comment on an issue event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#comment-on-issue
type IssueCommentEvent struct {
	base
	User             *user       `json:"user"`
	ProjectID        int         `json:"project_id"`
	Repository       *Repository `json:"repository"`
	ObjectAttributes *note       `json:"object_attributes"`
	Issue            struct {
		ID                  int      `json:"id"`
		IID                 int      `json:"iid"`
		ProjectID           int      `json:"project_id"`
		MilestoneID         int      `json:"milestone_id"`
		AuthorID            int      `json:"author_id"`
		Description         string   `json:"description"`
		State               string   `json:"state"`
		Title               string   `json:"title"`
		LastEditedAt        string   `json:"last_edit_at"`
		LastEditedByID      int      `json:"last_edited_by_id"`
		UpdatedAt           string   `json:"updated_at"`
		UpdatedByID         int      `json:"updated_by_id"`
		CreatedAt           string   `json:"created_at"`
		ClosedAt            string   `json:"closed_at"`
		DueDate             *ISOTime `json:"due_date"`
		URL                 string   `json:"url"`
		TimeEstimate        int      `json:"time_estimate"`
		Confidential        bool     `json:"confidential"`
		TotalTimeSpent      int      `json:"total_time_spent"`
		HumanTotalTimeSpent int      `json:"human_total_time_spent"`
		HumanTimeEstimate   int      `json:"human_time_estimate"`
		AssigneeIDs         []int    `json:"assignee_ids"`
		AssigneeID          int      `json:"assignee_id"`
	} `json:"issue"`
}

// SnippetCommentEvent represents a comment on a snippet event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#comment-on-code-snippet
type SnippetCommentEvent struct {
	base
	User             *user       `json:"user"`
	ProjectID        int         `json:"project_id"`
	Repository       *Repository `json:"repository"`
	ObjectAttributes *note       `json:"object_attributes"`
	Snippet          *Snippet    `json:"snippet"`
}

// MergeEvent represents a merge event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#merge-request-events
type MergeEvent struct {
	base
	User             *user         `json:"user"`
	ObjectAttributes *mergeRequest `json:"object_attributes"`
	Repository       *Repository   `json:"repository"`
	Assignee         *user         `json:"assignee"`
	Labels           []Label       `json:"labels"`
	Changes          struct {
		AssigneeID struct {
			Previous int `json:"previous"`
			Current  int `json:"current"`
		} `json:"assignee_id"`
		Description struct {
			Previous string `json:"previous"`
			Current  string `json:"current"`
		} `json:"description"`
		Labels struct {
			Previous []Label `json:"previous"`
			Current  []Label `json:"current"`
		} `json:"labels"`
		UpdatedByID struct {
			Previous int `json:"previous"`
			Current  int `json:"current"`
		} `json:"updated_by_id"`
	} `json:"changes"`
}

// MergeParams represents the merge params.
type MergeParams struct {
	ForceRemoveSourceBranch bool `json:"force_remove_source_branch"`
}

// UnmarshalJSON decodes the merge parameters
//
// This allows support of ForceRemoveSourceBranch for both type bool (>11.9) and string (<11.9)
func (p *MergeParams) UnmarshalJSON(b []byte) error {
	type Alias MergeParams
	raw := struct {
		*Alias
		ForceRemoveSourceBranch interface{} `json:"force_remove_source_branch"`
	}{
		Alias: (*Alias)(p),
	}

	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	switch v := raw.ForceRemoveSourceBranch.(type) {
	case nil:
		// No action needed.
	case bool:
		p.ForceRemoveSourceBranch = v
	case string:
		p.ForceRemoveSourceBranch, err = strconv.ParseBool(v)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("failed to unmarshal ForceRemoveSourceBranch of type: %T", v)
	}

	return nil
}

// WikiPageEvent represents a wiki page event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#wiki-page-events
type WikiPageEvent struct {
	base
	User *user `json:"user"`
	Wiki struct {
		WebURL            string `json:"web_url"`
		GitSSHURL         string `json:"git_ssh_url"`
		GitHTTPURL        string `json:"git_http_url"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch     string `json:"default_branch"`
	} `json:"wiki"`
	ObjectAttributes struct {
		wikiPage
		Title   string `json:"title"`
		Format  string `json:"format"`
		Message string `json:"message"`
		Slug    string `json:"slug"`
		URL     string `json:"url"`
		Action  string `json:"action"`
	} `json:"object_attributes"`
}

// PipelineEvent represents a pipeline event.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#pipeline-events
type PipelineEvent struct {
	base
	ObjectAttributes struct {
		ID         int      `json:"id"`
		Ref        string   `json:"ref"`
		Tag        bool     `json:"tag"`
		SHA        string   `json:"sha"`
		BeforeSHA  string   `json:"before_sha"`
		Status     string   `json:"status"`
		Stages     []string `json:"stages"`
		CreatedAt  string   `json:"created_at"`
		FinishedAt string   `json:"finished_at"`
		Duration   int      `json:"duration"`
	} `json:"object_attributes"`
	Commit *commit `json:"commit"`
	Builds []struct {
		ID         int    `json:"id"`
		Stage      string `json:"stage"`
		Name       string `json:"name"`
		Status     string `json:"status"`
		CreatedAt  string `json:"created_at"`
		StartedAt  string `json:"started_at"`
		FinishedAt string `json:"finished_at"`
		When       string `json:"when"`
		Manual     bool   `json:"manual"`
		User       *user  `json:"user"`
		Runner     struct {
			ID          int    `json:"id"`
			Description string `json:"description"`
			Active      bool   `json:"active"`
			IsShared    bool   `json:"is_shared"`
		} `json:"runner"`
		ArtifactsFile struct {
			Filename string `json:"filename"`
			Size     int    `json:"size"`
		} `json:"artifacts_file"`
	} `json:"builds"`
}

//BuildEvent represents a build event
//
// GitLab API docs:
// https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#build-events
type BuildEvent struct {
	base
	Ref               string  `json:"ref"`
	Tag               bool    `json:"tag"`
	BeforeSHA         string  `json:"before_sha"`
	SHA               string  `json:"sha"`
	BuildID           int     `json:"build_id"`
	BuildName         string  `json:"build_name"`
	BuildStage        string  `json:"build_stage"`
	BuildStatus       string  `json:"build_status"`
	BuildStartedAt    string  `json:"build_started_at"`
	BuildFinishedAt   string  `json:"build_finished_at"`
	BuildDuration     float64 `json:"build_duration"`
	BuildAllowFailure bool    `json:"build_allow_failure"`
	ProjectID         int     `json:"project_id"`
	ProjectName       string  `json:"project_name"`
	Commit            struct {
		ID          int    `json:"id"`
		SHA         string `json:"sha"`
		Message     string `json:"message"`
		AuthorName  string `json:"author_name"`
		AuthorEmail string `json:"author_email"`
		Status      string `json:"status"`
		Duration    int    `json:"duration"`
		StartedAt   string `json:"started_at"`
		FinishedAt  string `json:"finished_at"`
	} `json:"commit"`
	Repository *Repository `json:"repository"`
}

// Cf. Hook Data Builder
// https://gitlab.com/gitlab-org/gitlab/blob/master/lib/gitlab/hook_data/merge_request_builder.rb
type mergeRequest struct {
	issuable
	AuthorID                  int         `json:"author_id"`
	CreatedAt                 string      `json:"created_at"`
	Description               string      `json:"description"`
	HeadPipelineID            int         `json:"head_pipeline_id"`
	ID                        int         `json:"id"`
	IID                       int         `json:"iid"`
	LastEditedAt              string      `json:"last_edited_at"`
	LastEditedByID            int         `json:"last_edited_by_id"`
	SourceBranch              string      `json:"source_branch"`
	SourceProjectID           int         `json:"source_project_id"`
	State                     string      `json:"state"`
	TargetBranch              string      `json:"target_branch"`
	TargetProjectID           int         `json:"target_project_id"`
	TimeEstimate              int         `json:"time_estimate"`
	Title                     string      `json:"title"`
	UpdatedAt                 string      `json:"updated_at"`
	UpdatedByID               int         `json:"updated_by_id"`
	MergeCommitSHA            string      `json:"merge_commit_sha"`
	MergeError                string      `json:"merge_error"`
	MergeParams               MergeParams `json:"merge_params"`
	MergeStatus               string      `json:"merge_status"`
	MergeUserID               int         `json:"merge_user_id"`
	MergeWhenPipelineSucceeds bool        `json:"merge_when_pipeline_succeeds"`
	MilestoneID               int         `json:"milestone_id"`
	DeletedAt                 string      `json:"deleted_at"`                   // ???
	InProgressMergeCommitSHA  string      `json:"in_progress_merge_commit_sha"` // ???
	LockVersion               int         `json:"lock_version"`                 // ???
	ApprovalsBeforeMerge      string      `json:"approvals_before_merge"`       // ???
	RebaseCommitSHA           string      `json:"rebase_commit_sha"`            // ???
	Squash                    bool        `json:"squash"`                       // ???
	Position                  int         `json:"position"`                     // ???
	LockedAt                  string      `json:"locked_at"`                    // ???
	URL                       string      `json:"url"`
	Source                    *Repository `json:"source"`
	Target                    *Repository `json:"target"`
	Action                    string      `json:"action"` // ???
	OldRev                    string      `json:"oldrev"` // ???
	WorkInProgress            bool        `json:"work_in_progress"`
	LastCommit                *commit     `json:"last_commit"`
	TotalTimeSpent            int         `json:"total_time_spent"`
	HumanTotalTimeSpent       int         `json:"human_total_time_spent"`
	HumanTimeEstimate         int         `json:"human_time_estimate"`
	AssigneeIDs               []int       `json:"assignee_ids"`
	AssigneeID                int         `json:"assignee_id"` // Deprecated
}

// Cf. Hook Data Builder
// https://gitlab.com/gitlab-org/gitlab/blob/master/lib/gitlab/hook_data/note_builder.rb
type note struct {
	Attachment       string `json:"attachment"`
	AuthorID         int    `json:"author_id"`
	ChangePosition   int    `json:"change_position"`
	CommitID         string `json:"commit_id"`
	CreatedAt        string `json:"created_at"`
	Description      string `json:"description"`
	DiscussionID     string `json:"discussion_id"`
	ID               int    `json:"id"`
	LineCode         string `json:"line_code"`
	Note             string `json:"note"`
	NoteableID       int    `json:"noteable_id"`
	NoteableType     string `json:"noteable_type"`
	OriginalPosition int    `json:"original_position"`
	Position         int    `json:"position"`
	ProjectID        int    `json:"project_id"`
	ResolvedAt       string `json:"resolved_at"`
	ResolvedByID     int    `json:"resolved_by_id"`
	ResolvedByPush   bool   `json:"resolved_by_push"`
	StDiff           *Diff  `json:"st_diff"`
	URL              string `json:"url"`
	System           bool   `json:"system"`
	Type             string `json:"type"`
	UpdatedAt        string `json:"updated_at"`
	UpdatedByID      int    `json:"updated_by_id"`
}

// Cf. Hook Data Builder
// https://gitlab.com/gitlab-org/gitlab/blob/master/lib/gitlab/hook_data/issue_builder.rb
type issue struct {
	issuable
	ID                  int      `json:"id"`
	IID                 int      `json:"iid"`
	MovedToID           int      `json:"moved_to_id"`
	DuplicatedToID      int      `json:"duplicated_to_id"`
	Position            int      `json:"position"`
	RelativePosition    int      `json:"relative_position"`
	ProjectID           int      `json:"project_id"`
	MilestoneID         int      `json:"milestone_id"`
	AuthorID            int      `json:"author_id"`
	Description         string   `json:"description"`
	State               string   `json:"state"`
	Title               string   `json:"title"`
	LastEditedAt        string   `json:"last_edit_at"`
	LastEditedByID      int      `json:"last_edited_by_id"`
	UpdatedAt           string   `json:"updated_at"`
	UpdatedByID         int      `json:"updated_by_id"`
	CreatedAt           string   `json:"created_at"`
	ClosedAt            string   `json:"closed_at"`
	DueDate             *ISOTime `json:"due_date"`
	URL                 string   `json:"url"`
	TimeEstimate        int      `json:"time_estimate"`
	Confidential        bool     `json:"confidential"`
	TotalTimeSpent      int      `json:"total_time_spent"`
	HumanTotalTimeSpent int      `json:"human_total_time_spent"`
	HumanTimeEstimate   int      `json:"human_time_estimate"`
	AssigneeIDs         []int    `json:"assignee_ids"`
	AssigneeID          int      `json:"assignee_id"` // Deprecated
	Labels              []Label  `json:"labels"`
}

// Cf. Hook Data Builder
// https://gitlab.com/gitlab-org/gitlab/blob/master/lib/gitlab/hook_data/issuable_builder.rb
type wikiPage struct {
	Content string `json:"content"`
}

// Cf. Hook Data Builder
// https://gitlab.com/gitlab-org/gitlab/blob/master/lib/gitlab/hook_data/issuable_builder.rb
type issuable struct {
	ObjectKind       string            `json:"object_kind"`
	EventType        string            `json:"event_type"`
	User             *user             `json:"user"`
	Project          *project          `json:"project"`
	ObjectAttributes *objectAttributes `json:"object_attributes"`
	Labels           []Label           `json:"labels"`
	Assignees        []*user           `json:"assignees"`
	Assignee         *user             `json:"assignee"`
	Changes          struct {
		Labels struct {
			Previous []Label `json:"previous"`
			Current  []Label `json:"current"`
		} `json:"labels"`
		UpdatedByID struct {
			Previous int `json:"previous"`
			Current  int `json:"current"`
		} `json:"updated_by_id"`
	} `json:"changes"`
}

type project struct {
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	AvatarURL         string          `json:"avatar_url"`
	GitSSHURL         string          `json:"git_ssh_url"`
	GitHTTPURL        string          `json:"git_http_url"`
	Namespace         string          `json:"namespace"`
	PathWithNamespace string          `json:"path_with_namespace"`
	DefaultBranch     string          `json:"default_branch"`
	Homepage          string          `json:"homepage"`
	URL               string          `json:"url"`
	SSHURL            string          `json:"ssh_url"`
	HTTPURL           string          `json:"http_url"`
	WebURL            string          `json:"web_url"`
	Visibility        VisibilityValue `json:"visibility"`
}

type objectAttributes struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	AssigneeID  int    `json:"assignee_id"`
	AuthorID    int    `json:"author_id"`
	ProjectID   int    `json:"project_id"`
	CreatedAt   string `json:"created_at"` // Should be *time.Time (see Gitlab issue #21468)
	UpdatedAt   string `json:"updated_at"` // Should be *time.Time (see Gitlab issue #21468)
	Position    int    `json:"position"`
	BranchName  string `json:"branch_name"`
	Description string `json:"description"`
	MilestoneID int    `json:"milestone_id"`
	State       string `json:"state"`
	IID         int    `json:"iid"`
	URL         string `json:"url"`
	Action      string `json:"action"`
}

type base struct {
	ObjectKind string   `json:"object_kind"`
	EventType  string   `json:"event_type"`
	User       *user    `json:"user"`
	Project    *project `json:"project"`
}

type user struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	ID        int    `json:"id"`
	Email     string `json:"email"`
}

type commit struct {
	ID        string     `json:"id"`
	Message   string     `json:"message"`
	Timestamp *time.Time `json:"timestamp"`
	URL       string     `json:"url"`
	Author    struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"author"`
	SHA      string   `json:"sha"`
	Added    []string `json:"added"`
	Modified []string `json:"modified"`
	Removed  []string `json:"removed"`
}
