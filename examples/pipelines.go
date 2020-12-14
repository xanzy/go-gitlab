package main

import (
	"log"
	"time"

	"github.com/Fourcast/go-gitlab"
)

func pipelineExample() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	opt := &gitlab.ListProjectPipelinesOptions{
		Scope:         gitlab.String("branches"),
		Status:        gitlab.BuildState(gitlab.Running),
		Ref:           gitlab.String("master"),
		YamlErrors:    gitlab.Bool(true),
		Name:          gitlab.String("name"),
		Username:      gitlab.String("username"),
		UpdatedAfter:  gitlab.Time(time.Now().Add(-24 * 365 * time.Hour)),
		UpdatedBefore: gitlab.Time(time.Now().Add(-7 * 24 * time.Hour)),
		OrderBy:       gitlab.String("status"),
		Sort:          gitlab.String("asc"),
	}

	pipelines, _, err := git.Pipelines.ListProjectPipelines(2743054, opt)
	if err != nil {
		log.Fatal(err)
	}

	for _, pipeline := range pipelines {
		log.Printf("Found pipeline: %v", pipeline)
	}
}
