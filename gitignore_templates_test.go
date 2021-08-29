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
			"content": "*.gem\n*.rbc\n/.config\n/coverage/\n/InstalledFiles\n/pkg/\n/spec/reports/\n/spec/examples.txt\n/test/tmp/\n/test/version_tmp/\n/tmp/\n\n# Used by dotenv library to load environment variables.\n# .env\n\n## Specific to RubyMotion:\n.dat*\n.repl_history\nbuild/\n*.bridgesupport\nbuild-iPhoneOS/\nbuild-iPhoneSimulator/\n\n## Specific to RubyMotion (use of CocoaPods):\n#\n# We recommend against adding the Pods directory to your .gitignore. However\n# you should judge for yourself, the pros and cons are mentioned at:\n# https://guides.cocoapods.org/using/using-cocoapods.html#should-i-check-the-pods-directory-into-source-control\n#\n# vendor/Pods/\n\n## Documentation cache and generated files:\n/.yardoc/\n/_yardoc/\n/doc/\n/rdoc/\n\n## Environment normalization:\n/.bundle/\n/vendor/bundle\n/lib/bundler/man/\n\n# for a library or gem, you might want to ignore these files since the code is\n# intended to run in multiple environments; otherwise, check them in:\n# Gemfile.lock\n# .ruby-version\n# .ruby-gemset\n\n# unless supporting rvm < 1.11.0 or doing something fancy, ignore this:\n.rvmrc\n"
		  }`)
	})

	template, _, err := client.GitIgnoreTemplates.GetTemplate("Ruby")
	if err != nil {
		t.Errorf("GitIgnoreTempaltes.GetTemplate returned an error: %v", err)
	}

	want := &GitIgnoreTemplate{
		Name:    "Ruby",
		Content: "*.gem\n*.rbc\n/.config\n/coverage/\n/InstalledFiles\n/pkg/\n/spec/reports/\n/spec/examples.txt\n/test/tmp/\n/test/version_tmp/\n/tmp/\n\n# Used by dotenv library to load environment variables.\n# .env\n\n## Specific to RubyMotion:\n.dat*\n.repl_history\nbuild/\n*.bridgesupport\nbuild-iPhoneOS/\nbuild-iPhoneSimulator/\n\n## Specific to RubyMotion (use of CocoaPods):\n#\n# We recommend against adding the Pods directory to your .gitignore. However\n# you should judge for yourself, the pros and cons are mentioned at:\n# https://guides.cocoapods.org/using/using-cocoapods.html#should-i-check-the-pods-directory-into-source-control\n#\n# vendor/Pods/\n\n## Documentation cache and generated files:\n/.yardoc/\n/_yardoc/\n/doc/\n/rdoc/\n\n## Environment normalization:\n/.bundle/\n/vendor/bundle\n/lib/bundler/man/\n\n# for a library or gem, you might want to ignore these files since the code is\n# intended to run in multiple environments; otherwise, check them in:\n# Gemfile.lock\n# .ruby-version\n# .ruby-gemset\n\n# unless supporting rvm < 1.11.0 or doing something fancy, ignore this:\n.rvmrc\n",
	}
	if !reflect.DeepEqual(want, template) {
		t.Errorf("GitIgnoreTemplates.GetTemplate returned %+v, want %+v", template, want)
	}
}
