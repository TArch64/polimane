.PHONY: dev_api dev_worker db_new_migration db_migrate prod_api prod_worker prod_build

export CGO_ENABLED = 1
export CC = musl-gcc
export GOOSE_DRIVER = postgres
export GOOSE_DBSTRING = $(BACKEND_DATABASE_URL)
export GOOSE_MIGRATION_DIR = ./migrations
export SENTRY_ORG = myself-zmf
export SENTRY_PROJECT = polimane-backend

MAKEFLAGS += --no-print-directory
ldflags = '-linkmode external -extldflags "-static -Wl,-unresolved-symbols=ignore-all"'

dev_api:
	gow -e=go,mod,tmpl run -tags "dev" -ldflags $(ldflags) main_api.go

dev_worker:
	gow -e=go,mod,tmpl run -tags "dev" -ldflags $(ldflags) main_worker.go

db_new_migration:
	goose -s create $(name) sql

db_migrate:
	goose up

db_rollback:
	goose down

db_status:
	goose status

prod_api:
	sentry-cli releases new $(SENTRY_RELEASE)
	sentry-cli releases set-commits $(SENTRY_RELEASE) --commit "TArch64/polimane@$(SENTRY_COMMIT_SHA)"
	@make prod_build app=api out_dir=$(out_dir)
	sentry-cli deploys new --release $(SENTRY_RELEASE) --env production
	sentry-cli releases finalize $(SENTRY_RELEASE)

prod_worker:
	@make prod_build app=worker out_dir=$(out_dir)

prod_build:
	GOOS=linux GOARCH=arm64 \
	go build -tags "prod" -o "$(out_dir)/bootstrap" -ldflags $(ldflags) main_$(app).go
	zip -j "$(out_dir)/bootstrap.zip" "$(out_dir)/bootstrap"
	rm "$(out_dir)/bootstrap"
