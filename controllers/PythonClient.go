package controllers

import (
	"github.com/astaxie/beego"
)

type PyclientController struct {
	beego.Controller
}

func (this*PyclientController)GetDroneMessage()  {
	web := this.Input().Get("web")
	beego.Info(web)

	if web=="DroneMsg"{
		this.Data["json"] = result_dronemsg
	}else{
		this.Data["json"] = false
	}

	this.ServeJSON()
}


func (this*PyclientController)GetMoudleMessage()  {
	web := this.Input().Get("web")
	beego.Info(web)
	if web=="MoudleSta"{
		this.Data["json"] = moudlsta
	}else{
		this.Data["json"] = false
	}
	this.ServeJSON()
}

func (this*PyclientController)GetGasMessage()  {
	web := this.Input().Get("web")
	beego.Info(web)
	if web=="MoudleGas"{
		this.Data["json"] = resultsgas
	}else{
		this.Data["json"] = false
	}
	this.ServeJSON()
}
