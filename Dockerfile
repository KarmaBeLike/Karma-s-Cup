# Базовый образ
FROM golang:1.22-alpine


# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Сборка приложения
RUN go build -o bot .


# Запуск бота
CMD ["./bot"]
