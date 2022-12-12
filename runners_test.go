//
// Copyright 2021, Sander van Harmelen
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
	"time"
)

func TestDisableRunner(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/runners/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.Runners.DisableProjectRunner(1, 2, nil)
	if err != nil {
		t.Fatalf("Runners.DisableProjectRunner returns an error: %v", err)
	}
}

func TestListRunnersJobs(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/1/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, exampleListRunnerJobs)
	})

	opt := &ListRunnerJobsOptions{}

	jobs, _, err := client.Runners.ListRunnerJobs(1, opt)
	if err != nil {
		t.Fatalf("Runners.ListRunnersJobs returns an error: %v", err)
	}

	pipeline := struct {
		ID        int    `json:"id"`
		ProjectID int    `json:"project_id"`
		Ref       string `json:"ref"`
		Sha       string `json:"sha"`
		Status    string `json:"status"`
	}{
		ID:        8777,
		ProjectID: 3252,
		Ref:       "master",
		Sha:       "6c016b801a88f4bd31f927fc045b5c746a6f823e",
		Status:    "failed",
	}

	want := []*Job{
		{
			ID:             1,
			Status:         "failed",
			Stage:          "test",
			Name:           "run_tests",
			Ref:            "master",
			CreatedAt:      Time(time.Date(2021, time.October, 22, 11, 59, 25, 201000000, time.UTC)),
			StartedAt:      Time(time.Date(2021, time.October, 22, 11, 59, 33, 660000000, time.UTC)),
			FinishedAt:     Time(time.Date(2021, time.October, 22, 15, 59, 25, 201000000, time.UTC)),
			Duration:       171.540594,
			QueuedDuration: 2.535766,
			User: &User{
				ID:          368,
				Name:        "John SMITH",
				Username:    "john.smith",
				AvatarURL:   "https://gitlab.example.com/uploads/-/system/user/avatar/368/avatar.png",
				State:       "blocked",
				WebURL:      "https://gitlab.example.com/john.smith",
				PublicEmail: "john.smith@example.com",
			},
			Commit: &Commit{
				ID:             "6c016b801a88f4bd31f927fc045b5c746a6f823e",
				ShortID:        "6c016b80",
				CreatedAt:      Time(time.Date(2018, time.March, 21, 14, 41, 0, 0, time.UTC)),
				ParentIDs:      []string{"6008b4902d40799ab11688e502d9f1f27f6d2e18"},
				Title:          "Update env for specific runner",
				Message:        "Update env for specific runner\n",
				AuthorName:     "John SMITH",
				AuthorEmail:    "john.smith@example.com",
				AuthoredDate:   Time(time.Date(2018, time.March, 21, 14, 41, 0, 0, time.UTC)),
				CommitterName:  "John SMITH",
				CommitterEmail: "john.smith@example.com",
				CommittedDate:  Time(time.Date(2018, time.March, 21, 14, 41, 0, 0, time.UTC)),
				WebURL:         "https://gitlab.example.com/awesome/packages/common/-/commit/6c016b801a88f4bd31f927fc045b5c746a6f823e",
			},
			Pipeline: pipeline,
			WebURL:   "https://gitlab.example.com/awesome/packages/common/-/jobs/14606",
			Project: &Project{
				ID:                3252,
				Description:       "Common nodejs paquet for producer",
				Name:              "common",
				NameWithNamespace: "awesome",
				Path:              "common",
				PathWithNamespace: "awesome",
				CreatedAt:         Time(time.Date(2018, time.February, 13, 9, 21, 48, 107000000, time.UTC)),
			},
		},
	}
	if !reflect.DeepEqual(want[0], jobs[0]) {
		t.Errorf("Runners.ListRunnersJobs returned %+v, want %+v", jobs[0], want[0])
	}
}

func TestRemoveRunner(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.Runners.RemoveRunner(1, nil)
	if err != nil {
		t.Fatalf("Runners.RemoveARunner returns an error: %v", err)
	}
}

func TestUpdateRunnersDetails(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/6", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, exampleDetailResponse)
	})

	opt := &UpdateRunnerDetailsOptions{}

	details, _, err := client.Runners.UpdateRunnerDetails(6, opt, nil)
	if err != nil {
		t.Fatalf("Runners.UpdateRunnersDetails returns an error: %v", err)
	}

	projects := []struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		NameWithNamespace string `json:"name_with_namespace"`
		Path              string `json:"path"`
		PathWithNamespace string `json:"path_with_namespace"`
	}{{
		ID:                1,
		Name:              "GitLab Community Edition",
		NameWithNamespace: "GitLab.org / GitLab Community Edition",
		Path:              "gitlab-ce",
		PathWithNamespace: "gitlab-org/gitlab-ce",
	}}

	want := &RunnerDetails{
		Active:         true,
		Description:    "test-1-20150125-test",
		ID:             6,
		IsShared:       false,
		RunnerType:     "project_type",
		ContactedAt:    Time(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
		Online:         true,
		Status:         "online",
		Token:          "205086a8e3b9a2b818ffac9b89d102",
		TagList:        []string{"ruby", "mysql"},
		RunUntagged:    true,
		AccessLevel:    "ref_protected",
		Projects:       projects,
		MaximumTimeout: 3600,
		Locked:         false,
	}
	if !reflect.DeepEqual(want, details) {
		t.Errorf("Runners.UpdateRunnersDetails returned %+v, want %+v", details, want)
	}
}

func TestGetRunnerDetails(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/6", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, exampleDetailResponse)
	})

	details, _, err := client.Runners.GetRunnerDetails(6, nil)
	if err != nil {
		t.Fatalf("Runners.GetRunnerDetails returns an error: %v", err)
	}

	projects := []struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		NameWithNamespace string `json:"name_with_namespace"`
		Path              string `json:"path"`
		PathWithNamespace string `json:"path_with_namespace"`
	}{{
		ID:                1,
		Name:              "GitLab Community Edition",
		NameWithNamespace: "GitLab.org / GitLab Community Edition",
		Path:              "gitlab-ce",
		PathWithNamespace: "gitlab-org/gitlab-ce",
	}}

	want := &RunnerDetails{
		Active:         true,
		Description:    "test-1-20150125-test",
		ID:             6,
		IsShared:       false,
		RunnerType:     "project_type",
		ContactedAt:    Time(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
		Online:         true,
		Status:         "online",
		Token:          "205086a8e3b9a2b818ffac9b89d102",
		TagList:        []string{"ruby", "mysql"},
		RunUntagged:    true,
		AccessLevel:    "ref_protected",
		Projects:       projects,
		MaximumTimeout: 3600,
		Locked:         false,
	}
	if !reflect.DeepEqual(want, details) {
		t.Errorf("Runners.UpdateRunnersDetails returned %+v, want %+v", details, want)
	}
}

func TestRegisterNewRunner(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, exampleRegisterNewRunner)
	})

	opt := &RegisterNewRunnerOptions{}

	runner, resp, err := client.Runners.RegisterNewRunner(opt, nil)
	if err != nil {
		t.Fatalf("Runners.RegisterNewRunner returns an error: %v", err)
	}

	want := &Runner{
		ID:             12345,
		Token:          "6337ff461c94fd3fa32ba3b1ff4125",
		TokenExpiresAt: Time(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
	}
	if !reflect.DeepEqual(want, runner) {
		t.Errorf("Runners.RegisterNewRunner returned %+v, want %+v", runner, want)
	}

	wantCode := 201
	if !reflect.DeepEqual(wantCode, resp.StatusCode) {
		t.Errorf("Runners.DeleteRegisteredRunner returned status code %+v, want %+v", resp.StatusCode, wantCode)
	}
}

// Similar to TestRegisterNewRunner but sends info struct and some extra other
// fields too.
func TestRegisterNewRunnerInfo(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"id": 53,
			"description": "some description",
			"active": true,
			"ip_address": "1.2.3.4",
			"name": "some name",
			"online": true,
			"status": "online",
			"token": "1111122222333333444444",
			"token_expires_at": "2016-01-25T16:39:48.166Z"
		  }`)
	})

	opt := &RegisterNewRunnerOptions{
		Token:       String("6337ff461c94fd3fa32ba3b1ff4125"),
		Description: String("some description"),
		Info: &RegisterNewRunnerInfoOptions{
			String("some name"),
			String("13.7.0"),
			String("943fc252"),
			String("linux"),
			String("amd64"),
		},
		Active:         Bool(true),
		Locked:         Bool(true),
		RunUntagged:    Bool(false),
		MaximumTimeout: Int(45),
	}
	runner, resp, err := client.Runners.RegisterNewRunner(opt, nil)
	if err != nil {
		t.Fatalf("Runners.RegisterNewRunner returns an error: %v", err)
	}

	want := &Runner{
		ID:             53,
		Description:    "some description",
		Active:         true,
		IPAddress:      "1.2.3.4",
		Name:           "some name",
		Online:         true,
		Status:         "online",
		Token:          "1111122222333333444444",
		TokenExpiresAt: Time(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
	}
	if !reflect.DeepEqual(want, runner) {
		t.Errorf("Runners.RegisterNewRunner returned %+v, want %+v", runner, want)
	}

	wantCode := 201
	if !reflect.DeepEqual(wantCode, resp.StatusCode) {
		t.Errorf("Runners.DeleteRegisteredRunner returned status code %+v, want %+v", resp.StatusCode, wantCode)
	}
}

func TestDeleteRegisteredRunner(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	opt := &DeleteRegisteredRunnerOptions{}

	resp, err := client.Runners.DeleteRegisteredRunner(opt, nil)
	if err != nil {
		t.Fatalf("Runners.DeleteRegisteredRunner returns an error: %v", err)
	}

	want := 204
	if !reflect.DeepEqual(want, resp.StatusCode) {
		t.Errorf("Runners.DeleteRegisteredRunner returned returned status code  %+v, want %+v", resp.StatusCode, want)
	}
}

func TestDeleteRegisteredRunnerByID(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/11111", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	rid := 11111

	resp, err := client.Runners.DeleteRegisteredRunnerByID(rid, nil)
	if err != nil {
		t.Fatalf("Runners.DeleteRegisteredRunnerByID returns an error: %v", err)
	}

	want := 204
	if !reflect.DeepEqual(want, resp.StatusCode) {
		t.Errorf("Runners.DeleteRegisteredRunnerByID returned returned status code  %+v, want %+v", resp.StatusCode, want)
	}
}

func TestVerifyRegisteredRunner(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/verify", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
	})

	opt := &VerifyRegisteredRunnerOptions{}

	resp, err := client.Runners.VerifyRegisteredRunner(opt, nil)
	if err != nil {
		t.Fatalf("Runners.VerifyRegisteredRunner returns an error: %v", err)
	}

	want := 200
	if !reflect.DeepEqual(want, resp.StatusCode) {
		t.Errorf("Runners.VerifyRegisteredRunner returned returned status code  %+v, want %+v", resp.StatusCode, want)
	}
}

func TestResetInstanceRunnerRegistrationToken(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/reset_registration_token", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"token": "6337ff461c94fd3fa32ba3b1ff4125",
			"token_expires_at": "2016-01-25T16:39:48.166Z"
		}`)
	})

	token, resp, err := client.Runners.ResetInstanceRunnerRegistrationToken(nil)
	if err != nil {
		t.Fatalf("Runners.ResetInstanceRunnerRegistrationToken returns an error: %v", err)
	}

	want := &RunnerRegistrationToken{
		Token:          String("6337ff461c94fd3fa32ba3b1ff4125"),
		TokenExpiresAt: Time(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
	}
	if !reflect.DeepEqual(want, token) {
		t.Errorf("Runners.ResetInstanceRunnerRegistrationToken returned %+v, want %+v", token, want)
	}

	wantCode := 201
	if !reflect.DeepEqual(wantCode, resp.StatusCode) {
		t.Errorf("Runners.ResetInstanceRunnerRegistrationToken returned returned status code  %+v, want %+v", resp.StatusCode, wantCode)
	}
}

func TestResetGroupRunnerRegistrationToken(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/foobar/runners/reset_registration_token", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"token": "6337ff461c94fd3fa32ba3b1ff4125",
			"token_expires_at": "2016-01-25T16:39:48.166Z"
		}`)
	})

	token, resp, err := client.Runners.ResetGroupRunnerRegistrationToken("foobar", nil)
	if err != nil {
		t.Fatalf("Runners.ResetGroupRunnerRegistrationToken returns an error: %v", err)
	}

	want := &RunnerRegistrationToken{
		Token:          String("6337ff461c94fd3fa32ba3b1ff4125"),
		TokenExpiresAt: Time(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
	}
	if !reflect.DeepEqual(want, token) {
		t.Errorf("Runners.ResetGroupRunnerRegistrationToken returned %+v, want %+v", token, want)
	}

	wantCode := 201
	if !reflect.DeepEqual(wantCode, resp.StatusCode) {
		t.Errorf("Runners.ResetGroupRunnerRegistrationToken returned returned status code  %+v, want %+v", resp.StatusCode, wantCode)
	}
}

func TestResetProjectRunnerRegistrationToken(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/9/runners/reset_registration_token", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"token": "6337ff461c94fd3fa32ba3b1ff4125",
			"token_expires_at": "2016-01-25T16:39:48.166Z"
		}`)
	})

	token, resp, err := client.Runners.ResetProjectRunnerRegistrationToken("9", nil)
	if err != nil {
		t.Fatalf("Runners.ResetProjectRunnerRegistrationToken returns an error: %v", err)
	}

	want := &RunnerRegistrationToken{
		Token:          String("6337ff461c94fd3fa32ba3b1ff4125"),
		TokenExpiresAt: Time(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
	}
	if !reflect.DeepEqual(want, token) {
		t.Errorf("Runners.ResetProjectRunnerRegistrationToken returned %+v, want %+v", token, want)
	}

	wantCode := 201
	if !reflect.DeepEqual(wantCode, resp.StatusCode) {
		t.Errorf("Runners.ResetProjectRunnerRegistrationToken returned returned status code  %+v, want %+v", resp.StatusCode, wantCode)
	}
}

func TestResetRunnerAuthenticationToken(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runners/42/reset_authentication_token", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"token": "6337ff461c94fd3fa32ba3b1ff4125",
			"token_expires_at": "2016-01-25T16:39:48.166Z"
		}`)
	})

	token, resp, err := client.Runners.ResetRunnerAuthenticationToken(42, nil)
	if err != nil {
		t.Fatalf("Runners.ResetRunnerAuthenticationToken returns an error: %v", err)
	}

	want := &RunnerAuthenticationToken{
		Token:          String("6337ff461c94fd3fa32ba3b1ff4125"),
		TokenExpiresAt: Time(time.Date(2016, time.January, 25, 16, 39, 48, 166000000, time.UTC)),
	}
	if !reflect.DeepEqual(want, token) {
		t.Errorf("Runners.ResetRunnerAuthenticationToken returned %+v, want %+v", token, want)
	}

	wantCode := 201
	if !reflect.DeepEqual(wantCode, resp.StatusCode) {
		t.Errorf("Runners.ResetRunnerAuthenticationToken returned returned status code  %+v, want %+v", resp.StatusCode, wantCode)
	}
}
