//
// Copyright 2021, Pavel Kostohrys
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

import (
	"net/http"
)

// AccessRequest represents a access request for a user avatars.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/avatar.html
type AvatarOptions struct {
	Email string `json:"email"`
	Size  int    `json:"size"`
}

// AccessRequestsService handles communication with the avatar
// avatar related methods of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/avatar.html
type AvatarRequestsService struct {
	client *Client
}

func (s *AvatarRequestsService) GetAvatar(opts *AvatarOptions, options ...RequestOptionFunc) (error, *string) {
	req, err := s.client.NewRequest(http.MethodGet, "avatar", opts, options)
	if err != nil {
		return err, nil
	}

	avatar := struct {
		Url string `json:"avatar_url"`
	}{}

	_, err = s.client.Do(req, &avatar)
	if err != nil {
		return err, nil
	}

	return nil, &avatar.Url
}
