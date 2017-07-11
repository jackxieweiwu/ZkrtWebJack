FROM golang:1.8

MAINTAINER "xieweiwu@zkrtuav.com"

WORKDIR $GOPATH/src/ZkrtWebJack
ADD . $GOPATH/src/ZkrtWebJack

RUN go get github.com/astaxie/beego && go get github.com/beego/bee
RUN go get github.com/astaxie/beego/context/param
RUN go get github.com/astaxie/beego/orm
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/websocket
RUN go get github.com/satori/go.uuid

RUN go build .

EXPOSE 8080
EXPOSE 1935

ENTRYPOINT ["./zkrtwebjack"]