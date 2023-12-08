all: build dev

build:
	@go build -o ./bin ./cmd/webapp-api

dev:
	@./bin/webapp-api

.PHONY = all build dev