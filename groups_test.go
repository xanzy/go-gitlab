package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListGroupProjects(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/22/projects",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `[{"id":1},{"id":2}]`)
		})

	opt := &ListGroupProjectsOptions{
		ListOptions: ListOptions{2, 3},
		Archived:    Bool(true),
		OrderBy:     String("name"),
		Sort:        String("asc"),
		Search:      String("query"),
		Simple:      Bool(true),
		Visibility:  Visibility(PublicVisibility),
	}

	projects, _, err := client.Groups.ListGroupProjects(22, opt)
	if err != nil {
		t.Errorf("Groups.ListGroupProjects returned error: %v", err)
	}

	want := []*Project{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Groups.ListGroupProjects returned %+v, want %+v", projects, want)
	}
}
