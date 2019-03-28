//
// Copyright 2018, Sander van Harmelen
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

package gitlab

import (
	"fmt"
	"net/url"
	"time"
)

// DeploymentsService handles communication with the deployment related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/deployments.html
type DeploymentsService struct {
	client *Client
}

// Deployment represents the Gitlab deployment
type Deployment struct {
	ID          int          `json:"id" yaml:"id"`
	IID         int          `json:"iid" yaml:"iid"`
	Ref         string       `json:"ref" yaml:"ref"`
	SHA         string       `json:"sha" yaml:"sha"`
	CreatedAt   *time.Time   `json:"created_at" yaml:"created_at"`
	User        *ProjectUser `json:"user" yaml:"user"`
	Environment *Environment `json:"environment" yaml:"environment"`
	Deployable  struct {
		ID         int        `json:"id" yaml:"id"`
		Status     string     `json:"status" yaml:"status"`
		Stage      string     `json:"stage" yaml:"stage"`
		Name       string     `json:"name" yaml:"name"`
		Ref        string     `json:"ref" yaml:"ref"`
		Tag        bool       `json:"tag" yaml:"tag"`
		Coverage   float64    `json:"coverage" yaml:"coverage"`
		CreatedAt  *time.Time `json:"created_at" yaml:"created_at"`
		StartedAt  *time.Time `json:"started_at" yaml:"started_at"`
		FinishedAt *time.Time `json:"finished_at" yaml:"finished_at"`
		Duration   float64    `json:"duration" yaml:"duration"`
		User       *User      `json:"user" yaml:"user"`
		Commit     *Commit    `json:"commit" yaml:"commit"`
		Pipeline   struct {
			ID     int    `json:"id" yaml:"id"`
			SHA    string `json:"sha" yaml:"sha"`
			Ref    string `json:"ref" yaml:"ref"`
			Status string `json:"status" yaml:"status"`
		} `json:"pipeline" yaml:"pipeline"`
		Runner *Runner `json:"runner" yaml:"runner"`
	} `json:"deployable" yaml:"deployable"`
}

// ListProjectDeploymentsOptions represents the available ListProjectDeployments() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/deployments.html#list-project-deployments
type ListProjectDeploymentsOptions struct {
	ListOptions
	OrderBy *string `url:"order_by,omitempty" json:"order_by,omitempty" yaml:"order_by,omitempty"`
	Sort    *string `url:"sort,omitempty" json:"sort,omitempty" yaml:"sort,omitempty"`
}

// ListProjectDeployments gets a list of deployments in a project.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/deployments.html#list-project-deployments
func (s *DeploymentsService) ListProjectDeployments(pid interface{}, opts *ListProjectDeploymentsOptions, options ...OptionFunc) ([]*Deployment, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/deployments", url.QueryEscape(project))

	req, err := s.client.NewRequest("GET", u, opts, options)
	if err != nil {
		return nil, nil, err
	}

	var ds []*Deployment
	resp, err := s.client.Do(req, &ds)
	if err != nil {
		return nil, resp, err
	}

	return ds, resp, err
}

// GetProjectDeployment get a deployment for a project.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/deployments.html#get-a-specific-deployment
func (s *DeploymentsService) GetProjectDeployment(pid interface{}, deployment int, options ...OptionFunc) (*Deployment, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/deployments/%d", url.QueryEscape(project), deployment)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	d := new(Deployment)
	resp, err := s.client.Do(req, d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, err
}
