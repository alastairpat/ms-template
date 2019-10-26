#!/usr/bin/env bash

source ./scripts/common.sh

validate_jq_installed

get_application_version

get_application_name

if [ "$TRAVIS_BRANCH" != "master" ]
then
  TARGET_VERSION="$TARGET_VERSION-$(git rev-parse --short HEAD)"
fi

if ! docker tag "$APPLICATION_NAME" alastairpaterson/"$APPLICATION_NAME":"$TARGET_VERSION"
then
  echo "Docker tag failed - please see output above"
  exit 1
fi

if ! docker push alastairpaterson/"$APPLICATION_NAME":"$TARGET_VERSION"
then
  echo "Docker push failed - please see output above"
  exit 1
fi