FROM library/golang

# Godep for vendoring
RUN go get github.com/tools/godep

RUN go get github.com/astaxie/beego && go get github.com/beego/bee
RUN go get github.com/astaxie/beego/context/param
RUN go get github.com/astaxie/beego/orm
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/websocket
RUN go get github.com/satori/go.uuid

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/home/jack/goglandProjects/src/ZkrtWebJack
RUN mkdir -p $APP_DIR

# Set the entrypoint
ENTRYPOINT (cd $APP_DIR && ./ZkrtWebJack)
ADD . $APP_DIR

# Compile the binary and statically link
RUN cd $APP_DIR && CGO_ENABLED=0 godep go build -ldflags '-d -w -s'

EXPOSE 8080