package controllers

import (
	"log"
	"github.com/gorilla/websocket"
	"ZkrtWebJack/models"
	"fmt"
)

var(
	clients = make(map[*websocket.Conn]bool)
	clients2 = make(map[*websocket.Conn]bool)
	broadcast = make(chan models.Message)
	moudleSta = make(chan models.MessageStatus)
)

func init()  {
	go handleMessages()
	go handleMessagesMoudleStatus()
}

func handleMessages()  {
	for{
		msg := <-broadcast
		//fmt.Println("clients1 len",len(clients))
		for client := range clients{
			err:=client.WriteJSON(msg)
			if err != nil{
				log.Printf("client,WriteJSON error:%v",err)
				client.Close()
				delete(clients,client)
			}
		}
	}
}


func handleMessagesMoudleStatus()  {
	for{
		msgSta := <-moudleSta
		fmt.Println("clients2 len",len(clients2))
		for client := range clients2{
			err:=client.WriteJSON(msgSta)
			if err != nil{
				log.Printf("client,WriteJSON error:%v",err)
				client.Close()
				delete(clients2,client)
			}
		}
	}
}
