package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListGroups(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
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
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/g",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `{"id": 1, "name": "g"}`)
		})

	group, _, err := client.Groups.GetGroup("g")
	if err != nil {
		t.Errorf("Groups.GetGroup returned error: %v", err)
	}

	want := &Group{ID: 1, Name: "g"}
	if !reflect.DeepEqual(want, group) {
		t.Errorf("Groups.GetGroup returned %+v, want %+v", group, want)
	}
}

func TestCreateGroup(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
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
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/1/projects/2",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
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

func TestDeleteGroup(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
			w.WriteHeader(http.StatusAccepted)
		})

	resp, err := client.Groups.DeleteGroup(1)
	if err != nil {
		t.Errorf("Groups.DeleteGroup returned error: %v", err)
	}
	fmt.Println(resp.StatusCode)

	want := http.StatusAccepted
	got := resp.StatusCode
	if got != want {
		t.Errorf("Groups.DeleteGroup returned %d, want %d", got, want)
	}
}

func TestSearchGroup(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
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
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "PUT")
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
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/22/projects",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
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

func TestListSubgroups(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/1/subgroups",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `[{"id": 1}, {"id": 2}]`)
		})

	groups, _, err := client.Groups.ListSubgroups(1, &ListSubgroupsOptions{})
	if err != nil {
		t.Errorf("Groups.ListSubgroups returned error: %v", err)
	}

	want := []*Group{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, groups) {
		t.Errorf("Groups.ListSubgroups returned %+v, want %+v", groups, want)
	}
}
