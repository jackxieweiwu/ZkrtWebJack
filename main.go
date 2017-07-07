package main

import (
	_ "ZkrtWebJack/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"ZkrtWebJack/models"
	"log"
	"flag"
	_"ZkrtWebJack/controllers"
	"ZkrtWebJack/config"
	"ZkrtWebJack/rtmp"
)


func init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	flag.Parse()
	models.RegisterDataBase()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//orm.Debug = true
	go beego.Run()


	var err error
	/*if err = InitAppConfig(); err != nil {
		return
	}*/
	l := ":1935"
	err = rtmp.ListenAndServe(l)
	if err != nil {
		panic(err)
	}
	select {}
}

func InitAppConfig() (err error) {
	cfg := new(config.Config)
	err = cfg.Init("app.conf")
	if err != nil {
		return
	}
	return
}