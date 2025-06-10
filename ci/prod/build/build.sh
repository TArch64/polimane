#!/usr/bin/env bash

set -e

input=$(cat)
BUILD_ID=$(echo "$input" | jq -r '.build_id')
BUILD_IMAGE=$(echo "$input" | jq -r '.build_image')
BUILD_DOCKERFILE=$(echo "$input" | jq -r '.build_dockerfile')
BUILD_CONTEXT=$(echo "$input" | jq -r '.build_context')
BUILD_DIST=$(echo "$input" | jq -r '.build_dist')

cleanup() {
  docker container rm -f $BUILD_ID > /dev/null 2>&1 || true
}

trap cleanup EXIT INT TERM

docker build -t $BUILD_IMAGE -f $BUILD_DOCKERFILE $BUILD_CONTEXT > /dev/null
docker run -d --name $BUILD_ID $BUILD_IMAGE tail -f /dev/null > /dev/null
rm -rf $BUILD_DIST > /dev/null
docker cp $BUILD_ID:/app/dist $BUILD_DIST > /dev/null
docker container rm -f $BUILD_ID > /dev/null

echo '{}'
