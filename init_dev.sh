#!/usr/bin/env sh

set -e

echo "Building docker images..."
docker compose build scheduler # build first to omit a conflict in apk cache
docker compose build


echo "Init S3 storage..."
minio_init=$(cat <<'EOF'
mc alias set myminio http://localhost:9000 $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD
mc mb myminio/polimane-dev
mc anonymous set download myminio/polimane-dev
EOF
)

docker compose up -d s3
sleep 5 # wait for minio to be ready
docker compose exec s3 sh -c "$minio_init"
docker compose down s3

echo "Installing Front-end Dependencies..."
docker compose run --rm frontend bun install

echo "Installing Back-end Dependencies and Running DB Migrations..."
docker compose run --rm backend sh -c 'go mod download && make db_migrate'

docker compose down
