.PHONY: dev prod

dev:
	gow -e=go,mod run -tags "dev" main_dev.go

prod:
	@echo "Building production binary... $(out_file)"
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags "prod" -o "$(out_file)" main_prod.go
	zip -j $(out_file).zip $(out_file)
