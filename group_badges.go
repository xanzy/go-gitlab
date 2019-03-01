package gitlab

import (
	"fmt"
	"net/url"
)

// GroupBadge represents a group badge.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/group_badges.html
type GroupBadge struct {
	ID               int    `json:"id"`
	LinkURL          string `json:"link_url"`
	ImageURL         string `json:"image_url"`
	RenderedLinkURL  string `json:"rendered_link_url"`
	RenderedImageURL string `json:"rendered_image_url"`
	// Kind represents a project badge kind. Can be empty, when used PreviewProjectBadge().
	Kind string `json:"kind"`
}

type GroupBadgesService struct {
	client *Client
}

type ListGroupBadgesOptions ListOptions

func (s *GroupBadgesService) ListGroupBadges(gid interface{}, opt *ListGroupBadgesOptions, options ...OptionFunc) ([]*GroupBadge, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/badges", url.QueryEscape(group))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var gb []*GroupBadge
	resp, err := s.client.Do(req, &gb)
	if err != nil {
		return nil, resp, err
	}

	return gb, resp, err
}

func (s *GroupBadgesService) GetGroupBadge(gid interface{}, badge int, options ...OptionFunc) (*GroupBadge, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/badges/%d", url.QueryEscape(group), badge)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	gb := new(GroupBadge)
	resp, err := s.client.Do(req, gb)
	if err != nil {
		return nil, resp, err
	}

	return gb, resp, err
}

type AddGroupBadgeOptions struct {
	LinkURL  *string `url:"link_url,omitempty" json:"link_url,omitempty"`
	ImageURL *string `url:"image_url,omitempty" json:"image_url,omitempty"`
}

func (s *GroupBadgesService) AddGroupBadge(gid interface{}, opt *AddGroupBadgeOptions, options ...OptionFunc) (*GroupBadge, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/badges", url.QueryEscape(group))

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	gb := new(GroupBadge)
	resp, err := s.client.Do(req, gb)
	if err != nil {
		return nil, resp, err
	}

	return gb, resp, err
}

type EditGroupBadgeOptions struct {
	LinkURL  *string `url:"link_url,omitempty" json:"link_url,omitempty"`
	ImageURL *string `url:"image_url,omitempty" json:"image_url,omitempty"`
}

func (s *GroupBadgesService) EditGroupBadge(gid interface{}, badge int, opt *EditGroupBadgeOptions, options ...OptionFunc) (*GroupBadge, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/badges/%d", url.QueryEscape(group), badge)

	req, err := s.client.NewRequest("PUT", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	gb := new(GroupBadge)
	resp, err := s.client.Do(req, gb)
	if err != nil {
		return nil, resp, err
	}

	return gb, resp, err
}

func (s *GroupBadgesService) DeleteGroupBadge(gid interface{}, badge int, options ...OptionFunc) (*Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("group/%s/badges/%d", url.QueryEscape(group), badge)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

type GroupBadgePreviewOptions struct {
	LinkURL  *string `url:"link_url,omitempty" json:"link_url,omitempty"`
	ImageURL *string `url:"image_url,omitempty" json:"image_url,omitempty"`
}

func (s *GroupBadgesService) PreviewGroupBadge(gid interface{}, opt *GroupBadgePreviewOptions, options ...OptionFunc) (*GroupBadge, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/badges/render", url.QueryEscape(group))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	gb := new(GroupBadge)
	resp, err := s.client.Do(req, &gb)
	if err != nil {
		return nil, resp, err
	}

	return gb, resp, err
}
