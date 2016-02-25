package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jpitsch/mblextest/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.TplName = "user.tpl"
}

func (this *UserController) Post() {
	fmt.Println("Inside")
	username := this.GetString("username")
	//password := this.GetString("password")

	u := models.LoadUser(username)
	this.Data["UserName"]
	this.Data["Email"]
	this.Data["u"]=u

	this.Ctx.Redirect(302, "/userinfo")
}
