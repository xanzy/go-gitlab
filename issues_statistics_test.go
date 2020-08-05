package gitlab

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestGetIssuesStatistics(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/issues_statistics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testURL(t, r, "/api/v4/issues_statistics?assignee_id=1&author_id=1")
		fmt.Fprint(w, `{"statistics": {"counts": {"all": 20,"closed": 5,"opened": 15}}}`)
	})

	opt := &GetIssuesStatisticsOptions{
		AssigneeID: Int(1),
		AuthorID:   Int(1),
	}

	issue, _, err := client.IssuesStatistics.GetIssuesStatistics(opt)
	if err != nil {
		log.Fatal(err)
	}

	want := &IssuesStatistics{
		Statistics: Statistics{
			Counts: Counts{
				All:    20,
				Closed: 5,
				Opened: 15,
			},
		},
	}

	if !reflect.DeepEqual(want, issue) {
		t.Errorf("Issues.GetIssue returned %+v, want %+v", issue, want)
	}
}
