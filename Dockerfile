FROM golang:alpine3.22

WORKDIR /groupie-tracker

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

CMD [ "air" ]