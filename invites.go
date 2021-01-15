package gitlab

import "time"

type ListPendingInvitationsOptions struct {
    ListOptions
    Query *string `url:"query,omitempty" json:"query,omitempty"`
}

type PendingInvitations struct {
    Id            int    `json:"id"`
    InviteEmail   string `json:"invite_email"`
    CreatedAt     string `json:"created_at"`
    AccessLevel   int    `json:"access_level"`
    ExpiresAt     string `json:"expires_at"`
    UserName      string `json:"user_name"`
    CreatedByName string `json:"created_by_name"`
}

type InvitationsResponse struct {
    Status  string            `json:"status"`
    Message map[string]string `json:"message,omitempty"`
}

type InvitesOptions struct {
    Id          string     `url:"id,omitempty" json:"id,omitempty"`
    Email       string     `url:"email,omitempty" json:"email,omitempty"`
    AccessLevel string     `url:"access_level,omitempty" json:"access_level,omitempty"`
    ExpiresAt   *time.Time `url:"expires_at,omitempty" json:"expires_at,omitempty"`
}
