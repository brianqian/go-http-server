all: build dev

build:
	@go build -o ./bin/api ./cmd/main.go

dev:
	@./bin/api

.PHONY = all build dev