FROM golang:1.25-bookworm AS base

WORKDIR /app

ENV DEBIAN_FRONTEND=noninteractive

RUN --mount=type=cache,target=/var/cache/apt \
    --mount=type=cache,target=/var/lib/apt \
    apt-get update && \
    apt-get install -y --no-install-recommends zip musl-tools build-essential

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . /app


FROM base AS api

RUN --mount=type=cache,target=/go/pkg/mod \
    make prod_api out_dir=/app/dist


FROM base AS worker

RUN --mount=type=cache,target=/go/pkg/mod \
    make prod_worker out_dir=/app/dist
