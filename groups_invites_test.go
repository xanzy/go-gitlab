package gitlab

import (
    "fmt"
    "net/http"
    "reflect"
    "testing"
)

func TestListGroupPendingInvites(t *testing.T) {
    mux, server, client := setup(t)
    defer teardown(server)

    mux.HandleFunc("/api/v4/groups/test/invitations", func(w http.ResponseWriter, r *http.Request) {
        testMethod(t, r, "GET")
        fmt.Fprint(w, `[{"id":1},{"id":2}]`)
    })

    opt := &ListPendingInvitationsOptions{
        ListOptions: ListOptions{2, 3},
    }

    projects, _, err := client.GroupInvites.ListPendingInvitations("test", opt)
    if err != nil {
        t.Errorf("GroupInvites.ListPendingInvitations returned error: %v", err)
    }

    want := []*PendingInvitations{{Id: 1}, {Id: 2}}
    if !reflect.DeepEqual(want, projects) {
        t.Errorf("GroupInvites.ListPendingInvitations returned %+v, want %+v", projects, want)
    }
}

func TestInvitesGroups(t *testing.T) {
    mux, server, client := setup(t)
    defer teardown(server)

    mux.HandleFunc("/api/v4/groups/test/invitations", func(w http.ResponseWriter, r *http.Request) {
        testMethod(t, r, "POST")
        fmt.Fprint(w, `{"status": "success"}`)
    })

    opt := &InvitesOptions{
        Email: "example@member.org",
    }

    projects, _, err := client.GroupInvites.GroupInvites("test", *opt)
    if err != nil {
        t.Errorf("Groups.ListUserGroups returned error: %v", err)
    }

    want := &InvitationsResponse{Status: "success"}
    if !reflect.DeepEqual(want, projects) {
        t.Errorf("GroupInvites.GroupInvites returned %+v, want %+v", projects, want)
    }
}

func TestInvitesGroupsError(t *testing.T) {
    mux, server, client := setup(t)
    defer teardown(server)

    mux.HandleFunc("/api/v4/groups/test/invitations", func(w http.ResponseWriter, r *http.Request) {
        testMethod(t, r, "POST")
        fmt.Fprint(w, `{"status": "error","message": {"example@member.org": "Already invited"}}`)
    })

    opt := &InvitesOptions{
        Email: "example@member.org",
    }

    projects, _, err := client.GroupInvites.GroupInvites("test", *opt)
    if err != nil {
        t.Errorf("Groups.ListUserGroups returned error: %v", err)
    }

    want := &InvitationsResponse{Status: "error", Message: map[string]string{ "example@member.org": "Already invited" }}
    if !reflect.DeepEqual(want, projects) {
        t.Errorf("GroupInvites.GroupInvites returned %+v, want %+v", projects, want)
    }
}
