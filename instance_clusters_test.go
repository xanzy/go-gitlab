package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInstanceClustersService_ListClusters(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/admin/clusters", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 9,
				"name": "cluster-1",
				"managed": true,
				"enabled": true,
				"domain": null,
				"provider_type": "user",
				"platform_type": "kubernetes",
				"environment_scope": "*",
				"cluster_type": "instance_type",
				"user": {
				  "id": 1,
				  "name": "Administrator",
				  "username": "root",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				  "web_url": "https://gitlab.example.com/root"
				},
				"platform_kubernetes": {
				  "api_url": "https://example.com",
				  "namespace": null,
				  "authorization_type": "rbac",
				  "ca_cert":"-----BEGIN CERTIFICATE-----IxMDM1MV0ZDJkZjM...-----END CERTIFICATE-----"
				},
				"provider_gcp": null,
				"management_project": null
			  }
			]
		`)
	})

	want := []*InstanceCluster{{
		ID:               9,
		Name:             "cluster-1",
		Domain:           "",
		Managed:          true,
		ProviderType:     "user",
		PlatformType:     "kubernetes",
		EnvironmentScope: "*",
		ClusterType:      "instance_type",
		User: &User{
			ID:        1,
			Username:  "root",
			Email:     "",
			Name:      "Administrator",
			State:     "active",
			WebURL:    "https://gitlab.example.com/root",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
		},
		PlatformKubernetes: &PlatformKubernetes{
			APIURL:            "https://example.com",
			Token:             "",
			CaCert:            "-----BEGIN CERTIFICATE-----IxMDM1MV0ZDJkZjM...-----END CERTIFICATE-----",
			Namespace:         "",
			AuthorizationType: "rbac",
		},
		ManagementProject: nil,
	}}

	ics, resp, err := client.InstanceCluster.ListClusters(nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ics)

	ics, resp, err = client.InstanceCluster.ListClusters(nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ics)
}

func TestInstanceClustersService_ListClusters_StatusNotFound(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/admin/clusters", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusNotFound)
	})

	ics, resp, err := client.InstanceCluster.ListClusters(nil, nil)
	require.Error(t, err)
	require.Nil(t, ics)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestInstanceClustersService_GetCluster(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/admin/clusters/9", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": 9,
			"name": "cluster-1",
			"managed": true,
			"enabled": true,
			"domain": null,
			"provider_type": "user",
			"platform_type": "kubernetes",
			"environment_scope": "*",
			"cluster_type": "instance_type",
			"user": {
			  "id": 1,
			  "name": "Administrator",
			  "username": "root",
			  "state": "active",
			  "avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "https://gitlab.example.com/root"
			},
			"platform_kubernetes": {
			  "api_url": "https://example.com",
			  "namespace": null,
			  "authorization_type": "rbac",
			  "ca_cert":"-----BEGIN CERTIFICATE-----IxMDM1MV0ZDJkZjM...-----END CERTIFICATE-----"
			},
			"provider_gcp": null,
			"management_project": null
		  }
		`)
	})

	want := &InstanceCluster{
		ID:               9,
		Name:             "cluster-1",
		Domain:           "",
		Managed:          true,
		ProviderType:     "user",
		PlatformType:     "kubernetes",
		EnvironmentScope: "*",
		ClusterType:      "instance_type",
		User: &User{
			ID:        1,
			Username:  "root",
			Email:     "",
			Name:      "Administrator",
			State:     "active",
			WebURL:    "https://gitlab.example.com/root",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
		},
		PlatformKubernetes: &PlatformKubernetes{
			APIURL:            "https://example.com",
			Token:             "",
			CaCert:            "-----BEGIN CERTIFICATE-----IxMDM1MV0ZDJkZjM...-----END CERTIFICATE-----",
			Namespace:         "",
			AuthorizationType: "rbac",
		},
		ManagementProject: nil,
	}

	ic, resp, err := client.InstanceCluster.GetCluster(9, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ic)

	ic, resp, err = client.InstanceCluster.GetCluster(9, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ic)

	ic, resp, err = client.InstanceCluster.GetCluster(10, nil, nil)
	require.Error(t, err)
	require.Nil(t, ic)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestInstanceClustersService_AddCluster(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/admin/clusters/add", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id": 11,
			"name": "cluster-1",
			"managed": true,
			"enabled": true,
			"domain": null,
			"provider_type": "user",
			"platform_type": "kubernetes",
			"environment_scope": "*",
			"cluster_type": "instance_type",
			"user": {
			  "id": 1,
			  "name": "Administrator",
			  "username": "root",
			  "state": "active",
			  "avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "https://gitlab.example.com/root"
			},
			"platform_kubernetes": {
			  "api_url": "https://example.com",
			  "namespace": null,
			  "authorization_type": "rbac",
			  "ca_cert":"-----BEGIN CERTIFICATE-----IxMDM1MV0ZDJkZjM...-----END CERTIFICATE-----"
			},
			"provider_gcp": null,
			"management_project": null
		  }
		`)
	})

	want := &InstanceCluster{
		ID:               11,
		Name:             "cluster-1",
		Domain:           "",
		Managed:          true,
		ProviderType:     "user",
		PlatformType:     "kubernetes",
		EnvironmentScope: "*",
		ClusterType:      "instance_type",
		User: &User{
			ID:        1,
			Username:  "root",
			Email:     "",
			Name:      "Administrator",
			State:     "active",
			WebURL:    "https://gitlab.example.com/root",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
		},
		PlatformKubernetes: &PlatformKubernetes{
			APIURL:            "https://example.com",
			Token:             "",
			CaCert:            "-----BEGIN CERTIFICATE-----IxMDM1MV0ZDJkZjM...-----END CERTIFICATE-----",
			Namespace:         "",
			AuthorizationType: "rbac",
		},
		ManagementProject: nil,
	}

	ic, resp, err := client.InstanceCluster.AddCluster(nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ic)

	ic, resp, err = client.InstanceCluster.AddCluster(nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ic)
}

func TestInstanceClustersService_AddCluster_StatusInternalServerError(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/admin/clusters/add", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusInternalServerError)
	})

	ic, resp, err := client.InstanceCluster.AddCluster(nil, nil)
	require.Error(t, err)
	require.Nil(t, ic)
	require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestInstanceClustersService_EditCluster(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/admin/clusters/11", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
		  {
			"id": 11,
			"name": "cluster-1",
			"managed": true,
			"enabled": true,
			"domain": null,
			"provider_type": "user",
			"platform_type": "kubernetes",
			"environment_scope": "*",
			"cluster_type": "instance_type",
			"user": {
			  "id": 1,
			  "name": "Administrator",
			  "username": "root",
			  "state": "active",
			  "avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
			  "web_url": "https://gitlab.example.com/root"
			},
			"platform_kubernetes": {
			  "api_url": "https://example.com",
			  "namespace": null,
			  "authorization_type": "rbac",
			  "ca_cert":"-----BEGIN CERTIFICATE-----IxMDM1MV0ZDJkZjM...-----END CERTIFICATE-----"
			},
			"provider_gcp": null,
			"management_project": null
		  }
		`)
	})

	want := &InstanceCluster{
		ID:               11,
		Name:             "cluster-1",
		Domain:           "",
		Managed:          true,
		ProviderType:     "user",
		PlatformType:     "kubernetes",
		EnvironmentScope: "*",
		ClusterType:      "instance_type",
		User: &User{
			ID:        1,
			Username:  "root",
			Email:     "",
			Name:      "Administrator",
			State:     "active",
			WebURL:    "https://gitlab.example.com/root",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
		},
		PlatformKubernetes: &PlatformKubernetes{
			APIURL:            "https://example.com",
			Token:             "",
			CaCert:            "-----BEGIN CERTIFICATE-----IxMDM1MV0ZDJkZjM...-----END CERTIFICATE-----",
			Namespace:         "",
			AuthorizationType: "rbac",
		},
		ManagementProject: nil,
	}

	ic, resp, err := client.InstanceCluster.EditCluster(11, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ic)

	ic, resp, err = client.InstanceCluster.EditCluster(11, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ic)

	ic, resp, err = client.InstanceCluster.EditCluster(12, nil, nil)
	require.Error(t, err)
	require.Nil(t, ic)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestInstanceClustersService_DeleteCluster(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/admin/clusters/11", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.InstanceCluster.DeleteCluster(11, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.InstanceCluster.DeleteCluster(11, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.InstanceCluster.DeleteCluster(12, nil, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
