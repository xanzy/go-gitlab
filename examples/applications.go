package main

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func applicationsExample() {
	git := gitlab.NewClient(nil, "yourtokengoeshere")
	git.SetBaseURL("https://gitlab.com/api/v4")
	// Create an application
	opts := gitlab.CreateApplicationOptions{
		Name:        "Travis",
		RedirectURI: "http://example.org",
		Scopes:      "api",
	}
	created, _, err := git.Applications.CreateApplication(&opts)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Last created application : %v", created)

	// List all applications
	applications, _, err := git.Applications.ListApplications()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found %d applications :", len(applications))
	for _, app := range applications {
		log.Printf("Found app : %v", app)
	}

	// Delete an application
	resp, err := git.Applications.DeleteApplication(created.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status code response : %d", resp.StatusCode)
}
