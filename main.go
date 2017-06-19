package main

import (
	_ "ZkrtWebJack/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"ZkrtWebJack/models"
	//"github.com/astaxie/beego/orm"
	"log"
	"flag"
	"ZkrtWebJack/protocol/rtmp"
	//"time"
	_"ZkrtWebJack/controllers"
	"time"
	"net"
	"ZkrtWebJack/protocol/httpflv"
	"ZkrtWebJack/protocol/httpopera"
	"ZkrtWebJack/protocol/hls"
)


func init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	flag.Parse()
	models.RegisterDataBase();
}


var (
	version     = "master"
	rtmpAddr    = flag.String("rtmp-addr", ":1935", "RTMP server listen address")
	httpFlvAddr = flag.String("httpflv-addr", ":7001", "HTTP-FLV server listen address")
	hlsAddr     = flag.String("hls-addr", ":7002", "HLS server listen address")
	operaAddr   = flag.String("manage-addr", ":8080", "HTTP manage interface server listen address")
)

func startHTTPLiveServer() *hls.Server {
	hlsListen, err := net.Listen("tcp", *hlsAddr)
	if err != nil {
		log.Fatal(err)
	}

	hlsServer := hls.NewServer()
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("HLS server panic: ", r)
			}
		}()
		log.Println("HLS listen On", *hlsAddr)
		hlsServer.Serve(hlsListen)
	}()
	return hlsServer
}

//启动RTMP服务
func startRtmp(stream *rtmp.RtmpStream, hlsServer *hls.Server) {
	rtmpListen, err := net.Listen("tcp", *rtmpAddr)
	if err != nil {
		log.Fatal(err)
	}

	rtmpServer := rtmp.NewRtmpServer(stream, hlsServer)
	defer func() {
		if r := recover(); r != nil {
			log.Println("RTMP server panic: ", r)
		}
	}()
	log.Println("RTMP Listen On", *rtmpAddr)
	rtmpServer.Serve(rtmpListen)
}

//启动Http_FLV服务
func startHTTPFlvVideoServer(stream *rtmp.RtmpStream) {
	flvListen, err := net.Listen("tcp", *httpFlvAddr)
	if err != nil {
		log.Fatal(err)
	}

	hdlServer := httpflv.NewServer(stream)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("HTTP-FLV server panic: ", r)
			}
		}()
		log.Println("HTTP-FLV listen On", *httpFlvAddr)
		hdlServer.Serve(flvListen)
	}()
}

func startHTTPLiveRtmpOpera(stream *rtmp.RtmpStream) {
	if *operaAddr != "" {
		opListen, err := net.Listen("tcp", *operaAddr)
		if err != nil {
			log.Fatal(err)
		}
		opServer := httpopera.NewServer(stream)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Println("HTTP-Operation server panic: ", r)
				}
			}()
			log.Println("HTTP-Operation listen On", *operaAddr)
			opServer.Serve(opListen)
		}()
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//orm.Debug = true
	go beego.Run()

	defer func() {
		if r := recover(); r != nil {
			log.Println("livego panic: ", r)
			time.Sleep(1 * time.Second)

		}
	}()
	stream := rtmp.NewRtmpStream()
	hlsServer := startHTTPLiveServer()
	startHTTPFlvVideoServer(stream)
	//startHTTPLiveRtmpOpera(stream)
	startRtmp(stream, hlsServer)
}