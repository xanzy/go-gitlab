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

func TestListNamespaces(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

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
			}
		  ]`)
	})

	testCases := []struct {
		event      string
		search     *string
		owned_only *bool
	}{
		{"with_nothing", nil, nil},
		{"with_search", String("foobar"), nil},
		{"with_owned_only", nil, Bool(false)},
	}

	for _, tc := range testCases {
		t.Run(tc.event, func(t *testing.T) {
			namespaces, _, err := client.Namespaces.ListNamespaces(&ListNamespacesOptions{Search: tc.search, OwnedOnly: tc.owned_only})
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
					AvatarUrl:            String("https://secure.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon"),
					WebUrl:               "https://gitlab.example.com/user1",
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
					AvatarUrl:                   nil,
					WebUrl:                      "https://gitlab.example.com/groups/group1",
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
					AvatarUrl:                   nil,
					WebUrl:                      "https://gitlab.example.com/groups/foo/bar",
					MembersCountWithDescendants: 5,
					BillableMembersCount:        5,
					Plan:                        "default",
					TrialEndsOn:                 nil,
					Trial:                       false,
				},
			}

			if !reflect.DeepEqual(namespaces, want) {
				t.Errorf("Namespaces.ListNamespaces returned \ngot:\n%v\nwant:\n%v", Stringify(namespaces), Stringify(want))
			}
		})
	}
}
