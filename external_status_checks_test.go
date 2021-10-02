package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListExternalStatusChecks(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/1/status_checks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, exampleStatusChecks)
	})

	statusChecks, _, err := client.ExternalStatusChecks.ListExternalStatusChecks(1, 1, nil)
	if err != nil {
		t.Fatalf("ExternalStatusChecks.ListExternalStatusChecks returns an error: %v", err)
	}

	expectedStatusChecks := []*StatusCheck{
		{
			ID:          2,
			Name:        "Rule 1",
			ExternalURL: "https://gitlab.com/test-endpoint",
			Status:      "approved",
		},
		{
			ID:          1,
			Name:        "Rule 2",
			ExternalURL: "https://gitlab.com/test-endpoint-2",
			Status:      "pending",
		},
	}

	assert.Equal(t, expectedStatusChecks, statusChecks)
}
