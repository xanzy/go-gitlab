//
// Copyright 2021, Sander van Harmelen
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
//

package main

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func projectExample() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	// Create new project
	p := &gitlab.CreateProjectOptions{
		Name:                 gitlab.String("My Project"),
		Description:          gitlab.String("Just a test project to play with"),
		MergeRequestsEnabled: gitlab.Bool(true),
		SnippetsEnabled:      gitlab.Bool(true),
		Visibility:           gitlab.Visibility(gitlab.PublicVisibility),
	}
	project, _, err := git.Projects.CreateProject(p)
	if err != nil {
		log.Fatal(err)
	}

	// Add a new snippet
	s := &gitlab.CreateProjectSnippetOptions{
		Title:      gitlab.String("Dummy Snippet"),
		FileName:   gitlab.String("snippet.go"),
		Content:    gitlab.String("package main...."),
		Visibility: gitlab.Visibility(gitlab.PublicVisibility),
	}
	_, _, err = git.ProjectSnippets.CreateSnippet(project.ID, s)
	if err != nil {
		log.Fatal(err)
	}

	// List all project snippets
	snippets, _, err := git.ProjectSnippets.ListSnippets(project.PathWithNamespace, &gitlab.ListProjectSnippetsOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, snippet := range snippets {
		log.Printf("Found snippet: %s", snippet.Title)
	}
}
