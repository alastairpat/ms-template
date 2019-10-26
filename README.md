# Example Service 

[![Build Status](https://travis-ci.com/alastairpat/ms-template.svg?branch=master)](https://travis-ci.com/alastairpat/ms-template)

[Example Service on Docker Hub](https://hub.docker.com/r/alastairpaterson/microservice-template)

The Example Service provides an example microservice template.
It is designed to illustrate best practices, and can be adopted for use by other teams
with minimal changes.

## Ownership

Example Service is owned and supported by the Example Team.

* Email: [example.team@example.com](mailto:example.team@example.com)
* Slack: [#example-team](https://slack.com/app_redirect?channel=example-team)
* Documentation: [Example Team Confluence Space](https://confluence.example.com/example)

## Local Development

Example Service is written in GoLang and deployed as a Docker image. 
It is assumed your computer is set up for both Go development and Docker.

* Setup
  - After cloning this repository, you should create a branch for your work. Changes to the
    `master` branch, except via pull request, are not possible.
  - After creating your branch, you should increment the version number in 
  `metadata.json`. 
* Building and Running
  - You can build Example Service by running `go build` from the root directory. Similarly,
    you can run `go test [-cov]` to run unit tests (optionally with coverage checking).
  - The `Dockerfile` expects to find a Linux binary named `service` in the `output` directory.
  - You can build the application and Dockerfile by invoking `./scripts/build_application_and_image.sh` 
    from the root of this project

## Required Environment Variables

At present, Example Service requires no environmental variables to run. However, if it did,
you should document them here:

| Variable Name     | Description        | Default Value | Required? |
| ----------------- | ------------------ | ------------- | --------- |
| `TZ`              | Container timezone | UTC           | No        |
| `DATASOURCE_NAME` | Datasource name    | (null)        | Yes       |

## Continuous Integration

* Any pushed commits will automatically be tested with TravisCI. 
* If the build passes, it will also generate a new Docker image that is pushed to Docker Hub. 
  - If the build is against a branch, it will be tagged as `version-SHA short hash` (for example, `1.2.3-a1b2c3d`).
  - If the build is against `master`, it will simply be tagged as `version`.

## Release Management

* All changes to Example Service must happen via a pull request. 
* [Semantic Versioning](https://semver.org/) is used for version numbers.
* Version numbers on the `master` branch must be unique. Failure to use unique
  version numbers will result in a failed build.