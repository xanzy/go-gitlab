package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListFeatureFlags(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/features", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
		[
			{
			  "name": "experimental_feature",
			  "state": "off",
			  "gates": [
				{
				  "key": "boolean",
				  "value": false
				}
			  ]
			},
			{
			  "name": "new_library",
			  "state": "on"
			}
		  ]
	`)
	})

	features, _, err := client.Features.ListFeatures()
	if err != nil {
		t.Errorf("Features.ListFeatures returned error: %v", err)
	}

	want := []*Feature{
		{Name: "experimental_feature", State: "off", Gates: []Gate{
			{Key: "boolean", Value: false},
		}},
		{Name: "new_library", State: "on"},
	}
	if !reflect.DeepEqual(want, features) {
		t.Errorf("Features.ListFeatures returned %+v, want %+v", features, want)
	}
}
