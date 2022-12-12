package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListProjectIterations(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/42/iterations",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprintf(w, `[
        {
          "id": 53,
          "iid": 13,
          "sequence": 1,
          "group_id": 5,
          "title": "Iteration II",
          "description": "Ipsum Lorem ipsum",
          "state": 2,
          "web_url": "http://gitlab.example.com/groups/my-group/-/iterations/13"
        }
      ]`)
		})

	iterations, _, err := client.ProjectIterations.ListProjectIterations(42, &ListProjectIterationsOptions{})
	if err != nil {
		t.Errorf("GroupIterations.ListGroupIterations returned error: %v", err)
	}

	want := []*ProjectIteration{{
		ID:          53,
		IID:         13,
		Sequence:    1,
		GroupID:     5,
		Title:       "Iteration II",
		Description: "Ipsum Lorem ipsum",
		State:       2,
		WebURL:      "http://gitlab.example.com/groups/my-group/-/iterations/13",
	}}
	if !reflect.DeepEqual(want, iterations) {
		t.Errorf("ProjectIterations.ListProjectIterations returned %+v, want %+v", iterations, want)
	}
}
