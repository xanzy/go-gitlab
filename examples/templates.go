package main

import (
	"log"

	"github.com/xanzy/go-gitlab"
	"github.com/xanzy/go-gitlab/templates"
)

func templateExample() {
	git := gitlab.NewClient(nil, "86A4YVLhj4EhczYsxP_s")
	git.SetBaseURL("http://localhost:80/api/v4")

	// List all ci templates
	ci := templates.NewCITemplate(*git)
	c, _, err := ci.ListAllTemplates(&templates.ListCIYMLTemplatesOptions{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found %d CI templates :", len(c))
	for _, template := range c {
		log.Printf("Found template : %v", template.Name)
	}

	// List all gitignore templates
	gitignore := templates.NewGitignoreTemplate(*git)
	g, _, err := gitignore.ListTemplates(&templates.ListTemplatesOptions{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found %d Gitignore templates :", len(g))
	for _, template := range g {
		log.Printf("Found template : %s", template.Name)
	}

	// List all licenses templates
	license := templates.NewLicenseTemplate(*git)
	l, _, err := license.ListLicenseTemplates(&templates.ListLicenseTemplatesOptions{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found %d license templates :", len(l))
	for _, template := range l {
		log.Printf("Found template : %v", template.Name)
	}
}
