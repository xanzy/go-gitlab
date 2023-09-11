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

func repositoryArchiveExample() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	// Get repository archive
	opt := &gitlab.ArchiveOptions{
		Format: gitlab.String("tar.gz"),
		// Gitlab API supports archiving first level directories at the moment, e.g. "mygroup/myproject/mydir".
		// If you pass 2+ level directory in path, whole repository will be archived and downloaded.
		Path: gitlab.String("mydir"),
	}
	content, _, err := git.Repositories.Archive("mygroup/myproject", opt, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("File contains %d byte(s)", content)
}
