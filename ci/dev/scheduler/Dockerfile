FROM ubuntu:24.04

WORKDIR /app

RUN --mount=type=cache,target=/var/cache/apt \
    --mount=type=cache,target=/var/lib/apt \
    apt-get update && \
    apt-get install -y --no-install-recommends cron jq curl

RUN touch /var/log/cron.log
COPY ./start.sh ./

COPY ./schedule /etc/cron.d/
COPY ./send-message.sh ./
RUN crontab /etc/cron.d/schedule

CMD ./start.sh
