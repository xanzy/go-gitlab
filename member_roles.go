package gitlab

import (
	"fmt"
	"net/http"
)

// MemberRolesService handles communication with the member roles related methods of
// the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/member_roles.html
type MemberRolesService struct {
	client *Client
}

// MemberRole represents a GitLab member role.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/member_roles.html
type MemberRole struct {
	ID                       int              `json:"id"`
	Name                     string           `json:"name"`
	Description              string           `json:"description,omitempty"`
	GroupId                  int              `json:"group_id"`
	BaseAccessLevel          AccessLevelValue `json:"base_access_level"`
	AdminMergeRequests       bool             `json:"admin_merge_requests,omitempty"`
	AdminVulnerability       bool             `json:"admin_vulnerability,omitempty"`
	ReadCode                 bool             `json:"read_code,omitempty"`
	ReadDependency           bool             `json:"read_dependency,omitempty"`
	ReadVulnerability        bool             `json:"read_vulnerability,omitempty"`
	ManageProjectAccessToken bool             `json:"manage_project_access_token,omitempty"`
}

// CreateMemberRoleOptions represents the available CreateMemberRole() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/member_roles.html#add-a-member-role-to-a-group
type CreateMemberRoleOptions struct {
	Name               string           `json:"name,"`
	BaseAccessLevel    AccessLevelValue `json:"base_access_level"`
	Description        string           `json:"description,omitempty"`
	AdminMergeRequest  bool             `json:"admin_merge_request,omitempty"`
	AdminVulnerability bool             `json:"admin_vulnerability,omitempty"`
	ReadCode           bool             `json:"read_code,omitempty"`
	ReadDependency     bool             `json:"read_dependency,omitempty"`
	ReadVulnerability  bool             `json:"read_vulnerability,omitempty"`
}

// ListMemberRoles gets a list of member roles for a specified group.
//
// Gitlab API docs: https://docs.gitlab.com/ee/api/member_roles.html#list-all-member-roles-of-a-group
func (s *MemberRolesService) ListMemberRoles(groupId int, options ...RequestOptionFunc) ([]*MemberRole, *Response, error) {
	path := fmt.Sprintf("groups/%d/member_roles", groupId)
	req, err := s.client.NewRequest(http.MethodGet, path, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var memberRoles []*MemberRole
	resp, err := s.client.Do(req, &memberRoles)
	if err != nil {
		return nil, resp, err
	}

	return memberRoles, resp, nil
}

// CreateMemberRole creates a new member role for a specified group.
//
// Gitlab API docs: https://docs.gitlab.com/ee/api/member_roles.html#add-a-member-role-to-a-group
func (s *MemberRolesService) CreateMemberRole(groupId int, opt *CreateMemberRoleOptions, options ...RequestOptionFunc) (*MemberRole, *Response, error) {
	path := fmt.Sprintf("groups/%d/member_roles", groupId)
	req, err := s.client.NewRequest(http.MethodPost, path, opt, options)
	if err != nil {
		return nil, nil, err
	}

	memberRole := new(MemberRole)
	resp, err := s.client.Do(req, memberRole)
	if err != nil {
		return nil, resp, err
	}

	return memberRole, resp, nil
}

// DeleteMemberRole deletes a member role from a specified group.
//
// Gitlab API docs: https://docs.gitlab.com/ee/api/member_roles.html#remove-member-role-of-a-group
func (s *MemberRolesService) DeleteMemberRole(groupId, memberRoleId int, options ...RequestOptionFunc) (*Response, error) {
	path := fmt.Sprintf("groups/%d/member_roles/%d", groupId, memberRoleId)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
