FROM golang:1.12

WORKDIR /go/src/github.com/AndreBtt/Computsal

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

COPY . .

WORKDIR /go/src/github.com/AndreBtt/Computsal/server

RUN go build -o app

CMD ./app