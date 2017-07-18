package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["ZkrtWebJack/controllers:DroneMsgController"] = append(beego.GlobalControllerRouter["ZkrtWebJack/controllers:DroneMsgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtWebJack/controllers:DroneMsgController"] = append(beego.GlobalControllerRouter["ZkrtWebJack/controllers:DroneMsgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtWebJack/controllers:DroneMsgController"] = append(beego.GlobalControllerRouter["ZkrtWebJack/controllers:DroneMsgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtWebJack/controllers:DroneMsgController"] = append(beego.GlobalControllerRouter["ZkrtWebJack/controllers:DroneMsgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtWebJack/controllers:DroneMsgController"] = append(beego.GlobalControllerRouter["ZkrtWebJack/controllers:DroneMsgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtWebJack/controllers:ZkrtUserController"] = append(beego.GlobalControllerRouter["ZkrtWebJack/controllers:ZkrtUserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtWebJack/controllers:ZkrtUserController"] = append(beego.GlobalControllerRouter["ZkrtWebJack/controllers:ZkrtUserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtWebJack/controllers:ZkrtUserController"] = append(beego.GlobalControllerRouter["ZkrtWebJack/controllers:ZkrtUserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtWebJack/controllers:ZkrtUserController"] = append(beego.GlobalControllerRouter["ZkrtWebJack/controllers:ZkrtUserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtWebJack/controllers:ZkrtUserController"] = append(beego.GlobalControllerRouter["ZkrtWebJack/controllers:ZkrtUserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
