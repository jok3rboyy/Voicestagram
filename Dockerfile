FROM golang:1-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY . .

RUN go build -o app

EXPOSE 8080

CMD ["./app"]
