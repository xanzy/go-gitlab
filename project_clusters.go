//
// Copyright 2019, Matej Velikonja
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
	"net/url"
	"time"
)

// ProjectClustersService handles communication with the
// project clusters related methods of the GitLab API.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_clusters.html
type ProjectClustersService struct {
	client *Client
}

// ProjectCluster represents a GitLab Project Cluster.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_clusters.html
type ProjectCluster struct {
	ID                 int                `json:"id,omitempty"`
	Name               string             `json:"name,omitempty"`
	Enabled            *bool              `json:"enabled,omitempty"`
	CreatedAt          *time.Time         `json:"created_at,omitempty"`
	ProviderType       string             `json:"provider_type,omitempty"`
	PlatformType       string             `json:"platform_type,omitempty"`
	EnvironmentScope   string             `json:"environment_scope,omitempty"`
	ClusterType        string             `json:"cluster_type,omitempty"`
	User               *User              `json:"user,omitempty"`
	PlatformKubernetes PlatformKubernetes `json:"platform_kubernetes,omitempty"`
	Project            *Project           `json:"project,omitempty"`
}

func (v ProjectCluster) String() string {
	return Stringify(v)
}

// PlatformKubernetes represents a GitLab Project Cluster PlatformKubernetes.
type PlatformKubernetes struct {
	APIURL            string `json:"api_url,omitempty"`
	Token             string `json:"token,omitempty"`
	CaCert            string `json:"ca_cert,omitempty"`
	Namespace         string `json:"namespace,omitempty"`
	AuthorizationType string `json:"authorization_type,omitempty"`
}

// ListClusters gets a list of all clusters in a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_clusters.html#list-project-clusters
func (s *ProjectClustersService) ListClusters(pid interface{}, options ...OptionFunc) ([]*ProjectCluster, *Response, error) {
	projectID, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/clusters", url.QueryEscape(projectID))

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var vs []*ProjectCluster
	resp, err := s.client.Do(req, &vs)
	if err != nil {
		return nil, resp, err
	}

	return vs, resp, err
}

// GetCluster gets a cluster.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_clusters.html#get-a-single-project-cluster
func (s *ProjectClustersService) GetCluster(pid interface{}, clusterID int, options ...OptionFunc) (*ProjectCluster, *Response, error) {
	projectID, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/clusters/%d", url.QueryEscape(projectID), clusterID)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var vs *ProjectCluster
	resp, err := s.client.Do(req, &vs)
	if err != nil {
		return nil, resp, err
	}

	return vs, resp, err
}

// AddClusterOptions represents the available AddCluster() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_clusters.html#add-existing-cluster-to-project
type AddClusterOptions struct {
	Name               string             `json:"name,omitempty"`
	Enabled            *bool              `json:"enabled,omitempty"`
	EnvironmentScope   string             `json:"environment_scope,omitempty"`
	PlatformKubernetes PlatformKubernetes `json:"platform_kubernetes_attributes,omitempty"`
}

// AddCluster adds an existing cluster to the project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_clusters.html#add-existing-cluster-to-project
func (s *ProjectClustersService) AddCluster(pid interface{}, opt *AddClusterOptions, options ...OptionFunc) (*ProjectCluster, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/clusters/user", url.QueryEscape(project))

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}
	//b, _ := ioutil.ReadAll(req.Body)
	//fmt.Println(string(b))
	pc := new(ProjectCluster)
	resp, err := s.client.Do(req, pc)
	if err != nil {
		return nil, resp, err
	}

	return pc, resp, err
}

// EditClusterOptions represents the available EditCluster() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_clusters.html#edit-project-cluster
type EditClusterOptions struct {
	Name               string             `json:"name"`
	EnvironmentScope   string             `json:"environment_scope,omitempty"`
	PlatformKubernetes PlatformKubernetes `json:"platform_kubernetes_attributes,omitempty"`
}

// EditCluster updates an existing project cluster.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_clusters.html#edit-project-cluster
func (s *ProjectClustersService) EditCluster(pid interface{}, cid int, opt *EditClusterOptions, options ...OptionFunc) (*ProjectCluster, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/clusters/%d", url.QueryEscape(project), cid)

	req, err := s.client.NewRequest("PUT", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pc := new(ProjectCluster)
	resp, err := s.client.Do(req, pc)
	if err != nil {
		return nil, resp, err
	}

	return pc, resp, err
}

// DeleteCluster deletes an existing project cluster.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_clusters.html#delete-project-cluster
func (s *ProjectClustersService) DeleteCluster(pid interface{}, cid int, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/clusters/%d", url.QueryEscape(project), cid)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
