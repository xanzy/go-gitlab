package gitlab

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

// setup sets up a test HTTP server along with a gitlab.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (*http.ServeMux, *httptest.Server, *Client) {
	// mux is the HTTP request multiplexer used with the test server.
	mux := http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// client is the Gitlab client being tested.
	client := NewClient(nil, "")
	client.SetBaseURL(server.URL)

	return mux, server, client
}

// teardown closes the test HTTP server.
func teardown(server *httptest.Server) {
	server.Close()
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

type values map[string]string

func testFormValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Add(k, v)
	}

	err := r.ParseForm()
	if err != nil {
		t.Errorf("Error parsing form: %v", err)
	}

	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters: %v, want %v", got, want)
	}
}

func testJSONBody(t *testing.T, r *http.Request, want values) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}

	var got values
	err = json.Unmarshal(b, &got)
	if err != nil {
		t.Errorf("Error unmarshalling request body: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters: %v, want %v", got, want)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil, "")

	if c.BaseURL().String() != defaultBaseURL {
		t.Errorf("NewClient BaseURL is %s, want %s", c.BaseURL().String(), defaultBaseURL)
	}
	if c.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent is %s, want %s", c.UserAgent, userAgent)
	}
}

func TestCheckResponse(t *testing.T) {
	req, err := NewClient(nil, "").NewRequest("GET", "test", nil, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp := &http.Response{
		Request:    req,
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(strings.NewReader(`
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

	want := "GET https://gitlab.com/api/v3/test: 400 {error: message 1}, {message: {embed1: {prop3: [msg 1, msg2]}}, {embed2: {prop4: [some msg]}}, {prop1: [message 1, message 2]}, {prop2: [message 3]}}"

	if errResp.Error() != want {
		t.Errorf("Expected error: %s, got %s", want, errResp.Error())
	}
}

func TestRequestWithContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), interface{}("myKey"), interface{}("myValue"))
	req, err := NewClient(nil, "").NewRequest("GET", "test", nil, []OptionFunc{WithContext(ctx)})
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	if req.Context() != ctx {
		t.Fatal("Context was not set correctly")
	}
}
