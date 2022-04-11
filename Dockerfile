#FROM golang:1.16.3-alpine
#
#RUN apk add --update &&  apk add git
#
#RUN mkdir /go/src/app
## ワーキングディレクトリの設定
#WORKDIR /go/src/app
## ホストのファイルをコンテナの作業ディレクトリに移行
#ADD . /go/src/app

FROM golang:1.16.3-alpine AS build

ENV GO111MODULE=on

WORKDIR /

COPY . /go/src/github.com/Snails8/go-api

RUN apk update && apk add --no-cache git
RUN cd /go/src/github.com/Snails8/go-api/api && go build -o bin/sample main.go

FROM alpine:3.8

COPY --from=build /go/src/github.com/Snails8/go-api/api/bin/sample /usr/local/bin/

CMD ["sample"]