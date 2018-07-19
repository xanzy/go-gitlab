package main

import (
	"fmt"

	"github.com/xanzy/go-gitlab"
)

func pagination() {
	git := gitlab.NewClient(nil, "yourtokengoeshere")

	opt := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 10,
			Page:    1,
		},
	}

	_, resp, _ := git.Projects.ListProjects(opt) // Page 1
	for page := 2; page < resp.TotalPages; page++ {
		opt.ListOptions.Page = page
		projects, _, _ := git.Projects.ListProjects(opt)

		for _, p := range projects {
			fmt.Println(p.ID)
		}
	}
}
