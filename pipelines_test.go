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

	"github.com/stretchr/testify/assert"
)

func TestListProjectPipelines(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/pipelines", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opt := &ListProjectPipelinesOptions{Ref: String("master")}
	piplines, _, err := client.Pipelines.ListProjectPipelines(1, opt)
	if err != nil {
		t.Errorf("Pipelines.ListProjectPipelines returned error: %v", err)
	}

	want := []*PipelineInfo{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(want, piplines) {
		t.Errorf("Pipelines.ListProjectPipelines returned %+v, want %+v", piplines, want)
	}
}

func TestGetPipeline(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/pipelines/5949167", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1,"status":"success"}`)
	})

	pipeline, _, err := client.Pipelines.GetPipeline(1, 5949167)
	if err != nil {
		t.Errorf("Pipelines.GetPipeline returned error: %v", err)
	}

	want := &Pipeline{ID: 1, Status: "success"}
	if !reflect.DeepEqual(want, pipeline) {
		t.Errorf("Pipelines.GetPipeline returned %+v, want %+v", pipeline, want)
	}
}

func TestGetPipelineVariables(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/pipelines/5949167/variables", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"key":"RUN_NIGHTLY_BUILD","variable_type":"env_var","value":"true"},{"key":"foo","value":"bar"}]`)
	})

	variables, _, err := client.Pipelines.GetPipelineVariables(1, 5949167)
	if err != nil {
		t.Errorf("Pipelines.GetPipelineVariables returned error: %v", err)
	}

	want := []*PipelineVariable{{Key: "RUN_NIGHTLY_BUILD", Value: "true", VariableType: "env_var"}, {Key: "foo", Value: "bar"}}
	if !reflect.DeepEqual(want, variables) {
		t.Errorf("Pipelines.GetPipelineVariables returned %+v, want %+v", variables, want)
	}
}

func TestGetPipelineTestReport(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/pipelines/123456/test_report", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_pipeline_testreport.json")
	})

	testreport, _, err := client.Pipelines.GetPipelineTestReport(1, 123456)
	if err != nil {
		t.Errorf("Pipelines.GetPipelineTestReport returned error: %v", err)
	}

	want := &PipelineTestReport{
		TotalTime:    61.502,
		TotalCount:   9,
		SuccessCount: 5,
		ErrorCount:   4,
		TestSuites: []*PipelineTestSuites{
			{
				Name:         "Failing",
				TotalTime:    60.494,
				TotalCount:   8,
				SuccessCount: 4,
				ErrorCount:   4,
				TestCases: []*PipelineTestCases{
					{
						Status:        "error",
						Name:          "Error testcase 1",
						Classname:     "MyClassOne",
						File:          "/path/file.ext",
						ExecutionTime: 19.987,
						SystemOutput: &SystemOutput{
							Type:    "failure",
							Message: "Failed test",
						},
						StackTrace:    "java.lang.Exception: Stack trace\nat java.base/java.lang.Thread.dumpStack(Thread.java:1383)",
						AttachmentURL: "http://foo.bar",
						RecentFailures: &RecentFailures{
							Count:      10,
							BaseBranch: "master",
						},
					},
					{
						Status:        "error",
						Name:          "Error testcase 2",
						Classname:     "MyClass",
						ExecutionTime: 19.984,
					},
					{
						Status:    "error",
						Name:      "Error testcase 3",
						Classname: "MyClass",
						SystemOutput: &SystemOutput{
							Type:    "failure",
							Message: "Failed test",
						},
					},
					{
						Status:        "success",
						Name:          "Succes full testcase",
						Classname:     "MyClass",
						ExecutionTime: 19.7799999999999985,
					},
				},
			},
			{
				Name:         "Succes suite",
				TotalTime:    1.008,
				TotalCount:   1,
				SuccessCount: 1,
				TestCases: []*PipelineTestCases{{
					Status:        "success",
					Name:          "Succesfull testcase",
					Classname:     "MyClass",
					ExecutionTime: 1.008,
				}},
			},
		},
	}
	if !reflect.DeepEqual(want, testreport) {
		t.Errorf("Pipelines.GetPipelineTestReport returned %+v, want %+v", testreport, want)
	}
}

func TestGetLatestPipeline(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/pipelines/latest", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParams(t, r, "")
		fmt.Fprint(w, `{"id":1,"status":"success"}`)
	})

	pipeline, _, err := client.Pipelines.GetLatestPipeline(1, nil)

	assert.NoError(t, err)
	assert.Equal(t, &Pipeline{ID: 1, Status: "success"}, pipeline)
}

func TestGetLatestPipeline_WithRef(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/pipelines/latest", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParams(t, r, "ref=abc")
		fmt.Fprint(w, `{"id":1,"status":"success"}`)
	})

	pipeline, _, err := client.Pipelines.GetLatestPipeline(1, &GetLatestPipelineOptions{
		Ref: String("abc"),
	})

	assert.NoError(t, err)
	assert.Equal(t, &Pipeline{ID: 1, Status: "success"}, pipeline)
}

func TestCreatePipeline(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/pipeline", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1, "status":"pending"}`)
	})

	opt := &CreatePipelineOptions{Ref: String("master")}
	pipeline, _, err := client.Pipelines.CreatePipeline(1, opt)
	if err != nil {
		t.Errorf("Pipelines.CreatePipeline returned error: %v", err)
	}

	want := &Pipeline{ID: 1, Status: "pending"}
	if !reflect.DeepEqual(want, pipeline) {
		t.Errorf("Pipelines.CreatePipeline returned %+v, want %+v", pipeline, want)
	}
}

func TestRetryPipelineBuild(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/pipelines/5949167/retry", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintln(w, `{"id":1, "status":"pending"}`)
	})

	pipeline, _, err := client.Pipelines.RetryPipelineBuild(1, 5949167)
	if err != nil {
		t.Errorf("Pipelines.RetryPipelineBuild returned error: %v", err)
	}

	want := &Pipeline{ID: 1, Status: "pending"}
	if !reflect.DeepEqual(want, pipeline) {
		t.Errorf("Pipelines.RetryPipelineBuild returned %+v, want %+v", pipeline, want)
	}
}

func TestCancelPipelineBuild(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/pipelines/5949167/cancel", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintln(w, `{"id":1, "status":"canceled"}`)
	})

	pipeline, _, err := client.Pipelines.CancelPipelineBuild(1, 5949167)
	if err != nil {
		t.Errorf("Pipelines.CancelPipelineBuild returned error: %v", err)
	}

	want := &Pipeline{ID: 1, Status: "canceled"}
	if !reflect.DeepEqual(want, pipeline) {
		t.Errorf("Pipelines.CancelPipelineBuild returned %+v, want %+v", pipeline, want)
	}
}

func TestDeletePipeline(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/pipelines/5949167", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Pipelines.DeletePipeline("1", 5949167)
	if err != nil {
		t.Errorf("Pipelines.DeletePipeline returned error: %v", err)
	}
}
