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

func impersonationExample() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	uid := 1

	// list impersonation token from an user
	tokens, _, err := git.Users.GetAllImpersonationTokens(
		uid,
		&gitlab.GetAllImpersonationTokensOptions{State: gitlab.String("active")},
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, token := range tokens {
		log.Printf("Found token: %s", token.Token)
	}

	// create an impersonation token of an user
	token, _, err := git.Users.CreateImpersonationToken(
		uid,
		&gitlab.CreateImpersonationTokenOptions{Scopes: &[]string{"api"}},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Created token: %s", token.Token)
}
