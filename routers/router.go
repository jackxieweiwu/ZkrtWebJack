// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"ZkrtWebJack/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/DroneMsg",
			beego.NSInclude(
				&controllers.DroneMsgController{},
			),
		),

		beego.NSNamespace("/ZkrtUser",
			beego.NSInclude(
				&controllers.ZkrtUserController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/ws", &controllers.MywebSockerControllers{},)
	beego.Router("/ws/dronestatus", &controllers.MywebSockerControllers{},"GET,POST:GetMoudleStatus")
	beego.Router("/v1/drone_id", &controllers.ApiController{},"GET,POST:GetDroneId")
	beego.Router("/v1/drone_message", &controllers.ApiController{},"GET,POST:DroneSetMessage")
	beego.Router("/views/index", &controllers.IndexController{},"GET:GetLogin")
	beego.Router("/views/index", &controllers.IndexController{}, "GET,POST:LoginPost")
	beego.Router("/views/error", &controllers.IndexController{}, "GET,POST:GetError")
	beego.Router("/views/success", &controllers.IndexController{}, "GET:Successs")
	beego.Router("/drone/status", &controllers.ApiController{}, "GET,POST:SetMoudleValuesStatus")
	beego.Router("/drone/gas", &controllers.ApiController{}, "GET,POST:SetGsaValue")

	beego.Router("/v1/login", &controllers.IndexController{}, "GET:LoginPy")
	beego.Router("/v1/pyApiDroneMsg", &controllers.PyclientController{}, "GET,POST:GetDroneMessage")
	beego.Router("/v1/pyApiMoudleMsg", &controllers.PyclientController{}, "GET,POST:GetMoudleMessage")
	beego.Router("/v1/pyApiGasMsg", &controllers.PyclientController{}, "GET,POST:GetGasMessage")
}
