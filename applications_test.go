package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListApplications(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/applications",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `[{"id":1},{"id":2}]`)
		})

	apps, _, err := client.Applications.ListApplications()
	if err != nil {
		t.Errorf("Applications.ListApplications returned error: %v", err)
	}

	want := []*Application{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, apps) {
		t.Errorf("Groups.ListGroups returned %+v, want %+v", apps, want)
	}
}
