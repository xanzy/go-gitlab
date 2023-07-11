// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package gitlab

import (
	"fmt"
	"net/http"
)

// JobTokenScopeService handles communication with project CI settings
// such as token permissions.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_job_token_scopes.html
type JobTokenScopeService struct {
	client *Client
}

// GetJobTokenInboundAllowOptions represents the available options
// when querying the inbound CI allow-list for projects
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_job_token_scopes.html#get-a-projects-cicd-job-token-inbound-allowlist
type GetJobTokenInboundAllowOptions struct {
	ListOptions
}

// Fetch the CI/CD job token inbound allowlist (job token scope) of a project.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_job_token_scopes.html#get-a-projects-cicd-job-token-inbound-allowlist
func (j *JobTokenScopeService) GetProjectJobTokenInboundAllowlist(pid interface{}, opt *GetJobTokenInboundAllowOptions, options ...RequestOptionFunc) ([]*Project, *Response, error) {

	// Parse the project Id or namespace and create the URL
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf(`projects/%s/job_token_scope/allowlist`, PathEscape(project))

	req, err := j.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var p []*Project
	resp, err := j.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// JobTokenInboundAllowOptions represents the available options
// when adding or removing a project to the CI/CD job token inbound allowlist of a project.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_job_token_scopes.html#create-a-new-project-to-a-projects-cicd-job-token-inbound-allowlist
type JobTokenInboundAllowOptions struct {
	TargetProjectID int `json:"target_project_id"`
}

// AddJobTokenInboundAllowResponse represents the response from the
// Create a new project to a project’s CI/CD job token inbound allowlist API request
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_job_token_scopes.html#create-a-new-project-to-a-projects-cicd-job-token-inbound-allowlist
type AddJobTokenInboundAllowResponse struct {
	SourceProjectID int `json:"source_project_id"`
	TargetProjectID int `json:"target_project_id"`
}

// Adds a new Project to a Project's Job Token Inbound Allow list
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_job_token_scopes.html#create-a-new-project-to-a-projects-cicd-job-token-inbound-allowlist
func (j *JobTokenScopeService) AddProjectToJobScopeAllowList(pid interface{}, opt *JobTokenInboundAllowOptions, options ...RequestOptionFunc) (*AddJobTokenInboundAllowResponse, *Response, error) {
	// Parse the project Id or namespace and create the URL
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf(`projects/%s/job_token_scope/allowlist`, PathEscape(project))
	req, err := j.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	a := new(AddJobTokenInboundAllowResponse)
	resp, err := j.client.Do(req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}

// Removes a Project from a Project's Job Token Inbound Allow list
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_job_token_scopes.html#remove-a-project-from-a-projects-cicd-job-token-inbound-allowlist
func (j *JobTokenScopeService) RemoveProjectFromJobScopeAllowList(pid interface{}, opt *JobTokenInboundAllowOptions, options ...RequestOptionFunc) (*Response, error) {
	// Parse the project Id or namespace and create the URL
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf(`projects/%s/job_token_scope/allowlist/%d`, PathEscape(project), opt.TargetProjectID)
	req, err := j.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	// This API has no body in the request or response
	resp, err := j.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
