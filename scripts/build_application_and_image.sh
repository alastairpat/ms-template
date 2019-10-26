#!/usr/bin/env bash

source ./scripts/common.sh

validate_jq_installed

get_application_name

if ! CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o output/service
then
  echo "Go build failed - please see output above"
  exit 1
fi

if ! docker build -t "$APPLICATION_NAME" --build-arg buildSha="$(git rev-parse HEAD)" .
then
  echo "Docker build failed - please see output above"
  exit 1
fi