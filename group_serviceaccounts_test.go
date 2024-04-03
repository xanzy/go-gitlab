//
// Copyright 2023, James Hong
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
	"github.com/stretchr/testify/require"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestCreateServiceAccount(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/service_accounts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `
      {
	      "id": 57,
	      "username": "service_account_group_345_6018816a18e515214e0c34c2b33523fc",
	      "name": "Service account user"
      }`)
	})

	sa, _, err := client.Groups.CreateServiceAccount(1)
	if err != nil {
		t.Error(err)
	}

	want := &GroupServiceAccount{
		ID:       57,
		UserName: "service_account_group_345_6018816a18e515214e0c34c2b33523fc",
		Name:     "Service account user",
	}

	if !reflect.DeepEqual(sa, want) {
		t.Errorf("CreateServiceAccount returned \ngot:\n%v\nwant:\n%v", Stringify(sa), Stringify(want))
	}
}

func TestCreateServiceAccountPersonalAccessToken(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/service_accounts/57/personal_access_tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `
      {
      	"id":6,
      	"name":"service_account_token",
      	"revoked":false,
      	"created_at":"2023-06-13T07:47:13.000Z",
      	"scopes":["api"],
      	"user_id":71,
      	"last_used_at":null,
      	"active":true,
      	"expires_at":"2024-06-12",
      	"token":"random_token"
      }`)
	})

	expireTime, err := ParseISOTime("2024-06-12")
	require.NoError(t, err)

	options := &CreateServiceAccountPersonalAccessTokenOptions{
		Scopes:    Ptr([]string{"api"}),
		Name:      Ptr("service_account_token"),
		ExpiresAt: Ptr(expireTime),
	}
	pat, _, err := client.Groups.CreateServiceAccountPersonalAccessToken(1, 57, options)
	if err != nil {
		t.Error(err)
	}

	datePointer := time.Date(2023, 0o6, 13, 0o7, 47, 13, 0, time.UTC)
	expiresAt := ISOTime(time.Date(2024, time.June, 12, 0, 0, 0, 0, time.UTC))

	want := &PersonalAccessToken{
		ID:         6,
		Name:       "service_account_token",
		Revoked:    false,
		CreatedAt:  &datePointer,
		Scopes:     []string{"api"},
		UserID:     71,
		LastUsedAt: nil,
		Active:     true,
		ExpiresAt:  &expiresAt,
		Token:      "random_token",
	}

	if !reflect.DeepEqual(pat, want) {
		t.Errorf("CreateServiceAccountPersonalAccessToken returned \ngot:\n%v\nwant:\n%v", Stringify(pat), Stringify(want))
	}
}

func TestRotateServiceAccountPersonalAccessToken(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/service_accounts/57/personal_access_tokens/6/rotate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `
      {
      	"id":7,
      	"name":"service_account_token",
      	"revoked":false,
      	"created_at":"2023-06-13T07:54:49.000Z",
      	"scopes":["api"],
      	"user_id":71,
      	"last_used_at":null,
      	"active":true,
      	"expires_at":"2025-06-20",
      	"token":"random_token_2"
      }`)
	})

	pat, _, err := client.Groups.RotateServiceAccountPersonalAccessToken(1, 57, 6)
	if err != nil {
		t.Error(err)
	}

	datePointer := time.Date(2023, 0o6, 13, 0o7, 54, 49, 0, time.UTC)
	expiresAt := ISOTime(time.Date(2025, time.June, 20, 0, 0, 0, 0, time.UTC))

	want := &PersonalAccessToken{
		ID:         7,
		Name:       "service_account_token",
		Revoked:    false,
		CreatedAt:  &datePointer,
		Scopes:     []string{"api"},
		UserID:     71,
		LastUsedAt: nil,
		Active:     true,
		ExpiresAt:  &expiresAt,
		Token:      "random_token_2",
	}

	if !reflect.DeepEqual(pat, want) {
		t.Errorf("RotateServiceAccountPersonalAccessToken returned \ngot:\n%v\nwant:\n%v", Stringify(pat), Stringify(want))
	}
}
