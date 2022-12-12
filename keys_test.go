//
// Copyright 2021, Sander van Harmelen
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
	"net/http"
	"reflect"
	"testing"
)

func TestGetKeyWithUser(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/keys/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `{
			  "id": 1,
			  "title": "Sample key 25",
			  "key": "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAIEAiPWx6WM4lhHNedGfBpPJNPpZ7yKu+dnn1SJejgt1256k6YjzGGphH2TUxwKzxcKDKKezwkpfnxPkSMkuEspGRt/aZZ9wa++Oi7Qkr8prgHc4soW6NUlfDzpvZK2H5E7eQaSeP3SAwGmQKUFHCddNaP0L+hM7zhFNzjFvpaMgJw0=",
			  "user": {
			    "id": 25,
			    "username": "john_smith",
			    "name": "John Smith",
			    "email": "john@example.com",
			    "state": "active",
			    "bio": null,
			    "location": null,
			    "skype": "",
			    "linkedin": "",
			    "twitter": "",
			    "website_url": "http://localhost:3000/john_smith",
			    "organization": null,
			    "theme_id": 2,
			    "color_scheme_id": 1,
			    "avatar_url": "http://www.gravatar.com/avatar/cfa35b8cd2ec278026357769582fa563?s=40\u0026d=identicon",
			    "can_create_group": true,
			    "can_create_project": true,
			    "projects_limit": 10,
			    "two_factor_enabled": false,
			    "identities": [],
			    "external": false,
			    "public_email": "john@example.com"
			  }
			}`)
		})

	key, _, err := client.Keys.GetKeyWithUser(1)
	if err != nil {
		t.Errorf("Keys.GetKeyWithUser returned error: %v", err)
	}

	want := &Key{
		ID:    1,
		Title: "Sample key 25",
		Key:   "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAIEAiPWx6WM4lhHNedGfBpPJNPpZ7yKu+dnn1SJejgt1256k6YjzGGphH2TUxwKzxcKDKKezwkpfnxPkSMkuEspGRt/aZZ9wa++Oi7Qkr8prgHc4soW6NUlfDzpvZK2H5E7eQaSeP3SAwGmQKUFHCddNaP0L+hM7zhFNzjFvpaMgJw0=",
		User: User{
			ID:               25,
			Username:         "john_smith",
			Email:            "john@example.com",
			Name:             exampleEventUserName,
			State:            "active",
			Bio:              "",
			Location:         "",
			Skype:            "",
			Linkedin:         "",
			Twitter:          "",
			WebsiteURL:       "http://localhost:3000/john_smith",
			Organization:     "",
			ThemeID:          2,
			ColorSchemeID:    1,
			AvatarURL:        "http://www.gravatar.com/avatar/cfa35b8cd2ec278026357769582fa563?s=40\u0026d=identicon",
			CanCreateGroup:   true,
			CanCreateProject: true,
			ProjectsLimit:    10,
			TwoFactorEnabled: false,
			Identities:       []*UserIdentity{},
			External:         false,
			PublicEmail:      "john@example.com",
		},
	}

	if !reflect.DeepEqual(want, key) {
		t.Errorf("Keys.GetKeyWithUser returned %+v, want %+v", key, want)
	}
}
