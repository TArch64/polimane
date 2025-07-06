.PHONY: dev db_migration_new db_migration_blank db_migrate prod

ldflags = '-linkmode external -extldflags "-static -Wl,-unresolved-symbols=ignore-all"'

dev:
	CGO_ENABLED=1 \
	CC="musl-gcc" \
	gow -e=go,mod run -tags "dev" -ldflags $(ldflags) main_dev.go

db_migration_new:
	atlas migrate diff $(name) --env dev

db_migration_blank:
	atlas migrate new $(name) --env dev

db_migrate:
	atlas migrate apply --env dev --url $(BACKEND_DATABASE_URL)

prod:
	GOOS=linux \
	GOARCH=arm64 \
	CGO_ENABLED=1 \
	CC="musl-gcc" \
	go build -tags "prod" -o "$(out_dir)/bootstrap" -ldflags $(ldflags) main_prod.go
	zip -j "$(out_dir)/bootstrap.zip" "$(out_dir)/bootstrap"
	rm "$(out_dir)/bootstrap"
