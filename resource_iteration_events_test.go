// Copyright 2023, Hakki Ceylan, Yavuz Turk
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestListIssueIterationEventsService_ListIssueIterationEvents(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/issues/11/resource_iteration_events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			    "id": 142,
			    "user": {
			        "id": 1,
			        "username": "root",
			        "name": "Administrator",
			        "state": "active",
			        "avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			        "web_url": "https://gitlab.example.com/root"
			    },
			    "created_at": "2023-09-22T06:51:04.801Z",
			    "resource_type": "Issue",
			    "resource_id": 11,
			    "iteration": {
			        "id": 133,
			        "iid": 1,
			        "sequence": 1,
			        "group_id": 153,
			        "title": "Iteration 1",
			        "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			        "state": 1,
			        "created_at": "2023-07-15T00:05:06.509Z",
			        "updated_at": "2023-09-24T00:05:10.476Z",
			        "start_date": "2023-09-17",
			        "due_date": "2023-09-23",
			        "web_url": "https://gitlab.example.com/groups/project/-/iterations/1"
			    },
			    "action": "add"
			}
		]`)
	})

	opt := &ListIterationEventsOptions{ListOptions{Page: 1, PerPage: 10}}

	mes, _, err := client.ResourceIterationEvents.ListIssueIterationEvents(5, 11, opt)
	require.NoError(t, err)

	startDateISOTime, err := ParseISOTime("2023-09-17")
	require.NoError(t, err)

	dueDateISOTime, err := ParseISOTime("2023-09-23")
	require.NoError(t, err)

	want := []*IterationEvent{{
		ID: 142,
		User: &BasicUser{
			ID:        1,
			Username:  "root",
			Name:      "Administrator",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "https://gitlab.example.com/root",
		},
		ResourceType: "Issue",
		ResourceID:   11,
		CreatedAt:    Time(time.Date(2023, time.September, 22, 06, 51, 04, 801000000, time.UTC)),
		Iteration: &ProjectIssueIteration{
			Id:          133,
			Iid:         1,
			Sequence:    1,
			GroupId:     153,
			Title:       "Iteration 1",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			State:       1,
			CreatedAt:   Time(time.Date(2023, time.July, 15, 00, 05, 06, 509000000, time.UTC)),
			UpdatedAt:   Time(time.Date(2023, time.September, 24, 00, 05, 10, 476000000, time.UTC)),
			StartDate:   &startDateISOTime,
			DueDate:     &dueDateISOTime,
			WebUrl:      "https://gitlab.example.com/groups/project/-/iterations/1",
		},
		Action: "add",
	}}
	require.Equal(t, want, mes)
}

func TestListIssueIterationEventsService_GetIssueIterationEvent(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/issues/11/resource_iteration_events/143", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
			    "id": 143,
			    "user": {
			        "id": 1,
			        "username": "root",
			        "name": "Administrator",
			        "state": "active",
			        "avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			        "web_url": "https://gitlab.example.com/root"
			    },
			    "created_at": "2023-09-22T06:51:04.801Z",
			    "resource_type": "Issue",
			    "resource_id": 11,
			    "iteration": {
			        "id": 133,
			        "iid": 1,
			        "sequence": 1,
			        "group_id": 153,
			        "title": "Iteration 1",
			        "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			        "state": 1,
			        "created_at": "2023-07-15T00:05:06.509Z",
			        "updated_at": "2023-09-24T00:05:10.476Z",
			        "start_date": "2023-09-17",
			        "due_date": "2023-09-23",
			        "web_url": "https://gitlab.example.com/groups/project/-/iterations/2"
			    },
			    "action": "add"
			}`,
		)
	})

	me, _, err := client.ResourceIterationEvents.GetIssueIterationEvent(5, 11, 143)
	require.NoError(t, err)

	startDateISOTime, err := ParseISOTime("2023-09-17")
	require.NoError(t, err)

	dueDateISOTime, err := ParseISOTime("2023-09-23")
	require.NoError(t, err)

	want := &IterationEvent{
		ID: 143,
		User: &BasicUser{
			ID:        1,
			Username:  "root",
			Name:      "Administrator",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "https://gitlab.example.com/root",
		},
		ResourceType: "Issue",
		ResourceID:   11,
		CreatedAt:    Time(time.Date(2023, time.September, 22, 06, 51, 04, 801000000, time.UTC)),
		Iteration: &ProjectIssueIteration{
			Id:          133,
			Iid:         1,
			Sequence:    1,
			GroupId:     153,
			Title:       "Iteration 1",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			State:       1,
			CreatedAt:   Time(time.Date(2023, time.July, 15, 00, 05, 06, 509000000, time.UTC)),
			UpdatedAt:   Time(time.Date(2023, time.September, 24, 00, 05, 10, 476000000, time.UTC)),
			StartDate:   &startDateISOTime,
			DueDate:     &dueDateISOTime,
			WebUrl:      "https://gitlab.example.com/groups/project/-/iterations/2",
		},
		Action: "add",
	}
	require.Equal(t, want, me)
}
