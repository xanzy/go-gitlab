package gitlab

import "net/http"

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
	DisabledNotificationLevel      string = "disabled"
	ParticipatingNotificationLevel        = "participating"
	WatchNotificationLevel                = "watch"
	GlobalNotificationLevel               = "global"
	MentionNotificationLevel              = "mention"
	CustomNotificationLevel               = "custom"
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

// // List of notification email events (used together with CustomNotificationLevel)
// const (
// 	NewNoteEvent              string = "new_note"
// 	NewIssueEvent                    = "new_issue"
// 	ReOpenIssueEvent                 = "reopen_issue"
// 	CloseIssueEvent                  = "close_issue"
// 	ReAssignIssueEvent               = "reassign_issue"
// 	NewMergeRequestEvent             = "new_merge_request"
// 	ReOpenMergeRequestEvent          = "reopen_merge_request"
// 	CloseMergeRequestEvent           = "close_merge_request"
// 	ReAssignMergeRequestEvent        = "reassign_merge_request"
// 	MergeMergeRequestEvent           = "merge_merge_request"
// 	FailedPipelineEvent              = "failed_pipeline"
// 	SuccessPipelineEvent             = "success_pipeline"
// )

// NotificationSettings represents a Gitlab notification setting
type NotificationSettings struct {
	Level                     NotificationLevel `json:"level"`
	Email                     string            `json:"notification_email"`
	*NotificationEventOptions `json:"events,omitempty"`
}

// ImpersonateUserOption can be used to impersonate the API call as an arbitrary user.
// The authenticated user must be an administrator. User can be a username or user ID.
// See https://docs.gitlab.com/ce/api/README.html#sudo for more information
type ImpersonateUserOption struct {
	User string
}

// GetGlobalSettings returns current notification settings and email address.
// If opt is nil, the API call will be performed as the authenticated user.
// Otherwise the API call wlil be performed as the impersonated user (SUDO).
// GitLab API docs:
// https://docs.gitlab.com/ce/api/notification_settings.html#global-notification-settings
func (s *NotificationSettingsService) GetGlobalSettings(opt *ImpersonateUserOption) (*NotificationSettings, *Response, error) {
	req, err := s.client.NewRequest("GET", "notification_settings", nil)
	if err != nil {
		return nil, nil, err
	}

	setSudo(req, opt)

	ns := new(NotificationSettings)
	resp, err := s.client.Do(req, ns)
	if err != nil {
		return nil, resp, err
	}

	return ns, resp, err
}

// UpdateGlobalSettings updates current notification settings and email address.
// If opt is nil, the API call will be performed as the authenticated user.
// Otherwise the API call wlil be performed as the impersonated user (SUDO).
// GitLab API docs:
// https://docs.gitlab.com/ce/api/notification_settings.html#update-global-notification-settings
func (s *NotificationSettingsService) UpdateGlobalSettings(settings NotificationSettings, opt *ImpersonateUserOption) (*NotificationSettings, *Response, error) {

	req, err := s.client.NewRequest("PUT", "notification_settings", settings)

	if err != nil {
		return nil, nil, err
	}

	setSudo(req, opt)

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

// // MarshalJSON marshalls the NotificationSettings struct to a JSON representation expected by the API.
// func (s NotificationSettings) MarshalJSON([]byte, error) {

// }

func setSudo(req *http.Request, opt *ImpersonateUserOption) {
	if req != nil && opt != nil && len(opt.User) > 0 {
		req.Header.Set("SUDO", opt.User)
	}
}
