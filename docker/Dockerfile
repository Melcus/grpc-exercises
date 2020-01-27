FROM golang:1.13.6-alpine as builder

ENV GO111MODULE=on

WORKDIR /app

# Caching Go Modules
COPY /src/go.mod /app/go.mod
COPY /src/go.sum /app/go.sum

RUN go mod download

COPY /src /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o binary