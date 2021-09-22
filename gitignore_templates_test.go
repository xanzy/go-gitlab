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

func TestListTemplates(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/templates/gitignores", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			  "content": "Actionscript",
			  "name": "Actionscript"
			},
			{
			  "content": "Ada",
			  "name": "Ada"
			},
			{
			  "content": "Agda",
			  "name": "Agda"
			},
			{
			  "content": "Android",
			  "name": "Android"
			},
			{
			  "content": "AppEngine",
			  "name": "AppEngine"
			},
			{
			  "content": "AppceleratorTitanium",
			  "name": "AppceleratorTitanium"
			},
			{
			  "content": "ArchLinuxPackages",
			  "name": "ArchLinuxPackages"
			},
			{
			  "content": "Autotools",
			  "name": "Autotools"
			},
			{
			  "content": "C",
			  "name": "C"
			},
			{
			  "content": "C++",
			  "name": "C++"
			},
			{
			  "content": "CFWheels",
			  "name": "CFWheels"
			},
			{
			  "content": "CMake",
			  "name": "CMake"
			},
			{
			  "content": "CUDA",
			  "name": "CUDA"
			},
			{
			  "content": "CakePHP",
			  "name": "CakePHP"
			},
			{
			  "content": "ChefCookbook",
			  "name": "ChefCookbook"
			},
			{
			  "content": "Clojure",
			  "name": "Clojure"
			},
			{
			  "content": "CodeIgniter",
			  "name": "CodeIgniter"
			},
			{
			  "content": "CommonLisp",
			  "name": "CommonLisp"
			},
			{
			  "content": "Composer",
			  "name": "Composer"
			},
			{
			  "content": "Concrete5",
			  "name": "Concrete5"
			}
		  ]`)
	})

	templates, _, err := client.GitIgnoreTemplates.ListTemplates(&ListTemplatesOptions{})
	if err != nil {
		t.Errorf("GitIgnoreTemplates.ListTemplates returned error: %v", err)
	}

	want := []*GitIgnoreTemplate{
		{
			Name:    "Actionscript",
			Content: "Actionscript",
		},
		{
			Name:    "Ada",
			Content: "Ada",
		},
		{
			Name:    "Agda",
			Content: "Agda",
		},
		{
			Name:    "Android",
			Content: "Android",
		},
		{
			Name:    "AppEngine",
			Content: "AppEngine",
		},
		{
			Name:    "AppceleratorTitanium",
			Content: "AppceleratorTitanium",
		},
		{
			Name:    "ArchLinuxPackages",
			Content: "ArchLinuxPackages",
		},
		{
			Name:    "Autotools",
			Content: "Autotools",
		},
		{
			Name:    "C",
			Content: "C",
		},
		{
			Name:    "C++",
			Content: "C++",
		},
		{
			Name:    "CFWheels",
			Content: "CFWheels",
		},
		{
			Name:    "CMake",
			Content: "CMake",
		},
		{
			Name:    "CUDA",
			Content: "CUDA",
		},
		{
			Name:    "CakePHP",
			Content: "CakePHP",
		},
		{
			Name:    "ChefCookbook",
			Content: "ChefCookbook",
		},
		{
			Name:    "Clojure",
			Content: "Clojure",
		},
		{
			Name:    "CodeIgniter",
			Content: "CodeIgniter",
		},
		{
			Name:    "CommonLisp",
			Content: "CommonLisp",
		},
		{
			Name:    "Composer",
			Content: "Composer",
		},
		{
			Name:    "Concrete5",
			Content: "Concrete5",
		},
	}
	if !reflect.DeepEqual(want, templates) {
		t.Errorf("GitIgnoreTemplates.ListTemplates returned %+v, want %+v", templates, want)
	}
}

func TestGetTemplates(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/templates/gitignores/Ruby", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
			"name": "Ruby",
			"content": "*.gem\n*.rbc\n/.config\n/coverage/\n/InstalledFiles\n/pkg/\n/spec/reports/"
		  }`)
	})

	template, _, err := client.GitIgnoreTemplates.GetTemplate("Ruby")
	if err != nil {
		t.Errorf("GitIgnoreTempaltes.GetTemplate returned an error: %v", err)
	}

	want := &GitIgnoreTemplate{
		Name:    "Ruby",
		Content: "*.gem\n*.rbc\n/.config\n/coverage/\n/InstalledFiles\n/pkg/\n/spec/reports/",
	}
	if !reflect.DeepEqual(want, template) {
		t.Errorf("GitIgnoreTemplates.GetTemplate returned %+v, want %+v", template, want)
	}
}
