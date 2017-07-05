package controllers

import (
	"github.com/astaxie/beego"
	"ZkrtWebJack/models"
	"fmt"

)

type IndexController struct {
	beego.Controller
}

func (index *IndexController) GetLogin() {
	index.TplName = "home.html"
}

func (index *IndexController) GetError() {
	index.TplName = "error-page.html"
}

func (index *IndexController) Successs() {
	//index.TplName = "success.tpl"
	index.TplName = "zkrtsuccess.html"
}

func (index *IndexController) LoginPost() {
	//flash := beego.NewFlash()
	username, password := index.Input().Get("username"), index.Input().Get("password")
	bool,user:= models.GetWaiterForLogin(username, password);
	if bool {
		fmt.Println(user.ZkrtUser)
		index.Redirect("success", 302)
	} else {
		//flash.Error("用户名或密码错误")
		//flash.Store(&index.Controller)
		index.Redirect("error", 303)
	}
}

func (index *IndexController) LoginPy() {
	//flash := beego.NewFlash()
	username, password := index.Input().Get("username"), index.Input().Get("password")
	bool,user:= models.GetWaiterForLogin(username, password);
	fmt.Println(user)
	index.Data["json"] = bool
	index.ServeJSON()
}

