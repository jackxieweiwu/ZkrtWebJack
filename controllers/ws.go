package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"log"
	"time"
	"ZkrtWebJack/models"
	"encoding/json"
)

type MywebSockerControllers struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{}

//返回飞行器的基本状态
func (c*MywebSockerControllers) Get()  {
	ws,err := upgrader.Upgrade(c.Ctx.ResponseWriter,c.Ctx.Request,nil)
	if err != nil{
		log.Fatal(err)
	}
	clients[ws] = true
	for {
		time.Sleep(time.Second*1)
		b, _ :=json.Marshal(result_dronemsg)
		msg :=models.Message{string(b)}
		broadcast <-msg
	}
}

//返回模块的状态
func (c*MywebSockerControllers) GetMoudleStatus()  {
	ws,err := upgrader.Upgrade(c.Ctx.ResponseWriter,c.Ctx.Request,nil)
	if err != nil{
		log.Fatal(err)
	}
	clients2[ws] = true
	for {
		time.Sleep(time.Second*1)
		moudleS, _ :=json.Marshal(moudlsta)
		msgSta :=models.MessageStatus{string(moudleS)}
		moudleSta <-msgSta
	}
}
