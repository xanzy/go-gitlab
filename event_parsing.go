package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const eventTypeHeader = "X-Gitlab-Event"

// webhookEventType returns the event type for the given request.
func webhookEventType(r *http.Request) string {
	return r.Header.Get(eventTypeHeader)
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
//     event, err := gitlab.ParseWebhook(gitlab.webhookEventType(r), payload)
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
func ParseWebhook(typeHeaderValue string, payload []byte) (interface{}, error) {
	t, ok := eventTypeMapping[typeHeaderValue]
	if !ok {
		return nil, fmt.Errorf("unexpected value for %s header: %s", eventTypeHeader, typeHeaderValue)
	}

	event := rawWebhookEvent{
		Type:       t,
		RawPayload: payload,
	}
	return event.ParsePayload()
}

type eventStructType string

const (
	eventStructTypePushEvent             eventStructType = "PushEvent"
	eventStructTypeTagEvent              eventStructType = "TagEvent"
	eventStructTypeIssueEvent            eventStructType = "IssueEvent"
	eventStructTypeAmbiguousCommentEvent eventStructType = "ambiguousCommentEvent"
	eventStructTypeMergeEvent            eventStructType = "MergeEvent"
	eventStructTypeWikiPageEvent         eventStructType = "WikiPageEvent"
	eventStructTypePipelineEvent         eventStructType = "PipelineEvent"
	eventStructTypeBuildEvent            eventStructType = "BuildEvent"
)

// Keys here are possible values for the X-Gitlab-Event header.
var eventTypeMapping = map[string]eventStructType{
	"Push Hook":          eventStructTypePushEvent,
	"Tag Push Hook":      eventStructTypeTagEvent,
	"Issue Hook":         eventStructTypeIssueEvent,
	"Note Hook":          eventStructTypeAmbiguousCommentEvent,
	"Merge Request Hook": eventStructTypeMergeEvent,
	"Wiki Page Hook":     eventStructTypeWikiPageEvent,
	"Pipeline Hook":      eventStructTypePipelineEvent,
	"Build Hook":         eventStructTypeBuildEvent,
}

type rawWebhookEvent struct {
	Type       eventStructType
	RawPayload json.RawMessage
}

const (
	noteableTypeCommit       = "Commit"
	noteableTypeMergeRequest = "MergeRequest"
	noteableTypeIssue        = "Issue"
	noteableTypeSnippet      = "Snippet"
)

type ambiguousCommentEvent struct {
	ObjectKind       string `json:"object_kind"`
	ObjectAttributes struct {
		NoteableType string `json:"noteable_type"`
	} `json:"object_attributes"`
}

func (e *rawWebhookEvent) ParsePayload() (event interface{}, err error) {
	switch e.Type {
	case eventStructTypePushEvent:
		event = &PushEvent{}
	case eventStructTypeTagEvent:
		event = &TagEvent{}
	case eventStructTypeIssueEvent:
		event = &IssueEvent{}
	case eventStructTypeAmbiguousCommentEvent:
		ambi := &ambiguousCommentEvent{}
		err := json.Unmarshal(e.RawPayload, ambi)
		if err != nil {
			return nil, err
		} else if ambi.ObjectKind != "note" {
			return nil, fmt.Errorf("unexpected object kind %s", ambi.ObjectKind)
		}
		switch ambi.ObjectAttributes.NoteableType {
		case noteableTypeCommit:
			event = &CommitCommentEvent{}
		case noteableTypeMergeRequest:
			event = &MergeCommentEvent{}
		case noteableTypeIssue:
			event = &IssueCommentEvent{}
		case noteableTypeSnippet:
			event = &SnippetCommentEvent{}
		default:
			return nil, fmt.Errorf("unexpected noteable type %s", ambi.ObjectAttributes.NoteableType)
		}
	case eventStructTypeMergeEvent:
		event = &MergeEvent{}
	case eventStructTypeWikiPageEvent:
		event = &WikiPageEvent{}
	case eventStructTypePipelineEvent:
		event = &PipelineEvent{}
	case eventStructTypeBuildEvent:
		event = &BuildEvent{}
	default:
		return nil, fmt.Errorf("unexpected event type: %s", e.Type)
	}

	err = json.Unmarshal(e.RawPayload, event)
	return event, err
}
