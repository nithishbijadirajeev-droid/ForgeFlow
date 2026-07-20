# ---------- Build Stage ----------
FROM golang:1.26.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o forgeflow ./app/cmd/server

# ---------- Runtime Stage ----------
FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/forgeflow .

EXPOSE 8080

CMD ["./forgeflow"]