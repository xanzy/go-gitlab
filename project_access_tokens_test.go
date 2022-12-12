//
// Copyright 2021, Patrick Webster
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
	"time"
)

func TestListProjectAccessTokens(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/access_tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_project_access_tokens.json")
	})

	projectAccessTokens, _, err := client.ProjectAccessTokens.ListProjectAccessTokens(1, &ListProjectAccessTokensOptions{Page: 1, PerPage: 20})
	if err != nil {
		t.Errorf("ProjectAccessTokens.ListProjectAccessTokens returned error: %v", err)
	}

	time1, err := time.Parse(time.RFC3339, "2021-03-09T21:11:47.271Z")
	if err != nil {
		t.Errorf("ProjectAccessTokens.ListProjectAccessTokens returned error: %v", err)
	}
	time2, err := time.Parse(time.RFC3339, "2021-03-09T21:11:47.340Z")
	if err != nil {
		t.Errorf("ProjectAccessTokens.ListProjectAccessTokens returned error: %v", err)
	}
	time3, err := time.Parse(time.RFC3339, "2021-03-10T21:11:47.271Z")
	if err != nil {
		t.Errorf("GroupAccessTokens.ListGroupAccessTokens returned error: %v", err)
	}

	want := []*ProjectAccessToken{
		{
			ID:          1876,
			UserID:      2453,
			Name:        "token 10",
			Scopes:      []string{"api", "read_api", "read_repository", "write_repository"},
			CreatedAt:   &time1,
			LastUsedAt:  &time3,
			Active:      true,
			Revoked:     false,
			AccessLevel: AccessLevelValue(40),
		},
		{
			ID:          1877,
			UserID:      2456,
			Name:        "token 8",
			Scopes:      []string{"api", "read_api", "read_repository", "write_repository"},
			CreatedAt:   &time2,
			Active:      true,
			Revoked:     false,
			AccessLevel: AccessLevelValue(30),
		},
	}

	if !reflect.DeepEqual(want, projectAccessTokens) {
		t.Errorf("ProjectAccessTokens.ListProjectAccessTokens returned %+v, want %+v", projectAccessTokens, want)
	}
}

func TestGetProjectAccessToken(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/access_tokens/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_project_access_token.json")
	})

	projectAccessToken, _, err := client.ProjectAccessTokens.GetProjectAccessToken(1, 1)
	if err != nil {
		t.Errorf("ProjectAccessTokens.GetProjectAccessToken returned error: %v", err)
	}

	createdAt, err := time.Parse(time.RFC3339, "2021-03-09T21:11:47.271Z")
	if err != nil {
		t.Errorf("ProjectAccessTokens.GetProjectAccessToken returned error: %v", err)
	}

	want := &ProjectAccessToken{
		ID:          1,
		UserID:      2453,
		Name:        "token 10",
		Scopes:      []string{"api", "read_api", "read_repository", "write_repository"},
		CreatedAt:   &createdAt,
		Active:      true,
		Revoked:     false,
		AccessLevel: AccessLevelValue(40),
	}

	if !reflect.DeepEqual(want, projectAccessToken) {
		t.Errorf("ProjectAccessTokens.GetProjectAccessToken returned %+v, want %+v", projectAccessToken, want)
	}
}

func TestCreateProjectAccessToken(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/access_tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		mustWriteHTTPResponse(t, w, "testdata/create_project_access_token.json")
	})

	projectAccessToken, _, err := client.ProjectAccessTokens.CreateProjectAccessToken(1, nil)
	if err != nil {
		t.Errorf("ProjectAccessTokens.CreateProjectAccessToken returned error: %v", err)
	}

	time1, err := time.Parse(time.RFC3339, "2021-03-09T21:11:47.271Z")
	if err != nil {
		t.Errorf("ProjectAccessTokens.CreateProjectAccessToken returned error: %v", err)
	}
	want := &ProjectAccessToken{
		ID:          1876,
		UserID:      2453,
		Name:        "token 10",
		Scopes:      []string{"api", "read_api", "read_repository", "write_repository"},
		ExpiresAt:   nil,
		CreatedAt:   &time1,
		Active:      true,
		Revoked:     false,
		Token:       "2UsevZE1x1ZdFZW4MNzH",
		AccessLevel: AccessLevelValue(40),
	}

	if !reflect.DeepEqual(want, projectAccessToken) {
		t.Errorf("ProjectAccessTokens.CreateProjectAccessToken returned %+v, want %+v", projectAccessToken, want)
	}
}

func TestRevokeProjectAccessToken(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/access_tokens/1234", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.ProjectAccessTokens.RevokeProjectAccessToken("1", 1234)
	if err != nil {
		t.Errorf("ProjectAccessTokens.RevokeProjectAccessToken returned error: %v", err)
	}
}
