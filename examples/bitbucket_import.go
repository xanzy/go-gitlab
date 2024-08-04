package main

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func bitbucketCloudExample() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	cloudOpt := &gitlab.ImportRepositoryFromBitbucketCloudOptions{
		BitbucketUsername:    gitlab.Ptr("username"),
		BitbucketAppPassword: gitlab.Ptr("password"),
		RepoPath:             gitlab.Ptr("some/repo"),
		TargetNamespace:      gitlab.Ptr("some-group"),
		NewName:              gitlab.Ptr("some-repo"),
	}
	cloudResp, _, err := git.Import.ImportRepositoryFromBitbucketCloud(cloudOpt)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(cloudResp.String())
}

func bitbucketServerExample() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	serverOpt := &gitlab.ImportRepositoryFromBitbucketServerOptions{
		BitbucketServerUrl:      gitlab.Ptr("https://bitbucket.example.com"),
		BitbucketServerUsername: gitlab.Ptr("username"),
		PersonalAccessToken:     gitlab.Ptr("access-token"),
		BitbucketServerProject:  gitlab.Ptr("some-project"),
		BitbucketServerRepo:     gitlab.Ptr("some-repo"),
		NewName:                 gitlab.Ptr("some-other-repo"),
		NewNamespace:            gitlab.Ptr("some-group"),
		TimeoutStrategy:         gitlab.Ptr("pessimistic"),
	}
	serverResp, _, err := git.Import.ImportRepositoryFromBitbucketServer(serverOpt)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(serverResp.String())
}
