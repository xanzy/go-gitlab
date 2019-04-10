package gitlab

import (
	"fmt"
)

// AddGroupLDAPLinkOptions represents the available AddGroupLDAPLink() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/groups.html#add-ldap-group-link
type AddGroupLDAPLinkOptions struct {
	ID          int    `url:"id,omitempty" json:"id,omitempty"`
	CN          string `url:"cn,omitempty" json:"cn,omitempty"`
	GroupAccess int    `url:"group_access,omitempty" json:"group_access,omitempty"`
	Provider    string `url:"provider,omitempty" json:"provider,ommitempty"`
}

// AddGroupLDAPLink creates a new group LDAP link. Available only for users who can
// edit groups.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/groups.html#add-ldap-group-link
func (s *GroupsService) AddGroupLDAPLink(opt *AddGroupLDAPLinkOptions, options ...OptionFunc) (*Response, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("groups/%d/ldap_group_links", opt.ID), opt, options)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, nil)
}
