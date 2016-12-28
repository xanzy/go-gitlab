package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListVariables(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/variables", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"key":"MY_KEY", "value": "MY_VALUE"},{"key":"MY_KEY2", "value": "MY_VALUE2"}]`)
	})

	variables, _, err := client.Variables.ListVariables(1)

	if err != nil {
		t.Errorf("Projects.ListVariables returned error: %v", err)
	}

	want := []*Variable{{Key: "MY_KEY", Value: "MY_VALUE"}, {Key: "MY_KEY2", Value: "MY_VALUE2"}}
	if !reflect.DeepEqual(want, variables) {
		t.Errorf("Projects.ListVariables returned %+v, want %+v", variables, want)
	}
}

func TestGetSingleVariable(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/variables/MY_KEY", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"key":"MY_KEY", "value": "MY_VALUE"}`)
	})

	variable, _, err := client.Variables.GetSingleVariable(1, "MY_KEY")

	if err != nil {
		t.Errorf("Projects.GetSingleVariable returned error: %v", err)
	}

	want := &Variable{Key: "MY_KEY", Value: "MY_VALUE"}
	if !reflect.DeepEqual(want, variable) {
		t.Errorf("Projects.ListVariables returned %+v, want %+v", variable, want)
	}
}

func testCreateVariable(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/variables", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testJsonBody(t, r, values{
			"key":   "MY_KEY",
			"value": "MY_VALUE",
		})
		fmt.Fprint(w, `{"key":"MY_KEY", "value": "MY_VALUE"}`)
	})
	want := &Variable{Key: "MY_KEY", Value: "MY_VALUE"}
	variable, _, err := client.Variables.CreateVariable(1, *want)

	if err != nil {
		t.Errorf("Projects.CreateVariable returned error: %v", err)
	}

	if !reflect.DeepEqual(want, variable) {
		t.Errorf("Projects.CreateVariable returned %+v, want %+v", variable, want)
	}
}

func testUpdateVariable(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/variables/MY_KEY", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testJsonBody(t, r, values{
			"key":   "MY_KEY",
			"value": "MY_NEW_VALUE",
		})
		fmt.Fprint(w, `{"key":"MY_KEY", "value": "MY_NEW_VALUE"}`)
	})
	want := &Variable{Key: "MY_KEY", Value: "MY_NEW_VALUE"}
	variable, _, err := client.Variables.UpdateVariable(1, "MY_KEY", *want)

	if err != nil {
		t.Errorf("Projects.UpdateVariable returned error: %v", err)
	}

	if !reflect.DeepEqual(want, variable) {
		t.Errorf("Projects.UpdateVariable returned %+v, want %+v", variable, want)
	}
}
