#!/usr/bin/env bash

if ! TARGET_VERSION=$(jq --raw-output '.version' ./metadata.json)
then
  echo "Failed to determine application version. Make sure metadata.json exists in the current directory"
  exit 1
fi

if ! CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
then
  echo "Failed to determine current git branch"
  exit 1
fi

if [ "$CURRENT_BRANCH" != "master" ]
then
  TARGET_VERSION="$TARGET_VERSION-$(git rev-parse --short HEAD)"
fi

docker build --build-arg buildSha="$(git rev-parse HEAD)" -t alastairpaterson/microservice-template:"$TARGET_VERSION"

docker push alastairpaterson/microservice-template:"$TARGET_VERSION"