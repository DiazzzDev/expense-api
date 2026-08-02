run:
	go run ./cmd/server

test:
	go test ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run

tidy:
	go mod tidy

build:
	go build -o bin/server ./cmd/server