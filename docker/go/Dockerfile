FROM golang:1.22-alpine

RUN apk add --no-cache git curl bash

# GOPATHの設定
ENV GOPATH=/go

# ワーキングディレクトリの設定
WORKDIR /server

RUN go install github.com/air-verse/air@v1.52.2
RUN go install ariga.io/atlas/cmd/atlas@latest

USER root

EXPOSE 8080

CMD ["air"]
