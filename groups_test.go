package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListGroups(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `[{"id":1},{"id":2}]`)
		})

	groups, _, err := client.Groups.ListGroups(&ListGroupsOptions{})
	if err != nil {
		t.Errorf("Groups.ListGroups returned error: %v", err)
	}

	want := []*Group{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, groups) {
		t.Errorf("Groups.ListGroups returned %+v, want %+v", groups, want)
	}
}

func TestGetGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/g",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `{"id": 1, "name": "g", "default_branch": "branch"}`)
		})

	group, _, err := client.Groups.GetGroup("g", &GetGroupOptions{})
	if err != nil {
		t.Errorf("Groups.GetGroup returned error: %v", err)
	}

	want := &Group{ID: 1, Name: "g", DefaultBranch: "branch"}
	if !reflect.DeepEqual(want, group) {
		t.Errorf("Groups.GetGroup returned %+v, want %+v", group, want)
	}
}

func TestGetGroupWithFileTemplateId(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/g",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `{"id": 1, "name": "g","file_template_project_id": 12345}`)
		})

	group, _, err := client.Groups.GetGroup("g", &GetGroupOptions{})
	if err != nil {
		t.Errorf("Groups.GetGroup returned error: %v", err)
	}

	want := &Group{ID: 1, Name: "g", FileTemplateProjectID: 12345}
	if !reflect.DeepEqual(want, group) {
		t.Errorf("Groups.GetGroup returned %+v, want %+v", group, want)
	}
}

func TestCreateGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `{"id": 1, "name": "g", "path": "g"}`)
		})

	opt := &CreateGroupOptions{
		Name: Ptr("g"),
		Path: Ptr("g"),
	}

	group, _, err := client.Groups.CreateGroup(opt, nil)
	if err != nil {
		t.Errorf("Groups.CreateGroup returned error: %v", err)
	}

	want := &Group{ID: 1, Name: "g", Path: "g"}
	if !reflect.DeepEqual(want, group) {
		t.Errorf("Groups.CreateGroup returned %+v, want %+v", group, want)
	}
}

func TestCreateGroupWithDefaultBranch(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `{"id": 1, "name": "g", "path": "g", "default_branch": "branch"}`)
		})

	opt := &CreateGroupOptions{
		Name:          Ptr("g"),
		Path:          Ptr("g"),
		DefaultBranch: Ptr("branch"),
	}

	group, _, err := client.Groups.CreateGroup(opt, nil)
	if err != nil {
		t.Errorf("Groups.CreateGroup returned error: %v", err)
	}

	want := &Group{ID: 1, Name: "g", Path: "g", DefaultBranch: "branch"}
	if !reflect.DeepEqual(want, group) {
		t.Errorf("Groups.CreateGroup returned %+v, want %+v", group, want)
	}
}

func TestCreateGroupDefaultBranchSettings(t *testing.T) {
	mux, client := setup(t)

	var jsonRequestBody CreateGroupOptions
	mux.HandleFunc("/api/v4/groups",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)

			testData, _ := io.ReadAll(r.Body)
			err := json.Unmarshal(testData, &jsonRequestBody)
			if err != nil {
				t.Fatal("Failed to unmarshal request body into an interface.")
			}

			fmt.Fprint(w, `
			{
				"id": 1,
				"name": "g",
				"path": "g",
				"default_branch_protection_defaults": {
					"allowed_to_push": [
						{
							"access_level": 40
						}
					],
					"allow_force_push": false,
					"allowed_to_merge": [
						{
							"access_level": 40
						}
					]
				}
			}
			`)
		})

	opt := &CreateGroupOptions{
		Name: Ptr("g"),
		Path: Ptr("g"),
		DefaultBranchProtectionDefaults: &DefaultBranchProtectionDefaultsOptions{
			AllowedToPush: &[]*GroupAccessLevel{
				{
					AccessLevel: Ptr(AccessLevelValue(40)),
				},
			},
			AllowedToMerge: &[]*GroupAccessLevel{
				{
					AccessLevel: Ptr(AccessLevelValue(40)),
				},
			},
		},
	}

	group, _, err := client.Groups.CreateGroup(opt, nil)
	if err != nil {
		t.Errorf("Groups.CreateGroup returned error: %v", err)
	}

	// Create the group that we want to get back
	want := &Group{
		ID:   1,
		Name: "g",
		Path: "g",
		DefaultBranchProtectionDefaults: &BranchProtectionDefaults{
			AllowedToMerge: []*GroupAccessLevel{
				{
					AccessLevel: Ptr(MaintainerPermissions),
				},
			},
			AllowedToPush: []*GroupAccessLevel{
				{
					AccessLevel: Ptr(MaintainerPermissions),
				},
			},
		},
	}

	if !reflect.DeepEqual(want, group) {
		t.Errorf("Groups.CreateGroup returned %+v, want %+v", group, want)
	}

	// Validate the request does what we want it to
	allowedToMerge := *jsonRequestBody.DefaultBranchProtectionDefaults.AllowedToMerge
	allowedToPush := *jsonRequestBody.DefaultBranchProtectionDefaults.AllowedToPush
	assert.Equal(t, Ptr(MaintainerPermissions), allowedToMerge[0].AccessLevel)
	assert.Equal(t, Ptr(MaintainerPermissions), allowedToPush[0].AccessLevel)
}

func TestTransferGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/projects/2",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprintf(w, `{"id": 1}`)
		})

	group, _, err := client.Groups.TransferGroup(1, 2)
	if err != nil {
		t.Errorf("Groups.TransferGroup returned error: %v", err)
	}

	want := &Group{ID: 1}
	if !reflect.DeepEqual(group, want) {
		t.Errorf("Groups.TransferGroup returned %+v, want %+v", group, want)
	}
}

func TestTransferSubGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/transfer",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprintf(w, `{"id": 1, "parent_id": 2}`)
		})

	opt := &TransferSubGroupOptions{
		GroupID: Ptr(2),
	}

	group, _, err := client.Groups.TransferSubGroup(1, opt)
	if err != nil {
		t.Errorf("Groups.TransferSubGroup returned error: %v", err)
	}

	want := &Group{ID: 1, ParentID: 2}
	if !reflect.DeepEqual(group, want) {
		t.Errorf("Groups.TransferSubGroup returned %+v, want %+v", group, want)
	}
}

func TestDeleteGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			w.WriteHeader(http.StatusAccepted)
		})

	resp, err := client.Groups.DeleteGroup(1, nil)
	if err != nil {
		t.Errorf("Groups.DeleteGroup returned error: %v", err)
	}

	want := http.StatusAccepted
	got := resp.StatusCode
	if got != want {
		t.Errorf("Groups.DeleteGroup returned %d, want %d", got, want)
	}
}

func TestDeleteGroup_WithPermanentDelete(t *testing.T) {
	mux, client := setup(t)
	var params url.Values

	mux.HandleFunc("/api/v4/groups/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			w.WriteHeader(http.StatusAccepted)

			// Get the request parameters
			parsedParams, err := url.ParseQuery(r.URL.RawQuery)
			if err != nil {
				t.Errorf("Groups.DeleteGroup returned error when parsing test parameters: %v", err)
			}
			params = parsedParams
		})

	resp, err := client.Groups.DeleteGroup(1, &DeleteGroupOptions{
		PermanentlyRemove: Ptr(true),
		FullPath:          Ptr("testPath"),
	})
	if err != nil {
		t.Errorf("Groups.DeleteGroup returned error: %v", err)
	}

	// Test that our status code matches
	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Groups.DeleteGroup returned %d, want %d", resp.StatusCode, http.StatusAccepted)
	}

	// Test that "permanently_remove" is set to true
	if params.Get("permanently_remove") != "true" {
		t.Errorf("Groups.DeleteGroup returned %v, want %v", params.Get("permanently_remove"), true)
	}

	// Test that "full_path" is set to "testPath"
	if params.Get("full_path") != "testPath" {
		t.Errorf("Groups.DeleteGroup returned %v, want %v", params.Get("full_path"), "testPath")
	}
}

func TestSearchGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `[{"id": 1, "name": "Foobar Group"}]`)
		})

	groups, _, err := client.Groups.SearchGroup("foobar")
	if err != nil {
		t.Errorf("Groups.SearchGroup returned error: %v", err)
	}

	want := []*Group{{ID: 1, Name: "Foobar Group"}}
	if !reflect.DeepEqual(want, groups) {
		t.Errorf("Groups.SearchGroup returned +%v, want %+v", groups, want)
	}
}

func TestUpdateGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			fmt.Fprint(w, `{"id": 1}`)
		})

	group, _, err := client.Groups.UpdateGroup(1, &UpdateGroupOptions{})
	if err != nil {
		t.Errorf("Groups.UpdateGroup returned error: %v", err)
	}

	want := &Group{ID: 1}
	if !reflect.DeepEqual(want, group) {
		t.Errorf("Groups.UpdatedGroup returned %+v, want %+v", group, want)
	}
}

func TestUpdateGroupWithDefaultBranch(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			fmt.Fprint(w, `{"id": 1, "default_branch": "branch"}`)
		})

	opt := &UpdateGroupOptions{
		DefaultBranch: Ptr("branch"),
	}

	group, _, err := client.Groups.UpdateGroup(1, opt)
	if err != nil {
		t.Errorf("Groups.UpdateGroup returned error: %v", err)
	}

	want := &Group{ID: 1, DefaultBranch: "branch"}
	if !reflect.DeepEqual(want, group) {
		t.Errorf("Groups.UpdatedGroup returned %+v, want %+v", group, want)
	}
}

func TestListGroupProjects(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/22/projects",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `[{"id":1},{"id":2}]`)
		})

	projects, _, err := client.Groups.ListGroupProjects(22,
		&ListGroupProjectsOptions{})
	if err != nil {
		t.Errorf("Groups.ListGroupProjects returned error: %v", err)
	}

	want := []*Project{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Groups.ListGroupProjects returned %+v, want %+v", projects, want)
	}
}

func TestListSubGroups(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/subgroups",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `[{"id": 1}, {"id": 2}]`)
		})

	groups, _, err := client.Groups.ListSubGroups(1, &ListSubGroupsOptions{})
	if err != nil {
		t.Errorf("Groups.ListSubGroups returned error: %v", err)
	}

	want := []*Group{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, groups) {
		t.Errorf("Groups.ListSubGroups returned %+v, want %+v", groups, want)
	}
}

func TestListGroupLDAPLinks(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/ldap_group_links",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `[
	{
		"cn":"gitlab_group_example_30",
		"group_access":30,
		"provider":"example_ldap_provider"
	},
	{
		"cn":"gitlab_group_example_40",
		"group_access":40,
		"provider":"example_ldap_provider"
	}
]`)
		})

	links, _, err := client.Groups.ListGroupLDAPLinks(1)
	if err != nil {
		t.Errorf("Groups.ListGroupLDAPLinks returned error: %v", err)
	}

	want := []*LDAPGroupLink{
		{
			CN:          "gitlab_group_example_30",
			GroupAccess: 30,
			Provider:    "example_ldap_provider",
		},
		{
			CN:          "gitlab_group_example_40",
			GroupAccess: 40,
			Provider:    "example_ldap_provider",
		},
	}
	if !reflect.DeepEqual(want, links) {
		t.Errorf("Groups.ListGroupLDAPLinks returned %+v, want %+v", links, want)
	}
}

func TestAddGroupLDAPLink(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/ldap_group_links",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `
{
	"cn":"gitlab_group_example_30",
	"group_access":30,
	"provider":"example_ldap_provider"
}`)
		})

	opt := &AddGroupLDAPLinkOptions{
		CN:          Ptr("gitlab_group_example_30"),
		GroupAccess: Ptr(AccessLevelValue(30)),
		Provider:    Ptr("example_ldap_provider"),
	}

	link, _, err := client.Groups.AddGroupLDAPLink(1, opt)
	if err != nil {
		t.Errorf("Groups.AddGroupLDAPLink returned error: %v", err)
	}

	want := &LDAPGroupLink{
		CN:          "gitlab_group_example_30",
		GroupAccess: 30,
		Provider:    "example_ldap_provider",
	}
	if !reflect.DeepEqual(want, link) {
		t.Errorf("Groups.AddGroupLDAPLink returned %+v, want %+v", link, want)
	}
}

func TestAddGroupLDAPLinkFilter(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/ldap_group_links",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `
{
	"filter":"(memberOf=example_group_dn)",
	"group_access":30,
	"provider":"example_ldap_provider"
}`)
		})

	opt := &AddGroupLDAPLinkOptions{
		Filter:      Ptr("(memberOf=example_group_dn)"),
		GroupAccess: Ptr(AccessLevelValue(30)),
		Provider:    Ptr("example_ldap_provider"),
	}

	link, _, err := client.Groups.AddGroupLDAPLink(1, opt)
	if err != nil {
		t.Errorf("Groups.AddGroupLDAPLink returned error: %v", err)
	}

	want := &LDAPGroupLink{
		Filter:      "(memberOf=example_group_dn)",
		GroupAccess: 30,
		Provider:    "example_ldap_provider",
	}
	if !reflect.DeepEqual(want, link) {
		t.Errorf("Groups.AddGroupLDAPLink returned %+v, want %+v", link, want)
	}
}

func TestListGroupSAMLLinks(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/saml_group_links",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `[
	{
		"access_level":30,
		"name":"gitlab_group_example_developer"
	},
	{
		"access_level":40,
		"name":"gitlab_group_example_maintainer"
	}
]`)
		})

	links, _, err := client.Groups.ListGroupSAMLLinks(1)
	if err != nil {
		t.Errorf("Groups.ListGroupSAMLLinks returned error: %v", err)
	}

	want := []*SAMLGroupLink{
		{
			AccessLevel: DeveloperPermissions,
			Name:        "gitlab_group_example_developer",
		},
		{
			AccessLevel: MaintainerPermissions,
			Name:        "gitlab_group_example_maintainer",
		},
	}
	if !reflect.DeepEqual(want, links) {
		t.Errorf("Groups.ListGroupSAMLLinks returned %+v, want %+v", links, want)
	}
}

func TestListGroupSAMLLinksCustomRole(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/saml_group_links",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `[
	{
		"access_level":30,
		"name":"gitlab_group_example_developer",
		"member_role_id":123
	}
]`)
		})

	links, _, err := client.Groups.ListGroupSAMLLinks(1)
	if err != nil {
		t.Errorf("Groups.ListGroupSAMLLinks returned error: %v", err)
	}

	want := []*SAMLGroupLink{
		{
			AccessLevel:  DeveloperPermissions,
			Name:         "gitlab_group_example_developer",
			MemberRoleID: 123,
		},
	}
	if !reflect.DeepEqual(want, links) {
		t.Errorf("Groups.ListGroupSAMLLinks returned %+v, want %+v", links, want)
	}
}

func TestGetGroupSAMLLink(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/saml_group_links/gitlab_group_example_developer",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `
{
	"access_level":30,
	"name":"gitlab_group_example_developer"
}`)
		})

	links, _, err := client.Groups.GetGroupSAMLLink(1, "gitlab_group_example_developer")
	if err != nil {
		t.Errorf("Groups.GetGroupSAMLLinks returned error: %v", err)
	}

	want := &SAMLGroupLink{
		AccessLevel: DeveloperPermissions,
		Name:        "gitlab_group_example_developer",
	}
	if !reflect.DeepEqual(want, links) {
		t.Errorf("Groups.GetGroupSAMLLink returned %+v, want %+v", links, want)
	}
}

func TestGetGroupSAMLLinkCustomRole(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/saml_group_links/gitlab_group_example_developer",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `
{
	"access_level":30,
	"name":"gitlab_group_example_developer",
	"member_role_id":123
}`)
		})

	links, _, err := client.Groups.GetGroupSAMLLink(1, "gitlab_group_example_developer")
	if err != nil {
		t.Errorf("Groups.GetGroupSAMLLinks returned error: %v", err)
	}

	want := &SAMLGroupLink{
		AccessLevel:  DeveloperPermissions,
		Name:         "gitlab_group_example_developer",
		MemberRoleID: 123,
	}
	if !reflect.DeepEqual(want, links) {
		t.Errorf("Groups.GetGroupSAMLLink returned %+v, want %+v", links, want)
	}
}

func TestAddGroupSAMLLink(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/saml_group_links",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `
{
	"access_level":30,
	"name":"gitlab_group_example_developer"
}`)
		})

	opt := &AddGroupSAMLLinkOptions{
		SAMLGroupName: Ptr("gitlab_group_example_developer"),
		AccessLevel:   Ptr(DeveloperPermissions),
	}

	link, _, err := client.Groups.AddGroupSAMLLink(1, opt)
	if err != nil {
		t.Errorf("Groups.AddGroupSAMLLink returned error: %v", err)
	}

	want := &SAMLGroupLink{
		AccessLevel: DeveloperPermissions,
		Name:        "gitlab_group_example_developer",
	}
	if !reflect.DeepEqual(want, link) {
		t.Errorf("Groups.AddGroupSAMLLink returned %+v, want %+v", link, want)
	}
}

func TestAddGroupSAMLLinkCustomRole(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/saml_group_links",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `
{
	"access_level":30,
	"name":"gitlab_group_example_developer",
	"member_role_id":123
}`)
		})

	opt := &AddGroupSAMLLinkOptions{
		SAMLGroupName: Ptr("gitlab_group_example_developer"),
		AccessLevel:   Ptr(DeveloperPermissions),
		MemberRoleID:  Ptr(123),
	}

	link, _, err := client.Groups.AddGroupSAMLLink(1, opt)
	if err != nil {
		t.Errorf("Groups.AddGroupSAMLLink returned error: %v", err)
	}

	want := &SAMLGroupLink{
		AccessLevel:  DeveloperPermissions,
		Name:         "gitlab_group_example_developer",
		MemberRoleID: 123,
	}
	if !reflect.DeepEqual(want, link) {
		t.Errorf("Groups.AddGroupSAMLLink returned %+v, want %+v", link, want)
	}
}

func TestRestoreGroup(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/groups/1/restore",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `{"id": 1, "name": "g"}`)
		})

	group, _, err := client.Groups.RestoreGroup(1)
	if err != nil {
		t.Errorf("Groups.RestoreGroup returned error: %v", err)
	}
	want := &Group{ID: 1, Name: "g"}
	if !reflect.DeepEqual(want, group) {
		t.Errorf("Groups.RestoreGroup returned %+v, want %+v", group, want)
	}
}

func TestShareGroupWithGroup(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/groups/1/share",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `{"id": 1, "name": "g"}`)
		})

	group, _, err := client.Groups.ShareGroupWithGroup(1, &ShareGroupWithGroupOptions{
		GroupID:     Ptr(1),
		GroupAccess: Ptr(DeveloperPermissions),
	})
	if err != nil {
		t.Errorf("Groups.ShareGroupWithGroup returned error: %v", err)
	}
	want := &Group{ID: 1, Name: "g"}
	if !reflect.DeepEqual(want, group) {
		t.Errorf("Groups.ShareGroupWithGroup returned %+v, want %+v", group, want)
	}
}

func TestUnshareGroupFromGroup(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/groups/1/share/2",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			w.WriteHeader(204)
		})

	r, err := client.Groups.UnshareGroupFromGroup(1, 2)
	if err != nil {
		t.Errorf("Groups.UnshareGroupFromGroup returned error: %v", err)
	}
	if r.StatusCode != 204 {
		t.Errorf("Groups.UnshareGroupFromGroup returned status code %d", r.StatusCode)
	}
}

func TestUpdateGroupWithIPRestrictionRanges(t *testing.T) {
	mux, client := setup(t)
	const ipRange = "192.168.0.0/24"

	mux.HandleFunc("/api/v4/groups/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("Failed to read the request body. Error: %v", err)
			}

			var bodyJson map[string]interface{}
			err = json.Unmarshal(body, &bodyJson)
			if err != nil {
				t.Fatalf("Failed to parse the request body into JSON. Error: %v", err)
			}

			if bodyJson["ip_restriction_ranges"] != ipRange {
				t.Fatalf("Test failed. `ip_restriction_ranges` expected to be '%v', got %v", ipRange, bodyJson["ip_restriction_ranges"])
			}

			fmt.Fprintf(w, `{"id": 1, "ip_restriction_ranges" : "%v"}`, ipRange)
		})

	group, _, err := client.Groups.UpdateGroup(1, &UpdateGroupOptions{
		IPRestrictionRanges: Ptr(ipRange),
	})
	if err != nil {
		t.Errorf("Groups.UpdateGroup returned error: %v", err)
	}

	want := &Group{ID: 1, IPRestrictionRanges: ipRange}
	if !reflect.DeepEqual(want, group) {
		t.Errorf("Groups.UpdatedGroup returned %+v, want %+v", group, want)
	}
}

func TestGetGroupWithEmailsEnabled(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)

			// Modified from https://docs.gitlab.com/ee/api/groups.html#details-of-a-group
			fmt.Fprint(w, `
			{
				"id": 1,
				"name": "test",
				"path": "test",
				"emails_enabled": true,
				"description": "Aliquid qui quis dignissimos distinctio ut commodi voluptas est.",
				"visibility": "public",
				"avatar_url": null,
				"web_url": "https://gitlab.example.com/groups/test",
				"request_access_enabled": false,
				"repository_storage": "default",
				"full_name": "test",
				"full_path": "test",
				"runners_token": "ba324ca7b1c77fc20bb9",
				"file_template_project_id": 1,
				"parent_id": null,
				"enabled_git_access_protocol": "all",
				"created_at": "2020-01-15T12:36:29.590Z",
				"prevent_sharing_groups_outside_hierarchy": false,
				"ip_restriction_ranges": null,
				"math_rendering_limits_enabled": true,
				"lock_math_rendering_limits_enabled": false
			  }`)
		})

	group, _, err := client.Groups.GetGroup(1, &GetGroupOptions{})
	if err != nil {
		t.Errorf("Groups.UpdateGroup returned error: %v", err)
	}

	if !group.EmailsEnabled {
		t.Fatalf("Failed to parse `emails_enabled`. Wanted true, got %v", group.EmailsEnabled)
	}
}

func TestCreateGroupWithEmailsEnabled(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("Failed to read the request body. Error: %v", err)
			}

			// unmarshal into generic JSON since we don't want to test CreateGroupOptions using itself to validate.
			var bodyJson map[string]interface{}
			err = json.Unmarshal(body, &bodyJson)
			if err != nil {
				t.Fatalf("Failed to parse the request body into JSON. Error: %v", err)
			}

			if bodyJson["emails_enabled"] != true {
				t.Fatalf("Test failed. `emails_enabled` expected to be true, got %v", bodyJson["emails_enabled"])
			}

			// Response is tested via the "GET" test, only test the actual request here.
			fmt.Fprint(w, `
			{}`)
		})

	_, _, err := client.Groups.CreateGroup(&CreateGroupOptions{EmailsEnabled: Ptr(true)})
	if err != nil {
		t.Errorf("Groups.CreateGroup returned error: %v", err)
	}
}

func TestUpdateGroupWithEmailsEnabled(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("Failed to read the request body. Error: %v", err)
			}

			// unmarshal into generic JSON since we don't want to test UpdateGroupOptions using itself to validate.
			var bodyJson map[string]interface{}
			err = json.Unmarshal(body, &bodyJson)
			if err != nil {
				t.Fatalf("Failed to parse the request body into JSON. Error: %v", err)
			}

			if bodyJson["emails_enabled"] != true {
				t.Fatalf("Test failed. `emails_enabled` expected to be true, got %v", bodyJson["emails_enabled"])
			}

			// Response is tested via the "GET" test, only test the actual request here.
			fmt.Fprint(w, `
			{}`)
		})

	_, _, err := client.Groups.UpdateGroup(1, &UpdateGroupOptions{EmailsEnabled: Ptr(true)})
	if err != nil {
		t.Errorf("Groups.UpdateGroup returned error: %v", err)
	}
}

func TestGetGroupPushRules(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/push_rule", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": 1,
			"commit_message_regex": "Fixes \\d+\\..*",
			"commit_message_negative_regex": "ssh\\:\\/\\/",
			"branch_name_regex": "(feat|fix)\\/*",
			"deny_delete_tag": false,
			"member_check": false,
			"prevent_secrets": false,
			"author_email_regex": "@company.com$",
			"file_name_regex": "(jar|exe)$",
			"max_file_size": 5,
			"commit_committer_check": false,
			"commit_committer_name_check": false,
			"reject_unsigned_commits": false,
			"reject_non_dco_commits": false
		  }`)
	})

	rule, _, err := client.Groups.GetGroupPushRules(1)
	if err != nil {
		t.Errorf("Groups.GetGroupPushRules returned error: %v", err)
	}

	want := &GroupPushRules{
		ID:                         1,
		CommitMessageRegex:         "Fixes \\d+\\..*",
		CommitMessageNegativeRegex: "ssh\\:\\/\\/",
		BranchNameRegex:            "(feat|fix)\\/*",
		DenyDeleteTag:              false,
		MemberCheck:                false,
		PreventSecrets:             false,
		AuthorEmailRegex:           "@company.com$",
		FileNameRegex:              "(jar|exe)$",
		MaxFileSize:                5,
		CommitCommitterCheck:       false,
		CommitCommitterNameCheck:   false,
		RejectUnsignedCommits:      false,
		RejectNonDCOCommits:        false,
	}

	if !reflect.DeepEqual(want, rule) {
		t.Errorf("Groups.GetGroupPushRules returned %+v, want %+v", rule, want)
	}
}

func TestAddGroupPushRules(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/push_rule", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"id": 1,
			"commit_message_regex": "Fixes \\d+\\..*",
			"commit_message_negative_regex": "ssh\\:\\/\\/",
			"branch_name_regex": "(feat|fix)\\/*",
			"deny_delete_tag": false,
			"member_check": false,
			"prevent_secrets": false,
			"author_email_regex": "@company.com$",
			"file_name_regex": "(jar|exe)$",
			"max_file_size": 5,
			"commit_committer_check": false,
			"commit_committer_name_check": false,
			"reject_unsigned_commits": false,
			"reject_non_dco_commits": false
		  }`)
	})

	opt := &AddGroupPushRuleOptions{
		CommitMessageRegex:         Ptr("Fixes \\d+\\..*"),
		CommitMessageNegativeRegex: Ptr("ssh\\:\\/\\/"),
		BranchNameRegex:            Ptr("(feat|fix)\\/*"),
		DenyDeleteTag:              Ptr(false),
		MemberCheck:                Ptr(false),
		PreventSecrets:             Ptr(false),
		AuthorEmailRegex:           Ptr("@company.com$"),
		FileNameRegex:              Ptr("(jar|exe)$"),
		MaxFileSize:                Ptr(5),
		CommitCommitterCheck:       Ptr(false),
		CommitCommitterNameCheck:   Ptr(false),
		RejectUnsignedCommits:      Ptr(false),
		RejectNonDCOCommits:        Ptr(false),
	}

	rule, _, err := client.Groups.AddGroupPushRule(1, opt)
	if err != nil {
		t.Errorf("Groups.AddGroupPushRule returned error: %v", err)
	}

	want := &GroupPushRules{
		ID:                         1,
		CommitMessageRegex:         "Fixes \\d+\\..*",
		CommitMessageNegativeRegex: "ssh\\:\\/\\/",
		BranchNameRegex:            "(feat|fix)\\/*",
		DenyDeleteTag:              false,
		MemberCheck:                false,
		PreventSecrets:             false,
		AuthorEmailRegex:           "@company.com$",
		FileNameRegex:              "(jar|exe)$",
		MaxFileSize:                5,
		CommitCommitterCheck:       false,
		CommitCommitterNameCheck:   false,
		RejectUnsignedCommits:      false,
		RejectNonDCOCommits:        false,
	}

	if !reflect.DeepEqual(want, rule) {
		t.Errorf("Groups.AddGroupPushRule returned %+v, want %+v", rule, want)
	}
}

func TestEditGroupPushRules(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/push_rule", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{
			"id": 1,
			"commit_message_regex": "Fixes \\d+\\..*",
			"commit_message_negative_regex": "ssh\\:\\/\\/",
			"branch_name_regex": "(feat|fix)\\/*",
			"deny_delete_tag": false,
			"member_check": false,
			"prevent_secrets": false,
			"author_email_regex": "@company.com$",
			"file_name_regex": "(jar|exe)$",
			"max_file_size": 5,
			"commit_committer_check": false,
			"commit_committer_name_check": false,
			"reject_unsigned_commits": false,
			"reject_non_dco_commits": false
		  }`)
	})

	opt := &EditGroupPushRuleOptions{
		CommitMessageRegex:         Ptr("Fixes \\d+\\..*"),
		CommitMessageNegativeRegex: Ptr("ssh\\:\\/\\/"),
		BranchNameRegex:            Ptr("(feat|fix)\\/*"),
		DenyDeleteTag:              Ptr(false),
		MemberCheck:                Ptr(false),
		PreventSecrets:             Ptr(false),
		AuthorEmailRegex:           Ptr("@company.com$"),
		FileNameRegex:              Ptr("(jar|exe)$"),
		MaxFileSize:                Ptr(5),
		CommitCommitterCheck:       Ptr(false),
		CommitCommitterNameCheck:   Ptr(false),
		RejectUnsignedCommits:      Ptr(false),
		RejectNonDCOCommits:        Ptr(false),
	}

	rule, _, err := client.Groups.EditGroupPushRule(1, opt)
	if err != nil {
		t.Errorf("Groups.EditGroupPushRule returned error: %v", err)
	}

	want := &GroupPushRules{
		ID:                         1,
		CommitMessageRegex:         "Fixes \\d+\\..*",
		CommitMessageNegativeRegex: "ssh\\:\\/\\/",
		BranchNameRegex:            "(feat|fix)\\/*",
		DenyDeleteTag:              false,
		MemberCheck:                false,
		PreventSecrets:             false,
		AuthorEmailRegex:           "@company.com$",
		FileNameRegex:              "(jar|exe)$",
		MaxFileSize:                5,
		CommitCommitterCheck:       false,
		CommitCommitterNameCheck:   false,
		RejectUnsignedCommits:      false,
		RejectNonDCOCommits:        false,
	}

	if !reflect.DeepEqual(want, rule) {
		t.Errorf("Groups.EditGroupPushRule returned %+v, want %+v", rule, want)
	}
}

func TestUpdateGroupWithAllowedEmailDomainsList(t *testing.T) {
	mux, client := setup(t)
	const domain = "example.com"

	mux.HandleFunc("/api/v4/groups/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("Failed to read the request body. Error: %v", err)
			}

			var bodyJson map[string]interface{}
			err = json.Unmarshal(body, &bodyJson)
			if err != nil {
				t.Fatalf("Failed to parse the request body into JSON. Error: %v", err)
			}

			if bodyJson["allowed_email_domains_list"] != domain {
				t.Fatalf("Test failed. `allowed_email_domains_list` expected to be '%v', got %v", domain, bodyJson["allowed_email_domains_list"])
			}

			fmt.Fprintf(w, `{"id": 1, "allowed_email_domains_list" : "%v"}`, domain)
		})

	group, _, err := client.Groups.UpdateGroup(1, &UpdateGroupOptions{
		AllowedEmailDomainsList: Ptr(domain),
	})
	if err != nil {
		t.Errorf("Groups.UpdateGroup returned error: %v", err)
	}

	want := &Group{ID: 1, AllowedEmailDomainsList: domain}
	if !reflect.DeepEqual(want, group) {
		t.Errorf("Groups.UpdatedGroup returned %+v, want %+v", group, want)
	}
}
