# syntax=docker/dockerfile:1

FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download
COPY . .

RUN go build .

EXPOSE 25565
CMD ["./spotify-motd"]