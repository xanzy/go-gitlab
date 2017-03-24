package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const (
	myKey      = "MY_KEY"
	myValue    = "MY_VALUE"
	myKey2     = "MY_KEY2"
	myValue2   = "MY_VALUE2"
	myNewValue = "MY_NEW_VALUE"
)

func TestListBuildVariables(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/variables", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w,
			`[{"key":"%s","value":"%s"},{"key":"%s","value":"%s"}]`, myKey, myValue, myKey2, myValue2)
	})

	variables, _, err := client.BuildVariables.ListBuildVariables(1, nil)
	if err != nil {
		t.Errorf("ListBuildVariables returned error: %v", err)
	}

	want := []*BuildVariable{{Key: myKey, Value: myValue}, {Key: myKey2, Value: myValue2}}
	if !reflect.DeepEqual(want, variables) {
		t.Errorf("ListBuildVariables returned %+v, want %+v", variables, want)
	}
}

func TestGetBuildVariable(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/variables/"+myKey, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `{"key":"%s","value":"%s"}`, myKey, myValue)
	})

	variable, _, err := client.BuildVariables.GetBuildVariable(1, myKey)
	if err != nil {
		t.Errorf("GetBuildVariable returned error: %v", err)
	}

	want := &BuildVariable{Key: myKey, Value: myValue}
	if !reflect.DeepEqual(want, variable) {
		t.Errorf("GetBuildVariable returned %+v, want %+v", variable, want)
	}
}

func TestCreateBuildVariable(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/variables", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testJSONBody(t, r, values{
			"key":   myKey,
			"value": myValue,
		})
		fmt.Fprintf(w, `{"key":"%s","value":"%s"}`, myKey, myValue)
	})

	variable, _, err := client.BuildVariables.CreateBuildVariable(1, myKey, myValue)
	if err != nil {
		t.Errorf("CreateBuildVariable returned error: %v", err)
	}

	want := &BuildVariable{Key: myKey, Value: myValue}
	if !reflect.DeepEqual(want, variable) {
		t.Errorf("CreateBuildVariable returned %+v, want %+v", variable, want)
	}
}

func TestUpdateBuildVariable(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/variables/"+myKey, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testJSONBody(t, r, values{
			"key":   myKey,
			"value": myNewValue,
		})
		fmt.Fprintf(w, `{"key":"%s","value":"%s"}`, myKey, myNewValue)
	})

	variable, _, err := client.BuildVariables.UpdateBuildVariable(1, myKey, myNewValue)
	if err != nil {
		t.Errorf("UpdateBuildVariable returned error: %v", err)
	}

	want := &BuildVariable{Key: myKey, Value: myNewValue}
	if !reflect.DeepEqual(want, variable) {
		t.Errorf("UpdateBuildVariable returned %+v, want %+v", variable, want)
	}
}

func TestRemoveBuildVariable(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/variables/"+myKey, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.BuildVariables.RemoveBuildVariable(1, myKey)
	if err != nil {
		t.Errorf("RemoveBuildVariable returned error: %v", err)
	}
}
