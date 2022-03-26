//
// Copyright 2022, Daniela Filipe Bento
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

	gitlab "github.com/xanzy/go-gitlab"
)

func deploymentExample() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	opts := &gitlab.ListMergeRequestsOptions{}

	merge_requests, _, err := git.DeploymentMergeRequests.ListDeploymentMergeRequests(1, 1, opts)
	if err != nil {
		log.Fatal(err)
	}

	for _, merge_request := range merge_requests {
		log.Printf("Found merge_request: %v", merge_request)
	}
}
