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

func TestListRegistryRepositories(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/5/registry/repositories", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			  "id": 1,
			  "name": "",
			  "path": "group/project",
			  "project_id": 9,
			  "location": "gitlab.example.com:5000/group/project",
			  "created_at": "2019-01-10T13:38:57.391Z",
			  "cleanup_policy_started_at": "2020-01-10T15:40:57.391Z"
			},
			{
			  "id": 2,
			  "name": "releases",
			  "path": "group/project/releases",
			  "project_id": 9,
			  "location": "gitlab.example.com:5000/group/project/releases",
			  "created_at": "2019-01-10T13:39:08.229Z",
			  "cleanup_policy_started_at": "2020-08-17T03:12:35.489Z"
			}
		  ]`)
	})

	repositories, _, err := client.ContainerRegistry.ListRegistryRepositories(5, &ListRegistryRepositoriesOptions{})
	if err != nil {
		t.Errorf("ContainerRegistry.ListRegistryRepositories returned error: %v", err)
	}

	timeLayout := "2006-01-02T15:04:05.000Z"

	created_at1, err := time.Parse(timeLayout, "2019-01-10T13:38:57.391Z")
	if err != nil {
		t.Errorf("ContainerRepository.ListRegistryRepositories error while parsing time: %v", err)
	}

	created_at2, err := time.Parse(timeLayout, "2019-01-10T13:39:08.229Z")
	if err != nil {
		t.Errorf("ContainerRepository.ListRegistryRepositories error while parsing time: %v", err)
	}

	cleanup_policy_started_at1, err := time.Parse(timeLayout, "2020-01-10T15:40:57.391Z")
	if err != nil {
		t.Errorf("ContainerRepository.ListRegistryRepositories error while parsing time: %v", err)
	}

	cleanup_policy_started_at2, err := time.Parse(timeLayout, "2020-08-17T03:12:35.489Z")
	if err != nil {
		t.Errorf("ContainerRepository.ListRegistryRepositories error while parsing time: %v", err)
	}

	want := []*RegistryRepository{
		{ID: 1, Name: "", Path: "group/project", Location: "gitlab.example.com:5000/group/project", CreatedAt: &created_at1, CleanupPolicyStartedAt: &cleanup_policy_started_at1},
		{ID: 2, Name: "releases", Path: "group/project/releases", Location: "gitlab.example.com:5000/group/project/releases", CreatedAt: &created_at2, CleanupPolicyStartedAt: &cleanup_policy_started_at2},
	}
	if !reflect.DeepEqual(want, repositories) {
		t.Errorf("ContainerRepository.ListRegistryRepositories returned %+v, want %+v", repositories, want)
	}
}

func TestDeleteRegistryRepository(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/5/registry/repositories/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.ContainerRegistry.DeleteRegistryRepository(5, 2)
	if err != nil {
		t.Errorf("ContainerRegistry.DeleteRegistryRepository returned error: %v", err)
	}
}

func TestListRegistryRepositoryTags(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/5/registry/repositories/2/tags", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			  "name": "A",
			  "path": "group/project:A",
			  "location": "gitlab.example.com:5000/group/project:A"
			},
			{
			  "name": "latest",
			  "path": "group/project:latest",
			  "location": "gitlab.example.com:5000/group/project:latest"
			}
		  ]`)
	})

	registryRepositoryTags, _, err := client.ContainerRegistry.ListRegistryRepositoryTags(5, 2, &ListRegistryRepositoryTagsOptions{})
	if err != nil {
		t.Errorf("ContainerRegistry.ListRegistryRepositoryTags returned error: %v", err)
	}

	want := []*RegistryRepositoryTag{
		{Name: "A", Path: "group/project:A", Location: "gitlab.example.com:5000/group/project:A"},
		{Name: "latest", Path: "group/project:latest", Location: "gitlab.example.com:5000/group/project:latest"},
	}
	if !reflect.DeepEqual(want, registryRepositoryTags) {
		t.Errorf("ContainerRepository.ListRegistryRepositoryTags returned %+v, want %+v", registryRepositoryTags, want)
	}
}

func TestGetRegistryRepositoryTagDetail(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/5/registry/repositories/2/tags/v10.0.0", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
			"name": "v10.0.0",
			"path": "group/project:latest",
			"location": "gitlab.example.com:5000/group/project:latest",
			"revision": "e9ed9d87c881d8c2fd3a31b41904d01ba0b836e7fd15240d774d811a1c248181",
			"short_revision": "e9ed9d87c",
			"digest": "sha256:c3490dcf10ffb6530c1303522a1405dfaf7daecd8f38d3e6a1ba19ea1f8a1751",
			"created_at": "2019-01-06T16:49:51.272+00:00",
			"total_size": 350224384
		  }`)
	})

	repositoryTag, _, err := client.ContainerRegistry.GetRegistryRepositoryTagDetail(5, 2, "v10.0.0")
	if err != nil {
		t.Errorf("ContainerRegistry.GetRegistryRepositoryTagDetail returned error: %v", err)
	}

	timeLayout := "2006-01-02T15:04:05Z07:00"
	createdAt, err := time.Parse(timeLayout, "2019-01-06T16:49:51.272+00:00")
	if err != nil {
		t.Errorf("ContainerRepository.ListRegistryRepositories error while parsing time: %v", err)
	}

	want := &RegistryRepositoryTag{
		Name:          "v10.0.0",
		Path:          "group/project:latest",
		Location:      "gitlab.example.com:5000/group/project:latest",
		Revision:      "e9ed9d87c881d8c2fd3a31b41904d01ba0b836e7fd15240d774d811a1c248181",
		ShortRevision: "e9ed9d87c",
		Digest:        "sha256:c3490dcf10ffb6530c1303522a1405dfaf7daecd8f38d3e6a1ba19ea1f8a1751",
		CreatedAt:     &createdAt,
		TotalSize:     350224384,
	}
	if !reflect.DeepEqual(want, repositoryTag) {
		t.Errorf("ContainerRepository.ListRegistryRepositories returned %+v, want %+v", repositoryTag, want)
	}
}
