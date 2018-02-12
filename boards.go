//
// Copyright 2015, Sander van Harmelen
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
	"net/url"
)

// BoardsService handles communication with the issue related methods
// of the GitLab API.
//
// GitLab API docs:
// https://gitlab.com/gitlab-org/gitlab-ce/blob/8-16-stable/doc/api/boards.md
type BoardsService struct {
	client *Client
}

// Board represents a GitLab board.
//
// GitLab API docs:
// https://gitlab.com/gitlab-org/gitlab-ce/blob/8-16-stable/doc/api/boards.md
type Board struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Project   Project     `json:"project"`
	Milestone Milestone   `json:"milestone"`
	Lists     []BoardList `json:"lists"`
}

// BoardList represents a GitLab board list item.
//
// GitLab API docs:
// https://gitlab.com/gitlab-org/gitlab-ce/blob/8-16-stable/doc/api/boards.md
type BoardList struct {
	ID       int   `json:"id"`
	Label    Label `json:"label"`
	Position int   `json:"position"`
}

func (b Board) String() string {
	return Stringify(b)
}

func (b BoardList) String() string {
	return Stringify(b)
}

// ListProjectBoards gets a list of all boards of project.
//
// GitLab API docs:
// https://gitlab.com/gitlab-org/gitlab-ce/blob/8-16-stable/doc/api/boards.md#project-board
func (b *BoardsService) ListProjectBoards(pid interface{}) ([]*Board, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/boards", url.QueryEscape(project))

	req, err := b.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var boards []*Board
	resp, err := b.client.Do(req, &boards)
	if err != nil {
		return nil, resp, err
	}

	return boards, resp, err
}

// ListProjectBoard gets a single board list of project.
//
// GitLab API docs:
// https://gitlab.com/gitlab-org/gitlab-ce/blob/8-16-stable/doc/api/boards.md#project-board
func (b *BoardsService) ListProjectBoard(pid interface{}, bid interface{}) (*Board, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	boardID, err := parseID(bid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/boards/%s", url.QueryEscape(project), url.QueryEscape(boardID))

	req, err := b.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var board *Board
	resp, err := b.client.Do(req, &board)
	if err != nil {
		return nil, resp, err
	}

	return board, resp, err
}

// ListProjectBoardList gets a single board list of project.
//
// GitLab API docs:
// https://gitlab.com/gitlab-org/gitlab-ce/blob/8-16-stable/doc/api/boards.md#project-board
func (b *BoardsService) ListProjectBoardList(pid interface{}, bid interface{}, lid interface{}) (*BoardList, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	boardID, err := parseID(bid)
	if err != nil {
		return nil, nil, err
	}
	listID, err := parseID(lid)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("projects/%s/boards/%s/lists/%s", url.QueryEscape(project), url.QueryEscape(boardID), url.QueryEscape(listID))

	req, err := b.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var boardList *BoardList
	resp, err := b.client.Do(req, &boardList)
	if err != nil {
		return nil, resp, err
	}

	return boardList, resp, err
}
