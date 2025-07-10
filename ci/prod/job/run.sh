#!/usr/bin/env bash

set -e

BUILD_ARGS=$(echo "$BUILD_ARGS" | jq -r '. // [] | map("--build-arg " + .) | join(" ")')
BUILD_SECRET=$(echo "$BUILD_SECRET" | jq -r '. // [] | map("--secret id=" + .) | join(" ")')

docker build -t $BUILD_IMAGE -f $BUILD_DOCKERFILE $BUILD_ARGS $BUILD_SECRET $BUILD_CONTEXT
