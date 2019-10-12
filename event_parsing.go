package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// EventType represents a Gitlab event type.
type EventType string

// List of available event types.
const (
	EventTypeBuild        EventType = "Build Hook"
	EventTypeIssue        EventType = "Issue Hook"
	EventTypeJob          EventType = "Job Hook"
	EventTypeMergeRequest EventType = "Merge Request Hook"
	EventTypeNote         EventType = "Note Hook"
	EventTypePipeline     EventType = "Pipeline Hook"
	EventTypePush         EventType = "Push Hook"
	EventTypeTagPush      EventType = "Tag Push Hook"
	EventTypeWikiPage     EventType = "Wiki Page Hook"
	EventSystemHook       EventType = "System Hook"
)

const (
	noteableTypeCommit       = "Commit"
	noteableTypeMergeRequest = "MergeRequest"
	noteableTypeIssue        = "Issue"
	noteableTypeSnippet      = "Snippet"
)

type noteEvent struct {
	ObjectKind       string `json:"object_kind"`
	ObjectAttributes struct {
		NoteableType string `json:"noteable_type"`
	} `json:"object_attributes"`
}

const eventTypeHeader = "X-Gitlab-Event"

// WebhookEventType returns the event type for the given request.
func WebhookEventType(r *http.Request) EventType {
	return EventType(r.Header.Get(eventTypeHeader))
}

// ParseWebhook parses the event payload. For recognized event types, a
// value of the corresponding struct type will be returned. An error will
// be returned for unrecognized event types.
//
// Example usage:
//
// func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//     payload, err := ioutil.ReadAll(r.Body)
//     if err != nil { ... }
//     event, err := gitlab.ParseWebhook(gitlab.WebhookEventType(r), payload)
//     if err != nil { ... }
//     switch event := event.(type) {
//     case *gitlab.PushEvent:
//         processPushEvent(event)
//     case *gitlab.MergeEvent:
//         processMergeEvent(event)
//     ...
//     }
// }
//
func ParseWebhook(eventType EventType, payload []byte) (event interface{}, err error) {
	switch eventType {
	case EventTypeBuild:
		event = &BuildEvent{}
	case EventTypeIssue:
		event = &IssueEvent{}
	case EventTypeJob:
		event = &JobEvent{}
	case EventTypeMergeRequest:
		event = &MergeEvent{}
	case EventTypePipeline:
		event = &PipelineEvent{}
	case EventTypePush:
		event = &PushEvent{}
	case EventTypeTagPush:
		event = &TagEvent{}
	case EventTypeWikiPage:
		event = &WikiPageEvent{}
	case EventTypeNote:
		note := &noteEvent{}
		err := json.Unmarshal(payload, note)
		if err != nil {
			return nil, err
		}

		if note.ObjectKind != "note" {
			return nil, fmt.Errorf("unexpected object kind %s", note.ObjectKind)
		}

		switch note.ObjectAttributes.NoteableType {
		case noteableTypeCommit:
			event = &CommitCommentEvent{}
		case noteableTypeMergeRequest:
			event = &MergeCommentEvent{}
		case noteableTypeIssue:
			event = &IssueCommentEvent{}
		case noteableTypeSnippet:
			event = &SnippetCommentEvent{}
		default:
			return nil, fmt.Errorf("unexpected noteable type %s", note.ObjectAttributes.NoteableType)
		}

	default:
		return nil, fmt.Errorf("unexpected event type: %s", eventType)
	}

	if err := json.Unmarshal(payload, event); err != nil {
		return nil, err
	}

	return event, nil
}

// ParseSystemhook parses the System Hook event payload. For recognized event types, a
// value of the corresponding struct type will be returned. An error will
// be returned for unrecognized event types.
func ParseSystemhook(eventType EventType, payload []byte) (event interface{}, err error) {
	switch eventType {
	case EventSystemHook:
		e := &systemHookEvent{}
		err := json.Unmarshal(payload, e)
		if err != nil {
			return nil, err
		}

		switch e.EventName {
		case "push":
			event = &PushSystemHookEvent{}
		case "tag_push":
			event = &TagPushSystemHookEvent{}
		case "repository_update":
			event = &RepositoryUpdateSystemHookEvent{}
		case "project_create":
			fallthrough
		case "project_update":
			fallthrough
		case "project_destroy":
			fallthrough
		case "project_transfer":
			fallthrough
		case "project_rename":
			event = &ProjectSystemHookEvent{}
		case "group_create":
			fallthrough
		case "group_destroy":
			fallthrough
		case "group_rename":
			event = &GroupSystemHookEvent{}
		case "key_create":
			fallthrough
		case "key_destroy":
			event = &KeySystemHookEvent{}
		case "user_create":
			fallthrough
		case "user_destroy":
			fallthrough
		case "user_rename":
			event = &UserSystemHookEvent{}
		case "user_add_to_group":
			fallthrough
		case "user_remove_from_group":
			fallthrough
		case "user_update_for_group":
			event = &UserGroupSystemHookEvent{}
		case "user_add_to_team":
			fallthrough
		case "user_remove_from_team":
			fallthrough
		case "user_update_for_team":
			event = &UserTeamSystemHookEvent{}
		default:
			switch e.ObjectKind {
			case "merge_request":
				event = &MergeEvent{}
			default:
				return nil, fmt.Errorf("unexpected system hook type %s", e.EventName)
			}
		}

	default:
		return nil, fmt.Errorf("unexpected event type: %s", eventType)
	}

	if err := json.Unmarshal(payload, event); err != nil {
		return nil, err
	}

	return event, nil
}

// ParseHook processes both ParseWebhook & ParseSystemhook
func ParseHook(eventType EventType, payload []byte) (event interface{}, err error) {
	if event, err = ParseWebhook(eventType, payload); err != nil {
		event, err = ParseSystemhook(eventType, payload)
		return
	}
	return
}
