package gitlab

import (
	"time"
)

// Member represents any of the following:
// - project member
// - group member
// - billable group member
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/members.html

type Member struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	State          string `json:"state"`
	MembershipType string `json:"membership_type"`

	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`

	CreatedAt      *time.Time `json:"created_at"`
	ExpiresAt      *ISOTime   `json:"expires_at"`
	LastLoginAt    *time.Time `json:"last_login_at"`
	LastActivityOn *ISOTime   `json:"last_activity_on"`

	AccessLevel       AccessLevelValue         `json:"access_level"`
	GroupSAMLIdentity *GroupMemberSAMLIdentity `json:"group_saml_identity"`

	Removable   bool `json:"removable"`
	IsLastOwner bool `json:"is_last_owner"`
}
