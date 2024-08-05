FROM golang:1.22-alpine AS base

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo
ENV GOPATH=/go

WORKDIR /app

COPY packages/server $ROOTPATH

RUN go mod download

EXPOSE 8080

FROM base AS dev
RUN apk add --no-cache git curl bash
RUN go install github.com/air-verse/air@v1.52.2 \
    && go install ariga.io/atlas/cmd/atlas@latest \
    && go install github.com/google/wire/cmd/wire@latest
ENTRYPOINT ["air"]

FROM base AS prd
RUN go install ariga.io/atlas/cmd/atlas@latest
RUN go build -o ./build/main
ENTRYPOINT ["./build/main"]
