//
// Copyright 2021, Sander van Harmelen
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
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	ExpectedGroup     = "webhook-test"
	excpectedAvatar   = "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon"
	expectedEmail     = "user1@example.com"
	expectedEventName = "user_add_to_group"
	expectedID        = 1
	expectedName      = "User1"
	expectedUsername  = "user1"
)

func TestBuildEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/build.json")

	var event *BuildEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("Build Event can not unmarshaled: %v\n ", err.Error())
	}

	if event == nil {
		t.Errorf("Build Event is null")
	}

	if event.BuildID != 1977 {
		t.Errorf("BuildID is %v, want %v", event.BuildID, 1977)
	}

	if event.User.ID != 42 {
		t.Errorf("User ID is %d, want %d", event.User.ID, 42)
	}

	if event.User.Name != expectedName {
		t.Errorf("Username is %s, want %s", event.User.Name, expectedName)
	}

	if event.BuildCreatedAt != "2021-02-23T02:41:37.886Z" {
		t.Errorf("BuildCreatedAt is %s, want 2021-02-23T02:41:37.886Z", event.BuildCreatedAt)
	}
}

func TestCommitCommentEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/note_commit.json")

	var event *CommitCommentEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("Commit Comment Event can not unmarshaled: %v\n ", err.Error())
	}

	if event == nil {
		t.Errorf("Commit Comment Event is null")
	}

	if event.ObjectKind != string(NoteEventTargetType) {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, NoteEventTargetType)
	}

	if event.EventType != "note" {
		t.Errorf("EventType is %v, want %v", event.EventType, "note")
	}

	if event.ProjectID != 5 {
		t.Errorf("ProjectID is %v, want %v", event.ProjectID, 5)
	}

	if event.User.ID != 42 {
		t.Errorf("User ID is %d, want %d", event.User.ID, 42)
	}

	if event.Repository.Name != "Gitlab Test" {
		t.Errorf("Repository name is %v, want %v", event.Repository.Name, "Gitlab Test")
	}

	if event.ObjectAttributes.NoteableType != "Commit" {
		t.Errorf("NoteableType is %v, want %v", event.ObjectAttributes.NoteableType, "Commit")
	}

	if event.Commit.Title != "Add submodule" {
		t.Errorf("Issue title is %v, want %v", event.Commit.Title, "Add submodule")
	}
}

func TestJobEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/job.json")

	var event *JobEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("Job Event can not unmarshaled: %v\n ", err.Error())
	}

	if event == nil {
		t.Errorf("Job Event is null")
	}

	expectedEvent := JobEvent{
		ObjectKind:          "build",
		Ref:                 "main",
		Tag:                 false,
		BeforeSHA:           "0000000000000000000000000000000000000000",
		SHA:                 "95d49d1efbd941908580e79d65e4b5ecaf4a8305",
		BuildID:             3580121225,
		BuildName:           "auto_deploy:start",
		BuildStage:          "coordinated:tag",
		BuildStatus:         "success",
		BuildCreatedAt:      "2023-01-10 13:50:02 UTC",
		BuildStartedAt:      "2023-01-10 13:50:05 UTC",
		BuildFinishedAt:     "2023-01-10 13:50:54 UTC",
		BuildDuration:       49.503592,
		BuildQueuedDuration: 0.193009,
		BuildAllowFailure:   false,
		BuildFailureReason:  "unknown_failure",
		RetriesCount:        1,
		PipelineID:          743121198,
		ProjectID:           31537070,
		ProjectName:         "John Smith / release-tools-fake",
		User: &EventUser{
			ID:        2967854,
			Name:      "John Smith",
			Username:  "jsmithy2",
			AvatarURL: "https://gitlab.com/uploads/-/system/user/avatar/2967852/avatar.png",
			Email:     "john@smith.com",
		},
		Repository: &Repository{
			Name:              "release-tools-fake",
			Description:       "",
			WebURL:            "",
			AvatarURL:         "",
			GitSSHURL:         "git@gitlab.com:jsmithy2/release-tools-fake.git",
			GitHTTPURL:        "https://gitlab.com/jsmithy2/release-tools-fake.git",
			Namespace:         "",
			Visibility:        "",
			PathWithNamespace: "",
			DefaultBranch:     "",
			Homepage:          "https://gitlab.com/jsmithy2/release-tools-fake",
			URL:               "git@gitlab.com:jsmithy2/release-tools-fake.git",
			SSHURL:            "",
			HTTPURL:           "",
		},
	}
	expectedEvent.Commit.ID = 743121198
	expectedEvent.Commit.Name = "Build pipeline"
	expectedEvent.Commit.SHA = "95d49d1efbd941908580e79d65e4b5ecaf4a8305"
	expectedEvent.Commit.Message = "Remove test jobs and add back other jobs"
	expectedEvent.Commit.AuthorName = "John Smith"
	expectedEvent.Commit.AuthorEmail = "john@smith.com"
	expectedEvent.Commit.AuthorURL = "https://gitlab.com/jsmithy2"
	expectedEvent.Commit.Status = "running"
	expectedEvent.Commit.Duration = 128
	expectedEvent.Commit.StartedAt = "2023-01-10 13:50:05 UTC"
	expectedEvent.Commit.FinishedAt = "2022-10-12 08:09:29 UTC"

	expectedEvent.Runner.ID = 12270837
	expectedEvent.Runner.Description = "4-blue.shared.runners-manager.gitlab.com/default"
	expectedEvent.Runner.RunnerType = "instance_type"
	expectedEvent.Runner.Active = true
	expectedEvent.Runner.IsShared = true
	expectedEvent.Runner.Tags = []string{"linux", "docker"}

	expectedEvent.Environment.Name = "production"
	expectedEvent.Environment.Action = "start"
	expectedEvent.Environment.DeploymentTier = "production"

	assert.Equal(t, expectedEvent, *event, "event should be equal to the expected one")
}

func TestDeploymentEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/deployment.json")

	var event *DeploymentEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("Deployment Event can not unmarshaled: %v\n ", err.Error())
	}

	if event == nil {
		t.Errorf("Deployment Event is null")
	}

	if event.Project.ID != 30 {
		t.Errorf("Project.ID is %v, want %v", event.Project.ID, 30)
	}

	if event.User.ID != 42 {
		t.Errorf("User ID is %d, want %d", event.User.ID, 42)
	}

	if event.User.Name != expectedName {
		t.Errorf("Username is %s, want %s", event.User.Name, expectedName)
	}

	if event.CommitTitle != "Add new file" {
		t.Errorf("CommitTitle is %s, want %s", event.CommitTitle, "Add new file")
	}

	if event.Ref != "1.0.0" {
		t.Errorf("Ref is %s, want %s", event.Ref, "1.0.0")
	}

	if event.StatusChangedAt != "2021-04-28 21:50:00 +0200" {
		t.Errorf("StatusChangedAt is %s, want %s", event.StatusChangedAt, "2021-04-28 21:50:00 +0200")
	}

	if event.DeploymentID != 15 {
		t.Errorf("DeploymentID is %d, want %d", event.DeploymentID, 15)
	}

	if event.EnvironmentSlug != "staging" {
		t.Errorf("EnvironmentSlug is %s, want %s", event.EnvironmentSlug, "staging")
	}

	if event.EnvironmentExternalURL != "https://staging.example.com" {
		t.Errorf("EnvironmentExternalURL is %s, want %s", event.EnvironmentExternalURL, "https://staging.example.com")
	}
}

func TestFeatureFlagEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/feature_flag.json")

	var event *FeatureFlagEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("FeatureFlag Event can not unmarshaled: %v\n ", err.Error())
	}

	if event == nil {
		t.Errorf("FeatureFlag Event is null")
	}

	if event.ObjectKind != "feature_flag" {
		t.Errorf("ObjectKind is %s, want %s", event.ObjectKind, "feature_flag")
	}

	if event.Project.ID != 1 {
		t.Errorf("Project.ID is %v, want %v", event.Project.ID, 1)
	}

	if event.User.ID != 1 {
		t.Errorf("User ID is %d, want %d", event.User.ID, 1)
	}

	if event.User.Name != "Administrator" {
		t.Errorf("Username is %s, want %s", event.User.Name, "Administrator")
	}

	if event.ObjectAttributes.ID != 6 {
		t.Errorf("ObjectAttributes.ID is %d, want %d", event.ObjectAttributes.ID, 6)
	}

	if event.ObjectAttributes.Name != "test-feature-flag" {
		t.Errorf("ObjectAttributes.Name is %s, want %s", event.ObjectAttributes.Name, "test-feature-flag")
	}

	if event.ObjectAttributes.Description != "test-feature-flag-description" {
		t.Errorf("ObjectAttributes.Description is %s, want %s", event.ObjectAttributes.Description, "test-feature-flag-description")
	}

	if event.ObjectAttributes.Active != true {
		t.Errorf("ObjectAttributes.Active is %t, want %t", event.ObjectAttributes.Active, true)
	}
}

func TestIssueCommentEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/note_issue.json")

	var event *IssueCommentEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("Issue Comment Event can not unmarshaled: %v\n ", err.Error())
	}

	if event.ObjectKind != string(NoteEventTargetType) {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, NoteEventTargetType)
	}

	if event.EventType != "note" {
		t.Errorf("EventType is %v, want %v", event.EventType, "note")
	}

	if event.ProjectID != 5 {
		t.Errorf("ProjectID is %v, want %v", event.ProjectID, 5)
	}

	if event.User.ID != 42 {
		t.Errorf("User ID is %d, want %d", event.User.ID, 42)
	}

	if event.ObjectAttributes.NoteableType != "Issue" {
		t.Errorf("NoteableType is %v, want %v", event.ObjectAttributes.NoteableType, "Issue")
	}

	if event.Issue.Title != "test_issue" {
		t.Errorf("Issue title is %v, want %v", event.Issue.Title, "test_issue")
	}

	if event.Issue.Position != 0 {
		t.Errorf("Issue position is %v, want %v", event.Issue.Position, 0)
	}

	if event.Issue.BranchName != "" {
		t.Errorf("Issue branch name is %v, want %v", event.Issue.BranchName, "")
	}

	if len(event.Issue.Labels) == 0 || event.Issue.Labels[0].ID != 25 {
		t.Errorf("Label id is null")
	}

	assert.Equal(t, []*EventLabel{
		{
			ID:          25,
			Title:       "Afterpod",
			Color:       "#3e8068",
			ProjectID:   0,
			CreatedAt:   "2019-06-05T14:32:20.211Z",
			UpdatedAt:   "2019-06-05T14:32:20.211Z",
			Template:    false,
			Description: "",
			Type:        "GroupLabel",
			GroupID:     4,
		},
		{
			ID:          86,
			Title:       "Element",
			Color:       "#231afe",
			ProjectID:   4,
			CreatedAt:   "2019-06-05T14:32:20.637Z",
			UpdatedAt:   "2019-06-05T14:32:20.637Z",
			Template:    false,
			Description: "",
			Type:        "ProjectLabel",
			GroupID:     0,
		},
	}, event.Issue.Labels)
}

func TestIssueEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/issue.json")

	var event *IssueEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("Issue Event can not unmarshaled: %v\n ", err.Error())
	}

	if event.ObjectKind != string(IssueEventTargetType) {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, IssueEventTargetType)
	}

	if event.EventType != "issue" {
		t.Errorf("EventType is %v, want %v", event.EventType, "issue")
	}

	if event.Project.ID != 1 {
		t.Errorf("Project.ID is %v, want %v", event.Project.ID, 1)
	}

	if event.User.ID != 1 {
		t.Errorf("User ID is %d, want %d", event.User.ID, 1)
	}

	if event.Assignee.Username != "user1" {
		t.Errorf("Assignee username is %s, want %s", event.Assignee.Username, "user1")
	}

	if event.ObjectAttributes.ID != 301 {
		t.Errorf("ObjectAttributes.ID is %v, want %v", event.ObjectAttributes.ID, 301)
	}

	if event.ObjectAttributes.Title != "New API: create/update/delete file" {
		t.Errorf("ObjectAttributes.Title is %v, want %v", event.ObjectAttributes.Title, "New API: create/update/delete file")
	}

	if event.ObjectAttributes.StateID != StateIDOpen {
		t.Errorf("ObjectAttributes.StateID is %v, want %v", event.ObjectAttributes.StateID, StateIDOpen)
	}

	if event.ObjectAttributes.State != "opened" {
		t.Errorf("ObjectAttributes.State is %v, want %v", event.ObjectAttributes.State, "opened")
	}
	if event.ObjectAttributes.Confidential != false {
		t.Errorf("ObjectAttributes.Confidential is %v, want %v", event.ObjectAttributes.Confidential, false)
	}

	if event.ObjectAttributes.TotalTimeSpent != 0 {
		t.Errorf("ObjectAttributes.TotalTimeSpent is %v, want %v", event.ObjectAttributes.TotalTimeSpent, 0)
	}

	if event.ObjectAttributes.Action != "open" {
		t.Errorf("ObjectAttributes.Action is %v, want %v", event.ObjectAttributes.Action, "open")
	}

	if event.ObjectAttributes.EscalationStatus != "triggered" {
		t.Errorf("ObjectAttributes.EscalationStatus is %v, want %v", event.ObjectAttributes.EscalationStatus, "triggered")
	}

	if event.ObjectAttributes.EscalationPolicy.ID != 18 {
		t.Errorf("ObjectAttributes.EscalationPolicy.ID is %v, want %v", event.ObjectAttributes.EscalationPolicy.ID, 18)
	}

	if event.Changes.TotalTimeSpent.Previous != 8100 {
		t.Errorf("Changes.TotalTimeSpent.Previous is %v , want %v", event.Changes.TotalTimeSpent.Previous, 8100)
	}

	if event.Changes.TotalTimeSpent.Current != 9900 {
		t.Errorf("Changes.TotalTimeSpent.Current is %v , want %v", event.Changes.TotalTimeSpent.Current, 8100)
	}

	assert.Equal(t, []*EventLabel{
		{
			ID:          206,
			Title:       "API",
			Color:       "#ffffff",
			ProjectID:   14,
			CreatedAt:   "2013-12-03T17:15:43Z",
			UpdatedAt:   "2013-12-03T17:15:43Z",
			Template:    false,
			Description: "API related issues",
			Type:        "ProjectLabel",
			GroupID:     41,
		},
	}, event.Labels)

	assert.Equal(t, []*EventLabel{
		{
			ID:          206,
			Title:       "API",
			Color:       "#ffffff",
			ProjectID:   14,
			CreatedAt:   "2013-12-03T17:15:43Z",
			UpdatedAt:   "2013-12-03T17:15:43Z",
			Template:    false,
			Description: "API related issues",
			Type:        "ProjectLabel",
			GroupID:     41,
		},
	}, event.Changes.Labels.Previous)

	assert.Equal(t, []*EventLabel{
		{
			ID:          205,
			Title:       "Platform",
			Color:       "#123123",
			ProjectID:   14,
			CreatedAt:   "2013-12-03T17:15:43Z",
			UpdatedAt:   "2013-12-03T17:15:43Z",
			Template:    false,
			Description: "Platform related issues",
			Type:        "ProjectLabel",
			GroupID:     41,
		},
	}, event.Changes.Labels.Current)

	assert.Equal(t, "2017-09-15 16:54:55 UTC", event.Changes.ClosedAt.Previous)
	assert.Equal(t, "2017-09-15 16:56:00 UTC", event.Changes.ClosedAt.Current)

	assert.Equal(t, StateIDNone, event.Changes.StateID.Previous)
	assert.Equal(t, StateIDOpen, event.Changes.StateID.Current)

	assert.Equal(t, "2017-09-15 16:50:55 UTC", event.Changes.UpdatedAt.Previous)
	assert.Equal(t, "2017-09-15 16:52:00 UTC", event.Changes.UpdatedAt.Current)
}

// Generate unit test for MergeCommentEvent
func TestMergeCommentEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/note_merge_request.json")

	var event *MergeCommentEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("Merge Comment Event can not unmarshaled: %v\n ", err.Error())
	}

	if event == nil {
		t.Errorf("Merge Comment Event is null")
	}

	if event.ObjectAttributes.ID != 1244 {
		t.Errorf("ObjectAttributes.ID is %v, want %v", event.ObjectAttributes.ID, 1244)
	}

	if event.ObjectAttributes.Note != "This MR needs work." {
		t.Errorf("ObjectAttributes.Note is %v, want %v", event.ObjectAttributes.Note, "This MR needs work.")
	}

	if event.ObjectAttributes.NoteableType != "MergeRequest" {
		t.Errorf("ObjectAttributes.NoteableType is %v, want %v", event.ObjectAttributes.NoteableType, "MergeRequest")
	}

	if event.ObjectAttributes.AuthorID != 1 {
		t.Errorf("ObjectAttributes.AuthorID is %v, want %v", event.ObjectAttributes.AuthorID, 1)
	}

	if event.ObjectAttributes.CreatedAt != "2015-05-17 18:21:36 UTC" {
		t.Errorf("ObjectAttributes.CreatedAt is %v, want %v", event.ObjectAttributes.CreatedAt, "2015-05-17 18:21:36 UTC")
	}

	if event.ObjectAttributes.UpdatedAt != "2015-05-17 18:21:36 UTC" {
		t.Errorf("ObjectAttributes.UpdatedAt is %v, want %v", event.ObjectAttributes.UpdatedAt, "2015-05-17 18:21:36 UTC")
	}

	if event.ObjectAttributes.ProjectID != 5 {
		t.Errorf("ObjectAttributes.ProjectID is %v, want %v", event.ObjectAttributes.ProjectID, 5)
	}

	if event.MergeRequest.ID != 7 {
		t.Errorf("MergeRequest.ID is %v, want %v", event.MergeRequest.ID, 7)
	}

	if event.MergeRequest.TargetBranch != "markdown" {
		t.Errorf("MergeRequest.TargetBranch is %v, want %v", event.MergeRequest.TargetBranch, "markdown")
	}

	// generate test code for rest of the event.MergeRequest fields
	if event.MergeRequest.SourceBranch != "master" {
		t.Errorf("MergeRequest.SourceBranch is %v, want %v", event.MergeRequest.SourceBranch, "ms-viewport")
	}

	if event.MergeRequest.SourceProjectID != 5 {
		t.Errorf("MergeRequest.SourceProjectID is %v, want %v", event.MergeRequest.SourceProjectID, 5)
	}

	if event.MergeRequest.AuthorID != 8 {
		t.Errorf("MergeRequest.AuthorID is %v, want %v", event.MergeRequest.AuthorID, 8)
	}

	if event.MergeRequest.AssigneeID != 28 {
		t.Errorf("MergeRequest.AssigneeID is %v, want %v", event.MergeRequest.AssigneeID, 28)
	}

	if event.MergeRequest.State != "opened" {
		t.Errorf("MergeRequest.state is %v, want %v", event.MergeRequest.State, "opened")
	}

	if event.MergeRequest.MergeStatus != "cannot_be_merged" {
		t.Errorf("MergeRequest.merge_status is %v, want %v", event.MergeRequest.MergeStatus, "cannot_be_merged")
	}

	if event.MergeRequest.TargetProjectID != 5 {
		t.Errorf("MergeRequest.target_project_id is %v, want %v", event.MergeRequest.TargetProjectID, 5)
	}

	assert.Equal(t, []*EventLabel{
		{
			ID:          206,
			Title:       "Afterpod",
			Color:       "#3e8068",
			ProjectID:   0,
			CreatedAt:   "2019-06-05T14:32:20.211Z",
			UpdatedAt:   "2019-06-05T14:32:20.211Z",
			Template:    false,
			Description: "",
			Type:        "GroupLabel",
			GroupID:     4,
		},
		{
			ID:          86,
			Title:       "Element",
			Color:       "#231afe",
			ProjectID:   4,
			CreatedAt:   "2019-06-05T14:32:20.637Z",
			UpdatedAt:   "2019-06-05T14:32:20.637Z",
			Template:    false,
			Description: "",
			Type:        "ProjectLabel",
			GroupID:     0,
		},
	}, event.MergeRequest.Labels)

	assert.Equal(t, &EventUser{
		ID:        0,
		Name:      "User1",
		Username:  "user1",
		AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon",
		Email:     "",
	}, event.MergeRequest.Assignee)

	if event.MergeRequest.DetailedMergeStatus != "checking" {
		t.Errorf("MergeRequest.DetailedMergeStatus is %v, want %v", event.MergeRequest.DetailedMergeStatus, "checking")
	}
}

func TestMergeEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/merge_request.json")

	var event *MergeEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("Merge Event can not unmarshaled: %v\n ", err.Error())
	}

	if event == nil {
		t.Errorf("Merge Event is null")
	}

	if event.EventType != "merge_request" {
		t.Errorf("EventType is %v, want %v", event.EventType, "merge_request")
	}

	if event.Project.CIConfigPath != "" {
		t.Errorf("Project.CIConfigPath is %v, want %v", event.Project.CIConfigPath, "")
	}

	if event.ObjectAttributes.ID != 99 {
		t.Errorf("ObjectAttributes.ID is %v, want %v", event.ObjectAttributes.ID, 99)
	}

	if event.ObjectAttributes.Source.Homepage != "http://example.com/awesome_space/awesome_project" {
		t.Errorf("ObjectAttributes.Source.Homepage is %v, want %v", event.ObjectAttributes.Source.Homepage, "http://example.com/awesome_space/awesome_project")
	}

	if event.ObjectAttributes.LastCommit.ID != "da1560886d4f094c3e6c9ef40349f7d38b5d27d7" {
		t.Errorf("ObjectAttributes.LastCommit.ID is %v, want %s", event.ObjectAttributes.LastCommit.ID, "da1560886d4f094c3e6c9ef40349f7d38b5d27d7")
	}

	if event.ObjectAttributes.TotalTimeSpent != 0 {
		t.Errorf("ObjectAttributes.TotalTimeSpent is %v, want %v", event.ObjectAttributes.TotalTimeSpent, 0)
	}

	if event.ObjectAttributes.TimeChange != 0 {
		t.Errorf("ObjectAttributes.TimeChange is %v, want %v", event.ObjectAttributes.TimeChange, 0)
	}

	if event.ObjectAttributes.HumanTotalTimeSpent != "30m" {
		t.Errorf("ObjectAttributes.HumanTotalTimeSpent is %v, want %v", event.ObjectAttributes.HumanTotalTimeSpent, "30m")
	}

	if event.ObjectAttributes.HumanTimeChange != "30m" {
		t.Errorf("ObjectAttributes.HumanTimeChange is %v, want %v", event.ObjectAttributes.HumanTimeChange, "30m")
	}

	if event.ObjectAttributes.HumanTimeEstimate != "1h" {
		t.Errorf("ObjectAttributes.HumanTimeEstimate is %v, want %v", event.ObjectAttributes.HumanTimeEstimate, "1h")
	}

	if event.Assignees[0].Name != expectedName {
		t.Errorf("Assignee.Name is %v, want %v", event.Assignees[0].Name, expectedName)
	}

	if event.Assignees[0].Username != expectedUsername {
		t.Errorf("ObjectAttributes is %v, want %v", event.Assignees[0].Username, expectedUsername)
	}

	if event.User.ID != expectedID {
		t.Errorf("User ID is %d, want %d", event.User.ID, expectedID)
	}

	if event.User.Name != expectedName {
		t.Errorf("Username is %s, want %s", event.User.Name, expectedName)
	}

	if event.User.Email != expectedEmail {
		t.Errorf("User email is %s, want %s", event.User.Email, expectedEmail)
	}

	if event.ObjectAttributes.LastCommit.Timestamp == nil {
		t.Errorf("Timestamp isn't nil")
	}

	if name := event.ObjectAttributes.LastCommit.Author.Name; name != "GitLab dev user" {
		t.Errorf("Commit Username is %s, want %s", name, "GitLab dev user")
	}

	if event.ObjectAttributes.BlockingDiscussionsResolved != true {
		t.Errorf("BlockingDiscussionsResolved isn't true")
	}

	if event.ObjectAttributes.FirstContribution != true {
		t.Errorf("FirstContribution isn't true")
	}

	if event.Assignees[0].ID != expectedID {
		t.Errorf("Assignees[0].ID is %v, want %v", event.Assignees[0].ID, expectedID)
	}

	if event.Assignees[0].Name != expectedName {
		t.Errorf("Assignees[0].Name is %v, want %v", event.Assignees[0].Name, expectedName)
	}

	if event.Assignees[0].Username != expectedUsername {
		t.Errorf("Assignees[0].Username is %v, want %v", event.Assignees[0].Username, expectedName)
	}

	if event.Assignees[0].AvatarURL != excpectedAvatar {
		t.Errorf("Assignees[0].AvatarURL is %v, want %v", event.Assignees[0].AvatarURL, excpectedAvatar)
	}

	if len(event.Reviewers) < 1 {
		t.Errorf("Reviewers length is %d, want %d", len(event.Reviewers), 1)
	}

	if event.Reviewers[0].Name != expectedName {
		t.Errorf("Reviewers[0].Name is %v, want %v", event.Reviewers[0].Name, expectedName)
	}

	if event.Reviewers[0].Username != expectedUsername {
		t.Errorf("Reviewer[0].Username is %v, want %v", event.Reviewers[0].Username, expectedUsername)
	}

	if event.Reviewers[0].AvatarURL != excpectedAvatar {
		t.Errorf("Reviewers[0].AvatarURL is %v, want %v", event.Reviewers[0].AvatarURL, excpectedAvatar)
	}

	if event.ObjectAttributes.DetailedMergeStatus != "mergeable" {
		t.Errorf("DetailedMergeStatus is %s, want %s", event.ObjectAttributes.DetailedMergeStatus, "mergeable")
	}

	assert.Equal(t, []*EventLabel{
		{
			ID:          206,
			Title:       "API",
			Color:       "#ffffff",
			ProjectID:   14,
			CreatedAt:   "2013-12-03T17:15:43Z",
			UpdatedAt:   "2013-12-03T17:15:43Z",
			Template:    false,
			Description: "API related issues",
			Type:        "ProjectLabel",
			GroupID:     41,
		},
	}, event.Labels)

	assert.Equal(t, []*EventLabel{
		{
			ID:          206,
			Title:       "API",
			Color:       "#ffffff",
			ProjectID:   14,
			CreatedAt:   "2013-12-03T17:15:43Z",
			UpdatedAt:   "2013-12-03T17:15:43Z",
			Template:    false,
			Description: "API related issues",
			Type:        "ProjectLabel",
			GroupID:     41,
		},
	}, event.ObjectAttributes.Labels)

	assert.Equal(t, []*EventLabel{
		{
			ID:          206,
			Title:       "API",
			Color:       "#ffffff",
			ProjectID:   14,
			CreatedAt:   "2013-12-03T17:15:43Z",
			UpdatedAt:   "2013-12-03T17:15:43Z",
			Template:    false,
			Description: "API related issues",
			Type:        "ProjectLabel",
			GroupID:     41,
		},
	}, event.Changes.Labels.Previous)

	assert.Equal(t, []*EventLabel{
		{
			ID:          205,
			Title:       "Platform",
			Color:       "#123123",
			ProjectID:   14,
			CreatedAt:   "2013-12-03T17:15:43Z",
			UpdatedAt:   "2013-12-03T17:15:43Z",
			Template:    false,
			Description: "Platform related issues",
			Type:        "ProjectLabel",
			GroupID:     41,
		},
	}, event.Changes.Labels.Current)

	assert.Equal(t, StateIDLocked, event.Changes.StateID.Previous)
	assert.Equal(t, StateIDMerged, event.Changes.StateID.Current)

	assert.Equal(t, "2017-09-15 16:50:55 UTC", event.Changes.UpdatedAt.Previous)
	assert.Equal(t, "2017-09-15 16:52:00 UTC", event.Changes.UpdatedAt.Current)
}

func TestMemberEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/member.json")

	var event *MemberEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("Member Event can not unmarshaled: %v\n ", err.Error())
	}

	if event == nil {
		t.Errorf("Member Event is null")
	}

	if event.GroupName != ExpectedGroup {
		t.Errorf("Name is %v, want %v", event.GroupName, ExpectedGroup)
	}

	if event.GroupPath != ExpectedGroup {
		t.Errorf("GroupPath is %v, want %v", event.GroupPath, ExpectedGroup)
	}

	if event.GroupID != 100 {
		t.Errorf(
			"GroupID is %v, want %v", event.GroupID, 100)
	}

	if event.UserUsername != expectedUsername {
		t.Errorf(
			"UserUsername is %v, want %v", event.UserUsername, expectedUsername)
	}

	if event.UserName != expectedName {
		t.Errorf(
			"UserName is %v, want %v", event.UserName, expectedName)
	}

	if event.UserEmail != "testuser@webhooktest.com" {
		t.Errorf(
			"UserEmail is %v, want %v", event.UserEmail, "testuser@webhooktest.com")
	}

	if event.UserID != 64 {
		t.Errorf(
			"UserID is %v, want %v", event.UserID, 64)
	}

	if event.GroupAccess != "Guest" {
		t.Errorf(
			"GroupAccess is %v, want %v", event.GroupAccess, "Guest")
	}

	if event.EventName != expectedEventName {
		t.Errorf(
			"EventName is %v, want %v", event.EventName, expectedEventName)
	}

	if event.CreatedAt.Format(time.RFC3339) != "2020-12-11T04:57:22Z" {
		t.Errorf("CreatedAt is %v, want %v", event.CreatedAt.Format(time.RFC3339), "2020-12-11T04:57:22Z")
	}

	if event.UpdatedAt.Format(time.RFC3339) != "2020-12-11T04:57:22Z" {
		t.Errorf("UpdatedAt is %v, want %v", event.UpdatedAt.Format(time.RFC3339), "2020-12-11T04:57:22Z")
	}

	if event.ExpiresAt.Format(time.RFC3339) != "2020-12-14T00:00:00Z" {
		t.Errorf("ExpiresAt is %v, want %v", event.ExpiresAt.Format(time.RFC3339), "2020-12-14T00:00:00Z")
	}
}

func TestMergeEventUnmarshalFromGroup(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/group_merge_request.json")

	var event *MergeEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("Group Merge Event can not unmarshaled: %v\n ", err.Error())
	}

	if event == nil {
		t.Errorf("Group Merge Event is null")
	}

	if event.ObjectKind != eventObjectKindMergeRequest {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, eventObjectKindMergeRequest)
	}

	if event.User.Username != expectedUsername {
		t.Errorf("User.Username is %v, want %v", event.User.Username, expectedUsername)
	}

	if event.Project.Name != exampleProjectName {
		t.Errorf("Project.Name is %v, want %v", event.Project.Name, exampleProjectName)
	}

	if event.ObjectAttributes.ID != 15917 {
		t.Errorf("ObjectAttributes.ID is %v, want %v", event.ObjectAttributes.ID, 15917)
	}

	if event.ObjectAttributes.Source.Name != exampleProjectName {
		t.Errorf("ObjectAttributes.Source.Name is %v, want %v", event.ObjectAttributes.Source.Name, exampleProjectName)
	}

	if event.ObjectAttributes.LastCommit.Author.Email != "test.user@mail.com" {
		t.Errorf("ObjectAttributes.LastCommit.Author.Email is %v, want %v", event.ObjectAttributes.LastCommit.Author.Email, "test.user@mail.com")
	}

	if event.Repository.Name != exampleProjectName {
		t.Errorf("Repository.Name is %v, want %v", event.Repository.Name, exampleProjectName)
	}

	if event.User.Name != expectedName {
		t.Errorf("Username is %s, want %s", event.User.Name, expectedName)
	}

	if event.ObjectAttributes.LastCommit.Timestamp == nil {
		t.Errorf("Timestamp isn't nil")
	}

	if name := event.ObjectAttributes.LastCommit.Author.Name; name != "Test User" {
		t.Errorf("Commit Username is %s, want %s", name, "Test User")
	}
}

func TestPipelineEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/pipeline.json")

	var event *PipelineEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("Pipeline Event can not unmarshaled: %v\n ", err.Error())
	}

	if event == nil {
		t.Errorf("Pipeline Event is null")
	}

	if event.ObjectAttributes.ID != 31 {
		t.Errorf("ObjectAttributes.ID is %v, want %v", event.ObjectAttributes.ID, 31)
	}

	if event.ObjectAttributes.IID != 123 {
		t.Errorf("ObjectAttributes.IID is %v, want %v", event.ObjectAttributes.ID, 123)
	}

	if event.ObjectAttributes.DetailedStatus != "passed" {
		t.Errorf("ObjectAttributes.DetailedStatus is %v, want %v", event.ObjectAttributes.DetailedStatus, "passed")
	}

	if event.ObjectAttributes.QueuedDuration != 12 {
		t.Errorf("ObjectAttributes.QueuedDuration is %v, want %v", event.ObjectAttributes.QueuedDuration, 12)
	}

	if event.ObjectAttributes.Variables[0].Key != "NESTOR_PROD_ENVIRONMENT" {
		t.Errorf("ObjectAttributes.Variables[0].Key is %v, want %v", event.ObjectAttributes.Variables[0].Key, "NESTOR_PROD_ENVIRONMENT")
	}

	if event.User.ID != 42 {
		t.Errorf("User ID is %d, want %d", event.User.ID, 42)
	}

	if event.User.Name != expectedName {
		t.Errorf("Username is %s, want %s", event.User.Name, expectedName)
	}

	if event.Commit.Timestamp == nil {
		t.Errorf("Timestamp isn't nil")
	}

	if name := event.Commit.Author.Name; name != "User" {
		t.Errorf("Commit Username is %s, want %s", name, "User")
	}

	if len(event.Builds) != 5 {
		t.Errorf("Builds length is %d, want %d", len(event.Builds), 5)
	}

	if event.Builds[0].AllowFailure != true {
		t.Errorf("Builds.0.AllowFailure is %v, want %v", event.Builds[0].AllowFailure, true)
	}

	if event.Builds[0].Environment.Name != "production" {
		t.Errorf("Builds.0.Environment.Name is %v, want %v", event.Builds[0].Environment.Name, "production")
	}

	if event.Builds[0].Duration != 17.1 {
		t.Errorf("Builds[0].Duration is %v, want %v", event.Builds[0].Duration, 17.1)
	}

	if event.Builds[0].QueuedDuration != 3.5 {
		t.Errorf("Builds[0].QueuedDuration is %v, want %v", event.Builds[0].QueuedDuration, 3.5)
	}

	if event.Builds[0].FailureReason != "script_failure" {
		t.Errorf("Builds[0].Failurereason is %v, want %v", event.Builds[0].FailureReason, "script_failure")
	}

	if event.Builds[1].FailureReason != "" {
		t.Errorf("Builds[0].Failurereason is %v, want %v", event.Builds[0].FailureReason, "''")
	}

	if event.SourcePipline.PipelineID != 30 {
		t.Errorf("Source Pipline ID is %v, want %v", event.SourcePipline.PipelineID, 30)
	}

	if event.SourcePipline.JobID != 3401 {
		t.Errorf("Source Pipline JobID is %v, want %v", event.SourcePipline.JobID, 3401)
	}

	if event.SourcePipline.Project.ID != 41 {
		t.Errorf("Source Pipline Project ID is %v, want %v", event.SourcePipline.Project.ID, 41)
	}

	if event.MergeRequest.DetailedMergeStatus != "mergeable" {
		t.Errorf("MergeRequest.DetailedMergeStatus is %v, want %v", event.MergeRequest.DetailedMergeStatus, "mergeable")
	}
}

func TestPushEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/push.json")
	var event *PushEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("Push Event can not unmarshaled: %v\n ", err.Error())
	}

	if event == nil {
		t.Errorf("Push Event is null")
	}

	if event.EventName != "push" {
		t.Errorf("EventName is %v, want %v", event.EventName, "push")
	}

	if event.ProjectID != 15 {
		t.Errorf("ProjectID is %v, want %v", event.ProjectID, 15)
	}

	if event.UserName != exampleEventUserName {
		t.Errorf("Username is %s, want %s", event.UserName, exampleEventUserName)
	}

	if event.Project.ID != 15 {
		t.Errorf("Project.ID is %v, want %v", event.Project.ID, 15)
	}

	if event.Commits[0] == nil || event.Commits[0].Timestamp == nil {
		t.Errorf("Commit Timestamp isn't nil")
	}

	if event.Commits[0] == nil || event.Commits[0].Message != exampleCommitMessage {
		t.Errorf("Commit Message is %s, want %s", event.Commits[0].Message, exampleCommitMessage)
	}

	if event.Commits[0] == nil || event.Commits[0].Title != exampleCommitTitle {
		t.Errorf("Commit Title is %s, want %s", event.Commits[0].Title, exampleCommitTitle)
	}

	if event.Commits[0] == nil || event.Commits[0].Author.Name != "Jordi Mallach" {
		t.Errorf("Commit Username is %s, want %s", event.UserName, "Jordi Mallach")
	}
}

func TestReleaseEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/release.json")

	var event *ReleaseEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("Release Event can not unmarshaled: %v\n ", err.Error())
	}

	if event == nil {
		t.Errorf("Release Event is null")
	}

	if event.Project.ID != 327622 {
		t.Errorf("Project.ID is %v, want %v", event.Project.ID, 327622)
	}

	if event.Commit.Title != "Merge branch 'example-branch' into 'master'" {
		t.Errorf("Commit title is %s, want %s", event.Commit.Title, "Merge branch 'example-branch' into 'master'")
	}

	if len(event.Assets.Sources) != 4 {
		t.Errorf("Asset sources length is %d, want %d", len(event.Assets.Sources), 4)
	}

	if event.Assets.Sources[0].Format != "zip" {
		t.Errorf("First asset source format is %s, want %s", event.Assets.Sources[0].Format, "zip")
	}

	if len(event.Assets.Links) != 1 {
		t.Errorf("Asset links length is %d, want %d", len(event.Assets.Links), 1)
	}

	if event.Assets.Links[0].Name != "Changelog" {
		t.Errorf("First asset link name is %s, want %s", event.Assets.Links[0].Name, "Changelog")
	}

	if event.Commit.Author.Name != "User" {
		t.Errorf("Commit author name is %s, want %s", event.Commit.Author.Name, "User")
	}
}

func TestSubGroupEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/subgroup.json")

	var event *SubGroupEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("SubGroup Event can not unmarshaled: %v\n ", err.Error())
	}

	if event == nil {
		t.Errorf("SubGroup Event is null")
	}

	if event.Name != "SubGroup 1" {
		t.Errorf("Name is %v, want %v", event.Name, "SubGroup 1")
	}

	if event.GroupID != 2 {
		t.Errorf("GroupID is %v, want %v", event.GroupID, 2)
	}

	if event.ParentGroupID != 1 {
		t.Errorf("ParentGroupID is %v, want %v", event.ParentGroupID, 1)
	}

	if event.CreatedAt.Format(time.RFC3339) != "2022-01-24T14:23:59Z" {
		t.Errorf("CreatedAt is %v, want %v", event.CreatedAt.Format(time.RFC3339), "2022-01-24T14:23:59Z")
	}
}

func TestTagEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/tag_push.json")
	var event *TagEvent
	err := json.Unmarshal(jsonObject, &event)
	if err != nil {
		t.Errorf("Tag Event can not unmarshaled: %v\n ", err.Error())
	}

	if event == nil {
		t.Errorf("Tag Event is null")
	}

	if event.EventName != "tag_push" {
		t.Errorf("EventName is %v, want %v", event.EventName, "tag_push")
	}

	if event.ProjectID != 1 {
		t.Errorf("ProjectID is %v, want %v", event.ProjectID, 1)
	}

	if event.Project.ID != 1 {
		t.Errorf("Project.ID is %v, want %v", event.Project.ID, 1)
	}

	if event.UserName != exampleEventUserName {
		t.Errorf("Username is %s, want %s", event.UserName, exampleEventUserName)
	}

	if event.Commits[0] == nil || event.Commits[0].Timestamp == nil {
		t.Errorf("Commit Timestamp isn't nil")
	}

	if event.Commits[0] == nil || event.Commits[0].Message != exampleCommitMessage {
		t.Errorf("Commit Message is %s, want %s", event.Commits[0].Message, exampleCommitMessage)
	}

	if event.Commits[0] == nil || event.Commits[0].Title != exampleCommitTitle {
		t.Errorf("Commit Title is %s, want %s", event.Commits[0].Title, exampleCommitTitle)
	}

	if event.Commits[0] == nil || event.Commits[0].Author.Name != exampleEventUserName {
		t.Errorf("Commit Username is %s, want %s", event.UserName, exampleEventUserName)
	}
}
