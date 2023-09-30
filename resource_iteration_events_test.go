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

func TestListIssueIterationEvents(t *testing.T) {
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
			        "web_url": "http://gitlab.example.com/root"
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
			        "web_url": ""
			    },
			    "action": "add"
			}
		]`)
	})

	opt := &ListIterationEventsOptions{ListOptions{Page: 1, PerPage: 10}}

	mes, _, err := client.ResourceIterationEvents.ListIssueIterationEvents(5, 11, opt)
	require.NoError(t, err)

	eventCreatedAt, err := time.Parse(time.RFC3339, "2023-09-22T06:51:04.801Z")
	require.NoError(t, err)

	createdAt, err := time.Parse(time.RFC3339, "2023-07-15T00:05:06.509Z")
	require.NoError(t, err)

	updatedAt, err := time.Parse(time.RFC3339, "2023-09-24T00:05:10.476Z")
	require.NoError(t, err)

	startDate, err := time.Parse(time.RFC3339, "2023-09-17T00:00:00.000Z")
	require.NoError(t, err)
	startDateISOTime := ISOTime(startDate)

	dueDate, err := time.Parse(time.RFC3339, "2023-09-23T00:00:00.000Z")
	require.NoError(t, err)
	dueDateISOTime := ISOTime(dueDate)

	want := []*IterationEvent{{
		ID: 142,
		User: &BasicUser{
			ID:        1,
			Username:  "root",
			Name:      "Administrator",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/root",
		},
		ResourceType: "Issue",
		ResourceID:   11,
		CreatedAt:    &eventCreatedAt,
		Iteration: &ProjectIssueIteration{
			Id:          133,
			Iid:         1,
			Sequence:    1,
			GroupId:     153,
			Title:       "Iteration 1",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			State:       1,
			CreatedAt:   &createdAt,
			UpdatedAt:   &updatedAt,
			StartDate:   &startDateISOTime,
			DueDate:     &dueDateISOTime,
			WebUrl:      "",
		},
		Action: "add",
	}}
	require.Equal(t, want, mes)
}

func TestGetIssueIterationEvent(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/issues/11/resource_iteration_events/143", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
						{
			    "id": 142,
			    "user": {
			        "id": 1,
			        "username": "root",
			        "name": "Administrator",
			        "state": "active",
			        "avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			        "web_url": "http://gitlab.example.com/root"
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
			        "web_url": ""
			    },
			    "action": "add"
			}`,
		)
	})

	me, _, err := client.ResourceIterationEvents.GetIssueIterationEvent(5, 11, 143)
	require.NoError(t, err)

	eventCreatedAt, err := time.Parse(time.RFC3339, "2023-09-22T06:51:04.801Z")
	require.NoError(t, err)

	createdAt, err := time.Parse(time.RFC3339, "2023-07-15T00:05:06.509Z")
	require.NoError(t, err)

	updatedAt, err := time.Parse(time.RFC3339, "2023-09-24T00:05:10.476Z")
	require.NoError(t, err)

	startDate, err := time.Parse(time.RFC3339, "2023-09-17T00:00:00.000Z")
	require.NoError(t, err)
	startDateISOTime := ISOTime(startDate)

	dueDate, err := time.Parse(time.RFC3339, "2023-09-23T00:00:00.000Z")
	require.NoError(t, err)
	dueDateISOTime := ISOTime(dueDate)

	want := &IterationEvent{
		ID: 142,
		User: &BasicUser{
			ID:        1,
			Username:  "root",
			Name:      "Administrator",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			WebURL:    "http://gitlab.example.com/root",
		},
		ResourceType: "Issue",
		ResourceID:   11,
		CreatedAt:    &eventCreatedAt,
		Iteration: &ProjectIssueIteration{
			Id:          133,
			Iid:         1,
			Sequence:    1,
			GroupId:     153,
			Title:       "Iteration 1",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			State:       1,
			CreatedAt:   &createdAt,
			UpdatedAt:   &updatedAt,
			StartDate:   &startDateISOTime,
			DueDate:     &dueDateISOTime,
			WebUrl:      "",
		},
		Action: "add",
	}
	require.Equal(t, want, me)
}
