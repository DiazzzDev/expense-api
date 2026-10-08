FROM mcr.microsoft.com/devcontainers/go:1.26

RUN apt-get update && \
    apt-get install -y \
        git \
        curl \
        wget \
        make \
        unzip \
        postgresql-client \
    && rm -rf /var/lib/apt/lists/*

RUN go install github.com/air-verse/air@latest

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

ENV PATH="$PATH:/go/bin"