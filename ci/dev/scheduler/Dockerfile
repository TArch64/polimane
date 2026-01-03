FROM debian:13.2-slim

WORKDIR /app

RUN --mount=type=cache,target=/var/cache/apt \
    --mount=type=cache,target=/var/lib/apt \
    apt-get update && \
    apt-get install -y --no-install-recommends cron jq curl

RUN touch /var/log/cron.log

COPY ./bin ./bin
ENV PATH="/app/bin:$PATH"

COPY  ./start.sh ./send-message.sh ./
COPY ./schedule /etc/cron.d/
RUN crontab /etc/cron.d/schedule

CMD ./start.sh
