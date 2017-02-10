package gitlab

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

// NotificationEvent represents an event that should trigger email notifications if notification level is set to "custom""
type NotificationEvent string

// List of notification email events (used together with CustomNotificationLevel)
const (
	NewNoteEvent              string = "new_note"
	NewIssueEvent                    = "new_issue"
	ReOpenIssueEvent                 = "reopen_issue"
	CloseIssueEvent                  = "close_issue"
	ReAssignIssueEvent               = "reassign_issue"
	NewMergeRequestEvent             = "new_merge_request"
	ReOpenMergeRequestEvent          = "reopen_merge_request"
	CloseMergeRequestEvent           = "close_merge_request"
	ReAssignMergeRequestEvent        = "reassign_merge_request"
	MergeMergeRequestEvent           = "merge_merge_request"
	FailedPipelineEvent              = "failed_pipeline"
	SuccessPipelineEvent             = "success_pipeline"
)

// NotificationSettings represents a Gitlab notification setting
type NotificationSettings struct {
	Level  NotificationLevel          `json:"level"`
	Email  string                     `json:"notification_email"`
	Events map[NotificationEvent]bool `json:"events,omitempty"`
}

// NotificationSettingsOptions represents an object used to create/modify notification settings.
// By default, if AsUser is nil the API calls will be performed as the authenticated user.
// If AsUser contains a Gitlab user ID/username, the API calls will be performed as that user if the authenticated user is an administrator.
// If the authenticated user is not admin, a 403 error will be returned. If the provided user ID/username cannot be found, a 404 error will be returned.
// See https://docs.gitlab.com/ce/api/README.html#sudo for more details
type NotificationSettingsOptions struct {
	AsUser   *string
	Settings NotificationSettings
}

func (settings NotificationSettings) String() string {
	return Stringify(settings)
}

// GetGlobalNotificationSettings returns current notification settings and email address.
// GitLab API docs:
// https://docs.gitlab.com/ce/api/notification_settings.html#global-notification-settings
func (s *NotificationSettingsService) GetGlobalNotificationSettings() (*NotificationSettings, *Response, error) {
	req, err := s.client.NewRequest("GET", "notification_settings", nil)
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

func (s *NotificationSettingsService) UpdateGlobalNotificationSettings(opt NotificationSettingsOptions) (*NotificationSettings, *Response, error) {

	req, err := s.client.NewRequest("PUT", "notification_settings", opt)

	if err != nil {
		return nil, nil, err
	}

	if opt.AsUser != nil {
		req.Header.Set("SUDO", *opt.AsUser)
	}

	ns := new(NotificationSettings)
	resp, err := s.client.Do(req, ns)
	if err != nil {
		return nil, resp, err
	}

	return ns, resp, err

}
