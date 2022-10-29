package gitlab

import (
    "fmt"
    "net/http"
    "time"
)

// ProjectFeatureFlagService handles operations on gitlab project feature flags using the following api:
// GitLab API docs: https://docs.gitlab.com/ee/api/feature_flags.html
type ProjectFeatureFlagService struct {
    client *Client
}

// ProjectFeatureFlag represents a GitLab project iteration.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/feature_flags.html
type ProjectFeatureFlag struct {
    Name        string                       `json:"name"`
    Description string                       `json:"description"`
    Active      bool                         `json:"active"`
    Version     string                       `json:"version"`
    CreatedAt   *time.Time                   `json:"created_at"`
    UpdatedAt   *time.Time                   `json:"updated_at"`
    Scopes      []ProjectFeatureFlagScope    `json:"scopes"`
    Strategies  []ProjectFeatureFlagStrategy `json:"strategies"`
}

// ProjectFeatureFlagScope defines the scopes of a feature flag
type ProjectFeatureFlagScope struct {
    ID               int    `json:"id"`
    EnvironmentScope string `json:"environment_scope"`
}

// ProjectFeatureFlagStrategy defines the strategy used for a feature flag
type ProjectFeatureFlagStrategy struct {
    ID         int                                 `json:"id"`
    Name       string                              `json:"name"`
    Parameters ProjectFeatureFlagStrategyParameter `json:"parameters"`
    Scopes     []ProjectFeatureFlagScope           `json:"scopes"`
}

// ProjectFeatureFlagStrategyParameter is used in updating and creating feature flags
type ProjectFeatureFlagStrategyParameter struct {
    GroupID    *string `json:"groupId,omitempty"`
    UserIDs    *string `json:"userIds,omitempty"`
    Percentage *string `json:"percentage,omitempty"`
}

// UpdateProjectFeatureFlagOptions is used to specify the values when updating feature flags
type UpdateProjectFeatureFlagOptions struct {
    Name        string  `json:"name,omitempty"`
    Description *string `json:"description,omitempty"`
    Active      *bool   `json:"active,omitempty"`
    Strategies  []*FeatureFlagStrategy `json:"strategies,omitempty"`
}


// FeatureFlagStrategy represents a strategy for a feature flag
type FeatureFlagStrategy struct {
    Name       string                              `json:"Name,omitempty"`
    Parameters ProjectFeatureFlagStrategyParameter `json:"parameters,omitempty"`
    Scopes     []ProjectFeatureFlagScope           `json:"scopes,omitempty"`
}

// CreateProjectFeatureFlagOptions contains the options that can be specified when creating feature flags
type CreateProjectFeatureFlagOptions struct {
    Name        string  `json:"name"`
    Description *string `json:"description,omitempty"`
    Version     *string `json:"version,omitempty"`
    Active      *bool   `json:"active,omitempty"`
    Strategies  []*FeatureFlagStrategy `json:"strategies,omitempty"`
}

func (i ProjectFeatureFlag) String() string {
    return Stringify(i)
}

// ListProjectFeatureFlagOptions contains the options for ListProjectFeatureFlags
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/feature_flags.html#list-feature-flags-for-a-project
type ListProjectFeatureFlagOptions struct {
    ListOptions
    Scope *string `url:"scope,omitempty" json:"scope,omitempty"`
}

// ListProjectFeatureFlags returns a list with the feature flags of a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/feature_flags.html#list-feature-flags-for-a-project
func (s *ProjectFeatureFlagService) ListProjectFeatureFlags(pid interface{}, opt *ListProjectFeatureFlagOptions, options ...RequestOptionFunc) ([]*ProjectFeatureFlag, *Response, error) {
    project, err := parseID(pid)
    if err != nil {
        return nil, nil, err
    }
    u := fmt.Sprintf("projects/%s/feature_flags", PathEscape(project))

    req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
    if err != nil {
        return nil, nil, err
    }

    var pis []*ProjectFeatureFlag
    resp, err := s.client.Do(req, &pis)
    if err != nil {
        return nil, resp, err
    }

    return pis, resp, err
}

// GetProjectFeatureFlag gets a single feature flag for the specified project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/feature_flags.html#get-a-single-feature-flag
func (s *ProjectFeatureFlagService) GetProjectFeatureFlag(pid interface{}, name string, options ...RequestOptionFunc) (*ProjectFeatureFlag, *Response, error) {
    project, err := parseID(pid)
    if err != nil {
        return nil, nil, err
    }
    u := fmt.Sprintf("projects/%s/feature_flags/%s", PathEscape(project), name)

    req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
    if err != nil {
        return nil, nil, err
    }

    b := new(ProjectFeatureFlag)
    resp, err := s.client.Do(req, b)
    if err != nil {
        return nil, resp, err
    }

    return b, resp, err
}

// CreateProjectFeatureFlag creates a feature flag
//
// Gitlab API docs:
// https://docs.gitlab.com/ee/api/feature_flags.html#create-a-feature-flag
func (s *ProjectFeatureFlagService) CreateProjectFeatureFlag(pid interface{}, opt *CreateProjectFeatureFlagOptions, options ...RequestOptionFunc) (*ProjectFeatureFlag, *Response, error) {
    project, err := parseID(pid)
    if err != nil {
        return nil, nil, err
    }
    u := fmt.Sprintf("projects/%s/feature_flags",
        PathEscape(project),
    )

    req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
    if err != nil {
        return nil, nil, err
    }

    var flag ProjectFeatureFlag
    resp, err := s.client.Do(req, &flag)
    if err != nil {
        return &flag, resp, err
    }

    return &flag, resp, err
}

// UpdateProjectFeatureFlag updates a feature flag
//
// Gitlab API docs:
// https://docs.gitlab.com/ee/api/feature_flags.html#update-a-feature-flag
func (s *ProjectFeatureFlagService) UpdateProjectFeatureFlag(pid interface{}, name string, opt *UpdateProjectFeatureFlagOptions, options ...RequestOptionFunc) (*ProjectFeatureFlag, *Response, error) {
    group, err := parseID(pid)
    if err != nil {
        return nil, nil, err
    }
    u := fmt.Sprintf("projects/%s/feature_flags/%s",
        PathEscape(group),
        name,
    )

    req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
    if err != nil {
        return nil, nil, err
    }

    var flag ProjectFeatureFlag
    resp, err := s.client.Do(req, &flag)
    if err != nil {
        return &flag, resp, err
    }

    return &flag, resp, err
}

// DeleteProjectFeatureFlag deletes a feature flag
//
// Gitlab API docs:
// https://docs.gitlab.com/ee/api/feature_flags.html#delete-a-feature-flag
func (s *ProjectFeatureFlagService) DeleteProjectFeatureFlag(pid interface{}, name string, options ...RequestOptionFunc) (*Response, error) {
    project, err := parseID(pid)
    if err != nil {
        return nil, err
    }
    u := fmt.Sprintf("projects/%s/feature_flags/%s", PathEscape(project), name)

    req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
    if err != nil {
        return nil, err
    }

    return s.client.Do(req, nil)
}
