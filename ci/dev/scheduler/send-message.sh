#!/usr/bin/env sh

event_type="$1"

MESSAGE_BODY=$(cat <<EOF
{
  "eventType": "$event_type",
  "payload": {}
}
EOF
)

curl -s -o /dev/null -X POST "http://sqs:9324/queue/polimane-scheduled" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "Action=SendMessage&MessageBody=$(printf '%s' "$MESSAGE_BODY" | jq -sRr @uri)"
