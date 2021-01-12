FROM golang:1.15

ARG GO_SRC

WORKDIR /go/src/app
COPY ${GO_SRC} .

ENTRYPOINT go run serverentry.go
