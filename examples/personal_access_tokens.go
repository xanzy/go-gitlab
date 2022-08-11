//
// Copyright 2022, Ryan Glab <ryan.j.glab@gmail.com>
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
	"encoding/json"
	"fmt"
	"log"

	"github.com/xanzy/go-gitlab"
)

func patRevokeExample() {
	git, err := gitlab.NewClient("glpat-123xyz")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := git.PersonalAccessTokens.RevokePersonalAccessToken(99999999)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response Code: %s", resp.Status)
}

func patListExampleWithUserFilter() {
	git, err := gitlab.NewClient("glpat-123xyz")
	if err != nil {
		log.Fatal(err)
	}

	opt := &gitlab.ListPersonalAccessTokensOptions{
		ListOptions: gitlab.ListOptions{Page: 1, PerPage: 10},
		UserID:      gitlab.Int(12345),
	}

	personalAccessTokens, _, err := git.PersonalAccessTokens.ListPersonalAccessTokens(opt)
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(personalAccessTokens)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found personal access tokens: %s", data)
}
