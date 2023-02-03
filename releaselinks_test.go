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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReleaseLinksService_ListReleaseLinks(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1/assets/links",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, exampleReleaseLinkList)
		})

	releaseLinks, _, err := client.ReleaseLinks.ListReleaseLinks(
		1, exampleTagName, &ListReleaseLinksOptions{},
	)

	require.NoError(t, err)
	expectedReleaseLinks := []*ReleaseLink{
		{
			ID:       2,
			Name:     "awesome-v0.2.msi",
			URL:      "http://192.168.10.15:3000/msi",
			External: true,
		},
		{
			ID:             1,
			Name:           "awesome-v0.2.dmg",
			URL:            "http://192.168.10.15:3000",
			DirectAssetURL: "http://192.168.10.15:3000/namespace/example/-/releases/v0.1/downloads/awesome-v0.2.dmg",
			External:       false,
			LinkType:       OtherLinkType,
		},
	}
	assert.Equal(t, expectedReleaseLinks, releaseLinks)
}

func TestReleaseLinksService_CreateReleaseLink(t *testing.T) {
	testCases := []struct {
		description string
		options     *CreateReleaseLinkOptions
		response    string
		want        *ReleaseLink
	}{
		{
			description: "Mandatory Attributes",
			options: &CreateReleaseLinkOptions{
				Name: String("awesome-v0.2.dmg"),
				URL:  String("http://192.168.10.15:3000"),
			},
			response: `{
				"id":1,
				"name":"awesome-v0.2.dmg",
				"url":"http://192.168.10.15:3000",
				"external":true
			}`,
			want: &ReleaseLink{
				ID:       1,
				Name:     "awesome-v0.2.dmg",
				URL:      "http://192.168.10.15:3000",
				External: true,
			},
		},
		{
			description: "Optional Attributes",
			options: &CreateReleaseLinkOptions{
				Name:     String("release-notes.md"),
				URL:      String("http://192.168.10.15:3000"),
				FilePath: String("docs/release-notes.md"),
				LinkType: LinkType(OtherLinkType),
			},
			response: `{
				"id":1,
				"name":"release-notes.md",
				"url":"http://192.168.10.15:3000",
				"direct_asset_url": "http://192.168.10.15:3000/namespace/example/-/releases/v0.1/downloads/docs/release-notes.md",
				"external": false,
				"link_type": "other"
			}`,
			want: &ReleaseLink{
				ID:             1,
				Name:           "release-notes.md",
				URL:            "http://192.168.10.15:3000",
				DirectAssetURL: "http://192.168.10.15:3000/namespace/example/-/releases/v0.1/downloads/docs/release-notes.md",
				External:       false,
				LinkType:       OtherLinkType,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			mux, client := setup(t)

			mux.HandleFunc("/api/v4/projects/1/releases/v0.1/assets/links",
				func(w http.ResponseWriter, r *http.Request) {
					testMethod(t, r, http.MethodPost)
					fmt.Fprint(w, tc.response)
				})

			releaseLink, _, err := client.ReleaseLinks.CreateReleaseLink(1, exampleTagName, tc.options)

			require.NoError(t, err)
			assert.Equal(t, tc.want, releaseLink)
		})
	}
}

func TestReleaseLinksService_GetReleaseLink(t *testing.T) {
	mux, client := setup(t)

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
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1/assets/links/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			fmt.Fprint(w, exampleReleaseLink)
		})

	releaseLink, _, err := client.ReleaseLinks.UpdateReleaseLink(
		1, exampleTagName, 1,
		&UpdateReleaseLinkOptions{
			Name:     String(exampleReleaseName),
			FilePath: String("http://192.168.10.15:3000/namespace/example/-/releases/v0.1/downloads/awesome-v0.2.dmg"),
			LinkType: LinkType(OtherLinkType),
		})

	require.NoError(t, err)
	expectedRelease := &ReleaseLink{
		ID:             1,
		Name:           "awesome-v0.2.dmg",
		URL:            "http://192.168.10.15:3000",
		DirectAssetURL: "http://192.168.10.15:3000/namespace/example/-/releases/v0.1/downloads/awesome-v0.2.dmg",
		External:       true,
		LinkType:       OtherLinkType,
	}
	assert.Equal(t, expectedRelease, releaseLink)
}

func TestReleaseLinksService_DeleteReleaseLink(t *testing.T) {
	mux, client := setup(t)

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
