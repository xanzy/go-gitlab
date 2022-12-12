package gitlab

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProjectImportExportService_ScheduleExport(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/export", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusAccepted)
	})

	resp, err := client.ProjectImportExport.ScheduleExport(1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.ProjectImportExport.ScheduleExport(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.ProjectImportExport.ScheduleExport(1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.ProjectImportExport.ScheduleExport(2, nil, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectImportExportService_ExportStatus(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/export", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
			  "id": 1,
			  "description": "Itaque perspiciatis minima aspernatur corporis consequatur.",
			  "name": "Gitlab Test",
			  "name_with_namespace": "Gitlab Org / Gitlab Test",
			  "path": "gitlab-test",
			  "path_with_namespace": "gitlab-org/gitlab-test",
			  "export_status": "finished",
			  "_links": {
				"api_url": "https://gitlab.example.com/api/v4/projects/1/export/download",
				"web_url": "https://gitlab.example.com/gitlab-org/gitlab-test/download_export"
			  }
			}
		`)
	})

	want := &ExportStatus{
		ID:                1,
		Description:       "Itaque perspiciatis minima aspernatur corporis consequatur.",
		Name:              "Gitlab Test",
		NameWithNamespace: "Gitlab Org / Gitlab Test",
		Path:              "gitlab-test",
		PathWithNamespace: "gitlab-org/gitlab-test",
		ExportStatus:      "finished",
		Message:           "",
		Links: struct {
			APIURL string `json:"api_url"`
			WebURL string `json:"web_url"`
		}{
			APIURL: "https://gitlab.example.com/api/v4/projects/1/export/download",
			WebURL: "https://gitlab.example.com/gitlab-org/gitlab-test/download_export",
		},
	}

	es, resp, err := client.ProjectImportExport.ExportStatus(1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, es)

	es, resp, err = client.ProjectImportExport.ExportStatus(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, es)

	es, resp, err = client.ProjectImportExport.ExportStatus(1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, es)

	es, resp, err = client.ProjectImportExport.ExportStatus(2, nil, nil)
	require.Error(t, err)
	require.Nil(t, es)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectImportExportService_ExportDownload(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/export/download", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, "file.tar.gz")
	})

	want := []byte("file.tar.gz")

	es, resp, err := client.ProjectImportExport.ExportDownload(1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, es)

	es, resp, err = client.ProjectImportExport.ExportDownload(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, es)

	es, resp, err = client.ProjectImportExport.ExportDownload(1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, es)

	es, resp, err = client.ProjectImportExport.ExportDownload(2, nil, nil)
	require.Error(t, err)
	require.Nil(t, es)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestProjectImportExportService_ImportFile(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/import", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
			  "id": 1,
			  "description": null,
			  "name": "api-project",
			  "name_with_namespace": "Administrator / api-project",
			  "path": "api-project",
			  "path_with_namespace": "root/api-project",
			  "import_status": "scheduled",
			  "correlation_id": "mezklWso3Za"
			}
		`)
	})

	want := &ImportStatus{
		ID:                1,
		Description:       "",
		Name:              "api-project",
		NameWithNamespace: "Administrator / api-project",
		Path:              "api-project",
		PathWithNamespace: "root/api-project",
		ImportStatus:      "scheduled",
		CorrelationID:     "mezklWso3Za",
	}

	file := bytes.NewBufferString("dummy")
	es, resp, err := client.ProjectImportExport.ImportFromFile(file, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, es)

	es, resp, err = client.ProjectImportExport.ImportFromFile(file, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, es)
}

func TestProjectImportExportService_ImportStatus(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/import", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
			  "id": 1,
			  "description": "Itaque perspiciatis minima aspernatur corporis consequatur.",
			  "name": "Gitlab Test",
			  "name_with_namespace": "Gitlab Org / Gitlab Test",
			  "path": "gitlab-test",
			  "path_with_namespace": "gitlab-org/gitlab-test",
			  "import_status": "started",
			  "correlation_id": "mezklWso3Za"
			}
		`)
	})

	want := &ImportStatus{
		ID:                1,
		Description:       "Itaque perspiciatis minima aspernatur corporis consequatur.",
		Name:              "Gitlab Test",
		NameWithNamespace: "Gitlab Org / Gitlab Test",
		Path:              "gitlab-test",
		PathWithNamespace: "gitlab-org/gitlab-test",
		ImportStatus:      "started",
		CorrelationID:     "mezklWso3Za",
	}

	es, resp, err := client.ProjectImportExport.ImportStatus(1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, es)

	es, resp, err = client.ProjectImportExport.ImportStatus(1.01, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, es)

	es, resp, err = client.ProjectImportExport.ImportStatus(1, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, es)

	es, resp, err = client.ProjectImportExport.ImportStatus(2, nil)
	require.Error(t, err)
	require.Nil(t, es)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
