package gitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSystemhookPush(t *testing.T) {
	payload := loadFixture("testdata/systemhooks/push.json")

	parsedEvent, err := ParseSystemhook("System Hook", payload)
	if err != nil {
		t.Errorf("Error parsing build hook: %s", err)
	}

	event, ok := parsedEvent.(*PushSystemHookEvent)
	if !ok {
		t.Errorf("Expected PushSystemHookEvent, but parsing produced %T", parsedEvent)
	}
	assert.Equal(t, "push", event.EventName)
}

func TestParseSystemhookTagPush(t *testing.T) {
	payload := loadFixture("testdata/systemhooks/tag_push.json")

	parsedEvent, err := ParseSystemhook("System Hook", payload)
	if err != nil {
		t.Errorf("Error parsing build hook: %s", err)
	}

	event, ok := parsedEvent.(*TagPushSystemHookEvent)
	if !ok {
		t.Errorf("Expected TagPushSystemHookEvent, but parsing produced %T", parsedEvent)
	}
	assert.Equal(t, "tag_push", event.EventName)
}

func TestParseSystemhookMergeRequest(t *testing.T) {
	payload := loadFixture("testdata/systemhooks/merge_request.json")

	parsedEvent, err := ParseSystemhook("System Hook", payload)
	if err != nil {
		t.Errorf("Error parsing build hook: %s", err)
	}

	event, ok := parsedEvent.(*MergeEvent)
	if !ok {
		t.Errorf("Expected MergeRequestSystemHookEvent, but parsing produced %T", parsedEvent)
	}
	assert.Equal(t, "merge_request", event.ObjectKind)
}

func TestParseSystemhookRepositoryUpdate(t *testing.T) {
	payload := loadFixture("testdata/systemhooks/repository_update.json")

	parsedEvent, err := ParseSystemhook("System Hook", payload)
	if err != nil {
		t.Errorf("Error parsing build hook: %s", err)
	}

	event, ok := parsedEvent.(*RepositoryUpdateSystemHookEvent)
	if !ok {
		t.Errorf("Expected RepositoryUpdateSystemHookEvent, but parsing produced %T", parsedEvent)
	}
	assert.Equal(t, "repository_update", event.EventName)
}

func TestParseSystemhookProject(t *testing.T) {
	var tests = []struct {
		event   string
		payload []byte
	}{
		{"project_create", loadFixture("testdata/systemhooks/project_create.json")},
		{"project_update", loadFixture("testdata/systemhooks/project_update.json")},
		{"project_destroy", loadFixture("testdata/systemhooks/project_destroy.json")},
		{"project_transfer", loadFixture("testdata/systemhooks/project_transfer.json")},
		{"project_rename", loadFixture("testdata/systemhooks/project_rename.json")},
	}
	for _, tt := range tests {
		t.Run(tt.event, func(t *testing.T) {
			parsedEvent, err := ParseSystemhook("System Hook", tt.payload)
			if err != nil {
				t.Errorf("Error parsing build hook: %s", err)
			}
			event, ok := parsedEvent.(*ProjectSystemHookEvent)
			if !ok {
				t.Errorf("Expected ProjectSystemHookEvent, but parsing produced %T", parsedEvent)
			}
			assert.Equal(t, tt.event, event.EventName)
		})
	}
}

func TestParseSystemhookGroup(t *testing.T) {
	var tests = []struct {
		event   string
		payload []byte
	}{
		{"group_create", loadFixture("testdata/systemhooks/group_create.json")},
		{"group_destroy", loadFixture("testdata/systemhooks/group_destroy.json")},
		{"group_rename", loadFixture("testdata/systemhooks/group_rename.json")},
	}
	for _, tt := range tests {
		t.Run(tt.event, func(t *testing.T) {
			parsedEvent, err := ParseSystemhook("System Hook", tt.payload)
			if err != nil {
				t.Errorf("Error parsing build hook: %s", err)
			}
			event, ok := parsedEvent.(*GroupSystemHookEvent)
			if !ok {
				t.Errorf("Expected GroupSystemHookEvent, but parsing produced %T", parsedEvent)
			}
			assert.Equal(t, tt.event, event.EventName)
		})
	}
}

func TestParseSystemhookUser(t *testing.T) {
	var tests = []struct {
		event   string
		payload []byte
	}{
		{"user_create", loadFixture("testdata/systemhooks/user_create.json")},
		{"user_destroy", loadFixture("testdata/systemhooks/user_destroy.json")},
		{"user_rename", loadFixture("testdata/systemhooks/user_rename.json")},
	}
	for _, tt := range tests {
		t.Run(tt.event, func(t *testing.T) {
			parsedEvent, err := ParseSystemhook("System Hook", tt.payload)
			if err != nil {
				t.Errorf("Error parsing build hook: %s", err)
			}
			event, ok := parsedEvent.(*UserSystemHookEvent)
			if !ok {
				t.Errorf("Expected UserSystemHookEvent, but parsing produced %T", parsedEvent)
			}
			assert.Equal(t, tt.event, event.EventName)
		})
	}
}

func TestParseSystemhookUserGroup(t *testing.T) {
	var tests = []struct {
		event   string
		payload []byte
	}{
		{"user_add_to_group", loadFixture("testdata/systemhooks/user_add_to_group.json")},
		{"user_remove_from_group", loadFixture("testdata/systemhooks/user_remove_from_group.json")},
		{"user_update_for_group", loadFixture("testdata/systemhooks/user_update_for_group.json")},
	}
	for _, tt := range tests {
		t.Run(tt.event, func(t *testing.T) {
			parsedEvent, err := ParseSystemhook("System Hook", tt.payload)
			if err != nil {
				t.Errorf("Error parsing build hook: %s", err)
			}
			event, ok := parsedEvent.(*UserGroupSystemHookEvent)
			if !ok {
				t.Errorf("Expected UserGroupSystemHookEvent, but parsing produced %T", parsedEvent)
			}
			assert.Equal(t, tt.event, event.EventName)
		})
	}
}

func TestParseSystemhookUserTeam(t *testing.T) {
	var tests = []struct {
		event   string
		payload []byte
	}{
		{"user_add_to_team", loadFixture("testdata/systemhooks/user_add_to_team.json")},
		{"user_remove_from_team", loadFixture("testdata/systemhooks/user_remove_from_team.json")},
		{"user_update_for_team", loadFixture("testdata/systemhooks/user_update_for_team.json")},
	}
	for _, tt := range tests {
		t.Run(tt.event, func(t *testing.T) {
			parsedEvent, err := ParseSystemhook("System Hook", tt.payload)
			if err != nil {
				t.Errorf("Error parsing build hook: %s", err)
			}
			event, ok := parsedEvent.(*UserTeamSystemHookEvent)
			if !ok {
				t.Errorf("Expected UserTeamSystemHookEvent, but parsing produced %T", parsedEvent)
			}
			assert.Equal(t, tt.event, event.EventName)
		})
	}
}

func TestParseHookSystemHook(t *testing.T) {
	parsedEvent1, err := ParseHook("System Hook", loadFixture("testdata/systemhooks/merge_request.json"))
	if err != nil {
		t.Errorf("Error parsing build hook: %s", err)
	}
	parsedEvent2, err := ParseSystemhook("System Hook", loadFixture("testdata/systemhooks/merge_request.json"))
	if err != nil {
		t.Errorf("Error parsing build hook: %s", err)
	}
	assert.Equal(t, parsedEvent1, parsedEvent2)
}
