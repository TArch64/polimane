FROM golang:1.24-bullseye

WORKDIR /app

ENV PATH="/app/bin:$PATH"

RUN apt-get update && \
    apt-get install -y --no-install-recommends musl-tools build-essential && \
    rm -rf /var/lib/apt/lists/*

RUN go install github.com/mitranim/gow@latest
RUN curl -sSf https://atlasgo.sh | sh
