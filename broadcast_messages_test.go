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
	"time"
)

func TestListBroadcastMessages(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/broadcast_messages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[{
			"message": "Some Message",
			"starts_at": "2017-06-26T06:00:00.000Z",
			"ends_at": "2017-06-27T12:59:00.000Z",
			"color": "#E75E40",
			"font": "#FFFFFF",
			"id": 1,
			"active": false,
			"target_access_levels": [10,30],
			"target_path": "*/welcome",
			"broadcast_type": "banner",
			"dismissable": false
		},{
			"message": "SomeMessage2",
			"starts_at": "2015-04-27T06:43:00.000Z",
			"ends_at": "2015-04-28T20:43:00.000Z",
			"color": "#AA33EE",
			"font": "#224466",
			"id": 2,
			"active": true,
			"target_access_levels": [],
			"target_path": "*/*",
			"broadcast_type": "notification",
			"dismissable": true
		}]`)
	})

	got, _, err := client.BroadcastMessage.ListBroadcastMessages(nil, nil)
	if err != nil {
		t.Errorf("ListBroadcastMessages returned error: %v", err)
	}

	wantedFirstStartsAt := time.Date(2017, 0o6, 26, 6, 0, 0, 0, time.UTC)
	wantedFirstEndsAt := time.Date(2017, 0o6, 27, 12, 59, 0, 0, time.UTC)

	wantedSecondStartsAt := time.Date(2015, 0o4, 27, 6, 43, 0, 0, time.UTC)
	wantedSecondEndsAt := time.Date(2015, 0o4, 28, 20, 43, 0, 0, time.UTC)

	want := []*BroadcastMessage{{
		Message:            "Some Message",
		StartsAt:           &wantedFirstStartsAt,
		EndsAt:             &wantedFirstEndsAt,
		Color:              "#E75E40",
		Font:               "#FFFFFF",
		ID:                 1,
		Active:             false,
		TargetAccessLevels: []AccessLevelValue{GuestPermissions, DeveloperPermissions},
		TargetPath:         "*/welcome",
		BroadcastType:      "banner",
		Dismissable:        false,
	}, {
		Message:            "SomeMessage2",
		StartsAt:           &wantedSecondStartsAt,
		EndsAt:             &wantedSecondEndsAt,
		Color:              "#AA33EE",
		Font:               "#224466",
		ID:                 2,
		Active:             true,
		TargetAccessLevels: []AccessLevelValue{},
		TargetPath:         "*/*",
		BroadcastType:      "notification",
		Dismissable:        true,
	}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ListBroadcastMessages returned \ngot:\n%v\nwant:\n%v", Stringify(got), Stringify(want))
	}
}

func TestGetBroadcastMessages(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/broadcast_messages/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
			"message": "Some Message",
			"starts_at": "2017-06-26T06:00:00.000Z",
			"ends_at": "2017-06-27T12:59:00.000Z",
			"color": "#E75E40",
			"font": "#FFFFFF",
			"id": 1,
			"active": false,
			"target_access_levels": [10,30],
			"target_path": "*/welcome",
			"broadcast_type": "banner",
			"dismissable": false
		}`)
	})

	got, _, err := client.BroadcastMessage.GetBroadcastMessage(1)
	if err != nil {
		t.Errorf("GetBroadcastMessage returned error: %v", err)
	}

	wantedStartsAt := time.Date(2017, time.June, 26, 6, 0, 0, 0, time.UTC)
	wantedEndsAt := time.Date(2017, time.June, 27, 12, 59, 0, 0, time.UTC)

	want := &BroadcastMessage{
		Message:            "Some Message",
		StartsAt:           &wantedStartsAt,
		EndsAt:             &wantedEndsAt,
		Color:              "#E75E40",
		Font:               "#FFFFFF",
		ID:                 1,
		Active:             false,
		TargetAccessLevels: []AccessLevelValue{GuestPermissions, DeveloperPermissions},
		TargetPath:         "*/welcome",
		BroadcastType:      "banner",
		Dismissable:        false,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetBroadcastMessage returned \ngot:\n%v\nwant:\n%v", Stringify(got), Stringify(want))
	}
}

func TestCreateBroadcastMessages(t *testing.T) {
	mux, client := setup(t)

	wantedStartsAt := time.Date(2017, time.June, 26, 6, 0, 0, 0, time.UTC)
	wantedEndsAt := time.Date(2017, time.June, 27, 12, 59, 0, 0, time.UTC)

	mux.HandleFunc("/api/v4/broadcast_messages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `{
			"message": "Some Message",
			"starts_at": "2017-06-26T06:00:00.000Z",
			"ends_at": "2017-06-27T12:59:00.000Z",
			"color": "#E75E40",
			"font": "#FFFFFF",
			"id": 42,
			"active": false,
			"target_access_levels": [10,30],
			"target_path": "*/welcome",
			"broadcast_type": "banner",
			"dismissable": false
		}`)
	})

	opt := &CreateBroadcastMessageOptions{
		Message:            String("Some Message"),
		StartsAt:           &wantedStartsAt,
		EndsAt:             &wantedEndsAt,
		Color:              String("#E75E40"),
		Font:               String("#FFFFFF"),
		TargetAccessLevels: []AccessLevelValue{GuestPermissions, DeveloperPermissions},
		TargetPath:         String("*/welcome"),
		BroadcastType:      String("banner"),
		Dismissable:        Bool(false),
	}

	got, _, err := client.BroadcastMessage.CreateBroadcastMessage(opt)
	if err != nil {
		t.Errorf("CreateBroadcastMessage returned error: %v", err)
	}

	want := &BroadcastMessage{
		Message:            "Some Message",
		StartsAt:           &wantedStartsAt,
		EndsAt:             &wantedEndsAt,
		Color:              "#E75E40",
		Font:               "#FFFFFF",
		ID:                 42,
		Active:             false,
		TargetAccessLevels: []AccessLevelValue{GuestPermissions, DeveloperPermissions},
		TargetPath:         "*/welcome",
		BroadcastType:      "banner",
		Dismissable:        false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("CreateBroadcastMessage returned \ngot:\n%v\nwant:\n%v", Stringify(got), Stringify(want))
	}
}

func TestUpdateBroadcastMessages(t *testing.T) {
	mux, client := setup(t)

	wantedStartsAt := time.Date(2017, time.June, 26, 6, 0, 0, 0, time.UTC)
	wantedEndsAt := time.Date(2017, time.June, 27, 12, 59, 0, 0, time.UTC)

	mux.HandleFunc("/api/v4/broadcast_messages/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `{
			"message": "Some Message Updated",
			"starts_at": "2017-06-26T06:00:00.000Z",
			"ends_at": "2017-06-27T12:59:00.000Z",
			"color": "#E75E40",
			"font": "#FFFFFF",
			"id": 42,
			"active": false,
			"target_access_levels": [10,30],
			"target_path": "*/welcome",
			"broadcast_type": "banner",
			"dismissable": false
		}`)
	})

	opt := &UpdateBroadcastMessageOptions{
		Message:            String("Some Message Updated"),
		StartsAt:           &wantedStartsAt,
		EndsAt:             &wantedEndsAt,
		Color:              String("#E75E40"),
		Font:               String("#FFFFFF"),
		TargetAccessLevels: []AccessLevelValue{GuestPermissions, DeveloperPermissions},
		TargetPath:         String("*/welcome"),
		BroadcastType:      String("banner"),
		Dismissable:        Bool(false),
	}

	got, _, err := client.BroadcastMessage.UpdateBroadcastMessage(1, opt)
	if err != nil {
		t.Errorf("UpdateBroadcastMessage returned error: %v", err)
	}

	want := &BroadcastMessage{
		Message:            "Some Message Updated",
		StartsAt:           &wantedStartsAt,
		EndsAt:             &wantedEndsAt,
		Color:              "#E75E40",
		Font:               "#FFFFFF",
		ID:                 42,
		Active:             false,
		TargetAccessLevels: []AccessLevelValue{GuestPermissions, DeveloperPermissions},
		TargetPath:         "*/welcome",
		BroadcastType:      "banner",
		Dismissable:        false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("UpdateBroadcastMessage returned \ngot:\n%v\nwant:\n%v", Stringify(got), Stringify(want))
	}
}

func TestDeleteBroadcastMessages(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/broadcast_messages/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.BroadcastMessage.DeleteBroadcastMessage(1)
	if err != nil {
		t.Errorf("UpdateBroadcastMessage returned error: %v", err)
	}
}
