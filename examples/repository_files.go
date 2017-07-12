package main

import (
	"encoding/base64"
	"log"

	"github.com/xanzy/go-gitlab"
)

func repositoryFileExample() {
	git := gitlab.NewClient(nil, "yourtokengoeshere")

	// Create a new repository file
	cf := &gitlab.CreateFileOptions{
		FilePath:      gitlab.String("file.go"),
		BranchName:    gitlab.String("master"),
		Encoding:      gitlab.String("text"),
		Content:       gitlab.String("My file contents"),
		CommitMessage: gitlab.String("Adding a test file"),
	}
	file, _, err := git.RepositoryFiles.CreateFile("myname/myproject", cf)
	if err != nil {
		log.Fatal(err)
	}

	// Update a repository file
	uf := &gitlab.UpdateFileOptions{
		FilePath:      gitlab.String(file.FilePath),
		BranchName:    gitlab.String("master"),
		Encoding:      gitlab.String("text"),
		Content:       gitlab.String("My file content"),
		CommitMessage: gitlab.String("Fixing typo"),
	}
	_, _, err = git.RepositoryFiles.UpdateFile("myname/myproject", uf)
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

	if f.Encoding == "base64" {
		content, err := base64.StdEncoding.DecodeString(f.Content)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("File contains: %s", string(content))

	} else {
		log.Printf("File contains: %s", f.Content)
	}
}
