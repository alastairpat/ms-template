#!/usr/bin/env bash

function validate_jq_installed {
  if ! command -v jq > /dev/null
  then
    echo "This script relies on jq to function. Install it using brew (macOS), apt (Debian), or yum/dnf (Fedora)"
    exit 1
  fi
}

function get_application_version {
  if ! TARGET_VERSION=$(jq --raw-output '.version' ./metadata.json)
  then
    echo "Failed to determine application version. Make sure metadata.json exists in the current directory"
    exit 1
  fi

  export TARGET_VERSION=$TARGET_VERSION
}

function get_application_name {
  if ! APPLICATION_NAME=$(jq --raw-output '.name' ./metadata.json)
  then
    echo "Failed to determine application name. Make sure metadata.json exists in the current directory"
    exit 1
  fi

  export APPLICATION_NAME=$APPLICATION_NAME
}