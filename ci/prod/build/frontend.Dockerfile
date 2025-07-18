FROM oven/bun:1.2-slim

WORKDIR /app

ENV DEBIAN_FRONTEND=noninteractive

RUN --mount=type=cache,target=/var/cache/apt \
    --mount=type=cache,target=/var/lib/apt \
    apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates

COPY package.json bun.lock ./

RUN --mount=type=cache,target=/root/.bun/install/cache \
    bun install --frozen-lockfile

ENV NODE_ENV=production

ARG FRONTEND_PUBLIC_API_URL
ENV FRONTEND_PUBLIC_API_URL=$FRONTEND_PUBLIC_API_URL

COPY . /app

ARG FRONTEND_PUBLIC_SENTRY_RELEASE
ENV FRONTEND_PUBLIC_SENTRY_RELEASE=$FRONTEND_PUBLIC_SENTRY_RELEASE

RUN --mount=type=cache,target=/root/.bun/install/cache \
    --mount=type=secret,id=FRONTEND_PUBLIC_SENTRY_DSN,env=FRONTEND_PUBLIC_SENTRY_DSN \
    --mount=type=secret,id=SENTRY_AUTH_TOKEN,env=SENTRY_AUTH_TOKEN \
    bun run build && \
    find ./dist/assets/*.map -exec rm {} \;
