//
// Copyright 2020, Serena Fang
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
	"time"
)

// InstanceClustersService handles communication with the
// instance clusters related methods of the GitLab API.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/instance_clusters.html
type InstanceClustersService struct {
	client *Client
}

// InstanceCluster represents a GitLab Instance Cluster.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/instance_clusters.html
type InstanceCluster struct {
	ID                 int                 `json:"id"`
	Name               string              `json:"name"`
	Domain             string              `json:"domain"`
	CreatedAt          *time.Time          `json:"created_at"`
	ProviderType       string              `json:"provider_type"`
	PlatformType       string              `json:"platform_type"`
	EnvironmentScope   string              `json:"environment_scope"`
	ClusterType        string              `json:"cluster_type"`
	User               *User               `json:"user"`
	PlatformKubernetes *PlatformKubernetes `json:"platform_kubernetes"`
	ManagementProject  *ManagementProject  `json:"management_project"`
}

func (v InstanceCluster) String() string {
	return Stringify(v)
}

// PlatformKubernetes represents a GitLab Instance Cluster PlatformKubernetes.
type InstanceClusterPlatformKubernetes struct {
	APIURL            string `json:"api_url"`
	Token             string `json:"token"`
	CaCert            string `json:"ca_cert"`
	Namespace         string `json:"namespace"`
	AuthorizationType string `json:"authorization_type"`
}

// ListClusters gets a list of all instance clusters.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/instance_clusters.html#list-instance-clusters
func (s *InstanceClustersService) ListClusters(options ...RequestOptionFunc) ([]*InstanceCluster, *Response, error) {
	u := fmt.Sprintf("admin/clusters")

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var ics []*InstanceCluster
	resp, err := s.client.Do(req, &ics)
	if err != nil {
		return nil, resp, err
	}

	return ics, resp, err
}

// GetCluster gets an instance cluster.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/instance_clusters.html#get-a-single-instance-cluster
func (s *InstanceClustersService) GetCluster(cluster int, options ...RequestOptionFunc) (*InstanceCluster, *Response, error) {
	u := fmt.Sprintf("admin/clusters/%d", cluster)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	ic := new(InstanceCluster)
	resp, err := s.client.Do(req, &ic)
	if err != nil {
		return nil, resp, err
	}

	return ic, resp, err
}

// AddInstanceClusterOptions represents the available AddCluster() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/instance_clusters.html#add-existing-cluster-to-instance
type AddInstanceClusterOptions struct {
	Name                *string                               `url:"name,omitempty" json:"name,omitempty"`
	Domain              *string                               `url:"domain,omitempty" json:"domain,omitempty"`
	Enabled             *bool                                 `url:"enabled,omitempty" json:"enabled,omitempty"`
	Managed             *bool                                 `url:"managed,omitempty" json:"managed,omitempty"`
	EnvironmentScope    *string                               `url:"environment_scope,omitempty" json:"environment_scope,omitempty"`
	PlatformKubernetes  *AddInstancePlatformKubernetesOptions `url:"platform_kubernetes_attributes,omitempty" json:"platform_kubernetes_attributes,omitempty"`
	ManagementProjectID *string                               `url:"management_project_id,omitempty" json:"management_project_id,omitempty"`
}

// AddInstancePlatformKubernetesOptions represents the available PlatformKubernetes options for adding.
type AddInstancePlatformKubernetesOptions struct {
	APIURL            *string `url:"api_url,omitempty" json:"api_url,omitempty"`
	Token             *string `url:"token,omitempty" json:"token,omitempty"`
	CaCert            *string `url:"ca_cert,omitempty" json:"ca_cert,omitempty"`
	Namespace         *string `url:"namespace,omitempty" json:"namespace,omitempty"`
	AuthorizationType *string `url:"authorization_type,omitempty" json:"authorization_type,omitempty"`
}

// AddCluster adds an existing cluster to the instance.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/instance_clusters.html#add-existing-instance-cluster
func (s *InstanceClustersService) AddCluster(opt *AddInstanceClusterOptions, options ...RequestOptionFunc) (*InstanceCluster, *Response, error) {
	u := fmt.Sprintf("admin/clusters/add")

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	ic := new(InstanceCluster)
	resp, err := s.client.Do(req, ic)
	if err != nil {
		return nil, resp, err
	}

	return ic, resp, err
}

// EditInstanceClusterOptions represents the available EditCluster() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/instance_clusters.html#edit-instance-cluster
type EditInstanceClusterOptions struct {
	Name                *string                                `url:"name,omitempty" json:"name,omitempty"`
	Domain              *string                                `url:"domain,omitempty" json:"domain,omitempty"`
	EnvironmentScope    *string                                `url:"environment_scope,omitempty" json:"environment_scope,omitempty"`
	ManagementProjectID *string                                `url:"management_project_id,omitempty" json:"management_project_id,omitempty"`
	PlatformKubernetes  *EditInstancePlatformKubernetesOptions `url:"platform_kubernetes_attributes,omitempty" json:"platform_kubernetes_attributes,omitempty"`
}

// EditInstancePlatformKubernetesOptions represents the available PlatformKubernetes options for editing.
type EditInstancePlatformKubernetesOptions struct {
	APIURL    *string `url:"api_url,omitempty" json:"api_url,omitempty"`
	Token     *string `url:"token,omitempty" json:"token,omitempty"`
	CaCert    *string `url:"ca_cert,omitempty" json:"ca_cert,omitempty"`
	Namespace *string `url:"namespace,omitempty" json:"namespace,omitempty"`
}

// EditCluster updates an existing instance cluster.
// Note: name, api_url, ca_cert and token can only be updated if the cluster
// was added through the 'Add existing Kubernetes' cluster option or through the
// 'Add existing instance' cluster endpoint.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/instance_clusters.html#edit-instance-cluster
func (s *InstanceClustersService) EditCluster(cluster int, opt *EditInstanceClusterOptions, options ...RequestOptionFunc) (*InstanceCluster, *Response, error) {
	u := fmt.Sprintf("admin/clusters/%d", cluster)

	req, err := s.client.NewRequest("PUT", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	ic := new(InstanceCluster)
	resp, err := s.client.Do(req, ic)
	if err != nil {
		return nil, resp, err
	}

	return ic, resp, err
}

// DeleteCluster deletes an existing instance cluster.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/instance_clusters.html#delete-instance-cluster
func (s *InstanceClustersService) DeleteCluster(cluster int, options ...RequestOptionFunc) (*Response, error) {
	u := fmt.Sprintf("admin/clusters/%d", cluster)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
