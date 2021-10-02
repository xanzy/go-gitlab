package gitlab

import (
	"fmt"
	"net/http"
)

// ExternalStatusChecksService handles communication with the external status check
//  related methods of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/status_checks.html
type ExternalStatusChecksService struct {
	client *Client
}

type StatusCheck struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ExternalURL string `json:"external_url"`
	Status      string `json:"status"`
}

// ListExternalStatusChecks lists the external status checks that apply to it and their status
// for a single merge request
//
// GitLab API docs: https://docs.gitlab.com/ee/api/status_checks.html#list-status-checks-for-a-merge-request
func (s *ExternalStatusChecksService) ListExternalStatusChecks(pid interface{}, mr int, opt *ListOptions, options ...RequestOptionFunc) ([]*StatusCheck, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_requests/%d/status_checks", pathEscape(project), mr)

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var scs []*StatusCheck
	resp, err := s.client.Do(req, &scs)
	if err != nil {
		return nil, resp, err
	}

	return scs, resp, err
}
