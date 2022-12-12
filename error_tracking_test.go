//
// Copyright 2022, Ryan Glab <ryan.j.glab@gmail.com>
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

func TestGetErrorTracking(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/error_tracking/settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"active": true,
			"project_name": "sample sentry project",
			"sentry_external_url": "https://sentry.io/myawesomeproject/project",
			"api_url": "https://sentry.io/api/1/projects/myawesomeproject/project",
			"integrated": false
		}`)
	})

	et, _, err := client.ErrorTracking.GetErrorTrackingSettings(1)
	if err != nil {
		t.Errorf("ErrorTracking.GetErrorTracking returned error: %v", err)
	}

	want := &ErrorTrackingSettings{
		Active:            true,
		ProjectName:       "sample sentry project",
		SentryExternalURL: "https://sentry.io/myawesomeproject/project",
		APIURL:            "https://sentry.io/api/1/projects/myawesomeproject/project",
		Integrated:        false,
	}

	if !reflect.DeepEqual(want, et) {
		t.Errorf("ErrorTracking.GetErrorTracking returned %+v, want %+v", et, want)
	}
}

func TestDisableErrorTracking(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/error_tracking/settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, `{
			"active": false,
			"project_name": "sample sentry project",
			"sentry_external_url": "https://sentry.io/myawesomeproject/project",
			"api_url": "https://sentry.io/api/1/projects/myawesomeproject/project",
			"integrated": false
		}`)
	})

	et, _, err := client.ErrorTracking.EnableDisableErrorTracking(
		1,
		&EnableDisableErrorTrackingOptions{
			Active:     Bool(false),
			Integrated: Bool(false),
		},
	)
	if err != nil {
		t.Errorf("ErrorTracking.EnableDisableErrorTracking returned error: %v", err)
	}

	want := &ErrorTrackingSettings{
		Active:            false,
		ProjectName:       "sample sentry project",
		SentryExternalURL: "https://sentry.io/myawesomeproject/project",
		APIURL:            "https://sentry.io/api/1/projects/myawesomeproject/project",
		Integrated:        false,
	}

	if !reflect.DeepEqual(want, et) {
		t.Errorf("ErrorTracking.EnableDisableErrorTracking returned %+v, want %+v", et, want)
	}
}

func TestListErrorTrackingClientKeys(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/error_tracking/client_keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"active": true,
				"public_key": "glet_aa77551d849c083f76d0bc545ed053a3",
				"sentry_dsn": "https://glet_aa77551d849c083f76d0bc545ed053a3@gitlab.example.com/api/v4/error_tracking/collector/5"
			}
		]`)
	})

	cks, _, err := client.ErrorTracking.ListClientKeys(1, &ListClientKeysOptions{Page: 1, PerPage: 10})
	if err != nil {
		t.Errorf("ErrorTracking.ListErrorTrackingClientKeys returned error: %v", err)
	}

	want := []*ErrorTrackingClientKey{{
		ID:        1,
		Active:    true,
		PublicKey: "glet_aa77551d849c083f76d0bc545ed053a3",
		SentryDsn: "https://glet_aa77551d849c083f76d0bc545ed053a3@gitlab.example.com/api/v4/error_tracking/collector/5",
	}}

	if !reflect.DeepEqual(want, cks) {
		t.Errorf("ErrorTracking.ListErrorTrackingClientKeys returned %+v, want %+v", cks, want)
	}
}

func TestCreateClientKey(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/error_tracking/client_keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"id": 1,
			"active": true,
			"public_key": "glet_aa77551d849c083f76d0bc545ed053a3",
			"sentry_dsn": "https://glet_aa77551d849c083f76d0bc545ed053a3@gitlab.example.com/api/v4/error_tracking/collector/5"
		}`)
	})

	ck, _, err := client.ErrorTracking.CreateClientKey(1)
	if err != nil {
		t.Errorf("ErrorTracking.CreateClientKey returned error: %v", err)
	}

	want := &ErrorTrackingClientKey{
		ID:        1,
		Active:    true,
		PublicKey: "glet_aa77551d849c083f76d0bc545ed053a3",
		SentryDsn: "https://glet_aa77551d849c083f76d0bc545ed053a3@gitlab.example.com/api/v4/error_tracking/collector/5",
	}

	if !reflect.DeepEqual(want, ck) {
		t.Errorf("ErrorTracking.CreateClientKey returned %+v, want %+v", ck, want)
	}
}

func TestDeleteClientKey(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/error_tracking/client_keys/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, "/api/v4/projects/1/error_tracking/client_keys/3")
	})

	_, err := client.ErrorTracking.DeleteClientKey(1, 3)
	if err != nil {
		t.Errorf("ErrorTracking.DeleteClientKey returned error: %v", err)
	}
}
