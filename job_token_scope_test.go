// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProjectTokenAccessSettings(t *testing.T) {
	mux, client := setup(t)

	// Handle project ID 1, and print a result of access settings
	mux.HandleFunc("/api/v4/projects/1/job_token_scope", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		// Print on the response
		fmt.Fprint(w, `{"inbound_enabled":true,"outbound_enabled":false}`)
	})

	want := &JobTokenAccessSettings{
		InboundEnabled:  true,
		OutboundEnabled: false,
	}

	settings, _, err := client.JobTokenScope.GetProjectJobTokenAccessSettings(1)

	assert.NoError(t, err)
	assert.Equal(t, want, settings)
}

func TestPatchProjectJobTokenAccessSettings(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/job_token_scope", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)

		// Read the request to determine which target project is passed in
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("JobTokenScope.PatchProjectJobTokenAccessSettings failed to read body")
		}

		// Parse to object to ensure it's sent on the request appropriately.
		var options PatchProjectJobTokenAccessSettingsOptions
		err = json.Unmarshal(body, &options)
		if err != nil {
			t.Fatalf("JobTokenScope.PatchProjectJobTokenAccessSettings failed to unmarshal body: %v", err)
		}

		// Ensure we provide the proper response
		w.WriteHeader(http.StatusNoContent)

		// Print an empty body, since that's what the API provides.
		fmt.Fprint(w, "")
	})

	resp, err := client.JobTokenScope.PatchProjectJobTokenAccessSettings(
		1,
		&PatchProjectJobTokenAccessSettingsOptions{
			Enabled: false,
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode)
}

// This tests that when calling the GetProjectJobTokenInboundAllowList, we get a
// list of projects back properly. There isn't a "deep" test with every attribute
// specifieid, because the object returned is a *Project object, which is already
// tested in project.go.
func TestGetProjectJobTokenInboundAllowList(t *testing.T) {
	mux, client := setup(t)

	// Handle project ID 1, and print a result of two projects
	mux.HandleFunc("/api/v4/projects/1/job_token_scope/allowlist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		// Print on the response
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	want := []*Project{{ID: 1}, {ID: 2}}
	projects, _, err := client.JobTokenScope.GetProjectJobTokenInboundAllowList(
		1,
		&GetJobTokenInboundAllowListOptions{},
	)

	assert.NoError(t, err)
	assert.Equal(t, want, projects)
}

func TestAddProjectToJobScopeAllowList(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/job_token_scope/allowlist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		// Read the request to determine which target project is passed in
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("JobTokenScope.AddProjectToJobScopeAllowList failed to read body")
		}

		// Parse to object to ensure it's sent on the request appropriately.
		var createTokenRequest JobTokenInboundAllowOptions
		err = json.Unmarshal(body, &createTokenRequest)
		if err != nil {
			t.Fatalf("JobTokenScope.AddProjectToJobScopeAllowList failed to unmarshal body: %v", err)
		}

		// Ensure we provide the proper response
		w.WriteHeader(http.StatusCreated)

		// Print on the response with the proper target project
		fmt.Fprintf(w, `{
			"source_project_id": 1,
			"target_project_id": %d
		}`, *createTokenRequest.TargetProjectID)
	})

	want := &JobTokenInboundAllowItem{
		SourceProjectID: 1,
		TargetProjectID: 2,
	}

	addTokenResponse, resp, err := client.JobTokenScope.AddProjectToJobScopeAllowList(
		1,
		&JobTokenInboundAllowOptions{TargetProjectID: Int(2)},
	)
	assert.NoError(t, err)
	assert.Equal(t, want, addTokenResponse)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestRemoveProjectFromJobScopeAllowList(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/job_token_scope/allowlist/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)

		// Read the request to determine which target project is passed in
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("JobTokenScope.RemoveProjectFromJobScopeAllowList failed to read body")
		}

		// The body should be empty since all attributes are passed in the path
		if body != nil && string(body) != "" {
			t.Fatalf("JobTokenScope.RemoveProjectFromJobScopeAllowList failed to unmarshal body: %v", err)
		}

		// Ensure we provide the proper response
		w.WriteHeader(http.StatusNoContent)

		// Print an empty body, since that's what the API provides.
		fmt.Fprint(w, "")
	})

	resp, err := client.JobTokenScope.RemoveProjectFromJobScopeAllowList(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode)
}
