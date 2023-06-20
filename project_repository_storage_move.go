//
// Copyright 2023, Nick Westbury
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
	"time"
)

// ProjectRepositoryStorageMoveService handles communication with the repositories
// related methods of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_repository_storage_moves.html
type ProjectRepositoryStorageMoveService struct {
	client *Client
}

// ProjectRepositoryStorageMove represents the status of a Repository move.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_repository_storage_moves.html
type ProjectRepositoryStorageMove struct {
	ID                     int       `json:"id"`
	CreatedAt              time.Time `json:"created_at"`
	State                  string    `json:"state"`
	SourceStorageName      string    `json:"source_storage_name"`
	DestinationStorageName string    `json:"destination_storage_name"`
	Project                struct {
		ID                int       `json:"id"`
		Description       string    `json:"description"`
		Name              string    `json:"name"`
		NameWithNamespace string    `json:"name_with_namespace"`
		Path              string    `json:"path"`
		PathWithNamespace string    `json:"path_with_namespace"`
		CreatedAt         time.Time `json:"created_at"`
	} `json:"project"`
}

// ProjectRepositoryStorageMoveOptions represents the available parameters when interacting
// with the repository storage move API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_repository_storage_moves.html
type ProjectRepositoryStorageMoveOptions struct {
	ListOptions
	ProjectID              int `json:"project_id,omitempty"`
	DestinationStorageName int `json:"destination_storage_name,omitempty"`
}

// RetrieveAllProjectRepositoryStorageMoves Retrieves all repository storage moves accessible by the authenticated user.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_repository_storage_moves.html#retrieve-all-project-repository-storage-moves
func (s ProjectRepositoryStorageMoveService) RetrieveAllProjectRepositoryStorageMoves(opts ProjectRepositoryStorageMoveOptions, options ...RequestOptionFunc) ([]*ProjectRepositoryStorageMove, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "project_repository_storage_moves", opts, options)
	if err != nil {
		return nil, nil, err
	}

	var p []*ProjectRepositoryStorageMove
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

// RetrieveAllProjectRepositoryStorageMovesForProject Retrieves all repository storage moves for a single project
// accessible by the authenticated user.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_repository_storage_moves.html#retrieve-all-repository-storage-moves-for-a-project
func (s ProjectRepositoryStorageMoveService) RetrieveAllProjectRepositoryStorageMovesForProject(id interface{}, opts ProjectRepositoryStorageMoveOptions, options ...RequestOptionFunc) ([]*ProjectRepositoryStorageMove, *Response, error) {
	endpoint := fmt.Sprintf("/projects/%v/repository_storage_moves", id)
	req, err := s.client.NewRequest(http.MethodGet, endpoint, opts, options)
	if err != nil {
		return nil, nil, err
	}

	var p []*ProjectRepositoryStorageMove
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

// GetProjectRepositoryStorageMove Get a single repository storage move
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_repository_storage_moves.html#get-a-single-project-repository-storage-move
func (s ProjectRepositoryStorageMoveService) GetProjectRepositoryStorageMove(id interface{}, opts ProjectRepositoryStorageMoveOptions, options ...RequestOptionFunc) (*ProjectRepositoryStorageMove, *Response, error) {
	endpoint := fmt.Sprint("project_repository_storage_moves/", id)
	req, err := s.client.NewRequest(http.MethodGet, endpoint, opts, options)
	if err != nil {
		return nil, nil, err
	}

	var p *ProjectRepositoryStorageMove
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

// GetProjectRepositoryStorageMoveForProject Get a single repository storage move for a specified project
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_repository_storage_moves.html#get-a-single-repository-storage-move-for-a-project
func (s ProjectRepositoryStorageMoveService) GetProjectRepositoryStorageMoveForProject(projectID interface{}, repositoryStorageMoveID interface{}, opts ProjectRepositoryStorageMoveOptions, options ...RequestOptionFunc) (*ProjectRepositoryStorageMove, *Response, error) {
	endpoint := fmt.Sprintf("/projects/%v/repository_storage_moves/%v", projectID, repositoryStorageMoveID)
	req, err := s.client.NewRequest(http.MethodGet, endpoint, opts, options)
	if err != nil {
		return nil, nil, err
	}

	var p *ProjectRepositoryStorageMove
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

// ScheduleRepositoryStorageMove Schedule a repository to be moved
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_repository_storage_moves.html#schedule-a-repository-storage-move-for-a-project
func (s ProjectRepositoryStorageMoveService) ScheduleRepositoryStorageMove(opts ProjectRepositoryStorageMoveOptions, options ...RequestOptionFunc) (*ProjectRepositoryStorageMove, *Response, error) {
	endpoint := fmt.Sprintf("/projects/%v/repository_storage_moves", opts.ProjectID)
	req, err := s.client.NewRequest(http.MethodPost, endpoint, opts, options)
	if err != nil {
		return nil, nil, err
	}

	var p *ProjectRepositoryStorageMove
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}
