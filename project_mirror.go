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
	"time"
)

// ProjectMirrorService handles communication with the project mirror
// related methods of the GitLab API.
type ProjectMirrorService struct {
	client *Client
}

// ProjectMirror represents a project mirror configuration.
//
// GitLAb API docs:
// https://docs.gitlab.com/ce/api/remote_mirrors.html
type ProjectMirror struct {
	Enabled                bool        `json:"enabled"`
	ID                     int         `json:"id"`
	LastError              interface{} `json:"last_error"`
	LastSuccessfulUpdateAt time.Time   `json:"last_successful_update_at"`
	LastUpdateAt           time.Time   `json:"last_update_at"`
	LastUpdateStartedAt    time.Time   `json:"last_update_started_at"`
	OnlyProtectedBranches  bool        `json:"only_protected_branches"`
	KeepDivergentRefs      bool        `json:"keep_divergent_refs"`
	UpdateStatus           string      `json:"update_status"`
	URL                    string      `json:"url"`
}

// PostProjectMirrorOptions contains the properties requires to create
// a new project mirror.
type PostProjectMirrorOptions struct {
	URL                   string `json:"url"`
	Enabled               bool   `json:"enabled"`
	OnlyProtectedBranches bool   `json:"only_protected_branches"`
	KeepDivergentRefs     bool   `json:"keep_divergent_refs"`
}

// ListProjectMirror gets a list of mirrors configured on the project.
// these are copies of the repositories to external version control repositories
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/remote_mirrors.html
func (s *ProjectMirrorService) ListProjectMirror(pid interface{}) ([]*ProjectMirror, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/remote_mirrors", pathEscape(project))

	req, err := s.client.NewRequest("GET", u, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var pm []*ProjectMirror

	resp, err := s.client.Do(req, &pm)
	if err != nil {
		return nil, resp, err
	}

	return pm, resp, err

}
