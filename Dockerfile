# Stage 1: Build the Go application
FROM golang:1.23.1-alpine as builder
WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /app/bin/app ./cmd/auth/main.go

# Stage 2: Create a smaller image for running the binary
FROM alpine:latest

COPY --from=builder /app/bin/app /

ENTRYPOINT ["/app"] 