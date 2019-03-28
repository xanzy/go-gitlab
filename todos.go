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

// TodoTarget represents a todo target of type Issue or MergeRequest
type TodoTarget struct {
	// TODO: replace both Assignee and Author structs with v4 User struct
	Assignee struct {
		Name      string `json:"name" yaml:"name"`
		Username  string `json:"username" yaml:"username"`
		ID        int    `json:"id" yaml:"id"`
		State     string `json:"state" yaml:"state"`
		AvatarURL string `json:"avatar_url" yaml:"avatar_url"`
		WebURL    string `json:"web_url" yaml:"web_url"`
	} `json:"assignee" yaml:"assignee"`
	Author struct {
		Name      string `json:"name" yaml:"name"`
		Username  string `json:"username" yaml:"username"`
		ID        int    `json:"id" yaml:"id"`
		State     string `json:"state" yaml:"state"`
		AvatarURL string `json:"avatar_url" yaml:"avatar_url"`
		WebURL    string `json:"web_url" yaml:"web_url"`
	} `json:"author" yaml:"author"`
	CreatedAt      *time.Time `json:"created_at" yaml:"created_at"`
	Description    string     `json:"description" yaml:"description"`
	Downvotes      int        `json:"downvotes" yaml:"downvotes"`
	ID             int        `json:"id" yaml:"id"`
	IID            int        `json:"iid" yaml:"iid"`
	Labels         []string   `json:"labels" yaml:"labels"`
	Milestone      Milestone  `json:"milestone" yaml:"milestone"`
	ProjectID      int        `json:"project_id" yaml:"project_id"`
	State          string     `json:"state" yaml:"state"`
	Subscribed     bool       `json:"subscribed" yaml:"subscribed"`
	Title          string     `json:"title" yaml:"title"`
	UpdatedAt      *time.Time `json:"updated_at" yaml:"updated_at"`
	Upvotes        int        `json:"upvotes" yaml:"upvotes"`
	UserNotesCount int        `json:"user_notes_count" yaml:"user_notes_count"`
	WebURL         string     `json:"web_url" yaml:"web_url"`

	// Only available for type Issue
	Confidential bool   `json:"confidential" yaml:"confidential"`
	DueDate      string `json:"due_date" yaml:"due_date"`
	Weight       int    `json:"weight" yaml:"weight"`

	// Only available for type MergeRequest
	ApprovalsBeforeMerge      int    `json:"approvals_before_merge" yaml:"approvals_before_merge"`
	ForceRemoveSourceBranch   bool   `json:"force_remove_source_branch" yaml:"force_remove_source_branch"`
	MergeCommitSHA            string `json:"merge_commit_sha" yaml:"merge_commit_sha"`
	MergeWhenPipelineSucceeds bool   `json:"merge_when_pipeline_succeeds" yaml:"merge_when_pipeline_succeeds"`
	MergeStatus               string `json:"merge_status" yaml:"merge_status"`
	SHA                       string `json:"sha" yaml:"sha"`
	ShouldRemoveSourceBranch  bool   `json:"should_remove_source_branch" yaml:"should_remove_source_branch"`
	SourceBranch              string `json:"source_branch" yaml:"source_branch"`
	SourceProjectID           int    `json:"source_project_id" yaml:"source_project_id"`
	Squash                    bool   `json:"squash" yaml:"squash"`
	TargetBranch              string `json:"target_branch" yaml:"target_branch"`
	TargetProjectID           int    `json:"target_project_id" yaml:"target_project_id"`
	WorkInProgress            bool   `json:"work_in_progress" yaml:"work_in_progress"`
}

// Todo represents a GitLab todo.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/todos.html
type Todo struct {
	ID      int `json:"id" yaml:"id"`
	Project struct {
		ID                int    `json:"id" yaml:"id"`
		HTTPURLToRepo     string `json:"http_url_to_repo" yaml:"http_url_to_repo"`
		WebURL            string `json:"web_url" yaml:"web_url"`
		Name              string `json:"name" yaml:"name"`
		NameWithNamespace string `json:"name_with_namespace" yaml:"name_with_namespace"`
		Path              string `json:"path" yaml:"path"`
		PathWithNamespace string `json:"path_with_namespace" yaml:"path_with_namespace"`
	} `json:"project" yaml:"project"`
	Author struct {
		ID        int    `json:"id" yaml:"id"`
		Name      string `json:"name" yaml:"name"`
		Username  string `json:"username" yaml:"username"`
		State     string `json:"state" yaml:"state"`
		AvatarURL string `json:"avatar_url" yaml:"avatar_url"`
		WebURL    string `json:"web_url" yaml:"web_url"`
	} `json:"author" yaml:"author"`
	ActionName TodoAction `json:"action_name" yaml:"action_name"`
	TargetType string     `json:"target_type" yaml:"target_type"`
	Target     TodoTarget `json:"target" yaml:"target"`
	TargetURL  string     `json:"target_url" yaml:"target_url"`
	Body       string     `json:"body" yaml:"body"`
	State      string     `json:"state" yaml:"state"`
	CreatedAt  *time.Time `json:"created_at" yaml:"created_at"`
}

func (t Todo) String() string {
	return Stringify(t)
}

// ListTodosOptions represents the available ListTodos() options.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/todos.html#get-a-list-of-todos
type ListTodosOptions struct {
	Action    *TodoAction `url:"action,omitempty" json:"action,omitempty" yaml:"action,omitempty"`
	AuthorID  *int        `url:"author_id,omitempty" json:"author_id,omitempty" yaml:"author_id,omitempty"`
	ProjectID *int        `url:"project_id,omitempty" json:"project_id,omitempty" yaml:"project_id,omitempty"`
	State     *string     `url:"state,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
	Type      *string     `url:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
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
