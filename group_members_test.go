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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListBillableGroupMembers(t *testing.T) {
	mux, client := setup(t)

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
						"last_activity_on":"2021-01-27",
						"membership_type": "group_member",
						"removable": true,
						"created_at": "2017-10-23T11:41:28.793Z",
						"is_last_owner": false,
						"last_login_at": "2022-12-12T09:22:51.581Z"
					}
				]`)
		})

	billableMembers, _, err := client.Groups.ListBillableGroupMembers(1, &ListBillableGroupMembersOptions{})
	if err != nil {
		t.Errorf("Groups.ListBillableGroupMembers returned error: %v", err)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2017-10-23T11:41:28.793Z")
	lastLoginAt, _ := time.Parse(time.RFC3339, "2022-12-12T09:22:51.581Z")
	lastActivityOn, _ := time.Parse(time.RFC3339, "2021-01-27T00:00:00Z")
	lastActivityOnISOTime := ISOTime(lastActivityOn)

	want := []*BillableGroupMember{
		{
			ID:             1,
			Username:       "ray",
			Name:           "Raymond",
			State:          "active",
			AvatarURL:      "https://foo.bar/mypic",
			WebURL:         "http://192.168.1.8:3000/root",
			LastActivityOn: &lastActivityOnISOTime,
			MembershipType: "group_member",
			Removable:      true,
			CreatedAt:      &createdAt,
			IsLastOwner:    false,
			LastLoginAt:    &lastLoginAt,
		},
	}
	assert.Equal(t, want, billableMembers, "Expected returned Groups.ListBillableGroupMembers to equal")
}

func TestListMembershipsForBillableGroupMember(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/groups/1/billable_members/42/memberships",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w,
				`[
					{
						"id":21,
						"source_id":36,
						"source_full_name":"Root Group / Test Group",
						"source_members_url":"https://gitlab.example.com/groups/root-group/test-group/-/group_members",
						"created_at":"2021-03-31T17:28:44.812Z",
						"access_level": {
							"string_value": "Developer",
							"integer_value": 30
						}
					}
				]`)
		})

	memberships, _, err := client.Groups.ListMembershipsForBillableGroupMember(1, 42, &ListOptions{})
	if err != nil {
		t.Errorf("Groups.ListMembershipsForBillableGroupMember returned error: %v", err)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2021-03-31T17:28:44.812Z")

	want := []*BillableUserMembership{
		{
			ID:               21,
			SourceID:         36,
			SourceFullName:   "Root Group / Test Group",
			SourceMembersURL: "https://gitlab.example.com/groups/root-group/test-group/-/group_members",
			CreatedAt:        &createdAt,
			AccessLevel: &AccessLevelDetails{
				IntegerValue: 30,
				StringValue:  "Developer",
			},
		},
	}
	assert.Equal(t, want, memberships, "Expected returned Groups.ListMembershipsForBillableGroupMember to equal")
}

func TestListGroupMembersWithoutEmail(t *testing.T) {
	mux, client := setup(t)

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
			ID:          1,
			Username:    "raymond_smith",
			Name:        "Raymond Smith",
			State:       "active",
			AvatarURL:   "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
			WebURL:      "http://192.168.1.8:3000/root",
			CreatedAt:   &createdAt,
			ExpiresAt:   &expiresAtISOTime,
			AccessLevel: 30,
		},
	}
	if !reflect.DeepEqual(want, members) {
		t.Errorf("Groups.ListBillableGroupMembers returned %+v, want %+v", members[0], want[0])
	}
}

func TestListGroupMembersWithEmail(t *testing.T) {
	mux, client := setup(t)

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
						"email": "john@example.com",
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
			ID:          1,
			Username:    "raymond_smith",
			Name:        "Raymond Smith",
			State:       "active",
			AvatarURL:   "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
			WebURL:      "http://192.168.1.8:3000/root",
			CreatedAt:   &createdAt,
			ExpiresAt:   &expiresAtISOTime,
			AccessLevel: 30,
			Email:       "john@example.com",
		},
	}
	if !reflect.DeepEqual(want, members) {
		t.Errorf("Groups.ListBillableGroupMembers returned %+v, want %+v", members[0], want[0])
	}
}

func TestListGroupMembersWithoutSAML(t *testing.T) {
	mux, client := setup(t)

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
	mux, client := setup(t)

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

func TestGetGroupMemberCustomRole(t *testing.T) {
	mux, client := setup(t)

	path := fmt.Sprintf("/%sgroups/1/members/2", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		// This is pulled straight from a `/group/<group_id>/members/<user_id>` call, then obfuscated.
		fmt.Fprint(w, `
		{
			"id":1,
			"username":"test",
			"name":"testName",
			"access_level":30,
			"member_role":{
				"id":1,
				"group_id":2,
				"name":"TestingCustomRole",
				"description":"",
				"base_access_level":30,
				"admin_cicd_variables":true,
				"admin_group_member":null,
				"admin_merge_request":null,
				"admin_push_rules":null,
				"admin_terraform_state":null,
				"admin_vulnerability":null,
				"archive_project":null,
				"manage_group_access_tokens":null,
				"manage_project_access_tokens":null,
				"read_code":null,
				"read_dependency":null,
				"read_vulnerability":null,
				"remove_group":null,
				"remove_project":null
			}
		}
		`)
	})

	want := &GroupMember{
		ID:          1,
		Username:    "test",
		Name:        "testName",
		AccessLevel: AccessLevelValue(30),
		MemberRole: &MemberRole{
			ID:                 1,
			GroupID:            2,
			Name:               "TestingCustomRole",
			Description:        "",
			BaseAccessLevel:    AccessLevelValue(30),
			AdminCICDVariables: true,
		},
	}
	member, _, err := client.GroupMembers.GetGroupMember(1, 2)

	assert.NoError(t, err)
	assert.Equal(t, want, member)
}

func TestGetGroupMemberAll(t *testing.T) {
	mux, client := setup(t)

	path := fmt.Sprintf("/%sgroups/1/members/all/2", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		fmt.Fprint(w, `
		{
		  "id": 2,
		  "name": "aaa",
		  "username": "aaaName",
		  "state": "active",
		  "avatar_url": "https://secure.gravatar.com/avatar/e547676d82f1e16954b2280a5b4cbe79?s=80&d=identicon",
		  "web_url": "https://gitlab.example.cn/aaa",
		  "access_level": 30,
		  "created_at": "2024-06-19T07:14:02.793Z",
		  "expires_at": null
		}
		`)
	})

	createAt, _ := time.Parse(time.RFC3339, "2024-06-19T07:14:02.793Z")

	want := &GroupMember{
		ID:          2,
		Name:        "aaa",
		Username:    "aaaName",
		State:       "active",
		AvatarURL:   "https://secure.gravatar.com/avatar/e547676d82f1e16954b2280a5b4cbe79?s=80&d=identicon",
		WebURL:      "https://gitlab.example.cn/aaa",
		AccessLevel: AccessLevelValue(30),
		CreatedAt:   &createAt,
	}

	pm, resp, err := client.GroupMembers.GetInheritedGroupMember(1, 2, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pm)

	member, resp, err := client.GroupMembers.GetInheritedGroupMember(1.01, 2, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, member)

	member, resp, err = client.GroupMembers.GetInheritedGroupMember(1, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, member)

	member, resp, err = client.GroupMembers.GetInheritedGroupMember(2, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, member)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
