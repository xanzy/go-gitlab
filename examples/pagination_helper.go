//
// Copyright 2021, Timo Furrer <tuxtimo@gmail.com>
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
	"fmt"
	"log"

	"github.com/xanzy/go-gitlab"
)

func main() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	options := &gitlab.ListInstanceVariablesOptions{
		Page:    1,
	}
	instanceVariables, err := gitlab.Collect(
		git.InstanceVariables.ListVariables,
		options,
	)
	fmt.Printf("%d\n", len(instanceVariables))

	projectOptions := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
		},
		Archived: gitlab.Bool(false),
	}
	projects, err := gitlab.Collect(
		git.Projects.ListProjects,
		projectOptions,
	)
	fmt.Printf("%d\n", len(projects))


	userProjectOptions := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
		},
		Archived: gitlab.Bool(false),
	}
	userProjects, err := gitlab.Collect(
		func(opt *gitlab.ListProjectsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Project, *gitlab.Response, error) {
			return git.Projects.ListUserProjects(1, opt, options...)
		},
		userProjectOptions,
	)
	fmt.Printf("%d\n", len(userProjects))
}
