package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListMergeRequestsNullApprovals(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/merge_requests", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1, "title": "some title", "approvals_before_merge": null}]`)
	})

	mr, _, err := client.MergeRequests.ListMergeRequests(nil)

	if err != nil {
		t.Errorf("MergeRequests.ListMergeRequests returned error: %v", err)
	}

	want := []*MergeRequest{{ID: 1, Title: "some title", ApprovalsBeforeMerge: 0}}
	if !reflect.DeepEqual(want, mr) {
		t.Errorf("MergeRequests.ListMergeRequests returned %+v, want %+v", mr, want)
	}
}

func TestListMergeRequests2Approvals(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/merge_requests", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1, "title": "some title", "approvals_before_merge": 2}]`)
	})

	mr, _, err := client.MergeRequests.ListMergeRequests(nil)

	if err != nil {
		t.Errorf("MergeRequests.ListMergeRequests returned error: %v", err)
	}

	want := []*MergeRequest{{ID: 1, Title: "some title", ApprovalsBeforeMerge: 2}}
	if !reflect.DeepEqual(want, mr) {
		t.Errorf("MergeRequests.ListMergeRequests returned %+v, want %+v", mr, want)
	}
}

func TestListMergeRequestsError(t *testing.T) {
	_, server, client := setup()
	defer teardown(server)

	_, _, err := client.MergeRequests.ListMergeRequests(nil)

	if err == nil {
		t.Errorf("MergeRequests.ListMergeRequests expected to receive error, but returned nil")
	}
}
