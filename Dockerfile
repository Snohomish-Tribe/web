FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./



EXPOSE 3000

CMD ["go", "run"]