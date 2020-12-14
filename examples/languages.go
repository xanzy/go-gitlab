package main

import (
	"log"

	"github.com/Fourcast/go-gitlab"
)

func languagesExample() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	languages, _, err := git.Projects.GetProjectLanguages("2743054")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found languages: %v", languages)
}
