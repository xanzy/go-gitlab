package gitlab

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"
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

	rgs, _, err := client.ResourceGroup.GetAllResourceGroupsForAProject(1)
	if err != nil {
		log.Fatal(err)
	}

	date, _ := time.Parse(timeLayout, "2021-09-01T08:04:59.650Z")

	want := []*ResourceGroup{{
		ID:          3,
		Key:         "production",
		ProcessMode: "unordered",
		CreatedAt:   &date,
		UpdatedAt:   &date,
	}}

	if !reflect.DeepEqual(want, rgs) {
		t.Errorf("ResourceGroup.GetAllResourceGroupsForAProject returned %+v, want %+v", rgs, want)
	}
}

func TestResourceGroups_GetASpecificResourceGroup(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/resource_groups/3", func(w http.ResponseWriter, r *http.Request) {
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

	rgs, _, err := client.ResourceGroup.GetASpecificResourceGroup(1, 3)
	if err != nil {
		log.Fatal(err)
	}

	date, _ := time.Parse(timeLayout, "2021-09-01T08:04:59.650Z")

	want := []*ResourceGroup{{
		ID:          3,
		Key:         "production",
		ProcessMode: "unordered",
		CreatedAt:   &date,
		UpdatedAt:   &date,
	}}

	if !reflect.DeepEqual(want, rgs) {
		t.Errorf("ResourceGroup.GetAllResourceGroupsForAProject returned %+v, want %+v", rgs, want)
	}
}
