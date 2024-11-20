package gitlab

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateDependencyListExport(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/pipelines/1234/dependency_list_exports", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		mustWriteHTTPResponse(t, w, "testdata/create_dependency_list_export.json")
	})

	export, _, err := client.DependencyListExport.CreateDependencyListExport(1234)
	require.NoError(t, err)

	want := &DependencyListExport{
		ID:          5678,
		HasFinished: false,
		Self:        "http://gitlab.example.com/api/v4/dependency_list_exports/5678",
		Download:    "http://gitlab.example.com/api/v4/dependency_list_exports/5678/download",
	}
	require.Equal(t, want, export)
}

func TestGetDependencyListExport(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/dependency_list_exports/5678", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_dependency_list_export.json")
	})

	export, _, err := client.DependencyListExport.GetDependencyListExport(5678)
	require.NoError(t, err)

	want := &DependencyListExport{
		ID:          5678,
		HasFinished: true,
		Self:        "http://gitlab.example.com/api/v4/dependency_list_exports/5678",
		Download:    "http://gitlab.example.com/api/v4/dependency_list_exports/5678/download",
	}
	require.Equal(t, want, export)
}

func TestDownloadDependencyListExport(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/dependency_list_exports/5678/download", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/download_dependency_list_export.json")
	})

	sbom, _, err := client.DependencyListExport.DownloadDependencyListExport(5678)
	require.NoError(t, err)

	wantBytes, err := os.ReadFile("testdata/download_dependency_list_export.json")
	require.NoError(t, err)
	want := string(wantBytes)

	require.Equal(t, want, sbom)
}
