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

func TestGetSettings(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/application/settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1,    "default_projects_limit" : 100000}`)
	})

	settings, _, err := client.Settings.GetSettings()
	if err != nil {
		t.Fatal(err)
	}

	want := &Settings{ID: 1, DefaultProjectsLimit: 100000}
	if !reflect.DeepEqual(settings, want) {
		t.Errorf("Settings.GetSettings returned %+v, want %+v", settings, want)
	}
}

func TestUpdateSettings(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/application/settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{"default_projects_limit" : 100}`)
	})

	options := &UpdateSettingsOptions{
		DefaultProjectsLimit: Int(100),
	}
	settings, _, err := client.Settings.UpdateSettings(options)
	if err != nil {
		t.Fatal(err)
	}

	want := &Settings{DefaultProjectsLimit: 100}
	if !reflect.DeepEqual(settings, want) {
		t.Errorf("Settings.UpdateSettings returned %+v, want %+v", settings, want)
	}
}
