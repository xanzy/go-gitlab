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

func TestGetPipelinesEmailService(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/services/pipelines-email", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &PipelinesEmailService{Service: Service{ID: 1}}

	service, _, err := client.Services.GetPipelinesEmailService(1)
	if err != nil {
		t.Fatalf("Services.GetPipelinesEmailService returns an error: %v", err)
	}
	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.GetPipelinesEmailService returned %+v, want %+v", service, want)
	}
}

func TestSetPipelinesEmailService(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/services/pipelines-email", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	opt := &SetPipelinesEmailServiceOptions{
		Recipients:                String("test@email.com"),
		NotifyOnlyBrokenPipelines: Bool(true),
		NotifyOnlyDefaultBranch:   Bool(false),
		AddPusher:                 nil,
		BranchesToBeNotified:      nil,
		PipelineEvents:            nil,
	}

	_, err := client.Services.SetPipelinesEmailService(1, opt)
	if err != nil {
		t.Fatalf("Services.SetPipelinesEmailService returns an error: %v", err)
	}
}

func TestDeletePipelinesEmailService(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/services/pipelines-email", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Services.DeletePipelinesEmailService(1)
	if err != nil {
		t.Fatalf("Services.DeletePipelinesEmailService returns an error: %v", err)
	}
}
