package gitlab

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
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
			ID:                       1,
			Name:                     "GuestCodeReader",
			Description:              "A Guest user that can read code",
			GroupId:                  1,
			BaseAccessLevel:          10, // Guest Base Level
			AdminMergeRequests:       false,
			AdminVulnerability:       false,
			ReadCode:                 true,
			ReadDependency:           false,
			ReadVulnerability:        false,
			ManageProjectAccessToken: false,
		},
		{
			ID:                       2,
			Name:                     "GuestVulnerabilityReader",
			Description:              "A Guest user that can read vulnerabilities",
			GroupId:                  1,
			BaseAccessLevel:          10, // Guest Base Level
			AdminMergeRequests:       false,
			AdminVulnerability:       false,
			ReadCode:                 false,
			ReadDependency:           false,
			ReadVulnerability:        true,
			ManageProjectAccessToken: false,
		},
	}

	require.Equal(t, want, memberRoles)
}
