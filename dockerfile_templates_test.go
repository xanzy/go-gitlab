//
// Copyright 2022, FantasyTeddy
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

func TestDockerfileTemplatesService_ListTemplates(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/templates/dockerfiles", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			  "key":"Binary",
			  "name":"Binary"
			},
			{
			  "key":"Binary-alpine",
			  "name":"Binary-alpine"
			},
			{
			  "key":"Binary-scratch",
			  "name":"Binary-scratch"
			},
			{
			  "key":"Golang",
			  "name":"Golang"
			},
			{
			  "key":"Golang-alpine",
			  "name":"Golang-alpine"
			},
			{
			  "key":"Golang-scratch",
			  "name":"Golang-scratch"
			}
		  ]`)
	})

	templates, _, err := client.DockerfileTemplate.ListTemplates(&ListDockerfileTemplatesOptions{})
	if err != nil {
		t.Errorf("DockerfileTemplate.ListTemplates returned error: %v", err)
	}

	want := []*DockerfileTemplateListItem{
		{
			Key:  "Binary",
			Name: "Binary",
		},
		{
			Key:  "Binary-alpine",
			Name: "Binary-alpine",
		},
		{
			Key:  "Binary-scratch",
			Name: "Binary-scratch",
		},
		{
			Key:  "Golang",
			Name: "Golang",
		},
		{
			Key:  "Golang-alpine",
			Name: "Golang-alpine",
		},
		{
			Key:  "Golang-scratch",
			Name: "Golang-scratch",
		},
	}
	if !reflect.DeepEqual(want, templates) {
		t.Errorf("DockerfileTemplate.ListTemplates returned %+v, want %+v", templates, want)
	}
}

func TestDockerfileTemplatesService_GetTemplate(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/templates/dockerfiles/Binary", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
			"name": "Binary",
			"content": "# This file is a template, and might need editing before it works on your project."
		  }`)
	})

	template, _, err := client.DockerfileTemplate.GetTemplate("Binary")
	if err != nil {
		t.Errorf("DockerfileTemplate.GetTemplate returned error: %v", err)
	}

	want := &DockerfileTemplate{
		Name:    "Binary",
		Content: "# This file is a template, and might need editing before it works on your project.",
	}
	if !reflect.DeepEqual(want, template) {
		t.Errorf("DockerfileTemplate.GetTemplate returned %+v, want %+v", template, want)
	}
}
