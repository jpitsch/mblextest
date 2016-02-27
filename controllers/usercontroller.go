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
	this.Layout = "main-layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.tpl"
	this.TplName = "user.tpl"
	//flash := beego.ReadFromRequest(&this.Controller)

	this.Data["UserName"]=this.GetSession("username")
	this.Data["Email"]=this.GetSession("email")
	
}

func (this *UserController) Post() {
	flash := beego.NewFlash()

	username := this.GetString("username")
	//password := this.GetString("password")

	u := models.LoadUser(username)

	if u.UserName == "" {
		fmt.Println("No valid username")
		flash.Notice("Error processing login.")
		flash.Store(&this.Controller)
		this.Redirect("/", 302)
	}

	//if(u.Password != password) {
	//	fmt.Println("No valid username")
	//	this.Redirect("/", 302)
	//}
	
	this.SetSession("username", u.UserName)
	this.SetSession("email", u.Email)

	this.Redirect("/userinfo", 302)
}

func (this *UserController) Create() {

}