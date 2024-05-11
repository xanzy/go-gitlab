package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestImportService_ImportRepositoryFromGithub(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/import/github", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
				"id": 27,
				"name": "my-repo",
				"full_path": "/root/my-repo",
				"full_name": "Administrator / my-repo",
				"refs_url": "/root/my-repo/refs",
				"import_source": "my-github/repo",
				"import_status": "scheduled",
				"human_import_status_name": "scheduled",
				"provider_link": "/my-github/repo",
				"relation_type": null,
				"import_warning": null
			}
		`)
	})

	want := &GithubImport{
		ID:                    27,
		Name:                  "my-repo",
		FullPath:              "/root/my-repo",
		FullName:              "Administrator / my-repo",
		RefsUrl:               "/root/my-repo/refs",
		ImportSource:          "my-github/repo",
		ImportStatus:          "scheduled",
		HumanImportStatusName: "scheduled",
		ProviderLink:          "/my-github/repo",
	}

	opt := &ImportRepositoryFromGithubOptions{
		PersonalAccessToken: Ptr("token"),
		RepoID:              Ptr(34),
		TargetNamespace:     Ptr("root"),
	}

	gi, _, err := client.Import.ImportRepositoryFromGithub(opt)
	if err != nil {
		t.Errorf("Import.ImportRepositoryFromGithub returned error: %v", err)
	}

	if !reflect.DeepEqual(want, gi) {
		t.Errorf("Import.ImportRepositoryFromGithub return %+v, want %+v", gi, want)
	}
}
