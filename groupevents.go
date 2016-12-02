package gitlab

import "time"

//GroupMergeEvent represents a MergeEvent when added to group merge_request webhooks
type GroupMergeEvent struct {
	ObjectKind string `json:"object_kind"`
	User       struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	} `json:"user"`
	Project struct {
		Name              string               `json:"name"`
		Description       string               `json:"description"`
		WebURL            string               `json:"web_url"`
		AvatarURL         string               `json:"avatar_url"`
		GitSSHURL         string               `json:"git_ssh_url"`
		GitHTTPURL        string               `json:"git_http_url"`
		Namespace         string               `json:"namespace"`
		VisibilityLevel   VisibilityLevelValue `json:"visibility_level"`
		PathWithNamespace string               `json:"path_with_namespace"`
		DefaultBranch     string               `json:"default_branch"`
		Homepage          string               `json:"homepage"`
		URL               string               `json:"url"`
		SSHURL            string               `json:"ssh_url"`
		HTTPURL           string               `json:"http_url"`
	}
	ObjectAttributes struct {
		ID              int    `json:"id"`
		TargetBranch    string `json:"target_branch"`
		SourceBranch    string `json:"source_branch"`
		SourceProjectID int    `json:"source_project_id"`
		AuthorID        int    `json:"author_id"`
		AssigneeID      int    `json:"assignee_id"`
		Title           string `json:"title"`
		CreatedAt       string `json:"created_at"` // Should be *time.Time (see Gitlab issue #21468)
		UpdatedAt       string `json:"updated_at"` // Should be *time.Time (see Gitlab issue #21468)
		MilestoneID     int    `json:"milestone_id"`
		State           string `json:"state"`
		MergeStatus     string `json:"merge_status"`
		TargetProjectID int    `json:"target_project_id"`
		Iid             int    `json:"iid"`
		Description     string `json:"description"`
		Position        int    `json:"position"`
		LockedAt        string `json:"locked_at"`
		UpdatedByID     int    `json:"updated_by_id"`
		MergeError      string `json:"merge_error"`
		MergeParams     struct {
			ForceRemoveSourceBranch string `json:"force_remove_source_branch"`
		} `json:"merge_params"`
		MergeWhenBuildSucceeds   bool        `json:"merge_when_build_succeeds"`
		MergeUserID              int         `json:"merge_user_id"`
		MergeCommitSha           string      `json:"merge_commit_sha"`
		DeletedAt                string      `json:"deleted_at"`
		ApprovalsBeforeMerge     string      `json:"approvals_before_merge"`
		RebaseCommitSha          string      `json:"rebase_commit_sha"`
		InProgressMergeCommitSha string      `json:"in_progress_merge_commit_sha"`
		LockVersion              int         `json:"lock_version"`
		TimeEstimate             int         `json:"time_estimate"`
		Source                   *Repository `json:"source"`
		Target                   *Repository `json:"target"`
		LastCommit               struct {
			ID        string     `json:"id"`
			Message   string     `json:"message"`
			Timestamp *time.Time `json:"timestamp"`
			URL       string     `json:"url"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			}
		} `json:"last_commit"`
		WorkInProgress bool   `json:"work_in_progress"`
		URL            string `json:"url"`
		Action         string `json:"action"`
	} `json:"object_attributes"`
	Repository *Repository `json:"repository"`
	Assignee   struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	} `json:"assignee"`
}
