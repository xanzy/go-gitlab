package gitlab

// systemHookEvent is used to pre-process events to determine the right event type for System Hook events
type systemHookEvent struct {
	BaseSystemHookEvent
	ObjectKind string `json:"object_kind"`
}

// BaseSystemHookEvent contains System Hook's common properties
//
// GitLab API docs:
// https://docs.gitlab.com/ee/system_hooks/system_hooks.html
type BaseSystemHookEvent struct {
	EventName string `json:"event_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ProjectSystemHookEvent
//
// GitLab API docs:
// https://docs.gitlab.com/ee/system_hooks/system_hooks.html
type ProjectSystemHookEvent struct {
	BaseSystemHookEvent
	Name                 string `json:"name"`
	Path                 string `json:"path"`
	PathWithNamespace    string `json:"path_with_namespace"`
	ProjectID            int    `json:"project_id"`
	OwnerName            string `json:"owner_name"`
	OwnerEmail           string `json:"owner_email"`
	ProjectVisibility    string `json:"project_visibility"`
	OldPathWithNamespace string `json:"old_path_with_namespace,omitempty"`
}

// GroupSystemHookEvent
//
// GitLab API docs:
// https://docs.gitlab.com/ee/system_hooks/system_hooks.html
type GroupSystemHookEvent struct {
	BaseSystemHookEvent
	Name                 string `json:"name"`
	Path                 string `json:"path"`
	PathWithNamespace    string `json:"full_path"`
	GroupID              int    `json:"group_id"`
	OwnerName            string `json:"owner_name"`
	OwnerEmail           string `json:"owner_email"`
	ProjectVisibility    string `json:"project_visibility"`
	OldPath              string `json:"old_path,omitempty"`
	OldPathWithNamespace string `json:"old_full_path,omitempty"`
}

// KeySystemHookEvent
//
// GitLab API docs:
// https://docs.gitlab.com/ee/system_hooks/system_hooks.html
type KeySystemHookEvent struct {
	BaseSystemHookEvent
	ID       int    `json:"id"`
	Username string `json:"username"`
	Key      string `json:"key"`
}

// UserSystemHookEvent
//
// GitLab API docs:
// https://docs.gitlab.com/ee/system_hooks/system_hooks.html
type UserSystemHookEvent struct {
	BaseSystemHookEvent
	ID          int    `json:"user_id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	OldUsername string `json:"old_username,omitempty"`
	Email       string `json:"email"`
}

// UserGroupSystemHookEvent
//
// GitLab API docs:
// https://docs.gitlab.com/ee/system_hooks/system_hooks.html
type UserGroupSystemHookEvent struct {
	BaseSystemHookEvent
	ID          int    `json:"user_id"`
	Name        string `json:"user_name"`
	Username    string `json:"user_username"`
	Email       string `json:"user_email"`
	GroupID     int    `json:"group_id"`
	GroupName   string `json:"group_name"`
	GroupPath   string `json:"group_path"`
	GroupAccess string `json:"group_access"`
}

// UserTeamSystemHookEvent
//
// GitLab API docs:
// https://docs.gitlab.com/ee/system_hooks/system_hooks.html
type UserTeamSystemHookEvent struct {
	BaseSystemHookEvent
	ID                       int    `json:"user_id"`
	Name                     string `json:"user_name"`
	Username                 string `json:"user_username"`
	Email                    string `json:"user_email"`
	ProjectID                int    `json:"project_id"`
	ProjectName              string `json:"project_name"`
	ProjectPath              string `json:"project_path"`
	ProjectPathWithNamespace string `json:"project_path_with_namespace"`
	ProjectVisibility        string `json:"project_visibility"`
	AccessLevel              string `json:"access_level"`
}

// PushSystemHookEvent
//
// GitLab API docs:
// https://docs.gitlab.com/ee/system_hooks/system_hooks.html
type PushSystemHookEvent struct {
	BaseSystemHookEvent
}

// TagPushSystemHookEvent
//
// GitLab API docs:
// https://docs.gitlab.com/ee/system_hooks/system_hooks.html
type TagPushSystemHookEvent struct {
	BaseSystemHookEvent
}

// RepositoryUpdateSystemHookEvent
//
// GitLab API docs:
// https://docs.gitlab.com/ee/system_hooks/system_hooks.html
type RepositoryUpdateSystemHookEvent struct {
	BaseSystemHookEvent
}
