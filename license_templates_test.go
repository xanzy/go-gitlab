package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLicenseTemplatesService_ListLicenseTemplates(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/templates/licenses", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
				{
					"key": "apache-2.0",
					"name": "Apache License 2.0",
					"nickname": null,
					"featured": true,
					"html_url": "http://choosealicense.com/licenses/apache-2.0/",
					"source_url": "http://www.apache.org/licenses/LICENSE-2.0.html",
					"description": "A permissive license that also provides an express grant of patent rights from contributors to users.",
					"conditions": [
						"include-copyright",
						"document-changes"
					],
					"permissions": [
						"commercial-use",
						"modifications",
						"distribution",
						"patent-use",
						"private-use"
					],
					"limitations": [
						"trademark-use",
						"no-liability"
					],
					"content": "Apache License\n Version 2.0, January 2004\n [...]"
				}
			]
		`)
	})

	want := []*LicenseTemplate{{
		Key:         "apache-2.0",
		Name:        "Apache License 2.0",
		Nickname:    "",
		Featured:    true,
		HTMLURL:     "http://choosealicense.com/licenses/apache-2.0/",
		SourceURL:   "http://www.apache.org/licenses/LICENSE-2.0.html",
		Description: "A permissive license that also provides an express grant of patent rights from contributors to users.",
		Conditions:  []string{"include-copyright", "document-changes"},
		Permissions: []string{"commercial-use", "modifications", "distribution", "patent-use", "private-use"},
		Limitations: []string{"trademark-use", "no-liability"},
		Content:     "Apache License\n Version 2.0, January 2004\n [...]",
	}}

	lts, resp, err := client.LicenseTemplates.ListLicenseTemplates(nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, lts)

	lts, resp, err = client.LicenseTemplates.ListLicenseTemplates(nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, lts)
}

func TestLicenseTemplatesService_ListLicenseTemplates_StatusNotFound(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/templates/licenses", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusNotFound)
	})

	lts, resp, err := client.LicenseTemplates.ListLicenseTemplates(nil)
	require.Error(t, err)
	require.Nil(t, lts)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestLicenseTemplatesService_GetLicenseTemplate(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/templates/licenses/apache-2.0", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
				"key": "apache-2.0",
				"name": "Apache License 2.0",
				"nickname": null,
				"featured": true,
				"html_url": "http://choosealicense.com/licenses/apache-2.0/",
				"source_url": "http://www.apache.org/licenses/LICENSE-2.0.html",
				"description": "A permissive license that also provides an express grant of patent rights from contributors to users.",
				"conditions": [
					"include-copyright",
					"document-changes"
				],
				"permissions": [
					"commercial-use",
					"modifications",
					"distribution",
					"patent-use",
					"private-use"
				],
				"limitations": [
					"trademark-use",
					"no-liability"
				],
				"content": "Apache License\n Version 2.0, January 2004\n [...]"
			}
		`)
	})

	want := &LicenseTemplate{
		Key:         "apache-2.0",
		Name:        "Apache License 2.0",
		Nickname:    "",
		Featured:    true,
		HTMLURL:     "http://choosealicense.com/licenses/apache-2.0/",
		SourceURL:   "http://www.apache.org/licenses/LICENSE-2.0.html",
		Description: "A permissive license that also provides an express grant of patent rights from contributors to users.",
		Conditions:  []string{"include-copyright", "document-changes"},
		Permissions: []string{"commercial-use", "modifications", "distribution", "patent-use", "private-use"},
		Limitations: []string{"trademark-use", "no-liability"},
		Content:     "Apache License\n Version 2.0, January 2004\n [...]",
	}

	lt, resp, err := client.LicenseTemplates.GetLicenseTemplate("apache-2.0", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, lt)

	lt, resp, err = client.LicenseTemplates.GetLicenseTemplate("apache-2.0", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, lt)

	lt, resp, err = client.LicenseTemplates.GetLicenseTemplate("mit", nil)
	require.Error(t, err)
	require.Nil(t, lt)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
