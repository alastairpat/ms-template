#!/usr/bin/env bash
export DOCKER_CLI_EXPERIMENTAL=enabled

apt -y update; apt -y install jq

TARGET_VERSION=$(jq --raw-output '.version' ./metadata.json)

if [ $? -ne 0 ]
then
  echo "Failed to determine application version. Make sure metadata.json exists in the current directory"
  exit 1
fi

docker manifest inspect alastairpaterson/microservice-template:"$TARGET_VERSION" > /dev/null 2> /dev/null

IMAGE_EXISTS=$?

if [ -$IMAGE_EXISTS -eq 0 ]
then
  echo "Version $TARGET_VERSION has already been released! (Did you forget to bump the version in metadata.json?)"
  exit 1
fi

exit 0