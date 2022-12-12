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

func TestListProtectedTags(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/protected_tags", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"name":"1.0.0", "create_access_levels": [{"access_level": 40, "access_level_description": "Maintainers"}]},{"name":"*-release", "create_access_levels": [{"access_level": 30, "access_level_description": "Developers + Maintainers"}]}]`)
	})

	expected := []*ProtectedTag{
		{
			Name: "1.0.0",
			CreateAccessLevels: []*TagAccessDescription{
				{
					AccessLevel:            40,
					AccessLevelDescription: "Maintainers",
				},
			},
		},
		{
			Name: "*-release",
			CreateAccessLevels: []*TagAccessDescription{
				{
					AccessLevel:            30,
					AccessLevelDescription: "Developers + Maintainers",
				},
			},
		},
	}

	opt := &ListProtectedTagsOptions{}
	tags, _, err := client.ProtectedTags.ListProtectedTags(1, opt)
	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, tags)
}

func TestGetProtectedTag(t *testing.T) {
	mux, client := setup(t)

	tagName := "my-awesome-tag"

	mux.HandleFunc(fmt.Sprintf("/api/v4/projects/1/protected_tags/%s", tagName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"name":"my-awesome-tag", "create_access_levels": [{"access_level": 30, "access_level_description": "Developers + Maintainers"},{"access_level": 40, "access_level_description": "Sample Group", "group_id": 300}]}`)
	})

	expected := &ProtectedTag{
		Name: tagName,
		CreateAccessLevels: []*TagAccessDescription{
			{
				AccessLevel:            30,
				AccessLevelDescription: "Developers + Maintainers",
			},
			{
				AccessLevel:            40,
				GroupID:                300,
				AccessLevelDescription: "Sample Group",
			},
		},
	}

	tag, _, err := client.ProtectedTags.GetProtectedTag(1, tagName)

	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, tag)
}

func TestProtectRepositoryTags(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/protected_tags", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"name":"my-awesome-tag", "create_access_levels": [{"access_level": 30, "access_level_description": "Developers + Maintainers"},{"access_level": 40, "access_level_description": "Sample Group", "group_id": 300}]}`)
	})

	expected := &ProtectedTag{
		Name: "my-awesome-tag",
		CreateAccessLevels: []*TagAccessDescription{
			{
				AccessLevel:            30,
				AccessLevelDescription: "Developers + Maintainers",
			},
			{
				AccessLevel:            40,
				GroupID:                300,
				AccessLevelDescription: "Sample Group",
			},
		},
	}

	opt := &ProtectRepositoryTagsOptions{
		Name:              String("my-awesome-tag"),
		CreateAccessLevel: AccessLevel(30),
		AllowedToCreate: &[]*TagsPermissionOptions{
			{
				GroupID: Int(300),
			},
		},
	}
	tag, _, err := client.ProtectedTags.ProtectRepositoryTags(1, opt)

	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, expected, tag)
}

func TestUnprotectRepositoryTags(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/protected_tags/my-awesome-tag", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.ProtectedTags.UnprotectRepositoryTags(1, "my-awesome-tag")
	assert.NoError(t, err, "failed to get response")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
