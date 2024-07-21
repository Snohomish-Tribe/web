FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

EXPOSE 3000

