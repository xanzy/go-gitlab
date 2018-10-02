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

// DiscussionsService handles communication with the discussions related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/discussions.html
type DiscussionsService struct {
	client *Client
}

// Discussion represents a GitLab discussion.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/discussions.html
type Discussion struct {
	ID             string `json:"id"`
	IndividualNote bool   `json:"individual_note"`
	Notes          []Note `json:"notes"`
}

func (n Discussion) String() string {
	return Stringify(n)
}

// ListIssueDiscussionsOptions represents the available ListIssueDiscussions() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#list-project-issue-discussions
type ListIssueDiscussionsOptions struct {
	ListOptions
	OrderBy *string `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort    *string `url:"sort,omitempty" json:"sort,omitempty"`
}

// ListIssueDiscussions gets a list of all discussions for a single issue.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#list-project-issue-discussions
func (s *DiscussionsService) ListIssueDiscussions(pid interface{}, issue int, opt *ListIssueDiscussionsOptions, options ...OptionFunc) ([]*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/discussions", url.QueryEscape(project), issue)

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var n []*Discussion
	resp, err := s.client.Do(req, &n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}

// GetIssueDiscussion returns a single discussion for a specific project issue.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#get-single-issue-discussion
func (s *DiscussionsService) GetIssueDiscussion(pid interface{}, issue, discussion int, options ...OptionFunc) (*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/discussions/%d", url.QueryEscape(project), issue, discussion)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	n := new(Discussion)
	resp, err := s.client.Do(req, n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}

// CreateIssueDiscussionOptions represents the available CreateIssueDiscussion()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#create-new-issue-discussion
type CreateIssueDiscussionOptions struct {
	Body *string `url:"body,omitempty" json:"body,omitempty"`
}

// CreateIssueDiscussion creates a new discussion to a single project issue.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#create-new-issue-discussion
func (s *DiscussionsService) CreateIssueDiscussion(pid interface{}, issue int, opt *CreateIssueDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/discussions", url.QueryEscape(project), issue)

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	n := new(Discussion)
	resp, err := s.client.Do(req, n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}

// UpdateIssueDiscussionOptions represents the available UpdateIssueDiscussion()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#modify-existing-issue-discussion
type UpdateIssueDiscussionOptions struct {
	Body *string `url:"body,omitempty" json:"body,omitempty"`
}

// UpdateIssueDiscussion modifies existing discussion of an issue.
//
// https://docs.gitlab.com/ce/api/discussions.html#modify-existing-issue-discussion
func (s *DiscussionsService) UpdateIssueDiscussion(pid interface{}, issue, discussion int, opt *UpdateIssueDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/discussions/%d", url.QueryEscape(project), issue, discussion)

	req, err := s.client.NewRequest("PUT", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	n := new(Discussion)
	resp, err := s.client.Do(req, n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}

// DeleteIssueDiscussion deletes an existing discussion of an issue.
//
// https://docs.gitlab.com/ce/api/discussions.html#delete-an-issue-discussion
func (s *DiscussionsService) DeleteIssueDiscussion(pid interface{}, issue, discussion int, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/discussions/%d", url.QueryEscape(project), issue, discussion)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListSnippetDiscussionsOptions represents the available ListSnippetDiscussions() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#list-all-snippet-discussions
type ListSnippetDiscussionsOptions ListOptions

// ListSnippetDiscussions gets a list of all discussions for a single snippet. Snippet
// discussions are comments users can post to a snippet.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#list-all-snippet-discussions
func (s *DiscussionsService) ListSnippetDiscussions(pid interface{}, snippet int, opt *ListSnippetDiscussionsOptions, options ...OptionFunc) ([]*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/snippets/%d/discussions", url.QueryEscape(project), snippet)

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var n []*Discussion
	resp, err := s.client.Do(req, &n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}

// GetSnippetDiscussion returns a single discussion for a given snippet.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#get-single-snippet-discussion
func (s *DiscussionsService) GetSnippetDiscussion(pid interface{}, snippet, discussion int, options ...OptionFunc) (*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/snippets/%d/discussions/%d", url.QueryEscape(project), snippet, discussion)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	n := new(Discussion)
	resp, err := s.client.Do(req, n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}

// CreateSnippetDiscussionOptions represents the available CreateSnippetDiscussion()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#create-new-snippet-discussion
type CreateSnippetDiscussionOptions struct {
	Body *string `url:"body,omitempty" json:"body,omitempty"`
}

// CreateSnippetDiscussion creates a new discussion for a single snippet. Snippet discussions are
// comments users can post to a snippet.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#create-new-snippet-discussion
func (s *DiscussionsService) CreateSnippetDiscussion(pid interface{}, snippet int, opt *CreateSnippetDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/snippets/%d/discussions", url.QueryEscape(project), snippet)

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	n := new(Discussion)
	resp, err := s.client.Do(req, n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}

// UpdateSnippetDiscussionOptions represents the available UpdateSnippetDiscussion()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#modify-existing-snippet-discussion
type UpdateSnippetDiscussionOptions struct {
	Body *string `url:"body,omitempty" json:"body,omitempty"`
}

// UpdateSnippetDiscussion modifies existing discussion of a snippet.
//
// https://docs.gitlab.com/ce/api/discussions.html#modify-existing-snippet-discussion
func (s *DiscussionsService) UpdateSnippetDiscussion(pid interface{}, snippet, discussion int, opt *UpdateSnippetDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/snippets/%d/discussions/%d", url.QueryEscape(project), snippet, discussion)

	req, err := s.client.NewRequest("PUT", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	n := new(Discussion)
	resp, err := s.client.Do(req, n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}

// DeleteSnippetDiscussion deletes an existing discussion of a snippet.
//
// https://docs.gitlab.com/ce/api/discussions.html#delete-a-snippet-discussion
func (s *DiscussionsService) DeleteSnippetDiscussion(pid interface{}, snippet, discussion int, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/snippets/%d/discussions/%d", url.QueryEscape(project), snippet, discussion)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListMergeRequestDiscussionsOptions represents the available ListMergeRequestDiscussions()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#list-all-merge-request-discussions
type ListMergeRequestDiscussionsOptions ListOptions

// ListMergeRequestDiscussions gets a list of all discussions for a single merge request.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#list-all-merge-request-discussions
func (s *DiscussionsService) ListMergeRequestDiscussions(pid interface{}, mergeRequest int, opt *ListMergeRequestDiscussionsOptions, options ...OptionFunc) ([]*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_requests/%d/discussions", url.QueryEscape(project), mergeRequest)

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var n []*Discussion
	resp, err := s.client.Do(req, &n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}

// GetMergeRequestDiscussion returns a single discussion for a given merge request.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#get-single-merge-request-discussion
func (s *DiscussionsService) GetMergeRequestDiscussion(pid interface{}, mergeRequest, discussion int, options ...OptionFunc) (*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_requests/%d/discussions/%d", url.QueryEscape(project), mergeRequest, discussion)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	n := new(Discussion)
	resp, err := s.client.Do(req, n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}

// CreateMergeRequestDiscussionOptions represents the available
// CreateMergeRequestDiscussion() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#create-new-merge-request-discussion
type CreateMergeRequestDiscussionOptions struct {
	Body      *string      `url:"body,omitempty" json:"body,omitempty"`
	CreatedAt time.Time    `url:"created_at,omitempty"`
	Position  NotePosition `json:"position,omitempty"`
}

// CreateMergeRequestDiscussion creates a new discussion for a single merge request.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#create-new-merge-request-discussion
func (s *DiscussionsService) CreateMergeRequestDiscussion(pid interface{}, mergeRequest int, opt *CreateMergeRequestDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_requests/%d/discussions", url.QueryEscape(project), mergeRequest)

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	n := new(Discussion)
	resp, err := s.client.Do(req, n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}

// UpdateMergeRequestDiscussionOptions represents the available
// UpdateMergeRequestDiscussion() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#modify-existing-merge-request-discussion
type UpdateMergeRequestDiscussionOptions struct {
	Body *string `url:"body,omitempty" json:"body,omitempty"`
}

// UpdateMergeRequestDiscussion modifies existing discussion of a merge request.
//
// https://docs.gitlab.com/ce/api/discussions.html#modify-existing-merge-request-discussion
func (s *DiscussionsService) UpdateMergeRequestDiscussion(pid interface{}, mergeRequest, discussion int, opt *UpdateMergeRequestDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf(
		"projects/%s/merge_requests/%d/discussions/%d", url.QueryEscape(project), mergeRequest, discussion)
	req, err := s.client.NewRequest("PUT", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	n := new(Discussion)
	resp, err := s.client.Do(req, n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}

// DeleteMergeRequestDiscussion deletes an existing discussion of a merge request.
//
// https://docs.gitlab.com/ce/api/discussions.html#delete-a-merge-request-discussion
func (s *DiscussionsService) DeleteMergeRequestDiscussion(pid interface{}, mergeRequest, discussion int, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf(
		"projects/%s/merge_requests/%d/discussions/%d", url.QueryEscape(project), mergeRequest, discussion)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
