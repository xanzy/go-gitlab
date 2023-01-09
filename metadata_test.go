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
)

func TestGetMetadata(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/metadata",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `{
        "version": "15.6.0-pre",
        "revision": "016e8d8bdc3",
        "enterprise": true,
        "kas": {
          "enabled": true,
          "externalUrl": "wss://kas.gitlab.com",
          "version": "15.6.0-rc2"
        }
      }`)
		})

	version, _, err := client.Metadata.GetMetadata()
	if err != nil {
		t.Errorf("Metadata.GetMetadata returned error: %v", err)
	}

	want := &Metadata{
		Version: "15.6.0-pre", Revision: "016e8d8bdc3", KAS: struct {
			Enabled     bool   `json:"enabled"`
			ExternalURL string `json:"externalUrl"`
			Version     string `json:"version"`
		}{
			Enabled:     true,
			ExternalURL: "wss://kas.gitlab.com",
			Version:     "15.6.0-rc2",
		},
		Enterprise: true,
	}
	if !reflect.DeepEqual(want, version) {
		t.Errorf("Metadata.GetMetadata returned %+v, want %+v", version, want)
	}
}
