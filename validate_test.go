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

func TestValidate(t *testing.T) {
	testCases := []struct {
		description string
		opts        *LintOptions
		response    string
		want        *LintResult
	}{
		{
			description: "valid",
			opts: &LintOptions{
				Content: `
				build1:
					stage: build
					script:
						- echo "Do your build here"`,
				IncludeMergedYAML: true,
				IncludeJobs:       false,
			},
			response: `{
				"status": "valid",
				"errors": [],
				"merged_yaml":"---\nbuild1:\n    stage: build\n    script:\n    - echo\"Do your build here\""
			}`,
			want: &LintResult{
				Status:     "valid",
				MergedYaml: "---\nbuild1:\n    stage: build\n    script:\n    - echo\"Do your build here\"",
				Errors:     []string{},
			},
		},
		{
			description: "invalid",
			opts: &LintOptions{
				Content: `
					build1:
						- echo "Do your build here"`,
			},
			response: `{
				"status": "invalid",
				"errors": ["error message when content is invalid"]
			}`,
			want: &LintResult{
				Status: "invalid",
				Errors: []string{"error message when content is invalid"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			mux, client := setup(t)

			mux.HandleFunc("/api/v4/ci/lint", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodPost)
				fmt.Fprint(w, tc.response)
			})

			got, _, err := client.Validate.Lint(tc.opts)
			if err != nil {
				t.Errorf("Validate returned error: %v", err)
			}

			want := tc.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Validate returned \ngot:\n%v\nwant:\n%v", Stringify(got), Stringify(want))
			}
		})
	}
}

func TestValidateProject(t *testing.T) {
	testCases := []struct {
		description string
		response    string
		want        *ProjectLintResult
	}{
		{
			description: "valid",
			response: `{
				"valid": true,
				"errors": [],
				"warnings": [],
				"merged_yaml": 	"---\n:build:\n  :script:\n  - echo build"
			}`,
			want: &ProjectLintResult{
				Valid:      true,
				Warnings:   []string{},
				Errors:     []string{},
				MergedYaml: "---\n:build:\n  :script:\n  - echo build",
			},
		},
		{
			description: "invalid",
			response: `{
				"valid": false,
				"errors": ["jobs:build config contains unknown keys: bad_key"],
				"warnings": [],
				"merged_yaml": 	"---\n:build:\n  :script:\n  - echo build\n  :bad_key: value"
			}`,
			want: &ProjectLintResult{
				Valid:      false,
				Warnings:   []string{},
				Errors:     []string{"jobs:build config contains unknown keys: bad_key"},
				MergedYaml: "---\n:build:\n  :script:\n  - echo build\n  :bad_key: value",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			mux, client := setup(t)

			mux.HandleFunc("/api/v4/projects/1/ci/lint", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				fmt.Fprint(w, tc.response)
			})

			opt := &ProjectLintOptions{}
			got, _, err := client.Validate.ProjectLint(1, opt)
			if err != nil {
				t.Errorf("Validate returned error: %v", err)
			}

			want := tc.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Validate returned \ngot:\n%v\nwant:\n%v", Stringify(got), Stringify(want))
			}
		})
	}
}

func TestValidateProjectNamespace(t *testing.T) {
	testCases := []struct {
		description string
		request     *ProjectNamespaceLintOptions
		response    string
		want        *ProjectLintResult
	}{
		{
			description: "valid",
			request: &ProjectNamespaceLintOptions{
				Content:     String("{'build': {'script': 'echo build'}}"),
				DryRun:      Bool(false),
				IncludeJobs: Bool(true),
				Ref:         String("foo"),
			},
			response: `{
				"valid": true,
				"errors": [],
				"warnings": [],
				"merged_yaml": 	"---\n:build:\n  :script:\n  - echo build"
			}`,
			want: &ProjectLintResult{
				Valid:      true,
				Warnings:   []string{},
				Errors:     []string{},
				MergedYaml: "---\n:build:\n  :script:\n  - echo build",
			},
		},
		{
			description: "invalid",
			request: &ProjectNamespaceLintOptions{
				Content: String("{'build': {'script': 'echo build', 'bad_key': 'value'}}"),
				DryRun:  Bool(false),
			},
			response: `{
				"valid": false,
				"errors": ["jobs:build config contains unknown keys: bad_key"],
				"warnings": [],
				"merged_yaml": 	"---\n:build:\n  :script:\n  - echo build\n  :bad_key: value"
			}`,
			want: &ProjectLintResult{
				Valid:      false,
				Warnings:   []string{},
				Errors:     []string{"jobs:build config contains unknown keys: bad_key"},
				MergedYaml: "---\n:build:\n  :script:\n  - echo build\n  :bad_key: value",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			mux, client := setup(t)

			mux.HandleFunc("/api/v4/projects/1/ci/lint", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodPost)
				fmt.Fprint(w, tc.response)
			})

			got, _, err := client.Validate.ProjectNamespaceLint(1, tc.request)
			if err != nil {
				t.Errorf("Validate returned error: %v", err)
			}

			want := tc.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Validate returned \ngot:\n%v\nwant:\n%v", Stringify(got), Stringify(want))
			}
		})
	}
}
