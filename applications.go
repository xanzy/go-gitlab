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

// ApplicationsService handles communication with administrables applications
// of the Gitlab API.

// Gitlab API docs : https://docs.gitlab.com/ee/api/applications.html
type ApplicationsService struct {
	client *Client
}

type ListApplicationsOptions struct {
}

type Application struct {
	ID              int    `json:"id"`
	ApplicationID   string `json:"application_id"`
	ApplicationName string `json:"application_name"`
	CallbackURL     string `json:"callback_url"`
	Confidential    bool   `json:"confidential"`
}

// ListApplications get a list of administrables applications by the authenticated user
func (s *ApplicationsService) ListApplications(opts *ListApplicationsOptions, options ...OptionFunc) ([]*Application, *Response, error) {

	req, err := s.client.NewRequest("GET", "applications", opts, options)
	if err != nil {
		return nil, nil, err
	}

	var a []*Application
	resp, err := s.client.Do(req, &a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, err
}
