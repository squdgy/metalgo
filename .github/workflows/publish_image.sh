#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# If this is not a trusted build (Docker Credentials are not set)
if [[ -z "$DOCKER_USERNAME"  ]]; then
  exit 0;
fi

# Metal root directory
METAL_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd ../.. && pwd )

# Load the constants
source "$METAL_PATH"/scripts/constants.sh

if [[ $current_branch == "master" ]]; then
  echo "Tagging current avalanchego image as $metalgo_dockerhub_repo:latest"
  docker tag $metalgo_dockerhub_repo:$current_branch $metalgo_dockerhub_repo:latest
fi

echo "Pushing: $metalgo_dockerhub_repo:$current_branch"

echo "$DOCKER_PASS" | docker login --username "$DOCKER_USERNAME" --password-stdin

## pushing image with tags
docker image push -a $metalgo_dockerhub_repo
