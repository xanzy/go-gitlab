package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEpicIssuesService_ListEpicIssues(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/epics/5/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 76,
				"iid": 6,
				"project_id": 8,
				"title" : "Consequatur vero maxime deserunt laboriosam est voluptas dolorem.",
				"description" : "Ratione dolores corrupti mollitia soluta quia.",
				"state": "opened",
				"closed_at": null,
				"labels": [],
				"milestone": {
				  "id": 38,
				  "iid": 3,
				  "project_id": 8,
				  "title": "v2.0",
				  "description": "In tempore culpa inventore quo accusantium.",
				  "state": "closed"
				},
				"assignees": [{
				  "id": 7,
				  "name": "Pamella Huel",
				  "username": "arnita",
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/a2f5c6fcef64c9c69cb8779cb292be1b?s=80&d=identicon",
				  "web_url": "http://localhost:3001/arnita"
				}],
				"assignee": {
				  "id": 7,
				  "name": "Pamella Huel",
				  "username": "arnita",
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/a2f5c6fcef64c9c69cb8779cb292be1b?s=80&d=identicon",
				  "web_url": "http://localhost:3001/arnita"
				},
				"author": {
				  "id": 13,
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/30e3b2122ccd6b8e45e8e14a3ffb58fc?s=80&d=identicon",
				  "web_url": "http://localhost:3001/venky333"
				},
				"user_notes_count": 8,
				"upvotes": 0,
				"downvotes": 0,
				"due_date": null,
				"confidential": false,
				"weight": null,
				"discussion_locked": null,
				"web_url": "http://localhost:3001/h5bp/html5-boilerplate/issues/6",
				"time_stats": {
				  "time_estimate": 0,
				  "total_time_spent": 0,
				  "human_time_estimate": null,
				  "human_total_time_spent": null
				},
				"_links":{
				  "self": "http://localhost:3001/api/v4/projects/8/issues/6",
				  "notes": "http://localhost:3001/api/v4/projects/8/issues/6/notes",
				  "award_emoji": "http://localhost:3001/api/v4/projects/8/issues/6/award_emoji",
				  "project": "http://localhost:3001/api/v4/projects/8"
				},
				"epic_issue_id": 2
			  }
			]
		`)
	})

	want := []*Issue{{
		ID:          76,
		IID:         6,
		ExternalID:  "",
		State:       "opened",
		Description: "Ratione dolores corrupti mollitia soluta quia.",
		Author: &IssueAuthor{
			ID:        13,
			State:     "active",
			WebURL:    "http://localhost:3001/venky333",
			Name:      "Venkatesh Thalluri",
			AvatarURL: "http://www.gravatar.com/avatar/30e3b2122ccd6b8e45e8e14a3ffb58fc?s=80&d=identicon",
			Username:  "venky333",
		},
		Milestone: &Milestone{
			ID:          38,
			IID:         3,
			ProjectID:   8,
			Title:       "v2.0",
			Description: "In tempore culpa inventore quo accusantium.",
			State:       "closed",
			WebURL:      "",
		},
		ProjectID: 8,
		Assignees: []*IssueAssignee{{
			ID:        7,
			State:     "active",
			WebURL:    "http://localhost:3001/arnita",
			Name:      "Pamella Huel",
			AvatarURL: "http://www.gravatar.com/avatar/a2f5c6fcef64c9c69cb8779cb292be1b?s=80&d=identicon",
			Username:  "arnita",
		}},
		Assignee: &IssueAssignee{
			ID:        7,
			State:     "active",
			WebURL:    "http://localhost:3001/arnita",
			Name:      "Pamella Huel",
			AvatarURL: "http://www.gravatar.com/avatar/a2f5c6fcef64c9c69cb8779cb292be1b?s=80&d=identicon",
			Username:  "arnita",
		},
		ClosedBy:     nil,
		Title:        "Consequatur vero maxime deserunt laboriosam est voluptas dolorem.",
		CreatedAt:    nil,
		MovedToID:    0,
		Labels:       Labels{},
		LabelDetails: nil,
		Upvotes:      0,
		Downvotes:    0,
		DueDate:      nil,
		WebURL:       "http://localhost:3001/h5bp/html5-boilerplate/issues/6",
		References:   nil,
		TimeStats: &TimeStats{
			HumanTimeEstimate:   "",
			HumanTotalTimeSpent: "",
			TimeEstimate:        0,
			TotalTimeSpent:      0,
		},
		Confidential:     false,
		Weight:           0,
		DiscussionLocked: false,
		IssueType:        nil,
		Subscribed:       false,
		UserNotesCount:   8,
		Links: &IssueLinks{
			Self:       "http://localhost:3001/api/v4/projects/8/issues/6",
			Notes:      "http://localhost:3001/api/v4/projects/8/issues/6/notes",
			AwardEmoji: "http://localhost:3001/api/v4/projects/8/issues/6/award_emoji",
			Project:    "http://localhost:3001/api/v4/projects/8",
		},
		IssueLinkID:          0,
		MergeRequestCount:    0,
		EpicIssueID:          2,
		Epic:                 nil,
		TaskCompletionStatus: nil,
	}}

	is, resp, err := client.EpicIssues.ListEpicIssues(1, 5, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, is)

	is, resp, err = client.EpicIssues.ListEpicIssues(1.01, 5, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, is)

	is, resp, err = client.EpicIssues.ListEpicIssues(1, 5, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, is)

	is, resp, err = client.EpicIssues.ListEpicIssues(3, 5, nil, nil)
	require.Error(t, err)
	require.Nil(t, is)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestEpicIssuesService_AssignEpicIssue(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/epics/5/issues/55", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
			  "id": 11,
			  "epic": {
				"id": 30,
				"iid": 5,
				"title": "Ea cupiditate dolores ut vero consequatur quasi veniam voluptatem et non.",
				"description": "Molestias dolorem eos vitae expedita impedit necessitatibus quo voluptatum.",
				"author": {
				  "id": 7,
				  "name": "Pamella Huel",
				  "username": "arnita",
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/a2f5c6fcef64c9c69cb8779cb292be1b?s=80&d=identicon",
				  "web_url": "http://localhost:3001/arnita"
				}
			  },
			  "issue": {
				"id": 55,
				"iid": 13,
				"project_id": 8,
				"title": "Beatae laborum voluptatem voluptate eligendi ex accusamus.",
				"description": "Quam veritatis debitis omnis aliquam sit.",
				"state": "opened",
				"labels": [],
				"milestone": {
				  "id": 48,
				  "iid": 6,
				  "project_id": 8,
				  "title": "Sprint - Sed sed maxime temporibus ipsa ullam qui sit.",
				  "description": "Quos veritatis qui expedita sunt deleniti accusamus.",
				  "state": "active"
				},
				"assignees": [{
				  "id": 10,
				  "name": "Lu Mayer",
				  "username": "kam",
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/018729e129a6f31c80a6327a30196823?s=80&d=identicon",
				  "web_url": "http://localhost:3001/kam"
				}],
				"assignee": {
				  "id": 10,
				  "name": "Lu Mayer",
				  "username": "kam",
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/018729e129a6f31c80a6327a30196823?s=80&d=identicon",
				  "web_url": "http://localhost:3001/kam"
				},
				"author": {
				  "id": 25,
				  "name": "User 3",
				  "username": "user3",
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/97d6d9441ff85fdc730e02a6068d267b?s=80&d=identicon",
				  "web_url": "http://localhost:3001/user3"
				},
				"user_notes_count": 0,
				"upvotes": 0,
				"downvotes": 0,
				"due_date": null,
				"confidential": false,
				"weight": null,
				"discussion_locked": null,
				"web_url": "http://localhost:3001/h5bp/html5-boilerplate/issues/13",
				"time_stats": {
				  "time_estimate": 0,
				  "total_time_spent": 0,
				  "human_time_estimate": null,
				  "human_total_time_spent": null
				}
			  }
			}
		`)
	})

	want := &EpicIssueAssignment{
		ID: 11,
		Epic: &Epic{
			ID:          30,
			IID:         5,
			GroupID:     0,
			ParentID:    0,
			Title:       "Ea cupiditate dolores ut vero consequatur quasi veniam voluptatem et non.",
			Description: "Molestias dolorem eos vitae expedita impedit necessitatibus quo voluptatum.",
			State:       "",
			WebURL:      "",
			Author: &EpicAuthor{
				ID:        7,
				State:     "active",
				WebURL:    "http://localhost:3001/arnita",
				Name:      "Pamella Huel",
				AvatarURL: "http://www.gravatar.com/avatar/a2f5c6fcef64c9c69cb8779cb292be1b?s=80&d=identicon",
				Username:  "arnita",
			},
			StartDateIsFixed: false,
			DueDateIsFixed:   false,
			Upvotes:          0, Downvotes: 0,
			UserNotesCount: 0,
			URL:            "",
		},
		Issue: &Issue{
			ID:          55,
			IID:         13,
			ExternalID:  "",
			State:       "opened",
			Description: "Quam veritatis debitis omnis aliquam sit.",
			Author: &IssueAuthor{
				ID:        25,
				State:     "active",
				WebURL:    "http://localhost:3001/user3",
				Name:      "User 3",
				AvatarURL: "http://www.gravatar.com/avatar/97d6d9441ff85fdc730e02a6068d267b?s=80&d=identicon",
				Username:  "user3",
			},
			Milestone: &Milestone{
				ID:          48,
				IID:         6,
				ProjectID:   8,
				Title:       "Sprint - Sed sed maxime temporibus ipsa ullam qui sit.",
				Description: "Quos veritatis qui expedita sunt deleniti accusamus.",
				State:       "active",
				WebURL:      "",
			},
			ProjectID: 8,
			Assignees: []*IssueAssignee{
				{
					ID:        10,
					State:     "active",
					WebURL:    "http://localhost:3001/kam",
					Name:      "Lu Mayer",
					AvatarURL: "http://www.gravatar.com/avatar/018729e129a6f31c80a6327a30196823?s=80&d=identicon",
					Username:  "kam",
				},
			},
			Assignee: &IssueAssignee{
				ID:        10,
				State:     "active",
				WebURL:    "http://localhost:3001/kam",
				Name:      "Lu Mayer",
				AvatarURL: "http://www.gravatar.com/avatar/018729e129a6f31c80a6327a30196823?s=80&d=identicon",
				Username:  "kam",
			},
			Title:     "Beatae laborum voluptatem voluptate eligendi ex accusamus.",
			MovedToID: 0,
			Labels:    Labels{},
			Upvotes:   0,
			Downvotes: 0,
			WebURL:    "http://localhost:3001/h5bp/html5-boilerplate/issues/13",
			TimeStats: &TimeStats{
				HumanTimeEstimate:   "",
				HumanTotalTimeSpent: "",
				TimeEstimate:        0,
				TotalTimeSpent:      0,
			},
			Confidential:      false,
			Weight:            0,
			DiscussionLocked:  false,
			Subscribed:        false,
			UserNotesCount:    0,
			IssueLinkID:       0,
			MergeRequestCount: 0,
			EpicIssueID:       0,
		},
	}

	eia, resp, err := client.EpicIssues.AssignEpicIssue(1, 5, 55, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, eia)

	eia, resp, err = client.EpicIssues.AssignEpicIssue(1.01, 5, 55, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, eia)

	eia, resp, err = client.EpicIssues.AssignEpicIssue(1, 5, 55, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, eia)

	eia, resp, err = client.EpicIssues.AssignEpicIssue(3, 5, 55, nil, nil)
	require.Error(t, err)
	require.Nil(t, eia)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestEpicIssuesService_RemoveEpicIssue(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/epics/5/issues/55", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprintf(w, `
			{
			  "id": 11,
			  "epic": {
				"id": 30,
				"iid": 5,
				"title": "Ea cupiditate dolores ut vero consequatur quasi veniam voluptatem et non.",
				"description": "Molestias dolorem eos vitae expedita impedit necessitatibus quo voluptatum.",
				"author": {
				  "id": 7,
				  "name": "Pamella Huel",
				  "username": "arnita",
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/a2f5c6fcef64c9c69cb8779cb292be1b?s=80&d=identicon",
				  "web_url": "http://localhost:3001/arnita"
				}
			  },
			  "issue": {
				"id": 55,
				"iid": 13,
				"project_id": 8,
				"title": "Beatae laborum voluptatem voluptate eligendi ex accusamus.",
				"description": "Quam veritatis debitis omnis aliquam sit.",
				"state": "opened",
				"labels": [],
				"milestone": {
				  "id": 48,
				  "iid": 6,
				  "project_id": 8,
				  "title": "Sprint - Sed sed maxime temporibus ipsa ullam qui sit.",
				  "description": "Quos veritatis qui expedita sunt deleniti accusamus.",
				  "state": "active"
				},
				"assignees": [{
				  "id": 10,
				  "name": "Lu Mayer",
				  "username": "kam",
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/018729e129a6f31c80a6327a30196823?s=80&d=identicon",
				  "web_url": "http://localhost:3001/kam"
				}],
				"assignee": {
				  "id": 10,
				  "name": "Lu Mayer",
				  "username": "kam",
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/018729e129a6f31c80a6327a30196823?s=80&d=identicon",
				  "web_url": "http://localhost:3001/kam"
				},
				"author": {
				  "id": 25,
				  "name": "User 3",
				  "username": "user3",
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/97d6d9441ff85fdc730e02a6068d267b?s=80&d=identicon",
				  "web_url": "http://localhost:3001/user3"
				},
				"user_notes_count": 0,
				"upvotes": 0,
				"downvotes": 0,
				"due_date": null,
				"confidential": false,
				"weight": null,
				"discussion_locked": null,
				"web_url": "http://localhost:3001/h5bp/html5-boilerplate/issues/13",
				"time_stats": {
				  "time_estimate": 0,
				  "total_time_spent": 0,
				  "human_time_estimate": null,
				  "human_total_time_spent": null
				}
			  }
			}
		`)
	})

	want := &EpicIssueAssignment{
		ID: 11,
		Epic: &Epic{
			ID:          30,
			IID:         5,
			GroupID:     0,
			ParentID:    0,
			Title:       "Ea cupiditate dolores ut vero consequatur quasi veniam voluptatem et non.",
			Description: "Molestias dolorem eos vitae expedita impedit necessitatibus quo voluptatum.",
			State:       "",
			WebURL:      "",
			Author: &EpicAuthor{
				ID:        7,
				State:     "active",
				WebURL:    "http://localhost:3001/arnita",
				Name:      "Pamella Huel",
				AvatarURL: "http://www.gravatar.com/avatar/a2f5c6fcef64c9c69cb8779cb292be1b?s=80&d=identicon",
				Username:  "arnita",
			},
			StartDateIsFixed: false,
			DueDateIsFixed:   false,
			Upvotes:          0, Downvotes: 0,
			UserNotesCount: 0,
			URL:            "",
		},
		Issue: &Issue{
			ID:          55,
			IID:         13,
			ExternalID:  "",
			State:       "opened",
			Description: "Quam veritatis debitis omnis aliquam sit.",
			Author: &IssueAuthor{
				ID:        25,
				State:     "active",
				WebURL:    "http://localhost:3001/user3",
				Name:      "User 3",
				AvatarURL: "http://www.gravatar.com/avatar/97d6d9441ff85fdc730e02a6068d267b?s=80&d=identicon",
				Username:  "user3",
			},
			Milestone: &Milestone{
				ID:          48,
				IID:         6,
				ProjectID:   8,
				Title:       "Sprint - Sed sed maxime temporibus ipsa ullam qui sit.",
				Description: "Quos veritatis qui expedita sunt deleniti accusamus.",
				State:       "active",
				WebURL:      "",
			},
			ProjectID: 8,
			Assignees: []*IssueAssignee{
				{
					ID:        10,
					State:     "active",
					WebURL:    "http://localhost:3001/kam",
					Name:      "Lu Mayer",
					AvatarURL: "http://www.gravatar.com/avatar/018729e129a6f31c80a6327a30196823?s=80&d=identicon",
					Username:  "kam",
				},
			},
			Assignee: &IssueAssignee{
				ID:        10,
				State:     "active",
				WebURL:    "http://localhost:3001/kam",
				Name:      "Lu Mayer",
				AvatarURL: "http://www.gravatar.com/avatar/018729e129a6f31c80a6327a30196823?s=80&d=identicon",
				Username:  "kam",
			},
			Title:     "Beatae laborum voluptatem voluptate eligendi ex accusamus.",
			MovedToID: 0,
			Labels:    Labels{},
			Upvotes:   0,
			Downvotes: 0,
			WebURL:    "http://localhost:3001/h5bp/html5-boilerplate/issues/13",
			TimeStats: &TimeStats{
				HumanTimeEstimate:   "",
				HumanTotalTimeSpent: "",
				TimeEstimate:        0,
				TotalTimeSpent:      0,
			},
			Confidential:      false,
			Weight:            0,
			DiscussionLocked:  false,
			Subscribed:        false,
			UserNotesCount:    0,
			IssueLinkID:       0,
			MergeRequestCount: 0,
			EpicIssueID:       0,
		},
	}

	eia, resp, err := client.EpicIssues.RemoveEpicIssue(1, 5, 55, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, eia)

	eia, resp, err = client.EpicIssues.RemoveEpicIssue(1.01, 5, 55, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, eia)

	eia, resp, err = client.EpicIssues.RemoveEpicIssue(1, 5, 55, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, eia)

	eia, resp, err = client.EpicIssues.RemoveEpicIssue(3, 5, 55, nil, nil)
	require.Error(t, err)
	require.Nil(t, eia)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestEpicIssuesService_UpdateEpicIssueAssignment(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/epics/5/issues/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
			[
			  {
				"id": 76,
				"iid": 6,
				"project_id": 8,
				"title" : "Consequatur vero maxime deserunt laboriosam est voluptas dolorem.",
				"description" : "Ratione dolores corrupti mollitia soluta quia.",
				"state": "opened",
				"closed_at": null,
				"labels": [],
				"milestone": {
				  "id": 38,
				  "iid": 3,
				  "project_id": 8,
				  "title": "v2.0",
				  "description": "In tempore culpa inventore quo accusantium.",
				  "state": "closed"
				},
				"assignees": [{
				  "id": 7,
				  "name": "Pamella Huel",
				  "username": "arnita",
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/a2f5c6fcef64c9c69cb8779cb292be1b?s=80&d=identicon",
				  "web_url": "http://localhost:3001/arnita"
				}],
				"assignee": {
				  "id": 7,
				  "name": "Pamella Huel",
				  "username": "arnita",
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/a2f5c6fcef64c9c69cb8779cb292be1b?s=80&d=identicon",
				  "web_url": "http://localhost:3001/arnita"
				},
				"author": {
				  "id": 13,
				  "name": "Venkatesh Thalluri",
				  "username": "venky333",
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/30e3b2122ccd6b8e45e8e14a3ffb58fc?s=80&d=identicon",
				  "web_url": "http://localhost:3001/venky333"
				},
				"user_notes_count": 8,
				"upvotes": 0,
				"downvotes": 0,
				"due_date": null,
				"confidential": false,
				"weight": null,
				"discussion_locked": null,
				"web_url": "http://localhost:3001/h5bp/html5-boilerplate/issues/6",
				"time_stats": {
				  "time_estimate": 0,
				  "total_time_spent": 0,
				  "human_time_estimate": null,
				  "human_total_time_spent": null
				},
				"_links":{
				  "self": "http://localhost:3001/api/v4/projects/8/issues/6",
				  "notes": "http://localhost:3001/api/v4/projects/8/issues/6/notes",
				  "award_emoji": "http://localhost:3001/api/v4/projects/8/issues/6/award_emoji",
				  "project": "http://localhost:3001/api/v4/projects/8"
				},
				"epic_issue_id": 2
			  }
			]
		`)
	})

	want := []*Issue{{
		ID:          76,
		IID:         6,
		ExternalID:  "",
		State:       "opened",
		Description: "Ratione dolores corrupti mollitia soluta quia.",
		Author: &IssueAuthor{
			ID:        13,
			State:     "active",
			WebURL:    "http://localhost:3001/venky333",
			Name:      "Venkatesh Thalluri",
			AvatarURL: "http://www.gravatar.com/avatar/30e3b2122ccd6b8e45e8e14a3ffb58fc?s=80&d=identicon",
			Username:  "venky333",
		},
		Milestone: &Milestone{
			ID:          38,
			IID:         3,
			ProjectID:   8,
			Title:       "v2.0",
			Description: "In tempore culpa inventore quo accusantium.",
			State:       "closed",
			WebURL:      "",
		},
		ProjectID: 8,
		Assignees: []*IssueAssignee{{
			ID:        7,
			State:     "active",
			WebURL:    "http://localhost:3001/arnita",
			Name:      "Pamella Huel",
			AvatarURL: "http://www.gravatar.com/avatar/a2f5c6fcef64c9c69cb8779cb292be1b?s=80&d=identicon",
			Username:  "arnita",
		}},
		Assignee: &IssueAssignee{
			ID:        7,
			State:     "active",
			WebURL:    "http://localhost:3001/arnita",
			Name:      "Pamella Huel",
			AvatarURL: "http://www.gravatar.com/avatar/a2f5c6fcef64c9c69cb8779cb292be1b?s=80&d=identicon",
			Username:  "arnita",
		},
		ClosedBy:     nil,
		Title:        "Consequatur vero maxime deserunt laboriosam est voluptas dolorem.",
		CreatedAt:    nil,
		MovedToID:    0,
		Labels:       Labels{},
		LabelDetails: nil,
		Upvotes:      0,
		Downvotes:    0,
		DueDate:      nil,
		WebURL:       "http://localhost:3001/h5bp/html5-boilerplate/issues/6",
		References:   nil,
		TimeStats: &TimeStats{
			HumanTimeEstimate:   "",
			HumanTotalTimeSpent: "",
			TimeEstimate:        0,
			TotalTimeSpent:      0,
		},
		Confidential:     false,
		Weight:           0,
		DiscussionLocked: false,
		IssueType:        nil,
		Subscribed:       false,
		UserNotesCount:   8,
		Links: &IssueLinks{
			Self:       "http://localhost:3001/api/v4/projects/8/issues/6",
			Notes:      "http://localhost:3001/api/v4/projects/8/issues/6/notes",
			AwardEmoji: "http://localhost:3001/api/v4/projects/8/issues/6/award_emoji",
			Project:    "http://localhost:3001/api/v4/projects/8",
		},
		IssueLinkID:          0,
		MergeRequestCount:    0,
		EpicIssueID:          2,
		Epic:                 nil,
		TaskCompletionStatus: nil,
	}}

	is, resp, err := client.EpicIssues.UpdateEpicIssueAssignment(1, 5, 2, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, is)

	is, resp, err = client.EpicIssues.UpdateEpicIssueAssignment(1.01, 5, 2, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, is)

	is, resp, err = client.EpicIssues.UpdateEpicIssueAssignment(1, 5, 2, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, is)

	is, resp, err = client.EpicIssues.UpdateEpicIssueAssignment(3, 5, 2, nil, nil)
	require.Error(t, err)
	require.Nil(t, is)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
