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
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestGetGlobalSettings(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/notification_settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
			"level": "participating",
			"notification_email": "admin@example.com"
		  }`)
	})

	settings, _, err := client.NotificationSettings.GetGlobalSettings()
	if err != nil {
		t.Errorf("NotifcationSettings.GetGlobalSettings returned error: %v", err)
	}

	want := &NotificationSettings{
		Level:             1,
		NotificationEmail: "admin@example.com",
	}
	if !reflect.DeepEqual(settings, want) {
		t.Errorf("NotificationSettings.GetGlobalSettings returned %+v, want %+v", settings, want)
	}
}

func TestGetProjectSettings(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/notification_settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
		"level":"custom",
		"events":{
			"new_note":true,
			"new_issue":true,
			"reopen_issue":true,
			"close_issue":true,
			"reassign_issue":true,
			"issue_due":true,
			"new_merge_request":true,
			"push_to_merge_request":true,
			"reopen_merge_request":true,
			"close_merge_request":true,
			"reassign_merge_request":true,
			"merge_merge_request":true,
			"failed_pipeline":true,
			"fixed_pipeline":true,
			"success_pipeline":true,
			"moved_project":true,
			"merge_when_pipeline_succeeds":true,
			"new_epic":true
			}
		}`)
	})

	settings, _, err := client.NotificationSettings.GetSettingsForProject(1)
	if err != nil {
		t.Errorf("NotifcationSettings.GetSettingsForProject returned error: %v", err)
	}

	want := &NotificationSettings{
		Level: 5, //custom
		Events: &NotificationEvents{
			NewEpic:                   true,
			NewNote:                   true,
			NewIssue:                  true,
			ReopenIssue:               true,
			CloseIssue:                true,
			ReassignIssue:             true,
			IssueDue:                  true,
			NewMergeRequest:           true,
			PushToMergeRequest:        true,
			ReopenMergeRequest:        true,
			CloseMergeRequest:         true,
			ReassignMergeRequest:      true,
			MergeMergeRequest:         true,
			FailedPipeline:            true,
			FixedPipeline:             true,
			SuccessPipeline:           true,
			MovedProject:              true,
			MergeWhenPipelineSucceeds: true,
		},
	}
	if !reflect.DeepEqual(settings, want) {
		t.Errorf("NotificationSettings.GetSettingsForProject returned %+v, want %+v", settings, want)
	}
}

func TestUpdateProjectSettings(t *testing.T) {
	mux, client := setup(t)

	// Create the request to send
	var reqBody NotificationSettingsOptions
	customLevel := notificationLevelTypes["custom"]
	options := NotificationSettingsOptions{
		Level:        &customLevel,
		NewEpic:      Bool(true),
		MovedProject: Bool(true),
		CloseIssue:   Bool(true),
	}

	// Handle the request on the server, and return a fully hydrated response
	mux.HandleFunc("/api/v4/projects/1/notification_settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)

		// Store the body for later, so we can check only some values are marshalled properly for update
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &reqBody)

		fmt.Fprintf(w, `{
		"level":"custom",
		"events":{
			"new_note":true,
			"new_issue":true,
			"reopen_issue":true,
			"close_issue":true,
			"reassign_issue":true,
			"issue_due":true,
			"new_merge_request":true,
			"push_to_merge_request":true,
			"reopen_merge_request":true,
			"close_merge_request":true,
			"reassign_merge_request":true,
			"merge_merge_request":true,
			"failed_pipeline":true,
			"fixed_pipeline":true,
			"success_pipeline":true,
			"moved_project":true,
			"merge_when_pipeline_succeeds":true,
			"new_epic":true
			}
		}`)
	})

	// Make the actual request
	settings, _, err := client.NotificationSettings.UpdateSettingsForProject(1, &options)
	if err != nil {
		t.Errorf("NotifcationSettings.UpdateSettingsForProject returned error: %v", err)
	}

	// Test the response and the request
	wantResponse := &NotificationSettings{
		Level: customLevel,
		Events: &NotificationEvents{
			NewEpic:                   true,
			NewNote:                   true,
			NewIssue:                  true,
			ReopenIssue:               true,
			CloseIssue:                true,
			ReassignIssue:             true,
			IssueDue:                  true,
			NewMergeRequest:           true,
			PushToMergeRequest:        true,
			ReopenMergeRequest:        true,
			CloseMergeRequest:         true,
			ReassignMergeRequest:      true,
			MergeMergeRequest:         true,
			FailedPipeline:            true,
			FixedPipeline:             true,
			SuccessPipeline:           true,
			MovedProject:              true,
			MergeWhenPipelineSucceeds: true,
		},
	}

	if !reflect.DeepEqual(settings, wantResponse) {
		t.Errorf("NotificationSettings.UpdateSettingsForProject returned for the response %+v, want %+v", settings, wantResponse)
	}
	if !reflect.DeepEqual(settings, wantResponse) {
		t.Errorf("NotificationSettings.UpdateSettingsForProject send for the request %+v, want %+v", reqBody, options)
	}
}
