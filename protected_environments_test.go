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

package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProtectedEnvironments(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/protected_environments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"name":"1.0.0", "deploy_access_levels": [{"access_level": 40, "access_level_description": "Maintainers"}]},{"name":"*-release", "deploy_access_levels": [{"access_level": 30, "access_level_description": "Developers + Maintainers"}]}]`)
	})

	expected := []*ProtectedEnvironment{
		{
			Name: "1.0.0",
			DeployAccessLevels: []*EnvironmentAccessDescription{
				{
					AccessLevel:            40,
					AccessLevelDescription: "Maintainers",
				},
			},
		},
		{
			Name: "*-release",
			DeployAccessLevels: []*EnvironmentAccessDescription{
				{
					AccessLevel:            30,
					AccessLevelDescription: "Developers + Maintainers",
				},
			},
		},
	}

	opt := &ListProtectedEnvironmentsOptions{}
	environments, _, err := client.ProtectedEnvironments.ListProtectedEnvironments(1, opt)
	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, environments)
}

func TestGetProtectedEnvironment(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	environmentName := "my-awesome-environment"

	mux.HandleFunc(fmt.Sprintf("/api/v4/projects/1/protected_environments/%s", environmentName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"name":"my-awesome-environment", "deploy_access_levels": [{"access_level": 30, "access_level_description": "Developers + Maintainers"}]}`)
	})

	expected := &ProtectedEnvironment{
		Name: environmentName,
		DeployAccessLevels: []*EnvironmentAccessDescription{
			{
				AccessLevel:            30,
				AccessLevelDescription: "Developers + Maintainers",
			},
		},
	}

	environment, _, err := client.ProtectedEnvironments.GetProtectedEnvironment(1, environmentName)

	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, environment)
}

func TestProtectRepositoryEnvironments(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/protected_environments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"name":"my-awesome-environment", "deploy_access_levels": [{"access_level": 30, "access_level_description": "Developers + Maintainers"}]}`)
	})

	expected := &ProtectedEnvironment{
		Name: "my-awesome-environment",
		DeployAccessLevels: []*EnvironmentAccessDescription{
			{
				AccessLevel:            30,
				AccessLevelDescription: "Developers + Maintainers",
			},
		},
	}

	opt := &ProtectRepositoryEnvironmentsOptions{Name: String("my-awesome-environment"), DeployAccessLevels: &[]*EnvironmentAccessOptions{{AccessLevel: AccessLevel(30)}}}
	environment, _, err := client.ProtectedEnvironments.ProtectRepositoryEnvironments(1, opt)

	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, environment)
}

func TestUnprotectRepositoryEnvironments(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/protected_environments/my-awesome-environment", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.ProtectedEnvironments.UnprotectEnvironment(1, "my-awesome-environment")
	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
