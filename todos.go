package gitlab

import "time"
import "fmt"

// TodosService handles communication with the todos related methods of
// the Gitlab API.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/todos.html
type TodosService struct {
	client *Client
}

// TodoAction represents the available actions that can be performed on a todo.
// Todo represents a GitLab todo.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/todos.html
type TodoAction string

// The available todo actions.
const (
	TodoAssigned          TodoAction = "assigned"
	TodoMentioned         TodoAction = "mentioned"
	TodoBuildFailed       TodoAction = "build_failed"
	TodoMarked            TodoAction = "marked"
	TodoApprovalRequired  TodoAction = "approval_required"
	TodoDirectlyAddressed TodoAction = "directly_addressed"
)

// Todo represents a GitLab todo.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/todos.html
type Todo struct {
	ID      int `json:"id"`
	Project struct {
		ID                int    `json:"id"`
		HTTPURLToRepo     string `json:"http_url_to_repo"`
		WebURL            string `json:"web_url"`
		Name              string `json:"name"`
		NameWithNamespace string `json:"name_with_namespace"`
		Path              string `json:"path"`
		PathWithNamespace string `json:"path_with_namespace"`
	} `json:"project"`
	Author struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"author"`
	ActionName TodoAction  `json:"action_name"`
	TargetType string      `json:"target_type"`
	Target     interface{} `json:"target"`
	TargetURL  string      `json:"target_url"`
	Body       string      `json:"body"`
	State      string      `json:"state"`
	CreatedAt  *time.Time  `json:"created_at"`
}

func (t Todo) String() string {
	return Stringify(t)
}

// ListTodosOptions represents the available ListTodos() options.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/todos.html#get-a-list-of-todos
type ListTodosOptions struct {
	Action    *TodoAction `url:"action,omitempty" json:"action,omitempty"`
	AuthorID  *int        `url:"author_id,omitempty" json:"author_id,omitempty"`
	ProjectID *int        `url:"project_id,omitempty" json:"project_id,omitempty"`
	State     *string     `url:"state,omitempty" json:"state,omitempty"`
	Type      *string     `url:"type,omitempty" json:"type,omitempty"`
}

// ListTodos lists all todos created by authenticated user.
// When no filter is applied, it returns all pending todos for the current user.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/todos.html#get-a-list-of-todos
func (s *TodosService) ListTodos(opt *ListTodosOptions, options ...OptionFunc) ([]*Todo, *Response, error) {
	req, err := s.client.NewRequest("GET", "todos", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var t []*Todo
	resp, err := s.client.Do(req, &t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, err
}

// MarkTodoAsDone marks a single pending todo given by its ID for the current user as done.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/todos.html#mark-a-todo-as-done
func (s *TodosService) MarkTodoAsDone(id int, options ...OptionFunc) (*Response, error) {
	u := fmt.Sprintf("todos/%d/mark_as_done", id)

	req, err := s.client.NewRequest("POST", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// MarkAllTodosAsDone marks all pending todos for the current user as done.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/todos.html#mark-all-todos-as-done
func (s *TodosService) MarkAllTodosAsDone(options ...OptionFunc) (*Response, error) {
	req, err := s.client.NewRequest("POST", "todos/mark_as_done", nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
