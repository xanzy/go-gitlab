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

func TestGetDraftNote(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/projects/1/merge_requests/4329/draft_notes/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id": 37349978, "author_id": 10271899, "merge_request_id": 291473309, "resolve_discussion": false, "discussion_id": null, "note": "Some draft note", "commit_id": null, "position": null, "line_code": null}`)
	})

	note, _, err := client.DraftNotes.GetDraftNote("1", 4329, 3)
	if err != nil {
		t.Fatal(err)
	}

	want := &DraftNote{
		ID:                37349978,
		AuthorID:          10271899,
		MergeRequestID:    291473309,
		ResolveDiscussion: false,
		DiscussionID:      "",
		Note:              "Some draft note",
		CommitID:          "",
		LineCode:          "",
		Position:          nil,
	}

	if !reflect.DeepEqual(note, want) {
		t.Errorf("DraftNotes.GetDraftNote want %#v, got %#v", note, want)
	}
}
