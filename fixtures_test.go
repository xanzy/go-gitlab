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

const (
	// exampleDetailResponse provides fixture for Runners tests.
	exampleDetailResponse = `{
		"active": true,
		"architecture": null,
		"description": "test-1-20150125-test",
		"run_untagged": true,
		"id": 6,
		"is_shared": false,
		"contacted_at": "2016-01-25T16:39:48.066Z",
		"name": null,
		"online": true,
		"status": "online",
		"platform": null,
		"projects": [
			{
				"id": 1,
				"name": "GitLab Community Edition",
				"name_with_namespace": "GitLab.org / GitLab Community Edition",
				"path": "gitlab-ce",
				"path_with_namespace": "gitlab-org/gitlab-ce"
			}
		],
		"token": "205086a8e3b9a2b818ffac9b89d102",
		"revision": null,
		"tag_list": [
			"ruby",
			"mysql"
		],
		"version": null,
		"access_level": "ref_protected",
		"maximum_timeout": 3600,
		"locked": false
	}`

	// exampleEventUserName provides a fixture for a event user's name.
	exampleEventUserName = "John Smith"

	// exampleEventUserUsername provides a ficture for the event username.
	exampleEventUserUsername = "jsmith"

	// exampleProjectName provides a fixture for a project name.
	exampleProjectName = "example-project"

	// exampleRegisterNewRunner provides fixture for Runners tests.
	exampleRegisterNewRunner = `{
		"id": 12345,
		"token": "6337ff461c94fd3fa32ba3b1ff4125"
	}`

	// exampleReleaseLink provides fixture for Release Links tests.
	exampleReleaseLink = `{
		"id":1,
		"name":"awesome-v0.2.dmg",
		"url":"http://192.168.10.15:3000",
		"direct_asset_url": "http://192.168.10.15:3000/namespace/example/-/releases/v0.1/downloads/awesome-v0.2.dmg",
		"external":true,
		"link_type": "other"
	}`

	// exampleReleaseLinkList provides fixture for Release Links tests.
	exampleReleaseLinkList = `[
		{
			"id": 2,
			"name": "awesome-v0.2.msi",
			"url": "http://192.168.10.15:3000/msi",
			"external": true
		},
		{
			"id": 1,
			"name": "awesome-v0.2.dmg",
			"url": "http://192.168.10.15:3000",
			"direct_asset_url": "http://192.168.10.15:3000/namespace/example/-/releases/v0.1/downloads/awesome-v0.2.dmg",
			"external": false,
			"link_type": "other"
		}
	]`

	// exampleReleaseListResponse provides fixture for Releases tests.
	exampleReleaseListResponse = `[
		{
			"tag_name": "v0.2",
			"description": "description",
			"name": "Awesome app v0.2 beta",
			"description_html": "html",
			"created_at": "2019-01-03T01:56:19.539Z",
			"author": {
			"id": 1,
			"name": "Administrator",
			"username": "root",
			"state": "active",
			"avatar_url": "https://www.gravatar.com/avatar",
			"web_url": "http://localhost:3000/root"
			},
			"commit": {
			"id": "079e90101242458910cccd35eab0e211dfc359c0",
			"short_id": "079e9010",
			"title": "Update README.md",
			"created_at": "2019-01-03T01:55:38.000Z",
			"parent_ids": [
				"f8d3d94cbd347e924aa7b715845e439d00e80ca4"
			],
			"message": "Update README.md",
			"author_name": "Administrator",
			"author_email": "admin@example.com",
			"authored_date": "2019-01-03T01:55:38.000Z",
			"committer_name": "Administrator",
			"committer_email": "admin@example.com",
			"committed_date": "2019-01-03T01:55:38.000Z"
			},
			"assets": {
			"count": 4,
			"sources": [
				{
				"format": "zip",
				"url": "http://localhost:3000/archive/v0.2/awesome-app-v0.2.zip"
				},
				{
				"format": "tar.gz",
				"url": "http://localhost:3000/archive/v0.2/awesome-app-v0.2.tar.gz"
				}
			],
			"links": [
				{
				"id": 2,
				"name": "awesome-v0.2.msi",
				"url": "http://192.168.10.15:3000/msi",
				"external": true
				},
				{
				"id": 1,
				"name": "awesome-v0.2.dmg",
				"url": "http://192.168.10.15:3000",
				"external": true
				}
			]
			}
		},
		{
			"tag_name": "v0.1",
			"description": "description",
			"name": "Awesome app v0.1 alpha",
			"description_html": "description_html",
			"created_at": "2019-01-03T01:55:18.203Z",
			"author": {
			"id": 1,
			"name": "Administrator",
			"username": "root",
			"state": "active",
			"avatar_url": "https://www.gravatar.com/avatar",
			"web_url": "http://localhost:3000/root"
			},
			"commit": {
			"id": "f8d3d94cbd347e924aa7b715845e439d00e80ca4",
			"short_id": "f8d3d94c",
			"title": "Initial commit",
			"created_at": "2019-01-03T01:53:28.000Z",
			"parent_ids": [],
			"message": "Initial commit",
			"author_name": "Administrator",
			"author_email": "admin@example.com",
			"authored_date": "2019-01-03T01:53:28.000Z",
			"committer_name": "Administrator",
			"committer_email": "admin@example.com",
			"committed_date": "2019-01-03T01:53:28.000Z"
			},
			"assets": {
			"count": 2,
			"sources": [
				{
				"format": "zip",
				"url": "http://localhost:3000/archive/v0.1/awesome-app-v0.1.zip"
				},
				{
				"format": "tar.gz",
				"url": "http://localhost:3000/archive/v0.1/awesome-app-v0.1.tar.gz"
				}
			],
			"links": []
			}
		}
	]`

	// exampleReleaseName provides a fixture for a release name.
	exampleReleaseName = "awesome-v0.2.dmg"

	// exampleReleaseResponse provides fixture for Releases tests.
	exampleReleaseResponse = `{
		"tag_name": "v0.1",
		"description": "description",
		"name": "Awesome app v0.1 alpha",
		"description_html": "description_html",
		"created_at": "2019-01-03T01:55:18.203Z",
		"author": {
			"id": 1,
			"name": "Administrator",
			"username": "root",
			"state": "active",
			"avatar_url": "https://www.gravatar.com/avatar/",
			"web_url": "http://localhost:3000/root"
		},
		"commit": {
			"id": "f8d3d94cbd347e924aa7b715845e439d00e80ca4",
			"short_id": "f8d3d94c",
			"title": "Initial commit",
			"created_at": "2019-01-03T01:53:28.000Z",
			"parent_ids": [],
			"message": "Initial commit",
			"author_name": "Administrator",
			"author_email": "admin@example.com",
			"authored_date": "2019-01-03T01:53:28.000Z",
			"committer_name": "Administrator",
			"committer_email": "admin@example.com",
			"committed_date": "2019-01-03T01:53:28.000Z"
		},
		"assets": {
			"count": 2,
			"sources": [
			{
				"format": "zip",
				"url": "http://localhost:3000/archive/v0.1/awesome-app-v0.1.zip"
			},
			{
				"format": "tar.gz",
				"url": "http://localhost:3000/archive/v0.1/awesome-app-v0.1.tar.gz"
			}
			],
			"links": []
		}
	}`

	// exampleReleaseResponse provides fixture for Releases tests.
	exampleReleaseWithMetadataResponse = `{
		"tag_name": "v0.1.2+example-metadata",
		"description": "description",
		"name": "Awesome app v0.1 alpha",
		"description_html": "description_html",
		"created_at": "2019-01-03T01:55:18.203Z",
		"author": {
			"id": 1,
			"name": "Administrator",
			"username": "root",
			"state": "active",
			"avatar_url": "https://www.gravatar.com/avatar/",
			"web_url": "http://localhost:3000/root"
		},
		"commit": {
			"id": "f8d3d94cbd347e924aa7b715845e439d00e80ca4",
			"short_id": "f8d3d94c",
			"title": "Initial commit",
			"created_at": "2019-01-03T01:53:28.000Z",
			"parent_ids": [],
			"message": "Initial commit",
			"author_name": "Administrator",
			"author_email": "admin@example.com",
			"authored_date": "2019-01-03T01:53:28.000Z",
			"committer_name": "Administrator",
			"committer_email": "admin@example.com",
			"committed_date": "2019-01-03T01:53:28.000Z"
		},
		"assets": {
			"count": 2,
			"sources": [
			{
				"format": "zip",
				"url": "http://localhost:3000/archive/v0.1/awesome-app-v0.1.zip"
			},
			{
				"format": "tar.gz",
				"url": "http://localhost:3000/archive/v0.1/awesome-app-v0.1.tar.gz"
			}
			],
			"links": []
		}
	}`

	// exampleTagName provides a fixture for a tag name.
	exampleTagName = "v0.1"

	// exampleTagName provides a fixture for a tag name.
	exampleTagNameWithMetadata = "v0.1.2+example-metadata"
)
