# ========================
# Build stage
# ========================
FROM golang:1.25.6 AS builder

WORKDIR /app

# Copy go mod terlebih dahulu (cache dependency)
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o bin/server ./cmd/server

# ========================
# Runtime stage
# ========================
FROM alpine:latest

WORKDIR /app

# Copy binary hasil build
COPY --from=builder /app/bin/server ./server

# Expose port (sesuaikan dengan app kamu)
EXPOSE 8080

# Jalankan binary
CMD ["./server"]
