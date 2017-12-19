//
// Copyright 2017, Sander van Harmelen, Michael Lihs
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

// ProtectedBranchesService handles communication with the
// protected branch related methods of the GitLab API.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/protected_branches.html#protected-branches-api
type ProtectedBranchesService struct {
	client *Client
}

// BranchAccessDescription represents the access description for a
// protected branch
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/protected_branches.html#protected-branches-api
type BranchAccessDescription struct {
	AccessLevel            AccessLevelValue `json:"access_level"`
	AccessLevelDescription string           `json:"access_level_description"`
}

// ProtectedBranch represents a protected branch
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/protected_branches.html#list-protected-branches
type ProtectedBranch struct {
	Name              string                     `json:"name"`
	PushAccessLevels  []*BranchAccessDescription `json:"push_access_levels"`
	MergeAccessLevels []*BranchAccessDescription `json:"merge_access_levels"`
}

// ListProtectedBranches gets a list of protected branches from a project
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/protected_branches.html#list-protected-branches
func (s *ProtectedBranchesService) ListProtectedBranches(pid interface{}, options ...OptionFunc) ([]*ProtectedBranch, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/protected_branches", url.QueryEscape(project))

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var p []*ProtectedBranch
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

// GetProtectedBranch gets a single protected branch
// or wildcard protected branch
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/protected_branches.html#get-a-single-protected-branch-or-wildcard-protected-branch
func (s *ProtectedBranchesService) GetProtectedBranch(pid interface{}, name string, options ...OptionFunc) (*ProtectedBranch, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/protected_branches/%s", url.QueryEscape(project), name)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(ProtectedBranch)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}
