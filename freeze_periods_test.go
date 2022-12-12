package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFreezePeriodsService_ListFreezePeriods(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/19/freeze_periods", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			   {
				  "id":1,
				  "freeze_start":"0 23 * * 5",
				  "freeze_end":"0 8 * * 1",
				  "cron_timezone":"UTC"
			   }
			]
		`)
	})

	want := []*FreezePeriod{
		{
			ID:           1,
			FreezeStart:  "0 23 * * 5",
			FreezeEnd:    "0 8 * * 1",
			CronTimezone: "UTC",
		},
	}

	fps, resp, err := client.FreezePeriods.ListFreezePeriods(19, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, fps)

	fps, resp, err = client.FreezePeriods.ListFreezePeriods(19.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 19.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, fps)

	fps, resp, err = client.FreezePeriods.ListFreezePeriods(19, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, fps)

	fps, resp, err = client.FreezePeriods.ListFreezePeriods(3, nil, nil)
	require.Error(t, err)
	require.Nil(t, fps)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestFreezePeriodsService_GetFreezePeriod(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/19/freeze_periods/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		   {
			  "id":1,
			  "freeze_start":"0 23 * * 5",
			  "freeze_end":"0 8 * * 1",
			  "cron_timezone":"UTC"
		   }
		`)
	})

	want := &FreezePeriod{
		ID:           1,
		FreezeStart:  "0 23 * * 5",
		FreezeEnd:    "0 8 * * 1",
		CronTimezone: "UTC",
	}

	fp, resp, err := client.FreezePeriods.GetFreezePeriod(19, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, fp)

	fp, resp, err = client.FreezePeriods.GetFreezePeriod(19.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 19.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, fp)

	fp, resp, err = client.FreezePeriods.GetFreezePeriod(19, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, fp)

	fp, resp, err = client.FreezePeriods.GetFreezePeriod(3, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, fp)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestFreezePeriodsService_CreateFreezePeriodOptions(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/19/freeze_periods", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		   {
			  "id":1,
			  "freeze_start":"0 23 * * 5",
			  "freeze_end":"0 8 * * 1",
			  "cron_timezone":"UTC"
		   }
		`)
	})

	want := &FreezePeriod{
		ID:           1,
		FreezeStart:  "0 23 * * 5",
		FreezeEnd:    "0 8 * * 1",
		CronTimezone: "UTC",
	}

	fp, resp, err := client.FreezePeriods.CreateFreezePeriodOptions(19, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, fp)

	fp, resp, err = client.FreezePeriods.CreateFreezePeriodOptions(19.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 19.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, fp)

	fp, resp, err = client.FreezePeriods.CreateFreezePeriodOptions(19, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, fp)

	fp, resp, err = client.FreezePeriods.CreateFreezePeriodOptions(3, nil, nil)
	require.Error(t, err)
	require.Nil(t, fp)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestFreezePeriodsService_UpdateFreezePeriodOptions(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/19/freeze_periods/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
		   {
			  "id":1,
			  "freeze_start":"0 23 * * 5",
			  "freeze_end":"0 8 * * 1",
			  "cron_timezone":"UTC"
		   }
		`)
	})

	want := &FreezePeriod{
		ID:           1,
		FreezeStart:  "0 23 * * 5",
		FreezeEnd:    "0 8 * * 1",
		CronTimezone: "UTC",
	}

	fp, resp, err := client.FreezePeriods.UpdateFreezePeriodOptions(19, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, fp)

	fp, resp, err = client.FreezePeriods.UpdateFreezePeriodOptions(19.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 19.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, fp)

	fp, resp, err = client.FreezePeriods.UpdateFreezePeriodOptions(19, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, fp)

	fp, resp, err = client.FreezePeriods.UpdateFreezePeriodOptions(3, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, fp)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestFreezePeriodsService_DeleteFreezePeriod(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/19/freeze_periods/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.FreezePeriods.DeleteFreezePeriod(19, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.FreezePeriods.DeleteFreezePeriod(19.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 19.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.FreezePeriods.DeleteFreezePeriod(19, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.FreezePeriods.DeleteFreezePeriod(3, 1, nil, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
