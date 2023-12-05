//
// Copyright 2023, James Hong
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

// GroupServiceAccount represents a GitLab service account user.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/groups.html#create-service-account-user
type GroupServiceAccount struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

// CreateServiceAccount create a new service account user for a group.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/groups.html#create-service-account-user
func (s *GroupsService) CreateServiceAccount(gid interface{}, options ...RequestOptionFunc) (*GroupServiceAccount, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/service_accounts", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	sa := new(GroupServiceAccount)
	resp, err := s.client.Do(req, sa)
	if err != nil {
		return nil, resp, err
	}

	return sa, resp, nil
}

// GroupServiceAccountPAT represents a GitLab service account Personal Access Token.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/groups.html#create-personal-access-token-for-service-account-user
type GroupServiceAccountPAT struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Revoked    bool        `json:"revoked"`
	CreatedAt  *time.Time  `json:"created_at"`
	Scopes     []string    `json:"scopes"`
	UserID     int         `json:"user_id"`
	LastUsedAt interface{} `json:"last_used_at"`
	Active     bool        `json:"active"`
	ExpiresAt  string      `json:"expires_at"`
	Token      string      `json:"token"`
}

// AddServiceAccountsPATOptions represents the available AddServiceAccountsPAT() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/groups.html#create-personal-access-token-for-service-account-user
type AddServiceAccountsPATOptions struct {
	// Scopes cover the ranges of permission sets.
	// https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html#personal-access-token-scopes
	// e.g. api, read_user, read_api, read_repository, read_registry
	Scopes []string `json:"scopes,omitempty"`
	Name   string   `json:"name,omitempty"`
}

// AddServiceAccountsPAT add a new PAT for a service account user for a group.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/groups.html#add-group-hook
func (s *GroupsService) AddServiceAccountsPAT(gid interface{}, saID int, opt *AddServiceAccountsPATOptions, options ...RequestOptionFunc) (*GroupServiceAccountPAT, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/service_accounts/%d/personal_access_tokens", PathEscape(group), saID)

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pat := new(GroupServiceAccountPAT)
	resp, err := s.client.Do(req, pat)
	if err != nil {
		return nil, resp, err
	}

	return pat, resp, nil
}

// RotateServiceAccountsPAT rotate a PAT for a service account user for a group.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/groups.html#create-personal-access-token-for-service-account-user
func (s *GroupsService) RotateServiceAccountsPAT(gid interface{}, saID, tokenID int, options ...RequestOptionFunc) (*GroupServiceAccountPAT, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/service_accounts/%d/personal_access_tokens/%d/rotate", PathEscape(group), saID, tokenID)

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	pat := new(GroupServiceAccountPAT)
	resp, err := s.client.Do(req, pat)
	if err != nil {
		return nil, resp, err
	}

	return pat, resp, nil
}
