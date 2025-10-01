.PHONY: dev_api dev_worker db_migration_new db_migration_blank db_rehash_migrations db_migrate prod_api prod_worker test

MAKEFLAGS += --no-print-directory
ldflags = '-linkmode external -extldflags "-static -Wl,-unresolved-symbols=ignore-all"'
env ?= dev
test_pattern ?= ./...

cgo_envs:
	go env -w CGO_ENABLED=1
	go env -w CC=musl-gcc

dev_api: cgo_envs
	gow -e=go,mod,tmpl run -tags "dev" -ldflags $(ldflags) main_api.go

dev_worker: cgo_envs
	gow -e=go,mod,tmpl run -tags "dev" -ldflags $(ldflags) main_worker.go

db_new_migration:
	atlas migrate diff $(name) --env $(env)

db_blank_migration:
	atlas migrate new $(name) --env $(env)

db_rehash_migrations:
	atlas migrate hash

db_migrate:
	@make execute_with_retry cmd="atlas migrate apply --env $(env)" on_error="unexpected transaction status idle" interval=2

execute_with_retry:
	@while ! $(cmd) 2>&1 | grep -v "$(on_error)"; do \
		echo "Retrying..."; \
		sleep $(interval); \
	done

test: cgo_envs
	go test -tags "test" -ldflags $(ldflags) $(test_pattern)

test_coverage: cgo_envs
	go test -tags "test" -coverprofile=coverage/result.out -covermode=atomic -ldflags $(ldflags) $(test_pattern)
	go tool cover -html=coverage/result.out -o coverage/result.html
	go tool cover -func=coverage/result.out

prod_api: cgo_envs
	GOOS=linux GOARCH=arm64 \
	go build -tags "prod" -o "$(out_dir)/bootstrap" -ldflags $(ldflags) main_api.go
	zip -j "$(out_dir)/bootstrap.zip" "$(out_dir)/bootstrap"
	rm "$(out_dir)/bootstrap"

prod_worker: cgo_envs
	GOOS=linux GOARCH=arm64 \
	go build -tags "prod" -o "$(out_dir)/bootstrap" -ldflags $(ldflags) main_worker.go
	zip -j "$(out_dir)/bootstrap.zip" "$(out_dir)/bootstrap"
	rm "$(out_dir)/bootstrap"
