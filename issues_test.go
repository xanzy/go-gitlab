package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMoveIssue(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/issues/1/move", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"iid":1, "project_id":2}`)
	})

	opt := &MoveIssueOptions{IssueIID: Int(1), ToProjectID: Int(2)}
	issue, _, err := client.Issues.MoveIssue(1, 1, opt)

	if err != nil {
		t.Errorf("Issues.MoveIssue returned error: %v", err)
	}

	want := &Issue{IID: 1, ProjectID: 2}
	if !reflect.DeepEqual(want, issue) {
		t.Errorf("Issues.MoveIssue returned %+v, want %+v", issue, want)
	}
}
