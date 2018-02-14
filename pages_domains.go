package gitlab

import (
	"fmt"
	"net/url"
	"time"
)

// PagesDomainsService handles Pages domains.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/pages_domains.html
type PagesDomainsService struct {
	client *Client
}

type Cert struct {
	Expired    bool       `json:"expired"`
	Expiration *time.Time `json:"expiration"`
}

// PageDomain represents a pages domain.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/pages_domains.html
type PagesDomain struct {
	Domain      string `json:"domain"`
	URL         string `json:"url"`
	ProjectID   int    `json:"project_id"`
	Certificate *Cert  `json:"certificate"`
}

// ListPagesDomainsOptions represents the available ListPagesDomains() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/pages_domais.html#list-pages-domains
type ListPagesDomainsOptions struct {
	ListOptions
}

// ListPagesDomains gets a list of project Domains.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/pages_domains.html#list-pages-domains
func (s *PagesDomainsService) ListPagesDomains(pid interface{}, opt *ListPagesDomainsOptions, options ...OptionFunc) ([]*PagesDomain, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pages/domains", url.QueryEscape(project))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var pt []*PagesDomain
	resp, err := s.client.Do(req, &pt)
	if err != nil {
		return nil, resp, err
	}

	return pt, resp, err
}

// GetPagesDomains gets a specific Pages Domains for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/Pages_Domainss.html#get-Domains-details
func (s *PagesDomainsService) GetPagesDomains(pid interface{}, Domains int, options ...OptionFunc) (*PagesDomain, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pages/domains/%d", url.QueryEscape(project), Domains)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	pt := new(PagesDomain)
	resp, err := s.client.Do(req, pt)
	if err != nil {
		return nil, resp, err
	}

	return pt, resp, err
}

// AddPagesDomainOptions represents the available AddPagesDomain() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/pipeline_triggers.html#create-a-project-trigger
type AddPagesDomainOptions struct {
	Domain      *string `url:"domain,omitempty" json:"domain,omitempty"`
	Certificate *string `url:"certifiate,omitempty" json:"certifiate,omitempty"`
	Key         *string `url:"key,omitempty" json:"key,omitempty"`
}

// AddPagesDomain adds a pages domain to a specified project.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/pipeline_triggers.html#create-a-project-trigger
func (s *PagesDomainsService) AddPagesDomain(pid interface{}, opt *AddPagesDomainOptions, options ...OptionFunc) (*PagesDomain, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pages/domains", url.QueryEscape(project))

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pt := new(PagesDomain)
	resp, err := s.client.Do(req, pt)
	if err != nil {
		return nil, resp, err
	}

	return pt, resp, err
}

// EditPagesDomainOptions represents the available EditPagesDomain() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/pipeline_triggers.html#update-a-project-trigger
type EditPagesDomainOptions struct {
	ID         *int    `url:"id,omitempty" json:"id,omitempty"`
	Domain     *string `url:"domain,omitempty" json:"domain,omitempty"`
	Cerificate *string `url:"certifiate" json:"certifiate"`
	Key        *string `url:"key" json:"key"`
}

// EditPagesDomain edits a domain for a specified project.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/pipeline_triggers.html#update-a-project-trigger
func (s *PagesDomainsService) EditPagesDomain(pid interface{}, trigger int, opt *EditPagesDomainOptions, options ...OptionFunc) (*PagesDomain, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pages/domains/%d", url.QueryEscape(project), trigger)

	req, err := s.client.NewRequest("PUT", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pt := new(PagesDomain)
	resp, err := s.client.Do(req, pt)
	if err != nil {
		return nil, resp, err
	}

	return pt, resp, err
}

// DeletePagesDomain removes a domain from a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/pipeline_triggers.html#remove-a-project-trigger
func (s *PagesDomainsService) DeletePagesDomain(pid interface{}, trigger int, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/pages/domains/%d", url.QueryEscape(project), trigger)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
