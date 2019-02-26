package gitlab

import (
	"fmt"
	"net/url"
	"time"
)

// RegistryService handles communication with the container registry related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/container_registry.html
type RegistryService struct {
	client *Client
}

// RegistryRepository represents a GitLab content registry repository.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/container_registry.html
type RegistryRepository struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	Location  string     `json:"location"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

func (s RegistryRepository) String() string {
	return Stringify(s)
}

// RegistryRepositoryTag represents a GitLab registry image tag.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/container_registry.html
type RegistryRepositoryTag struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Location string `json:"location"`
}

func (s RegistryRepositoryTag) String() string {
	return Stringify(s)
}

// RegistryRepositoryTagDetail represents a GitLab registry tag detail.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/container_registry.html
type RegistryRepositoryTagDetail struct {
	Name          string     `json:"name"`
	Path          string     `json:"path"`
	Location      string     `json:"location"`
	Revision      string     `json:"revision"`
	ShortRevision string     `json:"short_revision"`
	Digest        string     `json:"digest"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	TotalSize     int        `json:"total_size"`
}

func (s RegistryRepositoryTagDetail) String() string {
	return Stringify(s)
}

// ListRegistryRepositoriesOptions represents the available ListRegistryRepositories() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/container_registry.html#list-registry-repositories
type ListRegistryRepositoriesOptions ListOptions

// ListRegistryRepositoryTagsOptions represents the available ListRegistryRepositoryTags() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/container_registry.html#list-repository-tags
type ListRegistryRepositoryTagsOptions ListOptions

// BulkDeleteRegistryRepositoryTagsOptions represents the available BulkDeleteRegistryRepositoryTags() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/container_registry.html#delete-repository-tags-in-bulk
type BulkDeleteRegistryRepositoryTagsOptions struct {
	NameRegexp *string `url:"name_regex,omitempty" json:"name_regex,omitempty"`
	KeepN      *int    `url:"keep_n,omitempty" json:"keep_n,omitempty"`
	OlderThan  *string `url:"older_than,omitempty" json:"older_than,omitempty"`
}

// ListRegistryRepositories gets a list of registry repositories in a projecty
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/container_registry.html#list-registry-repositories
func (s *RegistryService) ListRegistryRepositories(pid interface{}, opt *ListRegistryRepositoriesOptions, options ...OptionFunc) ([]*RegistryRepository, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/registry/repositories", url.QueryEscape(project))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var repos []*RegistryRepository
	resp, err := s.client.Do(req, &repos)
	if err != nil {
		return nil, resp, err
	}

	return repos, resp, err
}

// DeleteRegistryRepository deletes a repository in a registry.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/container_registry.html#delete-registry-repository
func (s *RegistryService) DeleteRegistryRepository(pid interface{}, repositoryID int, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("projects/%s/registry/repositories/%d", url.QueryEscape(project), repositoryID)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListRegistryRepositoryTags gets a list of tags for given registry repository.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/container_registry.html#list-repository-tags
func (s *RegistryService) ListRegistryRepositoryTags(pid interface{}, repositoryID int, opt *ListRegistryRepositoryTagsOptions, options ...OptionFunc) ([]*RegistryRepositoryTag, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/registry/repositories/%d/tags", url.QueryEscape(project), repositoryID)

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var repos []*RegistryRepositoryTag
	resp, err := s.client.Do(req, &repos)
	if err != nil {
		return nil, resp, err
	}

	return repos, resp, err
}

// GetRegistryRepositoryTagDetail get details of a registry repository tag
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/container_registry.html#get-details-of-a-repository-tag
func (s *RegistryService) GetRegistryRepositoryTagDetail(pid interface{}, repositoryID int, tagName string, options ...OptionFunc) (*RegistryRepositoryTagDetail, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/registry/repositories/%d/tags/%s", url.QueryEscape(project), repositoryID, tagName)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var tagDetail *RegistryRepositoryTagDetail
	resp, err := s.client.Do(req, &tagDetail)
	if err != nil {
		return nil, resp, err
	}

	return tagDetail, resp, err
}

// DeleteRegistryRepositoryTag deletes a registry repository tag.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/container_registry.html#delete-a-repository-tag
func (s *RegistryService) DeleteRegistryRepositoryTag(pid interface{}, repositoryID int, tagName string, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("projects/%s/registry/repositories/%d/tags/%s", url.QueryEscape(project), repositoryID, tagName)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// BulkDeleteRegistryRepositoryTags deletes repository tags in bulk based on given criteria.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/container_registry.html#delete-repository-tags-in-bulk
func (s *RegistryService) BulkDeleteRegistryRepositoryTags(pid interface{}, repositoryID int, opt *BulkDeleteRegistryRepositoryTagsOptions, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("projects/%s/registry/repositories/%d/tags", url.QueryEscape(project), repositoryID)

	req, err := s.client.NewRequest("DELETE", u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
