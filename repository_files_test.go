package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRepositoryFilesService_GetFile(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/13083/repository/files/app%2Fmodels%2Fkey%2Erb?ref=master", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
			  "file_name": "key.rb",
			  "file_path": "app/models/key.rb",
			  "size": 1476,
			  "encoding": "base64",
			  "content": "IyA9PSBTY2hlbWEgSW5mb3...",
			  "content_sha256": "4c294617b60715c1d218e61164a3abd4808a4284cbc30e6728a01ad9aada4481",
			  "execute_filemode": true,
			  "ref": "master",
			  "blob_id": "79f7bbd25901e8334750839545a9bd021f0e4c83",
			  "commit_id": "d5a3ff139356ce33e37e73add446f16869741b50",
			  "last_commit_id": "570e7b2abdd848b95f2f578043fc23bd6f6fd24d"
			}
		`)
	})

	want := &File{
		FileName:        "key.rb",
		FilePath:        "app/models/key.rb",
		Size:            1476,
		Encoding:        "base64",
		Content:         "IyA9PSBTY2hlbWEgSW5mb3...",
		ExecuteFilemode: true,
		Ref:             "master",
		BlobID:          "79f7bbd25901e8334750839545a9bd021f0e4c83",
		CommitID:        "d5a3ff139356ce33e37e73add446f16869741b50",
		SHA256:          "4c294617b60715c1d218e61164a3abd4808a4284cbc30e6728a01ad9aada4481",
		LastCommitID:    "570e7b2abdd848b95f2f578043fc23bd6f6fd24d",
	}

	f, resp, err := client.RepositoryFiles.GetFile(13083, "app%2Fmodels%2Fkey%2Erb?ref=master", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, f)

	f, resp, err = client.RepositoryFiles.GetFile(13083.01, "app%2Fmodels%2Fkey%2Erb?ref=master", nil)
	require.EqualError(t, err, "invalid ID type 13083.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, f)

	f, resp, err = client.RepositoryFiles.GetFile(13083, "app%2Fmodels%2Fkey%2Erb?ref=master", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, f)

	f, resp, err = client.RepositoryFiles.GetFile(13084, "app%2Fmodels%2Fkey%2Erb?ref=master", nil)
	require.Error(t, err)
	require.Nil(t, f)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestRepositoryFilesService_GetFileMetaData(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/13083/repository/files/app%2Fmodels%2Fkey%2Erb?ref=master", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodHead)
		w.Header().Set("X-Gitlab-Blob-Id", "79f7bbd25901e8334750839545a9bd021f0e4c83")
		w.Header().Set("X-Gitlab-Commit-Id", "d5a3ff139356ce33e37e73add446f16869741b50")
		w.Header().Set("X-Gitlab-Content-Sha256", "4c294617b60715c1d218e61164a3abd4808a4284cbc30e6728a01ad9aada4481")
		w.Header().Set("X-Gitlab-Encoding", "base64")
		w.Header().Set("X-Gitlab-File-Name", "key.rb")
		w.Header().Set("X-Gitlab-File-Path", "app/models/key.rb")
		w.Header().Set("X-Gitlab-Execute-Filemode", "true")
		w.Header().Set("X-Gitlab-Last-Commit-Id", "570e7b2abdd848b95f2f578043fc23bd6f6fd24d")
		w.Header().Set("X-Gitlab-Ref", "master")
		w.Header().Set("X-Gitlab-Size", "1476")
	})

	want := &File{
		FileName:        "key.rb",
		FilePath:        "app/models/key.rb",
		Size:            1476,
		Encoding:        "base64",
		ExecuteFilemode: true,
		Ref:             "master",
		BlobID:          "79f7bbd25901e8334750839545a9bd021f0e4c83",
		CommitID:        "d5a3ff139356ce33e37e73add446f16869741b50",
		SHA256:          "4c294617b60715c1d218e61164a3abd4808a4284cbc30e6728a01ad9aada4481",
		LastCommitID:    "570e7b2abdd848b95f2f578043fc23bd6f6fd24d",
	}

	f, resp, err := client.RepositoryFiles.GetFileMetaData(13083, "app%2Fmodels%2Fkey%2Erb?ref=master", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, f)

	f, resp, err = client.RepositoryFiles.GetFileMetaData(13083.01, "app%2Fmodels%2Fkey%2Erb?ref=master", nil)
	require.EqualError(t, err, "invalid ID type 13083.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, f)

	f, resp, err = client.RepositoryFiles.GetFileMetaData(13083, "app%2Fmodels%2Fkey%2Erb?ref=master", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, f)

	f, resp, err = client.RepositoryFiles.GetFileMetaData(13084, "app%2Fmodels%2Fkey%2Erb?ref=master", nil)
	require.Error(t, err)
	require.Nil(t, f)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestRepositoryFilesService_GetFileBlame(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/13083/repository/files/path%2Fto%2Ffile.rb/blame", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"commit": {
				  "id": "d42409d56517157c48bf3bd97d3f75974dde19fb",
				  "message": "Add feature also fix bug",
				  "parent_ids": [
					"cc6e14f9328fa6d7b5a0d3c30dc2002a3f2a3822"
				  ],
				  "author_name": "Venkatesh Thalluri",
				  "author_email": "venkatesh.thalluri@example.com",
				  "committer_name": "Venkatesh Thalluri",
				  "committer_email": "venkatesh.thalluri@example.com"
				},
				"lines": [
				  "require 'fileutils'",
				  "require 'open3'"
				]
			  }
			]
		`)
	})

	want := []*FileBlameRange{
		{
			Commit: struct {
				ID             string     `json:"id"`
				ParentIDs      []string   `json:"parent_ids"`
				Message        string     `json:"message"`
				AuthoredDate   *time.Time `json:"authored_date"`
				AuthorName     string     `json:"author_name"`
				AuthorEmail    string     `json:"author_email"`
				CommittedDate  *time.Time `json:"committed_date"`
				CommitterName  string     `json:"committer_name"`
				CommitterEmail string     `json:"committer_email"`
			}{
				ID:             "d42409d56517157c48bf3bd97d3f75974dde19fb",
				ParentIDs:      []string{"cc6e14f9328fa6d7b5a0d3c30dc2002a3f2a3822"},
				Message:        "Add feature also fix bug",
				AuthorName:     "Venkatesh Thalluri",
				AuthorEmail:    "venkatesh.thalluri@example.com",
				CommitterName:  "Venkatesh Thalluri",
				CommitterEmail: "venkatesh.thalluri@example.com",
			},
			Lines: []string{"require 'fileutils'", "require 'open3'"},
		},
	}

	fbr, resp, err := client.RepositoryFiles.GetFileBlame(13083, "path%2Fto%2Ffile.rb", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, fbr)

	fbr, resp, err = client.RepositoryFiles.GetFileBlame(13083.01, "path%2Fto%2Ffile.rb", nil)
	require.EqualError(t, err, "invalid ID type 13083.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, fbr)

	fbr, resp, err = client.RepositoryFiles.GetFileBlame(13083, "path%2Fto%2Ffile.rb", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, fbr)

	fbr, resp, err = client.RepositoryFiles.GetFileBlame(13084, "path%2Fto%2Ffile.rb", nil)
	require.Error(t, err)
	require.Nil(t, fbr)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestRepositoryFilesService_GetRawFile(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/13083/repository/files/app%2Fmodels%2Fkey%2Erb/raw", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, "HTTP/1.1 200 OK"+
			"..."+
			"X-Gitlab-Blob-Id: 79f7bbd25901e8334750839545a9bd021f0e4c83"+
			"X-Gitlab-Commit-Id: d5a3ff139356ce33e37e73add446f16869741b50"+
			"X-Gitlab-Content-Sha256: 4c294617b60715c1d218e61164a3abd4808a4284cbc30e6728a01ad9aada4481"+
			"X-Gitlab-Encoding: base64"+
			"X-Gitlab-File-Name: file.rb"+
			"X-Gitlab-File-Path: path/to/file.rb"+
			"X-Gitlab-Last-Commit-Id: 570e7b2abdd848b95f2f578043fc23bd6f6fd24d"+
			"X-Gitlab-Ref: master"+
			"X-Gitlab-Size: 1476"+
			"...")
	})

	want := []byte("HTTP/1.1 200 OK" +
		"..." +
		"X-Gitlab-Blob-Id: 79f7bbd25901e8334750839545a9bd021f0e4c83" +
		"X-Gitlab-Commit-Id: d5a3ff139356ce33e37e73add446f16869741b50" +
		"X-Gitlab-Content-Sha256: 4c294617b60715c1d218e61164a3abd4808a4284cbc30e6728a01ad9aada4481" +
		"X-Gitlab-Encoding: base64" +
		"X-Gitlab-File-Name: file.rb" +
		"X-Gitlab-File-Path: path/to/file.rb" +
		"X-Gitlab-Last-Commit-Id: 570e7b2abdd848b95f2f578043fc23bd6f6fd24d" +
		"X-Gitlab-Ref: master" +
		"X-Gitlab-Size: 1476" +
		"...",
	)

	b, resp, err := client.RepositoryFiles.GetRawFile(13083, "app%2Fmodels%2Fkey%2Erb", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, b)

	b, resp, err = client.RepositoryFiles.GetRawFile(13083.01, "app%2Fmodels%2Fkey%2Erb", nil)
	require.EqualError(t, err, "invalid ID type 13083.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.RepositoryFiles.GetRawFile(13083, "app%2Fmodels%2Fkey%2Erb", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, b)

	b, resp, err = client.RepositoryFiles.GetRawFile(13084, "app%2Fmodels%2Fkey%2Erb", nil)
	require.Error(t, err)
	require.Nil(t, b)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestRepositoryFilesService_CreateFile(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/13083/repository/files/app%2Fproject%2Erb", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
			  "file_path": "app/project.rb",
			  "branch": "master"
			}
		`)
	})

	want := &FileInfo{
		FilePath: "app/project.rb",
		Branch:   "master",
	}

	fi, resp, err := client.RepositoryFiles.CreateFile(13083, "app%2Fproject%2Erb", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, fi)

	fi, resp, err = client.RepositoryFiles.CreateFile(13083, "app%2Fproject%2Erb", &CreateFileOptions{ExecuteFilemode: Bool(true)})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, fi)

	fi, resp, err = client.RepositoryFiles.CreateFile(13083.01, "app%2Fproject%2Erb", nil)
	require.EqualError(t, err, "invalid ID type 13083.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, fi)

	fi, resp, err = client.RepositoryFiles.CreateFile(13083, "app%2Fproject%2Erb", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, fi)

	fi, resp, err = client.RepositoryFiles.CreateFile(13084, "app%2Fproject%2Erb", nil)
	require.Error(t, err)
	require.Nil(t, fi)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestRepositoryFilesService_UpdateFile(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/13083/repository/files/app%2Fproject%2Erb", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
			{
			  "file_path": "app/project.rb",
			  "branch": "master"
			}
		`)
	})

	want := &FileInfo{
		FilePath: "app/project.rb",
		Branch:   "master",
	}

	fi, resp, err := client.RepositoryFiles.UpdateFile(13083, "app%2Fproject%2Erb", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, fi)

	fi, resp, err = client.RepositoryFiles.UpdateFile(13083, "app%2Fproject%2Erb", &UpdateFileOptions{ExecuteFilemode: Bool(true)})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, fi)

	fi, resp, err = client.RepositoryFiles.UpdateFile(13083.01, "app%2Fproject%2Erb", nil)
	require.EqualError(t, err, "invalid ID type 13083.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, fi)

	fi, resp, err = client.RepositoryFiles.UpdateFile(13083, "app%2Fproject%2Erb", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, fi)

	fi, resp, err = client.RepositoryFiles.UpdateFile(13084, "app%2Fproject%2Erb", nil)
	require.Error(t, err)
	require.Nil(t, fi)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestRepositoryFilesService_DeleteFile(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/13083/repository/files/app%2Fproject%2Erb", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.RepositoryFiles.DeleteFile(13083, "app%2Fproject%2Erb", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.RepositoryFiles.DeleteFile(13083.01, "app%2Fproject%2Erb", nil)
	require.EqualError(t, err, "invalid ID type 13083.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.RepositoryFiles.DeleteFile(13083, "app%2Fproject%2Erb", nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.RepositoryFiles.DeleteFile(13084, "app%2Fproject%2Erb", nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
