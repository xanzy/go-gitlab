package gitlab

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// NotificationSettingsService handles communication with the notification settings related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/notification_settings.html
type NotificationSettingsService struct {
	client *Client
}

// NotificationLevel represents a notification level
type NotificationLevel string

// List of valid notification levels
const (
	DisabledNotificationLevel      = "disabled"
	ParticipatingNotificationLevel = "participating"
	WatchNotificationLevel         = "watch"
	GlobalNotificationLevel        = "global"
	MentionNotificationLevel       = "mention"
	CustomNotificationLevel        = "custom"
)

// NotificationEventOptions specifies which events should trigger email notifications if notification level is set to "custom""
type NotificationEventOptions struct {
	NewNote                   bool `json:"new_note"`
	NewIssueEvent             bool `json:"new_issue"`
	ReOpenIssueEvent          bool `json:"reopen_issue"`
	CloseIssueEvent           bool `json:"close_issue"`
	ReAssignIssueEvent        bool `json:"reassign_issue"`
	NewMergeRequestEvent      bool `json:"new_merge_request"`
	ReOpenMergeRequestEvent   bool `json:"reopen_merge_request"`
	CloseMergeRequestEvent    bool `json:"close_merge_request"`
	ReAssignMergeRequestEvent bool `json:"reassign_merge_request"`
	MergeMergeRequestEvent    bool `json:"merge_merge_request"`
	FailedPipelineEvent       bool `json:"failed_pipeline"`
	SuccessPipelineEvent      bool `json:"success_pipeline"`
}

// NotificationSettings represents a Gitlab notification setting
type NotificationSettings struct {
	Level  NotificationLevel         `json:"level"`
	Email  string                    `json:"notification_email,omitempty"`
	Events *NotificationEventOptions `json:"events,omitempty"`
}

// GetGlobalSettings returns current notification settings and email address.
// The API call can be performed as another user if the parameter "sudoUser" contains a valid user ID or username.
// Otherwise the call is performed as the authenticated user.
// The authenticated user must be administrator in order to impersonate another user.
// If the authenticated user is not administrator or the provided user ID/uername is not found an error is returned.
// GitLab API docs:
// https://docs.gitlab.com/ce/api/notification_settings.html#global-notification-settings
// https://docs.gitlab.com/ce/api/README.html#sudo
func (s *NotificationSettingsService) GetGlobalSettings(sudoUser interface{}) (*NotificationSettings, *Response, error) {
	req, err := s.client.NewRequest("GET", "notification_settings", nil)
	if err != nil {
		return nil, nil, err
	}

	if err := sudo(req, sudoUser); err != nil {
		return nil, nil, err
	}

	ns := new(NotificationSettings)
	resp, err := s.client.Do(req, ns)
	if err != nil {
		return nil, resp, err
	}

	return ns, resp, err
}

// UpdateGlobalSettings updates current notification settings and email address.
// The API call can be performed as another user if the parameter "sudoUser" contains a valid user ID or username.
// Otherwise the call is performed as the authenticated user.
// The authenticated user must be administrator in order to impersonate another user.
// If the authenticated user is not administrator or the provided user ID/uername is not found an error is returned.
// GitLab API docs:
// https://docs.gitlab.com/ce/api/notification_settings.html#update-global-notification-settings
// https://docs.gitlab.com/ce/api/README.html#sudo
func (s *NotificationSettingsService) UpdateGlobalSettings(settings NotificationSettings, sudoUser interface{}) (*NotificationSettings, *Response, error) {

	if settings.Level == GlobalNotificationLevel {
		return nil, nil, errors.New("notification level 'global' is not valid for global notification settings")
	}

	u := fmt.Sprintf("notification_settings?level=%s", url.QueryEscape(string(settings.Level)))

	if len(settings.Email) > 0 {
		u += fmt.Sprintf("&notification_email=%s", url.QueryEscape(settings.Email))
	}

	req, err := s.client.NewRequest("PUT", u, settings.Events)
	if err != nil {
		return nil, nil, err
	}

	if err := sudo(req, sudoUser); err != nil {
		return nil, nil, err
	}

	ns := new(NotificationSettings)
	resp, err := s.client.Do(req, ns)
	if err != nil {
		return nil, resp, err
	}

	return ns, resp, err
}

// GetSettingsForGroup returns current group notification settings.
// id - the group ID or path
// Note: the email field is ignored for group level settings.
// GitLab API docs:
// https://docs.gitlab.com/ce/api/notification_settings.html#group-project-level-notification-settings
func (s *NotificationSettingsService) GetSettingsForGroup(id interface{}) (*NotificationSettings, *Response, error) {

	group, err := parseID(id)
	if err != nil {
		return nil, nil, err
	}

	url := fmt.Sprintf("groups/%s/notification_settings", url.QueryEscape(group))
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	ns := new(NotificationSettings)
	resp, err := s.client.Do(req, ns)
	if err != nil {
		return nil, resp, err
	}

	return ns, resp, err
}

// GetSettingsForProject returns current project notification settings.
// id - the project ID or path
// Note: the email field is ignored for project level settings.
// GitLab API docs:
// https://docs.gitlab.com/ce/api/notification_settings.html#group-project-level-notification-settings
func (s *NotificationSettingsService) GetSettingsForProject(id interface{}) (*NotificationSettings, *Response, error) {

	project, err := parseID(id)
	if err != nil {
		return nil, nil, err
	}

	url := fmt.Sprintf("projects/%s/notification_settings", url.QueryEscape(project))
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	ns := new(NotificationSettings)
	resp, err := s.client.Do(req, ns)
	if err != nil {
		return nil, resp, err
	}

	return ns, resp, err
}

// UpdateSettingsForGroup updates current group notification settings.
// id - the group ID or path (e.g. <group_name> or 123)
// Note: the email field is ignored for group level settings.
// GitLab API docs:
// https://docs.gitlab.com/ce/api/notification_settings.html#update-group-project-level-notification-settings
func (s *NotificationSettingsService) UpdateSettingsForGroup(id interface{}, settings NotificationSettings) (*NotificationSettings, *Response, error) {

	group, err := parseID(id)
	if err != nil {
		return nil, nil, err
	}

	url := fmt.Sprintf("groups/%s/notification_settings?level=%s",
		url.QueryEscape(group), url.QueryEscape(string(settings.Level)))

	req, err := s.client.NewRequest("PUT", url, settings.Events)

	if err != nil {
		return nil, nil, err
	}

	ns := new(NotificationSettings)
	resp, err := s.client.Do(req, ns)
	if err != nil {
		return nil, resp, err
	}

	return ns, resp, err
}

// UpdateSettingsForProject updates current project notification settings.
// id - the project ID or path (e.g. <group_name>/<project_name> or 123)
// Note: the email field is ignored for project level settings.
// GitLab API docs:
// https://docs.gitlab.com/ce/api/notification_settings.html#update-group-project-level-notification-settings
func (s *NotificationSettingsService) UpdateSettingsForProject(id interface{}, settings NotificationSettings) (*NotificationSettings, *Response, error) {

	project, err := parseID(id)
	if err != nil {
		return nil, nil, err
	}

	url := fmt.Sprintf("projects/%s/notification_settings?level=%s",
		url.QueryEscape(project), url.QueryEscape(string(settings.Level)))

	req, err := s.client.NewRequest("PUT", url, settings.Events)

	if err != nil {
		return nil, nil, err
	}

	ns := new(NotificationSettings)
	resp, err := s.client.Do(req, ns)
	if err != nil {
		return nil, resp, err
	}

	return ns, resp, err
}

func (s NotificationSettings) String() string {
	return Stringify(s)
}

func sudo(req *http.Request, sudoUser interface{}) error {
	if req != nil && sudoUser != nil {
		user, err := parseID(sudoUser)
		if err != nil {
			return err
		}
		req.Header.Set("SUDO", user)
		return nil
	}
	return nil
}
