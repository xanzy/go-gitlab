//
// Copyright 2015, Sander van Harmelen
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
)

// ServicesService handles communication with the services related methods of
// the GitLab API.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/services.html
type ServicesService struct {
	client *Client
}

// SetGitLabCIServiceOptions represents the available SetGitLabCIService()
// options.
//
// GitLab API docs:
// http://doc.gitlab.com/ce/api/services.html#edit-gitlab-ci-service
type SetGitLabCIServiceOptions struct {
	Token      string `url:"token,omitempty"`
	ProjectURL string `url:"project_url,omitempty"`
}

// SetGitLabCIService sets GitLab CI service for a project.
//
// GitLab API docs:
// http://doc.gitlab.com/ce/api/services.html#edit-gitlab-ci-service
func (s *ServicesService) SetGitLabCIService(
	pid interface{},
	opt *SetGitLabCIServiceOptions) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/gitlab-ci", url.QueryEscape(project))

	req, err := s.client.NewRequest("PUT", u, opt)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// DeleteGitLabCIService deletes GitLab CI service settings for a project.
//
// GitLab API docs:
// http://doc.gitlab.com/ce/api/services.html#delete-gitlab-ci-service
func (s *ServicesService) DeleteGitLabCIService(pid interface{}) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/gitlab-ci", url.QueryEscape(project))

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// SetHipChatServiceOptions represents the available SetHipChatService()
// options.
//
// GitLab API docs:
// http://doc.gitlab.com/ce/api/services.html#edit-hipchat-service
type SetHipChatServiceOptions struct {
	Token string `url:"token,omitempty"`
	Room  string `url:"room,omitempty"`
}

// SetHipChatService sets HipChat service for a project
//
// GitLab API docs:
// http://doc.gitlab.com/ce/api/services.html#edit-hipchat-service
func (s *ServicesService) SetHipChatService(
	pid interface{},
	opt *SetHipChatServiceOptions) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/hipchat", url.QueryEscape(project))

	req, err := s.client.NewRequest("PUT", u, opt)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// DeleteHipChatService deletes HipChat service for project.
//
// GitLab API docs:
// http://doc.gitlab.com/ce/api/services.html#delete-hipchat-service
func (s *ServicesService) DeleteHipChatService(pid interface{}) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/hipchat", url.QueryEscape(project))

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
