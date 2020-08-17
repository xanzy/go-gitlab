package gitlab

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRunPipelineSchedule(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/pipeline_schedules/1/play", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"message": "201 Created"}`)
	})

	res, err := client.PipelineSchedules.RunPipelineSchedule(1, 1)

	if err != nil {
		t.Errorf("PipelineTriggers.RunPipelineTrigger returned error: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("PipelineSchedules.RunPipelineSchedule returned status %v, want %v", res.StatusCode, http.StatusCreated)
	}
}
