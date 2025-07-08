FROM debian:bullseye

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends curl build-essential ca-certificates && \
    rm -rf /var/lib/apt/lists/*

RUN curl -sSfL https://atlasgo.sh | sh

COPY Makefile ./
COPY atlas.hcl ./
COPY migrations ./migrations

RUN --mount=type=secret,id=BACKEND_DATABASE_URL,env=BACKEND_DATABASE_URL \
    --mount=type=secret,id=BACKEND_DATABASE_CERT,target=/tmp/postgres/ca-cert.pem \
    make db_migrate env=prod
