FROM golang:1.12 AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

EXPOSE 8080

ENTRYPOINT ["./main"]