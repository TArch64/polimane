#!/usr/bin/env bash

set -e

BUILD_ARGS=$(echo "$BUILD_ARGS" | jq -r '. // [] | map("--build-arg " + .) | join(" ")')

cleanup() {
  docker container rm -f $BUILD_ID || true
}

trap cleanup EXIT INT TERM

docker build -t $BUILD_IMAGE -f $BUILD_DOCKERFILE $BUILD_ARGS $BUILD_CONTEXT
docker run -d --name $BUILD_ID $BUILD_IMAGE tail -f /dev/null
rm -rf $BUILD_DIST
docker cp $BUILD_ID:/app/dist $BUILD_DIST
