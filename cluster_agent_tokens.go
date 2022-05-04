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

// ClusterAgentTokensService handles communication with the `cluster_agents` related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/cluster_agent_tokens.html
type ClusterAgentTokensService struct {
	client *Client
}

// ClusterAgentToken represents a GitLab Agent for Kubernetes
//
// The `last_used_at` field is not populated when listing all tokens.
// The `token` field only populated in the response of the `CreateClusterAgentToken` method.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/cluster_agent_tokens.html
type ClusterAgentToken struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	AgentID         int        `json:"agent_id"`
	Status          string     `json:"status"`
	CreatedAt       *time.Time `json:"created_at"`
	CreatedByUserID int        `json:"created_by_user_id"`
	LastUsedAt      *time.Time `json:"last_used_at"`
	Token           string     `json:"token"`
}

func (a ClusterAgentToken) String() string {
	return Stringify(a)
}

// ListClusterAgentTokensOptions represents the available ListClusterAgentTokens() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/cluster_agent_tokens.html#list-tokens-for-an-agent
type ListClusterAgentTokensOptions ListOptions

// ListClusterAgentTokens returns a list of Tokens for the Agents for Kubernetes registered for the project
//
// GitLab API docs: https://docs.gitlab.com/ee/api/cluster_agent_tokens.html#list-tokens-for-an-agent
func (s *ClusterAgentTokensService) ListClusterAgentTokens(pid interface{}, aid int, opt *ListClusterAgentTokensOptions, options ...RequestOptionFunc) ([]*ClusterAgentToken, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents/%d/tokens", PathEscape(project), aid)

	req, err := s.client.NewRequest(http.MethodGet, uri, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var t []*ClusterAgentToken
	resp, err := s.client.Do(req, &t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, err
}

// GetClusterAgentToken gets a Token of an Agent for Kubernetes by ID
//
// GitLab API docs: https://docs.gitlab.com/ee/api/cluster_agent_tokens.html#get-a-single-agent-token
func (s *ClusterAgentTokensService) GetClusterAgentToken(pid interface{}, aid int, id int, options ...RequestOptionFunc) (*ClusterAgentToken, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents/%d/tokens/%d", PathEscape(project), aid, id)

	req, err := s.client.NewRequest(http.MethodGet, uri, nil, options)
	if err != nil {
		return nil, nil, err
	}

	t := new(ClusterAgentToken)
	resp, err := s.client.Do(req, t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, err
}

// CreateClusterAgentTokenOptions represents the available CreateClusterAgentToken() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/cluster_agent_tokens.html#create-an-agent-token
type CreateClusterAgentTokenOptions struct {
	Name        *string `url:"name,omitempty" json:"name,omitempty"`
	Description *string `url:"description,omitempty" json:"description,omitempty"`
}

// CreateClusterAgentToken creates a new Token for an Agent for Kubernetes for a project
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/cluster_agent_tokens.html#create-an-agent-token
func (s *ClusterAgentTokensService) CreateClusterAgentToken(pid interface{}, aid int, opt *CreateClusterAgentTokenOptions, options ...RequestOptionFunc) (*ClusterAgentToken, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents/%d/tokens", PathEscape(project), aid)

	req, err := s.client.NewRequest(http.MethodPost, uri, opt, options)
	if err != nil {
		return nil, nil, err
	}

	t := new(ClusterAgentToken)
	resp, err := s.client.Do(req, t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, err
}

// RevokeClusterAgentToken revokes a Token of an Agent for Kubernetes
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/cluster_agent_tokens.html#revoke-an-agent-token
func (s *ClusterAgentTokensService) RevokeClusterAgentToken(pid interface{}, aid int, id int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents/%d/tokens/%d", PathEscape(project), aid, id)

	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
