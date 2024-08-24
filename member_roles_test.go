package gitlab

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListMemberRoles(t *testing.T) {
	mux, client := setup(t)

	path := "/api/v4/groups/1/member_roles"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_member_roles.json")
	})

	memberRoles, _, err := client.MemberRolesService.ListMemberRoles(1)
	require.NoError(t, err)

	want := []*MemberRole{
		{
			ID:              1,
			Name:            "GuestCodeReader",
			Description:     "A Guest user that can read code",
			GroupId:         1,
			BaseAccessLevel: 10, // Guest Base Level
			ReadCode:        true,
		},
		{
			ID:                2,
			Name:              "GuestVulnerabilityReader",
			Description:       "A Guest user that can read vulnerabilities",
			GroupId:           1,
			BaseAccessLevel:   10, // Guest Base Level
			ReadVulnerability: true,
		},
	}

	require.Equal(t, want, memberRoles)
}

func TestCreateMemberRole(t *testing.T) {
	mux, client := setup(t)

	path := "/api/v4/groups/84/member_roles"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		mustWriteHTTPResponse(t, w, "testdata/create_member_role.json")
	})

	memberRole, _, err := client.MemberRolesService.CreateMemberRole(84, &CreateMemberRoleOptions{
		Name:                       Ptr("Custom guest"),
		BaseAccessLevel:            Ptr(GuestPermissions),
		Description:                Ptr("a sample custom role"),
		AdminCICDVariables:         Ptr(false),
		AdminComplianceFramework:   Ptr(false),
		AdminGroupMembers:          Ptr(false),
		AdminMergeRequest:          Ptr(false),
		AdminPushRules:             Ptr(false),
		AdminTerraformState:        Ptr(false),
		AdminVulnerability:         Ptr(false),
		AdminWebHook:               Ptr(false),
		ArchiveProject:             Ptr(false),
		ManageDeployTokens:         Ptr(false),
		ManageGroupAccesToken:      Ptr(false),
		ManageMergeRequestSettings: Ptr(false),
		ManageProjectAccessToken:   Ptr(false),
		ManageSecurityPolicyLink:   Ptr(false),
		ReadCode:                   Ptr(true),
		ReadRunners:                Ptr(false),
		ReadDependency:             Ptr(false),
		ReadVulnerability:          Ptr(false),
		RemoveGroup:                Ptr(false),
		RemoveProject:              Ptr(false),
	})
	require.NoError(t, err)

	want := &MemberRole{
		ID:                         3,
		Name:                       "Custom guest",
		Description:                "a sample custom role",
		BaseAccessLevel:            GuestPermissions,
		GroupId:                    84,
		AdminCICDVariables:         false,
		AdminComplianceFramework:   false,
		AdminGroupMembers:          false,
		AdminMergeRequests:         false,
		AdminPushRules:             false,
		AdminTerraformState:        false,
		AdminVulnerability:         false,
		AdminWebHook:               false,
		ArchiveProject:             false,
		ManageDeployTokens:         false,
		ManageGroupAccesToken:      false,
		ManageMergeRequestSettings: false,
		ManageProjectAccessToken:   false,
		ManageSecurityPolicyLink:   false,
		ReadCode:                   true,
		ReadRunners:                false,
		ReadDependency:             false,
		ReadVulnerability:          false,
		RemoveGroup:                false,
		RemoveProject:              false,
	}

	require.Equal(t, want, memberRole)
}

func TestDeleteMemberRole(t *testing.T) {
	mux, client := setup(t)

	path := "/api/v4/groups/1/member_roles/2"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.MemberRolesService.DeleteMemberRole(1, 2)
	require.NoError(t, err)
}
