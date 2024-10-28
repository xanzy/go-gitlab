//
// Copyright 2021, Eric Stevens
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
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListGroupHooks(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/hooks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
[
	{
		"id": 1,
		"url": "http://example.com/hook",
		"group_id": 3,
		"push_events": true,
		"push_events_branch_filter": "main",
		"issues_events": true,
		"confidential_issues_events": true,
		"merge_requests_events": true,
		"tag_push_events": true,
		"note_events": true,
		"job_events": true,
		"pipeline_events": true,
		"wiki_page_events": true,
		"deployment_events": true,
		"releases_events": true,
		"subgroup_events": true,
		"member_events": true,
		"enable_ssl_verification": true,
		"alert_status": "executable",
		"created_at": "2012-10-12T17:04:47Z",
		"resource_access_token_events": true,
		"custom_headers": [
			{"key": "Authorization"},
			{"key": "OtherHeader"}
		]
	}
]`)
	})

	groupHooks, _, err := client.Groups.ListGroupHooks(1, nil)
	if err != nil {
		t.Error(err)
	}

	datePointer := time.Date(2012, 10, 12, 17, 4, 47, 0, time.UTC)
	want := []*GroupHook{{
		ID:                        1,
		URL:                       "http://example.com/hook",
		GroupID:                   3,
		PushEvents:                true,
		PushEventsBranchFilter:    "main",
		IssuesEvents:              true,
		ConfidentialIssuesEvents:  true,
		MergeRequestsEvents:       true,
		TagPushEvents:             true,
		NoteEvents:                true,
		JobEvents:                 true,
		PipelineEvents:            true,
		WikiPageEvents:            true,
		DeploymentEvents:          true,
		ReleasesEvents:            true,
		SubGroupEvents:            true,
		MemberEvents:              true,
		EnableSSLVerification:     true,
		AlertStatus:               "executable",
		CreatedAt:                 &datePointer,
		ResourceAccessTokenEvents: true,
		CustomHeaders: []*HookCustomHeader{
			{
				Key: "Authorization",
			},
			{
				Key: "OtherHeader",
			},
		},
	}}

	if !reflect.DeepEqual(groupHooks, want) {
		t.Errorf("listGroupHooks returned \ngot:\n%v\nwant:\n%v", Stringify(groupHooks), Stringify(want))
	}
}

func TestGetGroupHook(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
{
	"id": 1,
	"url": "http://example.com/hook",
	"group_id": 3,
	"push_events": true,
	"push_events_branch_filter": "main",
	"issues_events": true,
	"confidential_issues_events": true,
	"merge_requests_events": true,
	"tag_push_events": true,
	"note_events": true,
	"job_events": true,
	"pipeline_events": true,
	"wiki_page_events": true,
	"deployment_events": true,
	"releases_events": true,
	"subgroup_events": true,
	"member_events": true,
	"enable_ssl_verification": true,
	"alert_status": "executable",
	"created_at": "2012-10-12T17:04:47Z",
	"resource_access_token_events": true,
	"custom_headers": [
		{"key": "Authorization"},
		{"key": "OtherHeader"}
	]
}`)
	})

	groupHook, _, err := client.Groups.GetGroupHook(1, 1)
	if err != nil {
		t.Error(err)
	}

	datePointer := time.Date(2012, 10, 12, 17, 4, 47, 0, time.UTC)
	want := &GroupHook{
		ID:                        1,
		URL:                       "http://example.com/hook",
		GroupID:                   3,
		PushEvents:                true,
		PushEventsBranchFilter:    "main",
		IssuesEvents:              true,
		ConfidentialIssuesEvents:  true,
		MergeRequestsEvents:       true,
		TagPushEvents:             true,
		NoteEvents:                true,
		JobEvents:                 true,
		PipelineEvents:            true,
		WikiPageEvents:            true,
		DeploymentEvents:          true,
		ReleasesEvents:            true,
		SubGroupEvents:            true,
		MemberEvents:              true,
		EnableSSLVerification:     true,
		AlertStatus:               "executable",
		CreatedAt:                 &datePointer,
		ResourceAccessTokenEvents: true,
		CustomHeaders: []*HookCustomHeader{
			{
				Key: "Authorization",
			},
			{
				Key: "OtherHeader",
			},
		},
	}

	if !reflect.DeepEqual(groupHook, want) {
		t.Errorf("getGroupHooks returned \ngot:\n%v\nwant:\n%v", Stringify(groupHook), Stringify(want))
	}
}

func TestAddGroupHook(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/hooks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `
{
	"id": 1,
	"url": "http://example.com/hook",
	"group_id": 3,
	"push_events": true,
	"push_events_branch_filter": "main",
	"issues_events": true,
	"confidential_issues_events": true,
	"merge_requests_events": true,
	"tag_push_events": true,
	"note_events": true,
	"job_events": true,
	"pipeline_events": true,
	"wiki_page_events": true,
	"deployment_events": true,
	"releases_events": true,
	"subgroup_events": true,
	"member_events": true,
	"enable_ssl_verification": true,
	"created_at": "2012-10-12T17:04:47Z",
	"custom_webhook_template": "addTestValue",
	"resource_access_token_events": true,
	"custom_headers": [
		{"key": "Authorization", "value": "testMe"},
		{"key": "OtherHeader", "value": "otherTest"}
	]
}`)
	})

	url := "http://www.example.com/hook"
	opt := &AddGroupHookOptions{
		URL: &url,
	}

	groupHooks, _, err := client.Groups.AddGroupHook(1, opt)
	if err != nil {
		t.Error(err)
	}

	datePointer := time.Date(2012, 10, 12, 17, 4, 47, 0, time.UTC)
	want := &GroupHook{
		ID:                        1,
		URL:                       "http://example.com/hook",
		GroupID:                   3,
		PushEvents:                true,
		PushEventsBranchFilter:    "main",
		IssuesEvents:              true,
		ConfidentialIssuesEvents:  true,
		ConfidentialNoteEvents:    false,
		MergeRequestsEvents:       true,
		TagPushEvents:             true,
		NoteEvents:                true,
		JobEvents:                 true,
		PipelineEvents:            true,
		WikiPageEvents:            true,
		DeploymentEvents:          true,
		ReleasesEvents:            true,
		SubGroupEvents:            true,
		MemberEvents:              true,
		EnableSSLVerification:     true,
		CreatedAt:                 &datePointer,
		CustomWebhookTemplate:     "addTestValue",
		ResourceAccessTokenEvents: true,
		CustomHeaders: []*HookCustomHeader{
			{
				Key:   "Authorization",
				Value: "testMe",
			},
			{
				Key:   "OtherHeader",
				Value: "otherTest",
			},
		},
	}

	if !reflect.DeepEqual(groupHooks, want) {
		t.Errorf("AddGroupHook returned \ngot:\n%v\nwant:\n%v", Stringify(groupHooks), Stringify(want))
	}
}

func TestEditGroupHook(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `
{
	"id": 1,
	"url": "http://example.com/hook",
	"group_id": 3,
	"push_events": true,
	"push_events_branch_filter": "main",
	"issues_events": true,
	"confidential_issues_events": true,
	"merge_requests_events": true,
	"tag_push_events": true,
	"note_events": true,
	"job_events": true,
	"pipeline_events": true,
	"wiki_page_events": true,
	"deployment_events": true,
	"releases_events": true,
	"subgroup_events": true,
	"member_events": true,
	"enable_ssl_verification": true,
	"created_at": "2012-10-12T17:04:47Z",
	"custom_webhook_template": "testValue",
	"resource_access_token_events": true,
	"custom_headers": [
		{"key": "Authorization", "value": "testMe"},
		{"key": "OtherHeader", "value": "otherTest"}
	]
}`)
	})

	url := "http://www.example.com/hook"
	opt := &EditGroupHookOptions{
		URL: &url,
	}

	groupHooks, _, err := client.Groups.EditGroupHook(1, 1, opt)
	if err != nil {
		t.Error(err)
	}

	datePointer := time.Date(2012, 10, 12, 17, 4, 47, 0, time.UTC)
	want := &GroupHook{
		ID:                        1,
		URL:                       "http://example.com/hook",
		GroupID:                   3,
		PushEvents:                true,
		PushEventsBranchFilter:    "main",
		IssuesEvents:              true,
		ConfidentialIssuesEvents:  true,
		ConfidentialNoteEvents:    false,
		MergeRequestsEvents:       true,
		TagPushEvents:             true,
		NoteEvents:                true,
		JobEvents:                 true,
		PipelineEvents:            true,
		WikiPageEvents:            true,
		DeploymentEvents:          true,
		ReleasesEvents:            true,
		SubGroupEvents:            true,
		MemberEvents:              true,
		EnableSSLVerification:     true,
		CreatedAt:                 &datePointer,
		CustomWebhookTemplate:     "testValue",
		ResourceAccessTokenEvents: true,
		CustomHeaders: []*HookCustomHeader{
			{
				Key:   "Authorization",
				Value: "testMe",
			},
			{
				Key:   "OtherHeader",
				Value: "otherTest",
			},
		},
	}

	if !reflect.DeepEqual(groupHooks, want) {
		t.Errorf("EditGroupHook returned \ngot:\n%v\nwant:\n%v", Stringify(groupHooks), Stringify(want))
	}
}

func TestTriggerTestGroupHook(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/hooks/1/test/push_events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"message":"201 Created"}`)
	})

	mux.HandleFunc("/api/v4/groups/1/hooks/1/test/invalid_trigger", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error": "trigger does not have a valid value"}`)
	})

	tests := []struct {
		name       string
		groupID    interface{}
		hookID     int
		trigger    GroupHookTrigger
		wantErr    bool
		wantStatus int
		wantErrMsg string
	}{
		{
			name:       "Valid trigger",
			groupID:    1,
			hookID:     1,
			trigger:    GroupHookTriggerPush,
			wantErr:    false,
			wantStatus: http.StatusCreated,
		},
		{
			name:       "Invalid group ID",
			groupID:    "invalid",
			hookID:     1,
			trigger:    GroupHookTriggerPush,
			wantErr:    true,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Invalid trigger type",
			groupID:    1,
			hookID:     1,
			trigger:    "invalid_trigger",
			wantErr:    true,
			wantStatus: http.StatusBadRequest,
			wantErrMsg: "trigger does not have a valid value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.Groups.TriggerTestGroupHook(tt.groupID, tt.hookID, tt.trigger)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.wantStatus != 0 {
					assert.Equal(t, tt.wantStatus, resp.StatusCode)
				}
				if tt.wantErrMsg != "" {
					assert.Contains(t, err.Error(), tt.wantErrMsg)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.wantStatus, resp.StatusCode)
			}
		})
	}
}

func TestDeleteGroupHook(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Groups.DeleteGroupHook(1, 1)
	if err != nil {
		t.Error(err)
	}
}

func TestGetGroupWebhookHeader(t *testing.T) {
	mux, client := setup(t)

	// Removed most of the arguments to keep test slim
	mux.HandleFunc("/api/v4/groups/1/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": 1,
			"custom_webhook_template": "{\"event\":\"{{object_kind}}\"}",
			"custom_headers": [
			  {
				"key": "Authorization"
			  },
			  {
				"key": "OtherKey"
			  }
			]
		  }`)
	})

	hook, _, err := client.Groups.GetGroupHook(1, 1)
	if err != nil {
		t.Errorf("Projects.GetGroupHook returned error: %v", err)
	}

	want := &GroupHook{
		ID:                    1,
		CustomWebhookTemplate: "{\"event\":\"{{object_kind}}\"}",
		CustomHeaders: []*HookCustomHeader{
			{
				Key: "Authorization",
			},
			{
				Key: "OtherKey",
			},
		},
	}

	if !reflect.DeepEqual(want, hook) {
		t.Errorf("Projects.GetGroupHook returned %+v, want %+v", hook, want)
	}
}

func TestSetGroupWebhookHeader(t *testing.T) {
	mux, client := setup(t)
	var bodyJson map[string]interface{}

	// Removed most of the arguments to keep test slim
	mux.HandleFunc("/api/v4/groups/1/hooks/1/custom_headers/Authorization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		w.WriteHeader(http.StatusNoContent)

		// validate that the `value` body is sent properly
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Unable to read body properly. Error: %v", err)
		}

		// Unmarshal the body into JSON so we can check it
		_ = json.Unmarshal(body, &bodyJson)

		fmt.Fprint(w, ``)
	})

	req, err := client.Groups.SetGroupCustomHeader(1, 1, "Authorization", &SetHookCustomHeaderOptions{Value: Ptr("testValue")})
	if err != nil {
		t.Errorf("Groups.SetGroupCustomHeader returned error: %v", err)
	}

	assert.Equal(t, bodyJson["value"], "testValue")
	assert.Equal(t, http.StatusNoContent, req.StatusCode)
}
