package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProjectMembersService_ListProjectMembers(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/members", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 1,
				"username": "venkatesh_thalluri",
				"name": "Venkatesh Thalluri",
				"state": "active",
				"avatar_url": "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
				"web_url": "http://192.168.1.8:3000/root",
				"access_level": 30,
				"group_saml_identity": null
			  }
			]
		`)
	})

	want := []*ProjectMember{{
		ID:          1,
		Username:    "venkatesh_thalluri",
		Email:       "",
		Name:        "Venkatesh Thalluri",
		State:       "active",
		AccessLevel: 30,
		WebURL:      "http://192.168.1.8:3000/root",
		AvatarURL:   "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
	}}

	pms, resp, err := client.ProjectMembers.ListProjectMembers(1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pms)

	pms, resp, err = client.ProjectMembers.ListProjectMembers(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pms)

	pms, resp, err = client.ProjectMembers.ListProjectMembers(1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pms)

	pms, resp, err = client.ProjectMembers.ListProjectMembers(2, nil, nil)
	require.Error(t, err)
	require.Nil(t, pms)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectMembersService_ListAllProjectMembers(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/members/all", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 1,
				"username": "venkatesh_thalluri",
				"name": "Venkatesh Thalluri",
				"state": "active",
				"avatar_url": "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
				"web_url": "http://192.168.1.8:3000/root",
				"access_level": 30,
				"group_saml_identity": null
			  }
			]
		`)
	})

	want := []*ProjectMember{{
		ID:          1,
		Username:    "venkatesh_thalluri",
		Email:       "",
		Name:        "Venkatesh Thalluri",
		State:       "active",
		AccessLevel: 30,
		WebURL:      "http://192.168.1.8:3000/root",
		AvatarURL:   "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
	}}

	pms, resp, err := client.ProjectMembers.ListAllProjectMembers(1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pms)

	pms, resp, err = client.ProjectMembers.ListAllProjectMembers(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pms)

	pms, resp, err = client.ProjectMembers.ListAllProjectMembers(1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pms)

	pms, resp, err = client.ProjectMembers.ListAllProjectMembers(2, nil, nil)
	require.Error(t, err)
	require.Nil(t, pms)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectMembersService_GetProjectMember(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/members/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
			  "id": 1,
			  "username": "venkatesh_thalluri",
			  "name": "Venkatesh Thalluri",
			  "state": "active",
			  "avatar_url": "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
			  "web_url": "http://192.168.1.8:3000/root",
			  "access_level": 30,
			  "email": "venkatesh.thalluri@example.com",
			  "expires_at": null,
			  "group_saml_identity": null
			}
		`)
	})

	want := &ProjectMember{
		ID:          1,
		Username:    "venkatesh_thalluri",
		Email:       "venkatesh.thalluri@example.com",
		Name:        "Venkatesh Thalluri",
		State:       "active",
		ExpiresAt:   nil,
		AccessLevel: 30,
		WebURL:      "http://192.168.1.8:3000/root",
		AvatarURL:   "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
	}

	pm, resp, err := client.ProjectMembers.GetProjectMember(1, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pm)

	pm, resp, err = client.ProjectMembers.GetProjectMember(1.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pm)

	pm, resp, err = client.ProjectMembers.GetProjectMember(1, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pm)

	pm, resp, err = client.ProjectMembers.GetProjectMember(2, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, pm)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectMembersService_GetInheritedProjectMember(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/members/all/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
			  "id": 1,
			  "username": "venkatesh_thalluri",
			  "name": "Venkatesh Thalluri",
			  "state": "active",
			  "avatar_url": "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
			  "web_url": "http://192.168.1.8:3000/root",
			  "access_level": 30,
			  "email": "venkatesh.thalluri@example.com",
			  "expires_at": null,
			  "group_saml_identity": null
			}
		`)
	})

	want := &ProjectMember{
		ID:          1,
		Username:    "venkatesh_thalluri",
		Email:       "venkatesh.thalluri@example.com",
		Name:        "Venkatesh Thalluri",
		State:       "active",
		ExpiresAt:   nil,
		AccessLevel: 30,
		WebURL:      "http://192.168.1.8:3000/root",
		AvatarURL:   "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
	}

	pm, resp, err := client.ProjectMembers.GetInheritedProjectMember(1, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pm)

	pm, resp, err = client.ProjectMembers.GetInheritedProjectMember(1.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pm)

	pm, resp, err = client.ProjectMembers.GetInheritedProjectMember(1, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pm)

	pm, resp, err = client.ProjectMembers.GetInheritedProjectMember(2, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, pm)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectMembersService_AddProjectMember(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/members", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
			  "id": 1,
			  "username": "venkatesh_thalluri",
			  "name": "Venkatesh Thalluri",
			  "state": "active",
			  "avatar_url": "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
			  "web_url": "http://192.168.1.8:3000/root",
			  "access_level": 30,
			  "email": "venkatesh.thalluri@example.com",
			  "expires_at": null,
			  "group_saml_identity": null
			}
		`)
	})

	want := &ProjectMember{
		ID:          1,
		Username:    "venkatesh_thalluri",
		Email:       "venkatesh.thalluri@example.com",
		Name:        "Venkatesh Thalluri",
		State:       "active",
		ExpiresAt:   nil,
		AccessLevel: 30,
		WebURL:      "http://192.168.1.8:3000/root",
		AvatarURL:   "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
	}

	pm, resp, err := client.ProjectMembers.AddProjectMember(1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pm)

	pm, resp, err = client.ProjectMembers.AddProjectMember(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pm)

	pm, resp, err = client.ProjectMembers.AddProjectMember(1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pm)

	pm, resp, err = client.ProjectMembers.AddProjectMember(2, nil, nil)
	require.Error(t, err)
	require.Nil(t, pm)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectMembersService_EditProjectMember(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/members/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
			{
			  "id": 1,
			  "username": "venkatesh_thalluri",
			  "name": "Venkatesh Thalluri",
			  "state": "active",
			  "avatar_url": "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
			  "web_url": "http://192.168.1.8:3000/root",
			  "access_level": 30,
			  "email": "venkatesh.thalluri@example.com",
			  "expires_at": null,
			  "group_saml_identity": null
			}
		`)
	})

	want := &ProjectMember{
		ID:          1,
		Username:    "venkatesh_thalluri",
		Email:       "venkatesh.thalluri@example.com",
		Name:        "Venkatesh Thalluri",
		State:       "active",
		ExpiresAt:   nil,
		AccessLevel: 30,
		WebURL:      "http://192.168.1.8:3000/root",
		AvatarURL:   "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
	}

	pm, resp, err := client.ProjectMembers.EditProjectMember(1, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, pm)

	pm, resp, err = client.ProjectMembers.EditProjectMember(1.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, pm)

	pm, resp, err = client.ProjectMembers.EditProjectMember(1, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, pm)

	pm, resp, err = client.ProjectMembers.EditProjectMember(2, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, pm)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectMembersService_DeleteProjectMember(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/members/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprintf(w, `
			{
			  "id": 1,
			  "username": "venkatesh_thalluri",
			  "name": "Venkatesh Thalluri",
			  "state": "active",
			  "avatar_url": "https://www.gravatar.com/avatar/c2525a7f58ae3776070e44c106c48e15?s=80&d=identicon",
			  "web_url": "http://192.168.1.8:3000/root",
			  "access_level": 30,
			  "email": "venkatesh.thalluri@example.com",
			  "expires_at": null,
			  "group_saml_identity": null
			}
		`)
	})

	resp, err := client.ProjectMembers.DeleteProjectMember(1, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.ProjectMembers.DeleteProjectMember(1.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.ProjectMembers.DeleteProjectMember(1, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.ProjectMembers.DeleteProjectMember(2, 1, nil, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
