package gitlab

import (
	"fmt"
	"net/url"
	"time"
)

// AccessRequest represents a access request for a group or project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/access_requests.html
// TODO: try to find existed structure for that
type AccessRequest struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Name      string     `json:"name"`
	State     string     `json:"state"`
	CreatedAt *time.Time `json:"created_at"`
	// RequestedAt represents a access request requested time. Can be empty, when used
	// ApproveProjectAccessRequest() or ApproveGroupAccessRequest().
	RequestedAt *time.Time `json:"requested_at"`
	// AccessLevel represents a user access level. Can be empty, when used ListProjectAccessRequests(),
	// ListGroupAccessRequests(), RequestProjectAccess(), RequestGroupAccess().
	AccessLevel *AccessLevelValue `json:"access_level,omitempty"`
}

// AccessRequestsService handles communication with the project/group access requests
// related methods of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/access_requests.html
type AccessRequestsService struct {
	client *Client
}

// ListAccessRequestsOptions represents the available ListProjectAccessRequests() or ListGroupAccessRequests()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/access_requests.html#list-access-requests-for-a-group-or-project
type ListAccessRequestsOptions ListOptions

// ListProjectAccessRequests gets a list of access requests
// viewable by the authenticated user.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/access_requests.html#list-access-requests-for-a-group-or-project
func (s *AccessRequestsService) ListProjectAccessRequests(pid interface{}, opt *ListAccessRequestsOptions, options ...OptionFunc) ([]*AccessRequest, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/access_requests", url.QueryEscape(project))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var ar []*AccessRequest
	resp, err := s.client.Do(req, &ar)
	if err != nil {
		return nil, resp, err
	}

	return ar, resp, err
}

// ListGroupAccessRequests gets a list of access requests
// viewable by the authenticated user.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/access_requests.html#list-access-requests-for-a-group-or-project
func (s *AccessRequestsService) ListGroupAccessRequests(pid interface{}, opt *ListAccessRequestsOptions, options ...OptionFunc) ([]*AccessRequest, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/access_requests", url.QueryEscape(project))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var ar []*AccessRequest
	resp, err := s.client.Do(req, &ar)
	if err != nil {
		return nil, resp, err
	}

	return ar, resp, err
}

// RequestProjectAccess requests access for the authenticated user to a group or project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/access_requests.html#request-access-to-a-group-or-project
func (s *AccessRequestsService) RequestProjectAccess(pid interface{}, options ...OptionFunc) (*AccessRequest, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/access_requests", url.QueryEscape(project))

	req, err := s.client.NewRequest("POST", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	ar := new(AccessRequest)
	resp, err := s.client.Do(req, ar)
	if err != nil {
		return nil, resp, err
	}

	return ar, resp, err
}

// RequestGroupAccess requests access for the authenticated user to a group or project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/access_requests.html#request-access-to-a-group-or-project
func (s *AccessRequestsService) RequestGroupAccess(pid interface{}, options ...OptionFunc) (*AccessRequest, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/access_requests", url.QueryEscape(project))

	req, err := s.client.NewRequest("POST", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	ar := new(AccessRequest)
	resp, err := s.client.Do(req, ar)
	if err != nil {
		return nil, resp, err
	}

	return ar, resp, err
}

// ApproveAccessRequestOptions represents the available ApproveProjectAccessRequest()
// and ApproveGroupAccessRequest() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/access_requests.html#approve-an-access-request
type ApproveAccessRequestOptions struct {
	AccessLevel *AccessLevelValue `url:"access_level,omitempty" json:"access_level,omitempty"`
}

// ApproveProjectAccessRequest approves an access request for the given user.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/access_requests.html#approve-an-access-request
func (s *AccessRequestsService) ApproveProjectAccessRequest(pid interface{}, user int, opt *ApproveAccessRequestOptions, options ...OptionFunc) (*AccessRequest, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/access_requests/%d/approve", url.QueryEscape(project), user)

	req, err := s.client.NewRequest("PUT", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	ar := new(AccessRequest)
	resp, err := s.client.Do(req, ar)
	if err != nil {
		return nil, resp, err
	}

	return ar, resp, err
}

// ApproveGroupAccessRequest approves an access request for the given user.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/access_requests.html#approve-an-access-request
func (s *AccessRequestsService) ApproveGroupAccessRequest(pid interface{}, user int, opt *ApproveAccessRequestOptions, options ...OptionFunc) (*AccessRequest, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/access_requests/%d/approve", url.QueryEscape(project), user)

	req, err := s.client.NewRequest("PUT", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	ar := new(AccessRequest)
	resp, err := s.client.Do(req, ar)
	if err != nil {
		return nil, resp, err
	}

	return ar, resp, err
}

// DenyProjectAccessRequest denies an access request for the given user.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/access_requests.html#deny-an-access-request
func (s *AccessRequestsService) DenyProjectAccessRequest(pid interface{}, user int, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/access_requests/%d", url.QueryEscape(project), user)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DenyGroupAccessRequest denies an access request for the given user.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/access_requests.html#deny-an-access-request
func (s *AccessRequestsService) DenyGroupAccessRequest(pid interface{}, user int, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("groups/%s/access_requests/%d", url.QueryEscape(project), user)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
