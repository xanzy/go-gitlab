package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListGroupPendingInvites(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/test/invitations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListPendingInvitationsOptions{
		ListOptions: ListOptions{2, 3},
	}

	projects, _, err := client.Invites.ListPendingGroupInvitations("test", opt)
	if err != nil {
		t.Errorf("Invites.ListPendingGroupInvitations returned error: %v", err)
	}

	want := []*PendingInvite{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Invites.ListPendingGroupInvitations returned %+v, want %+v", projects, want)
	}
}

func TestGroupInvites(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/test/invitations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"status": "success"}`)
	})

	opt := &InvitesOptions{
		Email: String("example@member.org"),
	}

	projects, _, err := client.Invites.GroupInvites("test", opt)
	if err != nil {
		t.Errorf("Invites.GroupInvites returned error: %v", err)
	}

	want := &InvitesResult{Status: "success"}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Invites.GroupInvites returned %+v, want %+v", projects, want)
	}
}

func TestGroupInvitesError(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/test/invitations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"status": "error","message": {"example@member.org": "Already invited"}}`)
	})

	opt := &InvitesOptions{
		Email: String("example@member.org"),
	}

	projects, _, err := client.Invites.GroupInvites("test", opt)
	if err != nil {
		t.Errorf("Invites.GroupInvites returned error: %v", err)
	}

	want := &InvitesResult{Status: "error", Message: map[string]string{"example@member.org": "Already invited"}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Invites.GroupInvites returned %+v, want %+v", projects, want)
	}
}

func TestListProjectPendingInvites(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/test/invitations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListPendingInvitationsOptions{
		ListOptions: ListOptions{2, 3},
	}

	projects, _, err := client.Invites.ListPendingProjectInvitations("test", opt)
	if err != nil {
		t.Errorf("Invites.ListPendingProjectInvitations returned error: %v", err)
	}

	want := []*PendingInvite{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Invites.ListPendingProjectInvitations returned %+v, want %+v", projects, want)
	}
}

func TestProjectInvites(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/test/invitations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"status": "success"}`)
	})

	opt := &InvitesOptions{
		Email: String("example@member.org"),
	}

	projects, _, err := client.Invites.ProjectInvites("test", opt)
	if err != nil {
		t.Errorf("Invites.ProjectInvites returned error: %v", err)
	}

	want := &InvitesResult{Status: "success"}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Invites.ProjectInvites returned %+v, want %+v", projects, want)
	}
}

func TestProjectInvitesError(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/test/invitations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"status": "error","message": {"example@member.org": "Already invited"}}`)
	})

	opt := &InvitesOptions{
		Email: String("example@member.org"),
	}

	projects, _, err := client.Invites.ProjectInvites("test", opt)
	if err != nil {
		t.Errorf("Invites.ProjectInvites returned error: %v", err)
	}

	want := &InvitesResult{Status: "error", Message: map[string]string{"example@member.org": "Already invited"}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Invites.ProjectInvites returned %+v, want %+v", projects, want)
	}
}
