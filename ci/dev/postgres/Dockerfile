FROM postgres:17-alpine

ENV POSTGRES_HOST_AUTH_METHOD=trust
COPY initdb /docker-entrypoint-initdb.d
