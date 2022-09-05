FROM golang:1.19

COPY . /go/src/app

WORKDIR /go/src/app