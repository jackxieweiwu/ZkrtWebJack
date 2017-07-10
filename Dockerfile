FROM golang:1.8

MAINTAINER "xieweiwu@zkrtuav.com"

# 下载并安装第三方依赖到容器中
RUN go get github.com/astaxie/beego
RUN go get github.com/astaxie/beego/context/param
RUN go get github.com/astaxie/beego/orm
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/websocket
RUN go get github.com/satori/go.uuid

WORKDIR $GOPATH/src/ZkrtWebJack
ADD . $GOPATH/src/ZkrtWebJack
RUN go build .

EXPOSE 8080

ENTRYPOINT ["./zkrtdocker"]
