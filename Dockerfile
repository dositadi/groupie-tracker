FROM golang:tip-alpine3.23

WORKDIR /groupie-tracker

RUN go install github.com/air-verse/air@latest

COPY go.mod /groupie-tracker/

RUN go mod download

CMD [ "air" ]