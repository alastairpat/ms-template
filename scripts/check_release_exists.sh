#!/usr/bin/env bash
export DOCKER_CLI_EXPERIMENTAL=enabled

source ./scripts/common.sh

validate_jq_installed

get_application_version

get_application_name

if docker manifest inspect alastairpaterson/"$APPLICATION_NAME":"$TARGET_VERSION" > /dev/null 2> /dev/null
then
  echo "Version $TARGET_VERSION has already been released! (Did you forget to bump the version in metadata.json?)"
  exit 1
fi

exit 0