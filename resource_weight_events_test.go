package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestResourceWeightEventsService_ListIssueWightEvents(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/issues/11/resource_weight_events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			  "id": 142,
			  "user": {
				"id": 1,
				"name": "Administrator",
				"username": "root",
				"state": "active",
				"avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				"web_url": "http://gitlab.example.com/root"
			  },
			  "created_at": "2018-08-20T13:38:20.077Z",
			  "issue_id": 253,
			  "weight": 3
			}
		]`)
	})

	opt := &ListWeightEventsOptions{ListOptions{Page: 1, PerPage: 10}}

	wes, _, err := client.ResourceWeightEvents.ListIssueWeightEvents(5, 11, opt)
	require.NoError(t, err)

	want := []*WeightEvent{{
		ID: 142,
		User: &BasicUser{
			ID:        1,
			Username:  "root",
			Name:      "Administrator",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/root",
		},
		CreatedAt: Time(time.Date(2018, time.August, 20, 13, 38, 20, 77000000, time.UTC)),
		IssueID:   253,
		Weight:    3,
	}}
	require.Equal(t, want, wes)
}
