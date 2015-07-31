package main

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func main() {
	git := gitlab.NewClient(nil, "yourtokengoeshere")

	// Create new project
	p := &gitlab.CreateProjectOptions{
		Name:                 "My Project",
		Description:          "Just a test project to play with",
		MergeRequestsEnabled: true,
		SnippetsEnabled:      true,
		VisibilityLevel:      gitlab.PublicVisibility,
	}
	project, _, err := git.Projects.CreateProject(p)
	if err != nil {
		log.Fatal(err)
	}

	// Add a new snippet
	s := &gitlab.CreateSnippetOptions{
		Title:           "Dummy Snippet",
		FileName:        "snippet.go",
		Code:            "package main....",
		VisibilityLevel: gitlab.PublicVisibility,
	}
	_, _, err = git.ProjectSnippets.CreateSnippet(project.ID, s)
	if err != nil {
		log.Fatal(err)
	}

	// List all project snippets
	snippets, _, err := git.ProjectSnippets.ListSnippits(project.PathWithNamespace)
	if err != nil {
		log.Fatal(err)
	}

	for _, snippet := range snippets {
		log.Printf("Found snippet: %s", snippet.Title)
	}
}
