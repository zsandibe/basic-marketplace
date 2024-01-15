FROM golang:latest

WORKDIR /go/src/basic-marketplace/services/auth

ADD . /go/src/basic-marketplace

RUN go mod vendor

CMD go run cmd/main.go

EXPOSE 8082
