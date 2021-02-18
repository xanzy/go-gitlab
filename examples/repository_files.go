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

func repositoryFileExample() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new repository file
	cf := &gitlab.CreateFileOptions{
		Branch:        gitlab.String("master"),
		Content:       gitlab.String("My file contents"),
		CommitMessage: gitlab.String("Adding a test file"),
	}
	file, _, err := git.RepositoryFiles.CreateFile("myname/myproject", "file.go", cf)
	if err != nil {
		log.Fatal(err)
	}

	// Update a repository file
	uf := &gitlab.UpdateFileOptions{
		Branch:        gitlab.String("master"),
		Content:       gitlab.String("My file content"),
		CommitMessage: gitlab.String("Fixing typo"),
	}
	_, _, err = git.RepositoryFiles.UpdateFile("myname/myproject", file.FilePath, uf)
	if err != nil {
		log.Fatal(err)
	}

	gf := &gitlab.GetFileOptions{
		Ref: gitlab.String("master"),
	}
	f, _, err := git.RepositoryFiles.GetFile("myname/myproject", file.FilePath, gf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("File contains: %s", f.Content)

	gfb := &gitlab.GetFileBlameOptions{
		Ref: gitlab.String("master"),
	}
	fb, _, err := git.RepositoryFiles.GetFileBlame("myname/myproject", file.FilePath, gfb)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found %d blame ranges", len(fb))
}
