# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates openssl

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem \
    -days 365 -nodes -subj "/CN=localhost"

RUN go build -ldflags="-s -w"  -o fakenews

# Final stage
FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/fakenews .
COPY --from=builder /app/*.pem ./
COPY --from=builder /app/templates ./templates

EXPOSE 8080
EXPOSE 8443
CMD ["./fakenews"]