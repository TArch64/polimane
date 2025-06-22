FROM oven/bun:1.2-slim

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY package.json bun.lock ./
RUN bun install --frozen-lockfile

ENV NODE_ENV=production

ARG FRONTEND_PUBLIC_API_URL
ENV FRONTEND_PUBLIC_API_URL=$FRONTEND_PUBLIC_API_URL

COPY . /app

ARG FRONTEND_PUBLIC_SENTRY_RELEASE
ENV FRONTEND_PUBLIC_SENTRY_RELEASE=$FRONTEND_PUBLIC_SENTRY_RELEASE

RUN --mount=type=secret,id=FRONTEND_PUBLIC_SENTRY_DSN,env=FRONTEND_PUBLIC_SENTRY_DSN \
    --mount=type=secret,id=SENTRY_AUTH_TOKEN,env=SENTRY_AUTH_TOKEN \
    bun run build && \
    find ./dist/assets/*.map -exec rm {} \;
