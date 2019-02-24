package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListTodos(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/todos", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1,"state": "pending","target":{"id":1,"approvals_before_merge":2}},{"id":2,"state":"pending","target":{"id":2,"approvals_before_merge":null}}]`)
	})

	opts := &ListTodosOptions{}
	todos, _, err := client.Todos.ListTodos(opts)

	if err != nil {
		t.Errorf("Todos.ListTodos returned error: %v", err)
	}

	want := []*Todo{{ID: 1, State: "pending", Target: TodoTarget{ID: 1, ApprovalsBeforeMerge: 2}}, {ID: 2, State: "pending", Target: TodoTarget{ID: 2}}}
	if !reflect.DeepEqual(want, todos) {
		t.Errorf("Todos.ListTodos returned %+v, want %+v", todos, want)
	}

}

func TestMarkAllTodosAsDone(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/todos/mark_as_done", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.Todos.MarkAllTodosAsDone()

	if err != nil {
		t.Fatalf("Todos.MarkTodosRead returns an error: %v", err)
	}
}

func TestMarkTodoAsDone(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/todos/1/mark_as_done", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
	})

	_, err := client.Todos.MarkTodoAsDone(1)

	if err != nil {
		t.Fatalf("Todos.MarkTodoRead returns an error: %v", err)
	}
}
