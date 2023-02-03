package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPackagesService_ListProjectPackages(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/3/packages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 3,
				"name": "Hello/0.1@mycompany/stable",
				"conan_package_name": "Hello",
				"version": "0.1",
				"package_type": "conan",
				"last_downloaded_at": "2023-01-04T20:00:00.000Z",
				"_links": {
				  "web_path": "/foo/bar/-/packages/3",
				  "delete_api_path": "https://gitlab.example.com/api/v4/projects/1/packages/3"
				},
				"tags": [
					{
						"id": 1,
						"package_id": 37,
						"name": "Some Label",
						"created_at": "2023-01-04T20:00:00.000Z",
						"updated_at": "2023-01-04T20:00:00.000Z"
					}
				]
			  }
			]
		`)
	})

	timestamp := time.Date(2023, 1, 4, 20, 0, 0, 0, time.UTC)
	want := []*Package{{
		ID:               3,
		Name:             "Hello/0.1@mycompany/stable",
		Version:          "0.1",
		PackageType:      "conan",
		LastDownloadedAt: &timestamp,
		Links: &PackageLinks{
			WebPath:       "/foo/bar/-/packages/3",
			DeleteAPIPath: "https://gitlab.example.com/api/v4/projects/1/packages/3",
		},
		Tags: []PackageTag{
			{
				ID:        1,
				PackageID: 37,
				Name:      "Some Label",
				CreatedAt: &timestamp,
				UpdatedAt: &timestamp,
			},
		},
	}}

	ps, resp, err := client.Packages.ListProjectPackages(3, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ps)

	ps, resp, err = client.Packages.ListProjectPackages(3.01, nil)
	require.EqualError(t, err, "invalid ID type 3.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ps)

	ps, resp, err = client.Packages.ListProjectPackages(3, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ps)

	ps, resp, err = client.Packages.ListProjectPackages(5, nil)
	require.Error(t, err)
	require.Nil(t, ps)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestPackagesService_ListPackageFiles(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/3/packages/4/package_files", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 25,
				"package_id": 4,
				"file_name": "my-app-1.5-20181107.152550-1.jar",
				"size": 2421,
				"file_md5": "58e6a45a629910c6ff99145a688971ac",
				"file_sha1": "ebd193463d3915d7e22219f52740056dfd26cbfe",
				"file_sha256": "a903393463d3915d7e22219f52740056dfd26cbfeff321b"
			  }
			]
		`)
	})

	want := []*PackageFile{{
		ID:        25,
		PackageID: 4,
		FileName:  "my-app-1.5-20181107.152550-1.jar",
		Size:      2421,
		FileMD5:   "58e6a45a629910c6ff99145a688971ac",
		FileSHA1:  "ebd193463d3915d7e22219f52740056dfd26cbfe",
	}}

	ps, resp, err := client.Packages.ListPackageFiles(3, 4, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ps)

	ps, resp, err = client.Packages.ListPackageFiles(3.01, 4, nil)
	require.EqualError(t, err, "invalid ID type 3.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ps)

	ps, resp, err = client.Packages.ListPackageFiles(3, 4, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ps)

	ps, resp, err = client.Packages.ListPackageFiles(5, 4, nil)
	require.Error(t, err)
	require.Nil(t, ps)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestPackagesService_DeleteProjectPackage(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/3/packages/4", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Packages.DeleteProjectPackage(3, 4, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.Packages.DeleteProjectPackage(3.01, 4, nil)
	require.EqualError(t, err, "invalid ID type 3.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.Packages.DeleteProjectPackage(3, 4, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.Packages.DeleteProjectPackage(5, 4, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
