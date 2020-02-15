FROM golang:latest

LABEL maintainer="rostamiarmin@gmail.com"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 9001

CMD ./main
