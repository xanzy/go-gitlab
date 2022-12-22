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
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/templates/gitignores", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			  "key": "Actionscript",
			  "name": "Actionscript"
			},
			{
			  "key": "Ada",
			  "name": "Ada"
			},
			{
			  "key": "Agda",
			  "name": "Agda"
			},
			{
			  "key": "Android",
			  "name": "Android"
			},
			{
			  "key": "AppEngine",
			  "name": "AppEngine"
			},
			{
			  "key": "AppceleratorTitanium",
			  "name": "AppceleratorTitanium"
			},
			{
			  "key": "ArchLinuxPackages",
			  "name": "ArchLinuxPackages"
			},
			{
			  "key": "Autotools",
			  "name": "Autotools"
			},
			{
			  "key": "C",
			  "name": "C"
			},
			{
			  "key": "C++",
			  "name": "C++"
			},
			{
			  "key": "CFWheels",
			  "name": "CFWheels"
			},
			{
			  "key": "CMake",
			  "name": "CMake"
			},
			{
			  "key": "CUDA",
			  "name": "CUDA"
			},
			{
			  "key": "CakePHP",
			  "name": "CakePHP"
			},
			{
			  "key": "ChefCookbook",
			  "name": "ChefCookbook"
			},
			{
			  "key": "Clojure",
			  "name": "Clojure"
			},
			{
			  "key": "CodeIgniter",
			  "name": "CodeIgniter"
			},
			{
			  "key": "CommonLisp",
			  "name": "CommonLisp"
			},
			{
			  "key": "Composer",
			  "name": "Composer"
			},
			{
			  "key": "Concrete5",
			  "name": "Concrete5"
			}
		  ]`)
	})

	templates, _, err := client.GitIgnoreTemplates.ListTemplates(&ListTemplatesOptions{})
	if err != nil {
		t.Errorf("GitIgnoreTemplates.ListTemplates returned error: %v", err)
	}

	want := []*GitIgnoreTemplateListItem{
		{
			Key:  "Actionscript",
			Name: "Actionscript",
		},
		{
			Key:  "Ada",
			Name: "Ada",
		},
		{
			Key:  "Agda",
			Name: "Agda",
		},
		{
			Key:  "Android",
			Name: "Android",
		},
		{
			Key:  "AppEngine",
			Name: "AppEngine",
		},
		{
			Key:  "AppceleratorTitanium",
			Name: "AppceleratorTitanium",
		},
		{
			Key:  "ArchLinuxPackages",
			Name: "ArchLinuxPackages",
		},
		{
			Key:  "Autotools",
			Name: "Autotools",
		},
		{
			Key:  "C",
			Name: "C",
		},
		{
			Key:  "C++",
			Name: "C++",
		},
		{
			Key:  "CFWheels",
			Name: "CFWheels",
		},
		{
			Key:  "CMake",
			Name: "CMake",
		},
		{
			Key:  "CUDA",
			Name: "CUDA",
		},
		{
			Key:  "CakePHP",
			Name: "CakePHP",
		},
		{
			Key:  "ChefCookbook",
			Name: "ChefCookbook",
		},
		{
			Key:  "Clojure",
			Name: "Clojure",
		},
		{
			Key:  "CodeIgniter",
			Name: "CodeIgniter",
		},
		{
			Key:  "CommonLisp",
			Name: "CommonLisp",
		},
		{
			Key:  "Composer",
			Name: "Composer",
		},
		{
			Key:  "Concrete5",
			Name: "Concrete5",
		},
	}
	if !reflect.DeepEqual(want, templates) {
		t.Errorf("GitIgnoreTemplates.ListTemplates returned %+v, want %+v", templates, want)
	}
}

func TestGetTemplates(t *testing.T) {
	mux, client := setup(t)

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
