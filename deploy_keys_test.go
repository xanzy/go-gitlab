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

func TestListAllDeployKeys(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/deploy_keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			  "id": 1,
			  "title": "Public key",
			  "key": "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAIEAiPWx6WM4lhHNedGfBpPJNPpZ7yKu+dnn1SJejgt4596k6YjzGGphH2TUxwKzxcKDKKezwkpfnxPkSMkuEspGRt/aZZ9wa++Oi7Qkr8prgHc4soW6NUlfDzpvZK2H5E7eQaSeP3SAwGmQKUFHCddNaP0L+hM7zhFNzjFvpaMgJw0=",
			  "created_at": "2013-10-02T10:12:29Z"
			},
			{
			  "id": 3,
			  "title": "Another Public key",
			  "key": "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAIEAiPWx6WM4lhHNedGfBpPJNPpZ7yKu+dnn1SJejgt4596k6YjzGGphH2TUxwKzxcKDKKezwkpfnxPkSMkuEspGRt/aZZ9wa++Oi7Qkr8prgHc4soW6NUlfDzpvZK2H5E7eQaSeP3SAwGmQKUFHCddNaP0L+hM7zhFNzjFvpaMgJw0=",
			  "created_at": "2013-10-02T11:12:29Z"
			}
		  ]`)
	})

	deployKeys, _, err := client.DeployKeys.ListAllDeployKeys()
	if err != nil {
		t.Errorf("DeployKeys.ListAllDeployKeys returned error: %v", err)
	}

	timeLayout := "2006-01-02T15:04:05Z07:00"
	createdAt, err := time.Parse(timeLayout, "2013-10-02T10:12:29Z")
	if err != nil {
		t.Errorf("DeployKeys.ListAllDeployKeys returned an error while parsing time: %v", err)
	}

	createdAt2, err := time.Parse(timeLayout, "2013-10-02T11:12:29Z")
	if err != nil {
		t.Errorf("DeployKeys.ListAllDeployKeys returned an error while parsing time: %v", err)
	}

	want := []*DeployKey{
		{
			ID:        1,
			Title:     "Public key",
			Key:       "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAIEAiPWx6WM4lhHNedGfBpPJNPpZ7yKu+dnn1SJejgt4596k6YjzGGphH2TUxwKzxcKDKKezwkpfnxPkSMkuEspGRt/aZZ9wa++Oi7Qkr8prgHc4soW6NUlfDzpvZK2H5E7eQaSeP3SAwGmQKUFHCddNaP0L+hM7zhFNzjFvpaMgJw0=",
			CreatedAt: &createdAt,
		},
		{
			ID:        3,
			Title:     "Another Public key",
			Key:       "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAIEAiPWx6WM4lhHNedGfBpPJNPpZ7yKu+dnn1SJejgt4596k6YjzGGphH2TUxwKzxcKDKKezwkpfnxPkSMkuEspGRt/aZZ9wa++Oi7Qkr8prgHc4soW6NUlfDzpvZK2H5E7eQaSeP3SAwGmQKUFHCddNaP0L+hM7zhFNzjFvpaMgJw0=",
			CreatedAt: &createdAt2,
		},
	}
	if !reflect.DeepEqual(want, deployKeys) {
		t.Errorf("DeployKeys.ListAllDeployKeys returned %+v, want %+v", deployKeys, want)
	}
}

func TestListProjectDeployKeys(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/5/deploy_keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			  "id": 1,
			  "title": "Public key",
			  "key": "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAIEAiPWx6WM4lhHNedGfBpPJNPpZ7yKu+dnn1SJejgt4596k6YjzGGphH2TUxwKzxcKDKKezwkpfnxPkSMkuEspGRt/aZZ9wa++Oi7Qkr8prgHc4soW6NUlfDzpvZK2H5E7eQaSeP3SAwGmQKUFHCddNaP0L+hM7zhFNzjFvpaMgJw0=",
			  "created_at": "2013-10-02T10:12:29Z",
			  "can_push": false
			},
			{
			  "id": 3,
			  "title": "Another Public key",
			  "key": "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAIEAiPWx6WM4lhHNedGfBpPJNPpZ7yKu+dnn1SJejgt4596k6YjzGGphH2TUxwKzxcKDKKezwkpfnxPkSMkuEspGRt/aZZ9wa++Oi7Qkr8prgHc4soW6NUlfDzpvZK2H5E7eQaSeP3SAwGmQKUFHCddNaP0L+hM7zhFNzjFvpaMgJw0=",
			  "created_at": "2013-10-02T11:12:29Z",
			  "can_push": false
			}
		  ]`)
	})

	deployKeys, _, err := client.DeployKeys.ListProjectDeployKeys(5, &ListProjectDeployKeysOptions{})
	if err != nil {
		t.Errorf("DeployKeys.ListProjectDeployKeys returned error: %v", err)
	}

	timeLayout := "2006-01-02T15:04:05Z07:00"
	createdAt, err := time.Parse(timeLayout, "2013-10-02T10:12:29Z")
	if err != nil {
		t.Errorf("DeployKeys.ListAllDeployKeys returned an error while parsing time: %v", err)
	}

	createdAt2, err := time.Parse(timeLayout, "2013-10-02T11:12:29Z")
	if err != nil {
		t.Errorf("DeployKeys.ListAllDeployKeys returned an error while parsing time: %v", err)
	}

	canPush := false

	want := []*DeployKey{
		{
			ID:        1,
			Title:     "Public key",
			Key:       "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAIEAiPWx6WM4lhHNedGfBpPJNPpZ7yKu+dnn1SJejgt4596k6YjzGGphH2TUxwKzxcKDKKezwkpfnxPkSMkuEspGRt/aZZ9wa++Oi7Qkr8prgHc4soW6NUlfDzpvZK2H5E7eQaSeP3SAwGmQKUFHCddNaP0L+hM7zhFNzjFvpaMgJw0=",
			CreatedAt: &createdAt,
			CanPush:   &canPush,
		},
		{
			ID:        3,
			Title:     "Another Public key",
			Key:       "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAIEAiPWx6WM4lhHNedGfBpPJNPpZ7yKu+dnn1SJejgt4596k6YjzGGphH2TUxwKzxcKDKKezwkpfnxPkSMkuEspGRt/aZZ9wa++Oi7Qkr8prgHc4soW6NUlfDzpvZK2H5E7eQaSeP3SAwGmQKUFHCddNaP0L+hM7zhFNzjFvpaMgJw0=",
			CreatedAt: &createdAt2,
			CanPush:   &canPush,
		},
	}
	if !reflect.DeepEqual(want, deployKeys) {
		t.Errorf("DeployKeys.ListProjectDeployKeys returned %+v, want %+v", deployKeys, want)
	}
}

func TestGetDeployKey(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/5/deploy_keys/11", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
			"id": 1,
			"title": "Public key",
			"key": "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAIEAiPWx6WM4lhHNedGfBpPJNPpZ7yKu+dnn1SJejgt4596k6YjzGGphH2TUxwKzxcKDKKezwkpfnxPkSMkuEspGRt/aZZ9wa++Oi7Qkr8prgHc4soW6NUlfDzpvZK2H5E7eQaSeP3SAwGmQKUFHCddNaP0L+hM7zhFNzjFvpaMgJw0=",
			"created_at": "2013-10-02T10:12:29Z",
			"can_push": false
		  }`)
	})

	deployKey, _, err := client.DeployKeys.GetDeployKey(5, 11)
	if err != nil {
		t.Errorf("DeployKeys.GetDeployKey returned error: %v", err)
	}

	timeLayout := "2006-01-02T15:04:05Z07:00"
	createdAt, err := time.Parse(timeLayout, "2013-10-02T10:12:29Z")
	if err != nil {
		t.Errorf("DeployKeys.ListAllDeployKeys returned an error while parsing time: %v", err)
	}

	canPush := false

	want := &DeployKey{
		ID:        1,
		Title:     "Public key",
		Key:       "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAIEAiPWx6WM4lhHNedGfBpPJNPpZ7yKu+dnn1SJejgt4596k6YjzGGphH2TUxwKzxcKDKKezwkpfnxPkSMkuEspGRt/aZZ9wa++Oi7Qkr8prgHc4soW6NUlfDzpvZK2H5E7eQaSeP3SAwGmQKUFHCddNaP0L+hM7zhFNzjFvpaMgJw0=",
		CreatedAt: &createdAt,
		CanPush:   &canPush,
	}
	if !reflect.DeepEqual(want, deployKey) {
		t.Errorf("DeployKeys.GetDeployKey returned %+v, want %+v", deployKey, want)
	}
}

func TestAddDeployKey(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/5/deploy_keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `{
			"key" : "ssh-rsa AAAA...",
			"id" : 12,
			"title" : "My deploy key",
			"can_push": true,
			"created_at" : "2015-08-29T12:44:31.550Z"
		 }`)
	})

	key := "ssh-rsa AAAA..."
	title := "My deploy key"
	canPush := true
	opt := &AddDeployKeyOptions{
		Key:     &key,
		Title:   &title,
		CanPush: &canPush,
	}
	deployKey, _, err := client.DeployKeys.AddDeployKey(5, opt)
	if err != nil {
		t.Errorf("DeployKey.AddDeployKey returned error: %v", err)
	}

	timeLayout := "2006-01-02T15:04:05Z07:00"
	createdAt, err := time.Parse(timeLayout, "2015-08-29T12:44:31.550Z")
	if err != nil {
		t.Errorf("DeployKeys.ListAllDeployKeys returned an error while parsing time: %v", err)
	}

	want := &DeployKey{
		Title:     title,
		ID:        12,
		Key:       key,
		CanPush:   &canPush,
		CreatedAt: &createdAt,
	}
	if !reflect.DeepEqual(want, deployKey) {
		t.Errorf("DeployKeys.AddDeployKey returned %+v, want %+v", deployKey, want)
	}
}

func TestDeleteDeployKey(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/5/deploy_keys/13", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.DeployKeys.DeleteDeployKey(5, 13)
	if err != nil {
		t.Errorf("Deploykeys.DeleteDeployKey returned error: %v", err)
	}
}
