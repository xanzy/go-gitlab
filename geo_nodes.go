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

// GeoNode represents a GitLab Geo Node.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/geo_nodes.html
type GeoNode struct {
	ID                               int           `json:"id"`
	Name                             string        `json:"name"`
	URL                              string        `json:"url"`
	InternalURL                      string        `json:"internal_url"`
	Primary                          bool          `json:"primary"`
	Enabled                          bool          `json:"enabled"`
	Current                          bool          `json:"current"`
	FilesMaxCapacity                 int           `json:"files_max_capacity"`
	ReposMaxCapacity                 int           `json:"repos_max_capacity"`
	VerificationMaxCapacity          int           `json:"verification_max_capacity"`
	SelectiveSyncType                string        `json:"selective_sync_type"`
	SelectiveSyncShards              []interface{} `json:"selective_sync_shards"`
	SelectiveSyncNamespaceIds        []int         `json:"selective_sync_namespace_ids"`
	MinimumReverificationInterval    int           `json:"minimum_reverification_interval"`
	ContainerRepositoriesMaxCapacity int           `json:"container_repositories_max_capacity"`
	SyncObjectStorage                bool          `json:"sync_object_storage"`
	CloneProtocol                    string        `json:"clone_protocol"`
	WebEditURL                       string        `json:"web_edit_url"`
	WebGeoProjectsURL                string        `json:"web_geo_projects_url"`
	Links                            GeoNodeLinks  `json:"_links"`
}

// GeoNodeLinks represents links for GitLab GeoNode.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/geo_nodes.html
type GeoNodeLinks struct {
	Self   string `json:"self"`
	Status string `json:"status"`
	Repair string `json:"repair"`
}

// GeoNodesService handles communication with Geo Nodes related methods of GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/geo_nodes.html
type GeoNodesService struct {
	client *Client
}

// CreateGeoNodesOptions represents the available CreateGeoNode() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/geo_nodes.html#create-a-new-geo-node
type CreateGeoNodesOptions struct {
	Primary                          *bool     `url:"primary,omitempty" json:"primary,omitempty"`
	Enabled                          *bool     `url:"enabled,omitempty" json:"enabled,omitempty"`
	Name                             *string   `url:"name,omitempty" json:"name,omitempty"`
	URL                              *string   `url:"url,omitempty" json:"url,omitempty"`
	InternalURL                      *string   `url:"internal_url,omitempty" json:"internal_url,omitempty"`
	FilesMaxCapacity                 *int      `url:"files_max_capacity,omitempty" json:"files_max_capacity,omitempty"`
	ReposMaxCapacity                 *int      `url:"repos_max_capacity,omitempty" json:"repos_max_capacity,omitempty"`
	VerificationMaxCapacity          *int      `url:"verification_max_capacity,omitempty" json:"verification_max_capacity,omitempty"`
	ContainerRepositoriesMaxCapacity *int      `url:"container_repositories_max_capacity,omitempty" json:"container_repositories_max_capacity,omitempty"`
	SyncObjectStorage                *bool     `url:"sync_object_storage,omitempty" json:"sync_object_storage,omitempty"`
	SelectiveSyncType                *string   `url:"selective_sync_type,omitempty" json:"selective_sync_type,omitempty"`
	SelectiveSyncShards              []*string `url:"selective_sync_shards,omitempty" json:"selective_sync_shards,omitempty"`
	SelectiveSyncNamespaceIds        []*string `url:"selective_sync_namespace_ids,omitempty" json:"selective_sync_namespace_ids,omitempty"`
	MinimumReverificationInterval    *int      `url:"minimum_reverification_interval,omitempty" json:"minimum_reverification_interval,omitempty"`
}
