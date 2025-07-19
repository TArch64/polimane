FROM debian:bullseye

WORKDIR /app

ENV DEBIAN_FRONTEND=noninteractive

RUN --mount=type=cache,target=/var/cache/apt \
    --mount=type=cache,target=/var/lib/apt \
    apt-get update && \
    apt-get install -y --no-install-recommends curl build-essential ca-certificates

RUN curl -sSfL https://atlasgo.sh | sh

COPY Makefile ./
COPY atlas.hcl ./
COPY migrations ./migrations

RUN --mount=type=secret,id=BACKEND_DATABASE_URL,env=BACKEND_DATABASE_URL \
    --mount=type=secret,id=BACKEND_DATABASE_CERT,target=/tmp/postgres/ca-cert.pem \
    make db_migrate env=prod
