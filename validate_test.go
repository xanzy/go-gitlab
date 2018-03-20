package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

// TODO: better test posting empty string
func TestValidate(t *testing.T) {
	validContent := `
	build1:
		stage: build
		script:
			- echo "Do your build here"`

	invalidContent := `
	build1:
		- echo "Do your build here"`

	validRes := `{
			"status": "valid",
			"errors": []
		}`

	invalidRes := `{
			"status": "invalid",
			"errors": [
				"error message when content is invalid"
			]
		}`

	wantValid := &LintResult{
		Status: "valid",
		Errors: []string{},
	}

	wantInvalid := &LintResult{
		Status: "invalid",
		Errors: []string{"error message when content is invalid"},
	}

	testCases := []struct {
		desc    string
		content string
		res     string
		want    *LintResult
	}{
		{"valid", validContent, validRes, wantValid},
		{"invalid", invalidContent, invalidRes, wantInvalid},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mux, server, client := setup()
			defer teardown(server)

			mux.HandleFunc("/ci/lint", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				fmt.Fprintf(w, tc.res)
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
