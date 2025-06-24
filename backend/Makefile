.PHONY: dev prod

dev:
	gow -e=go,mod run -tags "dev" main_dev.go

prod:
	GOOS=linux \
	GOARCH=amd64 \
	CGO_ENABLED=1 \
	CC=clang \
	CXX=clang++ \
	go build -tags "prod" -o "$(out_dir)/bootstrap" main_prod.go
	zip -j "$(out_dir)/bootstrap.zip" "$(out_dir)/bootstrap"
	rm "$(out_dir)/bootstrap"
