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
	"errors"
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	js = User{
		ID:                             1,
		Username:                       "john_smith",
		Email:                          "",
		Name:                           "John Smith",
		State:                          "active",
		WebURL:                         "http://localhost:3000/john_smith",
		CreatedAt:                      timePtr(time.Date(2012, time.May, 23, 8, 00, 58, 0, time.UTC)),
		Bio:                            "Bio of John Smith",
		Location:                       "USA",
		PublicEmail:                    "john@example.com",
		Skype:                          "john_smith",
		Linkedin:                       "john_smith",
		Twitter:                        "john_smith",
		WebsiteURL:                     "john_smith.example.com",
		Organization:                   "Smith Inc",
		JobTitle:                       "Operations Specialist",
		ExternUID:                      "",
		Provider:                       "",
		ThemeID:                        0,
		LastActivityOn:                 nil,
		ColorSchemeID:                  0,
		IsAdmin:                        false,
		AvatarURL:                      "http://localhost:3000/uploads/user/avatar/1/cd8.jpeg",
		CanCreateGroup:                 false,
		CanCreateProject:               false,
		ProjectsLimit:                  0,
		CurrentSignInAt:                nil,
		CurrentSignInIP:                nil,
		LastSignInAt:                   nil,
		LastSignInIP:                   nil,
		ConfirmedAt:                    nil,
		TwoFactorEnabled:               false,
		Note:                           "",
		Identities:                     nil,
		External:                       false,
		PrivateProfile:                 false,
		SharedRunnersMinutesLimit:      0,
		ExtraSharedRunnersMinutesLimit: 0,
		UsingLicenseSeat:               false,
		CustomAttributes:               nil,
	}
	js_admin = User{
		ID:                             1,
		Username:                       "john_smith",
		Email:                          "john@example.com",
		Name:                           "John Smith",
		State:                          "active",
		WebURL:                         "http://localhost:3000/john_smith",
		CreatedAt:                      timePtr(time.Date(2012, time.May, 23, 8, 0, 58, 0, time.UTC)),
		Bio:                            "Bio of John Smith",
		Location:                       "USA",
		PublicEmail:                    "john@example.com",
		Skype:                          "john_smith",
		Linkedin:                       "john_smith",
		Twitter:                        "john_smith",
		WebsiteURL:                     "john_smith.example.com",
		Organization:                   "Smith Inc",
		JobTitle:                       "Operations Specialist",
		ExternUID:                      "",
		Provider:                       "",
		ThemeID:                        1,
		LastActivityOn:                 isoTimePtr(ISOTime(time.Date(2012, time.May, 23, 0, 0, 0, 0, time.UTC))),
		ColorSchemeID:                  2,
		IsAdmin:                        true,
		AvatarURL:                      "http://localhost:3000/uploads/user/avatar/1/index.jpg",
		CanCreateGroup:                 true,
		CanCreateProject:               true,
		ProjectsLimit:                  100,
		CurrentSignInAt:                timePtr(time.Date(2012, time.June, 2, 6, 36, 55, 0, time.UTC)),
		CurrentSignInIP:                ipPtr(net.ParseIP("8.8.8.8")),
		LastSignInAt:                   timePtr(time.Date(2012, time.June, 1, 11, 41, 1, 0, time.UTC)),
		LastSignInIP:                   ipPtr(net.ParseIP("2001:db8::68")),
		ConfirmedAt:                    timePtr(time.Date(2012, time.May, 23, 9, 05, 22, 0, time.UTC)),
		TwoFactorEnabled:               true,
		Note:                           "DMCA Request: 2018-11-05 | DMCA Violation | Abuse | https://gitlab.zendesk.com/agent/tickets/123",
		Identities:                     []*UserIdentity{&UserIdentity{Provider: "github", ExternUID: "2435223452345"}},
		External:                       false,
		PrivateProfile:                 false,
		SharedRunnersMinutesLimit:      0,
		ExtraSharedRunnersMinutesLimit: 0,
		UsingLicenseSeat:               false,
		CustomAttributes:               nil,
	}
	getUserOpts = GetUsersOptions{}
)

func timePtr(t time.Time) *time.Time {
	return &t
}

func isoTimePtr(i ISOTime) *ISOTime {
	return &i
}

func ipPtr(n net.IP) *net.IP {
	return &n
}

func TestGetUser(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := "/api/v4/users/1"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_user.json")
	})

	user, _, err := client.Users.GetUser(1, getUserOpts)

	require.NoError(t, err)

	require.Equal(t, user.ID, js.ID)
	require.Equal(t, user.Username, js.Username)
	require.Equal(t, user.Email, js.Email)
	require.Equal(t, user.Name, js.Name)
	require.Equal(t, user.State, js.State)
	require.Equal(t, user.WebURL, js.WebURL)
	require.Equal(t, user.CreatedAt, js.CreatedAt)
	require.Equal(t, user.Bio, js.Bio)
	require.Equal(t, user.Location, js.Location)
	require.Equal(t, user.PublicEmail, js.PublicEmail)
	require.Equal(t, user.Skype, js.Skype)
	require.Equal(t, user.Linkedin, js.Linkedin)
	require.Equal(t, user.Twitter, js.Twitter)
	require.Equal(t, user.WebsiteURL, js.WebsiteURL)
	require.Equal(t, user.Organization, js.Organization)
	require.Equal(t, user.JobTitle, js.JobTitle)
	require.Equal(t, user.ExternUID, js.ExternUID)
	require.Equal(t, user.Provider, js.Provider)
	require.Equal(t, user.ThemeID, js.ThemeID)
	require.Equal(t, user.LastActivityOn, js.LastActivityOn)
	require.Equal(t, user.ColorSchemeID, js.ColorSchemeID)
	require.Equal(t, user.IsAdmin, js.IsAdmin)
	require.Equal(t, user.AvatarURL, js.AvatarURL)
	require.Equal(t, user.CanCreateGroup, js.CanCreateGroup)
	require.Equal(t, user.CanCreateProject, js.CanCreateProject)
	require.Equal(t, user.ProjectsLimit, js.ProjectsLimit)
	require.Equal(t, user.CurrentSignInAt, js.CurrentSignInAt)
	require.Equal(t, user.CurrentSignInIP, js.CurrentSignInIP)
	require.Equal(t, user.LastSignInAt, js.LastSignInAt)
	require.Equal(t, user.LastSignInIP, js.LastSignInIP)
	require.Equal(t, user.ConfirmedAt, js.ConfirmedAt)
	require.Equal(t, user.TwoFactorEnabled, js.TwoFactorEnabled)
	require.Equal(t, user.Note, js.Note)
	require.Equal(t, user.Identities, js.Identities)
	require.Equal(t, user.External, js.External)
	require.Equal(t, user.PrivateProfile, js.PrivateProfile)
	require.Equal(t, user.SharedRunnersMinutesLimit, js.SharedRunnersMinutesLimit)
	require.Equal(t, user.ExtraSharedRunnersMinutesLimit, js.ExtraSharedRunnersMinutesLimit)
	require.Equal(t, user.UsingLicenseSeat, js.UsingLicenseSeat)
	require.Equal(t, user.CustomAttributes, js.CustomAttributes)
}

func TestGetUserAdmin(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := "/api/v4/users/1"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_user_admin.json")
	})

	user, _, err := client.Users.GetUser(1, getUserOpts)

	require.NoError(t, err)

	require.Equal(t, user.ID, js_admin.ID)
	require.Equal(t, user.Username, js_admin.Username)
	require.Equal(t, user.Email, js_admin.Email)
	require.Equal(t, user.Name, js_admin.Name)
	require.Equal(t, user.State, js_admin.State)
	require.Equal(t, user.WebURL, js_admin.WebURL)
	require.Equal(t, user.CreatedAt, js_admin.CreatedAt)
	require.Equal(t, user.Bio, js_admin.Bio)
	require.Equal(t, user.Location, js_admin.Location)
	require.Equal(t, user.PublicEmail, js_admin.PublicEmail)
	require.Equal(t, user.Skype, js_admin.Skype)
	require.Equal(t, user.Linkedin, js_admin.Linkedin)
	require.Equal(t, user.Twitter, js_admin.Twitter)
	require.Equal(t, user.WebsiteURL, js_admin.WebsiteURL)
	require.Equal(t, user.Organization, js_admin.Organization)
	require.Equal(t, user.JobTitle, js_admin.JobTitle)
	require.Equal(t, user.ExternUID, js_admin.ExternUID)
	require.Equal(t, user.Provider, js_admin.Provider)
	require.Equal(t, user.ThemeID, js_admin.ThemeID)
	require.Equal(t, user.LastActivityOn, js_admin.LastActivityOn)
	require.Equal(t, user.ColorSchemeID, js_admin.ColorSchemeID)
	require.Equal(t, user.IsAdmin, js_admin.IsAdmin)
	require.Equal(t, user.AvatarURL, js_admin.AvatarURL)
	require.Equal(t, user.CanCreateGroup, js_admin.CanCreateGroup)
	require.Equal(t, user.CanCreateProject, js_admin.CanCreateProject)
	require.Equal(t, user.ProjectsLimit, js_admin.ProjectsLimit)
	require.Equal(t, user.CurrentSignInAt, js_admin.CurrentSignInAt)
	require.Equal(t, user.CurrentSignInIP, js_admin.CurrentSignInIP)
	require.Equal(t, user.LastSignInAt, js_admin.LastSignInAt)
	require.Equal(t, user.LastSignInIP, js_admin.LastSignInIP)
	require.Equal(t, user.ConfirmedAt, js_admin.ConfirmedAt)
	require.Equal(t, user.TwoFactorEnabled, js_admin.TwoFactorEnabled)
	require.Equal(t, user.Note, js_admin.Note)
	require.Equal(t, user.Identities, js_admin.Identities)
	require.Equal(t, user.External, js_admin.External)
	require.Equal(t, user.PrivateProfile, js_admin.PrivateProfile)
	require.Equal(t, user.SharedRunnersMinutesLimit, js_admin.SharedRunnersMinutesLimit)
	require.Equal(t, user.ExtraSharedRunnersMinutesLimit, js_admin.ExtraSharedRunnersMinutesLimit)
	require.Equal(t, user.UsingLicenseSeat, js_admin.UsingLicenseSeat)
	require.Equal(t, user.CustomAttributes, js_admin.CustomAttributes)
}

func TestBlockUser(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/block", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.BlockUser(1)
	if err != nil {
		t.Errorf("Users.BlockUser returned error: %v", err)
	}
}

func TestBlockUser_UserNotFound(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/block", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.BlockUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.BlockUser error.\nExpected: %+v\nGot: %+v", ErrUserNotFound, err)
	}
}

func TestBlockUser_BlockPrevented(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/block", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.BlockUser(1)
	if !errors.Is(err, ErrUserBlockPrevented) {
		t.Errorf("Users.BlockUser error.\nExpected: %+v\nGot: %+v", ErrUserBlockPrevented, err)
	}
}

func TestBlockUser_UnknownError(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/block", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusTeapot)
	})

	want := fmt.Sprintf("Received unexpected result code: %d", http.StatusTeapot)

	err := client.Users.BlockUser(1)
	if err.Error() != want {
		t.Errorf("Users.BlockUser error.\nExpected: %s\nGot: %v", want, err)
	}
}

func TestUnblockUser(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/unblock", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.UnblockUser(1)
	if err != nil {
		t.Errorf("Users.UnblockUser returned error: %v", err)
	}
}

func TestUnblockUser_UserNotFound(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/unblock", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.UnblockUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.UnblockUser error.\nExpected: %v\nGot: %v", ErrUserNotFound, err)
	}
}

func TestUnblockUser_UnblockPrevented(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/unblock", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.UnblockUser(1)
	if !errors.Is(err, ErrUserUnblockPrevented) {
		t.Errorf("Users.UnblockUser error.\nExpected: %v\nGot: %v", ErrUserUnblockPrevented, err)
	}
}

func TestUnblockUser_UnknownError(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/unblock", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusTeapot)
	})

	want := fmt.Sprintf("Received unexpected result code: %d", http.StatusTeapot)

	err := client.Users.UnblockUser(1)
	if err.Error() != want {
		t.Errorf("Users.UnblockUser error.\nExpected: %s\n\tGot: %v", want, err)
	}
}

func TestDeactivateUser(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/deactivate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.DeactivateUser(1)
	if err != nil {
		t.Errorf("Users.DeactivateUser returned error: %v", err)
	}
}

func TestDeactivateUser_UserNotFound(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/deactivate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.DeactivateUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.DeactivateUser error.\nExpected: %+v\n\tGot: %+v", ErrUserNotFound, err)
	}
}

func TestDeactivateUser_DeactivatePrevented(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/deactivate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.DeactivateUser(1)
	if !errors.Is(err, ErrUserDeactivatePrevented) {
		t.Errorf("Users.DeactivateUser error.\nExpected: %+v\n\tGot: %+v", ErrUserDeactivatePrevented, err)
	}
}

func TestActivateUser(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/activate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.ActivateUser(1)
	if err != nil {
		t.Errorf("Users.ActivateUser returned error: %v", err)
	}
}

func TestActivateUser_ActivatePrevented(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/activate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.ActivateUser(1)
	if !errors.Is(err, ErrUserActivatePrevented) {
		t.Errorf("Users.ActivateUser error.\nExpected: %+v\n\tGot: %+v", ErrUserActivatePrevented, err)
	}
}

func TestActivateUser_UserNotFound(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/activate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.ActivateUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.ActivateUser error.\nExpected: %+v\n\tGot: %+v", ErrUserNotFound, err)
	}
}

func TestApproveUser(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/approve", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.ApproveUser(1)
	if err != nil {
		t.Errorf("Users.ApproveUser returned error: %v", err)
	}
}

func TestApproveUser_UserNotFound(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/approve", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.ApproveUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.ApproveUser error.\nExpected: %v\nGot: %v", ErrUserNotFound, err)
	}
}

func TestApproveUser_ApprovePrevented(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/approve", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.ApproveUser(1)
	if !errors.Is(err, ErrUserApprovePrevented) {
		t.Errorf("Users.ApproveUser error.\nExpected: %v\nGot: %v", ErrUserApprovePrevented, err)
	}
}

func TestApproveUser_UnknownError(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/approve", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusTeapot)
	})

	want := fmt.Sprintf("Received unexpected result code: %d", http.StatusTeapot)

	err := client.Users.ApproveUser(1)
	if err.Error() != want {
		t.Errorf("Users.ApproveUser error.\nExpected: %s\n\tGot: %v", want, err)
	}
}

func TestRejectUser(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/reject", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
	})

	err := client.Users.RejectUser(1)
	if err != nil {
		t.Errorf("Users.RejectUser returned error: %v", err)
	}
}

func TestRejectUser_UserNotFound(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/reject", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.RejectUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.RejectUser error.\nExpected: %v\nGot: %v", ErrUserNotFound, err)
	}
}

func TestRejectUser_RejectPrevented(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/reject", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.RejectUser(1)
	if !errors.Is(err, ErrUserRejectPrevented) {
		t.Errorf("Users.RejectUser error.\nExpected: %v\nGot: %v", ErrUserRejectPrevented, err)
	}
}

func TestRejectUser_Conflict(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/reject", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusConflict)
	})

	err := client.Users.RejectUser(1)
	if !errors.Is(err, ErrUserConflict) {
		t.Errorf("Users.RejectUser error.\nExpected: %v\nGot: %v", ErrUserConflict, err)
	}
}

func TestRejectUser_UnknownError(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/reject", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusTeapot)
	})

	want := fmt.Sprintf("Received unexpected result code: %d", http.StatusTeapot)

	err := client.Users.RejectUser(1)
	if err.Error() != want {
		t.Errorf("Users.RejectUser error.\nExpected: %s\n\tGot: %v", want, err)
	}
}

func TestGetMemberships(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/memberships", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_user_memberships.json")
	})

	opt := new(GetUserMembershipOptions)

	memberships, _, err := client.Users.GetUserMemberships(1, opt)
	require.NoError(t, err)

	want := []*UserMembership{{SourceID: 1, SourceName: "Project one", SourceType: "Project", AccessLevel: 20}, {SourceID: 3, SourceName: "Group three", SourceType: "Namespace", AccessLevel: 20}}
	assert.Equal(t, want, memberships)
}
