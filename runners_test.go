//
// Copyright 2017, Sander van Harmelen
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

func TestDisableRunner(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/runners/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.Runners.DisableProjectRunner(1, 2, nil)
	if err != nil {
		t.Fatalf("Runners.DisableProjectRunner returns an error: %v", err)
	}
}

func TestListRunnersJobs(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/runners/1/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListRunnersJobsOptions{}

	jobs, _, err := client.Runners.ListRunnersJobs(1, opt)
	if err != nil {
		t.Fatalf("Runners.ListRunnersJobs returns an error: %v", err)
	}

	want := []*Job{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, jobs) {
		t.Errorf("Runners.ListRunnersJobs returned %+v, want %+v", jobs, want)
	}
}
