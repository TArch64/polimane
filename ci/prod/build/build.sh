#!/usr/bin/env bash

set -e

BUILD_ARGS=$(echo "$BUILD_ARGS" | jq -r '. // [] | map("--build-arg " + .) | join(" ")')
BUILD_SECRET=$(echo "$BUILD_SECRET" | jq -r '. // [] | map("--secret id=" + .) | join(" ")')

cleanup() {
  docker container rm -f "build-$BUILD_ID" || true
}

trap cleanup EXIT INT TERM

docker build -t $BUILD_IMAGE -f $BUILD_DOCKERFILE $BUILD_ARGS $BUILD_SECRET $BUILD_CONTEXT
docker run -d --name "build-$BUILD_ID" $BUILD_IMAGE tail -f /dev/null
rm -rf $BUILD_DIST
docker cp "build-$BUILD_ID":/app/dist $BUILD_DIST
docker container rm -f "build-$BUILD_ID"
