# Build stage
FROM golang:1.22.6-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum* .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bot -a -installsuffix cgo .

# Final stage
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарный файл
COPY --from=builder /app/bot .

# Копируем .env файл
COPY .env .

CMD ["./bot"]