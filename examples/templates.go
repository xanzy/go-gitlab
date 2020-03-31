package main

import (
	"log"

	"github.com/prytoegrian/go-gitlab"
)

func templateExample() {
	git := gitlab.NewClient(nil, "86A4YVLhj4EhczYsxP_s")
	git.SetBaseURL("http://localhost:80/api/v4")

	// List all ci templates
	c := template.NewCITemplate(git)
	t, _, err := c.ListAllTemplates(&template.ListCIYMLTemplatesOptions{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found %d CI templates :", len(t))
	for _, template := range t {
		log.Printf("Found template : %v", template)
	}

	// List all gitignore templates
	gitignore := template.NewGitignoreTemplate(git)
	g, _, err := gitignore.ListTemplates(&template.ListTemplatesOptions{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found %d Gitignore templates :", len(g))
	for _, template := range t {
		log.Printf("Found template : %v", template)
	}
}
