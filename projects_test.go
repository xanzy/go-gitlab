package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListProjects(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":             "2",
			"per_page":         "3",
			"archived":         "true",
			"order_by":         "name",
			"sort":             "asc",
			"search":           "query",
			"ci_enabled_first": "true",
			"simple":           "true",
			"visibility":       "public",
		})
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{ListOptions{2, 3}, Bool(true), String("name"), String("asc"), String("query"), Bool(true), Bool(true), String("public")}
	projects, _, err := client.Projects.ListProjects(opt)

	if err != nil {
		t.Errorf("Projects.ListProjects returned error: %v", err)
	}

	want := []*Project{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Projects.ListProjects returned %+v, want %+v", projects, want)
	}
}

func TestListOwnedProjects(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/owned", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":             "2",
			"per_page":         "3",
			"archived":         "true",
			"order_by":         "name",
			"sort":             "asc",
			"search":           "query",
			"ci_enabled_first": "true",
			"simple":           "true",
			"visibility":       "public",
		})
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{ListOptions{2, 3}, Bool(true), String("name"), String("asc"), String("query"), Bool(true), Bool(true), String("public")}
	projects, _, err := client.Projects.ListOwnedProjects(opt)

	if err != nil {
		t.Errorf("Projects.ListOwnedProjects returned error: %v", err)
	}

	want := []*Project{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Projects.ListOwnedProjects returned %+v, want %+v", projects, want)
	}
}

func TestListStarredProjects(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/starred", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":             "2",
			"per_page":         "3",
			"archived":         "true",
			"order_by":         "name",
			"sort":             "asc",
			"search":           "query",
			"ci_enabled_first": "true",
			"simple":           "true",
			"visibility":       "public",
		})
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{ListOptions{2, 3}, Bool(true), String("name"), String("asc"), String("query"), Bool(true), Bool(true), String("public")}
	projects, _, err := client.Projects.ListStarredProjects(opt)

	if err != nil {
		t.Errorf("Projects.ListStarredProjects returned error: %v", err)
	}

	want := []*Project{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Projects.ListStarredProjects returned %+v, want %+v", projects, want)
	}
}

func TestListAllProjects(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/all", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":             "2",
			"per_page":         "3",
			"archived":         "true",
			"order_by":         "name",
			"sort":             "asc",
			"search":           "query",
			"ci_enabled_first": "true",
			"simple":           "true",
			"visibility":       "public",
		})
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectsOptions{ListOptions{2, 3}, Bool(true), String("name"), String("asc"), String("query"), Bool(true), Bool(true), String("public")}
	projects, _, err := client.Projects.ListAllProjects(opt)

	if err != nil {
		t.Errorf("Projects.ListAllProjects returned error: %v", err)
	}

	want := []*Project{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Projects.ListAllProjects returned %+v, want %+v", projects, want)
	}
}

func TestGetProject_byID(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &Project{ID: 1}

	project, _, err := client.Projects.GetProject(1)

	if err != nil {
		t.Fatalf("Projects.GetProject returns an error: %v", err)
	}

	if !reflect.DeepEqual(want, project) {
		t.Errorf("Projects.GetProject returned %+v, want %+v", project, want)
	}
}

func TestGetProject_byName(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/", func(w http.ResponseWriter, r *http.Request) {
		testUrl(t, r, "/projects/namespace%2Fname")
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &Project{ID: 1}

	project, _, err := client.Projects.GetProject("namespace/name")

	if err != nil {
		t.Fatalf("Projects.GetProject returns an error: %v", err)
	}

	if !reflect.DeepEqual(want, project) {
		t.Errorf("Projects.GetProject returned %+v, want %+v", project, want)
	}
}

func TestSearchProjects(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/search/query", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "2",
			"per_page": "3",
			"order_by": "name",
			"sort":     "asc",
		})
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &SearchProjectsOptions{ListOptions{2, 3}, String("name"), String("asc")}
	projects, _, err := client.Projects.SearchProjects("query", opt)

	if err != nil {
		t.Errorf("Projects.SearchProjects returned error: %v", err)
	}

	want := []*Project{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Projects.SearchProjects returned %+v, want %+v", projects, want)
	}
}

func TestCreateProject(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testJsonBody(t, r, values{
			"name": "n",
		})

		fmt.Fprint(w, `{"id":1}`)
	})

	opt := &CreateProjectOptions{Name: String("n")}
	project, _, err := client.Projects.CreateProject(opt)

	if err != nil {
		t.Errorf("Projects.CreateProject returned error: %v", err)
	}

	want := &Project{ID: 1}
	if !reflect.DeepEqual(want, project) {
		t.Errorf("Projects.CreateProject returned %+v, want %+v", project, want)
	}
}
