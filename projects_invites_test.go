package gitlab

import (
    "fmt"
    "net/http"
    "reflect"
    "testing"
)

func TestListProjectPendingInvites(t *testing.T) {
    mux, server, client := setup(t)
    defer teardown(server)

    mux.HandleFunc("/api/v4/projects/test/invitations", func(w http.ResponseWriter, r *http.Request) {
        testMethod(t, r, "GET")
        fmt.Fprint(w, `[{"id":1},{"id":2}]`)
    })

    opt := &ListPendingInvitationsOptions{
        ListOptions: ListOptions{2, 3},
    }

    projects, _, err := client.ProjectInvites.ListPendingInvitations("test", opt)
    if err != nil {
        t.Errorf("ProjectInvites.ListPendingInvitations returned error: %v", err)
    }

    want := []*PendingInvitations{{Id: 1}, {Id: 2}}
    if !reflect.DeepEqual(want, projects) {
        t.Errorf("ProjectInvites.ListPendingInvitations returned %+v, want %+v", projects, want)
    }
}

func TestInvitesProjects(t *testing.T) {
    mux, server, client := setup(t)
    defer teardown(server)

    mux.HandleFunc("/api/v4/projects/test/invitations", func(w http.ResponseWriter, r *http.Request) {
        testMethod(t, r, "POST")
        fmt.Fprint(w, `{"status": "success"}`)
    })

    opt := &InvitesOptions{
        Email: "example@member.org",
    }

    projects, _, err := client.ProjectInvites.ProjectInvites("test", *opt)
    if err != nil {
        t.Errorf("Projects.ListUserProjects returned error: %v", err)
    }

    want := &InvitationsResponse{Status: "success"}
    if !reflect.DeepEqual(want, projects) {
        t.Errorf("ProjectInvites.ProjectInvites returned %+v, want %+v", projects, want)
    }
}

func TestInvitesProjectsError(t *testing.T) {
    mux, server, client := setup(t)
    defer teardown(server)

    mux.HandleFunc("/api/v4/projects/test/invitations", func(w http.ResponseWriter, r *http.Request) {
        testMethod(t, r, "POST")
        fmt.Fprint(w, `{"status": "error","message": {"example@member.org": "Already invited"}}`)
    })

    opt := &InvitesOptions{
        Email: "example@member.org",
    }

    projects, _, err := client.ProjectInvites.ProjectInvites("test", *opt)
    if err != nil {
        t.Errorf("Projects.ListUserProjects returned error: %v", err)
    }

    want := &InvitationsResponse{Status: "error", Message: map[string]string{"example@member.org": "Already invited"}}
    if !reflect.DeepEqual(want, projects) {
        t.Errorf("ProjectInvites.ProjectInvites returned %+v, want %+v", projects, want)
    }
}
