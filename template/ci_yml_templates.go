package template

import (
	"fmt"
	"net/url"
	"strings"

)

// CIYMLTemplatesService handles communication with the gitlab
// CI YML templates related methods of the GitLab API.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/templates/gitlab_ci_ymls.html
type CIYMLTemplatesService struct {
	client *gitlab.Client
}

func NewCITemplate(c *gitlab.Client) CIYMLTemplatesService {
	return CIYMLTemplatesService{client: c}
}

// CIYMLTemplate represents a GitLab CI YML template.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/templates/gitlab_ci_ymls.html
type CIYMLTemplate struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// ListCIYMLTemplatesOptions represents the available ListAllTemplates() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/templates/gitignores.html#list-gitignore-templates
type ListCIYMLTemplatesOptions gitlab.ListOptions

// ListAllTemplates get all GitLab CI YML templates.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/templates/gitlab_ci_ymls.html#list-gitlab-ci-yml-templates
func (s *CIYMLTemplatesService) ListAllTemplates(opt *ListCIYMLTemplatesOptions, options ...gitlab.OptionFunc) ([]*CIYMLTemplate, *gitlab.Response, error) {
	req, err := s.client.NewRequest("GET", "templates/gitlab_ci_ymls", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var cts []*CIYMLTemplate
	resp, err := s.client.Do(req, &cts)
	if err != nil {
		return nil, resp, err
	}

	return cts, resp, err
}

// GetTemplate get a single GitLab CI YML template.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/templates/gitlab_ci_ymls.html#single-gitlab-ci-yml-template
func (s *CIYMLTemplatesService) GetTemplate(key string, options ...gitlab.OptionFunc) (*CIYMLTemplate, *gitlab.Response, error) {
	u := fmt.Sprintf("templates/gitlab_ci_ymls/%s", pathEscape(key))

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	ct := new(CIYMLTemplate)
	resp, err := s.client.Do(req, ct)
	if err != nil {
		return nil, resp, err
	}

	return ct, resp, err
}

// Helper function to escape a project identifier.
// Duplication of unexported gitlab.pathEscape()
func pathEscape(s string) string {
	return strings.Replace(url.PathEscape(s), ".", "%2E", -1)
}
