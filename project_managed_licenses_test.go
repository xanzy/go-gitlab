//
// Copyright 2021, Andrea Perizzato
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"net/http"
	"reflect"
	"testing"
)

func TestListProjectManagedLicenses(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/managed_licenses", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_project_managed_licenses.json")
	})

	licenses, _, err := client.ProjectManagedLicenses.ListManagedLicenses(1)
	if err != nil {
		t.Errorf("ProjectManagedLicenses.ListManagedLicenses returned error: %v", err)
	}

	want := []*ManagedLicense{
		{
			ID:             1,
			Name:           "MIT",
			ApprovalStatus: ManagedLicenseApproved,
		},
		{
			ID:             3,
			Name:           "ISC",
			ApprovalStatus: ManagedLicenseBlacklisted,
		},
	}

	if !reflect.DeepEqual(want, licenses) {
		t.Errorf("ProjectManagedLicenses.ListManagedLicenses returned %+v, want %+v", licenses, want)
	}
}

func TestGetProjectManagedLicenses(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/managed_licenses/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_project_managed_license.json")
	})

	license, _, err := client.ProjectManagedLicenses.GetManagedLicense(1, 3)
	if err != nil {
		t.Errorf("ProjectManagedLicenses.GetManagedLicense returned error: %v", err)
	}

	want := &ManagedLicense{
		ID:             3,
		Name:           "ISC",
		ApprovalStatus: ManagedLicenseBlacklisted,
	}

	if !reflect.DeepEqual(want, license) {
		t.Errorf("ProjectManagedLicenses.GetManagedLicense returned %+v, want %+v", license, want)
	}
}

func TestAddProjectManagedLicenses(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/managed_licenses", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"name":"MIT","approval_status":"approved"}`)
		mustWriteHTTPResponse(t, w, "testdata/add_project_managed_license.json")
	})

	ops := AddManagedLicenseOptions{
		Name:           "MIT",
		ApprovalStatus: ManagedLicenseApproved,
	}
	license, _, err := client.ProjectManagedLicenses.AddManagedLicense(1, &ops)
	if err != nil {
		t.Errorf("ProjectManagedLicenses.AddManagedLicense returned error: %v", err)
	}

	want := &ManagedLicense{
		ID:             123,
		Name:           "MIT",
		ApprovalStatus: ManagedLicenseApproved,
	}

	if !reflect.DeepEqual(want, license) {
		t.Errorf("ProjectManagedLicenses.AddManagedLicense returned %+v, want %+v", license, want)
	}
}

func TestRemoveProjectManagedLicenses(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/managed_licenses/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.ProjectManagedLicenses.RemoveManagedLicense(1, 3)
	if err != nil {
		t.Errorf("ProjectManagedLicenses.RemoveManagedLicense returned error: %v", err)
	}
}

func TestEditProjectManagedLicenses(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/managed_licenses/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		approvalStatus := r.URL.Query().Get("approval_status")
		if approvalStatus != "blacklisted" {
			t.Errorf("query param approval_status should be blacklisted but was %s", approvalStatus)
		}
		mustWriteHTTPResponse(t, w, "testdata/edit_project_managed_license.json")
	})

	ops := EditManagedLicenceOptions{
		ApprovalStatus: ManagedLicenseBlacklisted,
	}
	license, _, err := client.ProjectManagedLicenses.EditManagedLicense(1, 3, &ops)
	if err != nil {
		t.Errorf("ProjectManagedLicenses.EditManagedLicense returned error: %v", err)
	}

	want := &ManagedLicense{
		ID:             3,
		Name:           "CUSTOM",
		ApprovalStatus: ManagedLicenseBlacklisted,
	}

	if !reflect.DeepEqual(want, license) {
		t.Errorf("ProjectManagedLicenses.EditManagedLicense returned %+v, want %+v", license, want)
	}
}
