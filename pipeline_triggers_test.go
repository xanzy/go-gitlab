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

func TestRunPipeline(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/trigger/pipeline", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1, "status":"pending"}`)
	})

	opt := &RunPipelineTriggerOptions{Ref: String("master")}
	pipeline, _, err := client.PipelineTriggers.RunPipelineTrigger(1, opt)
	if err != nil {
		t.Errorf("PipelineTriggers.RunPipelineTrigger returned error: %v", err)
	}

	want := &Pipeline{ID: 1, Status: "pending"}
	if !reflect.DeepEqual(want, pipeline) {
		t.Errorf("PipelineTriggers.RunPipelineTrigger returned %+v, want %+v", pipeline, want)
	}
}
