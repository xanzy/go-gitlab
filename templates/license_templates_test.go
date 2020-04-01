// +build integration

package templates

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListLicenseTemplatesKO(t *testing.T) {
	t.Parallel()

	m, s, c := gitlabtest.Setup()
	defer gitlabtest.Teardown(s)

	m.HandleFunc("/api/v4/templates/licenses", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"key":"1",    "name" : ".ping"}`)
	})

	licenses := NewLicenseTemplate(*c)
	_, _, err := licenses.ListLicenseTemplates(&ListLicenseTemplatesOptions{})
	if err == nil {
		t.Error("templates.ListLicenseTemplates doesn't catch parse error")
	}
}

func TestListLicenseTemplatesOK(t *testing.T) {
	t.Parallel()
	m, s, c := gitlabtest.Setup()
	defer gitlabtest.Teardown(s)

	m.HandleFunc("/api/v4/templates/licenses", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"key":"1", "name" : "testLicense"}]`)
	})

	licenses := NewLicenseTemplate(*c)
	templates, _, _ := licenses.ListLicenseTemplates(&ListLicenseTemplatesOptions{})
	want := []*LicenseTemplate{{Key: "1", Name: "testLicense"}}
	if !reflect.DeepEqual(templates, want) {
		t.Errorf("templates.ListLicenseTemplates returned %+v, want %+v", templates, want)
	}
}

func TestGetLicenseTemplatesKO(t *testing.T) {
	t.Parallel()
	m, s, c := gitlabtest.Setup()
	defer gitlabtest.Teardown(s)

	m.HandleFunc("/api/v4/templates/licenses/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"key":"1", "name" : "testLicense"`)
	})

	licenses := NewLicenseTemplate(*c)
	_, _, err := licenses.GetLicenseTemplate("test", &GetLicenseTemplateOptions{})
	if err == nil {
		t.Error("templates.GetLicenseTemplate doesn't catch parse error")
	}
}

func TestGetLicenseTemplatesOK(t *testing.T) {
	t.Parallel()
	m, s, c := gitlabtest.Setup()
	defer gitlabtest.Teardown(s)

	m.HandleFunc("/api/v4/templates/licenses/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"key":"1", "name" : "testLicense"}`)
	})

	licenses := NewLicenseTemplate(*c)
	template, _, _ := licenses.GetLicenseTemplate("test", &GetLicenseTemplateOptions{})
	want := &LicenseTemplate{Key: "1", Name: "testLicense"}
	if !reflect.DeepEqual(template, want) {
		t.Errorf("templates.GetLicenseTemplate returned %+v, want %+v", template, want)
	}
}
