package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepositoriesService_ListTree(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/repository/tree", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": "a1e8f8d745cc87e3a9248358d9352bb7f9a0aeba",
				"name": "html",
				"type": "tree",
				"path": "files/html",
				"mode": "040000"
			  }
			]
		`)
	})

	want := []*TreeNode{{
		ID:   "a1e8f8d745cc87e3a9248358d9352bb7f9a0aeba",
		Name: "html",
		Type: "tree",
		Path: "files/html",
		Mode: "040000",
	},
	}

	tns, resp, err := client.Repositories.ListTree(1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, tns)

	tns, resp, err = client.Repositories.ListTree(1.01, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, tns)

	tns, resp, err = client.Repositories.ListTree(1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, tns)

	tns, resp, err = client.Repositories.ListTree(2, nil)
	require.Error(t, err)
	require.Nil(t, tns)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestRepositoriesService_Blob(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/repository/blobs/2dc6aa325a317eda67812f05600bdf0fcdc70ab0", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, "{"+
			"size: 100"+
			"content: content"+
			"}",
		)
	})

	want := []byte("{" +
		"size: 100" +
		"content: content" +
		"}")

	b, resp, err := client.Repositories.Blob(1, "2dc6aa325a317eda67812f05600bdf0fcdc70ab0", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, b)

	b, resp, err = client.Repositories.Blob(1.01, "2dc6aa325a317eda67812f05600bdf0fcdc70ab0", nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.Repositories.Blob(1, "2dc6aa325a317eda67812f05600bdf0fcdc70ab0", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.Repositories.Blob(2, "2dc6aa325a317eda67812f05600bdf0fcdc70ab0", nil)
	require.Error(t, err)
	require.Nil(t, b)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestRepositoriesService_RawBlobContent(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/repository/blobs/2dc6aa325a317eda67812f05600bdf0fcdc70ab0/raw", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, "{"+
			"size: 100"+
			"content: content"+
			"}",
		)
	})

	want := []byte("{" +
		"size: 100" +
		"content: content" +
		"}")

	b, resp, err := client.Repositories.RawBlobContent(1, "2dc6aa325a317eda67812f05600bdf0fcdc70ab0", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, b)

	b, resp, err = client.Repositories.RawBlobContent(1.01, "2dc6aa325a317eda67812f05600bdf0fcdc70ab0", nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.Repositories.RawBlobContent(1, "2dc6aa325a317eda67812f05600bdf0fcdc70ab0", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.Repositories.RawBlobContent(2, "2dc6aa325a317eda67812f05600bdf0fcdc70ab0", nil)
	require.Error(t, err)
	require.Nil(t, b)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestRepositoriesService_Archive(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/repository/archive.gz", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, "{"+
			"size: 100"+
			"content: content"+
			"}",
		)
	})

	opt := &ArchiveOptions{Format: String("gz")}
	want := []byte("{" +
		"size: 100" +
		"content: content" +
		"}")

	b, resp, err := client.Repositories.Archive(1, opt, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, b)

	b, resp, err = client.Repositories.Archive(1.01, opt, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.Repositories.Archive(1, opt, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.Repositories.Archive(2, opt, nil)
	require.Error(t, err)
	require.Nil(t, b)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestRepositoriesService_StreamArchive(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/repository/archive.gz", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
	})

	var w http.ResponseWriter
	opt := &ArchiveOptions{Format: String("gz")}

	resp, err := client.Repositories.StreamArchive(1, w, opt, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.Repositories.StreamArchive(1.01, w, opt, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.Repositories.StreamArchive(1, w, opt, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.Repositories.StreamArchive(2, w, opt, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
