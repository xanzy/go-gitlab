// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWithHeader(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/without-header", func(w http.ResponseWriter, r *http.Request) {
		assert.Empty(t, r.Header.Get("X-CUSTOM-HEADER"))
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"X-CUSTOM-HEADER": %s`, r.Header.Get("X-CUSTOM-HEADER"))
	})
	mux.HandleFunc("/api/v4/with-header", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "randomtokenstring", r.Header.Get("X-CUSTOM-HEADER"))
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"X-CUSTOM-HEADER": %s`, r.Header.Get("X-CUSTOM-HEADER"))
	})

	// ensure that X-CUSTOM-HEADER hasn't been set at all
	req, err := client.NewRequest(http.MethodGet, "/without-header", nil, nil)
	assert.NoError(t, err)

	_, err = client.Do(req, nil)
	assert.NoError(t, err)

	// ensure that X-CUSTOM-HEADER is set for only one request
	req, err = client.NewRequest(
		http.MethodGet,
		"/with-header",
		nil,
		[]RequestOptionFunc{WithHeader("X-CUSTOM-HEADER", "randomtokenstring")},
	)
	assert.NoError(t, err)

	_, err = client.Do(req, nil)
	assert.NoError(t, err)

	req, err = client.NewRequest(http.MethodGet, "/without-header", nil, nil)
	assert.NoError(t, err)

	_, err = client.Do(req, nil)
	assert.NoError(t, err)

	// ensure that X-CUSTOM-HEADER is set for all client requests
	addr := client.BaseURL().String()
	client, err = NewClient("",
		// same base options as setup
		WithBaseURL(addr),
		// Disable backoff to speed up tests that expect errors.
		WithCustomBackoff(func(_, _ time.Duration, _ int, _ *http.Response) time.Duration {
			return 0
		}),
		// add client headers
		WithRequestOptions(WithHeader("X-CUSTOM-HEADER", "randomtokenstring")))
	assert.NoError(t, err)
	assert.NotNil(t, client)

	req, err = client.NewRequest(http.MethodGet, "/with-header", nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, "randomtokenstring", req.Header.Get("X-CUSTOM-HEADER"))

	_, err = client.Do(req, nil)
	assert.NoError(t, err)

	req, err = client.NewRequest(http.MethodGet, "/with-header", nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, "randomtokenstring", req.Header.Get("X-CUSTOM-HEADER"))

	_, err = client.Do(req, nil)
	assert.NoError(t, err)
}

func TestWithHeaders(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/without-headers", func(w http.ResponseWriter, r *http.Request) {
		assert.Empty(t, r.Header.Get("X-CUSTOM-HEADER-1"))
		assert.Empty(t, r.Header.Get("X-CUSTOM-HEADER-2"))
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/api/v4/with-headers", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "randomtokenstring", r.Header.Get("X-CUSTOM-HEADER-1"))
		assert.Equal(t, "randomtokenstring2", r.Header.Get("X-CUSTOM-HEADER-2"))
		w.WriteHeader(http.StatusOK)
	})

	headers := map[string]string{
		"X-CUSTOM-HEADER-1": "randomtokenstring",
		"X-CUSTOM-HEADER-2": "randomtokenstring2",
	}

	// ensure that X-CUSTOM-HEADER hasn't been set at all
	req, err := client.NewRequest(http.MethodGet, "/without-headers", nil, nil)
	assert.NoError(t, err)

	_, err = client.Do(req, nil)
	assert.NoError(t, err)

	// ensure that X-CUSTOM-HEADER is set for only one request
	req, err = client.NewRequest(
		http.MethodGet,
		"/with-headers",
		nil,
		[]RequestOptionFunc{WithHeaders(headers)},
	)
	assert.NoError(t, err)

	_, err = client.Do(req, nil)
	assert.NoError(t, err)

	req, err = client.NewRequest(http.MethodGet, "/without-headers", nil, nil)
	assert.NoError(t, err)

	_, err = client.Do(req, nil)
	assert.NoError(t, err)

	// ensure that X-CUSTOM-HEADER is set for all client requests
	addr := client.BaseURL().String()
	client, err = NewClient("",
		// same base options as setup
		WithBaseURL(addr),
		// Disable backoff to speed up tests that expect errors.
		WithCustomBackoff(func(_, _ time.Duration, _ int, _ *http.Response) time.Duration {
			return 0
		}),
		// add client headers
		WithRequestOptions(WithHeaders(headers)),
	)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	req, err = client.NewRequest(http.MethodGet, "/with-headers", nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, "randomtokenstring", req.Header.Get("X-CUSTOM-HEADER-1"))

	_, err = client.Do(req, nil)
	assert.NoError(t, err)

	req, err = client.NewRequest(http.MethodGet, "/with-headers", nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, "randomtokenstring", req.Header.Get("X-CUSTOM-HEADER-1"))

	_, err = client.Do(req, nil)
	assert.NoError(t, err)
}
