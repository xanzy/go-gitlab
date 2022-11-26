package gitlab

import (
	"fmt"
	"net/http"
)

type DORAMetricType string

const (
	DORAMetricDeploymentFrequency  DORAMetricType = "deployment_frequency"
	DORAMetricLeadTimeForChanges   DORAMetricType = "lead_time_for_changes"
	DORAMetricTimeToRestoreService DORAMetricType = "time_to_restore_service"
	DORAMetricChangeFailureRate    DORAMetricType = "change_failure_rate"
)

type DORAMetricInterval string

const (
	// daily is the default
	DORAMetricIntervalDaily   DORAMetricInterval = "daily"
	DORAMetricIntervalMonthly DORAMetricInterval = "monthly"
	DORAMetricIntervalAll     DORAMetricInterval = "all"
)

// DORAMetric represents a single DORA metric data point.
//
// Gitlab API docs: https://docs.gitlab.com/ee/api/dora/metrics.html
type DORAMetric struct {
	Date  string `json:"date"`
	Value int    `json:"value"`
}

func (m DORAMetric) String() string {
	return Stringify(m)
}

type GetDORAMetricsOptions struct {
	Metric           DORAMetricType `url:"metric" json:"metric"`
	EndDate          *string        `url:"end_date,omitempty" json:"end_date,omitempty"`
	EnvironmentTiers *[]string      `url:"environment_tiers,omitempty" del:"," json:"environment_tiers,omitempty"`
	// Interval can be one of daily, monthly, all, default is daily
	Interval  *DORAMetricInterval `url:"interval,omitempty" json:"interval,omitempty"`
	StartDate *string             `url:"start_date,omitempty" json:"start_date,omitempty"`

	// Deprecated, use environment tiers instead
	EnvironmentTier *string `url:"environment_tier,omitempty" json:"environment_tier,omitempty"`
}

// DORAMetricsService handles communication with the DORA metrics related methods of the GitLab API
//
// Gitlab API docs: https://docs.gitlab.com/ee/api/dora/metrics.html
type DORAMetricsService struct {
	client *Client
}

// GetProjectDORAMetrics gets the DORA metrics for a project.
//
// GitLab API Docs: https://docs.gitlab.com/ee/api/dora/metrics.html#get-project-level-dora-metrics
func (s *DORAMetricsService) GetProjectDORAMetrics(pid interface{}, opt GetDORAMetricsOptions, options ...RequestOptionFunc) ([]DORAMetric, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	if opt.Metric == "" {
		return nil, nil, fmt.Errorf("metric cannot be empty")
	}

	u := fmt.Sprintf("projects/%s/dora/metrics", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var metrics []DORAMetric
	resp, err := s.client.Do(req, &metrics)
	if err != nil {
		return nil, resp, err
	}

	return metrics, resp, err
}

// GetGroupDORAMetrics gets the DORA metrics for a group.
//
// GitLab API Docs: https://docs.gitlab.com/ee/api/dora/metrics.html#get-group-level-dora-metrics
func (s *DORAMetricsService) GetGroupDORAMetrics(gid interface{}, opt GetDORAMetricsOptions, options ...RequestOptionFunc) ([]DORAMetric, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}

	if opt.Metric == "" {
		return nil, nil, fmt.Errorf("metric cannot be empty")
	}

	u := fmt.Sprintf("groups/%s/dora/metrics", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var metrics []DORAMetric
	resp, err := s.client.Do(req, &metrics)
	if err != nil {
		return nil, resp, err
	}

	return metrics, resp, err
}
