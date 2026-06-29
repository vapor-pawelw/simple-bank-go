include .env
export

GOOSE_DRIVER ?= postgres
GOOSE_DBSTRING ?= $(DATABASE_URL_LOCAL)
GOOSE_MIGRATION_DIR ?= ./internal/db/migrations

install_tools:
	go mod download
migrate_up:
	go tool goose up
migrate_down:
	go tool goose down
sqlc:
	go tool sqlc generate -f internal/db/sqlc.yaml
