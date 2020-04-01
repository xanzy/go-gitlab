// +build integration

package templates

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListAllCITemplatesKO(t *testing.T) {
	t.Parallel()

	m, s, c := gitlabtest.Setup()
	defer gitlabtest.Teardown(s)

	m.HandleFunc("/api/v4/templates/gitlab_ci_ymls", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"name":"1",    "content" : ".ping"}`)
	})

	ci := NewCITemplate(*c)

	_, _, err := ci.ListAllTemplates(&ListCIYMLTemplatesOptions{})
	if err == nil {
		t.Error("templates.ListAllTemplates doesn't catch parse error")
	}
}

func TestListAllCITemplatesOK(t *testing.T) {
	t.Parallel()
	m, s, c := gitlabtest.Setup()
	defer gitlabtest.Teardown(s)

	m.HandleFunc("/api/v4/templates/gitlab_ci_ymls", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"name":"1",    "content" : ".ping"}]`)
	})

	ci := NewCITemplate(*c)

	templates, _, _ := ci.ListAllTemplates(&ListCIYMLTemplatesOptions{})
	want := []*CIYMLTemplate{{Name: "1", Content: ".ping"}}
	if !reflect.DeepEqual(templates, want) {
		t.Errorf("templates.ListAllTemplates returned %+v, want %+v", templates, want)
	}
}

func TestGetCITemplateKO(t *testing.T) {
	t.Parallel()
	m, s, c := gitlabtest.Setup()
	defer gitlabtest.Teardown(s)

	m.HandleFunc("/api/v4/templates/gitlab_ci_ymls/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"name":"1",    "content" : ".ping"`)
	})

	ci := NewCITemplate(*c)

	_, _, err := ci.GetTemplate("test")
	if err == nil {
		t.Error("templates.GetTemplate doesn't catch parse error")
	}
}

func TestGetCITemplateOK(t *testing.T) {
	t.Parallel()
	m, s, c := gitlabtest.Setup()
	defer gitlabtest.Teardown(s)

	m.HandleFunc("/api/v4/templates/gitlab_ci_ymls/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"name":"1",    "content" : ".ping"}`)
	})

	ci := NewCITemplate(*c)

	template, _, _ := ci.GetTemplate("test")
	want := &CIYMLTemplate{Name: "1", Content: ".ping"}
	if !reflect.DeepEqual(template, want) {
		t.Errorf("templates.GetTemplate returned %+v, want %+v", template, want)
	}
}
