FROM golang:1.22.6-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum* .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bot -a -installsuffix cgo .

CMD ["./bot"]