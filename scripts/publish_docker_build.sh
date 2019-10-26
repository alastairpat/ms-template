#!/usr/bin/env bash

if ! TARGET_VERSION=$(jq --raw-output '.version' ./metadata.json)
then
  echo "Failed to determine application version. Make sure metadata.json exists in the current directory"
  exit 1
fi

if [ "$TRAVIS_BRANCH" != "master" ]
then
  TARGET_VERSION="$TARGET_VERSION-$(git rev-parse --short HEAD)"
fi

if ! docker build --build-arg buildSha="$(git rev-parse HEAD)" -t alastairpaterson/microservice-template:"$TARGET_VERSION" .
then
  echo "Docker build failed - please see output above"
  exit 1
fi

if ! docker push alastairpaterson/microservice-template:"$TARGET_VERSION"
then
  echo "Docker push failed - please see output above"
  exit 1
fi