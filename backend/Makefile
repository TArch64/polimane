.PHONY: dev prod

ldflags = '-linkmode external -extldflags "-static -Wl,-unresolved-symbols=ignore-all"'

dev:
	CGO_ENABLED=1 \
	CC="musl-gcc" \
	gow -e=go,mod run -tags "dev" -ldflags $(ldflags) main_dev.go

prod:
	GOOS=linux \
	GOARCH=arm64 \
	CGO_ENABLED=1 \
	CC="musl-gcc" \
	go build -tags "prod" -o "$(out_dir)/bootstrap" -ldflags $(ldflags) main_prod.go
	zip -j "$(out_dir)/bootstrap.zip" "$(out_dir)/bootstrap"
	rm "$(out_dir)/bootstrap"
