//
// Copyright 2021, Kordian Bruck
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

// PackagesService handles communication with the packages related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/packages.html
type PackagesService struct {
	client *Client
}

// Package represents a GitLab single package
//
// GitLab API docs: https://docs.gitlab.com/ee/api/packages.html
type Package struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Version   string     `json:"version"`
	Type      string     `json:"package_type"`
	CreatedAt *time.Time `json:"created_at"`
}

func (s Package) String() string {
	return Stringify(s)
}

// PackageFile represents one file contained within a package
//
// GitLab API docs: https://docs.gitlab.com/ee/api/packages.html
type PackageFile struct {
	ID        int        `json:"id"`
	PackageId int        `json:"package_id"`
	CreatedAt *time.Time `json:"created_at"`
	FileName  string     `json:"file_name"`
	Size      int        `json:"size"`
	MD5       string     `json:"file_md5"`
	SHA1      string     `json:"file_sha1"`
}

func (s PackageFile) String() string {
	return Stringify(s)
}

// ListProjectPackagesOptions are the parameters available in a ListProjectPackages() Operation
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/packages.html#within-a-project
type ListProjectPackagesOptions struct {
	ListOptions
	OrderBy            string `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort               string `url:"sort,omitempty" json:"sort,omitempty"`
	Type               string `url:"package_type,omitempty" json:"package_type,omitempty"`
	Name               string `url:"package_name,omitempty" json:"package_name,omitempty"`
	IncludeVersionless bool   `url:"include_versionless,omitempty" json:"include_versionless,omitempty"`
}

// ListProjectPackages gets a list of packages in a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/packages.html#within-a-project
func (s *PackagesService) ListProjectPackages(pid interface{}, opt *ListProjectPackagesOptions, options ...RequestOptionFunc) ([]*Package, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/packages", pathEscape(project))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var packages []*Package
	resp, err := s.client.Do(req, &packages)
	if err != nil {
		return nil, resp, err
	}

	return packages, resp, err
}

// DeleteRegistryRepository deletes a repository in a registry.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/packages.html#delete-a-project-package
func (s *PackagesService) DeleteProjectPackage(pid interface{}, packageID int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/packages/%d", pathEscape(project), packageID)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListPackageFilesOptions represents the available
// ListPackageFiles() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/packages.html#list-package-files
type ListPackageFilesOptions ListOptions

// ListPackageFiles gets a list of files that are within a package
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/packages.html#list-package-files
func (s *PackagesService) ListPackageFiles(pid interface{}, packageID int, opt *ListPackageFilesOptions, options ...RequestOptionFunc) ([]*PackageFile, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/packages/%d/package_files",
		pathEscape(project),
		packageID,
	)

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var tags []*PackageFile
	resp, err := s.client.Do(req, &tags)
	if err != nil {
		return nil, resp, err
	}

	return tags, resp, err
}
