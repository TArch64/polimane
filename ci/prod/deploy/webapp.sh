#!/usr/bin/env bash

set -e

input=$(cat)
export CLOUDFLARE_ACCOUNT_ID=$(echo "$input" | jq -r '.cloudflare_account_id')
export CLOUDFLARE_API_TOKEN=$(echo "$input" | jq -r '.cloudflare_api_token')
BUILD_DIST=$(echo "$input" | jq -r '.build_dist')
PROJECT_NAME=$(echo "$input" | jq -r '.project_name')

npx -y wrangler pages deploy $BUILD_DIST --project-name $PROJECT_NAME >&2
echo '{}'
