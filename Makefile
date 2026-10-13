.PHONY: run dev build fmt vet lint tidy test psql migrate-up migrate-down migrate-status migrate-reset migrate-create

MIGRATIONS_DIR := migrations

run:
	go run ./cmd/api

dev:
	air

build:
	go build -o bin/api ./cmd/api

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

tidy:
	go mod tidy

test:
	go test ./...

psql:
	psql "$$DATABASE_URL"

migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$$DATABASE_URL" up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$$DATABASE_URL" down

migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres "$$DATABASE_URL" status

# Solo para la DB local de practica: revierte todo y vuelve a aplicar
migrate-reset:
	goose -dir $(MIGRATIONS_DIR) postgres "$$DATABASE_URL" reset
	goose -dir $(MIGRATIONS_DIR) postgres "$$DATABASE_URL" up

migrate-create:
	@test -n "$(name)" || (echo "uso: make migrate-create name=add_algo" && exit 1)
	goose -dir $(MIGRATIONS_DIR) -s create $(name) sql
