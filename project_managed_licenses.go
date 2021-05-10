//
// Copyright 2021, Andrea Perizzato
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

// ProjectManagedLicensesService handles communication with the managed
// licenses methods of the GitLab API.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html
type ProjectManagedLicensesService struct {
	client *Client
}

// ManagedLicenseApprovalStatus describe the approval statuses of a license.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html
type ManagedLicenseApprovalStatus string

// Values of ManagedLicenseApprovalStatus.
const (
	ManagedLicenseApproved    ManagedLicenseApprovalStatus = "approved"
	ManagedLicenseBlacklisted ManagedLicenseApprovalStatus = "blacklisted"
)

// ManagedLicense represents a managed license.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html
type ManagedLicense struct {
	ID             int                          `json:"id"`
	Name           string                       `json:"name"`
	ApprovalStatus ManagedLicenseApprovalStatus `json:"approval_status"`
}

// ListManagedLicenses returns a list of managed licenses from a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html#list-managed-licenses
func (s *ProjectManagedLicensesService) ListManagedLicenses(pid interface{}, options ...RequestOptionFunc) ([]*ManagedLicense, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/managed_licenses", pathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var mls []*ManagedLicense
	resp, err := s.client.Do(req, &mls)
	if err != nil {
		return nil, resp, err
	}

	return mls, resp, err
}

// GetManagedLicense returns a single managed license for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html#show-an-existing-managed-license
func (s *ProjectManagedLicensesService) GetManagedLicense(pid, lid interface{}, options ...RequestOptionFunc) (*ManagedLicense, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	license, err := parseID(lid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/managed_licenses/%s", pathEscape(project), pathEscape(license))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	ml := new(ManagedLicense)
	resp, err := s.client.Do(req, ml)
	if err != nil {
		return nil, resp, err
	}

	return ml, resp, err
}

// AddManagedLicenseOptions represents the available options to add a managed license.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html#create-a-new-managed-license
type AddManagedLicenseOptions struct {
	Name           string                       `json:"name"`
	ApprovalStatus ManagedLicenseApprovalStatus `json:"approval_status"`
}

// AddManagedLicense adds a managed license to a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html#create-a-new-managed-license
func (s *ProjectManagedLicensesService) AddManagedLicense(pid interface{}, opt *AddManagedLicenseOptions, options ...RequestOptionFunc) (*ManagedLicense, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/managed_licenses", pathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	ml := new(ManagedLicense)
	resp, err := s.client.Do(req, ml)
	if err != nil {
		return nil, resp, err
	}

	return ml, resp, err
}

// RemoveManagedLicense removes a managed license from a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html#delete-a-managed-license
func (s *ProjectManagedLicensesService) RemoveManagedLicense(pid, lid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	license, err := parseID(lid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/managed_licenses/%s", pathEscape(project), pathEscape(license))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// EditManagedLicenceOptions represents the available options to edit a managed license.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html#edit-an-existing-managed-license
type EditManagedLicenceOptions struct {
	ApprovalStatus ManagedLicenseApprovalStatus `url:"approval_status"`
}

// EditManagedLicense updates an existing managed license in a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html#edit-an-existing-managed-license
func (s *ProjectManagedLicensesService) EditManagedLicense(pid, lid interface{}, opt *EditManagedLicenceOptions, options ...RequestOptionFunc) (*ManagedLicense, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	license, err := parseID(lid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/managed_licenses/%s", pathEscape(project), pathEscape(license))

	req, err := s.client.NewRequest(http.MethodPatch, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	ml := new(ManagedLicense)
	resp, err := s.client.Do(req, ml)
	if err != nil {
		return nil, resp, err
	}

	return ml, resp, err
}
