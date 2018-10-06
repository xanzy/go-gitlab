//
// Copyright 2018, steperdin
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
	"net/http"
	"net/url"
	"time"
)

// DiscussionService handles communication with disscussion on related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/discussions.html
type DiscussionService struct {
	client *Client
}

// Discussion represents a GitLab Discussion.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/discussions.html
type Discussion struct {
	ID             string `json:"id"`
	IndividualNote bool   `json:"individual_note"`
	Notes          []struct {
		ID         int         `json:"id"`
		Type       string      `json:"type"`
		Body       string      `json:"body"`
		Attachment interface{} `json:"attachment"`
		Author     struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Username  string `json:"username"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		} `json:"author"`
		CreatedAt    time.Time   `json:"created_at"`
		UpdatedAt    time.Time   `json:"updated_at"`
		System       bool        `json:"system"`
		NoteableID   int         `json:"noteable_id"`
		NoteableType string      `json:"noteable_type"`
		NoteableIid  interface{} `json:"noteable_iid"`
		Resolved     bool        `json:"resolved"`
		Resolvable   bool        `json:"resolvable"`
		ResolvedBy   interface{} `json:"resolved_by"`
	} `json:"notes"`
}

const (
	discussionMergeRequest = "merge_requests"
	discussionIssue        = "issues"
	discussionSnippet      = "snippets"
	discussionCommit       = "commits"
)

// ListDiscussionsOptions represents the available options for listing discussions
// for each resources
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html
type ListDiscussionsOptions ListOptions

// ListMergeRequestDiscussions gets a list of all discussions on the merge request.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#list-project-merge-request-discussions
func (s *DiscussionService) ListMergeRequestDiscussions(pid interface{}, mergeRequestIID int, opt *ListDiscussionsOptions, options ...OptionFunc) ([]*Discussion, *Response, error) {
	return s.listDiscussions(pid, discussionMergeRequest, mergeRequestIID, opt, options...)
}

// ListIssueDiscussions gets a list of all discussions on the issue.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#list-project-issue-discussions
func (s *DiscussionService) ListIssueDiscussions(pid interface{}, issueIID int, opt *ListDiscussionsOptions, options ...OptionFunc) ([]*Discussion, *Response, error) {
	return s.listDiscussions(pid, discussionIssue, issueIID, opt, options...)
}

// ListSnippetDiscussions gets a list of all discussions on the snippet.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#list-project-snippet-discussions
func (s *DiscussionService) ListSnippetDiscussions(pid interface{}, snippetID int, opt *ListDiscussionsOptions, options ...OptionFunc) ([]*Discussion, *Response, error) {
	return s.listDiscussions(pid, discussionSnippet, snippetID, opt, options...)
}

// ListCommitDiscussions gets a list of all discussions on the commit.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#list-project-commit-discussions
func (s *DiscussionService) ListCommitDiscussions(pid interface{}, commitID int, opt *ListDiscussionsOptions, options ...OptionFunc) ([]*Discussion, *Response, error) {
	return s.listDiscussions(pid, discussionCommit, commitID, opt, options...)
}

func (s *DiscussionService) listDiscussions(pid interface{}, resource string, resourceID int, opt *ListDiscussionsOptions, options ...OptionFunc) ([]*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/%s/%d/discussions",
		url.QueryEscape(project),
		resource,
		resourceID,
	)

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var d []*Discussion
	resp, err := s.client.Do(req, &d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, err
}

// GetMergeRequestDiscussion get a discussion from merge request.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#get-single-merge-request-discussion
func (s *DiscussionService) GetMergeRequestDiscussion(pid interface{}, mergeRequestIID, discussionID int, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.getDiscussion(pid, discussionMergeRequest, mergeRequestIID, discussionID, options...)
}

// GetIssueDiscussion get a discussion from issue.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#get-single-issue-discussion
func (s *DiscussionService) GetIssueDiscussion(pid interface{}, issueIID, discussionID int, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.getDiscussion(pid, discussionIssue, issueIID, discussionID, options...)
}

// GetSnippetDiscussion get a discussion from snippet.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#get-single-snippet-discussion
func (s *DiscussionService) GetSnippetDiscussion(pid interface{}, snippetID, discussionID int, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.getDiscussion(pid, discussionSnippet, snippetID, discussionID, options...)
}

// GetCommitDiscussion get a discussion from commit.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#get-single-commit-discussion
func (s *DiscussionService) GetCommitDiscussion(pid interface{}, commitID, discussionID int, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.getDiscussion(pid, discussionCommit, commitID, discussionID, options...)
}

func (s *DiscussionService) getDiscussion(pid interface{}, resource string, resourceID, discussionID int, options ...OptionFunc) (*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/%s/%d/discussions/%d",
		url.QueryEscape(project),
		resource,
		resourceID,
		discussionID,
	)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	d := new(Discussion)
	resp, err := s.client.Do(req, &d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, err
}

// CreateDiscussionOptions represents the available options for discussion
// for a resource
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#create-new-merge-request-discussion
type CreateDiscussionOptions struct {
	Body string `url:"body,omitempty" json:"body"`
}

// CreateMergeRequestDiscussion create a discussion from merge request.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#create-new-merge-request-discussion
func (s *DiscussionService) CreateMergeRequestDiscussion(pid interface{}, mergeRequestIID int, opt *CreateDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.createDiscussion(pid, discussionMergeRequest, mergeRequestIID, opt, options...)
}

// CreateIssueDiscussion creates a new discussion to a single project issue.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#create-new-issue-discussion
func (s *DiscussionService) CreateIssueDiscussion(pid interface{}, issueIID int, opt *CreateDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.createDiscussion(pid, discussionIssue, issueIID, opt, options...)
}

// CreateSnippetDiscussion creates a new discussion to a single project snippet.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#create-new-snippet-discussion
func (s *DiscussionService) CreateSnippetDiscussion(pid interface{}, snippetIID int, opt *CreateDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.createDiscussion(pid, discussionSnippet, snippetIID, opt, options...)
}

// CreateCommitDiscussion creates a new discussion to a single project commit.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#create-new-commit-discussion
func (s *DiscussionService) CreateCommitDiscussion(pid interface{}, commitIID int, opt *CreateDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.createDiscussion(pid, discussionCommit, commitIID, opt, options...)
}

func (s *DiscussionService) createDiscussion(pid interface{}, resource string, resourceID int, opt *CreateDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/%s/%d/discussions",
		url.QueryEscape(project),
		resource,
		resourceID,
	)

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	d := new(Discussion)
	resp, err := s.client.Do(req, &d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, err
}

// AddNoteMergeRequestDiscussion adds a new note to the discussion.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#add-note-to-existing-merge-request-discussion
func (s *DiscussionService) AddNoteMergeRequestDiscussion(pid interface{}, mergeRequestIID, discussionID int, opt *CreateDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.addNoteDiscussion(pid, discussionMergeRequest, mergeRequestIID, discussionID, opt, options...)
}

// AddNoteIssueDiscussion adds a new note to the discussion.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#add-note-to-existing-issue-discussion
func (s *DiscussionService) AddNoteIssueDiscussion(pid interface{}, issueIID, discussionID int, opt *CreateDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.addNoteDiscussion(pid, discussionIssue, issueIID, discussionID, opt, options...)
}

// AddNoteSnippetDiscussion adds a new note to the discussion.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#add-note-to-existing-snippet-discussion
func (s *DiscussionService) AddNoteSnippetDiscussion(pid interface{}, snippetIID, discussionID int, opt *CreateDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.addNoteDiscussion(pid, discussionSnippet, snippetIID, discussionID, opt, options...)
}

// AddNoteCommitDiscussion adds a new note to the discussion.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#add-note-to-existing-commit-discussion
func (s *DiscussionService) AddNoteCommitDiscussion(pid interface{}, commitIID, discussionID int, opt *CreateDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.addNoteDiscussion(pid, discussionCommit, commitIID, discussionID, opt, options...)
}

func (s *DiscussionService) addNoteDiscussion(pid interface{}, resource string, resourceID, discussionID int, opt *CreateDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/%s/%d/discussions/%d/notes",
		url.QueryEscape(project),
		resource,
		resourceID,
		discussionID,
	)

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	d := new(Discussion)
	resp, err := s.client.Do(req, &d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, err
}

// ModifyDiscussionOptions represents the available options for discussion
// for a resource
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#modify-an-existing-merge-request-discussion-note
type ModifyDiscussionOptions struct {
	Body     *string `json:"body,omitempty"`
	Resolved *bool   `json:"resolved,omitempty"`
}

// ModifyNoteMergeRequestDiscussion modify or resolve an existing discussion note of a merge request.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#modify-an-existing-merge-request-discussion-note
func (s *DiscussionService) ModifyNoteMergeRequestDiscussion(pid interface{}, mergeRequestIID, discussionID, noteID int, opt *ModifyDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.modifyDiscussionNote(pid, discussionMergeRequest, mergeRequestIID, discussionID, noteID, opt, options...)
}

// ModifyNoteIssueDiscussion modify an existing discussion note of an issue.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#modify-existing-issue-discussion-note
func (s *DiscussionService) ModifyNoteIssueDiscussion(pid interface{}, issueIID, discussionID, noteID int, opt *ModifyDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.modifyDiscussionNote(pid, discussionIssue, issueIID, discussionID, noteID, opt, options...)
}

// ModifyNoteSnippetDiscussion modify an existing discussion note of a snippet.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#modify-existing-snippet-discussion-note
func (s *DiscussionService) ModifyNoteSnippetDiscussion(pid interface{}, snippetIID, discussionID, noteID int, opt *ModifyDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.modifyDiscussionNote(pid, discussionSnippet, snippetIID, discussionID, noteID, opt, options...)
}

// ModifyNoteCommitDiscussion modify or resolve an existing discussion note of a commit
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#modify-an-existing-commit-discussion-note
func (s *DiscussionService) ModifyNoteCommitDiscussion(pid interface{}, commitIID, discussionID, noteID int, opt *ModifyDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.modifyDiscussionNote(pid, discussionCommit, commitIID, discussionID, noteID, opt, options...)
}

func (s *DiscussionService) modifyDiscussionNote(pid interface{}, resource string, resourceID, discussionID, noteID int, opt *ModifyDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/%s/%d/discussions/%d/notes/%d",
		url.QueryEscape(project),
		resource,
		resourceID,
		discussionID,
		noteID,
	)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	d := new(Discussion)
	resp, err := s.client.Do(req, &d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, err
}

// DeleteNoteMergeRequestDiscussion deletes an existing discussion note of a merge request.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#delete-a-merge-request-discussion-note
func (s *DiscussionService) DeleteNoteMergeRequestDiscussion(pid interface{}, mergeRequestIID, discussionID, noteID int, options ...OptionFunc) (*Response, error) {
	return s.deleteDiscussionNote(pid, discussionMergeRequest, mergeRequestIID, discussionID, noteID, options...)
}

// DeleteNoteIssueDiscussion deletes an existing discussion note of an issue.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#delete-an-issue-discussion-note
func (s *DiscussionService) DeleteNoteIssueDiscussion(pid interface{}, issueIID, discussionID, noteID int, options ...OptionFunc) (*Response, error) {
	return s.deleteDiscussionNote(pid, discussionIssue, issueIID, discussionID, noteID, options...)
}

// DeleteNoteSnippetDiscussion deletes an existing discussion note of a snippet.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#delete-an-snippet-discussion-note
func (s *DiscussionService) DeleteNoteSnippetDiscussion(pid interface{}, snippetIID, discussionID, noteID int, options ...OptionFunc) (*Response, error) {
	return s.deleteDiscussionNote(pid, discussionSnippet, snippetIID, discussionID, noteID, options...)
}

// DeleteNoteCommitDiscussion deletes an existing discussion note of a commit.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#delete-a-commit-discussion-note
func (s *DiscussionService) DeleteNoteCommitDiscussion(pid interface{}, commitIID, discussionID, noteID int, options ...OptionFunc) (*Response, error) {
	return s.deleteDiscussionNote(pid, discussionCommit, commitIID, discussionID, noteID, options...)
}

func (s *DiscussionService) deleteDiscussionNote(pid interface{}, resource string, resourceID, discussionID, noteID int, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/%s/%d/discussions/%d/notes/%d",
		url.QueryEscape(project),
		resource,
		resourceID,
		discussionID,
		noteID,
	)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, nil)
}
