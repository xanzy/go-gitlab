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
)

const invitationsProjectPath = "projects/%s/invitations"

// ProjectInvitesService handles communication with the group related methods of
// the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/groups.html
type ProjectInvitesService struct {
    client *Client
}

// ListPendingInvitations gets a list of pending invitations for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/invitations.html#list-all-invitations-pending-for-a-group-or-project
func (s *ProjectInvitesService) ListPendingInvitations(gid interface{}, opt *ListPendingInvitationsOptions, options ...RequestOptionFunc) ([]*PendingInvitations, *Response, error) {
    group, err := parseID(gid)
    if err != nil {
        return nil, nil, err
    }
    u := fmt.Sprintf(invitationsProjectPath, pathEscape(group))
    req, err := s.client.NewRequest("GET", u, opt, options)
    if err != nil {
        return nil, nil, err
    }

    var pendingInvitations []*PendingInvitations
    resp, err := s.client.Do(req, &pendingInvitations)
    if err != nil {
        return nil, resp, err
    }

    return pendingInvitations, resp, err
}

// ProjectInvites Send Invites to a project
//
// GitLab API docs: https://docs.gitlab.com/ee/api/invitations.html#list-all-invitations-pending-for-a-group-or-project
func (s *ProjectInvitesService) ProjectInvites(gid interface{}, opt InvitesOptions, options ...RequestOptionFunc) (*InvitationsResponse, *Response, error) {
    group, err := parseID(gid)
    if err != nil {
        return nil, nil, err
    }
    u := fmt.Sprintf(invitationsProjectPath, pathEscape(group))

    req, err := s.client.NewRequest("POST", u, opt, options)
    if err != nil {
        return nil, nil, err
    }

    g := new(InvitationsResponse)
    resp, err := s.client.Do(req, g)
    if err != nil {
        return nil, resp, err
    }

    return g, resp, err
}
