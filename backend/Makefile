.PHONY: dev db_migration_new db_migration_blank db_rehash_migrations db_migrate prod

ldflags = '-linkmode external -extldflags "-static -Wl,-unresolved-symbols=ignore-all"'
env ?= dev

dev:
	CGO_ENABLED=1 \
	CC="musl-gcc" \
	gow -e=go,mod run -tags "dev" -ldflags $(ldflags) main.go

db_new_migration:
	atlas migrate diff $(name) --env $(env)

db_blank_migration:
	atlas migrate new $(name) --env $(env)

db_rehash_migrations:
	atlas migrate hash --env $(env)

db_migrate:
	@make execute_with_retry cmd="atlas migrate apply --env $(env)" on_error="unexpected transaction status idle" interval=2

execute_with_retry:
	@while ! $(cmd) 2>&1 | grep -v "$(on_error)" > /dev/null; do \
		echo "Retrying..."; \
		sleep $(interval); \
	done

prod:
	GOOS=linux \
	GOARCH=arm64 \
	CGO_ENABLED=1 \
	CC="musl-gcc" \
	go build -tags "prod" -o "$(out_dir)/bootstrap" -ldflags $(ldflags) main.go
	zip -j "$(out_dir)/bootstrap.zip" "$(out_dir)/bootstrap"
	rm "$(out_dir)/bootstrap"
