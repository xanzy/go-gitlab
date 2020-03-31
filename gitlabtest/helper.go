// Test helper file
// Cf. https://stackoverflow.com/questions/56404355/how-do-i-package-golang-test-helper-code/56412373#56412373

package gitlabtest

import (
	"net/http"
	"net/http/httptest"

	"github.com/prytoegrian/go-gitlab"
)

// setup duplicates gitlab.setup for testability purpose
func Setup() (*http.ServeMux, *httptest.Server, *gitlab.Client) {
	// mux is the HTTP request multiplexer used with the test server.
	mux := http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// client is the Gitlab client being tested.
	client := gitlab.NewClient(nil, "")
	client.SetBaseURL(server.URL)

	return mux, server, client
}

// teardown duplicates gitlab.teardown for testability purpose
func Teardown(server *httptest.Server) {
	server.Close()
}

// testMethod duplicates gitlab.testMethod for testability purpose
// func testMethod(t *testing.T, r *http.Request, want string) {
// 	if got := r.Method; got != want {
// 		t.Errorf("Request method: %s, want %s", got, want)
// 	}
// }
