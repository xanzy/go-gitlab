package gitlab

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestListEnvironments(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/environments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testURL(t, r, "/api/v4/projects/1/environments?page=1&per_page=10")
		fmt.Fprint(w, `[{"id": 1,"name": "review/fix-foo", "slug": "review-fix-foo-dfjre3", "external_url": "https://review-fix-foo-dfjre3.example.gitlab.com"}]`)
	})

	envs, _, err := client.Environments.ListEnvironments(1, &ListEnvironmentsOptions{Page: 1, PerPage: 10})
	if err != nil {
		log.Fatal(err)
	}

	want := []*Environment{{ID: 1, Name: "review/fix-foo", Slug: "review-fix-foo-dfjre3", ExternalURL: "https://review-fix-foo-dfjre3.example.gitlab.com"}}
	if !reflect.DeepEqual(want, envs) {
		t.Errorf("Environments.ListEnvironments returned %+v, want %+v", envs, want)
	}
}

func TestCreateEnvironment(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/environments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testURL(t, r, "/api/v4/projects/1/environments")
		fmt.Fprint(w, `{"id": 1,"name": "deploy", "slug": "deploy", "external_url": "https://deploy.example.gitlab.com"}`)
	})

	envs, _, err := client.Environments.CreateEnvironment(1, &CreateEnvironmentOptions{Name: String("deploy"), ExternalURL: String("https://deploy.example.gitlab.com")})
	if err != nil {
		log.Fatal(err)
	}

	want := &Environment{ID: 1, Name: "deploy", Slug: "deploy", ExternalURL: "https://deploy.example.gitlab.com"}
	if !reflect.DeepEqual(want, envs) {
		t.Errorf("Environments.CreateEnvironment returned %+v, want %+v", envs, want)
	}
}

func TestEditEnvironment(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/environments/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testURL(t, r, "/api/v4/projects/1/environments/1")
		fmt.Fprint(w, `{"id": 1,"name": "staging", "slug": "staging", "external_url": "https://staging.example.gitlab.com"}`)
	})

	envs, _, err := client.Environments.EditEnvironment(1, 1, &EditEnvironmentOptions{Name: String("staging"), ExternalURL: String("https://staging.example.gitlab.com")})
	if err != nil {
		log.Fatal(err)
	}

	want := &Environment{ID: 1, Name: "staging", Slug: "staging", ExternalURL: "https://staging.example.gitlab.com"}
	if !reflect.DeepEqual(want, envs) {
		t.Errorf("Environments.EditEnvironment returned %+v, want %+v", envs, want)
	}
}

func TestDeleteEnvironment(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/environments/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testURL(t, r, "/api/v4/projects/1/environments/1")
	})
	_, err := client.Environments.DeleteEnvironment(1, 1)
	if err != nil {
		log.Fatal(err)
	}
}

func TestStopEnvironment(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/environments/1/stop", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testURL(t, r, "/api/v4/projects/1/environments/1/stop")
	})
	_, err := client.Environments.StopEnvironment(1, 1)
	if err != nil {
		log.Fatal(err)
	}
}
