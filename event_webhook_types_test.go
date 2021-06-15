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
)

const (
	expectedID       = 1
	expectedName     = "User1"
	expectedUsername = "user1"
	excpectedAvatar  = "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon"
	expectedEmail    = "test.user1@example.com"
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

	if len(event.Issue.Labels) == 0 || event.Issue.Labels[0].ID != 25 {
		t.Errorf("Label id is null")
	}
}

func TestIssueEventUnmarshal(t *testing.T) {
	jsonObject := loadFixture("testdata/webhooks/issue.json")

	var event *IssueEvent
	err := json.Unmarshal(jsonObject, &event)

	if err != nil {
		t.Errorf("Issue Event can not unmarshaled: %v\n ", err.Error())
	}

	if event.Project.ID != 1 {
		t.Errorf("Project.ID is %v, want %v", event.Project.ID, 1)
	}

	if event.User.ID != 42 {
		t.Errorf("User ID is %d, want %d", event.User.ID, 42)
	}

	if event.Assignee.Username != "user1" {
		t.Errorf("Assignee username is %s, want %s", event.Assignee.Username, "user1")
	}

	if event.Changes.TotalTimeSpent.Previous != 8100 {
		t.Errorf("Changes.TotalTimeSpent.Previous is %v , want %v", event.Changes.TotalTimeSpent.Previous, 8100)
	}

	if event.Changes.TotalTimeSpent.Current != 9900 {
		t.Errorf("Changes.TotalTimeSpent.Current is %v , want %v", event.Changes.TotalTimeSpent.Current, 8100)
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

	if event.ObjectAttributes.ID != 99 {
		t.Errorf("ObjectAttributes.ID is %v, want %v", event.ObjectAttributes.ID, 99)
	}

	if event.ObjectAttributes.Source.Homepage != "http://example.com/awesome_space/awesome_project" {
		t.Errorf("ObjectAttributes.Source.Homepage is %v, want %v", event.ObjectAttributes.Source.Homepage, "http://example.com/awesome_space/awesome_project")
	}

	if event.ObjectAttributes.LastCommit.ID != "da1560886d4f094c3e6c9ef40349f7d38b5d27d7" {
		t.Errorf("ObjectAttributes.LastCommit.ID is %v, want %s", event.ObjectAttributes.LastCommit.ID, "da1560886d4f094c3e6c9ef40349f7d38b5d27d7")
	}
	if event.ObjectAttributes.Assignee.Name != expectedName {
		t.Errorf("Assignee.Name is %v, want %v", event.ObjectAttributes.Assignee.Name, expectedName)
	}

	if event.ObjectAttributes.Assignee.Username != expectedUsername {
		t.Errorf("ObjectAttributes is %v, want %v", event.ObjectAttributes.Assignee.Username, expectedUsername)
	}

	if event.User.ID != 42 {
		t.Errorf("User ID is %d, want %d", event.User.ID, 42)
	}

	if event.User.Name != expectedName {
		t.Errorf("Username is %s, want %s", event.User.Name, expectedName)
	}

	if event.User.Email != "user1@example.com" {
		t.Errorf("User email is %s, want %s", event.User.Email, "user1@example.com")
	}

	if event.ObjectAttributes.LastCommit.Timestamp == nil {
		t.Errorf("Timestamp isn't nil")
	}

	if name := event.ObjectAttributes.LastCommit.Author.Name; name != "GitLab dev user" {
		t.Errorf("Commit Username is %s, want %s", name, "GitLab dev user")
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
		t.Errorf("Assignees[0].Email is %v, want %v", event.Assignees[0].AvatarURL, excpectedAvatar)
	}

	if event.Assignees[0].Email != expectedEmail {
		t.Errorf("Assignees[0].Email is %v, want %v", event.Assignees[0].Email, expectedEmail)
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

	if event.ObjectKind != "merge_request" {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, "merge_request")
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

	if event.Assignee.Username != expectedUsername {
		t.Errorf("Assignee.Username is %v, want %v", event.Assignee, expectedUsername)
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
		t.Errorf("ObjectAttributes is %v, want %v", event.ObjectAttributes.ID, 1977)
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

	if event.ProjectID != 15 {
		t.Errorf("ProjectID is %v, want %v", event.ProjectID, 15)
	}

	if event.UserName != exampleEventUserName {
		t.Errorf("Username is %s, want %s", event.UserName, exampleEventUserName)
	}

	if event.Commits[0] == nil || event.Commits[0].Timestamp == nil {
		t.Errorf("Commit Timestamp isn't nil")
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
