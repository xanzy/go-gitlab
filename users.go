//
// Copyright 2015, Sander van Harmelen
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
	"fmt"
	"time"
)

// UsersService handles communication with the user related methods of
// the GitLab API.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html
type UsersService struct {
	client *Client
}

// User represents a GitLab user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html
type User struct {
	ID               int       `json:"id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	Name             string    `json:"name"`
	State            string    `json:"state"`
	CreatedAt        time.Time `json:"created_at"`
	Bio              string    `json:"bio"`
	Skype            string    `json:"skype"`
	Linkedin         string    `json:"linkedin"`
	Twitter          string    `json:"twitter"`
	WebsiteURL       string    `json:"website_url"`
	ExternUID        string    `json:"extern_uid"`
	Provider         string    `json:"provider"`
	ThemeID          int       `json:"theme_id"`
	ColorSchemeID    int       `json:"color_scheme_id"`
	IsAdmin          bool      `json:"is_admin"`
	AvatarURL        string    `json:"avatar_url"`
	CanCreateGroup   bool      `json:"can_create_group"`
	CanCreateProject bool      `json:"can_create_project"`
	ProjectsLimit    int       `json:"projects_limit"`
	CurrentSignInAt  time.Time `json:"current_sign_in_at"`
	TwoFactorEnabled bool      `json:"two_factor_enabled"`
}

// ListUsers gets a list of users.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#list-users
func (s *UsersService) ListUsers() ([]*User, *Response, error) {
	req, err := s.client.NewRequest("GET", "users", nil)
	if err != nil {
		return nil, nil, err
	}

	var usr []*User
	resp, err := s.client.Do(req, &usr)
	if err != nil {
		return nil, resp, err
	}

	return usr, resp, err
}

// GetUser gets a single user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#single-user
func (s *UsersService) GetUser(user int) (*User, *Response, error) {
	u := fmt.Sprintf("users/%d", user)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	usr := new(User)
	resp, err := s.client.Do(req, usr)
	if err != nil {
		return nil, resp, err
	}

	return usr, resp, err
}

// CreateUserOptions represents the available CreateUser() options.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#user-creation
type CreateUserOptions struct {
	Email          string `url:"email,omitempty"`
	Password       string `url:"password,omitempty"`
	Username       string `url:"username,omitempty"`
	Name           string `url:"name,omitempty"`
	Skype          string `url:"skype,omitempty"`
	Linkedin       string `url:"linkedin,omitempty"`
	Twitter        string `url:"twitter,omitempty"`
	WebsiteURL     string `url:"website_url,omitempty"`
	ProjectsLimit  int    `url:"projects_limit,omitempty"`
	ExternUID      string `url:"extern_uid,omitempty"`
	Provider       string `url:"provider,omitempty"`
	Bio            string `url:"bio,omitempty"`
	Admin          bool   `url:"admin,omitempty"`
	CanCreateGroup bool   `url:"can_create_group,omitempty"`
	Confirm        bool   `url:"confirm,omitempty"`
}

// CreateUser creates a new user. Note only administrators can create new users.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#user-creation
func (s *UsersService) CreateUser(opt *CreateUserOptions) (*User, *Response, error) {
	req, err := s.client.NewRequest("POST", "users", opt)
	if err != nil {
		return nil, nil, err
	}

	usr := new(User)
	resp, err := s.client.Do(req, usr)
	if err != nil {
		return nil, resp, err
	}

	return usr, resp, err
}

// ModifyUserOptions represents the available ModifyUser() options.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#user-modification
type ModifyUserOptions struct {
	Email          string `url:"email,omitempty"`
	Password       string `url:"password,omitempty"`
	Username       string `url:"username,omitempty"`
	Name           string `url:"name,omitempty"`
	Skype          string `url:"skype,omitempty"`
	Linkedin       string `url:"linkedin,omitempty"`
	Twitter        string `url:"twitter,omitempty"`
	WebsiteURL     string `url:"website_url,omitempty"`
	ProjectsLimit  int    `url:"projects_limit,omitempty"`
	ExternUID      string `url:"extern_uid,omitempty"`
	Provider       string `url:"provider,omitempty"`
	Bio            string `url:"bio,omitempty"`
	Admin          bool   `url:"admin,omitempty"`
	CanCreateGroup bool   `url:"can_create_group,omitempty"`
}

// ModifyUser modifies an existing user. Only administrators can change attributes
// of a user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#user-modification
func (s *UsersService) ModifyUser(user int, opt *ModifyUserOptions) (*User, *Response, error) {
	u := fmt.Sprintf("users/%d", user)

	req, err := s.client.NewRequest("PUT", u, opt)
	if err != nil {
		return nil, nil, err
	}

	usr := new(User)
	resp, err := s.client.Do(req, usr)
	if err != nil {
		return nil, resp, err
	}

	return usr, resp, err
}

// DeleteUser deletes a user. Available only for administrators. This is an
// idempotent function, calling this function for a non-existent user id still
// returns a status code 200 OK. The JSON response differs if the user was
// actually deleted or not. In the former the user is returned and in the
// latter not.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#user-deletion
func (s *UsersService) DeleteUser(user int) (*Response, error) {
	u := fmt.Sprintf("users/%d", user)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// CurrentUser gets currently authenticated user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#current-user
func (s *UsersService) CurrentUser() (*User, *Response, error) {
	req, err := s.client.NewRequest("GET", "user", nil)
	if err != nil {
		return nil, nil, err
	}

	usr := new(User)
	resp, err := s.client.Do(req, usr)
	if err != nil {
		return nil, resp, err
	}

	return usr, resp, err
}

// SSHKey represents a SSH key.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#list-ssh-keys
type SSHKey struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
}

// ListSSHKeys gets a list of currently authenticated user's SSH keys.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#list-ssh-keys
func (s *UsersService) ListSSHKeys() ([]*SSHKey, *Response, error) {
	req, err := s.client.NewRequest("GET", "user/keys", nil)
	if err != nil {
		return nil, nil, err
	}

	var k []*SSHKey
	resp, err := s.client.Do(req, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err
}

// ListSSHKeysForUser gets a list of a specified user's SSH keys. Available
// only for admin
//
// GitLab API docs:
// http://doc.gitlab.com/ce/api/users.html#list-ssh-keys-for-user
func (s *UsersService) ListSSHKeysForUser(user int) ([]*SSHKey, *Response, error) {
	u := fmt.Sprintf("users/%d/keys", user)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var k []*SSHKey
	resp, err := s.client.Do(req, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err
}

// GetSSHKey gets a single key.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#single-ssh-key
func (s *UsersService) GetSSHKey(kid int) (*SSHKey, *Response, error) {
	u := fmt.Sprintf("user/keys/%d", kid)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	k := new(SSHKey)
	resp, err := s.client.Do(req, k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err
}

// AddSSHKeyOptions represents the available AddSSHKey() options.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/projects.html#add-ssh-key
type AddSSHKeyOptions struct {
	Title string `url:"title,omitempty"`
	Key   string `url:"key,omitempty"`
}

// AddSSHKey creates a new key owned by the currently authenticated user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#add-ssh-key
func (s *UsersService) AddSSHKey(opt *AddSSHKeyOptions) (*SSHKey, *Response, error) {
	req, err := s.client.NewRequest("POST", "user/keys", opt)
	if err != nil {
		return nil, nil, err
	}

	k := new(SSHKey)
	resp, err := s.client.Do(req, k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err
}

// AddSSHKeyForUser creates new key owned by specified user. Available only for
// admin.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#add-ssh-key-for-user
func (s *UsersService) AddSSHKeyForUser(
	user int,
	opt *AddSSHKeyOptions) (*SSHKey, *Response, error) {
	u := fmt.Sprintf("users/%d/keys", user)

	req, err := s.client.NewRequest("POST", u, opt)
	if err != nil {
		return nil, nil, err
	}

	k := new(SSHKey)
	resp, err := s.client.Do(req, k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err
}

// DeleteSSHKey deletes key owned by currently authenticated user. This is an
// idempotent function and calling it on a key that is already deleted or not
// available results in 200 OK.
//
// GitLab API docs:
// http://doc.gitlab.com/ce/api/users.html#delete-ssh-key-for-current-owner
func (s *UsersService) DeleteSSHKey(kid int) (*Response, error) {
	u := fmt.Sprintf("user/keys/%d", kid)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// DeleteSSHKeyForUser deletes key owned by a specified user. Available only
// for admin.
//
// GitLab API docs:
// http://doc.gitlab.com/ce/api/users.html#delete-ssh-key-for-given-user
func (s *UsersService) DeleteSSHKeyForUser(user int, kid int) (*Response, error) {
	u := fmt.Sprintf("users/%d/keys/%d", user, kid)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// BlockUser blocks the specified user. Available only for admin.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#block-user
func (s *UsersService) BlockUser(user int) (*User, *Response, error) {
	u := fmt.Sprintf("users/%d/block", user)

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, nil, err
	}

	usr := new(User)
	resp, err := s.client.Do(req, usr)
	if err != nil {
		return nil, resp, err
	}

	return usr, resp, err
}

// UnblockUser unblocks the specified user. Available only for admin.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#unblock-user
func (s *UsersService) UnblockUser(user int) (*User, *Response, error) {
	u := fmt.Sprintf("users/%d/unblock", user)

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, nil, err
	}

	usr := new(User)
	resp, err := s.client.Do(req, usr)
	if err != nil {
		return nil, resp, err
	}

	return usr, resp, err
}
