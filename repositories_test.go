package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepositoriesService_ListTree(t *testing.T) {
	mux, client := setup(t)

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

	want := []*TreeNode{
		{
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
	mux, client := setup(t)

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
	mux, client := setup(t)

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
	mux, client := setup(t)

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
	mux, client := setup(t)

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

func TestRepositoriesService_Compare(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/12d65c8dd2b2676fa3ac47d955accc085a37a9c1/repository/compare", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
			  "commit": {
				"id": "12d65c8dd2b2676fa3ac47d955accc085a37a9c1",
				"short_id": "12d65c8dd2b",
				"title": "JS fix",
				"author_name": "Example User",
				"author_email": "user@example.com"
			  },
			  "commits": [{
				"id": "12d65c8dd2b2676fa3ac47d955accc085a37a9c1",
				"short_id": "12d65c8dd2b",
				"title": "JS fix",
				"author_name": "Example User",
				"author_email": "user@example.com"
			  }],
			  "diffs": [{
				"old_path": "files/js/application.js",
				"new_path": "files/js/application.js",
				"a_mode": null,
				"b_mode": "100644",
				"diff": "--- a/files/js/application.js\n+++ c/files/js/application.js\n@@ -24,8 +24,10 @@\n //= require g.raphael-min\n //= require g.bar-min\n //= require branch-graph\n-//= require highlightjs.min\n-//= require ace/ace\n //= require_tree .\n //= require d3\n //= require underscore\n+\n+function fix() { \n+  alert(\"Fixed\")\n+}",
				"new_file": false,
				"renamed_file": false,
				"deleted_file": false
			  }],
			  "compare_timeout": false,
			  "compare_same_ref": false,
			  "web_url": "https://gitlab.example.com/thedude/gitlab-foss/-/compare/ae73cb07c9eeaf35924a10f713b364d32b2dd34f...0b4bc9a49b562e85de7cc9e834518ea6828729b9"
			}
		`)
	})

	opt := &CompareOptions{
		From: String("master"),
		To:   String("feature"),
	}
	want := &Compare{
		Commit: &Commit{
			ID:          "12d65c8dd2b2676fa3ac47d955accc085a37a9c1",
			ShortID:     "12d65c8dd2b",
			Title:       "JS fix",
			AuthorName:  "Example User",
			AuthorEmail: "user@example.com",
		},
		Commits: []*Commit{{
			ID:          "12d65c8dd2b2676fa3ac47d955accc085a37a9c1",
			ShortID:     "12d65c8dd2b",
			Title:       "JS fix",
			AuthorName:  "Example User",
			AuthorEmail: "user@example.com",
		}},
		Diffs: []*Diff{{
			Diff:        "--- a/files/js/application.js\n+++ c/files/js/application.js\n@@ -24,8 +24,10 @@\n //= require g.raphael-min\n //= require g.bar-min\n //= require branch-graph\n-//= require highlightjs.min\n-//= require ace/ace\n //= require_tree .\n //= require d3\n //= require underscore\n+\n+function fix() { \n+  alert(\"Fixed\")\n+}",
			NewPath:     "files/js/application.js",
			OldPath:     "files/js/application.js",
			AMode:       "",
			BMode:       "100644",
			NewFile:     false,
			RenamedFile: false,
			DeletedFile: false,
		}},
		CompareTimeout: false,
		CompareSameRef: false,
	}

	c, resp, err := client.Repositories.Compare("12d65c8dd2b2676fa3ac47d955accc085a37a9c1", opt, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, c)

	c, resp, err = client.Repositories.Compare(1.01, opt, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, c)

	c, resp, err = client.Repositories.Compare("12d65c8dd2b2676fa3ac47d955accc085a37a9c1", opt, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, c)

	c, resp, err = client.Repositories.Compare("12d65c8dd2b2676fa3ac47d955accc085a37a9c2", opt, nil)
	require.Error(t, err)
	require.Nil(t, c)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestRepositoriesService_Contributors(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/12d65c8dd2b2676fa3ac47d955accc085a37a9c1/repository/contributors", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[{
			  "name": "Example User",
			  "email": "example@example.com",
			  "commits": 117,
			  "additions": 0,
			  "deletions": 0
			}]
		`)
	})

	want := []*Contributor{{
		Name:      "Example User",
		Email:     "example@example.com",
		Commits:   117,
		Additions: 0,
		Deletions: 0,
	}}

	cs, resp, err := client.Repositories.Contributors("12d65c8dd2b2676fa3ac47d955accc085a37a9c1", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, cs)

	cs, resp, err = client.Repositories.Contributors(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, cs)

	cs, resp, err = client.Repositories.Contributors("12d65c8dd2b2676fa3ac47d955accc085a37a9c1", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, cs)

	cs, resp, err = client.Repositories.Contributors("12d65c8dd2b2676fa3ac47d955accc085a37a9c2", nil, nil)
	require.Error(t, err)
	require.Nil(t, cs)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestRepositoriesService_MergeBase(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1a0b36b3cdad1d2ee32457c102a8c0b7056fa863/repository/merge_base", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
			  "id": "1a0b36b3cdad1d2ee32457c102a8c0b7056fa863",
			  "short_id": "1a0b36b3",
			  "title": "Initial commit",
			  "parent_ids": [],
			  "message": "Initial commit\n",
			  "author_name": "Example User",
			  "author_email": "user@example.com",
			  "committer_name": "Example User",
			  "committer_email": "user@example.com"
			}
		`)
	})

	want := &Commit{
		ID:             "1a0b36b3cdad1d2ee32457c102a8c0b7056fa863",
		ShortID:        "1a0b36b3",
		Title:          "Initial commit",
		AuthorName:     "Example User",
		AuthorEmail:    "user@example.com",
		CommitterName:  "Example User",
		CommitterEmail: "user@example.com",
		Message:        "Initial commit\n",
		ParentIDs:      []string{},
	}

	c, resp, err := client.Repositories.MergeBase("1a0b36b3cdad1d2ee32457c102a8c0b7056fa863", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, c)

	c, resp, err = client.Repositories.MergeBase(1.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, c)

	c, resp, err = client.Repositories.MergeBase("1a0b36b3cdad1d2ee32457c102a8c0b7056fa863", nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, c)

	c, resp, err = client.Repositories.MergeBase("1a0b36b3cdad1d2ee32457c102a8c0b7056fa865", nil, nil)
	require.Error(t, err)
	require.Nil(t, c)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestAddChangelogData(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/changelog",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			w.WriteHeader(http.StatusOK)
		})

	resp, err := client.Repositories.AddChangelog(
		1,
		&AddChangelogOptions{
			Version: String("1.0.0"),
		})
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGenerateChangelogData(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/changelog",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, exampleChangelogResponse)
		})

	want := &ChangelogData{
		Notes: "## 1.0.0 (2021-11-17)\n\n### feature (2 changes)\n\n- [Title 2](namespace13/project13@ad608eb642124f5b3944ac0ac772fecaf570a6bf) ([merge request](namespace13/project13!2))\n- [Title 1](namespace13/project13@3c6b80ff7034fa0d585314e1571cc780596ce3c8) ([merge request](namespace13/project13!1))\n",
	}

	notes, _, err := client.Repositories.GenerateChangelogData(
		1,
		GenerateChangelogDataOptions{
			Version: String("1.0.0"),
		},
	)
	require.NoError(t, err)
	assert.Equal(t, want, notes)
}
