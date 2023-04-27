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
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListEnvironments(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/environments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/projects/1/environments?name=review%2Ffix-foo&page=1&per_page=10")
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "review/fix-foo",
				"slug": "review-fix-foo-dfjre3",
				"external_url": "https://review-fix-foo-dfjre3.example.gitlab.com",
				"state": "stopped",
				"created_at": "2013-10-02T10:12:29Z",
				"updated_at": "2013-12-02T10:12:29Z"
			}
		]`)
	})

	envs, _, err := client.Environments.ListEnvironments(1, &ListEnvironmentsOptions{Name: String("review/fix-foo"), ListOptions: ListOptions{Page: 1, PerPage: 10}})
	if err != nil {
		log.Fatal(err)
	}

	createdAtWant, _ := time.Parse(timeLayout, "2013-10-02T10:12:29Z")
	updatedAtWant, _ := time.Parse(timeLayout, "2013-12-02T10:12:29Z")
	want := []*Environment{{
		ID:          1,
		Name:        "review/fix-foo",
		Slug:        "review-fix-foo-dfjre3",
		ExternalURL: "https://review-fix-foo-dfjre3.example.gitlab.com",
		State:       "stopped",
		CreatedAt:   &createdAtWant,
		UpdatedAt:   &updatedAtWant,
	}}
	if !reflect.DeepEqual(want, envs) {
		t.Errorf("Environments.ListEnvironments returned %+v, want %+v", envs, want)
	}
}

func TestGetEnvironment(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/environments/5949167", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "review/fix-foo",
			"slug": "review-fix-foo-dfjre3",
			"external_url": "https://review-fix-foo-dfjre3.example.gitlab.com",
			"state": "stopped",
			"created_at": "2013-10-02T10:12:29Z",
			"updated_at": "2013-12-02T10:12:29Z"
		}`)
	})

	env, _, err := client.Environments.GetEnvironment(1, 5949167)
	if err != nil {
		t.Errorf("Environemtns.GetEnvironment returned error: %v", err)
	}

	createdAtWant, _ := time.Parse(timeLayout, "2013-10-02T10:12:29Z")
	updatedAtWant, _ := time.Parse(timeLayout, "2013-12-02T10:12:29Z")
	want := &Environment{
		ID:          1,
		Name:        "review/fix-foo",
		Slug:        "review-fix-foo-dfjre3",
		ExternalURL: "https://review-fix-foo-dfjre3.example.gitlab.com",
		State:       "stopped",
		CreatedAt:   &createdAtWant,
		UpdatedAt:   &updatedAtWant,
	}
	if !reflect.DeepEqual(want, env) {
		t.Errorf("Environments.GetEnvironment returned %+v, want %+v", env, want)
	}
}

func TestCreateEnvironment(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/environments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, "/api/v4/projects/1/environments")
		fmt.Fprint(w, `{
      "id": 1,
      "name": "deploy",
      "slug": "deploy",
      "external_url": "https://deploy.example.gitlab.com",
      "tier": "production"
    }`)
	})

	envs, _, err := client.Environments.CreateEnvironment(1, &CreateEnvironmentOptions{Name: String("deploy"), ExternalURL: String("https://deploy.example.gitlab.com"), Tier: String("production")})
	if err != nil {
		log.Fatal(err)
	}

	want := &Environment{ID: 1, Name: "deploy", Slug: "deploy", ExternalURL: "https://deploy.example.gitlab.com", Tier: "production"}
	if !reflect.DeepEqual(want, envs) {
		t.Errorf("Environments.CreateEnvironment returned %+v, want %+v", envs, want)
	}
}

func TestEditEnvironment(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/environments/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testURL(t, r, "/api/v4/projects/1/environments/1")
		fmt.Fprint(w, `{
      "id": 1,
      "name": "staging",
      "slug": "staging",
      "external_url": "https://staging.example.gitlab.com",
      "tier": "staging"
    }`)
	})

	envs, _, err := client.Environments.EditEnvironment(1, 1, &EditEnvironmentOptions{Name: String("staging"), ExternalURL: String("https://staging.example.gitlab.com"), Tier: String("staging")})
	if err != nil {
		log.Fatal(err)
	}

	want := &Environment{ID: 1, Name: "staging", Slug: "staging", ExternalURL: "https://staging.example.gitlab.com", Tier: "staging"}
	if !reflect.DeepEqual(want, envs) {
		t.Errorf("Environments.EditEnvironment returned %+v, want %+v", envs, want)
	}
}

func TestDeleteEnvironment(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/environments/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, "/api/v4/projects/1/environments/1")
	})
	_, err := client.Environments.DeleteEnvironment(1, 1)
	if err != nil {
		log.Fatal(err)
	}
}

func TestStopEnvironment(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/environments/1/stop", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, "/api/v4/projects/1/environments/1/stop")
		fmt.Fprint(w, `{
      "id": 1,
      "name": "staging",
      "state": "stopping",
      "slug": "staging",
      "external_url": "https://staging.example.gitlab.com",
      "tier": "staging"
    }`)
	})
	_, _, err := client.Environments.StopEnvironment(1, 1)
	if err != nil {
		log.Fatal(err)
	}
}

func TestUnmarshal(t *testing.T) {
	jsonObject := `
    {
        "id": 10,
        "name": "production",
        "slug": "production",
        "external_url": "https://example.com",
        "project": {
            "id": 1,
            "description": "",
            "name": "Awesome Project",
            "name_with_namespace": "FooBar Group / Awesome Project",
            "path": "awesome-project",
            "path_with_namespace": "foobar-group/awesome-project",
            "created_at": "2017-09-30T11:10:08.476-04:00",
            "default_branch": "develop",
            "tag_list": [],
            "ssh_url_to_repo": "git@example.gitlab.com:foobar-group/api.git",
            "http_url_to_repo": "https://example.gitlab.com/foobar-group/api.git",
            "web_url": "https://example.gitlab.com/foobar-group/api",
            "readme_url": null,
            "avatar_url": null,
            "star_count": 0,
            "forks_count": 1,
            "last_activity_at": "2019-11-03T22:22:46.564-05:00",
            "namespace": {
                "id": 15,
                "name": "FooBar Group",
                "path": "foobar-group",
                "kind": "group",
                "full_path": "foobar-group",
                "parent_id": null,
                "avatar_url": null,
                "web_url": "https://example.gitlab.com/groups/foobar-group"
            }
        },
        "state": "available"
    }`

	var env Environment
	err := json.Unmarshal([]byte(jsonObject), &env)

	if assert.NoError(t, err) {
		assert.Equal(t, 10, env.ID)
		assert.Equal(t, "production", env.Name)
		assert.Equal(t, "https://example.com", env.ExternalURL)
		assert.Equal(t, "available", env.State)
		if assert.NotNil(t, env.Project) {
			assert.Equal(t, "Awesome Project", env.Project.Name)
		}
	}
}
