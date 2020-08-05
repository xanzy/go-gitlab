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
	"time"
)

// IssuesStatisticsService handles communication with the issues statistics related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/issues_statistics.html
type IssuesStatisticsService struct {
	client *Client
}

// Counts represents a GitLab issues counts statistic.
type Counts struct {
	All    int `json:"all"`
	Closed int `json:"closed"`
	Opened int `json:"opened"`
}

// Statistics represents a GitLab issues statistic.
type Statistics struct {
	Counts Counts `json:"counts"`
}

// Statistics represents a GitLab issues statistic.
// GitLab API docs: https://docs.gitlab.com/ee/api/issues_statistics.html
type IssuesStatistics struct {
	Statistics Statistics `json:"statistics"`
}

func (n IssuesStatistics) String() string {
	return Stringify(n)
}

// GetIssuesStatisticsOptions represents the available GetIssuesStatistics() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/issues_statistics.html#get-issues-statistics
type GetIssuesStatisticsOptions struct {
	Labels           *Labels    `url:"labels,omitempty" json:"labels,omitempty"`
	Milestone        *Milestone `url:"milestone,omitempty" json:"milestone,omitempty"`
	Scope            *string    `url:"scope,omitempty" json:"scope,omitempty"`
	AuthorID         *int       `url:"author_id,omitempty" json:"author_id,omitempty"`
	AuthorUsername   *string    `url:"author_username,omitempty" json:"author_username,omitempty"`
	AssigneeID       *int       `url:"assignee_id,omitempty" json:"assignee_id,omitempty"`
	AssigneeUsername []string   `url:"assignee_username,omitempty" json:"assignee_username,omitempty"`
	MyReactionEmoji  *string    `url:"my_reaction_emoji,omitempty" json:"my_reaction_emoji,omitempty"`
	IIDs             []int      `url:"iids,omitempty" json:"iids,omitempty"`
	Search           *string    `url:"search,omitempty" json:"search,omitempty"`
	In               *string    `url:"in,omitempty" json:"in,omitempty"`
	CreatedAfter     *time.Time `url:"created_after,omitempty" json:"created_after,omitempty"`
	CreatedBefore    *time.Time `url:"created_before,omitempty" json:"created_before,omitempty"`
	UpdatedAfter     *time.Time `url:"updated_after,omitempty" json:"updated_after,omitempty"`
	UpdatedBefore    *time.Time `url:"updated_before,omitempty" json:"updated_before,omitempty"`
	Confidential     *bool      `url:"confidential,omitempty" json:"confidential,omitempty"`
}

// GetIssuesStatistics
// Gets issues count statistics on all issues the authenticated user has access to.
// By default it returns only issues created by the current user. To get all issues, use parameter scope=all.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/issues_statistics.html#get-issues-statistics
func (s *IssuesStatisticsService) GetIssuesStatistics(opt *GetIssuesStatisticsOptions, options ...RequestOptionFunc) (*IssuesStatistics, *Response, error) {
	req, err := s.client.NewRequest("GET", "issues_statistics", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var n *IssuesStatistics
	resp, err := s.client.Do(req, &n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}

// GetGroupIssuesStatisticsOptions represents the available GetGroupIssuesStatistics() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/issues_statistics.html#get-group-issues-statistics
type GetGroupIssuesStatisticsOptions struct {
	ID               int        `url:"id,omitempty" json:"id,omitempty"`
	Labels           *Labels    `url:"labels,omitempty" json:"labels,omitempty"`
	IIDs             []int      `url:"iids,omitempty" json:"iids,omitempty"`
	Milestone        *Milestone `url:"milestone,omitempty" json:"milestone,omitempty"`
	Scope            *string    `url:"scope,omitempty" json:"scope,omitempty"`
	AuthorID         *int       `url:"author_id,omitempty" json:"author_id,omitempty"`
	AuthorUsername   *string    `url:"author_username,omitempty" json:"author_username,omitempty"`
	AssigneeID       *int       `url:"assignee_id,omitempty" json:"assignee_id,omitempty"`
	AssigneeUsername []string   `url:"assignee_username,omitempty" json:"assignee_username,omitempty"`
	MyReactionEmoji  *string    `url:"my_reaction_emoji,omitempty" json:"my_reaction_emoji,omitempty"`
	Search           *string    `url:"search,omitempty" json:"search,omitempty"`
	In               *string    `url:"in,omitempty" json:"in,omitempty"`
	CreatedAfter     *time.Time `url:"created_after,omitempty" json:"created_after,omitempty"`
	CreatedBefore    *time.Time `url:"created_before,omitempty" json:"created_before,omitempty"`
	UpdatedAfter     *time.Time `url:"updated_after,omitempty" json:"updated_after,omitempty"`
	UpdatedBefore    *time.Time `url:"updated_before,omitempty" json:"updated_before,omitempty"`
	Confidential     *bool      `url:"confidential,omitempty" json:"confidential,omitempty"`
}

// GetIssuesStatistics
// Gets issues count statistics for given group.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/issues_statistics.html#get-group-issues-statistics
func (s *IssuesStatisticsService) GetGroupIssuesStatistics(pid interface{}, opt *GetIssuesStatisticsOptions, options ...RequestOptionFunc) (*IssuesStatistics, *Response, error) {
	group, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/issues_statistics", pathEscape(group))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var n *IssuesStatistics
	resp, err := s.client.Do(req, &n)
	if err != nil {
		return nil, resp, err
	}

	return n, resp, err
}
