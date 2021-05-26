//
// Copyright 2021, Pavel Kostohrys
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
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetAvatar(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	const host = "https://google.com"

	mux.HandleFunc("/api/v4/avatar",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.WriteHeader(http.StatusAccepted)
			avatar := struct {
				Url string `json:"avatar_url"`
			}{
				Url: host,
			}
			resp, _ := json.Marshal(avatar)
			_, _ = w.Write(resp)
		},
	)

	err, resp := client.Avatar.GetAvatar(&AvatarOptions{Email: "test"})
	if err != nil {
		t.Fatalf("Avatar.GetAvatar returned error: %v", err)
	}

	if host != *resp {
		t.Errorf("Avatar.GetAvatar wrong result %s, want %s", *resp, host)
	}
}
