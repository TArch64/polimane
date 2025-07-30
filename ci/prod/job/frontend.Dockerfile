FROM polimane-prod-frontend AS builder

FROM node:22-slim

WORKDIR /app

ENV DEBIAN_FRONTEND=noninteractive
ENV NODE_ENV=production
ENV PATH="/app/node_modules/.bin:$PATH"

ARG CLOUDFLARE_ACCOUNT_ID
ENV CLOUDFLARE_ACCOUNT_ID=$CLOUDFLARE_ACCOUNT_ID

ARG PROJECT_NAME
ENV PROJECT_NAME=$PROJECT_NAME

COPY --from=builder /app/package.json /app/bun.lock ./
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/dist ./dist

RUN --mount=type=secret,id=CLOUDFLARE_API_TOKEN,env=CLOUDFLARE_API_TOKEN npm run deploy
