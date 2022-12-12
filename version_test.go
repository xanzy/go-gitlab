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

func TestGetVersion(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/version",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `{"version":"11.3.4-ee", "revision":"14d3a1d"}`)
		})

	version, _, err := client.Version.GetVersion()
	if err != nil {
		t.Errorf("Version.GetVersion returned error: %v", err)
	}

	want := &Version{Version: "11.3.4-ee", Revision: "14d3a1d"}
	if !reflect.DeepEqual(want, version) {
		t.Errorf("Version.GetVersion returned %+v, want %+v", version, want)
	}
}
