FROM golang:1.25-bullseye

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

RUN --mount=type=cache,target=/go/pkg/mod \
    make prod out_dir=/app/dist
