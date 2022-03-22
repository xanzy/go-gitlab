# go-gitlab

A GitLab API client enabling Go programs to interact with GitLab in a simple and uniform way

[![Build Status](https://github.com/xanzy/go-gitlab/workflows/Lint%20and%20Test/badge.svg)](https://github.com/xanzy/go-gitlab/actions?workflow=Lint%20and%20Test)
[![Sourcegraph](https://sourcegraph.com/github.com/xanzy/go-gitlab/-/badge.svg)](https://sourcegraph.com/github.com/xanzy/go-gitlab?badge)
[![GoDoc](https://godoc.org/github.com/xanzy/go-gitlab?status.svg)](https://godoc.org/github.com/xanzy/go-gitlab)
[![Go Report Card](https://goreportcard.com/badge/github.com/xanzy/go-gitlab)](https://goreportcard.com/report/github.com/xanzy/go-gitlab)

## NOTE

Release v0.6.0 (released on 25-08-2017) no longer supports the older V3 GitLab API. If
you need V3 support, please use the `f-api-v3` branch. This release contains some backwards
incompatible changes that were needed to fully support the V4 GitLab API.

## Coverage

This API client package covers most of the existing GitLab API calls and is updated regularly
to add new and/or missing endpoints. Currently, the following services are supported:

- [x] Applications
- [x] Award Emojis
- [x] Branches
- [x] Broadcast Messages
- [x] Commits
- [x] Container Registry
- [x] Custom Attributes
- [x] Deploy Keys
- [x] Deployments
- [ ] Discussions (threaded comments)
- [x] Environments
- [ ] Epic Issues
- [ ] Epics
- [x] Events
- [x] Feature Flags
- [ ] Geo Nodes
- [x] Generic Packages
- [x] GitLab CI Config Templates
- [x] Gitignores Templates
- [x] Group Access Requests
- [x] Group Issue Boards
- [x] Group Members
- [x] Group Milestones
- [x] Group Wikis
- [x] Group-Level Variables
- [x] Groups
- [x] Instance Clusters
- [x] Invites
- [x] Issue Boards
- [x] Issues
- [x] Jobs
- [x] Keys
- [x] Labels
- [x] License
- [x] Markdown
- [x] Merge Request Approvals
- [x] Merge Requests
- [x] Namespaces
- [x] Notes (comments)
- [x] Notification Settings
- [x] Open Source License Templates
- [x] Packages
- [x] Pages
- [x] Pages Domains
- [x] Personal Access Tokens
- [x] Pipeline Schedules
- [x] Pipeline Triggers
- [x] Pipelines
- [x] Plan limits
- [x] Project Access Requests
- [x] Project Badges
- [x] Project Clusters
- [x] Project Import/export
- [x] Project Members
- [x] Project Milestones
- [x] Project Snippets
- [x] Project-Level Variables
- [x] Projects (including setting Webhooks)
- [x] Protected Branches
- [x] Protected Environments
- [x] Protected Tags
- [x] Repositories
- [x] Repository Files
- [x] Repository Submodules
- [x] Runners
- [x] Search
- [x] Services
- [x] Settings
- [x] Sidekiq Metrics
- [x] System Hooks
- [x] Tags
- [x] Todos
- [x] Users
- [x] Validate CI Configuration
- [x] Version
- [x] Wikis

## Development

### Use a Remote Environment via GitPod

You can choose to use your own development environment if desired, however a `.gitpod.yml` file is included within the repository to allow the use of [GitPod](https://gitpod.io/) easily.
This will allow you to use GitPod's integration with GitHub to quickly start a web-based development environment including Go and Docker, which are necessary
for running tests. To use GitPod's integration, you have two different options described below. After you've completed one of the two options, your development environment
will be ready within a minute or two. As part of starting up, your development environment will automatically start up the `gitlab-ce` container necessary for running
tests.

#### Option 1: Manually navigate to GitPod

You can manually sign in and open your workspace within GitPod by following these steps:

1. Navigate to [GitPod](https://gitpod.io/)
1. Click [Login](https://gitpod.io/login/) if you have an account, or [Sign Up](https://www.gitpod.io/#get-started) if you do not.
1. Click on "Continue with GitHub" and authorize GitPod to access your account.
1. After you've signed in, select "Projects" along the top menu, click on your forked `go-gitlab` project
1. Hover over either the main branch for your fork or the branch you created for your fork, and click "New Workspace"

#### Option 2: Open your GitPod Workspace directly via URL

1. Navigate to your fork of the `go-gitlab` project in GitHub
1. Select the branch you want to develop
1. Add `https://gitpod.io/#` to the front of your URL

Your workspace will automatically open the repository and branch that you selected in GitHub.

### Testing individual go-gitlab services with GitPods

Once a GitPod has started up with the gitlab-ce container service running, port 8080 will be exposed for API calls where each individual go-gitlab service can be tested against. Please note that the `ce` container service is started by default and to test any `ee` go-gitlab service you will need to have a valid license file provided in the top level directory named `Gitlab-license.txt`. Then modify the `GNUmakefile` to set the `SERVICE` variable to `gitlab-ee` instead of `gitlab-ce`. Once these modifications are done run `make testacc-down` to stop the current ce container service and then run `make testacc-up` to start the new ee container service

- Connection Information
  - PAT/API Token: `ACCTEST1234567890123`
  - API URL: `http://127.0.0.1:8080/api/v4/`

1. Create a dummy testing file (something.go, main.go, etc.) where you will be testing the services you are developing
2. Follow the format of one of the services under the examples folder
3. Set the connection information to the above for the GitLab client
4. On the terminal enter `go run <filename.go>` to run the test file that was created against the local gitlab container service

## Usage

```go
import "github.com/xanzy/go-gitlab"
```

Construct a new GitLab client, then use the various services on the client to
access different parts of the GitLab API. For example, to list all
users:

```go
git, err := gitlab.NewClient("yourtokengoeshere")
if err != nil {
  log.Fatalf("Failed to create client: %v", err)
}
users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{})
```

There are a few `With...` option functions that can be used to customize
the API client. For example, to set a custom base URL:

```go
git, err := gitlab.NewClient("yourtokengoeshere", gitlab.WithBaseURL("https://git.mydomain.com/api/v4"))
if err != nil {
  log.Fatalf("Failed to create client: %v", err)
}
users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{})
```

Some API methods have optional parameters that can be passed. For example,
to list all projects for user "svanharmelen":

```go
git := gitlab.NewClient("yourtokengoeshere")
opt := &ListProjectsOptions{Search: gitlab.String("svanharmelen")}
projects, _, err := git.Projects.ListProjects(opt)
```

### Examples

The [examples](https://github.com/xanzy/go-gitlab/tree/master/examples) directory
contains a couple for clear examples, of which one is partially listed here as well:

```go
package main

import (
  "log"

  "github.com/xanzy/go-gitlab"
)

func main() {
  git, err := gitlab.NewClient("yourtokengoeshere")
  if err != nil {
    log.Fatalf("Failed to create client: %v", err)
  }

  // Create new project
  p := &gitlab.CreateProjectOptions{
    Name:                 gitlab.String("My Project"),
    Description:          gitlab.String("Just a test project to play with"),
    MergeRequestsEnabled: gitlab.Bool(true),
    SnippetsEnabled:      gitlab.Bool(true),
    Visibility:           gitlab.Visibility(gitlab.PublicVisibility),
  }
  project, _, err := git.Projects.CreateProject(p)
  if err != nil {
    log.Fatal(err)
  }

  // Add a new snippet
  s := &gitlab.CreateProjectSnippetOptions{
    Title:           gitlab.String("Dummy Snippet"),
    FileName:        gitlab.String("snippet.go"),
    Content:         gitlab.String("package main...."),
    Visibility:      gitlab.Visibility(gitlab.PublicVisibility),
  }
  _, _, err = git.ProjectSnippets.CreateSnippet(project.ID, s)
  if err != nil {
    log.Fatal(err)
  }
}
```

For complete usage of go-gitlab, see the full [package docs](https://godoc.org/github.com/xanzy/go-gitlab).

## ToDo

- The biggest thing this package still needs is tests :disappointed:

## Issues

- If you have an issue: report it on the [issue tracker](https://github.com/xanzy/go-gitlab/issues)

## Author

Sander van Harmelen (<sander@vanharmelen.nl>)

## License

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>
