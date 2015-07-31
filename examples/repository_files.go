package main

import (
	"encoding/base64"
	"log"

	"github.com/xanzy/go-gitlab"
)

func main() {
	git := gitlab.NewClient(nil, "yourtokengoeshere")

	// Create a new repository file
	cf := &gitlab.CreateFileOptions{
		FilePath:      "file.go",
		BranchName:    "master",
		Encoding:      "text",
		Content:       "My file contenxst",
		CommitMessage: "Adding a test file",
	}
	file, _, err := git.RepositoryFiles.CreateFile("myname/myproject", cf)
	if err != nil {
		log.Fatal(err)
	}

	// Update a repository file
	uf := &gitlab.UpdateFileOptions{
		FilePath:      file.FilePath,
		BranchName:    "master",
		Encoding:      "text",
		Content:       "My file content",
		CommitMessage: "Fixing typo",
	}
	_, _, err = git.RepositoryFiles.UpdateFile("myname/myproject", uf)
	if err != nil {
		log.Fatal(err)
	}

	gf := &gitlab.GetFileOptions{
		FilePath: file.FilePath,
		Ref:      "master",
	}
	f, _, err := git.RepositoryFiles.GetFile("myname/myproject", gf)
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
