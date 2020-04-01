// +build integration

package templates

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/prytoegrian/go-gitlab/gitlabtest"
)

func TestListGitignoreTemplatesKO(t *testing.T) {
	t.Parallel()

	m, s, c := gitlabtest.Setup()
	defer gitlabtest.Teardown(s)

	m.HandleFunc("/api/v4/templates/gitignores", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"name":"go", "content": "testGitignores"}`)
	})

	gitignores := NewGitignoreTemplate(*c)
	_, _, err := gitignores.ListTemplates(&ListTemplatesOptions{})
	if err == nil {
		t.Error("templates.ListTemplates doesn't catch parse error")
	}
}

func TestListGitignoreTemplatesOK(t *testing.T) {
	t.Parallel()
	m, s, c := gitlabtest.Setup()
	defer gitlabtest.Teardown(s)

	m.HandleFunc("/api/v4/templates/gitignores", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"name":"go", "content": "testGitignores"}]`)
	})

	gitignores := NewGitignoreTemplate(*c)
	templates, _, _ := gitignores.ListTemplates(&ListTemplatesOptions{})
	want := []*GitIgnoreTemplate{{Name: "go", Content: "testGitignores"}}
	if !reflect.DeepEqual(templates, want) {
		t.Errorf("templates.ListTemplates returned %+v, want %+v", templates, want)
	}
}

func TestGetGitignoreTemplateKO(t *testing.T) {
	t.Parallel()
	m, s, c := gitlabtest.Setup()
	defer gitlabtest.Teardown(s)

	m.HandleFunc("/api/v4/templates/gitignores/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"name":"go", "content": "testGitignores"`)
	})

	gitignores := NewGitignoreTemplate(*c)
	_, _, err := gitignores.GetTemplate("test")
	if err == nil {
		t.Error("templates.GetTemplate doesn't catch parse error")
	}
}

func TestGetGitignoreTemplateOK(t *testing.T) {
	t.Parallel()
	m, s, c := gitlabtest.Setup()
	defer gitlabtest.Teardown(s)

	m.HandleFunc("/api/v4/templates/gitignores/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"name":"go", "content": "testGitignores"}`)
	})

	gitignores := NewGitignoreTemplate(*c)
	template, _, _ := gitignores.GetTemplate("test")
	want := &GitIgnoreTemplate{Name: "go", Content: "testGitignores"}
	if !reflect.DeepEqual(template, want) {
		t.Errorf("templates.GetTemplate returned %+v, want %+v", template, want)
	}
}
