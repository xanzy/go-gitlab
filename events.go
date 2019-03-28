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

// EventsService handles communication with the event related methods of
// the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/events.html
type EventsService struct {
	client *Client
}

// ContributionEvent represents a user's contribution
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/events.html#get-user-contribution-events
type ContributionEvent struct {
	Title       string     `json:"title" yaml:"title"`
	ProjectID   int        `json:"project_id" yaml:"project_id"`
	ActionName  string     `json:"action_name" yaml:"action_name"`
	TargetID    int        `json:"target_id" yaml:"target_id"`
	TargetIID   int        `json:"target_iid" yaml:"target_iid"`
	TargetType  string     `json:"target_type" yaml:"target_type"`
	AuthorID    int        `json:"author_id" yaml:"author_id"`
	TargetTitle string     `json:"target_title" yaml:"target_title"`
	CreatedAt   *time.Time `json:"created_at" yaml:"created_at"`
	PushData    struct {
		CommitCount int    `json:"commit_count" yaml:"commit_count"`
		Action      string `json:"action" yaml:"action"`
		RefType     string `json:"ref_type" yaml:"ref_type"`
		CommitFrom  string `json:"commit_from" yaml:"commit_from"`
		CommitTo    string `json:"commit_to" yaml:"commit_to"`
		Ref         string `json:"ref" yaml:"ref"`
		CommitTitle string `json:"commit_title" yaml:"commit_title"`
	} `json:"push_data" yaml:"push_data"`
	Note   *Note `json:"note" yaml:"note"`
	Author struct {
		Name      string `json:"name" yaml:"name"`
		Username  string `json:"username" yaml:"username"`
		ID        int    `json:"id" yaml:"id"`
		State     string `json:"state" yaml:"state"`
		AvatarURL string `json:"avatar_url" yaml:"avatar_url"`
		WebURL    string `json:"web_url" yaml:"web_url"`
	} `json:"author" yaml:"author"`
	AuthorUsername string `json:"author_username" yaml:"author_username"`
}

// ListContributionEventsOptions represents the options for GetUserContributionEvents
//
// GitLap API docs:
// https://docs.gitlab.com/ce/api/events.html#get-user-contribution-events
type ListContributionEventsOptions struct {
	ListOptions
	Action     *EventTypeValue       `url:"action,omitempty" json:"action,omitempty" yaml:"action,omitempty"`
	TargetType *EventTargetTypeValue `url:"target_type,omitempty" json:"target_type,omitempty" yaml:"target_type,omitempty"`
	Before     *ISOTime              `url:"before,omitempty" json:"before,omitempty" yaml:"before,omitempty"`
	After      *ISOTime              `url:"after,omitempty" json:"after,omitempty" yaml:"after,omitempty"`
	Sort       *string               `url:"sort,omitempty" json:"sort,omitempty" yaml:"sort,omitempty"`
}

// ListUserContributionEvents retrieves user contribution events
// for the specified user, sorted from newest to oldest.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/events.html#get-user-contribution-events
func (s *UsersService) ListUserContributionEvents(uid interface{}, opt *ListContributionEventsOptions, options ...OptionFunc) ([]*ContributionEvent, *Response, error) {
	user, err := parseID(uid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("users/%s/events", user)

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var cs []*ContributionEvent
	resp, err := s.client.Do(req, &cs)
	if err != nil {
		return nil, resp, err
	}

	return cs, resp, err
}

// ListCurrentUserContributionEvents gets a list currently authenticated user's events
//
// GitLab API docs: https://docs.gitlab.com/ce/api/events.html#list-currently-authenticated-user-39-s-events
func (s *EventsService) ListCurrentUserContributionEvents(opt *ListContributionEventsOptions, options ...OptionFunc) ([]*ContributionEvent, *Response, error) {
	req, err := s.client.NewRequest("GET", "events", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var cs []*ContributionEvent
	resp, err := s.client.Do(req, &cs)
	if err != nil {
		return nil, resp, err
	}

	return cs, resp, err
}

// ListProjectVisibleEvents gets a list of visible events for a particular project
//
// GitLab API docs: https://docs.gitlab.com/ee/api/events.html#list-a-project-s-visible-events
func (s *EventsService) ListProjectVisibleEvents(pid interface{}, opt *ListContributionEventsOptions, options ...OptionFunc) ([]*ContributionEvent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/events", url.QueryEscape(project))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var cs []*ContributionEvent
	resp, err := s.client.Do(req, &cs)
	if err != nil {
		return nil, resp, err
	}

	return cs, resp, err
}
