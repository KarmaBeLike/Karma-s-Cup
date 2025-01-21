# Базовый образ Golang
FROM golang:1.22

# Установка рабочего каталога
WORKDIR /app

# Копирование файлов проекта в контейнер
COPY . .

# Установка зависимостей
RUN go mod tidy

# Сборка приложения
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o myapp main.go

# Запуск приложения
CMD ["./myapp"]
