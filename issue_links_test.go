package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIssueLinksService_ListIssueRelations(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/4/issues/14/links", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id" : 84,
				"iid" : 14,
				"confidential": false,
				"issue_link_id": 1,
				"project_id" : 4,
				"title" : "Issues with auth",
				"state" : "opened",
				"assignees" : [],
				"assignee" : null,
				"labels" : [
				  "bug"
				],
				"author" : {
				  "name" : "Venkatesh Thalluri",
				  "avatar_url" : null,
				  "state" : "active",
				  "web_url" : "https://gitlab.example.com/eileen.lowe",
				  "id" : 18,
				  "username" : "venkatesh.thalluri"
				},
				"description" : null,
				"milestone" : null,
				"user_notes_count": 0,
				"due_date": null,
				"web_url": "http://example.com/example/example/issues/14",
				"confidential": false,
				"weight": null,
				"link_type": "relates_to"
			  }
			]
		`)
	})

	want := []*IssueRelation{{
		ID:           84,
		IID:          14,
		State:        "opened",
		Description:  "",
		Confidential: false,
		Author: &IssueAuthor{
			ID:        18,
			State:     "active",
			WebURL:    "https://gitlab.example.com/eileen.lowe",
			Name:      "Venkatesh Thalluri",
			AvatarURL: "",
			Username:  "venkatesh.thalluri",
		},
		Milestone:   nil,
		ProjectID:   4,
		Assignees:   []*IssueAssignee{},
		Assignee:    nil,
		Title:       "Issues with auth",
		Labels:      []string{"bug"},
		WebURL:      "http://example.com/example/example/issues/14",
		Weight:      0,
		IssueLinkID: 1,
		LinkType:    "relates_to",
	}}

	is, resp, err := client.IssueLinks.ListIssueRelations(4, 14, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, is)

	is, resp, err = client.IssueLinks.ListIssueRelations(4.01, 14, nil)
	require.EqualError(t, err, "invalid ID type 4.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, is)

	is, resp, err = client.IssueLinks.ListIssueRelations(4, 14, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, is)

	is, resp, err = client.IssueLinks.ListIssueRelations(8, 14, nil)
	require.Error(t, err)
	require.Nil(t, is)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestIssueLinksService_CreateIssueLink(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/4/issues/1/links", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
			  "source_issue" : {
				"id" : 83,
				"iid" : 11,
				"project_id" : 4,
				"title" : "Issues with auth",
				"state" : "opened",
				"assignees" : [],
				"assignee" : null,
				"labels" : [
				  "bug"
				],
				"author" : {
				  "name" : "Venkatesh Thalluri",
				  "avatar_url" : null,
				  "state" : "active",
				  "web_url" : "https://gitlab.example.com/eileen.lowe",
				  "id" : 18,
				  "username" : "venkatesh.thalluri"
				},
				"description" : null,
				"milestone" : null,
				"subscribed" : true,
				"user_notes_count": 0,
				"due_date": null,
				"web_url": "http://example.com/example/example/issues/11",
				"confidential": false,
				"weight": null
			  },
			  "target_issue" : {
				"id" : 84,
				"iid" : 14,
				"project_id" : 4,
				"title" : "Issues with auth",
				"state" : "opened",
				"assignees" : [],
				"assignee" : null,
				"labels" : [
				  "bug"
				],
				"author" : {
				  "name" : "Alexandra Bashirian",
				  "avatar_url" : null,
				  "state" : "active",
				  "web_url" : "https://gitlab.example.com/eileen.lowe",
				  "id" : 18,
				  "username" : "eileen.lowe"
				},
				"description" : null,
				"milestone" : null,
				"subscribed" : true,
				"user_notes_count": 0,
				"due_date": null,
				"web_url": "http://example.com/example/example/issues/14",
				"confidential": false,
				"weight": null
			  },
			  "link_type": "relates_to"
			}
		`)
	})

	want := &IssueLink{
		SourceIssue: &Issue{
			ID:          83,
			IID:         11,
			ExternalID:  "",
			State:       "opened",
			Description: "",
			Author: &IssueAuthor{
				ID:        18,
				State:     "active",
				WebURL:    "https://gitlab.example.com/eileen.lowe",
				Name:      "Venkatesh Thalluri",
				AvatarURL: "",
				Username:  "venkatesh.thalluri",
			},
			ProjectID:         4,
			Assignees:         []*IssueAssignee{},
			Title:             "Issues with auth",
			MovedToID:         0,
			Labels:            []string{"bug"},
			Upvotes:           0,
			Downvotes:         0,
			WebURL:            "http://example.com/example/example/issues/11",
			Confidential:      false,
			Weight:            0,
			DiscussionLocked:  false,
			Subscribed:        true,
			UserNotesCount:    0,
			IssueLinkID:       0,
			MergeRequestCount: 0,
			EpicIssueID:       0,
		},
		TargetIssue: &Issue{
			ID:          84,
			IID:         14,
			ExternalID:  "",
			State:       "opened",
			Description: "",
			Author: &IssueAuthor{
				ID:        18,
				State:     "active",
				WebURL:    "https://gitlab.example.com/eileen.lowe",
				Name:      "Alexandra Bashirian",
				AvatarURL: "",
				Username:  "eileen.lowe",
			},
			ProjectID:         4,
			Assignees:         []*IssueAssignee{},
			Title:             "Issues with auth",
			MovedToID:         0,
			Labels:            []string{"bug"},
			Upvotes:           0,
			Downvotes:         0,
			WebURL:            "http://example.com/example/example/issues/14",
			Confidential:      false,
			Weight:            0,
			DiscussionLocked:  false,
			Subscribed:        true,
			UserNotesCount:    0,
			IssueLinkID:       0,
			MergeRequestCount: 0,
			EpicIssueID:       0,
		},
		LinkType: "relates_to",
	}

	i, resp, err := client.IssueLinks.CreateIssueLink(4, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, i)

	i, resp, err = client.IssueLinks.CreateIssueLink(4.01, 1, nil)
	require.EqualError(t, err, "invalid ID type 4.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, i)

	i, resp, err = client.IssueLinks.CreateIssueLink(4, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, i)

	i, resp, err = client.IssueLinks.CreateIssueLink(8, 1, nil)
	require.Error(t, err)
	require.Nil(t, i)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestIssueLinksService_DeleteIssueLink(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/4/issues/1/links/83", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprintf(w, `
			{
			  "source_issue" : {
				"id" : 83,
				"iid" : 11,
				"project_id" : 4,
				"title" : "Issues with auth",
				"state" : "opened",
				"assignees" : [],
				"assignee" : null,
				"labels" : [
				  "bug"
				],
				"author" : {
				  "name" : "Venkatesh Thalluri",
				  "avatar_url" : null,
				  "state" : "active",
				  "web_url" : "https://gitlab.example.com/eileen.lowe",
				  "id" : 18,
				  "username" : "venkatesh.thalluri"
				},
				"description" : null,
				"milestone" : null,
				"subscribed" : true,
				"user_notes_count": 0,
				"due_date": null,
				"web_url": "http://example.com/example/example/issues/11",
				"confidential": false,
				"weight": null
			  },
			  "target_issue" : {
				"id" : 84,
				"iid" : 14,
				"project_id" : 4,
				"title" : "Issues with auth",
				"state" : "opened",
				"assignees" : [],
				"assignee" : null,
				"labels" : [
				  "bug"
				],
				"author" : {
				  "name" : "Alexandra Bashirian",
				  "avatar_url" : null,
				  "state" : "active",
				  "web_url" : "https://gitlab.example.com/eileen.lowe",
				  "id" : 18,
				  "username" : "eileen.lowe"
				},
				"description" : null,
				"milestone" : null,
				"subscribed" : true,
				"user_notes_count": 0,
				"due_date": null,
				"web_url": "http://example.com/example/example/issues/14",
				"confidential": false,
				"weight": null
			  },
			  "link_type": "relates_to"
			}
		`)
	})

	want := &IssueLink{
		SourceIssue: &Issue{
			ID:          83,
			IID:         11,
			ExternalID:  "",
			State:       "opened",
			Description: "",
			Author: &IssueAuthor{
				ID:        18,
				State:     "active",
				WebURL:    "https://gitlab.example.com/eileen.lowe",
				Name:      "Venkatesh Thalluri",
				AvatarURL: "",
				Username:  "venkatesh.thalluri",
			},
			ProjectID:         4,
			Assignees:         []*IssueAssignee{},
			Title:             "Issues with auth",
			MovedToID:         0,
			Labels:            []string{"bug"},
			Upvotes:           0,
			Downvotes:         0,
			WebURL:            "http://example.com/example/example/issues/11",
			Confidential:      false,
			Weight:            0,
			DiscussionLocked:  false,
			Subscribed:        true,
			UserNotesCount:    0,
			IssueLinkID:       0,
			MergeRequestCount: 0,
			EpicIssueID:       0,
		},
		TargetIssue: &Issue{
			ID:          84,
			IID:         14,
			ExternalID:  "",
			State:       "opened",
			Description: "",
			Author: &IssueAuthor{
				ID:        18,
				State:     "active",
				WebURL:    "https://gitlab.example.com/eileen.lowe",
				Name:      "Alexandra Bashirian",
				AvatarURL: "",
				Username:  "eileen.lowe",
			},
			ProjectID:         4,
			Assignees:         []*IssueAssignee{},
			Title:             "Issues with auth",
			MovedToID:         0,
			Labels:            []string{"bug"},
			Upvotes:           0,
			Downvotes:         0,
			WebURL:            "http://example.com/example/example/issues/14",
			Confidential:      false,
			Weight:            0,
			DiscussionLocked:  false,
			Subscribed:        true,
			UserNotesCount:    0,
			IssueLinkID:       0,
			MergeRequestCount: 0,
			EpicIssueID:       0,
		},
		LinkType: "relates_to",
	}

	i, resp, err := client.IssueLinks.DeleteIssueLink(4, 1, 83, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, i)

	i, resp, err = client.IssueLinks.DeleteIssueLink(4.01, 1, 83, nil)
	require.EqualError(t, err, "invalid ID type 4.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, i)

	i, resp, err = client.IssueLinks.DeleteIssueLink(4, 1, 83, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, i)

	i, resp, err = client.IssueLinks.DeleteIssueLink(8, 1, 83, nil)
	require.Error(t, err)
	require.Nil(t, i)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
