package main

import (
	"log"

	"github.com/prytoegrian/go-gitlab"
)

func templateExample() {
	git := gitlab.NewClient(nil, "86A4YVLhj4EhczYsxP_s")
	git.SetBaseURL("http://localhost:80/api/v4")

	// List all ci templates
	t, _, err := git.CIYMLTemplate.ListAllTemplates(&gitlab.ListCIYMLTemplatesOptions{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found %d templates :", len(t))
	for _, template := range t {
		log.Printf("Found template : %v", template)
	}
}
