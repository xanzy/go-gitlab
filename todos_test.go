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
