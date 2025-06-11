FROM oven/bun:1.2-slim

WORKDIR /app

COPY package.json bun.lock ./
RUN bun install --frozen-lockfile

COPY . /app

ARG API_URL
ENV FRONTEND_PUBLIC_API_URL=$API_URL
ENV NODE_ENV=production

RUN bun run build-only
