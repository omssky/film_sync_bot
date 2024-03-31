# syntax=docker/dockerfile:1
FROM golang:1.22-alpine AS builder

WORKDIR /build

# dependencies
COPY go.mod go.sum ./
RUN go mod download

# build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/bot ./cmd/film_sync_bot/main.go

FROM alpine AS runner

WORKDIR /

COPY --from=builder /build/bin/bot /bot
COPY config.yaml ./
CMD ["/bot"]
