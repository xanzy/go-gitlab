//
// Copyright 2023, Sander van Harmelen
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

func TestGroupListProtectedEnvironments(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/protected_environments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{
      "name":"staging",
      "deploy_access_levels": [
        {
          "access_level": 40,
          "access_level_description": "Maintainers"
        }
      ],
      "required_approval_count": 1,
      "approval_rules": [
        {
           "id": 38,
           "user_id": 42,
           "group_id": null,
           "access_level": null,
           "access_level_description": "qa-group",
           "required_approvals": 1,
           "group_inheritance_type": 0
        },
        {
           "id": 39,
           "user_id": null,
           "group_id": 135,
           "access_level": 30,
           "access_level_description": "security-group",
           "required_approvals": 2,
           "group_inheritance_type": 1
        }
      ]
    },{
      "name":"production",
      "deploy_access_levels": [
        {
          "access_level": 30,
          "access_level_description": "Developers + Maintainers"
        }
      ]
    }]`)
	})

	expected := []*GroupProtectedEnvironment{
		{
			Name: "staging",
			DeployAccessLevels: []*GroupEnvironmentAccessDescription{
				{
					AccessLevel:            40,
					AccessLevelDescription: "Maintainers",
				},
			},
			RequiredApprovalCount: 1,
			ApprovalRules: []*GroupEnvironmentApprovalRule{
				{
					ID:                     38,
					UserID:                 42,
					AccessLevelDescription: "qa-group",
					RequiredApprovalCount:  1,
				},
				{
					ID:                     39,
					GroupID:                135,
					AccessLevel:            30,
					AccessLevelDescription: "security-group",
					RequiredApprovalCount:  2,
					GroupInheritanceType:   1,
				},
			},
		},
		{
			Name: "production",
			DeployAccessLevels: []*GroupEnvironmentAccessDescription{
				{
					AccessLevel:            30,
					AccessLevelDescription: "Developers + Maintainers",
				},
			},
		},
	}

	opt := &ListGroupProtectedEnvironmentsOptions{}
	environments, _, err := client.GroupProtectedEnvironments.ListGroupProtectedEnvironments(1, opt)
	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, environments)
}

func TestGroupGetProtectedEnvironment(t *testing.T) {
	mux, client := setup(t)

	// Test with RequiredApprovalCount
	environmentName := "development"

	mux.HandleFunc(fmt.Sprintf("/api/v4/groups/1/protected_environments/%s", environmentName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
      "name":"%s",
      "deploy_access_levels": [
        {
          "access_level": 30,
          "access_level_description": "Developers + Maintainers"
        }
      ],
      "required_approval_count": 1,
      "approval_rules": [
        {
           "id": 1,
           "user_id": null,
           "group_id": 10,
           "access_level": 5,
           "access_level_description": "devops",
           "required_approvals": 0,
           "group_inheritance_type": 0
        }
      ]
    }`, environmentName)
	})

	expected := &GroupProtectedEnvironment{
		Name: environmentName,
		DeployAccessLevels: []*GroupEnvironmentAccessDescription{
			{
				AccessLevel:            30,
				AccessLevelDescription: "Developers + Maintainers",
			},
		},
		RequiredApprovalCount: 1,
		ApprovalRules: []*GroupEnvironmentApprovalRule{
			{
				ID:                     1,
				GroupID:                10,
				AccessLevel:            5,
				AccessLevelDescription: "devops",
			},
		},
	}

	environment, _, err := client.GroupProtectedEnvironments.GetGroupProtectedEnvironment(1, environmentName)
	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, environment)

	// Test without RequiredApprovalCount nor ApprovalRules
	environmentName = "testing"

	mux.HandleFunc(fmt.Sprintf("/api/v4/groups/2/protected_environments/%s", environmentName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
      "name":"%s",
      "deploy_access_levels": [
        {
          "access_level": 30,
          "access_level_description": "Developers + Maintainers"
        }
      ]
    }`, environmentName)
	})

	expected = &GroupProtectedEnvironment{
		Name: environmentName,
		DeployAccessLevels: []*GroupEnvironmentAccessDescription{
			{
				AccessLevel:            30,
				AccessLevelDescription: "Developers + Maintainers",
			},
		},
	}

	environment, _, err = client.GroupProtectedEnvironments.GetGroupProtectedEnvironment(2, environmentName)
	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, environment)
}

func TestGroupProtectEnvironments(t *testing.T) {
	mux, client := setup(t)

	// Test with RequiredApprovalCount and ApprovalRules
	environmentName := "other"

	mux.HandleFunc("/api/v4/groups/1/protected_environments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `{
      "name":"%s",
      "deploy_access_levels": [
        {
          "access_level": 30,
          "access_level_description": "Developers + Maintainers"
        }
      ],
      "required_approval_count": 2,
      "approval_rules": [
        {
           "id": 1,
           "user_id": null,
           "group_id": 10,
           "access_level": 5,
           "access_level_description": "devops",
           "required_approvals": 0,
           "group_inheritance_type": 0
        }
      ]
    }`, environmentName)
	})

	expected := &GroupProtectedEnvironment{
		Name: environmentName,
		DeployAccessLevels: []*GroupEnvironmentAccessDescription{
			{
				AccessLevel:            30,
				AccessLevelDescription: "Developers + Maintainers",
			},
		},
		RequiredApprovalCount: 2,
		ApprovalRules: []*GroupEnvironmentApprovalRule{
			{
				ID:                     1,
				GroupID:                10,
				AccessLevel:            5,
				AccessLevelDescription: "devops",
			},
		},
	}

	opt := &ProtectGroupEnvironmentOptions{
		Name: String(environmentName),
		DeployAccessLevels: &[]*GroupEnvironmentAccessOptions{
			{AccessLevel: AccessLevel(30)},
		},
		RequiredApprovalCount: Int(2),
		ApprovalRules: &[]*GroupEnvironmentApprovalRuleOptions{
			{
				GroupID:                Int(10),
				AccessLevel:            AccessLevel(0),
				AccessLevelDescription: String("devops"),
			},
		},
	}

	environment, _, err := client.GroupProtectedEnvironments.ProtectGroupEnvironment(1, opt)
	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, environment)

	// Test without RequiredApprovalCount nor ApprovalRules
	environmentName = "staging"

	mux.HandleFunc("/api/v4/groups/2/protected_environments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `{
      "name":"%s",
      "deploy_access_levels": [
        {
          "access_level": 30,
          "access_level_description": "Developers + Maintainers"
        }
      ]
    }`, environmentName)
	})

	expected = &GroupProtectedEnvironment{
		Name: environmentName,
		DeployAccessLevels: []*GroupEnvironmentAccessDescription{
			{
				AccessLevel:            30,
				AccessLevelDescription: "Developers + Maintainers",
			},
		},
	}

	opt = &ProtectGroupEnvironmentOptions{
		Name: String(environmentName),
		DeployAccessLevels: &[]*GroupEnvironmentAccessOptions{
			{AccessLevel: AccessLevel(30)},
		},
	}
	environment, _, err = client.GroupProtectedEnvironments.ProtectGroupEnvironment(2, opt)
	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, environment)
}

func TestGroupUpdateProtectedEnvironments(t *testing.T) {
	mux, client := setup(t)

	// Test with DeployAccessLevels, RequiredApprovalCount, and ApprovalRules as if adding new to existing protected environment
	environmentName := "other"

	mux.HandleFunc(fmt.Sprintf("/api/v4/groups/1/protected_environments/%s", environmentName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `{
      "name":"%s",
      "deploy_access_levels": [
        {
          "id": 42,
          "access_level": 30,
          "access_level_description": "Developers + Maintainers"
        }
      ],
      "required_approval_count": 2,
      "approval_rules": [
        {
           "id": 1,
           "user_id": null,
           "group_id": 10,
           "access_level": 5,
           "access_level_description": "devops",
           "required_approvals": 0,
           "group_inheritance_type": 0
        }
      ]
    }`, environmentName)
	})

	expected := &GroupProtectedEnvironment{
		Name: environmentName,
		DeployAccessLevels: []*GroupEnvironmentAccessDescription{
			{
				ID:                     42,
				AccessLevel:            30,
				AccessLevelDescription: "Developers + Maintainers",
			},
		},
		RequiredApprovalCount: 2,
		ApprovalRules: []*GroupEnvironmentApprovalRule{
			{
				ID:                     1,
				GroupID:                10,
				AccessLevel:            5,
				AccessLevelDescription: "devops",
			},
		},
	}

	opt := &UpdateGroupProtectedEnvironmentOptions{
		Name: String(environmentName),
		DeployAccessLevels: &[]*UpdateGroupEnvironmentAccessOptions{
			{AccessLevel: AccessLevel(30)},
		},
		RequiredApprovalCount: Int(2),
		ApprovalRules: &[]*UpdateGroupEnvironmentApprovalRuleOptions{
			{
				GroupID:                Int(10),
				AccessLevel:            AccessLevel(0),
				AccessLevelDescription: String("devops"),
			},
		},
	}

	environment, _, err := client.GroupProtectedEnvironments.UpdateGroupProtectedEnvironment(1, environmentName, opt)
	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, environment)

	// Test with DeployAccessLevels only, as if adding new to existing protected environment
	mux.HandleFunc(fmt.Sprintf("/api/v4/groups/2/protected_environments/%s", environmentName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `{
      "name":"%s",
      "deploy_access_levels": [
        {
          "id": 42,
          "access_level": 30,
          "access_level_description": "Developers + Maintainers"
        }
      ]
    }`, environmentName)
	})

	expected = &GroupProtectedEnvironment{
		Name: environmentName,
		DeployAccessLevels: []*GroupEnvironmentAccessDescription{
			{
				ID:                     42,
				AccessLevel:            30,
				AccessLevelDescription: "Developers + Maintainers",
			},
		},
	}

	opt = &UpdateGroupProtectedEnvironmentOptions{
		Name: String(environmentName),
		DeployAccessLevels: &[]*UpdateGroupEnvironmentAccessOptions{
			{AccessLevel: AccessLevel(30)},
		},
	}
	environment, _, err = client.GroupProtectedEnvironments.UpdateGroupProtectedEnvironment(2, environmentName, opt)
	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, environment)

	// Test update to DeployAccessLevel
	mux.HandleFunc(fmt.Sprintf("/api/v4/groups/3/protected_environments/%s", environmentName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `{
      "name":"%s",
      "deploy_access_levels": [
        {
          "id": 42,
          "access_level": 30,
          "access_level_description": "Developers + Maintainers"
        }
      ],
	  "required_approval_count": 2
    }`, environmentName)
	})

	expected = &GroupProtectedEnvironment{
		Name: environmentName,
		DeployAccessLevels: []*GroupEnvironmentAccessDescription{
			{
				ID:                     42,
				AccessLevel:            30,
				AccessLevelDescription: "Developers + Maintainers",
			},
		},
		RequiredApprovalCount: 2,
	}

	opt = &UpdateGroupProtectedEnvironmentOptions{
		Name: String(environmentName),
		DeployAccessLevels: &[]*UpdateGroupEnvironmentAccessOptions{
			{
				ID:          Int(42),
				AccessLevel: AccessLevel(30),
			},
		},
	}
	environment, _, err = client.GroupProtectedEnvironments.UpdateGroupProtectedEnvironment(3, environmentName, opt)
	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, environment)

	// Test update to ApprovalRules
	mux.HandleFunc(fmt.Sprintf("/api/v4/groups/4/protected_environments/%s", environmentName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `{
      "name":"%s",
      "deploy_access_levels": [
        {
          "id": 42,
          "access_level": 30,
          "access_level_description": "Developers + Maintainers"
        }
      ],
      "required_approval_count": 2,
      "approval_rules": [
        {
           "id": 1,
           "user_id": null,
           "group_id": 10,
           "access_level": 5,
           "access_level_description": "devops",
           "required_approvals": 0,
           "group_inheritance_type": 0
        }
      ]
    }`, environmentName)
	})

	expected = &GroupProtectedEnvironment{
		Name: environmentName,
		DeployAccessLevels: []*GroupEnvironmentAccessDescription{
			{
				ID:                     42,
				AccessLevel:            30,
				AccessLevelDescription: "Developers + Maintainers",
			},
		},
		RequiredApprovalCount: 2,
		ApprovalRules: []*GroupEnvironmentApprovalRule{
			{
				ID:                     1,
				GroupID:                10,
				AccessLevel:            5,
				AccessLevelDescription: "devops",
			},
		},
	}

	opt = &UpdateGroupProtectedEnvironmentOptions{
		Name: String(environmentName),
		ApprovalRules: &[]*UpdateGroupEnvironmentApprovalRuleOptions{
			{
				ID:                     Int(1),
				GroupID:                Int(10),
				AccessLevel:            AccessLevel(0),
				AccessLevelDescription: String("devops"),
			},
		},
	}

	environment, _, err = client.GroupProtectedEnvironments.UpdateGroupProtectedEnvironment(4, environmentName, opt)
	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, environment)

	// Test destroy ApprovalRule
	mux.HandleFunc(fmt.Sprintf("/api/v4/groups/5/protected_environments/%s", environmentName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `{
      "name":"%s",
      "deploy_access_levels": [
        {
          "id": 42,
          "access_level": 30,
          "access_level_description": "Developers + Maintainers"
        }
      ],
      "required_approval_count": 0,
      "approval_rules": []
    }`, environmentName)
	})

	expected = &GroupProtectedEnvironment{
		Name: environmentName,
		DeployAccessLevels: []*GroupEnvironmentAccessDescription{
			{
				ID:                     42,
				AccessLevel:            30,
				AccessLevelDescription: "Developers + Maintainers",
			},
		},
		RequiredApprovalCount: 0,
		ApprovalRules:         []*GroupEnvironmentApprovalRule{},
	}

	opt = &UpdateGroupProtectedEnvironmentOptions{
		Name: String(environmentName),
		ApprovalRules: &[]*UpdateGroupEnvironmentApprovalRuleOptions{
			{
				ID:      Int(1),
				Destroy: Bool(true),
			},
		},
		RequiredApprovalCount: Int(0),
	}

	environment, _, err = client.GroupProtectedEnvironments.UpdateGroupProtectedEnvironment(5, environmentName, opt)
	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, environment)
}

func TestGroupUnprotectEnvironments(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/protected_environments/staging", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.GroupProtectedEnvironments.UnprotectGroupEnvironment(1, "staging")
	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
