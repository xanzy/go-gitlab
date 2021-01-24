package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/xanzy/go-gitlab"
)

// webhook is a HTTP Handler for Gitlab Webhook events
type webhook struct {
	Secret         string
	EventsToAccept []gitlab.EventType
}

// webhookExample shows how to create a Webhook server to parse Gitlab events
// Listens /webhook endpoint on host 0.0.0.0 port 8080
func webhookExample() {
	wh := webhook{
		Secret: "your-gitlab-secret",
		// Define which types of events we want to subscribe to in our case, Merge Request and Pipeline
		EventsToAccept: []gitlab.EventType{gitlab.EventTypeMergeRequest, gitlab.EventTypePipeline},
	}

	mux := http.NewServeMux()
	mux.Handle("/webhook", wh)
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

// ServeHTTP tries to parse Gitlab events sent and calls handle function with the successfully parsed events
// Returns 500 on any error
// Returns 204 on successfully parsed event
func (hook webhook) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	event, err := hook.parse(request)
	if err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(fmt.Sprintf("could parse the webhook event: %v", err)))
		return
	}

	writer.WriteHeader(204)
	_, _ = writer.Write([]byte("no content"))

	hook.handle(event)
}

func (hook webhook) handle(event interface{}) {
	str, _ := json.Marshal(event)
	fmt.Println(string(str))
}

// parse verifies and parses the events specified in the r and returns the parsed event or an error
func (hook webhook) parse(r *http.Request) (interface{}, error) {
	defer func() {
		_, _ = io.Copy(ioutil.Discard, r.Body)
		_ = r.Body.Close()
	}()

	if r.Method != http.MethodPost {
		return nil, errors.New("invalid HTTP Method")
	}

	// If we have a Secret set, we should check if the request matches it
	if len(hook.Secret) > 0 {
		signature := r.Header.Get("X-Gitlab-Token")
		if signature != hook.Secret {
			return nil, errors.New("X-Gitlab-Token validation failed")
		}
	}

	event := r.Header.Get("X-Gitlab-Event")
	if strings.TrimSpace(event) == "" {
		return nil, errors.New("missing X-Gitlab-Event Header")
	}

	gitLabEvent := gitlab.EventType(event)
	if !isEventSubscribed(gitLabEvent, hook.EventsToAccept) {
		return nil, errors.New("event not defined to be parsed")
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, errors.New("error reading request body")
	}

	return gitlab.ParseWebhook(gitLabEvent, payload)
}

func isEventSubscribed(event gitlab.EventType, events []gitlab.EventType) bool {
	for _, e := range events {
		if event == e {
			return true
		}
	}
	return false
}
