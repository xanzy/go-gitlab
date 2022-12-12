package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
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
			fmt.Fprint(w, `{"id": 1, "name": "g"}`)
		})

	group, _, err := client.Groups.GetGroup("g", &GetGroupOptions{})
	if err != nil {
		t.Errorf("Groups.GetGroup returned error: %v", err)
	}

	want := &Group{ID: 1, Name: "g"}
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
		Name: String("g"),
		Path: String("g"),
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
		GroupID: Int(2),
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

	resp, err := client.Groups.DeleteGroup(1)
	if err != nil {
		t.Errorf("Groups.DeleteGroup returned error: %v", err)
	}

	want := http.StatusAccepted
	got := resp.StatusCode
	if got != want {
		t.Errorf("Groups.DeleteGroup returned %d, want %d", got, want)
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
		CN:          String("gitlab_group_example_30"),
		GroupAccess: AccessLevel(30),
		Provider:    String("example_ldap_provider"),
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
		Filter:      String("(memberOf=example_group_dn)"),
		GroupAccess: AccessLevel(30),
		Provider:    String("example_ldap_provider"),
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
		SAMLGroupName: String("gitlab_group_example_developer"),
		AccessLevel:   AccessLevel(DeveloperPermissions),
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
		GroupID:     Int(1),
		GroupAccess: AccessLevel(DeveloperPermissions),
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

func TestCreateGroupWithIPRestrictionRanges(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `{"id": 1, "name": "g", "path": "g", "ip_restriction_ranges" : "192.168.0.0/24"}`)
		})

	opt := &CreateGroupOptions{
		Name:                String("g"),
		Path:                String("g"),
		IPRestrictionRanges: String("192.168.0.0/24"),
	}

	group, _, err := client.Groups.CreateGroup(opt, nil)
	if err != nil {
		t.Errorf("Groups.CreateGroup returned error: %v", err)
	}

	want := &Group{ID: 1, Name: "g", Path: "g", IPRestrictionRanges: "192.168.0.0/24"}
	if !reflect.DeepEqual(want, group) {
		t.Errorf("Groups.CreateGroup returned %+v, want %+v", group, want)
	}
}

func TestUpdateGroupWithIPRestrictionRanges(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			fmt.Fprint(w, `{"id": 1, "ip_restriction_ranges" : "192.168.0.0/24"}`)
		})

	group, _, err := client.Groups.UpdateGroup(1, &UpdateGroupOptions{
		IPRestrictionRanges: String("192.168.0.0/24"),
	})
	if err != nil {
		t.Errorf("Groups.UpdateGroup returned error: %v", err)
	}

	want := &Group{ID: 1, IPRestrictionRanges: "192.168.0.0/24"}
	if !reflect.DeepEqual(want, group) {
		t.Errorf("Groups.UpdatedGroup returned %+v, want %+v", group, want)
	}
}
