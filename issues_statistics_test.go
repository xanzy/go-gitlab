//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestGetIssuesStatistics(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/issues_statistics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
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
		Statistics: struct {
			Counts struct {
				All    int `json:"all"`
				Closed int `json:"closed"`
				Opened int `json:"opened"`
			} `json:"counts"`
		}{
			Counts: struct {
				All    int `json:"all"`
				Closed int `json:"closed"`
				Opened int `json:"opened"`
			}{
				20, 5, 15,
			},
		},
	}

	if !reflect.DeepEqual(want, issue) {
		t.Errorf("IssuesStatistics.GetIssuesStatistics returned %+v, want %+v", issue, want)
	}
}

func TestGetGroupIssuesStatistics(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/issues_statistics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/groups/1/issues_statistics?assignee_id=1&author_id=1")
		fmt.Fprint(w, `{"statistics": {"counts": {"all": 20,"closed": 5,"opened": 15}}}`)
	})

	opt := &GetGroupIssuesStatisticsOptions{
		AssigneeID: Int(1),
		AuthorID:   Int(1),
	}

	issue, _, err := client.IssuesStatistics.GetGroupIssuesStatistics(1, opt)
	if err != nil {
		log.Fatal(err)
	}

	want := &IssuesStatistics{
		Statistics: struct {
			Counts struct {
				All    int `json:"all"`
				Closed int `json:"closed"`
				Opened int `json:"opened"`
			} `json:"counts"`
		}{
			Counts: struct {
				All    int `json:"all"`
				Closed int `json:"closed"`
				Opened int `json:"opened"`
			}{
				20, 5, 15,
			},
		},
	}

	if !reflect.DeepEqual(want, issue) {
		t.Errorf("IssuesStatistics.GetGroupIssuesStatistics returned %+v, want %+v", issue, want)
	}
}

func TestGetProjectIssuesStatistics(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues_statistics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/projects/1/issues_statistics?assignee_id=1&author_id=1")
		fmt.Fprint(w, `{"statistics": {"counts": {"all": 20,"closed": 5,"opened": 15}}}`)
	})

	opt := &GetProjectIssuesStatisticsOptions{
		AssigneeID: Int(1),
		AuthorID:   Int(1),
	}

	issue, _, err := client.IssuesStatistics.GetProjectIssuesStatistics(1, opt)
	if err != nil {
		log.Fatal(err)
	}

	want := &IssuesStatistics{
		Statistics: struct {
			Counts struct {
				All    int `json:"all"`
				Closed int `json:"closed"`
				Opened int `json:"opened"`
			} `json:"counts"`
		}{
			Counts: struct {
				All    int `json:"all"`
				Closed int `json:"closed"`
				Opened int `json:"opened"`
			}{
				20, 5, 15,
			},
		},
	}

	if !reflect.DeepEqual(want, issue) {
		t.Errorf("IssuesStatistics.GetProjectIssuesStatistics returned %+v, want %+v", issue, want)
	}
}
