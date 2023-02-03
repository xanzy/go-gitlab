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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSystemhookPush(t *testing.T) {
	payload := loadFixture("testdata/systemhooks/push.json")

	parsedEvent, err := ParseSystemhook(payload)
	if err != nil {
		t.Errorf("Error parsing build hook: %s", err)
	}

	event, ok := parsedEvent.(*PushSystemEvent)
	if !ok {
		t.Errorf("Expected PushSystemHookEvent, but parsing produced %T", parsedEvent)
	}
	assert.Equal(t, eventObjectKindPush, event.EventName)
}

func TestParseSystemhookTagPush(t *testing.T) {
	payload := loadFixture("testdata/systemhooks/tag_push.json")

	parsedEvent, err := ParseSystemhook(payload)
	if err != nil {
		t.Errorf("Error parsing build hook: %s", err)
	}

	event, ok := parsedEvent.(*TagPushSystemEvent)
	if !ok {
		t.Errorf("Expected TagPushSystemHookEvent, but parsing produced %T", parsedEvent)
	}
	assert.Equal(t, eventObjectKindTagPush, event.EventName)
}

func TestParseSystemhookMergeRequest(t *testing.T) {
	payload := loadFixture("testdata/systemhooks/merge_request.json")

	parsedEvent, err := ParseSystemhook(payload)
	if err != nil {
		t.Errorf("Error parsing build hook: %s", err)
	}

	event, ok := parsedEvent.(*MergeEvent)
	if !ok {
		t.Errorf("Expected MergeRequestSystemHookEvent, but parsing produced %T", parsedEvent)
	}
	assert.Equal(t, eventObjectKindMergeRequest, event.ObjectKind)
}

func TestParseSystemhookRepositoryUpdate(t *testing.T) {
	payload := loadFixture("testdata/systemhooks/repository_update.json")

	parsedEvent, err := ParseSystemhook(payload)
	if err != nil {
		t.Errorf("Error parsing build hook: %s", err)
	}

	event, ok := parsedEvent.(*RepositoryUpdateSystemEvent)
	if !ok {
		t.Errorf("Expected RepositoryUpdateSystemHookEvent, but parsing produced %T", parsedEvent)
	}
	assert.Equal(t, "repository_update", event.EventName)
}

func TestParseSystemhookProject(t *testing.T) {
	tests := []struct {
		event   string
		payload []byte
	}{
		{"project_create", loadFixture("testdata/systemhooks/project_create.json")},
		{"project_update", loadFixture("testdata/systemhooks/project_update.json")},
		{"project_destroy", loadFixture("testdata/systemhooks/project_destroy.json")},
		{"project_transfer", loadFixture("testdata/systemhooks/project_transfer.json")},
		{"project_rename", loadFixture("testdata/systemhooks/project_rename.json")},
	}
	for _, tc := range tests {
		t.Run(tc.event, func(t *testing.T) {
			parsedEvent, err := ParseSystemhook(tc.payload)
			if err != nil {
				t.Errorf("Error parsing build hook: %s", err)
			}
			event, ok := parsedEvent.(*ProjectSystemEvent)
			if !ok {
				t.Errorf("Expected ProjectSystemHookEvent, but parsing produced %T", parsedEvent)
			}
			assert.Equal(t, tc.event, event.EventName)
		})
	}
}

func TestParseSystemhookGroup(t *testing.T) {
	tests := []struct {
		event   string
		payload []byte
	}{
		{"group_create", loadFixture("testdata/systemhooks/group_create.json")},
		{"group_destroy", loadFixture("testdata/systemhooks/group_destroy.json")},
		{"group_rename", loadFixture("testdata/systemhooks/group_rename.json")},
	}
	for _, tc := range tests {
		t.Run(tc.event, func(t *testing.T) {
			parsedEvent, err := ParseSystemhook(tc.payload)
			if err != nil {
				t.Errorf("Error parsing build hook: %s", err)
			}
			event, ok := parsedEvent.(*GroupSystemEvent)
			if !ok {
				t.Errorf("Expected GroupSystemHookEvent, but parsing produced %T", parsedEvent)
			}
			assert.Equal(t, tc.event, event.EventName)
		})
	}
}

func TestParseSystemhookUser(t *testing.T) {
	tests := []struct {
		event   string
		payload []byte
	}{
		{"user_create", loadFixture("testdata/systemhooks/user_create.json")},
		{"user_destroy", loadFixture("testdata/systemhooks/user_destroy.json")},
		{"user_rename", loadFixture("testdata/systemhooks/user_rename.json")},
		{"user_failed_login", loadFixture("testdata/systemhooks/user_failed_login.json")},
	}
	for _, tc := range tests {
		t.Run(tc.event, func(t *testing.T) {
			parsedEvent, err := ParseSystemhook(tc.payload)
			if err != nil {
				t.Errorf("Error parsing build hook: %s", err)
			}
			event, ok := parsedEvent.(*UserSystemEvent)
			if !ok {
				t.Errorf("Expected UserSystemHookEvent, but parsing produced %T", parsedEvent)
			}
			assert.Equal(t, tc.event, event.EventName)
		})
	}
}

func TestParseSystemhookUserGroup(t *testing.T) {
	tests := []struct {
		event   string
		payload []byte
	}{
		{"user_add_to_group", loadFixture("testdata/systemhooks/user_add_to_group.json")},
		{"user_remove_from_group", loadFixture("testdata/systemhooks/user_remove_from_group.json")},
		{"user_update_for_group", loadFixture("testdata/systemhooks/user_update_for_group.json")},
	}
	for _, tc := range tests {
		t.Run(tc.event, func(t *testing.T) {
			parsedEvent, err := ParseSystemhook(tc.payload)
			if err != nil {
				t.Errorf("Error parsing build hook: %s", err)
			}
			event, ok := parsedEvent.(*UserGroupSystemEvent)
			if !ok {
				t.Errorf("Expected UserGroupSystemHookEvent, but parsing produced %T", parsedEvent)
			}
			assert.Equal(t, tc.event, event.EventName)
		})
	}
}

func TestParseSystemhookUserTeam(t *testing.T) {
	tests := []struct {
		event   string
		payload []byte
	}{
		{"user_add_to_team", loadFixture("testdata/systemhooks/user_add_to_team.json")},
		{"user_remove_from_team", loadFixture("testdata/systemhooks/user_remove_from_team.json")},
		{"user_update_for_team", loadFixture("testdata/systemhooks/user_update_for_team.json")},
	}
	for _, tc := range tests {
		t.Run(tc.event, func(t *testing.T) {
			parsedEvent, err := ParseSystemhook(tc.payload)
			if err != nil {
				t.Errorf("Error parsing build hook: %s", err)
			}
			event, ok := parsedEvent.(*UserTeamSystemEvent)
			if !ok {
				t.Errorf("Expected UserTeamSystemHookEvent, but parsing produced %T", parsedEvent)
			}
			assert.Equal(t, tc.event, event.EventName)
		})
	}
}

func TestParseHookSystemHook(t *testing.T) {
	parsedEvent1, err := ParseHook("System Hook", loadFixture("testdata/systemhooks/merge_request.json"))
	if err != nil {
		t.Errorf("Error parsing build hook: %s", err)
	}
	parsedEvent2, err := ParseSystemhook(loadFixture("testdata/systemhooks/merge_request.json"))
	if err != nil {
		t.Errorf("Error parsing build hook: %s", err)
	}
	assert.Equal(t, parsedEvent1, parsedEvent2)
}
