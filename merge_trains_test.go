package gitlab

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestListProjectMergeTrains(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_trains", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_merge_trains_in_project.json")
	})

	opts := &ListMergeTrainsOptions{}

	mergeTrains, _, err := client.MergeTrains.ListProjectMergeTrains(1, opts)

	if err != nil {
		t.Errorf("MergeTrains.ListProjectMergeTrains returned error: %v", err)
	}

	mergeRequestCreatedAt := time.Date(2020, 2, 6, 8, 39, 14, 883000000, time.UTC)
	mergeRequestUpdatedAt := time.Date(2020, 02, 6,8,40,57,38000000, time.UTC)

	pipelineCreatedAt := time.Date(2020, 2,6,8,40,42,410000000, time.UTC)
	pipelineUpdatedAt:= time.Date(2020,2,6,8,40,46,912000000, time.UTC)

	mergeTrainCreatedAt := time.Date(2020,2,6,8,39,47,217000000, time.UTC)
	mergeTrainUpdatedAt := time.Date(2020,2,6,8,40,57,720000000, time.UTC)
	mergeTrainMergedAt := time.Date(2020,2,6,8,40,57,719000000, time.UTC)

	want := []*MergeTrain{
		{
			ID: 110,
    	MergeRequest: &MergeTrainMergeRequest{
      ID: 126,
      IID: 59,
      ProjectID: 20,
      Title: "Test MR 1580978354",
      Description: "",
      State: "merged",
      CreatedAt: &mergeRequestCreatedAt,
      UpdatedAt:&mergeRequestUpdatedAt,
      WebURL: "http://local.gitlab.test:8181/root/merge-train-race-condition/-/merge_requests/59",
    },
    User: &MergeTrainUser{
      ID: 1,
      Name: "Administrator",
      Username: "root",
      State: "active",
      AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
      WebURL: "http://local.gitlab.test:8181/root",
    },
    Pipeline: &MergeTrainPipeline{
      ID: 246,
      SHA: "bcc17a8ffd51be1afe45605e714085df28b80b13",
      Ref: "refs/merge-requests/59/train",
      Status: "success",
      CreatedAt: &pipelineCreatedAt,
      UpdatedAt: &pipelineUpdatedAt,
      WebURL: "http://local.gitlab.test:8181/root/merge-train-race-condition/pipelines/246",
    },
    CreatedAt: &mergeTrainCreatedAt,
    UpdatedAt: &mergeTrainUpdatedAt,
    TargetBranch: "feature-1580973432",
    Status: "merged",
    MergedAt: &mergeTrainMergedAt,
    Duration: 70,
		},
	}

	if !reflect.DeepEqual(want, mergeTrains) {
		t.Errorf("MergeTrains.ListProjectMergeTrains returned %+v, want %+v", mergeTrains, want)
	}
}

func TestListMergeRequestInMergeTrain(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/597/merge_trains/main", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_merge_requests_in_merge_train.json")
	})

	opts := &ListMergeTrainsOptions{}

	mergeTrains, _, err := client.MergeTrains.ListMergeRequestInMergeTrain(597, "main", opts)

	if err != nil {
		t.Errorf("MergeTrains.ListMergeRequestInMergeTrain returned error: %v", err)
	}

	mergeRequestCreatedAt := time.Date(2022,10,31,19,6,5,725000000, time.UTC)
	mergeRequestUpdatedAt := time.Date(2022,10,31,19,6,5,725000000, time.UTC)

	pipelineCreatedAt := time.Date(2022,10,31,19,06,06,231000000, time.UTC)
	pipelineUpdatedAt := time.Date(2022,10,31,19,06,06,231000000, time.UTC)

	mergeTrainCreatedAt := time.Date(2022,10,31,19,06,06,237000000, time.UTC)
	mergeTrainUpdatedAt := time.Date(2022,10,31,19,06,06,237000000, time.UTC)

	want := []*MergeTrain{
		{
    ID: 267,
    MergeRequest: &MergeTrainMergeRequest{
      ID: 273,
      IID: 1,
      ProjectID: 597,
      Title: "My title 9",
      Description: "",
      State: "opened",
      CreatedAt: &mergeRequestCreatedAt,
      UpdatedAt: &mergeRequestUpdatedAt,
      WebURL: "http://localhost/namespace18/project21/-/merge_requests/1",
    },
    User: &MergeTrainUser{
      ID: 933,
      Username: "user12",
      Name: "Sidney Jones31",
      State: "active",
      AvatarURL: "https://www.gravatar.com/avatar/6c8365de387cb3db10ecc7b1880203c4?s=80\u0026d=identicon",
      WebURL: "http://localhost/user12",
    },
    Pipeline: &MergeTrainPipeline{
      ID: 273,
      IID: 1,
      ProjectID: 598,
      SHA: "b83d6e391c22777fca1ed3012fce84f633d7fed0",
      Ref: "main",
      Status: "pending",
      Source: "push",
      CreatedAt: &pipelineCreatedAt,
      UpdatedAt: &pipelineUpdatedAt,
      WebURL: "http://localhost/namespace19/project22/-/pipelines/273",
    },
    CreatedAt: &mergeTrainCreatedAt,
    UpdatedAt:&mergeTrainUpdatedAt,
    TargetBranch:"main",
    Status:"idle",
    MergedAt:nil,
    Duration:0,
  },
	}

	if !reflect.DeepEqual(want, mergeTrains) {
		t.Errorf("MergeTrains.ListMergeRequestInMergeTrain returned %+v, want %+v", mergeTrains, want)
	}
}

func TestGetMergeRequestOnAMergeTrain(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/597/merge_trains/merge_requests/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_merge_request_in_merge_train.json")
	})

	mergeTrain, _, err := client.MergeTrains.GetMergeRequestOnAMergeTrain(597, 1)

	if err != nil {
		t.Errorf("MergeTrains.GetMergeRequestOnAMergeTrain returned error: %v", err)
	}

	mergeRequestCreatedAt := time.Date(2022,10,31,19,6,5,725000000, time.UTC)
	mergeRequestUpdatedAt := time.Date(2022,10,31,19,6,5,725000000, time.UTC)

	pipelineCreatedAt := time.Date(2022,10,31,19,06,06,231000000, time.UTC)
	pipelineUpdatedAt := time.Date(2022,10,31,19,06,06,231000000, time.UTC)

	mergeTrainCreatedAt := time.Date(2022,10,31,19,06,06,237000000, time.UTC)
	mergeTrainUpdatedAt := time.Date(2022,10,31,19,06,06,237000000, time.UTC)

	want := &MergeTrain{
		ID: 267,
		MergeRequest: &MergeTrainMergeRequest{
			ID: 273,
			IID: 1,
			ProjectID: 597,
			Title: "My title 9",
			Description: "",
			State: "opened",
			CreatedAt: &mergeRequestCreatedAt,
			UpdatedAt: &mergeRequestUpdatedAt,
			WebURL: "http://localhost/namespace18/project21/-/merge_requests/1",
		},
		User: &MergeTrainUser{
			ID: 933,
			Username: "user12",
			Name: "Sidney Jones31",
			State: "active",
			AvatarURL: "https://www.gravatar.com/avatar/6c8365de387cb3db10ecc7b1880203c4?s=80\u0026d=identicon",
			WebURL: "http://localhost/user12",
		},
		Pipeline: &MergeTrainPipeline{
			ID: 273,
			IID: 1,
			ProjectID: 598,
			SHA: "b83d6e391c22777fca1ed3012fce84f633d7fed0",
			Ref: "main",
			Status: "pending",
			Source: "push",
			CreatedAt: &pipelineCreatedAt,
			UpdatedAt: &pipelineUpdatedAt,
			WebURL: "http://localhost/namespace19/project22/-/pipelines/273",
		},
		CreatedAt: &mergeTrainCreatedAt,
		UpdatedAt:&mergeTrainUpdatedAt,
		TargetBranch:"main",
		Status:"idle",
		MergedAt:nil,
		Duration:0,
	}

	if !reflect.DeepEqual(want, mergeTrain) {
		t.Errorf("MergeTrains.GetMergeRequestOnAMergeTrain returned %+v, want %+v", mergeTrain, want)
	}
}

func TestAddMergeRequestToMergeTrain(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/597/merge_trains/merge_requests/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		mustWriteHTTPResponse(t, w, "testdata/add_merge_request_in_merge_train.json")
	})

	opt := &AddMergeRequestToMergeTrainOptions{WhenPipelineSucceeds: Bool(true)}

	mergeTrains, _, err := client.MergeTrains.AddMergeRequestToMergeTrain(597, 1, opt)

	if err != nil {
		t.Errorf("MergeTrains.AddMergeRequestToMergeTrain returned error: %v", err)
	}

	mergeRequestCreatedAt := time.Date(2022,10,31,19,6,5,725000000, time.UTC)
	mergeRequestUpdatedAt := time.Date(2022,10,31,19,6,5,725000000, time.UTC)

	pipelineCreatedAt := time.Date(2022,10,31,19,06,06,231000000, time.UTC)
	pipelineUpdatedAt := time.Date(2022,10,31,19,06,06,231000000, time.UTC)

	mergeTrainCreatedAt := time.Date(2022,10,31,19,06,06,237000000, time.UTC)
	mergeTrainUpdatedAt := time.Date(2022,10,31,19,06,06,237000000, time.UTC)


	want := []*MergeTrain{
		{
			ID: 267,
			MergeRequest: &MergeTrainMergeRequest{
				ID: 273,
				IID: 1,
				ProjectID: 597,
				Title: "My title 9",
				Description: "",
				State: "opened",
				CreatedAt: &mergeRequestCreatedAt,
				UpdatedAt: &mergeRequestUpdatedAt,
				WebURL: "http://localhost/namespace18/project21/-/merge_requests/1",
			},
			User: &MergeTrainUser{
				ID: 933,
				Username: "user12",
				Name: "Sidney Jones31",
				State: "active",
				AvatarURL: "https://www.gravatar.com/avatar/6c8365de387cb3db10ecc7b1880203c4?s=80\u0026d=identicon",
				WebURL: "http://localhost/user12",
			},
			Pipeline: &MergeTrainPipeline{
				ID: 273,
				IID: 1,
				ProjectID: 598,
				SHA: "b83d6e391c22777fca1ed3012fce84f633d7fed0",
				Ref: "main",
				Status: "pending",
				Source: "push",
				CreatedAt: &pipelineCreatedAt,
				UpdatedAt: &pipelineUpdatedAt,
				WebURL: "http://localhost/namespace19/project22/-/pipelines/273",
			},
			CreatedAt: &mergeTrainCreatedAt,
			UpdatedAt:&mergeTrainUpdatedAt,
			TargetBranch:"main",
			Status:"idle",
			MergedAt:nil,
			Duration:0,
		},
	}

	if !reflect.DeepEqual(want, mergeTrains) {
		t.Errorf("MergeTrains.AddMergeRequestToMergeTrain returned %+v, want %+v", mergeTrains, want)
	}
}