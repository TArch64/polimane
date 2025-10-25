#!/usr/bin/env sh

set -e

cron
echo "Starting cron scheduler..."
cut -d'>' -f1 /etc/cron.d/schedule
tail -f /var/log/cron.log
