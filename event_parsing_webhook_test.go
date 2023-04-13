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
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhookEventType(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "https://gitlab.com", nil)
	if err != nil {
		t.Errorf("Error creating HTTP request: %s", err)
	}
	req.Header.Set("X-Gitlab-Event", "Push Hook")

	eventType := HookEventType(req)
	if eventType != "Push Hook" {
		t.Errorf("WebhookEventType is %s, want %s", eventType, "Push Hook")
	}
}

func TestParseBuildHook(t *testing.T) {
	raw := loadFixture("testdata/webhooks/build.json")

	parsedEvent, err := ParseWebhook("Build Hook", raw)
	if err != nil {
		t.Errorf("Error parsing build hook: %s", err)
	}

	event, ok := parsedEvent.(*BuildEvent)
	if !ok {
		t.Errorf("Expected BuildEvent, but parsing produced %T", parsedEvent)
	}

	if event.ObjectKind != "build" {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, "build")
	}

	if event.BuildID != 1977 {
		t.Errorf("BuildID is %v, want %v", event.BuildID, 1977)
	}

	if event.BuildAllowFailure {
		t.Errorf("BuildAllowFailure is %v, want %v", event.BuildAllowFailure, false)
	}

	if event.Commit.SHA != "2293ada6b400935a1378653304eaf6221e0fdb8f" {
		t.Errorf("Commit SHA is %v, want %v", event.Commit.SHA, "2293ada6b400935a1378653304eaf6221e0fdb8f")
	}

	if event.BuildCreatedAt != "2021-02-23T02:41:37.886Z" {
		t.Errorf("BuildCreatedAt is %s, want %s", event.User.Name, expectedName)
	}
}

func TestParseCommitCommentHook(t *testing.T) {
	raw := loadFixture("testdata/webhooks/note_commit.json")

	parsedEvent, err := ParseWebhook("Note Hook", raw)
	if err != nil {
		t.Errorf("Error parsing note hook: %s", err)
	}

	event, ok := parsedEvent.(*CommitCommentEvent)
	if !ok {
		t.Errorf("Expected CommitCommentEvent, but parsing produced %T", parsedEvent)
	}

	if event.ObjectKind != string(NoteEventTargetType) {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, NoteEventTargetType)
	}

	if event.ProjectID != 5 {
		t.Errorf("ProjectID is %v, want %v", event.ProjectID, 5)
	}

	if event.ObjectAttributes.NoteableType != "Commit" {
		t.Errorf("NoteableType is %v, want %v", event.ObjectAttributes.NoteableType, "Commit")
	}

	if event.Commit.ID != "cfe32cf61b73a0d5e9f13e774abde7ff789b1660" {
		t.Errorf("CommitID is %v, want %v", event.Commit.ID, "cfe32cf61b73a0d5e9f13e774abde7ff789b1660")
	}
}

func TestParseFeatureFLagHook(t *testing.T) {
	raw := loadFixture("testdata/webhooks/feature_flag.json")

	parsedEvent, err := ParseWebhook("Feature Flag Hook", raw)
	if err != nil {
		t.Errorf("Error parsing feature flag hook: %s", err)
	}

	event, ok := parsedEvent.(*FeatureFlagEvent)
	if !ok {
		t.Errorf("Expected FeatureFlagEvent, but parsing produced %T", parsedEvent)
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

func TestParseHookWebHook(t *testing.T) {
	parsedEvent1, err := ParseHook("Merge Request Hook", loadFixture("testdata/webhooks/merge_request.json"))
	if err != nil {
		t.Errorf("Error parsing build hook: %s", err)
	}
	parsedEvent2, err := ParseWebhook("Merge Request Hook", loadFixture("testdata/webhooks/merge_request.json"))
	if err != nil {
		t.Errorf("Error parsing build hook: %s", err)
	}
	assert.Equal(t, parsedEvent1, parsedEvent2)
}

func TestParseIssueCommentHook(t *testing.T) {
	raw := loadFixture("testdata/webhooks/note_issue.json")

	parsedEvent, err := ParseWebhook("Note Hook", raw)
	if err != nil {
		t.Errorf("Error parsing note hook: %s", err)
	}

	event, ok := parsedEvent.(*IssueCommentEvent)
	if !ok {
		t.Errorf("Expected IssueCommentEvent, but parsing produced %T", parsedEvent)
	}

	if event.ObjectKind != string(NoteEventTargetType) {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, NoteEventTargetType)
	}

	if event.ProjectID != 5 {
		t.Errorf("ProjectID is %v, want %v", event.ProjectID, 5)
	}

	if event.ObjectAttributes.NoteableType != "Issue" {
		t.Errorf("NoteableType is %v, want %v", event.ObjectAttributes.NoteableType, "Issue")
	}

	if event.Issue.Title != "test_issue" {
		t.Errorf("Issue title is %v, want %v", event.Issue.Title, "test_issue")
	}
	assert.Equal(t, 2, len(event.Issue.Labels))
}

func TestParseIssueHook(t *testing.T) {
	raw := loadFixture("testdata/webhooks/issue.json")

	parsedEvent, err := ParseWebhook("Issue Hook", raw)
	if err != nil {
		t.Errorf("Error parsing issue hook: %s", err)
	}

	event, ok := parsedEvent.(*IssueEvent)
	if !ok {
		t.Errorf("Expected IssueEvent, but parsing produced %T", parsedEvent)
	}

	if event.ObjectKind != "issue" {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, "issue")
	}

	if event.Project.Name != "Gitlab Test" {
		t.Errorf("Project name is %v, want %v", event.Project.Name, "Gitlab Test")
	}

	if event.ObjectAttributes.State != "opened" {
		t.Errorf("Issue state is %v, want %v", event.ObjectAttributes.State, "opened")
	}

	if event.Assignee.Username != "user1" {
		t.Errorf("Assignee username is %v, want %v", event.Assignee.Username, "user1")
	}
	assert.Equal(t, 1, len(event.Labels))
	assert.Equal(t, 0, event.Changes.UpdatedByID.Previous)
	assert.Equal(t, 1, event.Changes.UpdatedByID.Current)
	assert.Equal(t, 1, len(event.Changes.Labels.Previous))
	assert.Equal(t, 1, len(event.Changes.Labels.Current))
	assert.Equal(t, "", event.Changes.Description.Previous)
	assert.Equal(t, "New description", event.Changes.Description.Current)
	assert.Equal(t, "", event.Changes.Title.Previous)
	assert.Equal(t, "New title", event.Changes.Title.Current)
}

func TestParseMergeRequestCommentHook(t *testing.T) {
	raw := loadFixture("testdata/webhooks/note_merge_request.json")

	parsedEvent, err := ParseWebhook("Note Hook", raw)
	if err != nil {
		t.Errorf("Error parsing note hook: %s", err)
	}

	event, ok := parsedEvent.(*MergeCommentEvent)
	if !ok {
		t.Errorf("Expected MergeCommentEvent, but parsing produced %T", parsedEvent)
	}

	if event.ObjectKind != string(NoteEventTargetType) {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, "note")
	}

	if event.ProjectID != 5 {
		t.Errorf("ProjectID is %v, want %v", event.ProjectID, 5)
	}

	if event.ObjectAttributes.NoteableType != "MergeRequest" {
		t.Errorf("NoteableType is %v, want %v", event.ObjectAttributes.NoteableType, "MergeRequest")
	}

	if event.MergeRequest.ID != 7 {
		t.Errorf("MergeRequest ID is %v, want %v", event.MergeRequest.ID, 7)
	}

	expectedTitle := "Merge branch 'another-branch' into 'master'"
	if event.MergeRequest.LastCommit.Title != expectedTitle {
		t.Errorf("MergeRequest Title is %v, want %v", event.MergeRequest.Title, expectedTitle)
	}
}

func TestParseMemberHook(t *testing.T) {
	raw := loadFixture("testdata/webhooks/member.json")

	parsedEvent, err := ParseWebhook("Member Hook", raw)
	if err != nil {
		t.Errorf("Error parsing member hook: %s", err)
	}

	event, ok := parsedEvent.(*MemberEvent)
	if !ok {
		t.Errorf("Expected MemberEvent, but parsing produced %T", parsedEvent)
	}

	if event.EventName != "user_add_to_group" {
		t.Errorf("EventName is %v, want %v", event.EventName, "user_add_to_group")
	}
}

func TestParseMergeRequestHook(t *testing.T) {
	raw := loadFixture("testdata/webhooks/merge_request.json")

	parsedEvent, err := ParseWebhook("Merge Request Hook", raw)
	if err != nil {
		t.Errorf("Error parsing merge request hook: %s", err)
	}

	event, ok := parsedEvent.(*MergeEvent)
	if !ok {
		t.Errorf("Expected MergeEvent, but parsing produced %T", parsedEvent)
	}

	if event.ObjectKind != "merge_request" {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, "merge_request")
	}

	if event.ObjectAttributes.MergeStatus != "unchecked" {
		t.Errorf("MergeStatus is %v, want %v", event.ObjectAttributes.MergeStatus, "unchecked")
	}

	if event.ObjectAttributes.LastCommit.ID != "da1560886d4f094c3e6c9ef40349f7d38b5d27d7" {
		t.Errorf("LastCommit ID is %v, want %v", event.ObjectAttributes.LastCommit.ID, "da1560886d4f094c3e6c9ef40349f7d38b5d27d7")
	}

	if event.ObjectAttributes.WorkInProgress {
		t.Errorf("WorkInProgress is %v, want %v", event.ObjectAttributes.WorkInProgress, false)
	}
	assert.Equal(t, 1, len(event.Labels))
	assert.Equal(t, 0, event.Changes.UpdatedByID.Previous)
	assert.Equal(t, 1, event.Changes.UpdatedByID.Current)
	assert.Equal(t, 1, len(event.Changes.Labels.Previous))
	assert.Equal(t, 1, len(event.Changes.Labels.Current))
}

func TestParsePipelineHook(t *testing.T) {
	raw := loadFixture("testdata/webhooks/pipeline.json")

	parsedEvent, err := ParseWebhook("Pipeline Hook", raw)
	if err != nil {
		t.Errorf("Error parsing pipeline hook: %s", err)
	}

	event, ok := parsedEvent.(*PipelineEvent)
	if !ok {
		t.Errorf("Expected PipelineEvent, but parsing produced %T", parsedEvent)
	}

	if event.ObjectKind != "pipeline" {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, "pipeline")
	}

	if event.ObjectAttributes.Duration != 63 {
		t.Errorf("Duration is %v, want %v", event.ObjectAttributes.Duration, 63)
	}

	if event.Commit.ID != "bcbb5ec396a2c0f828686f14fac9b80b780504f2" {
		t.Errorf("Commit ID is %v, want %v", event.Commit.ID, "bcbb5ec396a2c0f828686f14fac9b80b780504f2")
	}

	if event.Builds[0].ID != 380 {
		t.Errorf("Builds[0] ID is %v, want %v", event.Builds[0].ID, 380)
	}

	if event.Builds[0].Runner.RunnerType != "instance_type" {
		t.Errorf("Builds[0] Runner RunnerType is %v, want %v", event.Builds[0].Runner.RunnerType, "instance_type")
	}
}

func TestParsePushHook(t *testing.T) {
	raw := loadFixture("testdata/webhooks/push.json")

	parsedEvent, err := ParseWebhook("Push Hook", raw)
	if err != nil {
		t.Errorf("Error parsing push hook: %s", err)
	}

	event, ok := parsedEvent.(*PushEvent)
	if !ok {
		t.Errorf("Expected PushEvent, but parsing produced %T", parsedEvent)
	}

	if event.ObjectKind != eventObjectKindPush {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, eventObjectKindPush)
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

func TestParseReleaseHook(t *testing.T) {
	raw := loadFixture("testdata/webhooks/release.json")

	parsedEvent, err := ParseWebhook("Release Hook", raw)
	if err != nil {
		t.Errorf("Error parsing release hook: %s", err)
	}

	event, ok := parsedEvent.(*ReleaseEvent)
	if !ok {
		t.Errorf("Expected ReleaseEvent, but parsing produced %T", parsedEvent)
	}

	if event.ObjectKind != "release" {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, "release")
	}

	if event.Project.Name != "Project Name" {
		t.Errorf("Project name is %v, want %v", event.Project.Name, "Project Name")
	}
}

func TestParseServiceWebHook(t *testing.T) {
	parsedEvent, err := ParseWebhook("Service Hook", loadFixture("testdata/webhooks/service_merge_request.json"))
	if err != nil {
		t.Errorf("Error parsing service hook merge request: %s", err)
	}

	switch event := parsedEvent.(type) {
	case *MergeEvent:
		assert.EqualValues(t, &EventUser{
			ID:        2,
			Name:      "the test",
			Username:  "test",
			Email:     "test@test.test",
			AvatarURL: "https://www.gravatar.com/avatar/dd46a756faad4727fb679320751f6dea?s=80&d=identicon",
		}, event.User)
		assert.EqualValues(t, "unchecked", event.ObjectAttributes.MergeStatus)
		assert.EqualValues(t, "next-feature", event.ObjectAttributes.SourceBranch)
		assert.EqualValues(t, "master", event.ObjectAttributes.TargetBranch)
	default:
		t.Errorf("unexpected event type: %s", reflect.TypeOf(parsedEvent))
	}
}

func TestParseSnippetCommentHook(t *testing.T) {
	raw := loadFixture("testdata/webhooks/note_snippet.json")

	parsedEvent, err := ParseWebhook("Note Hook", raw)
	if err != nil {
		t.Errorf("Error parsing note hook: %s", err)
	}

	event, ok := parsedEvent.(*SnippetCommentEvent)
	if !ok {
		t.Errorf("Expected SnippetCommentEvent, but parsing produced %T", parsedEvent)
	}

	if event.ObjectKind != string(NoteEventTargetType) {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, NoteEventTargetType)
	}

	if event.ProjectID != 5 {
		t.Errorf("ProjectID is %v, want %v", event.ProjectID, 5)
	}

	if event.ObjectAttributes.NoteableType != "Snippet" {
		t.Errorf("NoteableType is %v, want %v", event.ObjectAttributes.NoteableType, "Snippet")
	}

	if event.Snippet.Title != "test" {
		t.Errorf("Snippet title is %v, want %v", event.Snippet.Title, "test")
	}
}

func TestParseSubGroupHook(t *testing.T) {
	raw := loadFixture("testdata/webhooks/subgroup.json")

	parsedEvent, err := ParseWebhook("Subgroup Hook", raw)
	if err != nil {
		t.Errorf("Error parsing subgroup hook: %s", err)
	}

	event, ok := parsedEvent.(*SubGroupEvent)
	if !ok {
		t.Errorf("Expected SubGroupEvent, but parsing produced %T", parsedEvent)
	}

	if event.EventName != "subgroup_create" {
		t.Errorf("EventName is %v, want %v", event.EventName, "subgroup_create")
	}
}

func TestParseTagHook(t *testing.T) {
	raw := loadFixture("testdata/webhooks/tag_push.json")

	parsedEvent, err := ParseWebhook("Tag Push Hook", raw)
	if err != nil {
		t.Errorf("Error parsing tag hook: %s", err)
	}

	event, ok := parsedEvent.(*TagEvent)
	if !ok {
		t.Errorf("Expected TagEvent, but parsing produced %T", parsedEvent)
	}

	if event.ObjectKind != eventObjectKindTagPush {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, eventObjectKindTagPush)
	}

	if event.ProjectID != 1 {
		t.Errorf("ProjectID is %v, want %v", event.ProjectID, 1)
	}

	if event.UserName != exampleEventUserName {
		t.Errorf("Name is %s, want %s", event.UserName, exampleEventUserName)
	}

	if event.UserUsername != exampleEventUserUsername {
		t.Errorf("Username is %s, want %s", event.UserUsername, exampleEventUserUsername)
	}

	if event.Ref != "refs/tags/v1.0.0" {
		t.Errorf("Ref is %s, want %s", event.Ref, "refs/tags/v1.0.0")
	}
}

func TestParseWikiPageHook(t *testing.T) {
	raw := loadFixture("testdata/webhooks/wiki_page.json")

	parsedEvent, err := ParseWebhook("Wiki Page Hook", raw)
	if err != nil {
		t.Errorf("Error parsing wiki page hook: %s", err)
	}

	event, ok := parsedEvent.(*WikiPageEvent)
	if !ok {
		t.Errorf("Expected WikiPageEvent, but parsing produced %T", parsedEvent)
	}

	if event.ObjectKind != "wiki_page" {
		t.Errorf("ObjectKind is %v, want %v", event.ObjectKind, "wiki_page")
	}

	if event.Project.Name != "awesome-project" {
		t.Errorf("Project name is %v, want %v", event.Project.Name, "awesome-project")
	}

	if event.Wiki.WebURL != "http://example.com/root/awesome-project/wikis/home" {
		t.Errorf("Wiki web URL is %v, want %v", event.Wiki.WebURL, "http://example.com/root/awesome-project/wikis/home")
	}

	if event.ObjectAttributes.Message != "adding an awesome page to the wiki" {
		t.Errorf("Message is %v, want %v", event.ObjectAttributes.Message, "adding an awesome page to the wiki")
	}
}
