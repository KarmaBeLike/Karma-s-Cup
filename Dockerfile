# Базовый образ
FROM golang:1.22 as builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Сборка приложения
RUN go build -o bot .

# Минимальный финальный образ
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/bot .

# Копируем .env файл
COPY .env .

# Запуск бота
CMD ["./bot"]
