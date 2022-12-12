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
)

func TestTopicsService_ListTopics(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/topics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
      {
        "id": 1,
        "name": "gitlab",
        "title": "GitLab",
        "description": "GitLab is a version control system",
        "total_projects_count": 1000,
        "avatar_url": "http://www.gravatar.com/avatar/a0d477b3ea21970ce6ffcbb817b0b435?s=80&d=identicon"
      },
      {
        "id": 3,
        "name": "git",
        "title": "Git",
        "description": "Git is free and open source",
        "total_projects_count": 900,
        "avatar_url": "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon"
      },
      {
        "id": 2,
        "name": "git-lfs",
        "title": "Git LFS",
        "description": null,
        "total_projects_count": 300,
        "avatar_url": null
      }
    ]`)
	})

	opt := &ListTopicsOptions{Search: String("git")}
	topics, _, err := client.Topics.ListTopics(opt)
	if err != nil {
		t.Errorf("Tags.ListTags returned error: %v", err)
	}

	want := []*Topic{{
		ID:                 1,
		Name:               "gitlab",
		Title:              "GitLab",
		Description:        "GitLab is a version control system",
		TotalProjectsCount: 1000,
		AvatarURL:          "http://www.gravatar.com/avatar/a0d477b3ea21970ce6ffcbb817b0b435?s=80&d=identicon",
	}, {
		ID:                 3,
		Name:               "git",
		Title:              "Git",
		Description:        "Git is free and open source",
		TotalProjectsCount: 900,
		AvatarURL:          "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
	}, {
		ID:                 2,
		Name:               "git-lfs",
		Title:              "Git LFS",
		TotalProjectsCount: 300,
	}}
	if !reflect.DeepEqual(want, topics) {
		t.Errorf("Topics.ListTopics returned %+v, want %+v", topics, want)
	}
}

func TestTopicsService_GetTopic(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/topics/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
      "id": 1,
      "name": "gitlab",
      "title": "GitLab",
      "Description": "GitLab is a version control system",
      "total_projects_count": 1000,
      "avatar_url": "http://www.gravatar.com/avatar/a0d477b3ea21970ce6ffcbb817b0b435?s=80&d=identicon"
    }`)
	})

	release, _, err := client.Topics.GetTopic(1)
	if err != nil {
		t.Errorf("Topics.GetTopic returned error: %v", err)
	}

	want := &Topic{
		ID:                 1,
		Name:               "gitlab",
		Title:              "GitLab",
		Description:        "GitLab is a version control system",
		TotalProjectsCount: 1000,
		AvatarURL:          "http://www.gravatar.com/avatar/a0d477b3ea21970ce6ffcbb817b0b435?s=80&d=identicon",
	}
	if !reflect.DeepEqual(want, release) {
		t.Errorf("Topics.GetTopic returned %+v, want %+v", release, want)
	}
}

func TestTopicsService_CreateTopic(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/topics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
      "id": 1,
      "name": "topic1",
      "title": "Topic 1",
      "description": "description",
      "total_projects_count": 0,
      "avatar_url": null
    }`)
	})

	opt := &CreateTopicOptions{Name: String("topic1"), Title: String("Topic 1"), Description: String("description")}
	release, _, err := client.Topics.CreateTopic(opt)
	if err != nil {
		t.Errorf("Topics.CreateTopic returned error: %v", err)
	}

	want := &Topic{ID: 1, Name: "topic1", Title: "Topic 1", Description: "description", TotalProjectsCount: 0}
	if !reflect.DeepEqual(want, release) {
		t.Errorf("Topics.CreateTopic returned %+v, want %+v", release, want)
	}
}

func TestTopicsService_UpdateTopic(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/topics/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{
      "id": 1,
      "name": "topic1",
      "title": "Topic 1",
      "description": "description",
      "total_projects_count": 0,
      "avatar_url": null
    }`)
	})

	opt := &UpdateTopicOptions{Name: String("topic1"), Title: String("Topic 1"), Description: String("description")}
	release, _, err := client.Topics.UpdateTopic(1, opt)
	if err != nil {
		t.Errorf("Topics.UpdateTopic returned error: %v", err)
	}

	want := &Topic{ID: 1, Name: "topic1", Title: "Topic 1", Description: "description", TotalProjectsCount: 0}
	if !reflect.DeepEqual(want, release) {
		t.Errorf("Topics.UpdateTopic returned %+v, want %+v", release, want)
	}
}
