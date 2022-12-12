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
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/deploy_keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
		{
			"id": 1,
			"title": "Public key",
			"key": "ssh-rsa AAAA...",
			"fingerprint": "7f:72:08:7d:0e:47:48:ec:37:79:b2:76:68:b5:87:65",
			"created_at": "2013-10-02T10:12:29Z",
			"projects_with_write_access": [
			{
				"id": 73,
				"description": null,
				"name": "project2",
				"name_with_namespace": "Sidney Jones / project2",
				"path": "project2",
				"path_with_namespace": "sidney_jones/project2",
				"created_at": "2021-10-25T18:33:17.550Z"
			},
			{
				"id": 74,
				"description": null,
				"name": "project3",
				"name_with_namespace": "Sidney Jones / project3",
				"path": "project3",
				"path_with_namespace": "sidney_jones/project3",
				"created_at": "2021-10-25T18:33:17.666Z"
			}
			]
		},
			{
				"id": 3,
				"title": "Another Public key",
				"key": "ssh-rsa AAAA...",
				"fingerprint": "64:d3:73:d4:83:70:ab:41:96:68:d5:3d:a5:b0:34:ea",
				"created_at": "2013-10-02T11:12:29Z",
				"projects_with_write_access": []
			}
		  ]`)
	})

	deployKeys, _, err := client.DeployKeys.ListAllDeployKeys(&ListInstanceDeployKeysOptions{})
	if err != nil {
		t.Errorf("DeployKeys.ListAllDeployKeys returned error: %v", err)
	}

	createdAtKey1, _ := time.Parse(timeLayout, "2013-10-02T10:12:29Z")
	createdAtKey1Enable1, _ := time.Parse(timeLayout, "2021-10-25T18:33:17.550Z")
	createdAtKey1Enable2, _ := time.Parse(timeLayout, "2021-10-25T18:33:17.666Z")
	createdAtKey2, _ := time.Parse(timeLayout, "2013-10-02T11:12:29Z")

	want := []*InstanceDeployKey{
		{
			ID:          1,
			Title:       "Public key",
			Key:         "ssh-rsa AAAA...",
			CreatedAt:   &createdAtKey1,
			Fingerprint: "7f:72:08:7d:0e:47:48:ec:37:79:b2:76:68:b5:87:65",
			ProjectsWithWriteAccess: []*DeployKeyProject{
				{
					ID:                73,
					Description:       "",
					Name:              "project2",
					NameWithNamespace: "Sidney Jones / project2",
					Path:              "project2",
					PathWithNamespace: "sidney_jones/project2",
					CreatedAt:         &createdAtKey1Enable1,
				},
				{
					ID:                74,
					Description:       "",
					Name:              "project3",
					NameWithNamespace: "Sidney Jones / project3",
					Path:              "project3",
					PathWithNamespace: "sidney_jones/project3",
					CreatedAt:         &createdAtKey1Enable2,
				},
			},
		},
		{
			ID:                      3,
			Title:                   "Another Public key",
			Key:                     "ssh-rsa AAAA...",
			Fingerprint:             "64:d3:73:d4:83:70:ab:41:96:68:d5:3d:a5:b0:34:ea",
			CreatedAt:               &createdAtKey2,
			ProjectsWithWriteAccess: []*DeployKeyProject{},
		},
	}
	if !reflect.DeepEqual(want, deployKeys) {
		t.Errorf("DeployKeys.ListAllDeployKeys returned %+v, want %+v", deployKeys, want)
	}
}

func TestListProjectDeployKeys(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/deploy_keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			  "id": 1,
			  "title": "Public key",
			  "key": "ssh-rsa AAAA...",
			  "created_at": "2013-10-02T10:12:29Z",
			  "can_push": false
			},
			{
			  "id": 3,
			  "title": "Another Public key",
			  "key": "ssh-rsa AAAA...",
			  "created_at": "2013-10-02T11:12:29Z",
			  "can_push": false
			}
		  ]`)
	})

	deployKeys, _, err := client.DeployKeys.ListProjectDeployKeys(5, &ListProjectDeployKeysOptions{})
	if err != nil {
		t.Errorf("DeployKeys.ListProjectDeployKeys returned error: %v", err)
	}

	createdAt, err := time.Parse(timeLayout, "2013-10-02T10:12:29Z")
	if err != nil {
		t.Errorf("DeployKeys.ListAllDeployKeys returned an error while parsing time: %v", err)
	}

	createdAt2, err := time.Parse(timeLayout, "2013-10-02T11:12:29Z")
	if err != nil {
		t.Errorf("DeployKeys.ListAllDeployKeys returned an error while parsing time: %v", err)
	}

	want := []*ProjectDeployKey{
		{
			ID:        1,
			Title:     "Public key",
			Key:       "ssh-rsa AAAA...",
			CreatedAt: &createdAt,
			CanPush:   false,
		},
		{
			ID:        3,
			Title:     "Another Public key",
			Key:       "ssh-rsa AAAA...",
			CreatedAt: &createdAt2,
			CanPush:   false,
		},
	}
	if !reflect.DeepEqual(want, deployKeys) {
		t.Errorf("DeployKeys.ListProjectDeployKeys returned %+v, want %+v", deployKeys, want)
	}
}

func TestGetDeployKey(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/deploy_keys/11", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
			"id": 1,
			"title": "Public key",
			"key": "ssh-rsa AAAA...",
			"created_at": "2013-10-02T10:12:29Z",
			"can_push": false
		  }`)
	})

	deployKey, _, err := client.DeployKeys.GetDeployKey(5, 11)
	if err != nil {
		t.Errorf("DeployKeys.GetDeployKey returned error: %v", err)
	}

	createdAt, err := time.Parse(timeLayout, "2013-10-02T10:12:29Z")
	if err != nil {
		t.Errorf("DeployKeys.ListAllDeployKeys returned an error while parsing time: %v", err)
	}

	want := &ProjectDeployKey{
		ID:        1,
		Title:     "Public key",
		Key:       "ssh-rsa AAAA...",
		CreatedAt: &createdAt,
		CanPush:   false,
	}
	if !reflect.DeepEqual(want, deployKey) {
		t.Errorf("DeployKeys.GetDeployKey returned %+v, want %+v", deployKey, want)
	}
}

func TestAddDeployKey(t *testing.T) {
	mux, client := setup(t)

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

	opt := &AddDeployKeyOptions{
		Key:     String("ssh-rsa AAAA..."),
		Title:   String("My deploy key"),
		CanPush: Bool(true),
	}
	deployKey, _, err := client.DeployKeys.AddDeployKey(5, opt)
	if err != nil {
		t.Errorf("DeployKey.AddDeployKey returned error: %v", err)
	}

	createdAt, err := time.Parse(timeLayout, "2015-08-29T12:44:31.550Z")
	if err != nil {
		t.Errorf("DeployKeys.ListAllDeployKeys returned an error while parsing time: %v", err)
	}

	want := &ProjectDeployKey{
		Title:     *String("My deploy key"),
		ID:        12,
		Key:       *String("ssh-rsa AAAA..."),
		CreatedAt: &createdAt,
		CanPush:   true,
	}
	if !reflect.DeepEqual(want, deployKey) {
		t.Errorf("DeployKeys.AddDeployKey returned %+v, want %+v", deployKey, want)
	}
}

func TestDeleteDeployKey(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/deploy_keys/13", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.DeployKeys.DeleteDeployKey(5, 13)
	if err != nil {
		t.Errorf("Deploykeys.DeleteDeployKey returned error: %v", err)
	}
}

func TestEnableDeployKey(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/deploy_keys/13/enable", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `{
			"key" : "ssh-rsa AAAA...",
			"id" : 12,
			"title" : "My deploy key",
			"created_at" : "2015-08-29T12:44:31.550Z"
		 }`)
	})

	deployKey, _, err := client.DeployKeys.EnableDeployKey(5, 13)
	if err != nil {
		t.Errorf("DeployKeys.EnableDeployKey returned error: %v", err)
	}

	createdAt, err := time.Parse(timeLayout, "2015-08-29T12:44:31.550Z")
	if err != nil {
		t.Errorf("DeployKeys.ListAllDeployKeys returned an error while parsing time: %v", err)
	}

	want := &ProjectDeployKey{
		ID:        12,
		Title:     "My deploy key",
		Key:       "ssh-rsa AAAA...",
		CreatedAt: &createdAt,
	}
	if !reflect.DeepEqual(want, deployKey) {
		t.Errorf("DeployKeys.EnableDeployKey returned %+v, want %+v", deployKey, want)
	}
}

func TestUpdateDeployKey(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/deploy_keys/11", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `{
			"id": 11,
			"title": "New deploy key",
			"key": "ssh-rsa AAAA...",
			"created_at": "2015-08-29T12:44:31.550Z",
			"can_push": true
		 }`)
	})

	opt := &UpdateDeployKeyOptions{
		Title:   String("New deploy key"),
		CanPush: Bool(true),
	}
	deployKey, _, err := client.DeployKeys.UpdateDeployKey(5, 11, opt)
	if err != nil {
		t.Errorf("DeployKeys.UpdateDeployKey returned error: %v", err)
	}

	createdAt, err := time.Parse(timeLayout, "2015-08-29T12:44:31.550Z")
	if err != nil {
		t.Errorf("DeployKeys.ListAllDeployKeys returned an error while parsing time: %v", err)
	}

	want := &ProjectDeployKey{
		ID:        11,
		Title:     *String("New deploy key"),
		Key:       "ssh-rsa AAAA...",
		CreatedAt: &createdAt,
		CanPush:   true,
	}
	if !reflect.DeepEqual(want, deployKey) {
		t.Errorf("DeployKeys.UpdateDeployKey returned %+v, want %+v", deployKey, want)
	}
}
