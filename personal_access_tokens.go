//
// Copyright 2022, Ryan Glab <ryan.j.glab@gmail.com>
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
)

// PersonalAccessTokensService handles communication with the Personal Access Tokens related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/personal_access_tokens.html
type PersonalAccessTokensService struct {
	client *Client
}

func (p PersonalAccessToken) String() string {
	return Stringify(p)
}

// ListPersonalAccessTokensOptions represents the available options for
// listing Personal Access Tokens across a GitLab instance or for a specific user.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/personal_access_tokens.html#list-personal-access-tokens
type ListPersonalAccessTokensOptions struct {
	ListOptions
	UserID *int `url:"user_id,omitempty" json:"user_id,omitempty"`
}

// ListPersonalAccessTokens gets a list of all Personal Access Tokens in a
// GitLab instance optionally scoped down to a single users tokens.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/personal_access_tokens.html#list-personal-access-tokens
func (s *PersonalAccessTokensService) ListPersonalAccessTokens(opt *ListPersonalAccessTokensOptions, options ...RequestOptionFunc) ([]*PersonalAccessToken, *Response, error) {
	u := "personal_access_tokens"

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var pats []*PersonalAccessToken
	resp, err := s.client.Do(req, &pats)
	if err != nil {
		return nil, resp, err
	}

	return pats, resp, err
}

// DeletePersonalAccessToken deletes a Personal Access Token.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/personal_access_tokens.html#revoke-a-personal-access-token
func (s *PersonalAccessTokensService) DeletePersonalAccessToken(id int, options ...RequestOptionFunc) (*Response, error) {

	u := fmt.Sprintf("personal_access_tokens/%d", id)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
