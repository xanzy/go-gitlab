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

func TestClusterAgentTokensService_ListClusterAgentTokens(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

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

	opt := &ListClusterAgentTokensOptions{}
	clusterAgentTokens, _, err := client.ClusterAgentTokens.ListClusterAgentTokens(20, 5, opt)
	if err != nil {
		t.Errorf("ClusterAgentTokens.ListClusterAgentTokens returned error: %v", err)
	}

	want := []*ClusterAgentToken{
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
		t.Errorf("ClusterAgentTokens.ListClusterAgentTokens returned %+v, want %+v", clusterAgentTokens, want)
	}
}

func TestClusterAgentTokensService_GetClusterAgentToken(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

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

	clusterAgentToken, _, err := client.ClusterAgentTokens.GetClusterAgentToken(20, 5, 1)
	if err != nil {
		t.Errorf("ClusterAgentTokens.GetClusterAgentToken returned error: %v", err)
	}

	want := &ClusterAgentToken{
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
		t.Errorf("ClusterAgentTokens.GetClusterAgentToken returned %+v, want %+v", clusterAgentToken, want)
	}
}

func TestClusterAgentTokensService_RegisterClusterAgentToken(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

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

	opt := &CreateClusterAgentTokenOptions{Name: String("abcd"), Description: String("Some token")}
	clusterAgentToken, _, err := client.ClusterAgentTokens.CreateClusterAgentToken(20, 5, opt)
	if err != nil {
		t.Errorf("ClusterAgentTokens.CreateClusterAgentToken returned error: %v", err)
	}

	want := &ClusterAgentToken{
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
		t.Errorf("ClusterAgentTokens.CreateClusterAgentToken returned %+v, want %+v", clusterAgentToken, want)
	}
}
