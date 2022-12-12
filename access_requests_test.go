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
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListProjectAccessRequests(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/access_requests", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			  "id": 1,
			  "username": "raymond_smith",
			  "name": "Raymond Smith",
			  "state": "active",
			  "created_at": "2012-10-22T14:13:35Z",
			  "requested_at": "2012-10-22T14:13:35Z"
			},
			{
			  "id": 2,
			  "username": "john_doe",
			  "name": "John Doe",
			  "state": "active",
			  "created_at": "2012-10-22T14:13:35Z",
			  "requested_at": "2012-10-22T14:13:35Z"
			}
		]`)
	})

	created := time.Date(2012, 10, 22, 14, 13, 35, 0, time.UTC)
	expected := []*AccessRequest{
		{
			ID:          1,
			Username:    "raymond_smith",
			Name:        "Raymond Smith",
			State:       "active",
			CreatedAt:   &created,
			RequestedAt: &created,
		},
		{
			ID:          2,
			Username:    "john_doe",
			Name:        "John Doe",
			State:       "active",
			CreatedAt:   &created,
			RequestedAt: &created,
		},
	}

	requests, resp, err := client.AccessRequests.ListProjectAccessRequests(1, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expected, requests)

	requests, resp, err = client.AccessRequests.ListProjectAccessRequests(1.5, nil)
	assert.EqualError(t, err, "invalid ID type 1.5, the ID must be an int or a string")
	assert.Nil(t, resp)
	assert.Nil(t, requests)

	requests, resp, err = client.AccessRequests.ListProjectAccessRequests(2, nil)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Nil(t, requests)

	requests, resp, err = client.AccessRequests.ListProjectAccessRequests(1, nil, errorOption)
	assert.EqualError(t, err, "RequestOptionFunc returns an error")
	assert.Nil(t, resp)
	assert.Nil(t, requests)
}

func TestListGroupAccessRequests(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/access_requests", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			  "id": 1,
			  "username": "raymond_smith",
			  "name": "Raymond Smith",
			  "state": "active",
			  "created_at": "2012-10-22T14:13:35Z",
			  "requested_at": "2012-10-22T14:13:35Z"
			},
			{
			  "id": 2,
			  "username": "john_doe",
			  "name": "John Doe",
			  "state": "active",
			  "created_at": "2012-10-22T14:13:35Z",
			  "requested_at": "2012-10-22T14:13:35Z"
			}
		]`)
	})

	created := time.Date(2012, 10, 22, 14, 13, 35, 0, time.UTC)
	expected := []*AccessRequest{
		{
			ID:          1,
			Username:    "raymond_smith",
			Name:        "Raymond Smith",
			State:       "active",
			CreatedAt:   &created,
			RequestedAt: &created,
		},
		{
			ID:          2,
			Username:    "john_doe",
			Name:        "John Doe",
			State:       "active",
			CreatedAt:   &created,
			RequestedAt: &created,
		},
	}

	requests, resp, err := client.AccessRequests.ListGroupAccessRequests(1, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expected, requests)

	requests, resp, err = client.AccessRequests.ListGroupAccessRequests(1.5, nil)
	assert.EqualError(t, err, "invalid ID type 1.5, the ID must be an int or a string")
	assert.Nil(t, resp)
	assert.Nil(t, requests)

	requests, resp, err = client.AccessRequests.ListGroupAccessRequests(2, nil)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Nil(t, requests)

	requests, resp, err = client.AccessRequests.ListGroupAccessRequests(1, nil, errorOption)
	assert.EqualError(t, err, "RequestOptionFunc returns an error")
	assert.Nil(t, resp)
	assert.Nil(t, requests)
}

func TestRequestProjectAccess(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/access_requests", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `{
				"id": 1,
				"username": "raymond_smith",
				"name": "Raymond Smith",
				"state": "active",
				"created_at": "2012-10-22T14:13:35Z",
				"requested_at": "2012-10-22T14:13:35Z"
			}`)
	})

	created := time.Date(2012, 10, 22, 14, 13, 35, 0, time.UTC)
	expected := &AccessRequest{
		ID:          1,
		Username:    "raymond_smith",
		Name:        "Raymond Smith",
		State:       "active",
		CreatedAt:   &created,
		RequestedAt: &created,
	}

	accessRequest, resp, err := client.AccessRequests.RequestProjectAccess(1, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expected, accessRequest)

	accessRequest, resp, err = client.AccessRequests.RequestProjectAccess(1.5, nil)
	assert.EqualError(t, err, "invalid ID type 1.5, the ID must be an int or a string")
	assert.Nil(t, resp)
	assert.Nil(t, accessRequest)

	accessRequest, resp, err = client.AccessRequests.RequestProjectAccess(2, nil)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Nil(t, accessRequest)

	accessRequest, resp, err = client.AccessRequests.RequestProjectAccess(1, nil, errorOption)
	assert.EqualError(t, err, "RequestOptionFunc returns an error")
	assert.Nil(t, resp)
	assert.Nil(t, accessRequest)
}

func TestRequestGroupAccess(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/access_requests", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `{
				"id": 1,
				"username": "raymond_smith",
				"name": "Raymond Smith",
				"state": "active",
				"created_at": "2012-10-22T14:13:35Z",
				"requested_at": "2012-10-22T14:13:35Z"
			}`)
	})

	created := time.Date(2012, 10, 22, 14, 13, 35, 0, time.UTC)
	expected := &AccessRequest{
		ID:          1,
		Username:    "raymond_smith",
		Name:        "Raymond Smith",
		State:       "active",
		CreatedAt:   &created,
		RequestedAt: &created,
	}

	accessRequest, resp, err := client.AccessRequests.RequestGroupAccess(1, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expected, accessRequest)

	accessRequest, resp, err = client.AccessRequests.RequestGroupAccess(1.5, nil)
	assert.EqualError(t, err, "invalid ID type 1.5, the ID must be an int or a string")
	assert.Nil(t, resp)
	assert.Nil(t, accessRequest)

	accessRequest, resp, err = client.AccessRequests.RequestGroupAccess(2, nil)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Nil(t, accessRequest)

	accessRequest, resp, err = client.AccessRequests.RequestGroupAccess(1, nil, errorOption)
	assert.EqualError(t, err, "RequestOptionFunc returns an error")
	assert.Nil(t, resp)
	assert.Nil(t, accessRequest)
}

func TestApproveProjectAccessRequest(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/access_requests/10/approve", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)

		var opt ApproveAccessRequestOptions
		err := json.NewDecoder(r.Body).Decode(&opt)
		assert.NoError(t, err)
		defer r.Body.Close()

		fmt.Fprintf(w, `{
				"id": 10,
				"username": "raymond_smith",
				"name": "Raymond Smith",
				"state": "active",
				"created_at": "2012-10-22T14:13:35Z",
				"requested_at": "2012-10-22T14:13:35Z",
				"access_level": %d
			}`,
			*opt.AccessLevel)
	})

	created := time.Date(2012, 10, 22, 14, 13, 35, 0, time.UTC)
	expected := &AccessRequest{
		ID:          10,
		Username:    "raymond_smith",
		Name:        "Raymond Smith",
		State:       "active",
		CreatedAt:   &created,
		RequestedAt: &created,
		AccessLevel: DeveloperPermissions,
	}

	opt := &ApproveAccessRequestOptions{
		AccessLevel: AccessLevel(DeveloperPermissions),
	}

	request, resp, err := client.AccessRequests.ApproveProjectAccessRequest(1, 10, opt)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expected, request)

	request, resp, err = client.AccessRequests.ApproveProjectAccessRequest(1.5, 10, opt)
	assert.EqualError(t, err, "invalid ID type 1.5, the ID must be an int or a string")
	assert.Nil(t, resp)
	assert.Nil(t, request)

	request, resp, err = client.AccessRequests.ApproveProjectAccessRequest(2, 10, opt)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Nil(t, request)

	request, resp, err = client.AccessRequests.ApproveProjectAccessRequest(1, 10, opt, errorOption)
	assert.EqualError(t, err, "RequestOptionFunc returns an error")
	assert.Nil(t, resp)
	assert.Nil(t, request)
}

func TestApproveGroupAccessRequest(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/access_requests/10/approve", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)

		var opt ApproveAccessRequestOptions
		err := json.NewDecoder(r.Body).Decode(&opt)
		assert.NoError(t, err)
		defer r.Body.Close()

		fmt.Fprintf(w, `{
				"id": 10,
				"username": "raymond_smith",
				"name": "Raymond Smith",
				"state": "active",
				"created_at": "2012-10-22T14:13:35Z",
				"requested_at": "2012-10-22T14:13:35Z",
				"access_level": %d
			}`,
			*opt.AccessLevel)
	})

	created := time.Date(2012, 10, 22, 14, 13, 35, 0, time.UTC)
	expected := &AccessRequest{
		ID:          10,
		Username:    "raymond_smith",
		Name:        "Raymond Smith",
		State:       "active",
		CreatedAt:   &created,
		RequestedAt: &created,
		AccessLevel: DeveloperPermissions,
	}

	opt := &ApproveAccessRequestOptions{
		AccessLevel: AccessLevel(DeveloperPermissions),
	}

	request, resp, err := client.AccessRequests.ApproveGroupAccessRequest(1, 10, opt)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expected, request)

	request, resp, err = client.AccessRequests.ApproveGroupAccessRequest(1.5, 10, opt)
	assert.EqualError(t, err, "invalid ID type 1.5, the ID must be an int or a string")
	assert.Nil(t, resp)
	assert.Nil(t, request)

	request, resp, err = client.AccessRequests.ApproveGroupAccessRequest(2, 10, opt)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Nil(t, request)

	request, resp, err = client.AccessRequests.ApproveGroupAccessRequest(1, 10, opt, errorOption)
	assert.EqualError(t, err, "RequestOptionFunc returns an error")
	assert.Nil(t, resp)
	assert.Nil(t, request)
}

func TestDenyProjectAccessRequest(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/access_requests/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.AccessRequests.DenyProjectAccessRequest(1, 10)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	resp, err = client.AccessRequests.DenyProjectAccessRequest(1.5, 10)
	assert.EqualError(t, err, "invalid ID type 1.5, the ID must be an int or a string")
	assert.Nil(t, resp)

	resp, err = client.AccessRequests.DenyProjectAccessRequest(2, 10)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

	resp, err = client.AccessRequests.DenyProjectAccessRequest(1, 10, errorOption)
	assert.EqualError(t, err, "RequestOptionFunc returns an error")
	assert.Nil(t, resp)
}

func TestDenyGroupAccessRequest(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/access_requests/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.AccessRequests.DenyGroupAccessRequest(1, 10)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	resp, err = client.AccessRequests.DenyGroupAccessRequest(1.5, 10)
	assert.EqualError(t, err, "invalid ID type 1.5, the ID must be an int or a string")
	assert.Nil(t, resp)

	resp, err = client.AccessRequests.DenyGroupAccessRequest(2, 10)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

	resp, err = client.AccessRequests.DenyGroupAccessRequest(1, 10, errorOption)
	assert.EqualError(t, err, "RequestOptionFunc returns an error")
	assert.Nil(t, resp)
}
