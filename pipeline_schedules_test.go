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

func TestRunPipelineSchedule(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/pipeline_schedules/1/play", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"message": "201 Created"}`)
	})

	res, err := client.PipelineSchedules.RunPipelineSchedule(1, 1)
	if err != nil {
		t.Errorf("PipelineTriggers.RunPipelineTrigger returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("PipelineSchedules.RunPipelineSchedule returned status %v, want %v", res.StatusCode, http.StatusCreated)
	}
}
