FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd

FROM gcr.io/distroless/static

COPY --from=builder /app/server /server

EXPOSE 8080

ENTRYPOINT ["/server"]