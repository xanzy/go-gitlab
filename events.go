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
	Title       string     `url:"title" json:"title"`
	ProjectID   int        `url:"project_id" json:"project_id"`
	ActionName  string     `url:"action_name" json:"action_name"`
	TargetID    int        `url:"target_id" json:"target_id"`
	TargetIID   int        `url:"target_iid" json:"target_iid"`
	TargetType  string     `url:"target_type" json:"target_type"`
	AuthorID    int        `url:"author_id" json:"author_id"`
	TargetTitle string     `url:"target_title" json:"target_title"`
	CreatedAt   *time.Time `url:"created_at" json:"created_at"`
	PushData    struct {
		CommitCount int    `url:"commit_count" json:"commit_count"`
		Action      string `url:"action" json:"action"`
		RefType     string `url:"ref_type" json:"ref_type"`
		CommitFrom  string `url:"commit_from" json:"commit_from"`
		CommitTo    string `url:"commit_to" json:"commit_to"`
		Ref         string `url:"ref" json:"ref"`
		CommitTitle string `url:"commit_title" json:"commit_title"`
	} `url:"push_data" json:"push_data"`
	Note   *Note `url:"note" json:"note"`
	Author struct {
		Name      string `url:"name" json:"name"`
		Username  string `url:"username" json:"username"`
		ID        int    `url:"id" json:"id"`
		State     string `url:"state" json:"state"`
		AvatarURL string `url:"avatar_url" json:"avatar_url"`
		WebURL    string `url:"web_url" json:"web_url"`
	} `url:"author" json:"author"`
	AuthorUsername string `url:"author_username" json:"author_username"`
}

// ListContributionEventsOptions represents the options for GetUserContributionEvents
//
// GitLap API docs:
// https://docs.gitlab.com/ce/api/events.html#get-user-contribution-events
type ListContributionEventsOptions struct {
	ListOptions
	Action     *EventTypeValue       `url:"action,omitempty" json:"action,omitempty"`
	TargetType *EventTargetTypeValue `url:"target_type,omitempty" json:"target_type,omitempty"`
	Before     *ISOTime              `url:"before,omitempty" json:"before,omitempty"`
	After      *ISOTime              `url:"after,omitempty" json:"after,omitempty"`
	Sort       *string               `url:"sort,omitempty" json:"sort,omitempty"`
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
