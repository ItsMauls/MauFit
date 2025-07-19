# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/server/main.go

# Stage 2: Install Air for hot-reloading
FROM golang:1.24-alpine AS air_installer
RUN go install github.com/air-verse/air@latest

# Stage 3: Final image
FROM golang:1.24-alpine

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=air_installer /go/bin/air /usr/local/bin/
COPY . .

EXPOSE 8081

CMD ["air", "-c", ".air.toml"]