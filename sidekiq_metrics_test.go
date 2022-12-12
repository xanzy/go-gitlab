package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSidekiqService_GetQueueMetrics(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/sidekiq/queue_metrics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{"queues": {"default": {"backlog": 0,"latency": 0}}}`)
	})

	qm, _, err := client.Sidekiq.GetQueueMetrics()
	require.NoError(t, err)

	want := &QueueMetrics{Queues: map[string]struct {
		Backlog int `json:"backlog"`
		Latency int `json:"latency"`
	}{"default": {Backlog: 0, Latency: 0}}}
	require.Equal(t, want, qm)
}

func TestSidekiqService_GetProcessMetrics(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/sidekiq/process_metrics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"processes": [{"hostname": "gitlab.example.com","pid": 5649,"tag": "gitlab","concurrency": 25,"busy": 0}]}`)
	})

	pm, _, err := client.Sidekiq.GetProcessMetrics()
	require.NoError(t, err)

	want := &ProcessMetrics{[]struct {
		Hostname    string     `json:"hostname"`
		Pid         int        `json:"pid"`
		Tag         string     `json:"tag"`
		StartedAt   *time.Time `json:"started_at"`
		Queues      []string   `json:"queues"`
		Labels      []string   `json:"labels"`
		Concurrency int        `json:"concurrency"`
		Busy        int        `json:"busy"`
	}{{Hostname: "gitlab.example.com", Pid: 5649, Tag: "gitlab", Concurrency: 25, Busy: 0}}}
	require.Equal(t, want, pm)
}

func TestSidekiqService_GetJobStats(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/sidekiq/job_stats", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"jobs": {"processed": 2,"failed": 0,"enqueued": 0}}`)
	})

	js, _, err := client.Sidekiq.GetJobStats()
	require.NoError(t, err)

	want := &JobStats{struct {
		Processed int `json:"processed"`
		Failed    int `json:"failed"`
		Enqueued  int `json:"enqueued"`
	}(struct {
		Processed int
		Failed    int
		Enqueued  int
	}{Processed: 2, Failed: 0, Enqueued: 0})}
	require.Equal(t, want, js)
}

func TestSidekiqService_GetCompoundMetrics(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/sidekiq/compound_metrics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_compound_metrics.json")
	})

	cm, _, err := client.Sidekiq.GetCompoundMetrics()
	require.NoError(t, err)

	want := &CompoundMetrics{
		QueueMetrics: QueueMetrics{Queues: map[string]struct {
			Backlog int `json:"backlog"`
			Latency int `json:"latency"`
		}{"default": {
			Backlog: 0,
			Latency: 0,
		}}},
		ProcessMetrics: ProcessMetrics{Processes: []struct {
			Hostname    string     `json:"hostname"`
			Pid         int        `json:"pid"`
			Tag         string     `json:"tag"`
			StartedAt   *time.Time `json:"started_at"`
			Queues      []string   `json:"queues"`
			Labels      []string   `json:"labels"`
			Concurrency int        `json:"concurrency"`
			Busy        int        `json:"busy"`
		}{{Hostname: "gitlab.example.com", Pid: 5649, Tag: "gitlab", Concurrency: 25, Busy: 0}}},
		JobStats: JobStats{Jobs: struct {
			Processed int `json:"processed"`
			Failed    int `json:"failed"`
			Enqueued  int `json:"enqueued"`
		}(struct {
			Processed int
			Failed    int
			Enqueued  int
		}{Processed: 2, Failed: 0, Enqueued: 0})},
	}
	require.Equal(t, want, cm)
}
