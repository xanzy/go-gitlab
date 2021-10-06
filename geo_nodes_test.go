package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeoNodesService_CreateGeoNode(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/geo_nodes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
			  "id": 3,
			  "name": "Test Node 1",
			  "url": "https://secondary.example.com/",
			  "internal_url": "https://secondary.example.com/",
			  "primary": false,
			  "enabled": true,
			  "current": false,
			  "files_max_capacity": 10,
			  "repos_max_capacity": 25,
			  "verification_max_capacity": 100,
			  "selective_sync_type": "namespaces",
			  "selective_sync_shards": null,
			  "selective_sync_namespace_ids": [1, 25],
			  "minimum_reverification_interval": 7,
			  "container_repositories_max_capacity": 10,
			  "sync_object_storage": false,
			  "clone_protocol": "http",
			  "web_edit_url": "https://primary.example.com/admin/geo/nodes/3/edit",
			  "web_geo_projects_url": "http://secondary.example.com/admin/geo/projects",
			  "_links": {
				 "self": "https://primary.example.com/api/v4/geo_nodes/3",
				 "status": "https://primary.example.com/api/v4/geo_nodes/3/status",
				 "repair": "https://primary.example.com/api/v4/geo_nodes/3/repair"
			  }
			}
		`)
	})

	want := &GeoNode{
		ID:                               3,
		Name:                             "Test Node 1",
		URL:                              "https://secondary.example.com/",
		InternalURL:                      "https://secondary.example.com/",
		Primary:                          false,
		Enabled:                          true,
		Current:                          false,
		FilesMaxCapacity:                 10,
		ReposMaxCapacity:                 25,
		VerificationMaxCapacity:          100,
		SelectiveSyncType:                "namespaces",
		SelectiveSyncShards:              nil,
		SelectiveSyncNamespaceIds:        []int{1, 25},
		MinimumReverificationInterval:    7,
		ContainerRepositoriesMaxCapacity: 10,
		SyncObjectStorage:                false,
		CloneProtocol:                    "http",
		WebEditURL:                       "https://primary.example.com/admin/geo/nodes/3/edit",
		WebGeoProjectsURL:                "http://secondary.example.com/admin/geo/projects",
		Links: GeoNodeLinks{
			Self:   "https://primary.example.com/api/v4/geo_nodes/3",
			Status: "https://primary.example.com/api/v4/geo_nodes/3/status",
			Repair: "https://primary.example.com/api/v4/geo_nodes/3/repair",
		},
	}

	g, resp, err := client.GeoNodes.CreateGeoNode(nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, g)

	g, resp, err = client.GeoNodes.CreateGeoNode(nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, g)
}

func TestGeoNodesService_CreateGeoNode_StatusNotFound(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/geo_nodes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	g, resp, err := client.GeoNodes.CreateGeoNode(nil)
	require.Error(t, err)
	require.Nil(t, g)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGeoNodesService_ListGeoNodes(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/geo_nodes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
				{
				  "id": 3,
				  "name": "in-node",
				  "url": "https://secondary.example.com/",
				  "internal_url": "https://secondary.example.com/",
				  "primary": false,
				  "enabled": true,
				  "current": false,
				  "files_max_capacity": 10,
				  "repos_max_capacity": 25,
				  "verification_max_capacity": 100,
				  "selective_sync_type": "namespaces",
				  "selective_sync_shards": null,
				  "selective_sync_namespace_ids": [1, 25],
				  "minimum_reverification_interval": 7,
				  "container_repositories_max_capacity": 10,
				  "sync_object_storage": false,
				  "clone_protocol": "http",
				  "web_edit_url": "https://primary.example.com/admin/geo/nodes/3/edit",
				  "web_geo_projects_url": "http://secondary.example.com/admin/geo/projects",
				  "_links": {
					 "self": "https://primary.example.com/api/v4/geo_nodes/3",
					 "status": "https://primary.example.com/api/v4/geo_nodes/3/status",
					 "repair": "https://primary.example.com/api/v4/geo_nodes/3/repair"
				  }
				}
			]
		`)
	})

	want := []*GeoNode{
		{
			ID:                               3,
			Name:                             "in-node",
			URL:                              "https://secondary.example.com/",
			InternalURL:                      "https://secondary.example.com/",
			Primary:                          false,
			Enabled:                          true,
			Current:                          false,
			FilesMaxCapacity:                 10,
			ReposMaxCapacity:                 25,
			VerificationMaxCapacity:          100,
			SelectiveSyncType:                "namespaces",
			SelectiveSyncShards:              nil,
			SelectiveSyncNamespaceIds:        []int{1, 25},
			MinimumReverificationInterval:    7,
			ContainerRepositoriesMaxCapacity: 10,
			SyncObjectStorage:                false,
			CloneProtocol:                    "http",
			WebEditURL:                       "https://primary.example.com/admin/geo/nodes/3/edit",
			WebGeoProjectsURL:                "http://secondary.example.com/admin/geo/projects",
			Links: GeoNodeLinks{
				Self:   "https://primary.example.com/api/v4/geo_nodes/3",
				Status: "https://primary.example.com/api/v4/geo_nodes/3/status",
				Repair: "https://primary.example.com/api/v4/geo_nodes/3/repair",
			},
		},
	}

	gs, resp, err := client.GeoNodes.ListGeoNodes(nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, gs)

	gs, resp, err = client.GeoNodes.ListGeoNodes(nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, gs)
}

func TestGeoNodesService_ListGeoNodes_StatusNotFound(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/geo_nodes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusNotFound)
	})

	gs, resp, err := client.GeoNodes.ListGeoNodes(nil)
	require.Error(t, err)
	require.Nil(t, gs)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGeoNodesService_GetGeoNode(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/geo_nodes/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			{
			  "id": 3,
			  "name": "in-node",
			  "url": "https://secondary.example.com/",
			  "internal_url": "https://secondary.example.com/",
			  "primary": false,
			  "enabled": true,
			  "current": false,
			  "files_max_capacity": 10,
			  "repos_max_capacity": 25,
			  "verification_max_capacity": 100,
			  "selective_sync_type": "namespaces",
			  "selective_sync_shards": null,
			  "selective_sync_namespace_ids": [1, 25],
			  "minimum_reverification_interval": 7,
			  "container_repositories_max_capacity": 10,
			  "sync_object_storage": false,
			  "clone_protocol": "http",
			  "web_edit_url": "https://primary.example.com/admin/geo/nodes/3/edit",
			  "web_geo_projects_url": "http://secondary.example.com/admin/geo/projects",
			  "_links": {
				 "self": "https://primary.example.com/api/v4/geo_nodes/3",
				 "status": "https://primary.example.com/api/v4/geo_nodes/3/status",
				 "repair": "https://primary.example.com/api/v4/geo_nodes/3/repair"
			  }
			}
		`)
	})

	want := &GeoNode{
		ID:                               3,
		Name:                             "in-node",
		URL:                              "https://secondary.example.com/",
		InternalURL:                      "https://secondary.example.com/",
		Primary:                          false,
		Enabled:                          true,
		Current:                          false,
		FilesMaxCapacity:                 10,
		ReposMaxCapacity:                 25,
		VerificationMaxCapacity:          100,
		SelectiveSyncType:                "namespaces",
		SelectiveSyncShards:              nil,
		SelectiveSyncNamespaceIds:        []int{1, 25},
		MinimumReverificationInterval:    7,
		ContainerRepositoriesMaxCapacity: 10,
		SyncObjectStorage:                false,
		CloneProtocol:                    "http",
		WebEditURL:                       "https://primary.example.com/admin/geo/nodes/3/edit",
		WebGeoProjectsURL:                "http://secondary.example.com/admin/geo/projects",
		Links: GeoNodeLinks{
			Self:   "https://primary.example.com/api/v4/geo_nodes/3",
			Status: "https://primary.example.com/api/v4/geo_nodes/3/status",
			Repair: "https://primary.example.com/api/v4/geo_nodes/3/repair",
		},
	}

	g, resp, err := client.GeoNodes.GetGeoNode(3, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, g)

	g, resp, err = client.GeoNodes.GetGeoNode(3, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, g)

	g, resp, err = client.GeoNodes.GetGeoNode(5, nil)
	require.Error(t, err)
	require.Nil(t, g)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGeoNodesService_EditGeoNode(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/geo_nodes/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
			{
			  "id": 3,
			  "name": "in-node",
			  "url": "https://secondary.example.com/",
			  "internal_url": "https://secondary.example.com/",
			  "primary": false,
			  "enabled": true,
			  "current": false,
			  "files_max_capacity": 10,
			  "repos_max_capacity": 25,
			  "verification_max_capacity": 100,
			  "selective_sync_type": "namespaces",
			  "selective_sync_shards": null,
			  "selective_sync_namespace_ids": [1, 25],
			  "minimum_reverification_interval": 7,
			  "container_repositories_max_capacity": 10,
			  "sync_object_storage": false,
			  "clone_protocol": "http",
			  "web_edit_url": "https://primary.example.com/admin/geo/nodes/3/edit",
			  "web_geo_projects_url": "http://secondary.example.com/admin/geo/projects",
			  "_links": {
				 "self": "https://primary.example.com/api/v4/geo_nodes/3",
				 "status": "https://primary.example.com/api/v4/geo_nodes/3/status",
				 "repair": "https://primary.example.com/api/v4/geo_nodes/3/repair"
			  }
			}
		`)
	})

	want := &GeoNode{
		ID:                               3,
		Name:                             "in-node",
		URL:                              "https://secondary.example.com/",
		InternalURL:                      "https://secondary.example.com/",
		Primary:                          false,
		Enabled:                          true,
		Current:                          false,
		FilesMaxCapacity:                 10,
		ReposMaxCapacity:                 25,
		VerificationMaxCapacity:          100,
		SelectiveSyncType:                "namespaces",
		SelectiveSyncShards:              nil,
		SelectiveSyncNamespaceIds:        []int{1, 25},
		MinimumReverificationInterval:    7,
		ContainerRepositoriesMaxCapacity: 10,
		SyncObjectStorage:                false,
		CloneProtocol:                    "http",
		WebEditURL:                       "https://primary.example.com/admin/geo/nodes/3/edit",
		WebGeoProjectsURL:                "http://secondary.example.com/admin/geo/projects",
		Links: GeoNodeLinks{
			Self:   "https://primary.example.com/api/v4/geo_nodes/3",
			Status: "https://primary.example.com/api/v4/geo_nodes/3/status",
			Repair: "https://primary.example.com/api/v4/geo_nodes/3/repair",
		},
	}

	g, resp, err := client.GeoNodes.EditGeoNode(3, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, g)

	g, resp, err = client.GeoNodes.EditGeoNode(3, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, g)

	g, resp, err = client.GeoNodes.EditGeoNode(5, nil)
	require.Error(t, err)
	require.Nil(t, g)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGeoNodesService_DeleteGeoNode(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/geo_nodes/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.GeoNodes.DeleteGeoNode(3, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.GeoNodes.DeleteGeoNode(3, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.GeoNodes.DeleteGeoNode(5, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
