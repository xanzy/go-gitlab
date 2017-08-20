package gitlab

import (
	"fmt"
	"net/url"
	"strings"
)

// FeaturesService handles the communication with the application FeaturesService
// related methods of the GitLab API.
//
// GitLab API docs:
// https://gitlab.com/gitlab-org/gitlab-ce/blob/9-3-stable/doc/api/features.md
type FeaturesService struct {
	client *Client
}

// Feature represents a GitLab feature flag.
//
// GitLab API docs:
// https://gitlab.com/gitlab-org/gitlab-ce/blob/9-3-stable/doc/api/features.md
type Feature struct {
	Name  string `json:"name"`
	State string `json:"state"`
	Gates []Gate
}

// Gate represents a gate of a GitLab feature flag.
//
// GitLab API docs:
// https://gitlab.com/gitlab-org/gitlab-ce/blob/9-3-stable/doc/api/features.md
type Gate struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func (f Feature) String() string {
	return Stringify(f)
}

// ListFeatures gets a list of feature flags
//
// GitLab API docs:
// https://gitlab.com/gitlab-org/gitlab-ce/blob/9-3-stable/doc/api/features.md
func (s *FeaturesService) ListFeatures(options ...OptionFunc) ([]*Feature, *Response, error) {
	req, err := s.client.NewRequest("GET", "features", nil, options)
	if err != nil {
		return nil, nil, err
	}

	var f []*Feature
	resp, err := s.client.Do(req, &f)
	if err != nil {
		return nil, resp, err
	}
	return f, resp, err
}

// SetFeatureFlag sets or creates a feature flag gate
//
// GitLab API docs:
// https://gitlab.com/gitlab-org/gitlab-ce/blob/9-3-stable/doc/api/features.md
func (s *FeaturesService) SetFeatureFlag(name, key, value string, options ...OptionFunc) (*Feature, *Response, error) {
	u := fmt.Sprintf("features/%s", url.QueryEscape(name))

	values := url.Values{}
	values.Add(key, value)

	req, err := s.client.NewRequest("POST", u, strings.NewReader(values.Encode()), options)
	if err != nil {
		return nil, nil, err
	}

	f := &Feature{}
	resp, err := s.client.Do(req, f)
	if err != nil {
		return nil, resp, err
	}
	return f, resp, err
}
