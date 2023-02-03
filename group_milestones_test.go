package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGroupMilestonesService_ListGroupMilestones(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/milestones", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 12,
				"iid": 3,
				"group_id": 5,
				"title": "10.0",
				"description": "Version",
				"state": "active",
				"expired": false,
				"web_url": "https://gitlab.com/groups/gitlab-org/-/milestones/42"
			  }
			]
		`)
	})

	want := []*GroupMilestone{{
		ID:          12,
		IID:         3,
		GroupID:     5,
		Title:       "10.0",
		Description: "Version",
		State:       "active",
		Expired:     Bool(false),
	}}

	gms, resp, err := client.GroupMilestones.ListGroupMilestones(5, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, gms)

	gms, resp, err = client.GroupMilestones.ListGroupMilestones(5.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, gms)

	gms, resp, err = client.GroupMilestones.ListGroupMilestones(5, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, gms)

	gms, resp, err = client.GroupMilestones.ListGroupMilestones(7, nil, nil)
	require.Error(t, err)
	require.Nil(t, gms)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGroupMilestonesService_GetGroupMilestone(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/milestones/12", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": 12,
			"iid": 3,
			"group_id": 5,
			"title": "10.0",
			"description": "Version",
			"state": "active",
			"expired": false,
			"web_url": "https://gitlab.com/groups/gitlab-org/-/milestones/42"
		  }
		`)
	})

	want := &GroupMilestone{
		ID:          12,
		IID:         3,
		GroupID:     5,
		Title:       "10.0",
		Description: "Version",
		State:       "active",
		Expired:     Bool(false),
	}

	gm, resp, err := client.GroupMilestones.GetGroupMilestone(5, 12, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, gm)

	gm, resp, err = client.GroupMilestones.GetGroupMilestone(5.01, 12, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, gm)

	gm, resp, err = client.GroupMilestones.GetGroupMilestone(5, 12, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, gm)

	gm, resp, err = client.GroupMilestones.GetGroupMilestone(7, 12, nil, nil)
	require.Error(t, err)
	require.Nil(t, gm)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGroupMilestonesService_CreateGroupMilestone(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/milestones", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": 12,
			"iid": 3,
			"group_id": 5,
			"title": "10.0",
			"description": "Version",
			"state": "active",
			"expired": false,
			"web_url": "https://gitlab.com/groups/gitlab-org/-/milestones/42"
		  }
		`)
	})

	want := &GroupMilestone{
		ID:          12,
		IID:         3,
		GroupID:     5,
		Title:       "10.0",
		Description: "Version",
		State:       "active",
		Expired:     Bool(false),
	}

	gm, resp, err := client.GroupMilestones.CreateGroupMilestone(5, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, gm)

	gm, resp, err = client.GroupMilestones.CreateGroupMilestone(5.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, gm)

	gm, resp, err = client.GroupMilestones.CreateGroupMilestone(5, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, gm)

	gm, resp, err = client.GroupMilestones.CreateGroupMilestone(7, nil, nil)
	require.Error(t, err)
	require.Nil(t, gm)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGroupMilestonesService_UpdateGroupMilestone(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/milestones/12", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
		  {
			"id": 12,
			"iid": 3,
			"group_id": 5,
			"title": "10.0",
			"description": "Version",
			"state": "active",
			"expired": false,
			"web_url": "https://gitlab.com/groups/gitlab-org/-/milestones/42"
		  }
		`)
	})

	want := &GroupMilestone{
		ID:          12,
		IID:         3,
		GroupID:     5,
		Title:       "10.0",
		Description: "Version",
		State:       "active",
		Expired:     Bool(false),
	}

	gm, resp, err := client.GroupMilestones.UpdateGroupMilestone(5, 12, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, gm)

	gm, resp, err = client.GroupMilestones.UpdateGroupMilestone(5.01, 12, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, gm)

	gm, resp, err = client.GroupMilestones.UpdateGroupMilestone(5, 12, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, gm)

	gm, resp, err = client.GroupMilestones.UpdateGroupMilestone(7, 12, nil, nil)
	require.Error(t, err)
	require.Nil(t, gm)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGroupMilestonesService_GetGroupMilestoneIssues(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/milestones/12/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			   {
				  "project_id" : 5,
				  "milestone" : {
					 "due_date" : null,
					 "project_id" : 5,
					 "state" : "closed",
					 "description" : "Rerum est voluptatem provident consequuntur molestias similique ipsum dolor.",
					 "iid" : 3,
					 "id" : 11,
					 "title" : "v3.0"
				  },
				  "author" : {
					 "state" : "active",
					 "web_url" : "https://gitlab.example.com/root",
					 "avatar_url" : null,
					 "username" : "root",
					 "id" : 1,
					 "name" : "Administrator"
				  },
				  "description" : "Omnis vero earum sunt corporis dolor et placeat.",
				  "state" : "closed",
				  "iid" : 1,
				  "assignees" : [{
					 "avatar_url" : null,
					 "web_url" : "https://gitlab.example.com/venky333",
					 "state" : "active",
					 "username" : "venky333",
					 "id" : 9,
					 "name" : "Venkatesh Thalluri"
				  }],
				  "assignee" : {
					 "avatar_url" : null,
					 "web_url" : "https://gitlab.example.com/venky333",
					 "state" : "active",
					 "username" : "venky333",
					 "id" : 9,
					 "name" : "Venkatesh Thalluri"
				  },
				  "id" : 41
				}
			]
		`)
	})

	want := []*Issue{{
		ID:          41,
		IID:         1,
		ExternalID:  "",
		State:       "closed",
		Description: "Omnis vero earum sunt corporis dolor et placeat.",
		Author: &IssueAuthor{
			ID:        1,
			State:     "active",
			WebURL:    "https://gitlab.example.com/root",
			Name:      "Administrator",
			AvatarURL: "",
			Username:  "root",
		},
		Milestone: &Milestone{
			ID:          11,
			IID:         3,
			ProjectID:   5,
			Title:       "v3.0",
			Description: "Rerum est voluptatem provident consequuntur molestias similique ipsum dolor.",
			StartDate:   nil,
			DueDate:     nil,
			State:       "closed",
			WebURL:      "",
			UpdatedAt:   nil,
			CreatedAt:   nil,
			Expired:     nil,
		},
		ProjectID: 5,
		Assignees: []*IssueAssignee{{
			ID:        9,
			State:     "active",
			WebURL:    "https://gitlab.example.com/venky333",
			Name:      "Venkatesh Thalluri",
			AvatarURL: "",
			Username:  "venky333",
		}},
		Assignee: &IssueAssignee{
			ID:        9,
			State:     "active",
			WebURL:    "https://gitlab.example.com/venky333",
			Name:      "Venkatesh Thalluri",
			AvatarURL: "",
			Username:  "venky333",
		},
	}}

	is, resp, err := client.GroupMilestones.GetGroupMilestoneIssues(5, 12, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, is)

	is, resp, err = client.GroupMilestones.GetGroupMilestoneIssues(5.01, 12, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, is)

	is, resp, err = client.GroupMilestones.GetGroupMilestoneIssues(5, 12, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, is)

	is, resp, err = client.GroupMilestones.GetGroupMilestoneIssues(7, 12, nil, nil)
	require.Error(t, err)
	require.Nil(t, is)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGroupMilestonesService_GetGroupMilestoneMergeRequests(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/3/milestones/12/merge_requests", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 1,
				"iid": 1,
				"project_id": 3,
				"title": "test1",
				"description": "fixed login page css paddings",
				"state": "merged",
				"merged_by": {
				  "id": 87854,
				  "name": "Douwe Maan",
				  "username": "DouweM",
				  "state": "active",
				  "avatar_url": "https://gitlab.example.com/uploads/-/system/user/avatar/87854/avatar.png",
				  "web_url": "https://gitlab.com/DouweM"
				},
				"closed_by": null,
				"closed_at": null,
				"target_branch": "master",
				"source_branch": "test1",
				"upvotes": 0,
				"downvotes": 0,
				"author": {
				  "id": 1,
				  "name": "Administrator",
				  "username": "admin",
				  "state": "active",
				  "avatar_url": null,
				  "web_url" : "https://gitlab.example.com/admin"
				},
				"assignee": {
				  "id": 1,
				  "name": "Administrator",
				  "username": "admin",
				  "state": "active",
				  "avatar_url": null,
				  "web_url" : "https://gitlab.example.com/admin"
				},
				"assignees": [{
				  "name": "Venkatesh Thalluri",
				  "username": "venkatesh.thalluri",
				  "id": 12,
				  "state": "active",
				  "avatar_url": "http://www.gravatar.com/avatar/46f6f7dc858ada7be1853f7fb96e81da?s=80&d=identicon",
				  "web_url": "https://gitlab.example.com/axel.block"
				}],
				"reviewers": [{
				  "id": 2,
				  "name": "Sam Bauch",
				  "username": "kenyatta_oconnell",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/956c92487c6f6f7616b536927e22c9a0?s=80&d=identicon",
				  "web_url": "http://gitlab.example.com//kenyatta_oconnell"
				}],
				"source_project_id": 2,
				"target_project_id": 3,
				"draft": false,
				"work_in_progress": false,
				"milestone": {
				  "id": 5,
				  "iid": 1,
				  "project_id": 3,
				  "title": "v2.0",
				  "description": "Assumenda aut placeat expedita exercitationem labore sunt enim earum.",
				  "state": "closed",
				  "web_url": "https://gitlab.example.com/my-group/my-project/milestones/1"
				},
				"merge_when_pipeline_succeeds": true,
				"detailed_merge_status": "mergeable",
				"sha": "8888888888888888888888888888888888888888",
				"merge_commit_sha": null,
				"squash_commit_sha": null,
				"user_notes_count": 1,
				"discussion_locked": null,
				"should_remove_source_branch": true,
				"force_remove_source_branch": false,
				"allow_collaboration": false,
				"allow_maintainer_to_push": false,
				"web_url": "http://gitlab.example.com/my-group/my-project/merge_requests/1",
				"references": {
				  "short": "!1",
				  "relative": "my-group/my-project!1",
				  "full": "my-group/my-project!1"
				},
				"squash": false,
				"task_completion_status":{
				  "count":0,
				  "completed_count":0
				}
			  }
			]
		`)
	})

	want := []*MergeRequest{{
		ID:           1,
		IID:          1,
		TargetBranch: "master",
		SourceBranch: "test1",
		ProjectID:    3,
		Title:        "test1",
		State:        "merged",
		Upvotes:      0,
		Downvotes:    0,
		Author: &BasicUser{
			ID:        1,
			Username:  "admin",
			Name:      "Administrator",
			State:     "active",
			CreatedAt: nil,
			AvatarURL: "",
			WebURL:    "https://gitlab.example.com/admin",
		},
		Assignee: &BasicUser{
			ID: 1, Username: "admin",
			Name:      "Administrator",
			State:     "active",
			AvatarURL: "",
			WebURL:    "https://gitlab.example.com/admin",
		},
		Assignees: []*BasicUser{{
			ID:        12,
			Username:  "venkatesh.thalluri",
			Name:      "Venkatesh Thalluri",
			State:     "active",
			AvatarURL: "http://www.gravatar.com/avatar/46f6f7dc858ada7be1853f7fb96e81da?s=80&d=identicon", WebURL: "https://gitlab.example.com/axel.block",
		}},
		Reviewers: []*BasicUser{{
			ID:        2,
			Username:  "kenyatta_oconnell",
			Name:      "Sam Bauch",
			State:     "active",
			AvatarURL: "https://www.gravatar.com/avatar/956c92487c6f6f7616b536927e22c9a0?s=80&d=identicon", WebURL: "http://gitlab.example.com//kenyatta_oconnell",
		}},
		SourceProjectID: 2,
		TargetProjectID: 3,
		Description:     "fixed login page css paddings",
		WorkInProgress:  false,
		Milestone: &Milestone{
			ID:          5,
			IID:         1,
			ProjectID:   3,
			Title:       "v2.0",
			Description: "Assumenda aut placeat expedita exercitationem labore sunt enim earum.",
			State:       "closed",
			WebURL:      "https://gitlab.example.com/my-group/my-project/milestones/1",
		},
		MergeWhenPipelineSucceeds: true,
		DetailedMergeStatus:       "mergeable",
		MergeError:                "",
		MergedBy: &BasicUser{
			ID:        87854,
			Username:  "DouweM",
			Name:      "Douwe Maan",
			State:     "active",
			AvatarURL: "https://gitlab.example.com/uploads/-/system/user/avatar/87854/avatar.png",
			WebURL:    "https://gitlab.com/DouweM",
		},
		Subscribed:               false,
		SHA:                      "8888888888888888888888888888888888888888",
		MergeCommitSHA:           "",
		SquashCommitSHA:          "",
		UserNotesCount:           1,
		ChangesCount:             "",
		ShouldRemoveSourceBranch: true,
		ForceRemoveSourceBranch:  false,
		AllowCollaboration:       false,
		WebURL:                   "http://gitlab.example.com/my-group/my-project/merge_requests/1",
		References: &IssueReferences{
			Short:    "!1",
			Relative: "my-group/my-project!1",
			Full:     "my-group/my-project!1",
		},
		DiscussionLocked:     false,
		Squash:               false,
		DivergedCommitsCount: 0,
		RebaseInProgress:     false,
		ApprovalsBeforeMerge: 0,
		Reference:            "",
		FirstContribution:    false,
		TaskCompletionStatus: &TasksCompletionStatus{
			Count:          0,
			CompletedCount: 0,
		},
		HasConflicts:                false,
		BlockingDiscussionsResolved: false,
		Overflow:                    false,
	}}

	mrs, resp, err := client.GroupMilestones.GetGroupMilestoneMergeRequests(3, 12, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, mrs)

	mrs, resp, err = client.GroupMilestones.GetGroupMilestoneMergeRequests(3.01, 12, nil, nil)
	require.EqualError(t, err, "invalid ID type 3.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, mrs)

	mrs, resp, err = client.GroupMilestones.GetGroupMilestoneMergeRequests(3, 12, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, mrs)

	mrs, resp, err = client.GroupMilestones.GetGroupMilestoneMergeRequests(7, 12, nil, nil)
	require.Error(t, err)
	require.Nil(t, mrs)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGroupMilestonesService_GetGroupMilestoneBurndownChartEvents(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/3/milestones/12/burndown_events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
				{
					"weight": 10,
					"action": "update" 
				}
			]
		`)
	})

	want := []*BurndownChartEvent{{
		Weight: Int(10),
		Action: String("update"),
	}}

	bces, resp, err := client.GroupMilestones.GetGroupMilestoneBurndownChartEvents(3, 12, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, bces)

	bces, resp, err = client.GroupMilestones.GetGroupMilestoneBurndownChartEvents(3.01, 12, nil, nil)
	require.EqualError(t, err, "invalid ID type 3.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, bces)

	bces, resp, err = client.GroupMilestones.GetGroupMilestoneBurndownChartEvents(3, 12, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, bces)

	bces, resp, err = client.GroupMilestones.GetGroupMilestoneBurndownChartEvents(7, 12, nil, nil)
	require.Error(t, err)
	require.Nil(t, bces)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
