FROM golang:1.18

# Set working directory
WORKDIR /go/src/app

# Copy source code
COPY . .

# Expose port
EXPOSE 8000

# Build
RUN go build -o main cmd/main.go

# Run executable
CMD ["./main"]