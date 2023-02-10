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
	// exampleChangelogResponse provides fixture for Changelog tests.
	exampleChangelogResponse = `{
    "notes": "## 1.0.0 (2021-11-17)\n\n### feature (2 changes)\n\n- [Title 2](namespace13/project13@ad608eb642124f5b3944ac0ac772fecaf570a6bf) ([merge request](namespace13/project13!2))\n- [Title 1](namespace13/project13@3c6b80ff7034fa0d585314e1571cc780596ce3c8) ([merge request](namespace13/project13!1))\n"
  }`

	// exampleCommitMessage provides fixture for a commit message.
	exampleCommitMessage = "Merge branch 'some-feature' into 'master'\n\nRelease v1.0.0\n\nSee merge request jsmith/example!1"

	// exampleCommitTitle provides fixture for a commit title.
	exampleCommitTitle = "Merge branch 'some-feature' into 'master'"

	// exampleDetailResponse provides fixture for Runners tests.
	exampleDetailResponse = `{
		"active": true,
		"architecture": null,
		"description": "test-1-20150125-test",
		"run_untagged": true,
		"id": 6,
		"is_shared": false,
		"runner_type": "project_type",
		"contacted_at": "2016-01-25T16:39:48.166Z",
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

	// exampleRunnerJob provides fixture for the list runner jobs test.
	exampleListRunnerJobs = `
  [
    {
      "id": 1,
      "status": "failed",
      "stage": "test",
      "name": "run_tests",
      "ref": "master",
      "tag": false,
      "coverage": null,
      "allow_failure": false,
      "created_at": "2021-10-22T11:59:25.201Z",
      "started_at": "2021-10-22T11:59:33.660Z",
      "finished_at": "2021-10-22T15:59:25.201Z",
      "duration": 171.540594,
      "queued_duration": 2.535766,
      "user": {
        "id": 368,
        "name": "John SMITH",
        "username": "john.smith",
        "state": "blocked",
        "avatar_url": "https://gitlab.example.com/uploads/-/system/user/avatar/368/avatar.png",
        "web_url": "https://gitlab.example.com/john.smith",
        "bio": "",
        "location": "",
        "public_email": "john.smith@example.com",
        "skype": "",
        "linkedin": "",
        "twitter": "",
        "website_url": "",
        "organization": "",
        "job_title": "",
        "pronouns": null,
        "bot": false,
        "work_information": null,
        "bio_html": ""
      },
      "commit": {
        "id": "6c016b801a88f4bd31f927fc045b5c746a6f823e",
        "short_id": "6c016b80",
        "created_at": "2018-03-21T14:41:00.000Z",
        "parent_ids": [
          "6008b4902d40799ab11688e502d9f1f27f6d2e18"
        ],
        "title": "Update env for specific runner",
        "message": "Update env for specific runner\n",
        "author_name": "John SMITH",
        "author_email": "john.smith@example.com",
        "authored_date": "2018-03-21T14:41:00.000Z",
        "committer_name": "John SMITH",
        "committer_email": "john.smith@example.com",
        "committed_date": "2018-03-21T14:41:00.000Z",
        "web_url": "https://gitlab.example.com/awesome/packages/common/-/commit/6c016b801a88f4bd31f927fc045b5c746a6f823e"
      },
      "pipeline": {
        "id": 8777,
        "project_id": 3252,
        "sha": "6c016b801a88f4bd31f927fc045b5c746a6f823e",
        "ref": "master",
        "status": "failed",
        "source": "push",
        "created_at": "2018-03-21T13:41:15.356Z",
        "updated_at": "2018-03-21T15:12:52.021Z",
        "web_url": "https://gitlab.example.com/awesome/packages/common/-/pipelines/8777"
      },
      "web_url": "https://gitlab.example.com/awesome/packages/common/-/jobs/14606",
      "project": {
        "id": 3252,
        "description": "Common nodejs paquet for producer",
        "name": "common",
        "name_with_namespace": "awesome",
        "path": "common",
        "path_with_namespace": "awesome",
        "created_at": "2018-02-13T09:21:48.107Z"
      }
    }
  ]`

	// exampleProjectName provides a fixture for a project name.
	exampleProjectName = "example-project"

	// exampleProjectStatusChecks provides a fixture for a project status checks.
	exampleProjectStatusChecks = `[
		{
			"id": 1,
			"name": "Compliance Check",
			"project_id": 6,
			"external_url": "https://gitlab.com/example/test.json",
			"protected_branches": [
				{
					"id": 14,
					"project_id": 6,
					"name": "master",
					"created_at": "2020-10-12T14:04:50.787Z",
					"updated_at": "2020-10-12T14:04:50.787Z",
					"code_owner_approval_required": false
				}
			]
		}
	]`

	// exampleRegisterNewRunner provides fixture for Runners tests.
	exampleRegisterNewRunner = `{
		"id": 12345,
		"token": "6337ff461c94fd3fa32ba3b1ff4125",
		"token_expires_at": "2016-01-25T16:39:48.166Z"
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

	// exampleStatusChecks provides a fixture for status checks for a merge request.
	exampleStatusChecks = `[
    {
        "id": 2,
        "name": "Rule 1",
        "external_url": "https://gitlab.com/test-endpoint",
        "status": "approved"
    },
    {
        "id": 1,
        "name": "Rule 2",
        "external_url": "https://gitlab.com/test-endpoint-2",
        "status": "pending"
    }
	]`

	// exampleTagName provides a fixture for a tag name.
	exampleTagName = "v0.1"

	// exampleTagName provides a fixture for a tag name.
	exampleTagNameWithMetadata = "v0.1.2+example-metadata"
)
