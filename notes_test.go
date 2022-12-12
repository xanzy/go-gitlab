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
	"net/http"
	"reflect"
	"testing"
)

func TestGetEpicNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/epics/4329/notes/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":3,"type":null,"body":"foo bar","attachment":null,"system":false,"noteable_id":4392,"noteable_type":"Epic","resolvable":false,"noteable_iid":null}`)
	})

	note, _, err := client.Notes.GetEpicNote("1", 4329, 3, nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &Note{
		ID:           3,
		Body:         "foo bar",
		Attachment:   "",
		Title:        "",
		FileName:     "",
		System:       false,
		NoteableID:   4392,
		NoteableType: "Epic",
	}

	if !reflect.DeepEqual(note, want) {
		t.Errorf("Notes.GetEpicNote want %#v, got %#v", note, want)
	}
}

func TestGetMergeRequestNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/4329/notes/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":3,"type":"DiffNote","body":"foo bar","attachment":null,"system":false,"noteable_id":4392,"noteable_type":"Epic","resolvable":false,"noteable_iid":null}`)
	})

	note, _, err := client.Notes.GetMergeRequestNote("1", 4329, 3, nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &Note{
		ID:           3,
		Type:         DiffNote,
		Body:         "foo bar",
		System:       false,
		NoteableID:   4392,
		NoteableType: "Epic",
	}

	if !reflect.DeepEqual(note, want) {
		t.Errorf("Notes.GetEpicNote want %#v, got %#v", note, want)
	}
}
