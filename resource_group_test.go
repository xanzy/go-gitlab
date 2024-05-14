package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestResourceGroups_GetAllResourceGroupsForAProject(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/resource_groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
			{
				"id": 3,
				"key": "production",
				"process_mode": "unordered",
				"created_at": "2021-09-01T08:04:59.650Z",
				"updated_at": "2021-09-01T08:04:59.650Z"
			  }
		]`)
	})

	date, _ := time.Parse(timeLayout, "2021-09-01T08:04:59.650Z")

	want := []*ResourceGroup{{
		ID:          3,
		Key:         "production",
		ProcessMode: "unordered",
		CreatedAt:   &date,
		UpdatedAt:   &date,
	}}

	rgs, resp, err := client.ResourceGroup.GetAllResourceGroupsForAProject(1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.ElementsMatch(t, want, rgs)

	rgs, resp, err = client.ResourceGroup.GetAllResourceGroupsForAProject(1.01)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, rgs)

	rgs, resp, err = client.ResourceGroup.GetAllResourceGroupsForAProject(1, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, rgs)

	rgs, resp, err = client.ResourceGroup.GetAllResourceGroupsForAProject(2)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
	require.Nil(t, rgs)
}

func TestResourceGroups_GetASpecificResourceGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/resource_groups/production", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
			{
				"id": 3,
				"key": "production",
				"process_mode": "unordered",
				"created_at": "2021-09-01T08:04:59.650Z",
				"updated_at": "2021-09-01T08:04:59.650Z"
			  }
		`)
	})

	date, _ := time.Parse(timeLayout, "2021-09-01T08:04:59.650Z")

	want := &ResourceGroup{
		ID:          3,
		Key:         "production",
		ProcessMode: "unordered",
		CreatedAt:   &date,
		UpdatedAt:   &date,
	}

	rg, resp, err := client.ResourceGroup.GetASpecificResourceGroup(1, "production")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, rg)

	rg, resp, err = client.ResourceGroup.GetASpecificResourceGroup(1.01, "production")
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, rg)

	rg, resp, err = client.ResourceGroup.GetASpecificResourceGroup(1, "production", errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, rg)

	rg, resp, err = client.ResourceGroup.GetASpecificResourceGroup(2, "production")
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
	require.Nil(t, rg)
}

func TestResourceGroups_ListUpcomingJobsForASpecificResourceGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/resource_groups/production/upcoming_jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
			{
				"id": 6
			}
		]`)
	})

	want := []*Job{{
		ID: 6,
	}}

	jobs, resp, err := client.ResourceGroup.ListUpcomingJobsForASpecificResourceGroup(1, "production")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.ElementsMatch(t, want, jobs)

	jobs, resp, err = client.ResourceGroup.ListUpcomingJobsForASpecificResourceGroup(1.01, "production")
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, jobs)

	jobs, resp, err = client.ResourceGroup.ListUpcomingJobsForASpecificResourceGroup(1, "production", errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, jobs)

	jobs, resp, err = client.ResourceGroup.ListUpcomingJobsForASpecificResourceGroup(2, "production")
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
	require.Nil(t, jobs)
}

func TestResourceGroup_EditAnExistingResourceGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/resource_groups/production", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `
			{
				"id": 3,
				"key": "production",
				"process_mode": "unordered",
				"created_at": "2021-09-01T08:04:59.650Z",
				"updated_at": "2021-09-01T08:04:59.650Z"
			}
		`)
	})

	date, _ := time.Parse(timeLayout, "2021-09-01T08:04:59.650Z")

	want := &ResourceGroup{
		ID:          3,
		Key:         "production",
		ProcessMode: "unordered",
		CreatedAt:   &date,
		UpdatedAt:   &date,
	}

	opts := &EditAnExistingResourceGroupOptions{ProcessMode: Ptr(OldestFirst)}

	rg, resp, err := client.ResourceGroup.EditAnExistingResourceGroup(1, "production", opts)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, rg)

	rg, resp, err = client.ResourceGroup.EditAnExistingResourceGroup(1.01, "production", opts)
	require.EqualError(t, err, "invalid ID type 1.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, rg)

	rg, resp, err = client.ResourceGroup.EditAnExistingResourceGroup(1, "production", opts, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, rg)

	rg, resp, err = client.ResourceGroup.EditAnExistingResourceGroup(2, "production", opts)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
	require.Nil(t, rg)
}
