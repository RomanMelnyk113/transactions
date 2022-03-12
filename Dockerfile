# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR  /app

COPY go.mod go.sum /app/

RUN go mod download

COPY . /app

RUN go build -o /app/juni

CMD ["/app/juni"]
