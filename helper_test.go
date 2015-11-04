package gitlab

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func Stub(filename string) (*httptest.Server, *Client) {
	stub, _ := ioutil.ReadFile(filename)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(stub))
	}))
	client := NewClient(nil, "")
	client.SetBaseURL(ts.URL)
	return ts, client
}
