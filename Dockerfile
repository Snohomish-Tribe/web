FROM golang:1.22.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copies all files
COPY . ./

# ./ build here
RUN go build -o ./snohomishtribe ./cmd/web

EXPOSE 3000

CMD ["./snohomishtribe"]
