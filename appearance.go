//
// Copyright 2023, 徐晓伟 <xuxiaowei@xuxiaowei.com.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import "net/http"

// AppearanceService handles communication with appearance
// of the Gitlab API.
//
// Gitlab API docs : https://docs.gitlab.com/ee/api/appearance.html
type AppearanceService struct {
	client *Client
}

// Appearance represents a GitLab appearance
type Appearance struct {
	Title                       string `json:"title"`
	Description                 string `json:"description"`
	PwaName                     string `json:"pwa_name"`
	PwaShortName                string `json:"pwa_short_name"`
	PwaDescription              string `json:"pwa_description"`
	PwaIcon                     string `json:"pwa_icon"`
	Logo                        string `json:"logo"`
	HeaderLogo                  string `json:"header_logo"`
	Favicon                     string `json:"favicon"`
	NewProjectGuidelines        string `json:"new_project_guidelines"`
	ProfileImageGuidelines      string `json:"profile_image_guidelines"`
	HeaderMessage               string `json:"header_message"`
	FooterMessage               string `json:"footer_message"`
	MessageBackgroundColor      string `json:"message_background_color"`
	MessageFontColor            string `json:"message_font_color"`
	EmailHeaderAndFooterEnabled bool   `json:"email_header_and_footer_enabled"`
}

// GetAppearance
//
// Gitlab API docs : https://docs.gitlab.com/ee/api/appearance.html#get-current-appearance-configuration
func (s *AppearanceService) GetAppearance(options ...RequestOptionFunc) (*Appearance, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/application/appearance", nil, options)

	if err != nil {
		return nil, nil, err
	}

	var as *Appearance
	resp, err := s.client.Do(req, &as)
	if err != nil {
		return nil, resp, err
	}

	return as, resp, nil
}
