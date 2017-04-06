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

	opt := &MoveIssueOptions{ID: Int(1), IssueIID: Int(1), ToProjectID: Int(2)}
	issue, _, err := client.Issues.MoveIssue(1, 1, opt)
	if err != nil {
		t.Errorf("Issues.MoveIssue returned error: %v", err)
	}

	want := &Issue{IID: 1, ProjectID: 2}
	if !reflect.DeepEqual(want, issue) {
		t.Errorf("Issues.MoveIssue returned %+v, want %+v", issue, want)
	}
}

func TestSubscribeIssue(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/issues/1/subscribe", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"id":1, "project_id":1}`)
	})

	opt := &SubscribeIssueOptions{ID: Int(1), IssueIID: Int(1)}
	issue, _, err := client.Issues.SubscribeIssue(1, 1, opt)
	if err != nil {
		t.Errorf("Issues.SubscribeIssue returned error: %v", err)
	}

	want := &Issue{ID: 1, ProjectID: 1}
	if !reflect.DeepEqual(want, issue) {
		t.Errorf("Issues.SubscribeIssue returned %+v, want %+v", issue, want)
	}
}

func TestUnsubscribeIssue(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/issues/1/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"id":1, "project_id":1}`)
	})

	opt := &UnsubscribeIssueOptions{ID: Int(1), IssueIID: Int(1)}
	issue, _, err := client.Issues.UnsubscribeIssue(1, 1, opt)
	if err != nil {
		t.Errorf("Issues.UnsubscribeIssue returned error: %v", err)
	}

	want := &Issue{ID: 1, ProjectID: 1}
	if !reflect.DeepEqual(want, issue) {
		t.Errorf("Issues.UnsubscribeIssue returned %+v, want %+v", issue, want)
	}
}
