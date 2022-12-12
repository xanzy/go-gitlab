package gitlab

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestGroupScheduleExport(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/export",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `{"message": "202 Accepted"}`)
		})

	_, err := client.GroupImportExport.ScheduleExport(1)
	if err != nil {
		t.Errorf("GroupImportExport.ScheduleExport returned error: %v", err)
	}
}

func TestGroupExportDownload(t *testing.T) {
	mux, client := setup(t)
	content := []byte("fake content")

	mux.HandleFunc("/api/v4/groups/1/export/download",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.Write(content)
		})

	export, _, err := client.GroupImportExport.ExportDownload(1)
	if err != nil {
		t.Errorf("GroupImportExport.ExportDownload returned error: %v", err)
	}

	data, err := io.ReadAll(export)
	if err != nil {
		t.Errorf("Error reading export: %v", err)
	}

	want := []byte("fake content")
	if !reflect.DeepEqual(want, data) {
		t.Errorf("GroupImportExport.GroupExportDownload returned %+v, want %+v", data, want)
	}
}

func TestGroupImport(t *testing.T) {
	mux, client := setup(t)

	content := []byte("temporary file's content")
	tmpfile, err := os.CreateTemp(os.TempDir(), "example.*.tar.gz")
	if err != nil {
		tmpfile.Close()
		log.Fatal(err)
	}
	if _, err := tmpfile.Write(content); err != nil {
		tmpfile.Close()
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	mux.HandleFunc("/api/v4/groups/import",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `{"message": "202 Accepted"}`)
		})

	opt := &GroupImportFileOptions{
		Name:     String("test"),
		Path:     String("path"),
		File:     String(tmpfile.Name()),
		ParentID: Int(1),
	}

	_, err = client.GroupImportExport.ImportFile(opt)
	if err != nil {
		t.Errorf("GroupImportExport.ImportFile returned error: %v", err)
	}
}
