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
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestReleasesService_ListReleases(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, exampleReleaseListResponse)
		})

	opt := &ListReleasesOptions{}
	releases, _, err := client.Releases.ListReleases(1, opt)
	if err != nil {
		t.Error(err)
	}
	if len(releases) != 2 {
		t.Error("expected 2 releases")
	}
}

func TestReleasesService_GetRelease(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, exampleReleaseResponse)
		})

	release, _, err := client.Releases.GetRelease(1, exampleTagName)
	if err != nil {
		t.Error(err)
	}
	if release.TagName != exampleTagName {
		t.Errorf("expected tag %s, got %s", exampleTagName, release.TagName)
	}
}

func TestReleasesService_CreateRelease(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("unable to read request body")
			}
			if !strings.Contains(string(b), exampleTagName) {
				t.Errorf("expected request body to contain %s, got %s",
					exampleTagName, string(b))
			}
			if strings.Contains(string(b), "assets") {
				t.Errorf("expected request body not to have assets, got %s",
					string(b))
			}
			if strings.Contains(string(b), "milestones") {
				t.Errorf("expected request body not to have milestones, got %s",
					string(b))
			}
			if strings.Contains(string(b), "released_at") {
				t.Errorf("expected request body not to have released_at, got %s",
					string(b))
			}
			fmt.Fprint(w, exampleReleaseResponse)
		})

	opts := &CreateReleaseOptions{
		Name:        String("name"),
		TagName:     String(exampleTagName),
		Description: String("Description"),
	}

	release, _, err := client.Releases.CreateRelease(1, opts)
	if err != nil {
		t.Error(err)
	}
	if release.TagName != exampleTagName {
		t.Errorf("expected tag %s, got %s", exampleTagName, release.TagName)
	}
}

func TestReleasesService_CreateReleaseWithAsset(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("unable to read request body")
			}
			if !strings.Contains(string(b), exampleTagName) {
				t.Errorf("expected request body to contain %s, got %s",
					exampleTagName, string(b))
			}
			if !strings.Contains(string(b), "assets") {
				t.Errorf("expected request body to have assets, got %s",
					string(b))
			}
			if strings.Contains(string(b), "milestones") {
				t.Errorf("expected request body not to have milestones, got %s",
					string(b))
			}
			if strings.Contains(string(b), "released_at") {
				t.Errorf("expected request body not to have released_at, got %s",
					string(b))
			}
			fmt.Fprint(w, exampleReleaseResponse)
		})

	opts := &CreateReleaseOptions{
		Name:        String("name"),
		TagName:     String(exampleTagName),
		Description: String("Description"),
		Assets: &ReleaseAssetsOptions{
			Links: []*ReleaseAssetLinkOptions{
				{String("sldkf"), String("sldkfj"), String("sldkfh"), LinkType(OtherLinkType)},
			},
		},
	}

	release, _, err := client.Releases.CreateRelease(1, opts)
	if err != nil {
		t.Error(err)
	}
	if release.TagName != exampleTagName {
		t.Errorf("expected tag %s, got %s", exampleTagName, release.TagName)
	}
}

func TestReleasesService_CreateReleaseWithAssetAndNameMetadata(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("unable to read request body")
			}
			if !strings.Contains(string(b), exampleTagNameWithMetadata) {
				t.Errorf("expected request body to contain %s, got %s",
					exampleTagNameWithMetadata, string(b))
			}
			if !strings.Contains(string(b), "assets") {
				t.Errorf("expected request body to have assets, got %s",
					string(b))
			}
			if strings.Contains(string(b), "milestones") {
				t.Errorf("expected request body not to have milestones, got %s",
					string(b))
			}
			if strings.Contains(string(b), "released_at") {
				t.Errorf("expected request body not to have released_at, got %s",
					string(b))
			}
			fmt.Fprint(w, exampleReleaseWithMetadataResponse)
		})

	opts := &CreateReleaseOptions{
		Name:        String("name"),
		TagName:     String(exampleTagNameWithMetadata),
		Description: String("Description"),
		Assets: &ReleaseAssetsOptions{
			Links: []*ReleaseAssetLinkOptions{
				{String("sldkf"), String("sldkfj"), String("sldkfh"), LinkType(OtherLinkType)},
			},
		},
	}

	release, _, err := client.Releases.CreateRelease(1, opts)
	if err != nil {
		t.Error(err)
	}
	if release.TagName != exampleTagNameWithMetadata {
		t.Errorf("expected tag %s, got %s", exampleTagNameWithMetadata, release.TagName)
	}
}

func TestReleasesService_CreateReleaseWithMilestones(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("unable to read request body")
			}
			if !strings.Contains(string(b), exampleTagName) {
				t.Errorf("expected request body to contain %s, got %s",
					exampleTagName, string(b))
			}
			if strings.Contains(string(b), "assets") {
				t.Errorf("expected request body not to have assets, got %s",
					string(b))
			}
			if !strings.Contains(string(b), "milestones") {
				t.Errorf("expected request body to have milestones, got %s",
					string(b))
			}
			if strings.Contains(string(b), "released_at") {
				t.Errorf("expected request body not to have released_at, got %s",
					string(b))
			}
			fmt.Fprint(w, exampleReleaseResponse)
		})

	opts := &CreateReleaseOptions{
		Name:        String("name"),
		TagName:     String(exampleTagName),
		Description: String("Description"),
		Milestones:  &[]string{exampleTagName, "v0.1.0"},
	}

	release, _, err := client.Releases.CreateRelease(1, opts)
	if err != nil {
		t.Error(err)
	}
	if release.TagName != exampleTagName {
		t.Errorf("expected tag %s, got %s", exampleTagName, release.TagName)
	}
}

func TestReleasesService_CreateReleaseWithReleasedAt(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("unable to read request body")
			}
			if !strings.Contains(string(b), exampleTagName) {
				t.Errorf("expected request body to contain %s, got %s",
					exampleTagName, string(b))
			}
			if strings.Contains(string(b), "assets") {
				t.Errorf("expected request body not to have assets, got %s",
					string(b))
			}
			if strings.Contains(string(b), "milestones") {
				t.Errorf("expected request body not to have milestones, got %s",
					string(b))
			}
			if !strings.Contains(string(b), "released_at") {
				t.Errorf("expected request body to have released_at, got %s",
					string(b))
			}
			fmt.Fprint(w, exampleReleaseResponse)
		})

	opts := &CreateReleaseOptions{
		Name:        String("name"),
		TagName:     String(exampleTagName),
		Description: String("Description"),
		ReleasedAt:  &time.Time{},
	}

	release, _, err := client.Releases.CreateRelease(1, opts)
	if err != nil {
		t.Error(err)
	}
	if release.TagName != exampleTagName {
		t.Errorf("expected tag %s, got %s", exampleTagName, release.TagName)
	}
}

func TestReleasesService_UpdateRelease(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("unable to read request body")
			}
			if strings.Contains(string(b), "milestones") {
				t.Errorf("expected request body not to have milestones, got %s",
					string(b))
			}
			if strings.Contains(string(b), "released_at") {
				t.Errorf("expected request body not to have released_at, got %s",
					string(b))
			}
			fmt.Fprint(w, exampleReleaseResponse)
		})

	opts := &UpdateReleaseOptions{
		Name:        String("name"),
		Description: String("Description"),
	}

	release, _, err := client.Releases.UpdateRelease(1, exampleTagName, opts)
	if err != nil {
		t.Error(err)
	}
	if release.TagName != exampleTagName {
		t.Errorf("expected tag %s, got %s", exampleTagName, release.TagName)
	}
}

func TestReleasesService_UpdateReleaseWithMilestones(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("unable to read request body")
			}
			if !strings.Contains(string(b), "milestones") {
				t.Errorf("expected request body to have milestones, got %s",
					string(b))
			}
			if strings.Contains(string(b), "released_at") {
				t.Errorf("expected request body not to have released_at, got %s",
					string(b))
			}
			fmt.Fprint(w, exampleReleaseResponse)
		})

	opts := &UpdateReleaseOptions{
		Name:        String("name"),
		Description: String("Description"),
		Milestones:  &[]string{exampleTagName, "v0.1.0"},
	}

	release, _, err := client.Releases.UpdateRelease(1, exampleTagName, opts)
	if err != nil {
		t.Error(err)
	}
	if release.TagName != exampleTagName {
		t.Errorf("expected tag %s, got %s", exampleTagName, release.TagName)
	}
}

func TestReleasesService_UpdateReleaseWithReleasedAt(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("unable to read request body")
			}
			if strings.Contains(string(b), "milestones") {
				t.Errorf("expected request body not to have milestones, got %s",
					string(b))
			}
			if !strings.Contains(string(b), "released_at") {
				t.Errorf("expected request body to have released_at, got %s",
					string(b))
			}
			fmt.Fprint(w, exampleReleaseResponse)
		})

	opts := &UpdateReleaseOptions{
		Name:        String("name"),
		Description: String("Description"),
		ReleasedAt:  &time.Time{},
	}

	release, _, err := client.Releases.UpdateRelease(1, exampleTagName, opts)
	if err != nil {
		t.Error(err)
	}
	if release.TagName != exampleTagName {
		t.Errorf("expected tag %s, got %s", exampleTagName, release.TagName)
	}
}

func TestReleasesService_DeleteRelease(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			fmt.Fprint(w, exampleReleaseResponse)
		})

	release, _, err := client.Releases.DeleteRelease(1, exampleTagName)
	if err != nil {
		t.Error(err)
	}
	if release.TagName != exampleTagName {
		t.Errorf("expected tag %s, got %s", exampleTagName, release.TagName)
	}
}
