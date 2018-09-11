package main

import (
	"log"

	"github.com/topochan/go-gitlab"
)

func main() {
	git := gitlab.NewClient(nil, "yourtokengoeshere")
	git.SetBaseURL("https://gitlab.com/api/v4")

	languages, _, err := git.Projects.GetProjectLanguages("2743054")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found languages: %v", languages)
}
