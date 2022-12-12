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
	mux, client := setup(t)

	const url = "https://www.gravatar.com/avatar/10e6bf7bcf22c2f00a3ef684b4ada178"

	mux.HandleFunc("/api/v4/avatar",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.WriteHeader(http.StatusAccepted)
			avatar := Avatar{AvatarURL: url}
			resp, _ := json.Marshal(avatar)
			_, _ = w.Write(resp)
		},
	)

	opt := &GetAvatarOptions{Email: String("sander@vanharmelen.nnl")}
	avatar, resp, err := client.Avatar.GetAvatar(opt)
	if err != nil {
		t.Fatalf("Avatar.GetAvatar returned error: %v", err)
	}

	if resp.Status != "202 Accepted" {
		t.Fatalf("Avatar.GetAvatar returned wrong status code: %v", resp.Status)
	}

	if url != avatar.AvatarURL {
		t.Errorf("Avatar.GetAvatar wrong result %s, want %s", avatar.AvatarURL, url)
	}
}
