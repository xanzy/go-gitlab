//
// Copyright 2021, Arkbriar
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
	"time"
)

// IssueLinksService handles communication with the issue relations related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/issue_links.html
type IssueLinksService struct {
	client *Client
}

// IssueLink represents a two-way relation between two issues.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/issue_links.html
type IssueLink struct {
	SourceIssue *Issue `json:"source_issue"`
	TargetIssue *Issue `json:"target_issue"`
	LinkType    string `json:"link_type"`
}

// IssueRelation gets a relation between two issues.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/issue_links.html#list-issue-relations
type IssueRelation struct {
	ID                   int                    `json:"id"`
	IID                  int                    `json:"iid"`
	ExternalID           string                 `json:"external_id"`
	State                string                 `json:"state"`
	Description          string                 `json:"description"`
	Author               *IssueAuthor           `json:"author"`
	Milestone            *Milestone             `json:"milestone"`
	ProjectID            int                    `json:"project_id"`
	Assignees            []*IssueAssignee       `json:"assignees"`
	Assignee             *IssueAssignee         `json:"assignee"`
	UpdatedAt            *time.Time             `json:"updated_at"`
	ClosedAt             *time.Time             `json:"closed_at"`
	ClosedBy             *IssueCloser           `json:"closed_by"`
	Title                string                 `json:"title"`
	CreatedAt            *time.Time             `json:"created_at"`
	MovedToID            int                    `json:"moved_to_id"`
	Labels               Labels                 `json:"labels"`
	LabelDetails         []*LabelDetails        `json:"label_details"`
	Upvotes              int                    `json:"upvotes"`
	Downvotes            int                    `json:"downvotes"`
	DueDate              *ISOTime               `json:"due_date"`
	WebURL               string                 `json:"web_url"`
	References           *IssueReferences       `json:"references"`
	TimeStats            *TimeStats             `json:"time_stats"`
	Confidential         bool                   `json:"confidential"`
	Weight               int                    `json:"weight"`
	DiscussionLocked     bool                   `json:"discussion_locked"`
	IssueType            *string                `json:"issue_type,omitempty"`
	Subscribed           bool                   `json:"subscribed"`
	UserNotesCount       int                    `json:"user_notes_count"`
	Links                *IssueLinks            `json:"_links"`
	IssueLinkID          int                    `json:"issue_link_id"`
	MergeRequestCount    int                    `json:"merge_requests_count"`
	EpicIssueID          int                    `json:"epic_issue_id"`
	Epic                 *Epic                  `json:"epic"`
	TaskCompletionStatus *TasksCompletionStatus `json:"task_completion_status"`
	LinkType             string                 `json:"link_type"`
	LinkCreatedAt        *time.Time             `json:"link_created_at"`
	LinkUpdatedAt        *time.Time             `json:"link_updated_at"`
}

// ListIssueRelations gets a list of related issues of a given issue,
// sorted by the relationship creation datetime (ascending).
//
// Issues will be filtered according to the user authorizations.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/issue_links.html#list-issue-relations
func (s *IssueLinksService) ListIssueRelations(pid interface{}, issueIID int, options ...RequestOptionFunc) ([]*IssueRelation, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/links", PathEscape(project), issueIID)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var is []*IssueRelation
	resp, err := s.client.Do(req, &is)
	if err != nil {
		return nil, resp, err
	}

	return is, resp, err
}

// CreateIssueLinkOptions represents the available CreateIssueLink() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/issue_links.html
type CreateIssueLinkOptions struct {
	TargetProjectID *string `json:"target_project_id"`
	TargetIssueIID  *string `json:"target_issue_iid"`
	LinkType        *string `json:"link_type"`
}

// CreateIssueLink creates a two-way relation between two issues.
// User must be allowed to update both issues in order to succeed.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/issue_links.html#create-an-issue-link
func (s *IssueLinksService) CreateIssueLink(pid interface{}, issueIID int, opt *CreateIssueLinkOptions, options ...RequestOptionFunc) (*IssueLink, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/links", PathEscape(project), issueIID)

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	i := new(IssueLink)
	resp, err := s.client.Do(req, &i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, err
}

// DeleteIssueLink deletes an issue link, thus removes the two-way relationship.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/issue_links.html#delete-an-issue-link
func (s *IssueLinksService) DeleteIssueLink(pid interface{}, issueIID, issueLinkID int, options ...RequestOptionFunc) (*IssueLink, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/links/%d",
		PathEscape(project),
		issueIID,
		issueLinkID)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	i := new(IssueLink)
	resp, err := s.client.Do(req, &i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, err
}
