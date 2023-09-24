//
// Copyright 2023, Joel Gerber
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

func dataDogExample() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	// Create new DataDog integration
	opts := &gitlab.SetDataDogServiceOptions{
		APIKey:             gitlab.String("testing"),
		ArchiveTraceEvents: gitlab.Bool(true),
		DataDogEnv:         gitlab.String("sandbox"),
		DataDogService:     gitlab.String("test"),
		DataDogSite:        gitlab.String("datadoghq.com"),
		DataDogTags:        gitlab.String("country:canada\nprovince:ontario"),
	}

	_, err = git.Services.SetDataDogService(1, opts)
	if err != nil {
		log.Fatal(err)
	}

	// Query the integration
	svc, _, err := git.Services.GetDataDogService(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		"api_url: %s, archive_trace_events: %v, datadog_env: %s, datadog_service: %s, datadog_site: %s, datadog_tags: %s",
		svc.Properties.APIURL, svc.Properties.ArchiveTraceEvents, svc.Properties.DataDogEnv,
		svc.Properties.DataDogService, svc.Properties.DataDogSite, svc.Properties.DataDogTags,
	)

	// Delete the integration
	_, err = git.Services.DeleteDataDogService(1)
	if err != nil {
		log.Fatal(err)
	}
}
