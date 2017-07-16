FROM golang:1.8

MAINTAINER china "xieweiwu@zkrtuav.com"

RUN go get github.com/astaxie/beego && go get github.com/beego/bee
RUN go get github.com/astaxie/beego/context/param
RUN go get github.com/astaxie/beego/orm
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/websocket
RUN go get github.com/satori/go.uuid

EXPOSE 8080
