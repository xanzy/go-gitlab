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

func TestListAllTemplates(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/templates/gitlab_ci_ymls", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			   "key":"5-Minute-Production-App",
			   "name":"5-Minute-Production-App"
			},
			{
			   "key":"Android",
			   "name":"Android"
			},
			{
			   "key":"Android-Fastlane",
			   "name":"Android-Fastlane"
			},
			{
			   "key":"Auto-DevOps",
			   "name":"Auto-DevOps"
			}
		 ]`)
	})

	templates, _, err := client.CIYMLTemplate.ListAllTemplates(&ListCIYMLTemplatesOptions{})
	if err != nil {
		t.Errorf("CIYMLTemplates.ListAllTemplates returned error: %v", err)
	}

	want := []*CIYMLTemplateListItem{
		{
			Key:  "5-Minute-Production-App",
			Name: "5-Minute-Production-App",
		},
		{
			Key:  "Android",
			Name: "Android",
		},
		{
			Key:  "Android-Fastlane",
			Name: "Android-Fastlane",
		},
		{
			Key:  "Auto-DevOps",
			Name: "Auto-DevOps",
		},
	}
	if !reflect.DeepEqual(want, templates) {
		t.Errorf("CIYMLTemplates.ListAllTemplates returned %+v, want %+v", templates, want)
	}
}

func TestGetTemplate(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/templates/gitlab_ci_ymls/Ruby", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
			"name": "Ruby",
			"content": "# This file is a template, and might need editing before it works on your project."
		  }`)
	})

	template, _, err := client.CIYMLTemplate.GetTemplate("Ruby")
	if err != nil {
		t.Errorf("CIYMLTemplates.GetTemplate returned error: %v", err)
	}

	want := &CIYMLTemplate{
		Name:    "Ruby",
		Content: "# This file is a template, and might need editing before it works on your project.",
	}
	if !reflect.DeepEqual(want, template) {
		t.Errorf("CIYMLTemplates.GetTemplate returned %+v, want %+v", template, want)
	}
}
