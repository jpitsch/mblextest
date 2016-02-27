package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "pitsch.jason@gmail.com"
	c.Data["Title"] = "Login"
	c.Layout = "main-layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"
	c.TplName = "login.tpl"
}
