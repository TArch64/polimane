#!/usr/bin/env bash

set -e

input=$(cat)
BUILD_ID=$(echo "$input" | jq -r '.build_id')
BUILD_IMAGE=$(echo "$input" | jq -r '.build_image')
BUILD_DOCKERFILE=$(echo "$input" | jq -r '.build_dockerfile')
BUILD_CONTEXT=$(echo "$input" | jq -r '.build_context')
BUILD_DIST=$(echo "$input" | jq -r '.build_dist')
BUILD_ARGS=$(echo "$input" | jq -r '.build_args // "" | split("|") | map("--build-arg \"" + . + "\"") | join(" ")')

cleanup() {
  docker container rm -f $BUILD_ID >&2 || true
}

trap cleanup EXIT INT TERM

docker build -t $BUILD_IMAGE -f $BUILD_DOCKERFILE $BUILD_ARGS $BUILD_CONTEXT >&2
docker run -d --name $BUILD_ID $BUILD_IMAGE tail -f /dev/null >&2
rm -rf $BUILD_DIST >&2
docker cp $BUILD_ID:/app/dist $BUILD_DIST >&2
docker container rm -f $BUILD_ID >&2

echo '{}'
