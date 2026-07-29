FROM golang:1.26.2 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o app ./cmd

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/app .
COPY migrations ./migrations
CMD ["./app"]