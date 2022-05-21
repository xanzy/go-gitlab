# Contributing to go-gitlab

First of all, thanks for considering to contribute to the [go-gitlab library](https://github.com/xanzy/go-gitlab) :tada:.

## Repository Structure

This library aims to provide a go client providing a 1:1 mapping to the [GitLab REST API](https://docs.gitlab.com/ee/api/).
Every REST API resource has it's own Go file with an associated `_test.go` file. Acceptance Tests are located in the `acc/`
directory. When implementing a new endpoint, please inspire yourself with an already existing one.

## Unit Tests

Unit Tests are implemented in the corresponding `_test.go` file located in the root of the project.
They can be run using `go test -v .` or `make test`.

## Acceptance Tests

The Acceptance Tests run against a real GitLab instance. At the moment only a GitLab CE instance is supported in the pipeline,
but given a license key you can easily run the test suite against an EE instance.

To start a GitLab CE instance you can use `make testacc-up`. It requires `docker-compose`.
Once an instance is running you can run the test suite with `make testacc`.
The instance can be destroyed using `make testacc-down`.