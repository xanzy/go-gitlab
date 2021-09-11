//
// Copyright 2021, Igor Varavko
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

import "net/http"

// PlanLimitsService handles communication with the repositories related
// methods of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/plan_limits.html
type PlanLimitsService struct {
	client *Client
}

// PlanLimit represents a GitLab pipeline.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/plan_limits.html
type PlanLimit struct {
	PlanName                   *string `json:"plan_name,omitempty"`
	ConanMaxFileSize           *int    `json:"conan_max_file_size,omitempty"`
	GenericPackagesMaxFileSize *int    `json:"generic_packages_max_file_size,omitempty"`
	HelmMaxFileSize            *int    `json:"helm_max_file_size,omitempty"`
	MavenMaxFileSize           *int    `json:"maven_max_file_size,omitempty"`
	NpmMaxFileSize             *int    `json:"npm_max_file_size,omitempty"`
	NugetMaxFileSize           *int    `json:"nuget_max_file_size,omitempty"`
	PyPiMaxFileSize            *int    `json:"pypi_max_file_size,omitempty"`
	TerraformModuleMaxFileSize *int    `json:"terraform_module_max_file_size,omitempty"`
}

// PlanLimitOptions represents the available options that can be passed
// to the API when updating the Plan Limits.
type PlanLimitOptions struct {
	PlanName                   string `json:"plan_name,omitempty" url:"plan_name,omitempty"`
	ConanMaxFileSize           *int   `json:"conan_max_file_size,omitempty" url:"conan_max_file_size,omitempty"`
	GenericPackagesMaxFileSize *int   `json:"generic_packages_max_file_size,omitempty" url:"generic_packages_max_file_size,omitempty"`
	HelmMaxFileSize            *int   `json:"helm_max_file_size,omitempty" url:"helm_max_file_size,omitempty"`
	MavenMaxFileSize           *int   `json:"maven_max_file_size,omitempty" url:"maven_max_file_size,omitempty"`
	NpmMaxFileSize             *int   `json:"npm_max_file_size,omitempty" url:"npm_max_file_size,omitempty"`
	NugetMaxFileSize           *int   `json:"nuget_max_file_size,omitempty" url:"nuget_max_file_size,omitempty"`
	PyPiMaxFileSize            *int   `json:"pypi_max_file_size,omitempty" url:"pypi_max_file_size,omitempty"`
	TerraformModuleMaxFileSize *int   `json:"terraform_module_max_file_size,omitempty" url:"terraform_module_max_file_size,omitempty"`
}

// List the current limits of a plan on the GitLab instance.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/plan_limits.html#get-current-plan-limits
func (s *PlanLimitsService) GetCurrentPlanLimits(options ...RequestOptionFunc) (*PlanLimit, *Response, error) {
	u := "application/plan_limits"

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var planlimit *PlanLimit
	resp, err := s.client.Do(req, &planlimit)
	if err != nil {
		return nil, resp, err
	}

	return planlimit, resp, err
}

// Modify the limits of a plan on the GitLab instance.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/plan_limits.html#change-plan-limits
func (s *PlanLimitsService) ChangePlanLimits(opt *PlanLimitOptions, options ...RequestOptionFunc) (*PlanLimit, *Response, error) {
	u := "application/plan_limits"

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var planlimit *PlanLimit
	resp, err := s.client.Do(req, &planlimit)
	if err != nil {
		return nil, resp, err
	}

	return planlimit, resp, err
}
