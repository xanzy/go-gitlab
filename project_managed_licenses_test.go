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

func TestListManagedLicenses(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/managed_licenses", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_project_managed_licenses.json")
	})

	licenses, _, err := client.ManagedLicenses.ListManagedLicenses(1)
	if err != nil {
		t.Errorf("ManagedLicenses.ListManagedLicenses returned error: %v", err)
	}

	want := []*ManagedLicense{
		{
			ID:             1,
			Name:           "MIT",
			ApprovalStatus: LicenseApproved,
		},
		{
			ID:             3,
			Name:           "ISC",
			ApprovalStatus: LicenseBlacklisted,
		},
	}

	if !reflect.DeepEqual(want, licenses) {
		t.Errorf("ManagedLicenses.ListManagedLicenses returned %+v, want %+v", licenses, want)
	}
}

func TestGetManagedLicenses(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/managed_licenses/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_project_managed_license.json")
	})

	license, _, err := client.ManagedLicenses.GetManagedLicense(1, 3)
	if err != nil {
		t.Errorf("ManagedLicenses.GetManagedLicense returned error: %v", err)
	}

	want := &ManagedLicense{
		ID:             3,
		Name:           "ISC",
		ApprovalStatus: LicenseBlacklisted,
	}

	if !reflect.DeepEqual(want, license) {
		t.Errorf("ManagedLicenses.GetManagedLicense returned %+v, want %+v", license, want)
	}
}

func TestAddManagedLicenses(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/managed_licenses", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"name":"MIT","approval_status":"approved"}`)
		mustWriteHTTPResponse(t, w, "testdata/add_project_managed_license.json")
	})

	ops := AddManagedLicenseOptions{
		Name:           String("MIT"),
		ApprovalStatus: LicenseApprovalStatus(LicenseApproved),
	}
	license, _, err := client.ManagedLicenses.AddManagedLicense(1, &ops)
	if err != nil {
		t.Errorf("ManagedLicenses.AddManagedLicense returned error: %v", err)
	}

	want := &ManagedLicense{
		ID:             123,
		Name:           "MIT",
		ApprovalStatus: LicenseApproved,
	}

	if !reflect.DeepEqual(want, license) {
		t.Errorf("ManagedLicenses.AddManagedLicense returned %+v, want %+v", license, want)
	}
}

func TestDeleteManagedLicenses(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/managed_licenses/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.ManagedLicenses.DeleteManagedLicense(1, 3)
	if err != nil {
		t.Errorf("ManagedLicenses.RemoveManagedLicense returned error: %v", err)
	}
}

func TestEditManagedLicenses(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/managed_licenses/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testBody(t, r, `{"approval_status":"blacklisted"}`)
		mustWriteHTTPResponse(t, w, "testdata/edit_project_managed_license.json")
	})

	ops := EditManagedLicenceOptions{
		ApprovalStatus: LicenseApprovalStatus(LicenseBlacklisted),
	}
	license, _, err := client.ManagedLicenses.EditManagedLicense(1, 3, &ops)
	if err != nil {
		t.Errorf("ManagedLicenses.EditManagedLicense returned error: %v", err)
	}

	want := &ManagedLicense{
		ID:             3,
		Name:           "CUSTOM",
		ApprovalStatus: LicenseBlacklisted,
	}

	if !reflect.DeepEqual(want, license) {
		t.Errorf("ManagedLicenses.EditManagedLicense returned %+v, want %+v", license, want)
	}
}
