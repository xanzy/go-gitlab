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
	"testing"
)

func TestReleaseLinksService_ListReleaseLinks(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1/assets/links",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, exampleReleaseLinkList)
		})

	releaseLinks, _, err := client.ReleaseLinks.ListReleaseLinks(
		1, exampleTagName, &ListReleaseLinksOptions{},
	)
	if err != nil {
		t.Error(err)
	}
	if len(releaseLinks) != 2 {
		t.Error("expected 2 links")
	}
	if releaseLinks[0].Name != "awesome-v0.2.msi" {
		t.Errorf("release link name, expected '%s', got '%s'", "awesome-v0.2.msi",
			releaseLinks[0].Name)
	}
}

func TestReleaseLinksService_CreateReleaseLink(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1/assets/links",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, exampleReleaseLink)
		})

	releaseLink, _, err := client.ReleaseLinks.CreateReleaseLink(
		1, exampleTagName,
		&CreateReleaseLinkOptions{
			Name: String(exampleReleaseName),
			URL:  String("http://192.168.10.15:3000"),
		})
	if err != nil {
		t.Error(err)
	}
	if releaseLink.Name != exampleReleaseName {
		t.Errorf("release link name, expected '%s', got '%s'", exampleReleaseName,
			releaseLink.Name)
	}
}

func TestReleaseLinksService_GetReleaseLink(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1/assets/links/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, exampleReleaseLink)
		})

	releaseLink, _, err := client.ReleaseLinks.GetReleaseLink(1, exampleTagName, 1)
	if err != nil {
		t.Error(err)
	}
	if releaseLink.Name != exampleReleaseName {
		t.Errorf("release link name, expected '%s', got '%s'", exampleReleaseName,
			releaseLink.Name)
	}
}

func TestReleaseLinksService_UpdateReleaseLink(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1/assets/links/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			fmt.Fprint(w, exampleReleaseLink)
		})

	releaseLink, _, err := client.ReleaseLinks.UpdateReleaseLink(
		1, exampleTagName, 1,
		&UpdateReleaseLinkOptions{
			Name: String(exampleReleaseName),
		})
	if err != nil {
		t.Error(err)
	}
	if releaseLink.Name != exampleReleaseName {
		t.Errorf("release link name, expected '%s', got '%s'", exampleReleaseName,
			releaseLink.Name)
	}
}

func TestReleaseLinksService_DeleteReleaseLink(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1/assets/links/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			fmt.Fprint(w, exampleReleaseLink)
		})

	releaseLink, _, err := client.ReleaseLinks.DeleteReleaseLink(1, exampleTagName, 1)
	if err != nil {
		t.Error(err)
	}
	if releaseLink.Name != exampleReleaseName {
		t.Errorf("release link name, expected '%s', got '%s'", exampleReleaseName,
			releaseLink.Name)
	}
}
