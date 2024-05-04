package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDORAMetrics_GetProjectDORAMetrics(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/dora/metrics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		query := r.URL.Query()
		for k, v := range map[string]string{
			"metric":     "deployment_frequency",
			"start_date": "2021-03-01",
			"end_date":   "2021-03-08",
		} {
			if query.Get(k) != v {
				t.Errorf("Query parameter %s: %s, want %s", k, query.Get(k), v)
			}
		}

		fmt.Fprint(w, `
			[
				{ "date": "2021-03-01", "value": 3 },
				{ "date": "2021-03-02", "value": 6 },
				{ "date": "2021-03-03", "value": 0 },
				{ "date": "2021-03-04", "value": 0 },
				{ "date": "2021-03-05", "value": 0 },
				{ "date": "2021-03-06", "value": 0 },
				{ "date": "2021-03-07", "value": 0 },
				{ "date": "2021-03-08", "value": 4 }
			]
		`)
	})

	want := []DORAMetric{
		{Date: "2021-03-01", Value: 3},
		{Date: "2021-03-02", Value: 6},
		{Date: "2021-03-03", Value: 0},
		{Date: "2021-03-04", Value: 0},
		{Date: "2021-03-05", Value: 0},
		{Date: "2021-03-06", Value: 0},
		{Date: "2021-03-07", Value: 0},
		{Date: "2021-03-08", Value: 4},
	}

	startDate := ISOTime(time.Date(2021, time.March, 1, 0, 0, 0, 0, time.UTC))
	endDate := ISOTime(time.Date(2021, time.March, 8, 0, 0, 0, 0, time.UTC))

	d, resp, err := client.DORAMetrics.GetProjectDORAMetrics(1, GetDORAMetricsOptions{
		Metric:    DORAMetricDeploymentFrequency,
		StartDate: &startDate,
		EndDate:   &endDate,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)
}

func TestDORAMetrics_GetGroupDORAMetrics(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/dora/metrics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		query := r.URL.Query()
		for k, v := range map[string]string{
			"metric":     "deployment_frequency",
			"start_date": "2021-03-01",
			"end_date":   "2021-03-08",
		} {
			if query.Get(k) != v {
				t.Errorf("Query parameter %s: %s, want %s", k, query.Get(k), v)
			}
		}

		fmt.Fprint(w, `
			[
				{ "date": "2021-03-01", "value": 3 },
				{ "date": "2021-03-02", "value": 6 },
				{ "date": "2021-03-03", "value": 0 },
				{ "date": "2021-03-04", "value": 0 },
				{ "date": "2021-03-05", "value": 0 },
				{ "date": "2021-03-06", "value": 0 },
				{ "date": "2021-03-07", "value": 0 },
				{ "date": "2021-03-08", "value": 4 }
			]
		`)
	})

	want := []DORAMetric{
		{Date: "2021-03-01", Value: 3},
		{Date: "2021-03-02", Value: 6},
		{Date: "2021-03-03", Value: 0},
		{Date: "2021-03-04", Value: 0},
		{Date: "2021-03-05", Value: 0},
		{Date: "2021-03-06", Value: 0},
		{Date: "2021-03-07", Value: 0},
		{Date: "2021-03-08", Value: 4},
	}

	startDate := ISOTime(time.Date(2021, time.March, 1, 0, 0, 0, 0, time.UTC))
	endDate := ISOTime(time.Date(2021, time.March, 8, 0, 0, 0, 0, time.UTC))

	d, resp, err := client.DORAMetrics.GetGroupDORAMetrics(1, GetDORAMetricsOptions{
		Metric:    DORAMetricDeploymentFrequency,
		StartDate: &startDate,
		EndDate:   &endDate,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, d)
}
