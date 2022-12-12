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
	"time"
)

func TestListNamespaces(t *testing.T) {
	mux, client := setup(t)

	trialEndsOn, _ := time.Parse(time.RFC3339, "2022-05-08T00:00:00Z")
	trialEndsOnISOTime := ISOTime(trialEndsOn)

	mux.HandleFunc("/api/v4/namespaces", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			  "id": 1,
			  "name": "user1",
			  "path": "user1",
			  "kind": "user",
			  "full_path": "user1",
			  "avatar_url": "https://secure.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "https://gitlab.example.com/user1",
			  "billable_members_count": 1,
			  "plan": "default",
			  "trial_ends_on": null,
			  "trial": false
			},
			{
			  "id": 2,
			  "name": "group1",
			  "path": "group1",
			  "kind": "group",
			  "full_path": "group1",
			  "web_url": "https://gitlab.example.com/groups/group1",
			  "members_count_with_descendants": 2,
			  "billable_members_count": 2,
			  "plan": "default",
			  "trial_ends_on": null,
			  "trial": false
			},
			{
			  "id": 3,
			  "name": "bar",
			  "path": "bar",
			  "kind": "group",
			  "full_path": "foo/bar",
			  "parent_id": 9,
			  "web_url": "https://gitlab.example.com/groups/foo/bar",
			  "members_count_with_descendants": 5,
			  "billable_members_count": 5,
			  "plan": "default",
			  "trial_ends_on": null,
			  "trial": false
			},
			{
				"id": 4,
				"name": "group2",
				"path": "group2",
				"kind": "group",
				"full_path": "group2",
				"avatar_url": "https://gitlab.example.com/groups/group2",
				"web_url": "https://gitlab.example.com/group2",
				"billable_members_count": 1,
				"plan": "default",
				"trial_ends_on": "2022-05-08",
				"trial": true
			  }
		  ]`)
	})

	testCases := []struct {
		event     string
		search    *string
		ownedOnly *bool
	}{
		{"with_nothing", nil, nil},
		{"with_search", String("foobar"), nil},
		{"with_owned_only", nil, Bool(false)},
	}

	for _, tc := range testCases {
		t.Run(tc.event, func(t *testing.T) {
			namespaces, _, err := client.Namespaces.ListNamespaces(&ListNamespacesOptions{Search: tc.search, OwnedOnly: tc.ownedOnly})
			if err != nil {
				t.Errorf("Namespaces.ListNamespaces returned error: %v", err)
			}

			want := []*Namespace{
				{
					ID:                   1,
					Name:                 "user1",
					Path:                 "user1",
					Kind:                 "user",
					FullPath:             "user1",
					AvatarURL:            String("https://secure.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon"),
					WebURL:               "https://gitlab.example.com/user1",
					Plan:                 "default",
					BillableMembersCount: 1,
					TrialEndsOn:          nil,
					Trial:                false,
				},
				{
					ID:                          2,
					Name:                        "group1",
					Path:                        "group1",
					Kind:                        "group",
					FullPath:                    "group1",
					AvatarURL:                   nil,
					WebURL:                      "https://gitlab.example.com/groups/group1",
					MembersCountWithDescendants: 2,
					BillableMembersCount:        2,
					Plan:                        "default",
					TrialEndsOn:                 nil,
					Trial:                       false,
				},
				{
					ID:                          3,
					Name:                        "bar",
					Path:                        "bar",
					Kind:                        "group",
					FullPath:                    "foo/bar",
					ParentID:                    9,
					AvatarURL:                   nil,
					WebURL:                      "https://gitlab.example.com/groups/foo/bar",
					MembersCountWithDescendants: 5,
					BillableMembersCount:        5,
					Plan:                        "default",
					TrialEndsOn:                 nil,
					Trial:                       false,
				},
				{
					ID:                   4,
					Name:                 "group2",
					Path:                 "group2",
					Kind:                 "group",
					FullPath:             "group2",
					AvatarURL:            String("https://gitlab.example.com/groups/group2"),
					WebURL:               "https://gitlab.example.com/group2",
					Plan:                 "default",
					BillableMembersCount: 1,
					TrialEndsOn:          &trialEndsOnISOTime,
					Trial:                true,
				},
			}

			if !reflect.DeepEqual(namespaces, want) {
				t.Errorf("Namespaces.ListNamespaces returned \ngot:\n%v\nwant:\n%v", Stringify(namespaces), Stringify(want))
			}
		})
	}
}

func TestGetNamespace(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/namespaces/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
			"id": 2,
			"name": "group1",
			"path": "group1",
			"kind": "group",
			"full_path": "group1",
			"avatar_url": null,
			"web_url": "https://gitlab.example.com/groups/group1",
			"members_count_with_descendants": 2,
			"billable_members_count": 2,
			"max_seats_used": 0,
			"seats_in_use": 0,
			"plan": "default",
			"trial_ends_on": null,
			"trial": false
		  }`)
	})

	namespace, _, err := client.Namespaces.GetNamespace(2)
	if err != nil {
		t.Errorf("Namespaces.GetNamespace returned error: %v", err)
	}

	want := &Namespace{
		ID:                          2,
		Name:                        "group1",
		Path:                        "group1",
		Kind:                        "group",
		FullPath:                    "group1",
		AvatarURL:                   nil,
		WebURL:                      "https://gitlab.example.com/groups/group1",
		MembersCountWithDescendants: 2,
		BillableMembersCount:        2,
		MaxSeatsUsed:                Int(0),
		SeatsInUse:                  Int(0),
		Plan:                        "default",
		TrialEndsOn:                 nil,
		Trial:                       false,
	}

	if !reflect.DeepEqual(namespace, want) {
		t.Errorf("Namespaces.ListNamespaces returned \ngot:\n%v\nwant:\n%v", Stringify(namespace), Stringify(want))
	}
}

func TestNamespaceExists(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/namespaces/my-group/exists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
			"exists": true,
			"suggests": [
				"my-group1"
			]
		}`)
	})

	opt := &NamespaceExistsOptions{
		ParentID: Int(1),
	}
	exists, _, err := client.Namespaces.NamespaceExists("my-group", opt)
	if err != nil {
		t.Errorf("Namespaces.NamespaceExists returned error: %v", err)
	}

	want := &NamespaceExistance{
		Exists:   true,
		Suggests: []string{"my-group1"},
	}
	if !reflect.DeepEqual(exists, want) {
		t.Errorf("Namespaces.NamespaceExists returned \ngot:\n%v\nwant:\n%v", Stringify(exists), Stringify(want))
	}
}

func TestSearchNamespace(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/namespaces", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			  "id": 4,
			  "name": "twitter",
			  "path": "twitter",
			  "kind": "group",
			  "full_path": "twitter",
			  "avatar_url": null,
			  "web_url": "https://gitlab.example.com/groups/twitter",
			  "members_count_with_descendants": 2,
			  "billable_members_count": 2,
			  "max_seats_used": 0,
			  "seats_in_use": 0,
			  "plan": "default",
			  "trial_ends_on": null,
			  "trial": false
			}
		  ]`)
	})

	namespaces, _, err := client.Namespaces.SearchNamespace("?search=twitter")
	if err != nil {
		t.Errorf("Namespaces.SearchNamespaces returned error: %v", err)
	}

	want := []*Namespace{
		{
			ID:                          4,
			Name:                        "twitter",
			Path:                        "twitter",
			Kind:                        "group",
			FullPath:                    "twitter",
			AvatarURL:                   nil,
			WebURL:                      "https://gitlab.example.com/groups/twitter",
			MembersCountWithDescendants: 2,
			BillableMembersCount:        2,
			MaxSeatsUsed:                Int(0),
			SeatsInUse:                  Int(0),
			Plan:                        "default",
			TrialEndsOn:                 nil,
			Trial:                       false,
		},
	}
	if !reflect.DeepEqual(namespaces, want) {
		t.Errorf("Namespaces.SearchNamespaces returned \ngot:\n%v\nwant:\n%v", Stringify(namespaces), Stringify(want))
	}
}
