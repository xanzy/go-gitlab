package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const eventTypeHeader = "X-Gitlab-Event"

// WebhookType returns the event type for the given request.
func WebhookType(r *http.Request) string {
	return r.Header.Get(eventTypeHeader)
}

// ParseWebhook parses the event payload. For recognized event types, a
// value of the corresponding struct type will be returned. An error will
// be returned for unrecognized event types.
//
// Example usage:
//
//     func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//       payload, err := ioutil.ReadAll(r.Body)
//       if err != nil { ... }
//       event, err := gitlab.ParseWebhook(gitlab.WebhookType(r), payload)
//       if err != nil { ... }
//       switch event := event.(type) {
//       case *gitlab.PushEvent:
//           processPushEvent(event)
//       case *gitlab.MergeEvent:
//           processMergeEvent(event)
//       ...
//       }
//     }
//
func ParseWebhook(typeHeaderValue string, payload []byte) (interface{}, error) {
	eventType, ok := eventTypeMapping[typeHeaderValue]
	if !ok {
		return nil, fmt.Errorf("unknown value for %s header: %s", eventTypeHeader, typeHeaderValue)
	}

	event := rawEvent{
		Type:       eventType,
		RawPayload: payload,
	}
	return event.ParsePayload()
}

var eventTypeMapping = map[string]string{
	"Push Hook":          "PushEvent",
	"Tag Push Hook":      "TagEvent",
	"Issue Hook":         "IssueEvent",
	"Note Hook":          "ambiguousCommentEvent",
	"Merge Request Hook": "MergeEvent",
	"Wiki Page Hook":     "WikiPageEvent",
	"Pipeline Hook":      "PipelineEvent",
	"Build Hook":         "BuildEvent",
}

type rawEvent struct {
	Type       string
	RawPayload json.RawMessage
}

type ambiguousCommentEvent struct {
	ObjectKind       string `json:"object_kind"`
	ObjectAttributes struct {
		NoteableType string `json:"noteable_type"`
	} `json:"object_attributes"`
}

func (e *rawEvent) ParsePayload() (event interface{}, err error) {
	switch e.Type {
	case "PushEvent":
		event = &PushEvent{}
	case "TagEvent":
		event = &TagEvent{}
	case "IssueEvent":
		event = &IssueEvent{}
	case "ambiguousCommentEvent":
		ambi := &ambiguousCommentEvent{}
		err := json.Unmarshal(e.RawPayload, ambi)
		if err != nil {
			return nil, err
		} else if ambi.ObjectKind != "note" {
			return nil, fmt.Errorf("unexpected object kind %s", ambi.ObjectKind)
		}
		switch ambi.ObjectAttributes.NoteableType {
		case "Commit":
			event = &CommitCommentEvent{}
		case "MergeRequest":
			event = &MergeCommentEvent{}
		case "Issue":
			event = &IssueCommentEvent{}
		case "Snippet":
			event = &SnippetCommentEvent{}
		default:
			return nil, fmt.Errorf("unexpected noteable type %s", ambi.ObjectAttributes.NoteableType)
		}
	case "MergeEvent":
		event = &MergeEvent{}
	case "WikiPageEvent":
		event = &WikiPageEvent{}
	case "PipelineEvent":
		event = &PipelineEvent{}
	case "BuildEvent":
		event = &BuildEvent{}
	default:
		return nil, fmt.Errorf("unexpected event type: %s", e.Type)
	}

	err = json.Unmarshal(e.RawPayload, event)
	return event, err
}
