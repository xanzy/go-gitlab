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
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchService_Users(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/search_users.json")
	})

	opts := &SearchOptions{ListOptions: ListOptions{PerPage: 2}}
	users, _, err := client.Search.Users("doe", opts)

	require.NoError(t, err)

	want := []*User{{
		ID:        1,
		Username:  "user1",
		Name:      "John Doe1",
		State:     "active",
		AvatarURL: "http://www.gravatar.com/avatar/c922747a93b40d1ea88262bf1aebee62?s=80&d=identicon",
		WebURL:    "http://localhost/user1",
	}}
	require.Equal(t, want, users)
}

func TestSearchService_UsersByGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/3/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/search_users.json")
	})

	opts := &SearchOptions{ListOptions: ListOptions{PerPage: 2}}
	users, _, err := client.Search.UsersByGroup("3", "doe", opts)

	require.NoError(t, err)

	want := []*User{{
		ID:        1,
		Username:  "user1",
		Name:      "John Doe1",
		State:     "active",
		AvatarURL: "http://www.gravatar.com/avatar/c922747a93b40d1ea88262bf1aebee62?s=80&d=identicon",
		WebURL:    "http://localhost/user1",
	}}
	require.Equal(t, want, users)
}

func TestSearchService_UsersByProject(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/6/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/search_users.json")
	})

	opts := &SearchOptions{ListOptions: ListOptions{PerPage: 2}}
	users, _, err := client.Search.UsersByProject("6", "doe", opts)

	require.NoError(t, err)

	want := []*User{{
		ID:        1,
		Username:  "user1",
		Name:      "John Doe1",
		State:     "active",
		AvatarURL: "http://www.gravatar.com/avatar/c922747a93b40d1ea88262bf1aebee62?s=80&d=identicon",
		WebURL:    "http://localhost/user1",
	}}
	require.Equal(t, want, users)
}
