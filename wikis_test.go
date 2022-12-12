//
// Copyright 2021, Stany MARCEL
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListWikis(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/wikis", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			  "content": "Here is an instruction how to deploy this project.",
			  "format": "markdown",
			  "slug": "deploy",
			  "title": "deploy"
			},
			{
			  "content": "Our development process is described here.",
			  "format": "markdown",
			  "slug": "development",
			  "title": "development"
			},
			{
			  "content": "*  [Deploy](deploy)\n*  [Development](development)",
			  "format": "markdown",
			  "slug": "home",
			  "title": "home"
			}
		  ]`)
	})

	wikis, _, err := client.Wikis.ListWikis(1, &ListWikisOptions{WithContent: Bool(true)})
	if err != nil {
		t.Errorf("Wikis.ListWikis returned error: %v", err)
	}

	want := []*Wiki{
		{
			Content: "Here is an instruction how to deploy this project.",
			Format:  "markdown",
			Slug:    "deploy",
			Title:   "deploy",
		},
		{
			Content: "Our development process is described here.",
			Format:  "markdown",
			Slug:    "development",
			Title:   "development",
		},
		{
			Content: "*  [Deploy](deploy)\n*  [Development](development)",
			Format:  "markdown",
			Slug:    "home",
			Title:   "home",
		},
	}

	if !reflect.DeepEqual(want, wikis) {
		t.Errorf("Labels.CreateLabel returned %+v, want %+v", wikis, want)
	}
}

func TestGetWikiPage(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/wikis/home", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
			"content": "home page",
			"format": "markdown",
			"slug": "home",
			"title": "home",
			"encoding": "UTF-8"
		  }`)
	})

	wiki, _, err := client.Wikis.GetWikiPage(1, "home", &GetWikiPageOptions{})
	if err != nil {
		t.Errorf("Wiki.GetWikiPage returned error: %v", err)
	}

	want := &Wiki{
		Content:  "home page",
		Encoding: "UTF-8",
		Format:   "markdown",
		Slug:     "home",
		Title:    "home",
	}

	if !reflect.DeepEqual(want, wiki) {
		t.Errorf("Labels.CreateLabel returned %+v, want %+v", wiki, want)
	}
}

func TestCreateWikiPage(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/wikis", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `{
			"content": "Hello world",
			"format": "markdown",
			"slug": "Hello",
			"title": "Hello"
		  }`)
	})

	opt := &CreateWikiPageOptions{
		Content: String("Hello world"),
		Title:   String("Hello"),
		Format:  WikiFormat(WikiFormatMarkdown),
	}
	wiki, _, err := client.Wikis.CreateWikiPage(1, opt)
	if err != nil {
		t.Errorf("Wiki.CreateWikiPage returned error: %v", err)
	}

	want := &Wiki{
		Content: "Hello world",
		Format:  "markdown",
		Slug:    "Hello",
		Title:   "Hello",
	}

	if !reflect.DeepEqual(want, wiki) {
		t.Errorf("Wiki.CreateWikiPage returned %+v, want %+v", wiki, want)
	}
}

func TestEditWikiPage(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/wikis/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `{
			"content": "documentation",
			"format": "markdown",
			"slug": "Docs",
			"title": "Docs"
		  }`)
	})

	opt := &EditWikiPageOptions{
		Content: String("documentation"),
		Format:  WikiFormat(WikiFormatMarkdown),
		Title:   String("Docs"),
	}
	wiki, _, err := client.Wikis.EditWikiPage(1, "foo", opt)
	if err != nil {
		t.Errorf("Wiki.EditWikiPage returned error: %v", err)
	}

	want := &Wiki{
		Content: "documentation",
		Format:  "markdown",
		Slug:    "Docs",
		Title:   "Docs",
	}

	if !reflect.DeepEqual(want, wiki) {
		t.Errorf("Wiki.EditWikiPage returned %+v, want %+v", wiki, want)
	}
}

func TestDeleteWikiPage(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/wikis/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Wikis.DeleteWikiPage(1, "foo")
	if err != nil {
		t.Errorf("Wiki.DeleteWikiPage returned error: %v", err)
	}
}
