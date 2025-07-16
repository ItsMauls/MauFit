# --- STAGE 1: Build ---
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/server/main.go

# --- STAGE 2: Final Image ---
FROM alpine:latest

WORKDIR /app

# Hanya salin binary aplikasi yang sudah jadi
COPY --from=builder /app/main .

EXPOSE 8080

# Langsung jalankan aplikasi Go
CMD ["/app/main"]
