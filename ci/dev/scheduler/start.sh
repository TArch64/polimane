#!/usr/bin/env sh

set -e

cron
echo "Starting cron scheduler..."
cat /etc/cron.d/schedule
tail -f /var/log/cron.log
