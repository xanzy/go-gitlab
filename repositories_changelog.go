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
	"time"
)

// ChangelogOptions represents a Changelog file.
//
// https://docs.gitlab.com/ee/api/repositories.html#add-changelog-data-to-a-changelog-file
type ChangelogOptions struct {
	// The version to generate the changelog for. [The format must follow semantic versioning.](https://semver.org/)
	Version string `json:"version" url:"version"`
	// The branch to commit the changelog to. Defaults to the project's default branch.
	Branch string `json:"branch,omitempty" url:"-"`
	// Path to the changelog configuration file in the project's Git repository.
	// Defaults to .gitlab/changelog_config.yml
	ConfigFile string `json:"config_file,omitempty" url:"config_file,omitempty"`
	// The date and time of the release. Defaults to the current time.
	Date time.Time `json:"date,omitempty" url:"date,omitempty"`
	// The file to commit the changes to. Defaults to CHANGELOG.md.
	File string `json:"file,omitempty" url:"-"`
	// the SHA of the commit that marks the beginning range of commits to
	// include in the changelog. The commit isn't included in the changelog.
	From string `json:"from,omitempty" url:"from,omitempty"`
	// The commit message to use when committing the changes. Defaults to "Add
	// changelog for Version X", where X is the value of the version argument.
	Message string `json:"message,omitempty" url:"-"`
	// The SHA of the commit that marks the end of the range of commits to
	// include in the changelog. This commit IS included in the changelog.
	// Defaults to the branch specified in the branch attribute. Limited to
	// 15000 commits unless the feature flag changelog_commits_limitation is
	// disabled.
	To string `json:"to,omitempty" url:"to,omitempty"`
	// The Git trailer to use for including commits. Defaults to "Changelog".
	// Case-sensitive: "Example" does not match "example" or "eXaMpLE".
	Trailer string `json:"trailer,omitempty" url:"trailer,omitempty"`
}

// AddChangelog genertates changelog data based on commits in a repository.
//
// Gitlab API docs:
// https://docs.gitlab.com/ee/api/repositories.html#add-changelog-data-to-a-changelog-file
func (s *RepositoriesService) AddChangelog(pid interface{}, opt *ChangelogOptions, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/changelog", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, nil)
}

// The generated Changelog data
type GeneratedChangelogNotes struct {
	Notes string `json:"notes"`
}

// GenerateChangelog works almost exactly like AddChangelog, except that the
// the file isn't committed to the repository.
//
// Gitlab API docs:
// https://docs.gitlab.com/ee/api/repositories.html#generate-changelog-data
func (s *RepositoriesService) GenerateChangelog(pid interface{}, opt ChangelogOptions, options ...RequestOptionFunc) (*GeneratedChangelogNotes, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/changelog", project)

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	cl := new(GeneratedChangelogNotes)
	resp, err := s.client.Do(req, cl)
	if err != nil {
		return nil, resp, err
	}
	return cl, resp, err
}
