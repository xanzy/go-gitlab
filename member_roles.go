package gitlab

import (
	"fmt"
	"net/http"
)

const (
	GuestBaseLevel BaseAccessLevel = 10
)

type BaseAccessLevel int

type MemberRolesService struct {
	client *Client
}

type MemberRole struct {
	Id                       int    `json:"id"`
	Name                     string `json:"name"`
	Description              string `json:"description,omitempty"`
	GroupId                  int    `json:"group_id"`
	BaseAccessLevel          int    `json:"base_access_level"`
	AdminMergeRequests       bool   `json:"admin_merge_requests,omitempty"`
	AdminVulnerability       bool   `json:"admin_vulnerability,omitempty"`
	ReadCode                 bool   `json:"read_code,omitempty"`
	ReadDependency           bool   `json:"read_dependency,omitempty"`
	ReadVulnerability        bool   `json:"read_vulnerability,omitempty"`
	ManageProjectAccessToken bool   `json:"manage_project_access_token,omitempty"`
}

type CreateMemberRoleOptions struct {
	Name               string          `json:"name,"`
	BaseAccessLevel    BaseAccessLevel `json:"base_access_level"`
	Description        string          `json:"description,omitempty"`
	AdminMergeRequest  bool            `json:"admin_merge_request,omitempty"`
	AdminVulnerability bool            `json:"admin_vulnerability,omitempty"`
	ReadCode           bool            `json:"read_code,omitempty"`
	ReadDependency     bool            `json:"read_dependency,omitempty"`
	ReadVulnerability  bool            `json:"read_vulnerability,omitempty"`
}

func (s *MemberRolesService) ListMemberRoles(groupId int, options ...RequestOptionFunc) ([]*MemberRole, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("/groups/%d/member_roles", groupId), nil, options)
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

func (s *MemberRolesService) CreateMemberRole(groupId int, opt *CreateMemberRoleOptions, options ...RequestOptionFunc) (*MemberRole, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("/groups/%d/member_roles", groupId), opt, options)
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

func (s *MemberRolesService) DeleteMemberRole(groupId, memberRoleId int, options ...RequestOptionFunc) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("/groups/%d/member_roles/%d", groupId, memberRoleId), nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
