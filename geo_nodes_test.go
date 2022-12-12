package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeoNodesService_CreateGeoNode(t *testing.T) {
	mux, client := setup(t)

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
	mux, client := setup(t)

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
	mux, client := setup(t)

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
	mux, client := setup(t)

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
	mux, client := setup(t)

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
	mux, client := setup(t)

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
	mux, client := setup(t)

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

func TestGeoNodesService_RepairGeoNode(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/geo_nodes/3/repair", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
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

	g, resp, err := client.GeoNodes.RepairGeoNode(3, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, g)

	g, resp, err = client.GeoNodes.RepairGeoNode(3, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, g)

	g, resp, err = client.GeoNodes.RepairGeoNode(5, nil)
	require.Error(t, err)
	require.Nil(t, g)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGeoNodesService_RetrieveStatusOfAllGeoNodes(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/geo_nodes/status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_geo_nodes_status.json")
	})

	want := []*GeoNodeStatus{
		{
			GeoNodeID:                                  1,
			Healthy:                                    true,
			Health:                                     "Healthy",
			HealthStatus:                               "Healthy",
			MissingOauthApplication:                    false,
			AttachmentsCount:                           1,
			AttachmentsSyncedInPercentage:              "0.00%",
			LfsObjectsSyncedInPercentage:               "0.00%",
			JobArtifactsCount:                          2,
			JobArtifactsSyncedInPercentage:             "0.00%",
			ContainerRepositoriesCount:                 3,
			ContainerRepositoriesSyncedInPercentage:    "0.00%",
			DesignRepositoriesCount:                    3,
			DesignRepositoriesSyncedInPercentage:       "0.00%",
			ProjectsCount:                              41,
			RepositoriesCount:                          41,
			RepositoriesSyncedInPercentage:             "0.00%",
			WikisCount:                                 41,
			WikisSyncedInPercentage:                    "0.00%",
			ReplicationSlotsCount:                      1,
			ReplicationSlotsUsedCount:                  1,
			ReplicationSlotsUsedInPercentage:           "100.00%",
			RepositoriesCheckedCount:                   20,
			RepositoriesCheckedFailedCount:             20,
			RepositoriesCheckedInPercentage:            "100.00%",
			RepositoriesChecksummedCount:               20,
			RepositoriesChecksumFailedCount:            5,
			RepositoriesChecksummedInPercentage:        "48.78%",
			WikisChecksummedCount:                      10,
			WikisChecksumFailedCount:                   3,
			WikisChecksummedInPercentage:               "24.39%",
			RepositoriesVerifiedCount:                  20,
			RepositoriesVerificationFailedCount:        5,
			RepositoriesVerifiedInPercentage:           "48.78%",
			RepositoriesChecksumMismatchCount:          3,
			WikisVerifiedCount:                         10,
			WikisVerificationFailedCount:               3,
			WikisVerifiedInPercentage:                  "24.39%",
			WikisChecksumMismatchCount:                 1,
			RepositoriesRetryingVerificationCount:      1,
			WikisRetryingVerificationCount:             3,
			LastEventID:                                23,
			LastEventTimestamp:                         1509681166,
			LastSuccessfulStatusCheckTimestamp:         1510125024,
			Version:                                    "10.3.0",
			Revision:                                   "33d33a096a",
			MergeRequestDiffsCount:                     5,
			MergeRequestDiffsChecksumTotalCount:        5,
			MergeRequestDiffsChecksummedCount:          5,
			MergeRequestDiffsSyncedInPercentage:        "0.00%",
			MergeRequestDiffsVerifiedInPercentage:      "0.00%",
			PackageFilesCount:                          5,
			PackageFilesChecksumTotalCount:             5,
			PackageFilesChecksummedCount:               5,
			PackageFilesSyncedInPercentage:             "0.00%",
			PackageFilesVerifiedInPercentage:           "0.00%",
			PagesDeploymentsCount:                      5,
			PagesDeploymentsChecksumTotalCount:         5,
			PagesDeploymentsChecksummedCount:           5,
			PagesDeploymentsSyncedInPercentage:         "0.00%",
			PagesDeploymentsVerifiedInPercentage:       "0.00%",
			TerraformStateVersionsCount:                5,
			TerraformStateVersionsChecksumTotalCount:   5,
			TerraformStateVersionsChecksummedCount:     5,
			TerraformStateVersionsSyncedInPercentage:   "0.00%",
			TerraformStateVersionsVerifiedInPercentage: "0.00%",
			SnippetRepositoriesCount:                   5,
			SnippetRepositoriesChecksumTotalCount:      5,
			SnippetRepositoriesChecksummedCount:        5,
			SnippetRepositoriesSyncedInPercentage:      "0.00%",
			SnippetRepositoriesVerifiedInPercentage:    "0.00%",
			GroupWikiRepositoriesCount:                 5,
			GroupWikiRepositoriesChecksumTotalCount:    5,
			GroupWikiRepositoriesChecksummedCount:      5,
			GroupWikiRepositoriesSyncedInPercentage:    "0.00%",
			GroupWikiRepositoriesVerifiedInPercentage:  "0.00%",
			PipelineArtifactsCount:                     5,
			PipelineArtifactsChecksumTotalCount:        5,
			PipelineArtifactsChecksummedCount:          5,
			PipelineArtifactsSyncedInPercentage:        "0.00%",
			PipelineArtifactsVerifiedInPercentage:      "0.00%",
			UploadsCount:                               5,
			UploadsSyncedInPercentage:                  "0.00%",
		},
	}

	gnss, resp, err := client.GeoNodes.RetrieveStatusOfAllGeoNodes(nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, gnss)

	gnss, resp, err = client.GeoNodes.RetrieveStatusOfAllGeoNodes(errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, gnss)
}

func TestGeoNodesService_RetrieveStatusOfAllGeoNodes_StatusNotFound(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/geo_nodes/status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusNotFound)
	})

	gnss, resp, err := client.GeoNodes.RetrieveStatusOfAllGeoNodes(nil)
	require.Error(t, err)
	require.Nil(t, gnss)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGeoNodesService_RetrieveStatusOfGeoNode(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/geo_nodes/1/status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_geo_node_status.json")
	})

	want := &GeoNodeStatus{
		GeoNodeID:                                  1,
		Healthy:                                    true,
		Health:                                     "Healthy",
		HealthStatus:                               "Healthy",
		MissingOauthApplication:                    false,
		AttachmentsCount:                           1,
		AttachmentsSyncedInPercentage:              "0.00%",
		LfsObjectsSyncedInPercentage:               "0.00%",
		JobArtifactsCount:                          2,
		JobArtifactsSyncedInPercentage:             "0.00%",
		ContainerRepositoriesCount:                 3,
		ContainerRepositoriesSyncedInPercentage:    "0.00%",
		DesignRepositoriesCount:                    3,
		DesignRepositoriesSyncedInPercentage:       "0.00%",
		ProjectsCount:                              41,
		RepositoriesCount:                          41,
		RepositoriesSyncedInPercentage:             "0.00%",
		WikisCount:                                 41,
		WikisSyncedInPercentage:                    "0.00%",
		ReplicationSlotsCount:                      1,
		ReplicationSlotsUsedCount:                  1,
		ReplicationSlotsUsedInPercentage:           "100.00%",
		RepositoriesCheckedCount:                   20,
		RepositoriesCheckedFailedCount:             20,
		RepositoriesCheckedInPercentage:            "100.00%",
		RepositoriesChecksummedCount:               20,
		RepositoriesChecksumFailedCount:            5,
		RepositoriesChecksummedInPercentage:        "48.78%",
		WikisChecksummedCount:                      10,
		WikisChecksumFailedCount:                   3,
		WikisChecksummedInPercentage:               "24.39%",
		RepositoriesVerifiedCount:                  20,
		RepositoriesVerificationFailedCount:        5,
		RepositoriesVerifiedInPercentage:           "48.78%",
		RepositoriesChecksumMismatchCount:          3,
		WikisVerifiedCount:                         10,
		WikisVerificationFailedCount:               3,
		WikisVerifiedInPercentage:                  "24.39%",
		WikisChecksumMismatchCount:                 1,
		RepositoriesRetryingVerificationCount:      1,
		WikisRetryingVerificationCount:             3,
		LastEventID:                                23,
		LastEventTimestamp:                         1509681166,
		LastSuccessfulStatusCheckTimestamp:         1510125024,
		Version:                                    "10.3.0",
		Revision:                                   "33d33a096a",
		MergeRequestDiffsCount:                     5,
		MergeRequestDiffsChecksumTotalCount:        5,
		MergeRequestDiffsChecksummedCount:          5,
		MergeRequestDiffsSyncedInPercentage:        "0.00%",
		MergeRequestDiffsVerifiedInPercentage:      "0.00%",
		PackageFilesCount:                          5,
		PackageFilesChecksumTotalCount:             5,
		PackageFilesChecksummedCount:               5,
		PackageFilesSyncedInPercentage:             "0.00%",
		PackageFilesVerifiedInPercentage:           "0.00%",
		PagesDeploymentsCount:                      5,
		PagesDeploymentsChecksumTotalCount:         5,
		PagesDeploymentsChecksummedCount:           5,
		PagesDeploymentsSyncedInPercentage:         "0.00%",
		PagesDeploymentsVerifiedInPercentage:       "0.00%",
		TerraformStateVersionsCount:                5,
		TerraformStateVersionsChecksumTotalCount:   5,
		TerraformStateVersionsChecksummedCount:     5,
		TerraformStateVersionsSyncedInPercentage:   "0.00%",
		TerraformStateVersionsVerifiedInPercentage: "0.00%",
		SnippetRepositoriesCount:                   5,
		SnippetRepositoriesChecksumTotalCount:      5,
		SnippetRepositoriesChecksummedCount:        5,
		SnippetRepositoriesSyncedInPercentage:      "0.00%",
		SnippetRepositoriesVerifiedInPercentage:    "0.00%",
		GroupWikiRepositoriesCount:                 5,
		GroupWikiRepositoriesChecksumTotalCount:    5,
		GroupWikiRepositoriesChecksummedCount:      5,
		GroupWikiRepositoriesSyncedInPercentage:    "0.00%",
		GroupWikiRepositoriesVerifiedInPercentage:  "0.00%",
		PipelineArtifactsCount:                     5,
		PipelineArtifactsChecksumTotalCount:        5,
		PipelineArtifactsChecksummedCount:          5,
		PipelineArtifactsSyncedInPercentage:        "0.00%",
		PipelineArtifactsVerifiedInPercentage:      "0.00%",
		UploadsCount:                               5,
		UploadsSyncedInPercentage:                  "0.00%",
	}

	gns, resp, err := client.GeoNodes.RetrieveStatusOfGeoNode(1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, gns)

	gns, resp, err = client.GeoNodes.RetrieveStatusOfGeoNode(1, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, gns)

	gns, resp, err = client.GeoNodes.RetrieveStatusOfGeoNode(3, nil)
	require.Error(t, err)
	require.Nil(t, gns)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
