package gitlab

import (
	"fmt"
	"net/url"
)

// CustomeAttributesService handles communication with the
// group, project and user custom attributes related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/custom_attributes.html
type CustomeAttributesService struct {
	client *Client
}

// CustomAttributes struct is used to unmarshal response to api calls
//
// GitLab API docs: https://docs.gitlab.com/ee/api/custom_attributes.html
type CustomAttributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// CustomAttributeOptions struct contains all options required to call the methods below
// agianst the gitlab api for customattributes
//
// GitLab API docs: https://docs.gitlab.com/ee/api/custom_attributes.html
type CustomAttributeOptions struct {
	CustomeAttributeResourceID interface{}
	CustomeAttributeResource   string
	CA                         CustomAttributes
}

// ListCustomAttributes lists the custom attributes of the specified resource
//
// GitLab API docs:  https://docs.gitlab.com/ee/api/custom_attributes.html#list-custom-attributes
func (s *CustomeAttributesService) ListCustomAttributes(opt CustomAttributeOptions, options ...OptionFunc) ([]*CustomAttributes, *Response, error) {
	resourceID, err := parseID(opt.CustomeAttributeResourceID)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("%s/%s/custom_attributes", opt.CustomeAttributeResource, url.QueryEscape(resourceID))
	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var ca []*CustomAttributes
	resp, err := s.client.Do(req, &ca)
	if err != nil {
		return nil, resp, err
	}

	return ca, resp, err

}

// GetCustomAttribute returns the attribute with a speciifc key
//
// GitLab API docs:  https://docs.gitlab.com/ee/api/custom_attributes.html#single-custom-attribute
func (s *CustomeAttributesService) GetCustomAttribute(opt CustomAttributeOptions, options ...OptionFunc) (*CustomAttributes, *Response, error) {
	resourceID, err := parseID(opt.CustomeAttributeResourceID)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("%s/%s/custom_attributes/%s", opt.CustomeAttributeResource, url.QueryEscape(resourceID), opt.CA.Key)
	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var ca *CustomAttributes
	resp, err := s.client.Do(req, &ca)
	if err != nil {
		return nil, resp, err
	}

	return ca, resp, err

}

// SetCustomAttribute sets the custom attributes of the specified resource
//
// GitLab API docs: https://docs.gitlab.com/ee/api/custom_attributes.html#set-custom-attribute
func (s *CustomeAttributesService) SetCustomAttribute(opt CustomAttributeOptions, options ...OptionFunc) (*CustomAttributes, *Response, error) {
	resourceID, err := parseID(opt.CustomeAttributeResourceID)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("%s/%s/custom_attributes/%s", opt.CustomeAttributeResource, url.QueryEscape(resourceID), opt.CA.Key)
	req, err := s.client.NewRequest("PUT", u, &opt.CA, options)
	if err != nil {
		return nil, nil, err
	}

	ca := new(CustomAttributes)
	resp, err := s.client.Do(req, ca)
	if err != nil {
		return nil, resp, err
	}

	return ca, resp, err
}

// DeleteCustomAttribute removes the customattribute fromt the specified resource
//
// GitLab API docs: https://docs.gitlab.com/ee/api/custom_attributes.html#delete-custom-attribute
func (s *CustomeAttributesService) DeleteCustomAttribute(opt CustomAttributeOptions, options ...OptionFunc) (*Response, error) {
	resourceID, err := parseID(opt.CustomeAttributeResourceID)
	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("%s/%s/custom_attributes/%s", opt.CustomeAttributeResource, url.QueryEscape(resourceID), opt.CA.Key)
	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
