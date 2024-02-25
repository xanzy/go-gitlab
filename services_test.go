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

func TestListServices(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})
	want := []*Service{{ID: 1}, {ID: 2}}

	services, _, err := client.Services.ListServices(1)
	if err != nil {
		t.Fatalf("Services.ListServices returns an error: %v", err)
	}
	if !reflect.DeepEqual(want, services) {
		t.Errorf("Services.ListServices returned %+v, want %+v", services, want)
	}
}

func TestCustomIssueTrackerService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/custom-issue-tracker", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
      "id": 1,
      "title": "5",
      "push_events": true,
      "properties": {
        "new_issue_url": "1",
        "issues_url": "2",
        "project_url": "3"
      }
    }`)
	})
	want := &CustomIssueTrackerService{
		Service: Service{
			ID:         1,
			Title:      "5",
			PushEvents: true,
		},
		Properties: &CustomIssueTrackerServiceProperties{
			NewIssueURL: "1",
			IssuesURL:   "2",
			ProjectURL:  "3",
		},
	}

	service, _, err := client.Services.GetCustomIssueTrackerService(1)
	if err != nil {
		t.Fatalf("Services.GetCustomIssueTrackerService returns an error: %v", err)
	}
	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.GetCustomIssueTrackerService returned %+v, want %+v", service, want)
	}
}

func TestSetCustomIssueTrackerService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/custom-issue-tracker", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	opt := &SetCustomIssueTrackerServiceOptions{
		NewIssueURL: Ptr("1"),
		IssuesURL:   Ptr("2"),
		ProjectURL:  Ptr("3"),
		Description: Ptr("4"),
		Title:       Ptr("5"),
		PushEvents:  Ptr(true),
	}

	_, err := client.Services.SetCustomIssueTrackerService(1, opt)
	if err != nil {
		t.Fatalf("Services.SetCustomIssueTrackerService returns an error: %v", err)
	}
}

func TestDeleteCustomIssueTrackerService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/custom-issue-tracker", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Services.DeleteCustomIssueTrackerService(1)
	if err != nil {
		t.Fatalf("Services.DeleteCustomIssueTrackerService returns an error: %v", err)
	}
}

func TestGetDataDogService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/datadog", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
      "id": 1,
      "active": true,
      "properties": {
        "api_url": "",
        "datadog_env": "production",
        "datadog_service": "gitlab",
        "datadog_site": "datadoghq.com",
        "datadog_tags": "country=canada\nprovince=ontario",
        "archive_trace_events": true
      }
    }`)
	})
	want := &DataDogService{
		Service: Service{ID: 1, Active: true},
		Properties: &DataDogServiceProperties{
			APIURL:             "",
			DataDogEnv:         "production",
			DataDogService:     "gitlab",
			DataDogSite:        "datadoghq.com",
			DataDogTags:        "country=canada\nprovince=ontario",
			ArchiveTraceEvents: true,
		},
	}

	service, _, err := client.Services.GetDataDogService(1)
	if err != nil {
		t.Fatalf("Services.GetDataDogService returns an error: %v", err)
	}
	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.GetDataDogService returned %+v, want %+v", service, want)
	}
}

func TestSetDataDogService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/datadog", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	opt := &SetDataDogServiceOptions{
		APIKey:             String("secret"),
		APIURL:             String("https://some-api.com"),
		DataDogEnv:         String("sandbox"),
		DataDogService:     String("source-code"),
		DataDogSite:        String("datadoghq.eu"),
		DataDogTags:        String("country=france"),
		ArchiveTraceEvents: Bool(false),
	}

	_, err := client.Services.SetDataDogService(1, opt)
	if err != nil {
		t.Fatalf("Services.SetDataDogService returns an error: %v", err)
	}
}

func TestDeleteDataDogService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/datadog", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Services.DeleteDataDogService(1)
	if err != nil {
		t.Fatalf("Services.DeleteDataDogService returns an error: %v", err)
	}
}

func TestGetDiscordService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/discord", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &DiscordService{Service: Service{ID: 1}}

	service, _, err := client.Services.GetDiscordService(1)
	if err != nil {
		t.Fatalf("Services.GetDiscordService returns an error: %v", err)
	}
	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.GetDiscordService returned %+v, want %+v", service, want)
	}
}

func TestSetDiscordService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/discord", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	opt := &SetDiscordServiceOptions{
		WebHook: Ptr("webhook_uri"),
	}

	_, err := client.Services.SetDiscordService(1, opt)
	if err != nil {
		t.Fatalf("Services.SetDiscordService returns an error: %v", err)
	}
}

func TestDeleteDiscordService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/discord", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Services.DeleteDiscordService(1)
	if err != nil {
		t.Fatalf("Services.DeleteDiscordService returns an error: %v", err)
	}
}

func TestGetDroneCIService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/drone-ci", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &DroneCIService{Service: Service{ID: 1}}

	service, _, err := client.Services.GetDroneCIService(1)
	if err != nil {
		t.Fatalf("Services.GetDroneCIService returns an error: %v", err)
	}
	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.GetDroneCIService returned %+v, want %+v", service, want)
	}
}

func TestSetDroneCIService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/drone-ci", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	opt := &SetDroneCIServiceOptions{Ptr("token"), Ptr("drone-url"), Ptr(true), nil, nil, nil}

	_, err := client.Services.SetDroneCIService(1, opt)
	if err != nil {
		t.Fatalf("Services.SetDroneCIService returns an error: %v", err)
	}
}

func TestDeleteDroneCIService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/drone-ci", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Services.DeleteDroneCIService(1)
	if err != nil {
		t.Fatalf("Services.DeleteDroneCIService returns an error: %v", err)
	}
}

func TestGetEmailsOnPushService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/integrations/emails-on-push", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &EmailsOnPushService{Service: Service{ID: 1}}

	service, _, err := client.Services.GetEmailsOnPushService(1)
	if err != nil {
		t.Fatalf("Services.GetEmailsOnPushService returns an error: %v", err)
	}
	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.GetEmailsOnPushService returned %+v, want %+v", service, want)
	}
}

func TestSetEmailsOnPushService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/integrations/emails-on-push", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	opt := &SetEmailsOnPushServiceOptions{Ptr("t"), Ptr(true), Ptr(true), Ptr(true), Ptr(true), Ptr("t")}

	_, err := client.Services.SetEmailsOnPushService(1, opt)
	if err != nil {
		t.Fatalf("Services.SetEmailsOnPushService returns an error: %v", err)
	}
}

func TestDeleteEmailsOnPushService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/integrations/emails-on-push", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Services.DeleteEmailsOnPushService(1)
	if err != nil {
		t.Fatalf("Services.DeleteEmailsOnPushService returns an error: %v", err)
	}
}

func TestGetJiraService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/0/services/jira", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1, "properties": {"jira_issue_transition_id": "2"}}`)
	})

	mux.HandleFunc("/api/v4/projects/1/services/jira", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1, "properties": {"jira_issue_transition_id": 2}}`)
	})

	mux.HandleFunc("/api/v4/projects/2/services/jira", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1, "properties": {"jira_issue_transition_id": "2,3"}}`)
	})

	mux.HandleFunc("/api/v4/projects/3/services/jira", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1, "properties": {}}`)
	})

	want := []*JiraService{
		{
			Service: Service{ID: 1},
			Properties: &JiraServiceProperties{
				JiraIssueTransitionID: "2",
			},
		},
		{
			Service: Service{ID: 1},
			Properties: &JiraServiceProperties{
				JiraIssueTransitionID: "2",
			},
		},
		{
			Service: Service{ID: 1},
			Properties: &JiraServiceProperties{
				JiraIssueTransitionID: "2,3",
			},
		},
		{
			Service:    Service{ID: 1},
			Properties: &JiraServiceProperties{},
		},
	}

	for testcase := 0; testcase < len(want); testcase++ {
		service, _, err := client.Services.GetJiraService(testcase)
		if err != nil {
			t.Fatalf("Services.GetJiraService returns an error: %v", err)
		}

		if !reflect.DeepEqual(want[testcase], service) {
			t.Errorf("Services.GetJiraService returned %+v, want %+v", service, want[testcase])
		}
	}
}

func TestSetJiraService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/jira", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	opt := &SetJiraServiceOptions{
		URL:                   Ptr("asd"),
		APIURL:                Ptr("asd"),
		ProjectKey:            Ptr("as"),
		Username:              Ptr("aas"),
		Password:              Ptr("asd"),
		Active:                Ptr(true),
		JiraIssueTransitionID: Ptr("2,3"),
		CommitEvents:          Ptr(true),
		CommentOnEventEnabled: Ptr(true),
		MergeRequestsEvents:   Ptr(true),
	}

	_, err := client.Services.SetJiraService(1, opt)
	if err != nil {
		t.Fatalf("Services.SetJiraService returns an error: %v", err)
	}
}

func TestDeleteJiraService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/jira", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Services.DeleteJiraService(1)
	if err != nil {
		t.Fatalf("Services.DeleteJiraService returns an error: %v", err)
	}
}

func TestGetMattermostService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/mattermost", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &MattermostService{Service: Service{ID: 1}}

	service, _, err := client.Services.GetMattermostService(1)
	if err != nil {
		t.Fatalf("Services.GetMattermostService returns an error: %v", err)
	}
	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.GetMattermostService returned %+v, want %+v", service, want)
	}
}

func TestSetMattermostService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/mattermost", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	opt := &SetMattermostServiceOptions{
		WebHook:  Ptr("webhook_uri"),
		Username: Ptr("username"),
		Channel:  Ptr("#development"),
	}

	_, err := client.Services.SetMattermostService(1, opt)
	if err != nil {
		t.Fatalf("Services.SetMasttermostService returns an error: %v", err)
	}
}

func TestDeleteMattermostService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/mattermost", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Services.DeleteMattermostService(1)
	if err != nil {
		t.Fatalf("Services.DeleteMattermostService returns an error: %v", err)
	}
}

func TestGetMattermostSlashCommandsService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/mattermost-slash-commands", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &MattermostSlashCommandsService{Service: Service{ID: 1}}

	service, _, err := client.Services.GetMattermostSlashCommandsService(1)
	if err != nil {
		t.Fatalf("Services.mattermost-slash-commands returns an error: %v", err)
	}
	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.mattermost-slash-commands returned %+v, want %+v", service, want)
	}
}

func TestSetMattermostSlashCommandsService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/mattermost-slash-commands", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	opt := &SetMattermostSlashCommandsServiceOptions{
		Token:    Ptr("token"),
		Username: Ptr("username"),
	}

	_, err := client.Services.SetMattermostSlashCommandsService(1, opt)
	if err != nil {
		t.Fatalf("Services.SetMattermostSlashCommandsService returns an error: %v", err)
	}
}

func TestDeleteMattermostSlashCommandsService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/mattermost-slash-commands", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Services.DeleteMattermostSlashCommandsService(1)
	if err != nil {
		t.Fatalf("Services.DeleteMattermostSlashCommandsService returns an error: %v", err)
	}
}

func TestGetPipelinesEmailService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/pipelines-email", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &PipelinesEmailService{Service: Service{ID: 1}}

	service, _, err := client.Services.GetPipelinesEmailService(1)
	if err != nil {
		t.Fatalf("Services.GetPipelinesEmailService returns an error: %v", err)
	}
	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.GetPipelinesEmailService returned %+v, want %+v", service, want)
	}
}

func TestSetPipelinesEmailService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/pipelines-email", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	opt := &SetPipelinesEmailServiceOptions{
		Recipients:                Ptr("test@email.com"),
		NotifyOnlyBrokenPipelines: Ptr(true),
		NotifyOnlyDefaultBranch:   Ptr(false),
		AddPusher:                 nil,
		BranchesToBeNotified:      nil,
		PipelineEvents:            nil,
	}

	_, err := client.Services.SetPipelinesEmailService(1, opt)
	if err != nil {
		t.Fatalf("Services.SetPipelinesEmailService returns an error: %v", err)
	}
}

func TestDeletePipelinesEmailService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/pipelines-email", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Services.DeletePipelinesEmailService(1)
	if err != nil {
		t.Fatalf("Services.DeletePipelinesEmailService returns an error: %v", err)
	}
}

func TestGetPrometheusService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/prometheus", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &PrometheusService{Service: Service{ID: 1}}

	service, _, err := client.Services.GetPrometheusService(1)
	if err != nil {
		t.Fatalf("Services.GetPrometheusService returns an error: %v", err)
	}
	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.GetPrometheusService returned %+v, want %+v", service, want)
	}
}

func TestSetPrometheusService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/prometheus", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	opt := &SetPrometheusServiceOptions{Ptr("t"), Ptr("u"), Ptr("a")}

	_, err := client.Services.SetPrometheusService(1, opt)
	if err != nil {
		t.Fatalf("Services.SetDroneCIService returns an error: %v", err)
	}
}

func TestDeletePrometheusService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/prometheus", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Services.DeletePrometheusService(1)
	if err != nil {
		t.Fatalf("Services.DeletePrometheusService returns an error: %v", err)
	}
}

func TestGetSlackService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/slack", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &SlackService{Service: Service{ID: 1}}

	service, _, err := client.Services.GetSlackService(1)
	if err != nil {
		t.Fatalf("Services.GetSlackService returns an error: %v", err)
	}
	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.GetSlackService returned %+v, want %+v", service, want)
	}
}

func TestSetSlackService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/slack", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	opt := &SetSlackServiceOptions{
		WebHook:  Ptr("webhook_uri"),
		Username: Ptr("username"),
		Channel:  Ptr("#development"),
	}

	_, err := client.Services.SetSlackService(1, opt)
	if err != nil {
		t.Fatalf("Services.SetSlackService returns an error: %v", err)
	}
}

func TestDeleteSlackService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/slack", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Services.DeleteSlackService(1)
	if err != nil {
		t.Fatalf("Services.DeleteSlackService returns an error: %v", err)
	}
}

func TestGetSlackSlashCommandsService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/slack-slash-commands", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &SlackSlashCommandsService{Service: Service{ID: 1}}

	service, _, err := client.Services.GetSlackSlashCommandsService(1)
	if err != nil {
		t.Fatalf("Services.GetSlackSlashCommandsService returns an error: %v", err)
	}
	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.GetSlackSlashCommandsService returned %+v, want %+v", service, want)
	}
}

func TestSetSlackSlashCommandsService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/slack-slash-commands", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	opt := &SetSlackSlashCommandsServiceOptions{
		Token: Ptr("token"),
	}

	_, err := client.Services.SetSlackSlashCommandsService(1, opt)
	if err != nil {
		t.Fatalf("Services.SetSlackSlashCommandsService returns an error: %v", err)
	}
}

func TestDeleteSlackSlashCommandsService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/slack-slash-commands", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Services.DeleteSlackSlashCommandsService(1)
	if err != nil {
		t.Fatalf("Services.DeleteSlackSlashCommandsService returns an error: %v", err)
	}
}

func TestGetTelegramService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/telegram", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
			{
			  "id": 1,
			  "title": "Telegram",
			  "slug": "telegram",
			  "created_at": "2023-12-16T20:21:03.117Z",
			  "updated_at": "2023-12-16T20:22:19.140Z",
			  "active": true,
			  "commit_events": true,
			  "push_events": false,
			  "issues_events": false,
			  "incident_events": false,
			  "alert_events": true,
			  "confidential_issues_events": false,
			  "merge_requests_events": false,
			  "tag_push_events": false,
			  "deployment_events": false,
			  "note_events": false,
			  "confidential_note_events": false,
			  "pipeline_events": true,
			  "wiki_page_events": false,
			  "job_events": true,
			  "comment_on_event_enabled": true,
			  "vulnerability_events": false,
			  "properties": {
				"room": "-1000000000000",
				"notify_only_broken_pipelines": false,
				"branches_to_be_notified": "all"
			  }
			}
		`)
	})
	wantCreatedAt, _ := time.Parse(time.RFC3339, "2023-12-16T20:21:03.117Z")
	wantUpdatedAt, _ := time.Parse(time.RFC3339, "2023-12-16T20:22:19.140Z")
	want := &TelegramService{
		Service: Service{
			ID:                       1,
			Title:                    "Telegram",
			Slug:                     "telegram",
			CreatedAt:                &wantCreatedAt,
			UpdatedAt:                &wantUpdatedAt,
			Active:                   true,
			CommitEvents:             true,
			PushEvents:               false,
			IssuesEvents:             false,
			AlertEvents:              true,
			ConfidentialIssuesEvents: false,
			MergeRequestsEvents:      false,
			TagPushEvents:            false,
			DeploymentEvents:         false,
			NoteEvents:               false,
			ConfidentialNoteEvents:   false,
			PipelineEvents:           true,
			WikiPageEvents:           false,
			JobEvents:                true,
			CommentOnEventEnabled:    true,
			VulnerabilityEvents:      false,
		},
		Properties: &TelegramServiceProperties{
			Room:                      "-1000000000000",
			NotifyOnlyBrokenPipelines: false,
			BranchesToBeNotified:      "all",
		},
	}

	service, _, err := client.Services.GetTelegramService(1)
	if err != nil {
		t.Fatalf("Services.GetTelegramService returns an error: %v", err)
	}
	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.GetTelegramService returned %+v, want %+v", service, want)
	}
}

func TestSetTelegramService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/telegram", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	opt := &SetTelegramServiceOptions{
		Token:                     Ptr("token"),
		Room:                      Ptr("-1000"),
		NotifyOnlyBrokenPipelines: Ptr(true),
		BranchesToBeNotified:      Ptr("all"),
		PushEvents:                Ptr(true),
		IssuesEvents:              Ptr(true),
		ConfidentialIssuesEvents:  Ptr(true),
		MergeRequestsEvents:       Ptr(true),
		TagPushEvents:             Ptr(true),
		NoteEvents:                Ptr(true),
		ConfidentialNoteEvents:    Ptr(true),
		PipelineEvents:            Ptr(true),
		WikiPageEvents:            Ptr(true),
	}

	_, err := client.Services.SetTelegramService(1, opt)
	if err != nil {
		t.Fatalf("Services.SetTelegramService returns an error: %v", err)
	}
}

func TestDeleteTelegramService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/telegram", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Services.DeleteTelegramService(1)
	if err != nil {
		t.Fatalf("Services.DeleteTelegramService returns an error: %v", err)
	}
}

func TestGetYouTrackService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/youtrack", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &YouTrackService{Service: Service{ID: 1}}

	service, _, err := client.Services.GetYouTrackService(1)
	if err != nil {
		t.Fatalf("Services.GetYouTrackService returns an error: %v", err)
	}
	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.GetYouTrackService returned %+v, want %+v", service, want)
	}
}

func TestSetYouTrackService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/youtrack", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	opt := &SetYouTrackServiceOptions{
		IssuesURL:   Ptr("https://example.org/youtrack/issue/:id"),
		ProjectURL:  Ptr("https://example.org/youtrack/projects/1"),
		Description: Ptr("description"),
		PushEvents:  Ptr(true),
	}

	_, err := client.Services.SetYouTrackService(1, opt)
	if err != nil {
		t.Fatalf("Services.SetYouTrackService returns an error: %v", err)
	}
}

func TestDeleteYouTrackService(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/services/youtrack", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Services.DeleteYouTrackService(1)
	if err != nil {
		t.Fatalf("Services.DeleteYouTrackService returns an error: %v", err)
	}
}
