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
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListTodos(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/todos", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_todos.json")
	})

	opts := &ListTodosOptions{ListOptions: ListOptions{PerPage: 2}}
	todos, _, err := client.Todos.ListTodos(opts)

	require.NoError(t, err)

	want := []*Todo{{ID: 1, State: "pending", Target: &TodoTarget{ID: 1, ApprovalsBeforeMerge: 2}}, {ID: 2, State: "pending", Target: &TodoTarget{ID: 2}}}
	require.Equal(t, want, todos)
}

func TestMarkAllTodosAsDone(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/todos/mark_as_done", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.Todos.MarkAllTodosAsDone()
	require.NoError(t, err)
}

func TestMarkTodoAsDone(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/todos/1/mark_as_done", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	_, err := client.Todos.MarkTodoAsDone(1)
	require.NoError(t, err)
}

func TestIDStringMarshal(t *testing.T) {
	cases := []struct {
		name       string
		todoTarget *TodoTarget
		expect     interface{}
		expectErr  bool
	}{{
		name:       "int id",
		todoTarget: &TodoTarget{ID: 1},
		expect:     float64(1),
	}, {
		name:       "string id",
		todoTarget: &TodoTarget{IDString: "a"},
		expect:     "a",
	}, {
		name:       "cannot set both",
		todoTarget: &TodoTarget{ID: 1, IDString: "a"},
		expectErr:  true,
	}, {
		name:       "set neither yields float",
		todoTarget: &TodoTarget{},
		expect:     float64(0),
	}}

	extractID := func(b []byte) (interface{}, error) {
		var untyped interface{}
		err := json.Unmarshal(b, &untyped)
		if err != nil {
			return nil, err
		}
		if todoTarget, ok := untyped.(map[string]interface{}); ok {
			if id, ok := todoTarget["id"]; ok {
				return id, nil
			} else {
				return nil, fmt.Errorf("id not found")
			}
		} else {
			return nil, fmt.Errorf("unexpected shape")
		}
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := json.Marshal(tc.todoTarget)
			if err != nil && tc.expectErr {
				return // pass
			}
			if b == nil {
				t.Fatalf("wanted bytes. got nil")
			}
			id, err := extractID(b)
			if err != nil {
				t.Errorf("could not extract id: %v", err)
			}
			fail := func(want, got interface{}) {
				t.Errorf("wanted type %T (%v). got type %T (%v)", want, want, got, got)
			}
			switch got := id.(type) {
			case string:
				expect, ok := tc.expect.(string)
				if !ok {
					fail(tc.expect, id)
				}
				if expect != got {
					t.Errorf("wanted %v. got %v", expect, got)
				}
			case float64:
				expect, ok := tc.expect.(float64)
				if !ok {
					fail(tc.expect, id)
				}
				if expect != got {
					t.Errorf("wanted %v. got %v", expect, got)
				}
			default:
				t.Errorf("wanted type string or float64. got type %T (%v)", id, id)
			}
		})
	}
}

func TestIDStringUnmarshal(t *testing.T) {
	cases := []struct {
		name           string
		given          string
		expectID       int
		expectIDString string
		expectErr      bool
	}{{
		name:     "int id",
		given:    `{"id":1}`,
		expectID: 1,
	}, {
		name:           "string id",
		given:          `{"id":"a"}`,
		expectIDString: "a",
	}, {
		name:  "no id",
		given: `{}`,
	}, {
		name:      "wrong shape",
		given:     `[]`,
		expectErr: true,
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			todoTarget := &TodoTarget{}
			err := json.Unmarshal([]byte(tc.given), todoTarget)
			if err != nil && tc.expectErr {
				return // pass
			}
			if todoTarget.ID != tc.expectID {
				t.Errorf("wanted ID %v. got %v", tc.expectID, todoTarget.ID)
			}
			if todoTarget.IDString != tc.expectIDString {
				t.Errorf("wanted IDString %v. got %v", tc.expectIDString, todoTarget.IDString)
			}
		})
	}
}
