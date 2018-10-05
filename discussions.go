//
// Copyright 2018, steperdin
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
	"time"
)

// DiscussionService handles communication with disscussion on related methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/discussions.html
type DiscussionService struct {
	client *Client
}

// Discussion represents a GitLab Discussion.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/discussions.html
type Discussion struct {
	ID             string `json:"id"`
	IndividualNote bool   `json:"individual_note"`
	Notes          []struct {
		ID         int         `json:"id"`
		Type       string      `json:"type"`
		Body       string      `json:"body"`
		Attachment interface{} `json:"attachment"`
		Author     struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Username  string `json:"username"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		} `json:"author"`
		CreatedAt    time.Time   `json:"created_at"`
		UpdatedAt    time.Time   `json:"updated_at"`
		System       bool        `json:"system"`
		NoteableID   int         `json:"noteable_id"`
		NoteableType string      `json:"noteable_type"`
		NoteableIid  interface{} `json:"noteable_iid"`
		Resolved     bool        `json:"resolved"`
		Resolvable   bool        `json:"resolvable"`
		ResolvedBy   interface{} `json:"resolved_by"`
	} `json:"notes"`
}

const (
	discussionMergeRequest = "merge_requests"
	discussionIssue        = "issues"
	discussionSnippets     = "snippets"
	discussionCommits      = "commits"
)

// ListDiscussionsOptions represents the available options for listing discussions
// for each resources
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html
type ListDiscussionsOptions ListOptions

// ListMergeRequestDiscussions gets a list of all discussions on the merge request.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#list-project-merge-request-discussions
func (s *DiscussionService) ListMergeRequestDiscussions(pid interface{}, mergeRequestIID int, opt *ListDiscussionsOptions, options ...OptionFunc) ([]*Discussion, *Response, error) {
	return s.listDiscussions(pid, discussionMergeRequest, mergeRequestIID, opt, options...)
}

// ListIssueDiscussions gets a list of all discussions on the issue.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#list-project-issue-discussions
func (s *DiscussionService) ListIssueDiscussions(pid interface{}, issueIID int, opt *ListDiscussionsOptions, options ...OptionFunc) ([]*Discussion, *Response, error) {
	return s.listDiscussions(pid, discussionIssue, issueIID, opt, options...)
}

// ListSnippetDiscussions gets a list of all discussions on the snippet.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#list-project-snippet-discussions
func (s *DiscussionService) ListSnippetDiscussions(pid interface{}, snippetID int, opt *ListDiscussionsOptions, options ...OptionFunc) ([]*Discussion, *Response, error) {
	return s.listDiscussions(pid, discussionSnippets, snippetID, opt, options...)
}

// ListCommitDiscussions gets a list of all discussions on the commit.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#list-project-commit-discussions
func (s *DiscussionService) ListCommitDiscussions(pid interface{}, commitID int, opt *ListDiscussionsOptions, options ...OptionFunc) ([]*Discussion, *Response, error) {
	return s.listDiscussions(pid, discussionSnippets, commitID, opt, options...)
}

func (s *DiscussionService) listDiscussions(pid interface{}, resource string, resourceID int, opt *ListDiscussionsOptions, options ...OptionFunc) ([]*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/%s/%d/discussions",
		url.QueryEscape(project),
		resource,
		resourceID,
	)

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var d []*Discussion
	resp, err := s.client.Do(req, &d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, err
}

// GetMergeRequestDiscussion get an discussion from merge request.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#get-single-merge-request-discussion
func (s *DiscussionService) GetMergeRequestDiscussion(pid interface{}, mergeRequestIID, discussionID int, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.getDiscussion(pid, discussionMergeRequest, mergeRequestIID, discussionID, options...)
}

// GetIssueDiscussion get an discussion from issue.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#get-single-issue-discussion
func (s *DiscussionService) GetIssueDiscussion(pid interface{}, issueIID, discussionID int, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.getDiscussion(pid, discussionIssue, issueIID, discussionID, options...)
}

// GetSnippetDiscussion get an discussion from snippet.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#get-single-snippet-discussion
func (s *DiscussionService) GetSnippetDiscussion(pid interface{}, snippetID, discussionID int, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.getDiscussion(pid, discussionSnippets, snippetID, discussionID, options...)
}

// GetCommitDiscussion get an discussion from commit.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#get-single-commit-discussion
func (s *DiscussionService) GetCommitDiscussion(pid interface{}, commitID, discussionID int, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.getDiscussion(pid, discussionCommits, commitID, discussionID, options...)
}

func (s *DiscussionService) getDiscussion(pid interface{}, resource string, resourceID, discussionID int, options ...OptionFunc) (*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/%s/%d/discussions/%d",
		url.QueryEscape(project),
		resource,
		resourceID,
		discussionID,
	)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	d := new(Discussion)
	resp, err := s.client.Do(req, &d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, err
}

// CreateDiscussionOptions represents the available options for discussion
// for a resource
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#create-new-merge-request-discussion
type CreateDiscussionOptions struct {
	Body string `url:"body,omitempty" json:"body"`
}

// CreateMergeRequestDiscussion create a discussion from merge request.
//
// GitLab API docs:
// https://docs.gitlab.com/ce/api/discussions.html#create-new-merge-request-discussion
func (s *DiscussionService) CreateMergeRequestDiscussion(pid interface{}, mergeRequestIID int, opt *CreateDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	return s.createDiscussion(pid, discussionMergeRequest, mergeRequestIID, opt, options...)
}

// // CreateIssueAwardEmoji get an award emoji from issue.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-a-new-emoji
// func (s *AwardEmojiService) CreateIssueAwardEmoji(pid interface{}, issueIID int, opt *CreateAwardEmojiOptions, options ...OptionFunc) (*AwardEmoji, *Response, error) {
// 	return s.createAwardEmoji(pid, awardIssue, issueIID, opt, options...)
// }

// // CreateSnippetAwardEmoji get an award emoji from snippet.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-a-new-emoji
// func (s *AwardEmojiService) CreateSnippetAwardEmoji(pid interface{}, snippetID int, opt *CreateAwardEmojiOptions, options ...OptionFunc) (*AwardEmoji, *Response, error) {
// 	return s.createAwardEmoji(pid, awardSnippets, snippetID, opt, options...)
// }

func (s *DiscussionService) createDiscussion(pid interface{}, resource string, resourceID int, opt *CreateDiscussionOptions, options ...OptionFunc) (*Discussion, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/%s/%d/discussions",
		url.QueryEscape(project),
		resource,
		resourceID,
	)

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	d := new(Discussion)
	resp, err := s.client.Do(req, &d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, err
}

// // DeleteIssueAwardEmoji delete award emoji on an issue.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-a-new-emoji-on-a-note
// func (s *AwardEmojiService) DeleteIssueAwardEmoji(pid interface{}, issueIID, awardID int, options ...OptionFunc) (*Response, error) {
// 	return s.deleteAwardEmoji(pid, awardMergeRequest, issueIID, awardID, options...)
// }

// // DeleteMergeRequestAwardEmoji delete award emoji on a merge request.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-a-new-emoji-on-a-note
// func (s *AwardEmojiService) DeleteMergeRequestAwardEmoji(pid interface{}, mergeRequestIID, awardID int, options ...OptionFunc) (*Response, error) {
// 	return s.deleteAwardEmoji(pid, awardMergeRequest, mergeRequestIID, awardID, options...)
// }

// // DeleteSnippetAwardEmoji delete award emoji on a snippet.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-a-new-emoji-on-a-note
// func (s *AwardEmojiService) DeleteSnippetAwardEmoji(pid interface{}, snippetID, awardID int, options ...OptionFunc) (*Response, error) {
// 	return s.deleteAwardEmoji(pid, awardMergeRequest, snippetID, awardID, options...)
// }

// // DeleteAwardEmoji Delete an award emoji on the specified resource.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#delete-an-award-emoji
// func (s *AwardEmojiService) deleteAwardEmoji(pid interface{}, resource string, resourceID, awardID int, options ...OptionFunc) (*Response, error) {
// 	project, err := parseID(pid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	u := fmt.Sprintf("projects/%s/%s/%d/award_emoji/%d", url.QueryEscape(project), resource,
// 		resourceID, awardID)

// 	req, err := s.client.NewRequest("DELETE", u, nil, options)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return s.client.Do(req, nil)
// }

// // ListIssuesAwardEmojiOnNote gets a list of all award emoji on a note from the
// // issue.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-emoji-on-notes
// func (s *AwardEmojiService) ListIssuesAwardEmojiOnNote(pid interface{}, issueID, noteID int, opt *ListAwardEmojiOptions, options ...OptionFunc) ([]*AwardEmoji, *Response, error) {
// 	return s.listAwardEmojiOnNote(pid, awardIssue, issueID, noteID, opt, options...)
// }

// // ListMergeRequestAwardEmojiOnNote gets a list of all award emoji on a note
// // from the merge request.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-emoji-on-notes
// func (s *AwardEmojiService) ListMergeRequestAwardEmojiOnNote(pid interface{}, mergeRequestIID, noteID int, opt *ListAwardEmojiOptions, options ...OptionFunc) ([]*AwardEmoji, *Response, error) {
// 	return s.listAwardEmojiOnNote(pid, awardMergeRequest, mergeRequestIID, noteID, opt, options...)
// }

// // ListSnippetAwardEmojiOnNote gets a list of all award emoji on a note from the
// // snippet.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-emoji-on-notes
// func (s *AwardEmojiService) ListSnippetAwardEmojiOnNote(pid interface{}, snippetIID, noteID int, opt *ListAwardEmojiOptions, options ...OptionFunc) ([]*AwardEmoji, *Response, error) {
// 	return s.listAwardEmojiOnNote(pid, awardSnippets, snippetIID, noteID, opt, options...)
// }

// func (s *AwardEmojiService) listAwardEmojiOnNote(pid interface{}, resources string, ressourceID, noteID int, opt *ListAwardEmojiOptions, options ...OptionFunc) ([]*AwardEmoji, *Response, error) {
// 	project, err := parseID(pid)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	u := fmt.Sprintf("projects/%s/%s/%d/notes/%d/award_emoji", url.QueryEscape(project), resources,
// 		ressourceID, noteID)

// 	req, err := s.client.NewRequest("GET", u, opt, options)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	var as []*AwardEmoji
// 	resp, err := s.client.Do(req, &as)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return as, resp, err
// }

// // GetIssuesAwardEmojiOnNote gets an award emoji on a note from an issue.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-emoji-on-notes
// func (s *AwardEmojiService) GetIssuesAwardEmojiOnNote(pid interface{}, issueID, noteID, awardID int, options ...OptionFunc) (*AwardEmoji, *Response, error) {
// 	return s.getSingleNoteAwardEmoji(pid, awardIssue, issueID, noteID, awardID, options...)
// }

// // GetMergeRequestAwardEmojiOnNote gets an award emoji on a note from a
// // merge request.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-emoji-on-notes
// func (s *AwardEmojiService) GetMergeRequestAwardEmojiOnNote(pid interface{}, mergeRequestIID, noteID, awardID int, options ...OptionFunc) (*AwardEmoji, *Response, error) {
// 	return s.getSingleNoteAwardEmoji(pid, awardMergeRequest, mergeRequestIID, noteID, awardID,
// 		options...)
// }

// // GetSnippetAwardEmojiOnNote gets an award emoji on a note from a snippet.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-emoji-on-notes
// func (s *AwardEmojiService) GetSnippetAwardEmojiOnNote(pid interface{}, snippetIID, noteID, awardID int, options ...OptionFunc) (*AwardEmoji, *Response, error) {
// 	return s.getSingleNoteAwardEmoji(pid, awardSnippets, snippetIID, noteID, awardID, options...)
// }

// func (s *AwardEmojiService) getSingleNoteAwardEmoji(pid interface{}, ressource string, resourceID, noteID, awardID int, options ...OptionFunc) (*AwardEmoji, *Response, error) {
// 	project, err := parseID(pid)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	u := fmt.Sprintf("projects/%s/%s/%d/notes/%d/award_emoji/%d",
// 		url.QueryEscape(project),
// 		ressource,
// 		resourceID,
// 		noteID,
// 		awardID,
// 	)

// 	req, err := s.client.NewRequest("GET", u, nil, options)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	a := new(AwardEmoji)
// 	resp, err := s.client.Do(req, &a)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return a, resp, err
// }

// // CreateIssuesAwardEmojiOnNote gets an award emoji on a note from an issue.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-emoji-on-notes
// func (s *AwardEmojiService) CreateIssuesAwardEmojiOnNote(pid interface{}, issueID, noteID int, opt *CreateAwardEmojiOptions, options ...OptionFunc) (*AwardEmoji, *Response, error) {
// 	return s.createAwardEmojiOnNote(pid, awardIssue, issueID, noteID, opt, options...)
// }

// // CreateMergeRequestAwardEmojiOnNote gets an award emoji on a note from a
// // merge request.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-emoji-on-notes
// func (s *AwardEmojiService) CreateMergeRequestAwardEmojiOnNote(pid interface{}, mergeRequestIID, noteID int, opt *CreateAwardEmojiOptions, options ...OptionFunc) (*AwardEmoji, *Response, error) {
// 	return s.createAwardEmojiOnNote(pid, awardMergeRequest, mergeRequestIID, noteID, opt, options...)
// }

// // CreateSnippetAwardEmojiOnNote gets an award emoji on a note from a snippet.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-emoji-on-notes
// func (s *AwardEmojiService) CreateSnippetAwardEmojiOnNote(pid interface{}, snippetIID, noteID int, opt *CreateAwardEmojiOptions, options ...OptionFunc) (*AwardEmoji, *Response, error) {
// 	return s.createAwardEmojiOnNote(pid, awardSnippets, snippetIID, noteID, opt, options...)
// }

// // CreateAwardEmojiOnNote award emoji on a note.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-a-new-emoji-on-a-note
// func (s *AwardEmojiService) createAwardEmojiOnNote(pid interface{}, resource string, resourceID, noteID int, opt *CreateAwardEmojiOptions, options ...OptionFunc) (*AwardEmoji, *Response, error) {
// 	project, err := parseID(pid)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	u := fmt.Sprintf("projects/%s/%s/%d/notes/%d/award_emoji",
// 		url.QueryEscape(project),
// 		resource,
// 		resourceID,
// 		noteID,
// 	)

// 	req, err := s.client.NewRequest("POST", u, nil, options)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	a := new(AwardEmoji)
// 	resp, err := s.client.Do(req, &a)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return a, resp, err
// }

// // DeleteIssuesAwardEmojiOnNote deletes an award emoji on a note from an issue.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-emoji-on-notes
// func (s *AwardEmojiService) DeleteIssuesAwardEmojiOnNote(pid interface{}, issueID, noteID, awardID int, options ...OptionFunc) (*Response, error) {
// 	return s.deleteAwardEmojiOnNote(pid, awardIssue, issueID, noteID, awardID, options...)
// }

// // DeleteMergeRequestAwardEmojiOnNote deletes an award emoji on a note from a
// // merge request.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-emoji-on-notes
// func (s *AwardEmojiService) DeleteMergeRequestAwardEmojiOnNote(pid interface{}, mergeRequestIID, noteID, awardID int, options ...OptionFunc) (*Response, error) {
// 	return s.deleteAwardEmojiOnNote(pid, awardMergeRequest, mergeRequestIID, noteID, awardID,
// 		options...)
// }

// // DeleteSnippetAwardEmojiOnNote deletes an award emoji on a note from a snippet.
// //
// // GitLab API docs:
// // https://docs.gitlab.com/ce/api/award_emoji.html#award-emoji-on-notes
// func (s *AwardEmojiService) DeleteSnippetAwardEmojiOnNote(pid interface{}, snippetIID, noteID, awardID int, options ...OptionFunc) (*Response, error) {
// 	return s.deleteAwardEmojiOnNote(pid, awardSnippets, snippetIID, noteID, awardID, options...)
// }

// func (s *AwardEmojiService) deleteAwardEmojiOnNote(pid interface{}, resource string, resourceID, noteID, awardID int, options ...OptionFunc) (*Response, error) {
// 	project, err := parseID(pid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	u := fmt.Sprintf("projects/%s/%s/%d/notes/%d/award_emoji/%d",
// 		url.QueryEscape(project),
// 		resource,
// 		resourceID,
// 		noteID,
// 		awardID,
// 	)

// 	req, err := s.client.NewRequest("DELETE", u, nil, options)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return s.client.Do(req, nil)
// }
