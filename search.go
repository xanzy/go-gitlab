//
// Copyright 2018, Sander van Harmelen
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

// SearchService handles communication with the search related methods of the
// GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/search.html
type SearchService struct {
	client *Client
}

// Projects searches the expression within projects
//
// GitLab API docs: https://docs.gitlab.com/ee/api/search.html#scope-projects
func (s *SearchService) Projects(query string, opt *ListOptions, options ...OptionFunc) ([]*Project, *Response, error) {
	var ps []*Project
	resp, err := s.search("projects", query, opt, &ps)
	return ps, resp, err
}

// ProjectsByGroup searches the expression within projects for
// the specified group
//
// GitLab API docs: https://docs.gitlab.com/ee/api/search.html#group-search-api
func (s *SearchService) ProjectsByGroup(gid interface{}, query string, opt *ListOptions, options ...OptionFunc) ([]*Project, *Response, error) {
	var ps []*Project
	resp, err := s.searchByGroup(gid, "projects", query, opt, &ps)
	return ps, resp, err
}

// Issues searches the expression within issues
//
// GitLab API docs: https://docs.gitlab.com/ee/api/search.html#scope-issues
func (s *SearchService) Issues(query string, opt *ListOptions, options ...OptionFunc) ([]*Issue, *Response, error) {
	var is []*Issue
	resp, err := s.search("issues", query, opt, &is)
	return is, resp, err
}

// IssuesByGroup searches the expression within issues for
// the specified group
//
// GitLab API docs: https://docs.gitlab.com/ee/api/search.html#scope-issues
func (s *SearchService) IssuesByGroup(gid interface{}, query string, opt *ListOptions, options ...OptionFunc) ([]*Issue, *Response, error) {
	var is []*Issue
	resp, err := s.searchByGroup(gid, "issues", query, opt, &is)
	return is, resp, err
}

// IssuesByProject searches the expression within issues for
// the specified project
//
// GitLab API docs: https://docs.gitlab.com/ee/api/search.html#scope-issues
func (s *SearchService) IssuesByProject(pid interface{}, query string, opt *ListOptions, options ...OptionFunc) ([]*Issue, *Response, error) {
	var is []*Issue
	resp, err := s.searchByProject(pid, "issues", query, opt, &is)
	return is, resp, err
}

// MergeRequests searches the expression within merge requests
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/search.html#scope-merge_requests
func (s *SearchService) MergeRequests(query string, opt *ListOptions, options ...OptionFunc) ([]*MergeRequest, *Response, error) {
	var ms []*MergeRequest
	resp, err := s.search("merge_requests", query, opt, &ms)
	return ms, resp, err
}

// MergeRequestsByGroup searches the expression within merge requests for
// the specified group
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/search.html#scope-merge_requests
func (s *SearchService) MergeRequestsByGroup(gid interface{}, query string, opt *ListOptions, options ...OptionFunc) ([]*MergeRequest, *Response, error) {
	var ms []*MergeRequest
	resp, err := s.searchByGroup(gid, "merge_requests", query, opt, &ms)
	return ms, resp, err
}

// MergeRequestsByProject searches the expression within merge requests for
// the specified project
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/search.html#scope-merge_requests
func (s *SearchService) MergeRequestsByProject(pid interface{}, query string, opt *ListOptions, options ...OptionFunc) ([]*MergeRequest, *Response, error) {
	var ms []*MergeRequest
	resp, err := s.searchByProject(pid, "merge_requests", query, opt, &ms)
	return ms, resp, err
}

// Milestones searches the expression within milestones
//
// GitLab API docs: https://docs.gitlab.com/ee/api/search.html#scope-milestones
func (s *SearchService) Milestones(query string, opt *ListOptions, options ...OptionFunc) ([]*Milestone, *Response, error) {
	var ms []*Milestone
	resp, err := s.search("milestones", query, opt, ms)
	return ms, resp, err
}

// MilestonesByGroup searches the expression within milestones for
// the specified group
//
// GitLab API docs: https://docs.gitlab.com/ee/api/search.html#scope-milestones
func (s *SearchService) MilestonesByGroup(gid interface{}, query string, opt *ListOptions, options ...OptionFunc) ([]*Milestone, *Response, error) {
	var ms []*Milestone
	resp, err := s.searchByGroup(gid, "milestones", query, opt, &ms)
	return ms, resp, err
}

// MilestonesByProject searches the expression within milestones for
// the specified project
//
// GitLab API docs: https://docs.gitlab.com/ee/api/search.html#scope-milestones
func (s *SearchService) MilestonesByProject(pid interface{}, query string, opt *ListOptions, options ...OptionFunc) ([]*Milestone, *Response, error) {
	var ms []*Milestone
	resp, err := s.searchByProject(pid, "milestones", query, opt, &ms)
	return ms, resp, err
}

// SnippetTitles searches the expression within snippet titles
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/search.html#scope-snippet_titles
func (s *SearchService) SnippetTitles(query string, opt *ListOptions, options ...OptionFunc) ([]*Snippet, *Response, error) {
	var ss []*Snippet
	resp, err := s.search("snippet_titles", query, opt, &ss)
	return ss, resp, err
}

// SnippetBlobs searches the expression within snippet blobs
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/search.html#scope-snippet_blobs
func (s *SearchService) SnippetBlobs(query string, opt *ListOptions, options ...OptionFunc) ([]*Snippet, *Response, error) {
	var ss []*Snippet
	resp, err := s.search("snippet_blobs", query, opt, &ss)
	return ss, resp, err
}

// NotesByProject searches the expression within notes for the specified
// project
//
// GitLab API docs: // https://docs.gitlab.com/ee/api/search.html#scope-notes
func (s *SearchService) NotesByProject(pid interface{}, query string, opt *ListOptions, options ...OptionFunc) ([]*Note, *Response, error) {
	var ns []*Note
	resp, err := s.searchByProject(pid, "notes", query, opt, &ns)
	return ns, resp, err
}

// WikiBlobs searches the expression within all wiki blobs
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/search.html#scope-wiki_blobs
func (s *SearchService) WikiBlobs(query string, opt *ListOptions, options ...OptionFunc) ([]*Wiki, *Response, error) {
	var ws []*Wiki
	resp, err := s.search("wiki_blobs", query, opt, &ws)
	return ws, resp, err
}

// WikiBlobsByGroup searches the expression within wiki blobs for
// specified group
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/search.html#scope-wiki_blobs
func (s *SearchService) WikiBlobsByGroup(gid interface{}, query string, opt *ListOptions, options ...OptionFunc) ([]*Wiki, *Response, error) {
	var ws []*Wiki
	resp, err := s.searchByGroup(gid, "wiki_blobs", query, opt, &ws)
	return ws, resp, err
}

// WikiBlobsByProject searches the expression within wiki blobs for
// the specified project
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/search.html#scope-wiki_blobs
func (s *SearchService) WikiBlobsByProject(pid interface{}, query string, opt *ListOptions, options ...OptionFunc) ([]*Wiki, *Response, error) {
	var ws []*Wiki
	resp, err := s.searchByProject(pid, "wiki_blobs", query, opt, &ws)
	return ws, resp, err
}

// Commits searches the expression within all commits
//
// GitLab API docs: https://docs.gitlab.com/ee/api/search.html#scope-commits
func (s *SearchService) Commits(query string, opt *ListOptions, options ...OptionFunc) ([]*Commit, *Response, error) {
	var cs []*Commit
	resp, err := s.search("commits", query, opt, &cs)
	return cs, resp, err
}

// CommitsByGroup searches the expression within commits for the specified
// group
//
// GitLab API docs: https://docs.gitlab.com/ee/api/search.html#scope-commits
func (s *SearchService) CommitsByGroup(gid interface{}, query string, opt *ListOptions, options ...OptionFunc) ([]*Commit, *Response, error) {
	var cs []*Commit
	resp, err := s.searchByGroup(gid, "commits", query, opt, &cs)
	return cs, resp, err
}

// CommitsByProject searches the expression within commits for the
// specified project
//
// GitLab API docs: https://docs.gitlab.com/ee/api/search.html#scope-commits
func (s *SearchService) CommitsByProject(pid interface{}, query string, opt *ListOptions, options ...OptionFunc) ([]*Commit, *Response, error) {
	var cs []*Commit
	resp, err := s.searchByProject(pid, "commits", query, opt, &cs)
	return cs, resp, err
}

type Blob struct {
	Basename  string `json:"basename"`
	Data      string `json:"data"`
	Filename  string `json:"filename"`
	ID        int    `json:"id"`
	Ref       string `json:"ref"`
	Startline int    `json:"startline"`
	ProjectID int    `json:"project_id"`
}

// Blobs searches the expression within all blobs
//
// GitLab API docs: https://docs.gitlab.com/ee/api/search.html#scope-blobs
func (s *SearchService) Blobs(query string, opt *ListOptions, options ...OptionFunc) ([]*Blob, *Response, error) {
	var bs []*Blob
	resp, err := s.search("blobs", query, opt, &bs)
	return bs, resp, err
}

// BlobsByGroup searches the expression within blobs for the specified
// group
//
// GitLab API docs: https://docs.gitlab.com/ee/api/search.html#scope-blobs
func (s *SearchService) BlobsByGroup(gid interface{}, query string, opt *ListOptions, options ...OptionFunc) ([]*Blob, *Response, error) {
	var bs []*Blob
	resp, err := s.searchByGroup(gid, "blobs", query, opt, &bs)
	return bs, resp, err
}

// BlobsByProject searches the expression within blobs for the specified
// project
//
// GitLab API docs: https://docs.gitlab.com/ee/api/search.html#scope-blobs
func (s *SearchService) BlobsByProject(pid interface{}, query string, opt *ListOptions, options ...OptionFunc) ([]*Blob, *Response, error) {
	var bs []*Blob
	resp, err := s.searchByProject(pid, "blobs", query, opt, &bs)
	return bs, resp, err
}

func (s *SearchService) searchByProject(pid interface{}, scope, query string, opt *ListOptions, result interface{}, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}

	params := make(url.Values)
	params.Set("scope", scope)
	params.Set("search", query)

	u := fmt.Sprintf("projects/%s/-/search?%s", url.QueryEscape(project), params.Encode())

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, result)
}

func (s *SearchService) searchByGroup(gid interface{}, scope, query string, opt *ListOptions, result interface{}, options ...OptionFunc) (*Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, err
	}

	params := make(url.Values)
	params.Set("scope", scope)
	params.Set("search", query)

	u := fmt.Sprintf("groups/%s/-/search?%s", url.QueryEscape(group), params.Encode())

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, result)
}

func (s *SearchService) search(scope, query string, opt *ListOptions, result interface{}, options ...OptionFunc) (*Response, error) {
	params := make(url.Values)
	params.Set("scope", scope)
	params.Set("search", query)

	u := fmt.Sprintf("search?%s", params.Encode())

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, result)
}
