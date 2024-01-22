FROM golang:1.19.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app

ENV ENV_VARIABLE=value

EXPOSE 8080

CMD ["./app"]