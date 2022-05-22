//
// Copyright 2022, Ryan Glab <ryan.j.glab@gmail.com>
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

// ErrorTrackingService handles communication with the error tracking
// methods of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/error_tracking.html
type ErrorTrackingService struct {
	client *Client
}

// ErrorTracking represents error tracking settings for a GitLab project.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/error_tracking.html
type ErrorTracking struct {
	Active            bool   `json:"active"`
	ProjectName       string `json:"project_name"`
	SentryExternalURL string `json:"sentry_external_url"`
	APIURL            string `json:"api_url"`
	Integrated        bool   `json:"integrated"`
}

func (p ErrorTracking) String() string {
	return Stringify(p)
}

// GetErrorTracking gets Sentry error tracking settings for a specific project
//
// GitLab API docs: https://docs.gitlab.com/ee/api/error_tracking.html#get-error-tracking-settings
func (s *ErrorTrackingService) GetErrorTracking(pid interface{}, options ...RequestOptionFunc) (*ErrorTracking, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/error_tracking/settings", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	ErrorTracking := new(ErrorTracking)
	resp, err := s.client.Do(req, ErrorTracking)
	if err != nil {
		return nil, resp, err
	}

	return ErrorTracking, resp, err
}

type ConfigureErrorTrackingOptions struct {
	Active     *bool `url:"active,omitempty" json:"active,omitempty"`
	Integrated *bool `url:"integrated,omitempty" json:"integrated,omitempty"`
}

// EnableDisableErrorTracking allows the enabling or disabling of sentry integration on a specified project
//
// GitLab API docs: https://docs.gitlab.com/ee/api/error_tracking.html#enable-or-disable-the-error-tracking-project-settings
func (s *ErrorTrackingService) EnableDisableErrorTracking(pid interface{}, opt *ConfigureErrorTrackingOptions, options ...RequestOptionFunc) (*ErrorTracking, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%d/error_tracking/settings", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPatch, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	ET := new(ErrorTracking)
	resp, err := s.client.Do(req, &ET)
	if err != nil {
		return nil, resp, err
	}

	return ET, resp, err
}

// ListErrorTrackingClientKeysOptions represents the available
// ListErrorTrackingClientKeys() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/error_tracking.html#list-project-client-keys
type ListErrorTrackingClientKeysOptions ListOptions

type ProjectClientKey struct {
	ID        int    `json:"id"`
	Active    bool   `json:"active"`
	PublicKey string `json:"public_key"`
	SentryDsn string `json:"sentry_dsn"`
}

// ListErrorTrackingClientKeys gets a list of project client keys for sentry.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/error_tracking.html#list-project-client-keys
func (s *ErrorTrackingService) ListErrorTrackingClientKeys(pid interface{}, opt *ListErrorTrackingClientKeysOptions, options ...RequestOptionFunc) ([]*ProjectClientKey, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/error_tracking/client_keys", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var projectClientKeys []*ProjectClientKey
	resp, err := s.client.Do(req, &projectClientKeys)
	if err != nil {
		return nil, resp, err
	}

	return projectClientKeys, resp, err
}

// CreateClientKey creates a project client key for sentry integration in a GitLab project.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/error_tracking.html#create-a-client-key
func (s *ErrorTrackingService) CreateClientKey(pid interface{}, options ...RequestOptionFunc) (*ProjectClientKey, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/error_tracking/client_keys", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	projectClientKey := new(ProjectClientKey)
	resp, err := s.client.Do(req, projectClientKey)
	if err != nil {
		return nil, resp, err
	}

	return projectClientKey, resp, err
}

// DeleteClientKey removes a sentry client key from a project.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/error_tracking.html#delete-a-client-key
func (s *ErrorTrackingService) DeleteClientKey(pid interface{}, keyID int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("projects/%s/error_tracking/client_keys/%d", PathEscape(project), keyID)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
