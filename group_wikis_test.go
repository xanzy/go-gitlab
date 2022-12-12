package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListGroupWikis(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/wikis",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `[
				{
					"content": "content here",
					"format": "markdown",
					"slug": "deploy",
					"title": "deploy title"
				}
			]`)
		})

	groupwikis, _, err := client.GroupWikis.ListGroupWikis(1, &ListGroupWikisOptions{})
	if err != nil {
		t.Errorf("GroupWikis.ListGroupWikis returned error: %v", err)
	}

	want := []*GroupWiki{
		{
			Content: "content here",
			Format:  WikiFormatMarkdown,
			Slug:    "deploy",
			Title:   "deploy title",
		},
	}
	if !reflect.DeepEqual(want, groupwikis) {
		t.Errorf("GroupWikis.ListGroupWikis returned %+v, want %+v", groupwikis, want)
	}
}

func TestGetGroupWikiPage(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/wikis/deploy",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `{
				"content": "content here",
				"format": "asciidoc",
				"slug": "deploy",
				"title": "deploy title",
				"encoding": "UTF-8"
			}`)
		})

	groupwiki, _, err := client.GroupWikis.GetGroupWikiPage(1, "deploy", &GetGroupWikiPageOptions{})
	if err != nil {
		t.Errorf("GroupWikis.GetGroupWikiPage returned error: %v", err)
	}

	want := &GroupWiki{
		Content:  "content here",
		Encoding: "UTF-8",
		Format:   WikiFormatASCIIDoc,
		Slug:     "deploy",
		Title:    "deploy title",
	}
	if !reflect.DeepEqual(want, groupwiki) {
		t.Errorf("GroupWikis.GetGroupWikiPage returned %+v, want %+v", groupwiki, want)
	}
}

func TestCreateGroupWikiPage(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/wikis",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `{
				"content": "content here",
				"format": "rdoc",
				"slug": "deploy",
				"title": "deploy title"
			}`)
		})

	groupwiki, _, err := client.GroupWikis.CreateGroupWikiPage(1, &CreateGroupWikiPageOptions{
		Content: String("content here"),
		Title:   String("deploy title"),
		Format:  WikiFormat(WikiFormatRDoc),
	})
	if err != nil {
		t.Errorf("GroupWikis.CreateGroupWikiPage returned error: %v", err)
	}

	want := &GroupWiki{
		Content: "content here",
		Format:  WikiFormatRDoc,
		Slug:    "deploy",
		Title:   "deploy title",
	}
	if !reflect.DeepEqual(want, groupwiki) {
		t.Errorf("GroupWikis.CreateGroupWikiPage returned %+v, want %+v", groupwiki, want)
	}
}

func TestEditGroupWikiPage(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/wikis/deploy",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			fmt.Fprint(w, `{
				"content": "content here",
				"format": "asciidoc",
				"slug": "deploy",
				"title": "deploy title"
			}`)
		})

	groupwiki, _, err := client.GroupWikis.EditGroupWikiPage(1, "deploy", &EditGroupWikiPageOptions{
		Content: String("content here"),
		Title:   String("deploy title"),
		Format:  WikiFormat(WikiFormatRDoc),
	})
	if err != nil {
		t.Errorf("GroupWikis.EditGroupWikiPage returned error: %v", err)
	}

	want := &GroupWiki{
		Content: "content here",
		Format:  WikiFormatASCIIDoc,
		Slug:    "deploy",
		Title:   "deploy title",
	}
	if !reflect.DeepEqual(want, groupwiki) {
		t.Errorf("GroupWikis.EditGroupWikiPage returned %+v, want %+v", groupwiki, want)
	}
}

func TestDeleteGroupWikiPage(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/wikis/deploy",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			w.WriteHeader(204)
		})

	r, err := client.GroupWikis.DeleteGroupWikiPage(1, "deploy")
	if err != nil {
		t.Errorf("GroupWikis.DeleteGroupWikiPage returned error: %v", err)
	}
	if r.StatusCode != 204 {
		t.Errorf("GroupWikis.DeleteGroupWikiPage returned wrong status code %d != 204", r.StatusCode)
	}
}
