FROM golang:1.8

MAINTAINER Razil "xieweiwu@zkrtuav.com"

WORKDIR $GOPATH/src/ZkrtWebJack
ADD . $GOPATH/src/ZkrtWebJack
RUN go build .

EXPOSE 8080

ENTRYPOINT ["./zkrtdocker"]
