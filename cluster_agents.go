//
// Copyright 2022, Timo Furrer <tuxtimo@gmail.com>
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

// ClusterAgentsService handles communication with the `cluster_agents` related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/cluster_agents.html
type ClusterAgentsService struct {
	client *Client
}

// ClusterAgent represents a GitLab Agent for Kubernetes
//
// GitLab API docs: https://docs.gitlab.com/ee/api/cluster_agents.html
type ClusterAgent struct {
	ID              int           `json:"id"`
	Name            string        `json:"name"`
	CreatedAt       *time.Time    `json:"created_at"`
	CreatedByUserID int           `json:"created_by_user_id"`
	ConfigProject   ConfigProject `json:"config_project"`
}

type ConfigProject struct {
	ID                int        `json:"id"`
	Description       string     `json:"description"`
	Name              string     `json:"name"`
	NameWithNamespace string     `json:"name_with_namespace"`
	Path              string     `json:"path"`
	PathWithNamespace string     `json:"path_with_namespace"`
	CreatedAt         *time.Time `json:"created_at"`
}

func (a ClusterAgent) String() string {
	return Stringify(a)
}

// ListClusterAgentsOptions represents the available ClusterAgents() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/cluster_agents.html#list-the-agents-for-a-project
type ListClusterAgentsOptions ListOptions

// ListClusterAgents returns a list of Agents for Kubernetes registered for the project
//
// GitLab API docs: https://docs.gitlab.com/ee/api/cluster_agents.html#list-the-agents-for-a-project
func (s *ClusterAgentsService) ListClusterAgents(pid interface{}, opt *ListClusterAgentsOptions, options ...RequestOptionFunc) ([]*ClusterAgent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, uri, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var t []*ClusterAgent
	resp, err := s.client.Do(req, &t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, err
}

// GetClusterAgent gets an Agent for Kubernetes by ID
//
// GitLab API docs: https://docs.gitlab.com/ee/api/cluster_agents.html#get-details-about-an-agent
func (s *ClusterAgentsService) GetClusterAgent(pid interface{}, id int, options ...RequestOptionFunc) (*ClusterAgent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents/%d", PathEscape(project), id)

	req, err := s.client.NewRequest(http.MethodGet, uri, nil, options)
	if err != nil {
		return nil, nil, err
	}

	t := new(ClusterAgent)
	resp, err := s.client.Do(req, t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, err
}

// RegisterClusterAgentOptions represents the available RegisterClusterAgent() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/cluster_agents.html#register-an-agent-with-a-project
type RegisterClusterAgentOptions struct {
	Name *string `url:"name,omitempty" json:"name,omitempty"`
}

// RegisterClusterAgent registers a new Agent for Kubernetes for a project
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/cluster_agents.html#register-an-agent-with-a-project
func (s *ClusterAgentsService) RegisterClusterAgent(pid interface{}, opt *RegisterClusterAgentOptions, options ...RequestOptionFunc) (*ClusterAgent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, uri, opt, options)
	if err != nil {
		return nil, nil, err
	}

	t := new(ClusterAgent)
	resp, err := s.client.Do(req, t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, err
}

// DeleteClusterAgent deletes a registered Agent for Kubernetes
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/cluster_agents.html#delete-a-registered-agent
func (s *ClusterAgentsService) DeleteClusterAgent(pid interface{}, id int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents/%d", PathEscape(project), id)

	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
