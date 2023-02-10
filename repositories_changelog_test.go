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
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddChangelogData(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/changelog",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			w.WriteHeader(http.StatusOK)
		})

	resp, err := client.Repositories.AddChangelog(
		1,
		&ChangelogOptions{
			Version: "1.0.0",
		})

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGenerateChangelogData(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/changelog",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, exampleChangelogResponse)
		})

	notes, _, err := client.Repositories.GenerateChangelog(
		1,
		ChangelogOptions{
			Version: "1.0.0",
		},
	)
	require.NoError(t, err)
	expectedNotes := &GeneratedChangelogNotes{
		Notes: "## 1.0.0 (2021-11-17)\n\n### feature (2 changes)\n\n- [Title 2](namespace13/project13@ad608eb642124f5b3944ac0ac772fecaf570a6bf) ([merge request](namespace13/project13!2))\n- [Title 1](namespace13/project13@3c6b80ff7034fa0d585314e1571cc780596ce3c8) ([merge request](namespace13/project13!1))\n",
	}
	assert.Equal(t, expectedNotes, notes)
}
