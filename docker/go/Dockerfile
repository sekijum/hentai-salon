FROM golang:1.22-alpine

RUN apk add --no-cache git curl bash

ENV GOPATH=/go

WORKDIR /server

RUN go install github.com/air-verse/air@v1.52.2
RUN go install ariga.io/atlas/cmd/atlas@latest
RUN go install github.com/google/wire/cmd/wire@latest

USER root

EXPOSE 8080

CMD ["air"]
