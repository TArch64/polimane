FROM golang:1.25-bookworm

WORKDIR /app

ENV PATH="/app/bin:$PATH"
ENV DEBIAN_FRONTEND=noninteractive

RUN --mount=type=cache,target=/var/cache/apt \
    --mount=type=cache,target=/var/lib/apt \
    apt-get update && \
    apt-get install -y --no-install-recommends musl-tools build-essential

RUN --mount=type=cache,target=/go/pkg/mod \
    go install github.com/mitranim/gow@latest

RUN curl -sSf https://atlasgo.sh | sh
