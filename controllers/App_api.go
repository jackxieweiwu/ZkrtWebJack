package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/cache"
	"ZkrtWebJack/models"
	"ZkrtWebJack/util/uid"
)

/*var (
	urlcache cache.Cache
)*/
//var countryCapitalMap map[string]models.DroneMessage
var result_dronemsg models.DroneMessage
var moudlsta models.DroneMoudleStatus//获取传送过来的模块状态
var resultsgas models.DroneGas //传送毒气值
/*
func init() {
	//urlcache, _ = cache.NewCache("memory", `{"interval":0}`)
	//countryCapitalMap = make(map[string]models.DroneMessage)
}
*/

type ApiController struct {
	beego.Controller
}

//得到用户传进来的飞行器的GPS
/*func (this *ApiController) DroneSetGps() {
	var result models.DroneGpsMessage
	longurl_lat := this.Input().Get("drone_lat")
	longurl_lng := this.Input().Get("drone_lng")
	beego.Info(longurl_lat)
	beego.Info(longurl_lng)
	result.DroneGpsLat = longurl_lat
	result.DroneGpsLng = longurl_lng
	*//*urlmd5 := models.GetMD5(longurl)
	beego.Info(urlmd5)
	if urlcache.IsExist(urlmd5) {
		result.UrlShort = urlcache.Get(urlmd5).(string)
	} else {
		result.UrlShort = models.Generate()
		err := urlcache.Put(urlmd5, result.UrlShort, 0)
		if err != nil {
			beego.Info(err)
		}
		err = urlcache.Put(result.UrlShort, longurl, 0)
		if err != nil {
			beego.Info(err)
		}
	}*//*
	this.Data["json"] = result
	this.ServeJSON()
}*/

func (this *ApiController) DroneSetMessage()  {
	drone_id := this.Input().Get("drone_id")
	longurl_lat := this.Input().Get("drone_lat")
	longurl_lng := this.Input().Get("drone_lng")
	longurl_alt := this.Input().Get("drone_alt")
	longurl_y := this.Input().Get("drone_y")
	longurl_r := this.Input().Get("drone_r")
	longurl_p := this.Input().Get("drone_p")

	beego.Info(drone_id)
	beego.Info(longurl_lat)
	beego.Info(longurl_lng)
	beego.Info(longurl_alt)
	beego.Info(longurl_y)
	beego.Info(longurl_r)
	beego.Info(longurl_p)
	result_dronemsg.DroneId = drone_id
	result_dronemsg.DroneGpsLat = longurl_lat
	result_dronemsg.DroneGpsLng = longurl_lng
	result_dronemsg.Drone_alt = longurl_alt
	result_dronemsg.Drone_y = longurl_y
	result_dronemsg.Drone_r = longurl_r
	result_dronemsg.Drone_p = longurl_p
	this.Data["json"] = result_dronemsg
	this.ServeJSON()
}

//获取DroneiD
func (this *ApiController) GetDroneId()  {
	var results models.DroneTypeId
	drone_key := this.Input().Get("drone_key")
	if drone_key != "ZKRT_JACK"{
		results.Drone_Id = "000000"
	}else{
		results.Drone_Id = uid.NewId()
	}
	this.Data["json"] = results
	this.ServeJSON()
}

//获取传送过来的模块状态
func (this *ApiController) SetMoudleValuesStatus(){
	drone_id := this.Input().Get("drone_id")
	drone_tmp := this.Input().Get("drone_tmp")
	drone_ram := this.Input().Get("drone_ram")
	drone_gas := this.Input().Get("drone_gas")
	drone_prop := this.Input().Get("drone_prop")
	drone_expo := this.Input().Get("drone_expo")
	drone_obst := this.Input().Get("drone_obst")
	drone_doul := this.Input().Get("drone_doul")
	drone_3d := this.Input().Get("drone_3d")
	beego.Info(drone_id)
	beego.Info(drone_tmp)
	beego.Info(drone_ram)
	beego.Info(drone_gas)
	beego.Info(drone_prop)
	beego.Info(drone_expo)
	beego.Info(drone_obst)
	beego.Info(drone_doul)
	beego.Info(drone_3d)

	moudlsta.DroneId = drone_id
	moudlsta.DroneTmp = drone_tmp
	moudlsta.DroneRam = drone_ram
	moudlsta.DroneGas = drone_gas
	moudlsta.DroneProp = drone_prop
	moudlsta.DroneExpo = drone_expo
	moudlsta.DroneObst = drone_obst
	moudlsta.DroneDoul = drone_doul
	moudlsta.Drone3D = drone_3d
	this.Data["json"] = moudlsta
	this.ServeJSON()
}

//传送毒气值
func (this *ApiController) SetGsaValue(){
	drone_id := this.Input().Get("drone_id")
	drone_co := this.Input().Get("drone_co")
	drone_nh3 := this.Input().Get("drone_nh3")
	drone_h2s := this.Input().Get("drone_h2s")
	drone_co2 := this.Input().Get("drone_co2")
	drone_voc := this.Input().Get("drone_voc")
	drone_cl2 := this.Input().Get("drone_cl2")
	drone_so2 := this.Input().Get("drone_so2")
	drone_ch4 := this.Input().Get("drone_ch4")
	beego.Info(drone_id)
	beego.Info(drone_co)
	beego.Info(drone_nh3)
	beego.Info(drone_h2s)
	beego.Info(drone_co2)
	beego.Info(drone_voc)
	beego.Info(drone_cl2)
	beego.Info(drone_so2)
	beego.Info(drone_ch4)
	resultsgas.DroneId = drone_id
	resultsgas.DroneCO = drone_co
	resultsgas.DroneNH3 = drone_nh3
	resultsgas.DroneH2S = drone_h2s
	resultsgas.DroneCO2 = drone_co2
	resultsgas.DroneVOC = drone_voc
	resultsgas.DroneCL2 = drone_cl2
	resultsgas.DroneSO2 = drone_so2
	resultsgas.DroneCH4 = drone_ch4
	this.Data["json"] = resultsgas
	this.ServeJSON()
}

/*func (this *ApiController) GetDroneId()  {
	var result models.DroneTypeId

	shorturl := this.Input().Get("dronemd5")
	result.Drone_Id = uid.NewId()
	if urlcache.IsExist(shorturl) {
		result.Drone_Id = urlcache.Get(shorturl).(string)
	} else {
		result.Drone_Id = ""
	}
	this.Data["json"] = result
	this.ServeJSON()
}*/

/*
//获取app终端传过来的Drone_id
func (this *ApiController) SetDroneId() {
	var result models.DroneTypeId
	drone_id := this.Input().Get("drone_id")
	beego.Info(drone_id)
	result.Drone_Id = drone_id
	urlmd5 := models.GetMD5(drone_id)
	beego.Info(urlmd5)
	if urlcache.IsExist(urlmd5) {
		result.Drone_md = urlcache.Get(urlmd5).(string)
	} else {
		result.Drone_md = models.Generate()
		err := urlcache.Put(urlmd5, result.Drone_md, 0)
		if err != nil {
			beego.Info(err)
		}
		err = urlcache.Put(result.Drone_md, drone_id, 0)
		if err != nil {
			beego.Info(err)
		}
	}
	this.Data["json"] = result
	this.ServeJSON()
}
*/
