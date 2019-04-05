FROM golang:1.12

WORKDIR /go/src/github.com/AndreBtt/Computsal/api

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

COPY . .

WORKDIR /go/src/github.com/AndreBtt/Computsal/api/config

ENTRYPOINT go run *.go

EXPOSE 8080