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
	"fmt"
	"net/url"
	"time"
)

// CommitsService handles communication with the commit related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html
type CommitsService struct {
	client *Client
}

// Commit represents a GitLab commit.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html
type Commit struct {
	ID             string           `json:"id" yaml:"id"`
	ShortID        string           `json:"short_id" yaml:"short_id"`
	Title          string           `json:"title" yaml:"title"`
	AuthorName     string           `json:"author_name" yaml:"author_name"`
	AuthorEmail    string           `json:"author_email" yaml:"author_email"`
	AuthoredDate   *time.Time       `json:"authored_date" yaml:"authored_date"`
	CommitterName  string           `json:"committer_name" yaml:"committer_name"`
	CommitterEmail string           `json:"committer_email" yaml:"committer_email"`
	CommittedDate  *time.Time       `json:"committed_date" yaml:"committed_date"`
	CreatedAt      *time.Time       `json:"created_at" yaml:"created_at"`
	Message        string           `json:"message" yaml:"message"`
	ParentIDs      []string         `json:"parent_ids" yaml:"parent_ids"`
	Stats          *CommitStats     `json:"stats" yaml:"stats"`
	Status         *BuildStateValue `json:"status" yaml:"status"`
}

// CommitStats represents the number of added and deleted files in a commit.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html
type CommitStats struct {
	Additions int `json:"additions" yaml:"additions"`
	Deletions int `json:"deletions" yaml:"deletions"`
	Total     int `json:"total" yaml:"total"`
}

func (c Commit) String() string {
	return Stringify(c)
}

// ListCommitsOptions represents the available ListCommits() options.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#list-repository-commits
type ListCommitsOptions struct {
	ListOptions
	RefName   *string    `url:"ref_name,omitempty" json:"ref_name,omitempty" yaml:"ref_name,omitempty"`
	Since     *time.Time `url:"since,omitempty" json:"since,omitempty" yaml:"since,omitempty"`
	Until     *time.Time `url:"until,omitempty" json:"until,omitempty" yaml:"until,omitempty"`
	Path      *string    `url:"path,omitempty" json:"path,omitempty" yaml:"path,omitempty"`
	All       *bool      `url:"all,omitempty" json:"all,omitempty" yaml:"all,omitempty"`
	WithStats *bool      `url:"with_stats,omitempty" json:"with_stats,omitempty" yaml:"with_stats,omitempty"`
}

// ListCommits gets a list of repository commits in a project.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#list-commits
func (s *CommitsService) ListCommits(pid interface{}, opt *ListCommitsOptions, options ...OptionFunc) ([]*Commit, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits", url.QueryEscape(project))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var c []*Commit
	resp, err := s.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

// FileAction represents the available actions that can be performed on a file.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#create-a-commit-with-multiple-files-and-actions
type FileAction string

// The available file actions.
const (
	FileCreate FileAction = "create"
	FileDelete FileAction = "delete"
	FileMove   FileAction = "move"
	FileUpdate FileAction = "update"
)

// CommitAction represents a single file action within a commit.
type CommitAction struct {
	Action       FileAction `url:"action" json:"action" yaml:"action"`
	FilePath     string     `url:"file_path" json:"file_path" yaml:"file_path"`
	PreviousPath string     `url:"previous_path,omitempty" json:"previous_path,omitempty" yaml:"previous_path,omitempty"`
	Content      string     `url:"content,omitempty" json:"content,omitempty" yaml:"content,omitempty"`
	Encoding     string     `url:"encoding,omitempty" json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

// CommitRef represents the reference of branches/tags in a commit.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/commits.html#get-references-a-commit-is-pushed-to
type CommitRef struct {
	Type string `json:"type" yaml:"type"`
	Name string `json:"name" yaml:"name"`
}

// GetCommitRefsOptions represents the available GetCommitRefs() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/commits.html#get-references-a-commit-is-pushed-to
type GetCommitRefsOptions struct {
	ListOptions
	Type *string `url:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
}

// GetCommitRefs gets all references (from branches or tags) a commit is pushed to
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/commits.html#get-references-a-commit-is-pushed-to
func (s *CommitsService) GetCommitRefs(pid interface{}, sha string, opt *GetCommitRefsOptions, options ...OptionFunc) ([]CommitRef, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits/%s/refs", url.QueryEscape(project), sha)

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var cs []CommitRef
	resp, err := s.client.Do(req, &cs)
	if err != nil {
		return nil, resp, err
	}

	return cs, resp, err
}

// GetCommit gets a specific commit identified by the commit hash or name of a
// branch or tag.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#get-a-single-commit
func (s *CommitsService) GetCommit(pid interface{}, sha string, options ...OptionFunc) (*Commit, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits/%s", url.QueryEscape(project), sha)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	c := new(Commit)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

// CreateCommitOptions represents the available options for a new commit.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#create-a-commit-with-multiple-files-and-actions
type CreateCommitOptions struct {
	Branch        *string         `url:"branch" json:"branch" yaml:"branch"`
	CommitMessage *string         `url:"commit_message" json:"commit_message" yaml:"commit_message"`
	StartBranch   *string         `url:"start_branch,omitempty" json:"start_branch,omitempty" yaml:"start_branch,omitempty"`
	Actions       []*CommitAction `url:"actions" json:"actions" yaml:"actions"`
	AuthorEmail   *string         `url:"author_email,omitempty" json:"author_email,omitempty" yaml:"author_email,omitempty"`
	AuthorName    *string         `url:"author_name,omitempty" json:"author_name,omitempty" yaml:"author_name,omitempty"`
}

// CreateCommit creates a commit with multiple files and actions.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#create-a-commit-with-multiple-files-and-actions
func (s *CommitsService) CreateCommit(pid interface{}, opt *CreateCommitOptions, options ...OptionFunc) (*Commit, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits", url.QueryEscape(project))

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var c *Commit
	resp, err := s.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

// Diff represents a GitLab diff.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html
type Diff struct {
	Diff        string `json:"diff" yaml:"diff"`
	NewPath     string `json:"new_path" yaml:"new_path"`
	OldPath     string `json:"old_path" yaml:"old_path"`
	AMode       string `json:"a_mode" yaml:"a_mode"`
	BMode       string `json:"b_mode" yaml:"b_mode"`
	NewFile     bool   `json:"new_file" yaml:"new_file"`
	RenamedFile bool   `json:"renamed_file" yaml:"renamed_file"`
	DeletedFile bool   `json:"deleted_file" yaml:"deleted_file"`
}

func (d Diff) String() string {
	return Stringify(d)
}

// GetCommitDiffOptions represents the available GetCommitDiff() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/commits.html#get-the-diff-of-a-commit
type GetCommitDiffOptions ListOptions

// GetCommitDiff gets the diff of a commit in a project..
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/commits.html#get-the-diff-of-a-commit
func (s *CommitsService) GetCommitDiff(pid interface{}, sha string, opt *GetCommitDiffOptions, options ...OptionFunc) ([]*Diff, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits/%s/diff", url.QueryEscape(project), sha)

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var d []*Diff
	resp, err := s.client.Do(req, &d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, err
}

// CommitComment represents a GitLab commit comment.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html
type CommitComment struct {
	Note     string `json:"note" yaml:"note"`
	Path     string `json:"path" yaml:"path"`
	Line     int    `json:"line" yaml:"line"`
	LineType string `json:"line_type" yaml:"line_type"`
	Author   Author `json:"author" yaml:"author"`
}

// Author represents a GitLab commit author
type Author struct {
	ID        int        `json:"id" yaml:"id"`
	Username  string     `json:"username" yaml:"username"`
	Email     string     `json:"email" yaml:"email"`
	Name      string     `json:"name" yaml:"name"`
	State     string     `json:"state" yaml:"state"`
	Blocked   bool       `json:"blocked" yaml:"blocked"`
	CreatedAt *time.Time `json:"created_at" yaml:"created_at"`
}

func (c CommitComment) String() string {
	return Stringify(c)
}

// GetCommitCommentsOptions represents the available GetCommitComments() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/commits.html#get-the-comments-of-a-commit
type GetCommitCommentsOptions ListOptions

// GetCommitComments gets the comments of a commit in a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/commits.html#get-the-comments-of-a-commit
func (s *CommitsService) GetCommitComments(pid interface{}, sha string, opt *GetCommitCommentsOptions, options ...OptionFunc) ([]*CommitComment, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits/%s/comments", url.QueryEscape(project), sha)

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var c []*CommitComment
	resp, err := s.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

// PostCommitCommentOptions represents the available PostCommitComment()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/commits.html#post-comment-to-commit
type PostCommitCommentOptions struct {
	Note     *string `url:"note,omitempty" json:"note,omitempty" yaml:"note,omitempty"`
	Path     *string `url:"path" json:"path" yaml:"path"`
	Line     *int    `url:"line" json:"line" yaml:"line"`
	LineType *string `url:"line_type" json:"line_type" yaml:"line_type"`
}

// PostCommitComment adds a comment to a commit. Optionally you can post
// comments on a specific line of a commit. Therefor both path, line_new and
// line_old are required.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/commits.html#post-comment-to-commit
func (s *CommitsService) PostCommitComment(pid interface{}, sha string, opt *PostCommitCommentOptions, options ...OptionFunc) (*CommitComment, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits/%s/comments", url.QueryEscape(project), sha)

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	c := new(CommitComment)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

// GetCommitStatusesOptions represents the available GetCommitStatuses() options.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#get-the-status-of-a-commit
type GetCommitStatusesOptions struct {
	ListOptions
	Ref   *string `url:"ref,omitempty" json:"ref,omitempty" yaml:"ref,omitempty"`
	Stage *string `url:"stage,omitempty" json:"stage,omitempty" yaml:"stage,omitempty"`
	Name  *string `url:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	All   *bool   `url:"all,omitempty" json:"all,omitempty" yaml:"all,omitempty"`
}

// CommitStatus represents a GitLab commit status.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#get-the-status-of-a-commit
type CommitStatus struct {
	ID          int        `json:"id" yaml:"id"`
	SHA         string     `json:"sha" yaml:"sha"`
	Ref         string     `json:"ref" yaml:"ref"`
	Status      string     `json:"status" yaml:"status"`
	Name        string     `json:"name" yaml:"name"`
	TargetURL   string     `json:"target_url" yaml:"target_url"`
	Description string     `json:"description" yaml:"description"`
	CreatedAt   *time.Time `json:"created_at" yaml:"created_at"`
	StartedAt   *time.Time `json:"started_at" yaml:"started_at"`
	FinishedAt  *time.Time `json:"finished_at" yaml:"finished_at"`
	Author      Author     `json:"author" yaml:"author"`
}

// GetCommitStatuses gets the statuses of a commit in a project.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#get-the-status-of-a-commit
func (s *CommitsService) GetCommitStatuses(pid interface{}, sha string, opt *GetCommitStatusesOptions, options ...OptionFunc) ([]*CommitStatus, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits/%s/statuses", url.QueryEscape(project), sha)

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var cs []*CommitStatus
	resp, err := s.client.Do(req, &cs)
	if err != nil {
		return nil, resp, err
	}

	return cs, resp, err
}

// SetCommitStatusOptions represents the available SetCommitStatus() options.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#post-the-status-to-commit
type SetCommitStatusOptions struct {
	State       BuildStateValue `url:"state" json:"state" yaml:"state"`
	Ref         *string         `url:"ref,omitempty" json:"ref,omitempty" yaml:"ref,omitempty"`
	Name        *string         `url:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Context     *string         `url:"context,omitempty" json:"context,omitempty" yaml:"context,omitempty"`
	TargetURL   *string         `url:"target_url,omitempty" json:"target_url,omitempty" yaml:"target_url,omitempty"`
	Description *string         `url:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
}

// SetCommitStatus sets the status of a commit in a project.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#post-the-status-to-commit
func (s *CommitsService) SetCommitStatus(pid interface{}, sha string, opt *SetCommitStatusOptions, options ...OptionFunc) (*CommitStatus, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/statuses/%s", url.QueryEscape(project), sha)

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var cs *CommitStatus
	resp, err := s.client.Do(req, &cs)
	if err != nil {
		return nil, resp, err
	}

	return cs, resp, err
}

// GetMergeRequestsByCommit gets merge request associated with a commit.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/commits.html#list-merge-requests-associated-with-a-commit
func (s *CommitsService) GetMergeRequestsByCommit(pid interface{}, sha string, options ...OptionFunc) ([]*MergeRequest, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits/%s/merge_requests",
		url.QueryEscape(project), url.QueryEscape(sha))

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var mrs []*MergeRequest
	resp, err := s.client.Do(req, &mrs)
	if err != nil {
		return nil, resp, err
	}

	return mrs, resp, err
}

// CherryPickCommitOptions represents the available options for cherry-picking a commit.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#cherry-pick-a-commit
type CherryPickCommitOptions struct {
	TargetBranch *string `url:"branch" json:"branch,omitempty" yaml:"branch,omitempty"`
}

// CherryPickCommit sherry picks a commit to a given branch.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#cherry-pick-a-commit
func (s *CommitsService) CherryPickCommit(pid interface{}, sha string, opt *CherryPickCommitOptions, options ...OptionFunc) (*Commit, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits/%s/cherry_pick",
		url.QueryEscape(project), url.QueryEscape(sha))

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var c *Commit
	resp, err := s.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}
