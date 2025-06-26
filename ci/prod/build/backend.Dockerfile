FROM golang:1.24-bullseye

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends zip musl-tools build-essential && \
    rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY . /app
RUN make prod out_dir=/app/dist
