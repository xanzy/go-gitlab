package main

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func pipelineTestReportExample() {
	git, err := gitlab.NewClient(gitlab.PrivateTokenAuth("yourtokengoeshere"))
	if err != nil {
		log.Fatal(err)
	}

	opt := &gitlab.ListProjectPipelinesOptions{Ref: gitlab.String("master")}
	projectID := 1234

	pipelines, _, err := git.Pipelines.ListProjectPipelines(projectID, opt)
	if err != nil {
		log.Fatal(err)
	}

	for _, pipeline := range pipelines {
		log.Printf("Found pipeline: %v", pipeline)

		report, _, err := git.Pipelines.GetPipelineTestReport(projectID, pipeline.ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Found test report: %v", report)

	}
}
