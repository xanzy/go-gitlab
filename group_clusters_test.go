package gitlab

import (
	"fmt"
	"net/http"
	"testing"
)

func TestListGroupClusters(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)
	pid := 1234

	mux.HandleFunc("/api/v4/groups/1234/clusters", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `[
  {
    "id":18,
    "name":"cluster-1",
    "domain":"example.com",
    "created_at":"2019-01-02T20:18:12.563Z",
    "provider_type":"user",
    "platform_type":"kubernetes",
    "environment_scope":"*",
    "cluster_type":"group_type",
    "user":
    {
      "id":1,
      "name":"Administrator",
      "username":"root",
      "state":"active",
      "avatar_url":"https://www.gravatar.com/avatar/4249f4df72b..",
      "web_url":"https://gitlab.example.com/root"
    },
    "platform_kubernetes":
    {
      "api_url":"https://104.197.68.152",
      "authorization_type":"rbac",
      "ca_cert":"-----BEGIN CERTIFICATE-----\r\nhFiK1L61owwDQYJKoZIhvcNAQELBQAw\r\nLzEtMCsGA1UEAxMkZDA1YzQ1YjctNzdiMS00NDY0LThjNmEtMTQ0ZDJkZjM4ZDBj\r\nMB4XDTE4MTIyNzIwMDM1MVoXDTIzMTIyNjIxMDM1MVowLzEtMCsGA1UEAxMkZDA1\r\nYzQ1YjctNzdiMS00NDY0LThjNmEtMTQ0ZDJkZjM.......-----END CERTIFICATE-----"
    }
  }
]`
		fmt.Fprint(w, response)
	})

	clusters, _, err := client.GroupCluster.ListClusters(pid)

	if err != nil {
		t.Errorf("GroupClusters.ListClusters returned error: %v", err)
	}

	if len(clusters) != 1 {
		t.Errorf("expected 1 cluster; got %d", len(clusters))
	}

	if clusters[0].ID != 18 {
		t.Errorf("expected clusterID 1; got %d", clusters[0].ID)
	}

	if clusters[0].Domain != "example.com" {
		t.Errorf("expected cluster domain example.com; got %q", clusters[0].Domain)
	}
}

func TestGetGroupCluster(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)
	pid := 1234

	mux.HandleFunc("/api/v4/groups/1234/clusters/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{
  "id":18,
  "name":"cluster-1",
  "domain":"example.com",
  "created_at":"2019-01-02T20:18:12.563Z",
  "provider_type":"user",
  "platform_type":"kubernetes",
  "environment_scope":"*",
  "cluster_type":"group_type",
  "user":
  {
    "id":1,
    "name":"Administrator",
    "username":"root",
    "state":"active",
    "avatar_url":"https://www.gravatar.com/avatar/4249f4df72b..",
    "web_url":"https://gitlab.example.com/root"
  },
  "platform_kubernetes":
  {
    "api_url":"https://104.197.68.152",
    "authorization_type":"rbac",
    "ca_cert":"-----BEGIN CERTIFICATE-----\r\nhFiK1L61owwDQYJKoZIhvcNAQELBQAw\r\nLzEtMCsGA1UEAxMkZDA1YzQ1YjctNzdiMS00NDY0LThjNmEtMTQ0ZDJkZjM4ZDBj\r\nMB4XDTE4MTIyNzIwMDM1MVoXDTIzMTIyNjIxMDM1MVowLzEtMCsGA1UEAxMkZDA1\r\nYzQ1YjctNzdiMS00NDY0LThjNmEtMTQ0ZDJkZjM.......-----END CERTIFICATE-----"
  },
  "group":
  {
    "id":26,
    "description":"",
    "name":"group-with-clusters-api",
    "name_with_namespace":"Administrator / group-with-clusters-api",
    "path":"group-with-clusters-api",
    "path_with_namespace":"root/group-with-clusters-api",
    "created_at":"2019-01-02T20:13:32.600Z",
    "default_branch":null,
    "tag_list":[],
    "ssh_url_to_repo":"ssh://gitlab.example.com/root/group-with-clusters-api.git",
    "http_url_to_repo":"https://gitlab.example.com/root/group-with-clusters-api.git",
    "web_url":"https://gitlab.example.com/root/group-with-clusters-api",
    "readme_url":null,
    "avatar_url":null,
    "star_count":0,
    "forks_count":0,
    "last_activity_at":"2019-01-02T20:13:32.600Z",
    "namespace":
    {
      "id":1,
      "name":"root",
      "path":"root",
      "kind":"user",
      "full_path":"root",
      "parent_id":null
    }
  }
}`
		fmt.Fprint(w, response)
	})

	cluster, _, err := client.GroupCluster.GetCluster(pid, 1)

	if err != nil {
		t.Errorf("GroupClusters.ListClusters returned error: %v", err)
	}

	if cluster.ID != 18 {
		t.Errorf("expected clusterID 18; got %d", cluster.ID)
	}

	if cluster.Domain != "example.com" {
		t.Errorf("expected cluster domain example.com; got %q", cluster.Domain)
	}
}

func TestAddGroupCluster(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)
	pid := 1234

	mux.HandleFunc("/api/v4/groups/1234/clusters/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{
  "id":24,
  "name":"cluster-5",
  "domain":"example.com",
  "created_at":"2019-01-03T21:53:40.610Z",
  "provider_type":"user",
  "platform_type":"kubernetes",
  "environment_scope":"*",
  "cluster_type":"group_type",
  "user":
  {
    "id":1,
    "name":"Administrator",
    "username":"root",
    "state":"active",
    "avatar_url":"https://www.gravatar.com/avatar/4249f4df72b..",
    "web_url":"https://gitlab.example.com/root"
  },
  "platform_kubernetes":
  {
    "api_url":"https://35.111.51.20",
    "authorization_type":"rbac",
    "ca_cert":"-----BEGIN CERTIFICATE-----\r\nhFiK1L61owwDQYJKoZIhvcNAQELBQAw\r\nLzEtMCsGA1UEAxMkZDA1YzQ1YjctNzdiMS00NDY0LThjNmEtMTQ0ZDJkZjM4ZDBj\r\nMB4XDTE4MTIyNzIwMDM1MVoXDTIzMTIyNjIxMDM1MVowLzEtMCsGA1UEAxMkZDA1\r\nYzQ1YjctNzdiMS00NDY0LThjNmEtMTQ0ZDJkZjM.......-----END CERTIFICATE-----"
  },
  "group":
  {
    "id":26,
    "description":"",
    "name":"group-with-clusters-api",
    "name_with_namespace":"Administrator / group-with-clusters-api",
    "path":"group-with-clusters-api",
    "path_with_namespace":"root/group-with-clusters-api",
    "created_at":"2019-01-02T20:13:32.600Z",
    "default_branch":null,
    "tag_list":[],
    "ssh_url_to_repo":"ssh:://gitlab.example.com/root/group-with-clusters-api.git",
    "http_url_to_repo":"https://gitlab.example.com/root/group-with-clusters-api.git",
    "web_url":"https://gitlab.example.com/root/group-with-clusters-api",
    "readme_url":null,
    "avatar_url":null,
    "star_count":0,
    "forks_count":0,
    "last_activity_at":"2019-01-02T20:13:32.600Z",
    "namespace":
    {
      "id":1,
      "name":"root",
      "path":"root",
      "kind":"user",
      "full_path":"root",
      "parent_id":null
    }
  }
}`
		fmt.Fprint(w, response)
	})

	cluster, _, err := client.GroupCluster.AddCluster(pid, &AddClusterOptions{})

	if err != nil {
		t.Errorf("GroupClusters.AddCluster returned error: %v", err)
	}

	if cluster.ID != 24 {
		t.Errorf("expected ClusterID 24; got %d", cluster.ID)
	}
}

func TestEditGroupCluster(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)
	pid := 1234

	mux.HandleFunc("/api/v4/groups/1234/clusters/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		response := `{
  "id":24,
  "name":"new-cluster-name",
  "domain":"example.com",
  "created_at":"2019-01-03T21:53:40.610Z",
  "provider_type":"user",
  "platform_type":"kubernetes",
  "environment_scope":"*",
  "cluster_type":"group_type",
  "user":
  {
    "id":1,
    "name":"Administrator",
    "username":"root",
    "state":"active",
    "avatar_url":"https://www.gravatar.com/avatar/4249f4df72b..",
    "web_url":"https://gitlab.example.com/root"
  },
  "platform_kubernetes":
  {
    "api_url":"https://new-api-url.com",
    "authorization_type":"rbac",
    "ca_cert":null
  },
  "group":
  {
    "id":26,
    "description":"",
    "name":"group-with-clusters-api",
    "name_with_namespace":"Administrator / group-with-clusters-api",
    "path":"group-with-clusters-api",
    "path_with_namespace":"root/group-with-clusters-api",
    "created_at":"2019-01-02T20:13:32.600Z",
    "default_branch":null,
    "tag_list":[],
    "ssh_url_to_repo":"ssh:://gitlab.example.com/root/group-with-clusters-api.git",
    "http_url_to_repo":"https://gitlab.example.com/root/group-with-clusters-api.git",
    "web_url":"https://gitlab.example.com/root/group-with-clusters-api",
    "readme_url":null,
    "avatar_url":null,
    "star_count":0,
    "forks_count":0,
    "last_activity_at":"2019-01-02T20:13:32.600Z",
    "namespace":
    {
      "id":1,
      "name":"root",
      "path":"root",
      "kind":"user",
      "full_path":"root",
      "parent_id":null
    }
  }
}`
		fmt.Fprint(w, response)
	})

	cluster, _, err := client.GroupCluster.EditCluster(pid, 1, &EditClusterOptions{})

	if err != nil {
		t.Errorf("GroupClusters.EditCluster returned error: %v", err)
	}

	if cluster.ID != 24 {
		t.Errorf("expected ClusterID 24; got %d", cluster.ID)
	}
}

func TestDeleteGroupCluster(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/1234/clusters/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusAccepted)
	})

	resp, err := client.GroupCluster.DeleteCluster(1234, 1)
	if err != nil {
		t.Errorf("GroupCluster.DeleteCluster returned error: %v", err)
	}

	want := http.StatusAccepted
	got := resp.StatusCode
	if got != want {
		t.Errorf("GroupCluster.DeleteCluster returned %d, want %d", got, want)
	}
}
