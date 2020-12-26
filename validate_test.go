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
		content     string
		response    string
		want        *LintResult
	}{
		{
			description: "valid",
			content: `
				build1:
					stage: build
					script:
						- echo "Do your build here"`,
			response: `{
				"status": "valid",
				"errors": []
			}`,
			want: &LintResult{
				Status: "valid",
				Errors: []string{},
			},
		},
		{
			description: "invalid",
			content: `
				build1:
					- echo "Do your build here"`,
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
			mux, server, client := setup(t)
			defer teardown(server)

			mux.HandleFunc("/api/v4/ci/lint", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				fmt.Fprint(w, tc.response)
			})

			got, _, err := client.Validate.Lint(tc.content)

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
			mux, server, client := setup(t)
			defer teardown(server)

			mux.HandleFunc("/api/v4/projects/1/ci/lint", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
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
				Content: String("{'build': {'script': 'echo build'}}"),
				DryRun:  Bool(false),
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
			mux, server, client := setup(t)
			defer teardown(server)

			mux.HandleFunc("/api/v4/projects/1/ci/lint", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
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
