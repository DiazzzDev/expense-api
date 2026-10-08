migrate-up:
	goose -dir migrations postgres "$$DATABASE_URL" up

migrate-down:
	goose -dir migrations postgres "$$DATABASE_URL" down

migrate-status:
	goose -dir migrations postgres "$$DATABASE_URL" status

run:
	go run ./cmd/api