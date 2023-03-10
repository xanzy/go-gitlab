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
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestListAllDeployTokens(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/deploy_tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
[
	{
		"id": 1,
		"name": "MyToken",
		"username": "gitlab+deploy-token-1",
		"expires_at": "2020-02-14T00:00:00.000Z",
		"revoked": true,
		"expired": true,
		"scopes": [
			"read_repository",
			"read_registry"
		]
	}
]
`)
	})

	deployTokens, _, err := client.DeployTokens.ListAllDeployTokens()
	if err != nil {
		t.Errorf("DeployTokens.ListAllDeployTokens returned an error: %v", err)
	}

	wantExpiresAt := time.Date(2020, 0o2, 14, 0, 0, 0, 0, time.UTC)

	want := []*DeployToken{
		{
			ID:        1,
			Name:      "MyToken",
			Username:  "gitlab+deploy-token-1",
			ExpiresAt: &wantExpiresAt,
			Revoked:   true,
			Expired:   true,
			Scopes: []string{
				"read_repository",
				"read_registry",
			},
		},
	}

	if !reflect.DeepEqual(want, deployTokens) {
		t.Errorf("DeployTokens.ListAllDeployTokens returned %+v, want %+v", deployTokens, want)
	}
}

func TestListProjectDeployTokens(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/deploy_tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
[
  {
    "id": 1,
    "name": "MyToken",
    "username": "gitlab+deploy-token-1",
    "expires_at": "2020-02-14T00:00:00.000Z",
    "scopes": [
      "read_repository",
      "read_registry"
    ]
  }
]
`)
	})

	deployTokens, _, err := client.DeployTokens.ListProjectDeployTokens(1, nil)
	if err != nil {
		t.Errorf("DeployTokens.ListProjectDeployTokens returned an error: %v", err)
	}

	wantExpiresAt := time.Date(2020, 0o2, 14, 0, 0, 0, 0, time.UTC)

	want := []*DeployToken{
		{
			ID:        1,
			Name:      "MyToken",
			Username:  "gitlab+deploy-token-1",
			ExpiresAt: &wantExpiresAt,
			Scopes: []string{
				"read_repository",
				"read_registry",
			},
		},
	}

	if !reflect.DeepEqual(want, deployTokens) {
		t.Errorf("DeployTokens.ListProjectDeployTokens returned %+v, want %+v", deployTokens, want)
	}
}

func TestGetProjectDeployTokens(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/deploy_tokens/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
{
  "id": 1,
  "name": "MyToken",
  "username": "gitlab+deploy-token-1",
  "expires_at": "2020-02-14T00:00:00.000Z",
  "scopes": [
    "read_repository",
    "read_registry"
  ]
}
`)
	})

	deployToken, _, err := client.DeployTokens.GetProjectDeployToken(1, 1)
	if err != nil {
		t.Errorf("DeployTokens.GetProjectDeployToken returned an error: %v", err)
	}

	wantExpiresAt := time.Date(2020, 0o2, 14, 0, 0, 0, 0, time.UTC)

	want := &DeployToken{
		ID:        1,
		Name:      "MyToken",
		Username:  "gitlab+deploy-token-1",
		ExpiresAt: &wantExpiresAt,
		Scopes: []string{
			"read_repository",
			"read_registry",
		},
	}

	if !reflect.DeepEqual(want, deployToken) {
		t.Errorf("DeployTokens.GetProjectDeployToken returned %+v, want %+v", deployToken, want)
	}
}

func TestCreateProjectDeployToken(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/deploy_tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `
{
	"id": 1,
	"name": "My deploy token",
	"username": "custom-user",
	"expires_at": "2021-01-01T00:00:00.000Z",
	"token": "jMRvtPNxrn3crTAGukpZ",
	"scopes": [
		"read_repository"
	]
}
`)
	})

	expiresAt := time.Date(2021, 0o1, 0o1, 0, 0, 0, 0, time.UTC)

	deployToken, _, err := client.DeployTokens.CreateProjectDeployToken(5, &CreateProjectDeployTokenOptions{
		Name:      String("My deploy token"),
		Username:  String("custom-user"),
		ExpiresAt: &expiresAt,
		Scopes: &[]string{
			"read_repository",
		},
	})
	if err != nil {
		t.Errorf("DeployTokens.CreateProjectDeployToken returned an error: %v", err)
	}

	want := &DeployToken{
		ID:        1,
		Name:      "My deploy token",
		Username:  "custom-user",
		ExpiresAt: &expiresAt,
		Token:     "jMRvtPNxrn3crTAGukpZ",
		Scopes: []string{
			"read_repository",
		},
	}

	if !reflect.DeepEqual(want, deployToken) {
		t.Errorf("DeployTokens.CreateProjectDeployToken returned %+v, want %+v", deployToken, want)
	}
}

func TestDeleteProjectDeployToken(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/deploy_tokens/13", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusAccepted)
	})

	resp, err := client.DeployTokens.DeleteProjectDeployToken(5, 13)
	if err != nil {
		t.Errorf("DeployTokens.DeleteProjectDeployToken returned an error: %v", err)
	}

	want := http.StatusAccepted
	got := resp.StatusCode

	if want != got {
		t.Errorf("DeployTokens.DeleteProjectDeployToken returned %+v, want %+v", got, want)
	}
}

func TestListGroupDeployTokens(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/deploy_tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
[
  {
    "id": 1,
    "name": "MyToken",
    "username": "gitlab+deploy-token-1",
    "expires_at": "2020-02-14T00:00:00.000Z",
    "scopes": [
      "read_repository",
      "read_registry"
    ]
  }
]
`)
	})

	deployTokens, _, err := client.DeployTokens.ListGroupDeployTokens(1, nil)
	if err != nil {
		t.Errorf("DeployTokens.ListGroupDeployTokens returned an error: %v", err)
	}

	wantExpiresAt := time.Date(2020, 0o2, 14, 0, 0, 0, 0, time.UTC)

	want := []*DeployToken{
		{
			ID:        1,
			Name:      "MyToken",
			Username:  "gitlab+deploy-token-1",
			ExpiresAt: &wantExpiresAt,
			Scopes: []string{
				"read_repository",
				"read_registry",
			},
		},
	}

	if !reflect.DeepEqual(want, deployTokens) {
		t.Errorf("DeployTokens.ListGroupDeployTokens returned %+v, want %+v", deployTokens, want)
	}
}

func TestGetGroupDeployTokens(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/deploy_tokens/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
{
  "id": 1,
  "name": "MyToken",
  "username": "gitlab+deploy-token-1",
  "expires_at": "2020-02-14T00:00:00.000Z",
  "scopes": [
    "read_repository",
    "read_registry"
  ]
}
`)
	})

	deployToken, _, err := client.DeployTokens.GetGroupDeployToken(1, 1)
	if err != nil {
		t.Errorf("DeployTokens.GetGroupDeployToken returned an error: %v", err)
	}

	wantExpiresAt := time.Date(2020, 0o2, 14, 0, 0, 0, 0, time.UTC)

	want := &DeployToken{
		ID:        1,
		Name:      "MyToken",
		Username:  "gitlab+deploy-token-1",
		ExpiresAt: &wantExpiresAt,
		Scopes: []string{
			"read_repository",
			"read_registry",
		},
	}

	if !reflect.DeepEqual(want, deployToken) {
		t.Errorf("DeployTokens.GetGroupDeployToken returned %+v, want %+v", deployToken, want)
	}
}

func TestCreateGroupDeployToken(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/deploy_tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `
{
	"id": 1,
	"name": "My deploy token",
	"username": "custom-user",
	"expires_at": "2021-01-01T00:00:00.000Z",
	"token": "jMRvtPNxrn3crTAGukpZ",
	"scopes": [
		"read_repository"
	]
}
`)
	})

	expiresAt := time.Date(2021, 0o1, 0o1, 0, 0, 0, 0, time.UTC)

	deployToken, _, err := client.DeployTokens.CreateGroupDeployToken(5, &CreateGroupDeployTokenOptions{
		Name:      String("My deploy token"),
		Username:  String("custom-user"),
		ExpiresAt: &expiresAt,
		Scopes: &[]string{
			"read_repository",
		},
	})
	if err != nil {
		t.Errorf("DeployTokens.CreateGroupDeployToken returned an error: %v", err)
	}

	want := &DeployToken{
		ID:        1,
		Name:      "My deploy token",
		Username:  "custom-user",
		ExpiresAt: &expiresAt,
		Token:     "jMRvtPNxrn3crTAGukpZ",
		Scopes: []string{
			"read_repository",
		},
	}

	if !reflect.DeepEqual(want, deployToken) {
		t.Errorf("DeployTokens.CreateGroupDeployToken returned %+v, want %+v", deployToken, want)
	}
}

func TestDeleteGroupDeployToken(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/deploy_tokens/13", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusAccepted)
	})

	resp, err := client.DeployTokens.DeleteGroupDeployToken(5, 13)
	if err != nil {
		t.Errorf("DeployTokens.DeleteGroupDeployToken returned an error: %v", err)
	}

	want := http.StatusAccepted
	got := resp.StatusCode

	if want != got {
		t.Errorf("DeployTokens.DeleteGroupDeployToken returned %+v, want %+v", got, want)
	}
}
