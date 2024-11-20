## Image optimized with multi-stage build

# Stage 1: golang alpine
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum to install dependencies and run
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Optimization and pointing to ./cmd where main.go is
RUN CGO_ENABLED=0 GOOS=linux go build -o golang-products-api ./cmd

# Stage 2: using scratch to final image, compiling binary
FROM scratch

COPY --from=builder /app/golang-products-api /

ENTRYPOINT [ "/golang-products-api" ]