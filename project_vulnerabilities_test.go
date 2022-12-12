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

func TestListProjectVulnerabilities(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/vulnerabilities", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectVulnerabilitiesOptions{
		ListOptions: ListOptions{2, 3},
	}

	projectVulnerabilities, _, err := client.ProjectVulnerabilities.ListProjectVulnerabilities(1, opt)
	if err != nil {
		t.Errorf("ProjectVulnerabilities.ListProjectVulnerabilities returned error: %v", err)
	}

	want := []*ProjectVulnerability{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, projectVulnerabilities) {
		t.Errorf("ProjectVulnerabilities.ListProjectVulnerabilities returned %+v, want %+v", projectVulnerabilities, want)
	}
}

func TestCreateVulnerability(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/vulnerabilities", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1}`)
	})

	opt := &CreateVulnerabilityOptions{
		FindingID: Int(1),
	}

	projectVulnerability, _, err := client.ProjectVulnerabilities.CreateVulnerability(1, opt)
	if err != nil {
		t.Errorf("ProjectVulnerabilities.CreateVulnerability returned error: %v", err)
	}

	want := &ProjectVulnerability{ID: 1}
	if !reflect.DeepEqual(want, projectVulnerability) {
		t.Errorf("ProjectVulnerabilities.CreateVulnerability returned %+v, want %+v", projectVulnerability, want)
	}
}
