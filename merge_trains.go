package gitlab

import (
	"fmt"
	"net/http"
	"time"
)

// MergeTrainsService handles communication with the merge trains related
// methods of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/merge_trains.html
type MergeTrainsService struct {
	client *Client
}

// MergeTrain represents a Gitlab merge train.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/merge_trains.html
type MergeTrain struct {
	ID           int                     `json:"id"`
	MergeRequest *MergeTrainMergeRequest `json:"merge_request"`
	User         *MergeTrainUser         `json:"user"`
	Pipeline     *MergeTrainPipeline     `json:"pipeline"`
	CreatedAt    *time.Time              `json:"created_at"`
	UpdatedAt    *time.Time              `json:"updated_at"`
	TargetBranch string                  `json:"target_branch"`
	Status       string                  `json:"status"`
	MergedAt     *time.Time              `json:"merged_at"`
	Duration     int                     `json:"duration"`
}

// MergeTrainMergeRequest represents a Gitlab merge request inside merge train
type MergeTrainMergeRequest struct {
	ID          int        `json:"id"`
	IID         int        `json:"iid"`
	ProjectID   int        `json:"project_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	State       string     `json:"state"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	WebURL      string     `json:"web_url"`
}

// MergeTrainUser represents a Gitlab user inside merge train
type MergeTrainUser struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

// MergeTrainPipeline represents a Gitlab pipeline inside merge train
type MergeTrainPipeline struct {
	ID        int        `json:"id"`
	IID       int        `json:"iid"`
	ProjectID int        `json:"project_id"`
	SHA       string     `json:"sha"`
	Ref       string     `json:"ref"`
	Status    string     `json:"status"`
	Source    string     `json:"source"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	WebURL    string     `json:"web_url"`
}

// ListMergeTrainsOptions represents the available ListMergeTrain() options.
//
// Gitab API docs:
// https://docs.gitlab.com/ee/api/merge_trains.html#list-merge-trains-for-a-project
type ListMergeTrainsOptions struct {
	ListOptions
	Scope *string `url:"scope,omitempty" json:"scope,omitempty"`
	Sort  *string `url:"sort,omitempty" json:"sort,omitempty"`
}

// ListProjectMergeTrains get a list of merge trains in a project
//
//	The scope of trains to show: active (to be merged) and complete (have been merged)
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/merge_trains.html#list-merge-trains-for-a-project
func (s *MergeTrainsService) ListProjectMergeTrains(pid interface{}, opt *ListMergeTrainsOptions, options ...RequestOptionFunc) ([]*MergeTrain, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/merge_trains", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var mergeTrains []*MergeTrain
	resp, err := s.client.Do(req, &mergeTrains)
	if err != nil {
		return nil, resp, err
	}

	return mergeTrains, resp, nil
}

// ListMergeRequestInMergeTrain gets a list of merge requests added to a merge train
// for the requested target branch
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/merge_trains.html#list-merge-requests-in-a-merge-train
func (s *MergeTrainsService) ListMergeRequestInMergeTrain(pid interface{}, targetBranch string, opts *ListMergeTrainsOptions, options ...RequestOptionFunc) ([]*MergeTrain, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/merge_trains/%s", PathEscape(project), targetBranch)

	req, err := s.client.NewRequest(http.MethodGet, u, opts, options)
	if err != nil {
		return nil, nil, err
	}

	var mergeTrains []*MergeTrain
	resp, err := s.client.Do(req, &mergeTrains)
	if err != nil {
		return nil, resp, err
	}

	return mergeTrains, resp, nil
}

// GetMergeRequestOnAMergeTrain Get merge train information for the requested merge request.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/merge_trains.html#get-the-status-of-a-merge-request-on-a-merge-train
func (s *MergeTrainsService) GetMergeRequestOnAMergeTrain(pid interface{}, mergeRequest int, options ...RequestOptionFunc) (*MergeTrain, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/merge_trains/merge_requests/%d", PathEscape(project), mergeRequest)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var mergeTrain *MergeTrain
	resp, err := s.client.Do(req, &mergeTrain)
	if err != nil {
		return nil, resp, err
	}

	return mergeTrain, resp, nil
}

// AddMergeRequestToMergeTrainOptions represents the available AddMergeRequestToMergeTrain() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/merge_trains.html#add-a-merge-request-to-a-merge-train
type AddMergeRequestToMergeTrainOptions struct {
	WhenPipelineSucceeds *bool   `json:"when_pipeline_succeeds,omitempty"`
	SHA                  *string `json:"sha,omitempty"`
	Squash               *string `json:"squash,omitempty"`
}

// AddMergeRequestToMergeTrain Add a merge request to the merge train targeting the merge request’s target branch
//
// GitLab API docs: https://docs.gitlab.com/ee/api/merge_trains.html#add-a-merge-request-to-a-merge-train
func (s *MergeTrainsService) AddMergeRequestToMergeTrain(pid interface{}, mergeRequest int, opts *AddMergeRequestToMergeTrainOptions, options ...RequestOptionFunc) ([]*MergeTrain, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/merge_trains/merge_requests/%d", PathEscape(project), mergeRequest)

	req, err := s.client.NewRequest(http.MethodPost, u, opts, options)
	if err != nil {
		return nil, nil, err
	}

	var mergeTrains []*MergeTrain
	resp, err := s.client.Do(req, &mergeTrains)
	if err != nil {
		return nil, resp, err
	}

	return mergeTrains, resp, nil
}
