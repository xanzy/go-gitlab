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

func ListClusterAgents(t *testing.T) {
	mux, client := setup(t)

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

	opt := &ListAgentsOptions{}
	clusterAgents, _, err := client.ClusterAgents.ListAgents(20, opt)
	if err != nil {
		t.Errorf("ClusterAgents.ListClusterAgents returned error: %v", err)
	}

	want := []*Agent{
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

func GetClusterAgent(t *testing.T) {
	mux, client := setup(t)

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

	clusterAgent, _, err := client.ClusterAgents.GetAgent(20, 1)
	if err != nil {
		t.Errorf("ClusterAgents.GetClusterAgent returned error: %v", err)
	}

	want := &Agent{
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

func RegisterClusterAgent(t *testing.T) {
	mux, client := setup(t)

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

	opt := &RegisterAgentOptions{Name: String("agent-1")}
	clusterAgent, _, err := client.ClusterAgents.RegisterAgent(20, opt)
	if err != nil {
		t.Errorf("ClusterAgents.RegisterClusterAgent returned error: %v", err)
	}

	want := &Agent{
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

func ListAgentTokens(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/20/cluster_agents/5/tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
		[
			{
			  "id": 1,
			  "name": "abcd",
			  "description": "Some token",
			  "agent_id": 5,
			  "status": "active",
			  "created_at": "2022-03-25T14:12:11.497Z",
			  "created_by_user_id": 1
			},
			{
			  "id": 2,
			  "name": "foobar",
			  "description": null,
			  "agent_id": 5,
			  "status": "active",
			  "created_at": "2022-03-25T14:12:11.497Z",
			  "created_by_user_id": 1
			}
		]
		`)
	})

	opt := &ListAgentTokensOptions{}
	clusterAgentTokens, _, err := client.ClusterAgents.ListAgentTokens(20, 5, opt)
	if err != nil {
		t.Errorf("ClusterAgents.ListAgentTokens returned error: %v", err)
	}

	want := []*AgentToken{
		{
			ID:              1,
			Name:            "abcd",
			Description:     "Some token",
			AgentID:         5,
			Status:          "active",
			CreatedAt:       Time(time.Date(2022, time.March, 25, 14, 12, 11, 497000000, time.UTC)),
			CreatedByUserID: 1,
		},
		{
			ID:              2,
			Name:            "foobar",
			Description:     "",
			AgentID:         5,
			Status:          "active",
			CreatedAt:       Time(time.Date(2022, time.March, 25, 14, 12, 11, 497000000, time.UTC)),
			CreatedByUserID: 1,
		},
	}

	if !reflect.DeepEqual(want, clusterAgentTokens) {
		t.Errorf("ClusterAgents.ListAgentTokens returned %+v, want %+v", clusterAgentTokens, want)
	}
}

func GetAgentToken(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/20/cluster_agents/5/tokens/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
		{
			"id": 1,
			"name": "abcd",
			"description": "Some token",
			"agent_id": 5,
			"status": "active",
			"created_at": "2022-03-25T14:12:11.497Z",
			"created_by_user_id": 1,
			"last_used_at": null
		 }
    	`)
	})

	clusterAgentToken, _, err := client.ClusterAgents.GetAgentToken(20, 5, 1)
	if err != nil {
		t.Errorf("ClusterAgents.GetAgentToken returned error: %v", err)
	}

	want := &AgentToken{
		ID:              1,
		Name:            "abcd",
		Description:     "Some token",
		AgentID:         5,
		Status:          "active",
		CreatedAt:       Time(time.Date(2022, time.March, 25, 14, 12, 11, 497000000, time.UTC)),
		CreatedByUserID: 1,
		LastUsedAt:      nil,
	}
	if !reflect.DeepEqual(want, clusterAgentToken) {
		t.Errorf("ClusterAgents.GetAgentToken returned %+v, want %+v", clusterAgentToken, want)
	}
}

func RegisterAgentToken(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/20/cluster_agents/5/tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `
		{
			"id": 1,
			"name": "abcd",
			"description": "Some token",
			"agent_id": 5,
			"status": "active",
			"created_at": "2022-03-25T14:12:11.497Z",
			"created_by_user_id": 1,
			"last_used_at": null,
			"token": "qeY8UVRisx9y3Loxo1scLxFuRxYcgeX3sxsdrpP_fR3Loq4xyg"
		}
    	`)
	})

	opt := &CreateAgentTokenOptions{Name: String("abcd"), Description: String("Some token")}
	clusterAgentToken, _, err := client.ClusterAgents.CreateAgentToken(20, 5, opt)
	if err != nil {
		t.Errorf("ClusterAgents.CreateAgentToken returned error: %v", err)
	}

	want := &AgentToken{
		ID:              1,
		Name:            "abcd",
		Description:     "Some token",
		AgentID:         5,
		Status:          "active",
		CreatedAt:       Time(time.Date(2022, time.March, 25, 14, 12, 11, 497000000, time.UTC)),
		CreatedByUserID: 1,
		LastUsedAt:      nil,
		Token:           "qeY8UVRisx9y3Loxo1scLxFuRxYcgeX3sxsdrpP_fR3Loq4xyg",
	}
	if !reflect.DeepEqual(want, clusterAgentToken) {
		t.Errorf("ClusterAgents.CreateAgentToken returned %+v, want %+v", clusterAgentToken, want)
	}
}
