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

// PipelinesEmailService represents Pipelines Email service settings.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/services.html#pipeline-emails
type PipelinesEmailService struct {
	Service
	Properties *PipelinesEmailProperties `json:"properties"`
}

// PipelinesEmailProperties represents PipelinesEmail specific properties.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/services.html#pipeline-emails
type PipelinesEmailProperties struct {
	Recipients                string    `json:"recipients,omitempty"`
	NotifyOnlyBrokenPipelines BoolValue `json:"notify_only_broken_pipelines,omitempty"`
	NotifyOnlyDefaultBranch   BoolValue `json:"notify_only_default_branch,omitempty"`
}

// SetPipelinesEmailServiceOptions represents the available SetPipelinesEmailService()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/services.html#pipeline-emails
type SetPipelinesEmailServiceOptions struct {
	Recipients                *string `url:"recipients,omitempty" json:"recipients,omitempty"`
	NotifyOnlyBrokenPipelines *bool   `url:"notify_only_broken_pipelines,omitempty" json:"notify_only_broken_pipelines,omitempty"`
	NotifyOnlyDefaultBranch   *bool   `url:"notify_only_default_branch,omitempty" json:"notify_only_default_branch,omitempty"`
	AddPusher                 *bool   `url:"add_pusher,omitempty" json:"add_pusher,omitempty"`
	BranchesToBeNotified      *string `url:"branches_to_be_notified,omitempty" json:"branches_to_be_notified,omitempty"`
	PipelineEvents            *bool   `url:"pipeline_events,omitempty" json:"pipeline_events,omitempty"`
}

// GetPipelinesEmailService gets Pipelines Email service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/services.html#get-pipeline-emails-service-settings
func (s *ServicesService) GetPipelinesEmailService(pid interface{}, options ...OptionFunc) (*PipelinesEmailService, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/services/pipelines-email", pathEscape(project))

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	svc := new(PipelinesEmailService)
	resp, err := s.client.Do(req, svc)
	if err != nil {
		return nil, resp, err
	}

	return svc, resp, err
}

// SetPipelinesEmailService sets Pipelines Email service for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/services.html#pipeline-emails
func (s *ServicesService) SetPipelinesEmailService(pid interface{}, opt *SetPipelinesEmailServiceOptions, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/pipelines-email", pathEscape(project))

	req, err := s.client.NewRequest("PUT", u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DeletePipelinesEmailService deletes Pipelines Email service settings for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/services.html#delete-pipeline-emails-service
func (s *ServicesService) DeletePipelinesEmailService(pid interface{}, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/services/pipelines-email", pathEscape(project))

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
