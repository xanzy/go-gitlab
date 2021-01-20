package gitlab

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// AuditEvent represents an audit event for a group or project
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/audit_events.html
type AuditEvent struct {
	ID         int               `json:"id"`
	AuthorID   int               `json:"author_id"`
	EntityID   int               `json:"entity_id"`
	EntityType string            `json:"entity_type"`
	Details    AuditEventDetails `json:"details"`
	CreatedAt  *time.Time        `json:"created_at"`
}

// AuditEvent represents the details portion of an audit event for a group or project
// The exact fields that are returned for an audit event depend on the action being recorded
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/audit_events.html
type AuditEventDetails struct {
	With          string      `json:"with"`
	Add           string      `json:"add"`
	As            string      `json:"as"`
	Change        string      `json:"change"`
	From          string      `json:"from"`
	To            string      `json:"to"`
	Remove        string      `json:"remove"`
	CustomMessage string      `json:"custom_message"`
	AuthorName    string      `json:"author_name"`
	TargetID      TargetID    `json:"target_id"` // Sometimes this is an int and sometimes is a string with project path
	TargetType    string      `json:"target_type"`
	TargetDetails string      `json:"target_details"`
	IPAddress     string      `json:"ip_address"`
	EntityPath    string      `json:"entity_path"`
}

// TargetID is a custom type so we can handle unmarshalling this since it is sometimes a string and sometimes an int
type TargetID string

// UnmarshalJSON Unmarshals the TargetID to a string, even if its an int in the json
// TargetID is a mixed return value in the GitLab API
func (t *TargetID) UnmarshalJSON(data []byte) error {
	var i interface{}
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}

	switch v := i.(type) {
	case float64:
		*t = TargetID(strconv.FormatFloat(v, 'f', -1, 64))
	case string:
		*t = TargetID(v)
	}

	return nil
}

// AuditEventsService handles communication with the project/group
// audit event related methods of the GitLab API.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/audit_events.html
type AuditEventsService struct {
	client *Client
}

// ListAuditEventsOptions represents the available
// ListProjectAuditEvents() or ListGroupAuditEvents() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/audit_events.html
type ListAuditEventsOptions struct {
	ListOptions
	CreatedAfter  *time.Time `url:"created_after,omitempty" json:"created_after,omitempty"`
	CreatedBefore *time.Time `url:"created_before,omitempty" json:"created_before,omitempty"`
}

// ListProjectAuditEvents gets a list of audit events for the specified project
// viewable by the authenticated user.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/audit_events.html
func (s *AuditEventsService) ListProjectAuditEvents(pid interface{}, opt *ListAuditEventsOptions, options ...RequestOptionFunc) ([]*AuditEvent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/audit_events", pathEscape(project))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var aes []*AuditEvent
	resp, err := s.client.Do(req, &aes)
	if err != nil {
		return nil, resp, err
	}

	return aes, resp, err
}

// GetProjectAuditEvent gets a specific project audit event
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/audit_events.html
func (s *AuditEventsService) GetProjectAuditEvent(pid interface{}, aid interface{}, options ...RequestOptionFunc) (*AuditEvent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	auditEventID, err := parseID(aid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/audit_events/%s", pathEscape(project), pathEscape(auditEventID))

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	a := new(AuditEvent)
	resp, err := s.client.Do(req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, err
}

// ListGroupAuditEvents gets a list of audit events for the specified group
// viewable by the authenticated user.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/audit_events.html
func (s *AuditEventsService) ListGroupAuditEvents(gid interface{}, opt *ListAuditEventsOptions, options ...RequestOptionFunc) ([]*AuditEvent, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/audit_events", pathEscape(group))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var aes []*AuditEvent
	resp, err := s.client.Do(req, &aes)
	if err != nil {
		return nil, resp, err
	}

	return aes, resp, err
}

// GetGroupAuditEvent gets a specific group audit event
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/audit_events.html
func (s *AuditEventsService) GetGroupAuditEvent(gid interface{}, aid interface{}, options ...RequestOptionFunc) (*AuditEvent, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}

	auditEventID, err := parseID(aid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("groups/%s/audit_events/%s", pathEscape(group), pathEscape(auditEventID))

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	a := new(AuditEvent)
	resp, err := s.client.Do(req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, err
}
