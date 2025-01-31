FROM golang:1.22.6-alpine

# Устанавливаем tzdata и устанавливаем часовой пояс
RUN apk add --no-cache tzdata
ENV TZ=Asia/Atyrau

WORKDIR /app

COPY go.mod .
COPY go.sum* .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bot -a -installsuffix cgo .

CMD ["./bot"]