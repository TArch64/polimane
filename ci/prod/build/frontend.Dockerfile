FROM oven/bun:1.2-slim

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    rm -rf /var/lib/apt/lists/*

COPY package.json bun.lock ./
RUN bun install --frozen-lockfile

COPY . /app

ARG FRONTEND_PUBLIC_API_URL
ENV FRONTEND_PUBLIC_API_URL=$FRONTEND_PUBLIC_API_URL

ARG FRONTEND_PUBLIC_SENTRY_RELEASE
ENV FRONTEND_PUBLIC_SENTRY_RELEASE=$FRONTEND_PUBLIC_SENTRY_RELEASE

ENV NODE_ENV=production

RUN --mount=type=secret,id=build_secret export $(cat /run/secrets/build_secret | xargs) && \
    bun run build && \
    find ./dist/assets/*.map -exec rm {} \;
