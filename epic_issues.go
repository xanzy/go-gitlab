package gitlab

import "fmt"

// EpicIssuesService handles communication with the epic issue related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/epic_issues.html
type EpicIssuesService struct {
	client *Client
}

// ListEpicIssues get a list of epic issues.
//
// Gitlab API docs:
// https://docs.gitlab.com/ee/api/epic_issues.html#list-issues-for-an-epic
func (s *EpicIssuesService) ListEpicIssues(gid interface{}, epic int, opt *ListOptions, options ...RequestOptionFunc) ([]*Issue, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/epics/%d/issues", pathEscape(group), epic)

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var issues []*Issue
	resp, err := s.client.Do(req, &issues)
	if err != nil {
		return nil, resp, err
	}

	return issues, resp, err
}

// EpicIssueAssignment contains both the Epic and Issue objects returned from Gitlab w/ the assignment ID
type EpicIssueAssignment struct {
	ID    int    `json:"id"`
	Epic  *Epic  `json:"epic"`
	Issue *Issue `json:"issue"`
}

// AssignEpicIssue assigns an existing issue to an Epic
//
// Gitlab API Docs: https://docs.gitlab.com/ee/api/epic_issues.html#assign-an-issue-to-the-epic
func (s *EpicIssuesService) AssignEpicIssue(gid interface{}, epic, issue int, options ...RequestOptionFunc) (*EpicIssueAssignment, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("groups/%s/epics/%d/issues/%d", pathEscape(group), epic, issue)

	req, err := s.client.NewRequest("POST", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var a *EpicIssueAssignment

	resp, err := s.client.Do(req, &a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, err
}

// RemoveEpicIssue removes an issue from an Epic
//
// Gitlab API Docs: https://docs.gitlab.com/ee/api/epic_issues.html#remove-an-issue-from-the-epic
func (s *EpicIssuesService) RemoveEpicIssue(gid interface{}, epic int, epicIssue int, options ...RequestOptionFunc) (*EpicIssueAssignment, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("groups/%s/epics/%d/issues/%d", pathEscape(group), epic, epicIssue)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var a *EpicIssueAssignment

	resp, err := s.client.Do(req, &a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, err
}

// UpdateEpicIsssueAssignmentOptions describes options to move issues within an epic
type UpdateEpicIsssueAssignmentOptions struct {
	*ListOptions
	MoveBeforeID int `json:"move_before_id"`
	MoveAfterID  int `json:"move_after_id"`
}

// UpdateEpicIssueAssignment moves an issue before or after another issue in an epic issue list
//
// Gitlab API Docs:
// https://docs.gitlab.com/ee/api/epic_issues.html#update-epic---issue-association
func (s *EpicIssuesService) UpdateEpicIssueAssignment(gid interface{}, epic int, epicIssue int, opt *UpdateEpicIsssueAssignmentOptions, options ...RequestOptionFunc) ([]*Issue, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/epics/%d/issues/%d", pathEscape(group), epic, epicIssue)

	req, err := s.client.NewRequest("PUT", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var issues []*Issue
	resp, err := s.client.Do(req, &issues)
	if err != nil {
		return nil, resp, err
	}

	return issues, resp, err
}
