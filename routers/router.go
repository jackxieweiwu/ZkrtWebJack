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
	beego.Router("/v1/drone_gps", &controllers.ApiController{},"GET,POST:DroneSetGps")
	beego.Router("/v1/drone_message", &controllers.ApiController{},"GET,POST:DroneSetMessage")
	beego.Router("/v1/drone_id", &controllers.ApiController{},"GET,POST:DroneId")
	beego.Router("/views/index", &controllers.IndexController{},"GET:GetLogin")
	beego.Router("/views/index", &controllers.IndexController{}, "POST:LoginPost")
	beego.Router("/views/success", &controllers.IndexController{}, "GET:Successs")
}
