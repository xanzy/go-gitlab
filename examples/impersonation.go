package main

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func impersonationExample() {
	git := gitlab.NewClient(nil, "yourtokengoeshere")

	userID := 1

	//list impersonation token from an user
	tokens, _, err := git.Users.GetAllImpersonationTokens(userID, &gitlab.GetAllImpersonationTokensOptions{
		State: gitlab.String("active"),
	})
	if err != nil {
		panic(err)
	}

	for _, token := range tokens {
		log.Println(token.Token)
	}

	//create an impersonation token of an user
	token, _, err := git.Users.CreateImpersonationToken(userID, &gitlab.CreateImpersonationTokenOptions{
		Scopes: &[]string{"api"},
	})
	if err != nil {
		panic(err)
	}

	log.Println(token.Token)
}
