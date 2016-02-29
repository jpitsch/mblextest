package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jpitsch/mblextest/models"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	this.Layout = "main-layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.tpl"
	this.TplName = "admin-home.tpl"
	//flash := beego.ReadFromRequest(&this.Controller)
}

func (this *AdminController) Post() {

}
