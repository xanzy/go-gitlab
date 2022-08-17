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

func applicationsExample() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	// Create an application
	opts := &gitlab.CreateApplicationOptions{
		Name:        gitlab.String("Travis"),
		RedirectURI: gitlab.String("http://example.org"),
		Scopes:      gitlab.String("api"),
	}
	created, _, err := git.Applications.CreateApplication(opts)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Last created application : %v", created)

	// List all applications
	applications, _, err := git.Applications.ListApplications(&gitlab.ListApplicationsOptions{})
	if err != nil {
		log.Fatal(err)
	}

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
