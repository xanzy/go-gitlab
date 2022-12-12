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

func TestGetEpic(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/7/epics/8", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":8, "title": "Incredible idea", "description": "This is a test epic", "author" : {"id" : 26, "name": "jramsay"}}`)
	})

	epic, _, err := client.Epics.GetEpic("7", 8)
	if err != nil {
		log.Fatal(err)
	}

	want := &Epic{
		ID:          8,
		Title:       "Incredible idea",
		Description: "This is a test epic",
		Author:      &EpicAuthor{ID: 26, Name: "jramsay"},
	}

	if !reflect.DeepEqual(want, epic) {
		t.Errorf("Epics.GetEpic returned %+v, want %+v", epic, want)
	}
}

func TestDeleteEpic(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/7/epics/8", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Epics.DeleteEpic("7", 8)
	if err != nil {
		log.Fatal(err)
	}
}

func TestListGroupEpics(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/7/epics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/groups/7/epics?author_id=26&state=opened")
		fmt.Fprint(w, `[{"id":8, "title": "Incredible idea", "description": "This is a test epic", "author" : {"id" : 26, "name": "jramsay"}}]`)
	})

	listGroupEpics := &ListGroupEpicsOptions{
		AuthorID: Int(26),
		State:    String("opened"),
	}

	epics, _, err := client.Epics.ListGroupEpics("7", listGroupEpics)
	if err != nil {
		log.Fatal(err)
	}

	want := []*Epic{{
		ID:          8,
		Title:       "Incredible idea",
		Description: "This is a test epic",
		Author:      &EpicAuthor{ID: 26, Name: "jramsay"},
	}}

	if !reflect.DeepEqual(want, epics) {
		t.Errorf("Epics.ListGroupEpics returned %+v, want %+v", epics, want)
	}
}

func TestCreateEpic(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/7/epics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":8, "title": "Incredible idea", "description": "This is a test epic", "author" : {"id" : 26, "name": "jramsay"}}`)
	})

	createEpicOptions := &CreateEpicOptions{
		Title:       String("Incredible idea"),
		Description: String("This is a test epic"),
	}

	epic, _, err := client.Epics.CreateEpic("7", createEpicOptions)
	if err != nil {
		log.Fatal(err)
	}

	want := &Epic{
		ID:          8,
		Title:       "Incredible idea",
		Description: "This is a test epic",
		Author:      &EpicAuthor{ID: 26, Name: "jramsay"},
	}

	if !reflect.DeepEqual(want, epic) {
		t.Errorf("Epics.CreateEpic returned %+v, want %+v", epic, want)
	}
}

func TestUpdateEpic(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/7/epics/8", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{"id":8, "title": "Incredible idea", "description": "This is a test epic", "author" : {"id" : 26, "name": "jramsay"}}`)
	})

	updateEpicOptions := &UpdateEpicOptions{
		Title:       String("Incredible idea"),
		Description: String("This is a test epic"),
	}

	epic, _, err := client.Epics.UpdateEpic("7", 8, updateEpicOptions)
	if err != nil {
		log.Fatal(err)
	}

	want := &Epic{
		ID:          8,
		Title:       "Incredible idea",
		Description: "This is a test epic",
		Author:      &EpicAuthor{ID: 26, Name: "jramsay"},
	}

	if !reflect.DeepEqual(want, epic) {
		t.Errorf("Epics.UpdateEpic returned %+v, want %+v", epic, want)
	}
}
