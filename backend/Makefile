.PHONY: dev prod

dev:
	gow -e=go,mod run -tags "dev" main_dev.go

prod:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags "prod" -o "$(out_dir)/bootstrap" main_prod.go
	zip -j "$(out_dir)/lambda.zip" "$(out_dir)/bootstrap"
