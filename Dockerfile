FROM golang:1.22.1

WORKDIR /app
# WORKDIR /cmd/web

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 3000

CMD ["/docker-gs-ping"]