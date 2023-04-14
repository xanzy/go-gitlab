package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProjectTemplatesService_ListTemplates(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/templates/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
				{
					"key": "epl-1.0",
					"name": "Eclipse Public License 1.0"
				  },
				  {
					"key": "lgpl-3.0",
					"name": "GNU Lesser General Public License v3.0"
				  }
			]
		`)
	})

	want := []*ProjectTemplate{
		{
			Key:  "epl-1.0",
			Name: "Eclipse Public License 1.0",
		},
		{
			Key:  "lgpl-3.0",
			Name: "GNU Lesser General Public License v3.0",
		},
	}

	ss, resp, err := client.ProjectTemplates.ListTemplates(1, "issues", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ss)
}

func TestProjectTemplatesService_GetProjectTemplate(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/templates/issues/test_issue", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
			  "name": "test_issue",
			  "content": "## Test"
			}
		`)
	})

	want := &ProjectTemplate{
		Name:    "test_issue",
		Content: "## Test",
	}

	ss, resp, err := client.ProjectTemplates.GetProjectTemplate(1, "issues", "test_issue", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ss)
}
