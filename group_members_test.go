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

func TestListBillableGroupMembers(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/1/billable_members",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w,
				`[
					{
						"id":1,
						"username":"ray",
						"name":"Raymond",
						"state":"active",
						"avatar_url":"https://foo.bar/mypic",
						"web_url":"http://192.168.1.8:3000/root",
						"last_activity_on":"2021-01-27"
					}
				]`)
		})

	billableMembers, _, err := client.Groups.ListBillableGroupMembers(1, &ListBillableGroupMembersOptions{})
	if err != nil {
		t.Errorf("Groups.ListBillableGroupMembers returned error: %v", err)
	}

	testTime := ISOTime{}
	err = testTime.UnmarshalJSON([]byte(`"2021-01-27"`))
	if err != nil {
		t.Errorf("Could not unmarshal date string to ISOTime: %v", err)
	}
	want := []*BillableGroupMember{
		{
			ID:             1,
			Username:       "ray",
			Name:           "Raymond",
			State:          "active",
			AvatarURL:      "https://foo.bar/mypic",
			WebURL:         "http://192.168.1.8:3000/root",
			LastActivityOn: testTime,
		},
	}
	if !reflect.DeepEqual(want, billableMembers) {
		t.Errorf("Groups.ListBillableGroupMembers returned %+v, want %+v", billableMembers, want)
	}
}

func TestListGroupMembersWithoutSAML(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/1/members",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w,
				`[
					{
						"id": 1,
						"username": "raymond_smith",
						"name": "Raymond Smith",
						"state": "active",
						"avatar_url": "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
						"web_url": "http://192.168.1.8:3000/root",
						"created_at": "2012-10-21T14:13:35Z",
						"expires_at": "2012-10-22",
						"access_level": 30,
						"group_saml_identity": null
					}
				]`)
		})

	members, _, err := client.Groups.ListGroupMembers(1, &ListGroupMembersOptions{})
	if err != nil {
		t.Errorf("Groups.ListGroupMembers returned error: %v", err)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2012-10-21T14:13:35Z")
	expiresAt, _ := time.Parse(time.RFC3339, "2012-10-22T00:00:00Z")
	expiresAtISOTime := ISOTime(expiresAt)
	want := []*GroupMember{
		{
			ID:                1,
			Username:          "raymond_smith",
			Name:              "Raymond Smith",
			State:             "active",
			AvatarURL:         "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
			WebURL:            "http://192.168.1.8:3000/root",
			CreatedAt:         &createdAt,
			ExpiresAt:         &expiresAtISOTime,
			AccessLevel:       30,
			GroupSAMLIdentity: nil,
		},
	}
	if !reflect.DeepEqual(want, members) {
		t.Errorf("Groups.ListBillableGroupMembers returned %+v, want %+v", members[0], want[0])
	}
}

func TestListGroupMembersWithSAML(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/1/members",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w,
				`[
					{
						"id": 2,
						"username": "john_doe",
						"name": "John Doe",
						"state": "active",
						"avatar_url": "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
						"web_url": "http://192.168.1.8:3000/root",
						"created_at": "2012-10-21T14:13:35Z",
						"expires_at": "2012-10-22",
						"access_level": 30,
						"group_saml_identity": {
							"extern_uid":"ABC-1234567890",
							"provider": "group_saml",
							"saml_provider_id": 10
						}
					}
				]`)
		})

	members, _, err := client.Groups.ListGroupMembers(1, &ListGroupMembersOptions{})
	if err != nil {
		t.Errorf("Groups.ListGroupMembers returned error: %v", err)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2012-10-21T14:13:35Z")
	expiresAt, _ := time.Parse(time.RFC3339, "2012-10-22T00:00:00Z")
	expiresAtISOTime := ISOTime(expiresAt)
	want := []*GroupMember{
		{
			ID:          2,
			Username:    "john_doe",
			Name:        "John Doe",
			State:       "active",
			AvatarURL:   "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
			WebURL:      "http://192.168.1.8:3000/root",
			CreatedAt:   &createdAt,
			ExpiresAt:   &expiresAtISOTime,
			AccessLevel: 30,
			GroupSAMLIdentity: &GroupMemberSAMLIdentity{
				ExternUID:      "ABC-1234567890",
				Provider:       "group_saml",
				SAMLProviderID: 10,
			},
		},
	}
	if !reflect.DeepEqual(want, members) {
		t.Errorf("Groups.ListBillableGroupMembers returned %+v, want %+v", members[0], want[0])
	}
}
