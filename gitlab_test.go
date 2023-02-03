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
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

var timeLayout = "2006-01-02T15:04:05Z07:00"

// setup sets up a test HTTP server along with a gitlab.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup(t *testing.T) (*http.ServeMux, *Client) {
	// mux is the HTTP request multiplexer used with the test server.
	mux := http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	// client is the Gitlab client being tested.
	client, err := NewClient("",
		WithBaseURL(server.URL),
		// Disable backoff to speed up tests that expect errors.
		WithCustomBackoff(func(_, _ time.Duration, _ int, _ *http.Response) time.Duration {
			return 0
		}),
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	return mux, client
}

func testURL(t *testing.T, r *http.Request, want string) {
	if got := r.RequestURI; got != want {
		t.Errorf("Request url: %+v, want %s", got, want)
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %s, want %s", got, want)
	}
}

func testBody(t *testing.T, r *http.Request, want string) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	if err != nil {
		t.Fatalf("Failed to Read Body: %v", err)
	}

	if got := buffer.String(); got != want {
		t.Errorf("Request body: %s, want %s", got, want)
	}
}

func testParams(t *testing.T, r *http.Request, want string) {
	if got := r.URL.RawQuery; got != want {
		t.Errorf("Request query: %s, want %s", got, want)
	}
}

func mustWriteHTTPResponse(t *testing.T, w io.Writer, fixturePath string) {
	f, err := os.Open(fixturePath)
	if err != nil {
		t.Fatalf("error opening fixture file: %v", err)
	}

	if _, err = io.Copy(w, f); err != nil {
		t.Fatalf("error writing response: %v", err)
	}
}

func errorOption(*retryablehttp.Request) error {
	return errors.New("RequestOptionFunc returns an error")
}

func TestNewClient(t *testing.T) {
	c, err := NewClient("")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	expectedBaseURL := defaultBaseURL + apiVersionPath

	if c.BaseURL().String() != expectedBaseURL {
		t.Errorf("NewClient BaseURL is %s, want %s", c.BaseURL().String(), expectedBaseURL)
	}
	if c.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent is %s, want %s", c.UserAgent, userAgent)
	}
}

func TestCheckResponse(t *testing.T) {
	c, err := NewClient("")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	req, err := c.NewRequest(http.MethodGet, "test", nil, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp := &http.Response{
		Request:    req.Request,
		StatusCode: http.StatusBadRequest,
		Body: io.NopCloser(strings.NewReader(`
		{
			"message": {
				"prop1": [
					"message 1",
					"message 2"
				],
				"prop2":[
					"message 3"
				],
				"embed1": {
					"prop3": [
						"msg 1",
						"msg2"
					]
				},
				"embed2": {
					"prop4": [
						"some msg"
					]
				}
			},
			"error": "message 1"
		}`)),
	}

	errResp := CheckResponse(resp)
	if errResp == nil {
		t.Fatal("Expected error response.")
	}

	want := "GET https://gitlab.com/api/v4/test: 400 {error: message 1}, {message: {embed1: {prop3: [msg 1, msg2]}}, {embed2: {prop4: [some msg]}}, {prop1: [message 1, message 2]}, {prop2: [message 3]}}"

	if errResp.Error() != want {
		t.Errorf("Expected error: %s, got %s", want, errResp.Error())
	}
}

func TestCheckResponseOnUnknownErrorFormat(t *testing.T) {
	c, err := NewClient("")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	req, err := c.NewRequest(http.MethodGet, "test", nil, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp := &http.Response{
		Request:    req.Request,
		StatusCode: http.StatusBadRequest,
		Body:       io.NopCloser(strings.NewReader("some error message but not JSON")),
	}

	errResp := CheckResponse(resp)
	if errResp == nil {
		t.Fatal("Expected error response.")
	}

	want := "GET https://gitlab.com/api/v4/test: 400 failed to parse unknown error format: some error message but not JSON"

	if errResp.Error() != want {
		t.Errorf("Expected error: %s, got %s", want, errResp.Error())
	}
}

func TestRequestWithContext(t *testing.T) {
	c, err := NewClient("")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	req, err := c.NewRequest(http.MethodGet, "test", nil, []RequestOptionFunc{WithContext(ctx)})
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	defer cancel()

	if req.Context() != ctx {
		t.Fatal("Context was not set correctly")
	}
}

func loadFixture(filePath string) []byte {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func TestPathEscape(t *testing.T) {
	want := "diaspora%2Fdiaspora"
	got := PathEscape("diaspora/diaspora")
	if want != got {
		t.Errorf("Expected: %s, got %s", want, got)
	}
}
