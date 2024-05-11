//
// Copyright 2021, Sander van Harmelen
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
	"net/http"
)

// ImportService handles communication with the import
// related methods of the GitLab API.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/import.html
type ImportService struct {
	client *Client
}

// GithubImport represents the response from an import from Github.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/import.html#import-repository-from-github
type GithubImport struct {
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	FullPath              string `json:"full_path"`
	FullName              string `json:"full_name"`
	RefsUrl               string `json:"refs_url"`
	ImportSource          string `json:"import_source"`
	ImportStatus          string `json:"import_status"`
	HumanImportStatusName string `json:"human_import_status_name"`
	ProviderLink          string `json:"provider_link"`
	RelationType          string `json:"relation_type"`
	ImportWarning         string `json:"import_warning"`
}

func (s GithubImport) String() string {
	return Stringify(s)
}

// ImportRepositoryFromGithubOptions represents the available
// ImportRepositoryFromGithub() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/import.html#import-repository-from-github
type ImportRepositoryFromGithubOptions struct {
	PersonalAccessToken *string `url:"personal_access_token,omitempty" json:"personal_access_token,omitempty"`
	RepoID              *int    `url:"repo_id,omitempty" json:"repo_id,omitempty"`
	NewName             *string `url:"new_name,omitempty" json:"new_name,omitempty"`
	TargetNamespace     *string `url:"target_namespace,omitempty" json:"target_namespace,omitempty"`
	GithubHostname      *string `url:"github_hostname,omitempty" json:"github_hostname,omitempty"`
	OptionalStages      struct {
		SingleEndpointNotesImport *bool `url:"single_endpoint_notes_import,omitempty" json:"single_endpoint_notes_import,omitempty"`
		AttachmentsImport         *bool `url:"attachments_import,omitempty" json:"attachments_import,omitempty"`
		CollaboratorsImport       *bool `url:"collaborators_import,omitempty" json:"collaborators_import,omitempty"`
	} `url:"optional_stages,omitempty" json:"optional_stages,omitempty"`
	TimeoutStrategy *string `url:"timeout_strategy,omitempty" json:"timeout_strategy,omitempty"`
}

// Import a repository from Github.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/import.html#import-repository-from-github
func (s *ImportService) ImportRepositoryFromGithub(opt *ImportRepositoryFromGithubOptions, options ...RequestOptionFunc) (*GithubImport, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "import/github", opt, options)
	if err != nil {
		return nil, nil, err
	}

	gi := new(GithubImport)
	resp, err := s.client.Do(req, gi)
	if err != nil {
		return nil, resp, err
	}

	return gi, resp, nil
}
