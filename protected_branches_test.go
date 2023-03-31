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
	"reflect"
	"testing"
)

func TestListProtectedBranches(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/protected_branches", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
	{
		"id":1,
		"name":"master",
		"push_access_levels":[{
			"id":1,
			"access_level":40,
			"access_level_description":"Maintainers"
		}],
		"merge_access_levels":[{
			"id":1,
			"access_level":40,
			"access_level_description":"Maintainers"
		}],
		"code_owner_approval_required":false
	}
]`)
	})
	opt := &ListProtectedBranchesOptions{}
	protectedBranches, _, err := client.ProtectedBranches.ListProtectedBranches("1", opt)
	if err != nil {
		t.Errorf("ProtectedBranches.ListProtectedBranches returned error: %v", err)
	}
	want := []*ProtectedBranch{
		{
			ID:   1,
			Name: "master",
			PushAccessLevels: []*BranchAccessDescription{
				{
					ID:                     1,
					AccessLevel:            40,
					AccessLevelDescription: "Maintainers",
				},
			},
			MergeAccessLevels: []*BranchAccessDescription{
				{
					ID:                     1,
					AccessLevel:            40,
					AccessLevelDescription: "Maintainers",
				},
			},
			AllowForcePush:            false,
			CodeOwnerApprovalRequired: false,
		},
	}
	if !reflect.DeepEqual(want, protectedBranches) {
		t.Errorf("ProtectedBranches.ListProtectedBranches returned %+v, want %+v", protectedBranches, want)
	}
}

func TestListProtectedBranchesWithoutCodeOwnerApproval(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/protected_branches", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
	{
		"id":1,
		"name":"master",
		"push_access_levels":[{
			"access_level":40,
			"access_level_description":"Maintainers"
		}],
		"merge_access_levels":[{
			"access_level":40,
			"access_level_description":"Maintainers"
		}]
	}
]`)
	})
	opt := &ListProtectedBranchesOptions{}
	protectedBranches, _, err := client.ProtectedBranches.ListProtectedBranches("1", opt)
	if err != nil {
		t.Errorf("ProtectedBranches.ListProtectedBranches returned error: %v", err)
	}
	want := []*ProtectedBranch{
		{
			ID:   1,
			Name: "master",
			PushAccessLevels: []*BranchAccessDescription{
				{
					AccessLevel:            40,
					AccessLevelDescription: "Maintainers",
				},
			},
			MergeAccessLevels: []*BranchAccessDescription{
				{
					AccessLevel:            40,
					AccessLevelDescription: "Maintainers",
				},
			},
			AllowForcePush:            false,
			CodeOwnerApprovalRequired: false,
		},
	}
	if !reflect.DeepEqual(want, protectedBranches) {
		t.Errorf("Projects.ListProjects returned %+v, want %+v", protectedBranches, want)
	}
}

func TestProtectRepositoryBranches(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/protected_branches", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `
	{
		"id":1,
		"name":"master",
		"push_access_levels":[{
			"access_level":40,
			"access_level_description":"Maintainers"
		}],
		"merge_access_levels":[{
			"access_level":40,
			"access_level_description":"Maintainers"
		}],
		"allow_force_push":true,
		"code_owner_approval_required":true
	}`)
	})
	opt := &ProtectRepositoryBranchesOptions{
		Name:                      String("master"),
		PushAccessLevel:           AccessLevel(MaintainerPermissions),
		MergeAccessLevel:          AccessLevel(MaintainerPermissions),
		AllowForcePush:            Bool(true),
		CodeOwnerApprovalRequired: Bool(true),
	}
	projects, _, err := client.ProtectedBranches.ProtectRepositoryBranches("1", opt)
	if err != nil {
		t.Errorf("ProtectedBranches.ProtectRepositoryBranches returned error: %v", err)
	}
	want := &ProtectedBranch{
		ID:   1,
		Name: "master",
		PushAccessLevels: []*BranchAccessDescription{
			{
				AccessLevel:            40,
				AccessLevelDescription: "Maintainers",
			},
		},
		MergeAccessLevels: []*BranchAccessDescription{
			{
				AccessLevel:            40,
				AccessLevelDescription: "Maintainers",
			},
		},
		AllowForcePush:            true,
		CodeOwnerApprovalRequired: true,
	}
	if !reflect.DeepEqual(want, projects) {
		t.Errorf("Projects.ListProjects returned %+v, want %+v", projects, want)
	}
}

func TestUpdateRepositoryBranches(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/protected_branches/master", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testBody(t, r, `{"code_owner_approval_required":true}`)
		fmt.Fprintf(w, `{
			"name": "master",
			"code_owner_approval_required": true
		}`)
	})
	opt := &UpdateProtectedBranchOptions{
		CodeOwnerApprovalRequired: Bool(true),
	}
	protectedBranch, _, err := client.ProtectedBranches.UpdateProtectedBranch("1", "master", opt)
	if err != nil {
		t.Errorf("ProtectedBranches.UpdateProtectedBranch returned error: %v", err)
	}

	want := &ProtectedBranch{
		Name:                      "master",
		CodeOwnerApprovalRequired: true,
	}

	if !reflect.DeepEqual(want, protectedBranch) {
		t.Errorf("ProtectedBranches.UpdateProtectedBranch returned %+v, want %+v", protectedBranch, want)
	}
}
