//
// Copyright 2022, Timo Furrer <tuxtimo@gmail.com>
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

package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestClusterAgentsService_ListClusterAgents(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/20/cluster_agents", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
		[
			{
			  "id": 1,
			  "name": "agent-1",
			  "config_project": {
				"id": 20,
				"description": "",
				"name": "test",
				"name_with_namespace": "Administrator / test",
				"path": "test",
				"path_with_namespace": "root/test",
				"created_at": "2022-03-20T20:42:40.221Z"
			  },
			  "created_at": "2022-04-20T20:42:40.221Z",
			  "created_by_user_id": 42
			},
			{
			  "id": 2,
			  "name": "agent-2",
			  "config_project": {
				"id": 20,
				"description": "",
				"name": "test",
				"name_with_namespace": "Administrator / test",
				"path": "test",
				"path_with_namespace": "root/test",
				"created_at": "2022-03-20T20:42:40.221Z"
			  },
			  "created_at": "2022-04-20T20:42:40.221Z",
			  "created_by_user_id": 42
			}
		  ]
		`)
	})

	opt := &ListClusterAgentsOptions{}
	clusterAgents, _, err := client.ClusterAgents.ListClusterAgents(20, opt)
	if err != nil {
		t.Errorf("ClusterAgents.ListClusterAgents returned error: %v", err)
	}

	want := []*ClusterAgent{
		{
			ID:   1,
			Name: "agent-1",
			ConfigProject: ConfigProject{
				ID:                20,
				Description:       "",
				Name:              "test",
				NameWithNamespace: "Administrator / test",
				Path:              "test",
				PathWithNamespace: "root/test",
				CreatedAt:         Time(time.Date(2022, time.March, 20, 20, 42, 40, 221000000, time.UTC)),
			},
			CreatedAt:       Time(time.Date(2022, time.April, 20, 20, 42, 40, 221000000, time.UTC)),
			CreatedByUserID: 42,
		},
		{
			ID:   2,
			Name: "agent-2",
			ConfigProject: ConfigProject{
				ID:                20,
				Description:       "",
				Name:              "test",
				NameWithNamespace: "Administrator / test",
				Path:              "test",
				PathWithNamespace: "root/test",
				CreatedAt:         Time(time.Date(2022, time.March, 20, 20, 42, 40, 221000000, time.UTC)),
			},
			CreatedAt:       Time(time.Date(2022, time.April, 20, 20, 42, 40, 221000000, time.UTC)),
			CreatedByUserID: 42,
		},
	}

	if !reflect.DeepEqual(want, clusterAgents) {
		t.Errorf("ClusterAgents.ListClusterAgents returned %+v, want %+v", clusterAgents, want)
	}
}

func TestClusterAgentsService_GetClusterAgent(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/20/cluster_agents/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
			{
				"id": 1,
				"name": "agent-1",
				"config_project": {
				"id": 20,
				"description": "",
				"name": "test",
				"name_with_namespace": "Administrator / test",
				"path": "test",
				"path_with_namespace": "root/test",
				"created_at": "2022-03-20T20:42:40.221Z"
				},
				"created_at": "2022-04-20T20:42:40.221Z",
				"created_by_user_id": 42
			}
    	`)
	})

	clusterAgent, _, err := client.ClusterAgents.GetClusterAgent(20, 1)
	if err != nil {
		t.Errorf("ClusterAgents.GetClusterAgent returned error: %v", err)
	}

	want := &ClusterAgent{
		ID:   1,
		Name: "agent-1",
		ConfigProject: ConfigProject{
			ID:                20,
			Description:       "",
			Name:              "test",
			NameWithNamespace: "Administrator / test",
			Path:              "test",
			PathWithNamespace: "root/test",
			CreatedAt:         Time(time.Date(2022, time.March, 20, 20, 42, 40, 221000000, time.UTC)),
		},
		CreatedAt:       Time(time.Date(2022, time.April, 20, 20, 42, 40, 221000000, time.UTC)),
		CreatedByUserID: 42,
	}
	if !reflect.DeepEqual(want, clusterAgent) {
		t.Errorf("ClusterAgents.GetClusterAgent returned %+v, want %+v", clusterAgent, want)
	}
}

func TestClusterAgentsService_RegisterClusterAgent(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/20/cluster_agents", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `
			{
				"id": 1,
				"name": "agent-1",
				"config_project": {
				  "id": 20,
				  "description": "",
				  "name": "test",
				  "name_with_namespace": "Administrator / test",
				  "path": "test",
				  "path_with_namespace": "root/test",
				  "created_at": "2022-03-20T20:42:40.221Z"
				},
				"created_at": "2022-04-20T20:42:40.221Z",
				"created_by_user_id": 42
			  }
    	`)
	})

	opt := &RegisterClusterAgentOptions{Name: String("agent-1")}
	clusterAgent, _, err := client.ClusterAgents.RegisterClusterAgent(20, opt)
	if err != nil {
		t.Errorf("ClusterAgents.RegisterClusterAgent returned error: %v", err)
	}

	want := &ClusterAgent{
		ID:   1,
		Name: "agent-1",
		ConfigProject: ConfigProject{
			ID:                20,
			Description:       "",
			Name:              "test",
			NameWithNamespace: "Administrator / test",
			Path:              "test",
			PathWithNamespace: "root/test",
			CreatedAt:         Time(time.Date(2022, time.March, 20, 20, 42, 40, 221000000, time.UTC)),
		},
		CreatedAt:       Time(time.Date(2022, time.April, 20, 20, 42, 40, 221000000, time.UTC)),
		CreatedByUserID: 42,
	}
	if !reflect.DeepEqual(want, clusterAgent) {
		t.Errorf("ClusterAgents.RegisterClusterAgent returned %+v, want %+v", clusterAgent, want)
	}
}
