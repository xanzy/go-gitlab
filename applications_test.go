package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateApplication(t *testing.T) {
	m, s, c := setup()
	defer teardown(s)

	m.HandleFunc("/api/v4/applications",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			fmt.Fprint(w, `{"id":1, "application_name":"testApplication"}`)
		},
	)

	posted := CreateApplicationOptions{
		Name: "testApplication",
	}
	a, _, err := c.Applications.CreateApplication(&posted)
	if err != nil {
		t.Errorf("Applications.CreateApplication returned error : %v", err)
	}
	want := Application{ID: 1, ApplicationName: "testApplication"}
	if reflect.DeepEqual(want, a) {
		t.Errorf("Applications.CreateApplication returned %v, want %v", a, want)
	}
}

func TestListApplications(t *testing.T) {
	m, s, c := setup()
	defer teardown(s)

	m.HandleFunc("/api/v4/applications",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `[{"id":1},{"id":2}]`)
		},
	)

	apps, _, err := c.Applications.ListApplications()
	if err != nil {
		t.Errorf("Applications.ListApplications returned error: %v", err)
	}

	want := []*Application{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, apps) {
		t.Errorf("Applications.ListApplications returned %+v, want %+v", apps, want)
	}
}

func TestDeleteApplication(t *testing.T) {
	m, s, c := setup()
	defer teardown(s)

	m.HandleFunc("/api/v4/applications/4",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
			w.WriteHeader(http.StatusAccepted)
		},
	)

	r, err := c.Applications.DeleteApplication(4)
	if err != nil {
		t.Errorf("Applications.DeleteApplication returned error : %v", err)
	}

	want := http.StatusAccepted
	got := r.StatusCode
	if got != want {
		t.Errorf("Applications.DeleteApplication returned %d, want %d", got, want)
	}
}
