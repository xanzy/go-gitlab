//
// Copyright 2021, Igor Varavko
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
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetCurrentPlanLimits(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/application/plan_limits", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
			"conan_max_file_size": 3221225472,
			"generic_packages_max_file_size": 5368709120,
			"helm_max_file_size": 5242880,
			"maven_max_file_size": 3221225472,
			"npm_max_file_size": 524288000,
			"nuget_max_file_size": 524288000,
			"pypi_max_file_size": 3221225472,
			"terraform_module_max_file_size": 1073741824
		  }`)
	})

	opt := &GetCurrentPlanLimitsOptions{
		PlanName: String("default"),
	}
	planlimit, _, err := client.PlanLimits.GetCurrentPlanLimits(opt)
	if err != nil {
		t.Errorf("PlanLimits.GetCurrentPlanLimits returned error: %v", err)
	}

	want := &PlanLimit{
		ConanMaxFileSize:           3221225472,
		GenericPackagesMaxFileSize: 5368709120,
		HelmMaxFileSize:            5242880,
		MavenMaxFileSize:           3221225472,
		NPMMaxFileSize:             524288000,
		NugetMaxFileSize:           524288000,
		PyPiMaxFileSize:            3221225472,
		TerraformModuleMaxFileSize: 1073741824,
	}

	if !reflect.DeepEqual(want, planlimit) {
		t.Errorf("PlanLimits.GetCurrentPlanLimits returned %+v, want %+v", planlimit, want)
	}
}

func TestChangePlanLimits(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/application/plan_limits", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `{
			"conan_max_file_size": 3221225472,
			"generic_packages_max_file_size": 5368709120,
			"helm_max_file_size": 5242880,
			"maven_max_file_size": 3221225472,
			"npm_max_file_size": 524288000,
			"nuget_max_file_size": 524288000,
			"pypi_max_file_size": 3221225472,
			"terraform_module_max_file_size": 1073741824
		  }`)
	})

	opt := &ChangePlanLimitOptions{
		PlanName:         String("default"),
		ConanMaxFileSize: Int(3221225472),
	}
	planlimit, _, err := client.PlanLimits.ChangePlanLimits(opt)
	if err != nil {
		t.Errorf("PlanLimits.ChangePlanLimits returned error: %v", err)
	}

	want := &PlanLimit{
		ConanMaxFileSize:           3221225472,
		GenericPackagesMaxFileSize: 5368709120,
		HelmMaxFileSize:            5242880,
		MavenMaxFileSize:           3221225472,
		NPMMaxFileSize:             524288000,
		NugetMaxFileSize:           524288000,
		PyPiMaxFileSize:            3221225472,
		TerraformModuleMaxFileSize: 1073741824,
	}

	if !reflect.DeepEqual(want, planlimit) {
		t.Errorf("PlanLimits.ChangePlanLimits returned %+v, want %+v", planlimit, want)
	}
}
