all: build dev

build:
	@go build -o ./bin/api ./cmd/main.go

dev:
	@./bin/api

migration:
ifdef name
	@TERN_CONFIG="./data/migrations/tern.dev.conf" TERN_MIGRATIONS="./data/migrations" tern new $(name)
	@echo "Creating migration with name $(name)..."
else
	@TERN_CONFIG="./data/migrations/tern.dev.conf" TERN_MIGRATIONS="./data/migrations" tern new temp_name
	@echo "Creating migration with name 'temp_name'..."
endif

migrate:
	@TERN_CONFIG="./data/migrations/tern.dev.conf" TERN_MIGRATIONS="./data/migrations" tern migrate

seed:build
	@./bin/api seed

.PHONY = all build dev migration