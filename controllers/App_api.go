package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/cache"
	//"ZkrtWebJack/models"
)

/*var (
	urlcache cache.Cache
)

func init() {
	urlcache, _ = cache.NewCache("memory", `{"interval":0}`)
}*/

type DroneGpsMessage struct {
	//UrlShort string
	DroneGpsLat  string
	DroneGpsLng  string
}

type DroneMessage struct {
	Drone_alt string
	Drone_y string
	Drone_r string
	Drone_p string
}

type DroneTypeId struct {
	DroneId string
}

type ApiController struct {
	beego.Controller
}

//得到用户传进来的飞行器的GPS
func (this *ApiController) DroneSetGps() {
	var result DroneGpsMessage
	longurl_lat := this.Input().Get("drone_lat")
	longurl_lng := this.Input().Get("drone_lng")
	beego.Info(longurl_lat)
	beego.Info(longurl_lng)
	result.DroneGpsLat = longurl_lat
	result.DroneGpsLng = longurl_lng
	/*urlmd5 := models.GetMD5(longurl)
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
	}*/
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *ApiController) DroneSetMessage()  {
	var result DroneMessage
	longurl_alt := this.Input().Get("drone_alt")
	longurl_y := this.Input().Get("drone_y")
	longurl_r := this.Input().Get("drone_r")
	longurl_p := this.Input().Get("drone_p")
	beego.Info(longurl_alt)
	beego.Info(longurl_y)
	beego.Info(longurl_r)
	beego.Info(longurl_p)
	result.Drone_alt = longurl_alt
	result.Drone_y = longurl_y
	result.Drone_r = longurl_r
	result.Drone_p = longurl_p
	this.Data["json"] = result
	this.ServeJSON()
}

//

//
func (this *ApiController) DroneId() {
	var result DroneTypeId
	drone_id := this.Input().Get("drone_id")
	beego.Info(drone_id)
	result.DroneId = drone_id
	this.Data["json"] = result
	this.ServeJSON()
}
