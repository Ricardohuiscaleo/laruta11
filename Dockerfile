FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod ./
RUN go mod download && go mod tidy

# Copy source code
COPY *.go ./

# Build
RUN go build -o api .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary
COPY --from=builder /app/api .

# Expose port
EXPOSE 3000

# Run
CMD ["./api"]
