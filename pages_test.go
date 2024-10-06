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
	"time"

	"github.com/stretchr/testify/require"
)

func TestUnpublishPages(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/2/pages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Pages.UnpublishPages(2)
	if err != nil {
		t.Errorf("Pages.UnpublishPages returned error: %v", err)
	}
}

func TestGetPages(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/projects/2/pages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
		  {
			"url": "https://ssl.domain.example",
			"deployments": [
			  {
				"created_at": "2021-04-27T21:27:38.584Z",
				"url": "https://ssl.domain.example/",
				"path_prefix": "",
				"root_directory": null
			  }
			],
			"is_unique_domain_enabled": false,
			"force_https": false
		  }
		`)
	})

	want := &Pages{
		URL:                   "https://ssl.domain.example",
		IsUniqueDomainEnabled: false,
		ForceHTTPS:            false,
		Deployments: []*PagesDeployment{
			{
				CreatedAt:     time.Date(2021, time.April, 27, 21, 27, 38, 584000000, time.UTC),
				URL:           "https://ssl.domain.example/",
				PathPrefix:    "",
				RootDirectory: "",
			},
		},
	}

	p, resp, err := client.Pages.GetPages(2)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, p)
}
