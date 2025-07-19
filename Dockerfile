# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/auth

# Air binary for dev
RUN go install github.com/cosmtrek/air@latest

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /go/bin/air /usr/bin/air
# ถ้ามีไฟล์ air config
COPY .air.toml . 
EXPOSE 8080

# By default run the binary, can override CMD for dev
CMD ["./main"]