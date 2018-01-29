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
)

// SubgroupsService handles communication with the subgroup related methods of
// the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/groups.html
type SubgroupsService struct {
	client *Client
}

//Subgroup represents a GitLab subgroup.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/groups.html
type Subgroup struct {
	ID                   int              `json:"id"`
	Name                 string           `json:"name"`
	Path                 string           `json:"path"`
	Description          string           `json:"description"`
	Visibility           *VisibilityValue `json:"visibility"`
	LFSEnabled           bool             `json:"lfs_enabled"`
	AvatarURL            string           `json:"avatar_url"`
	WebURL               string           `json:"web_url"`
	RequestAccessEnabled bool             `json:"request_access_enabled"`
	FullName             string           `json:"full_name"`
	FullPath             string           `json:"full_path"`
	ParentID             int              `json:"parent_id"`
}


// ListGroupsOptions represents the available ListGroups() options.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/groups.html#list-project-groups
type ListSubgroupsOptions struct {
	ListOptions
	AllAvailable *bool   `url:"all_available,omitempty" json:"all_available,omitempty"`
	OrderBy      *string `url:"order_by,omitempty" json:"order_by,omitempty"`
	Owned        *bool   `url:"owned,omitempty" json:"owned,omitempty"`
	Search       *string `url:"search,omitempty" json:"search,omitempty"`
	Sort         *string `url:"sort,omitempty" json:"sort,omitempty"`
	Statistics   *bool   `url:"statistics,omitempty" json:"statistics,omitempty"`
}

// ListSubgroups gets a list of subgroups for a given project.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/groups.html#list-project-groups
func (s *GroupsService) ListSubgroups(gid interface{}, opt *ListSubgroupsOptions, options ...OptionFunc) ([]*Subgroup, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/subgroups", url.QueryEscape(group))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var sg []*Subgroup
	resp, err := s.client.Do(req, &sg)
	if err != nil {
		return nil, resp, err
	}

	return sg, resp, err
}