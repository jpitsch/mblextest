package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jpitsch/mblextest/models"
)

type TestController struct {
	beego.Controller
}

func (this *TestController) Get() {
	this.Layout = "main-layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.tpl"
	this.TplName = "admin-home.tpl"
}

func (this *AdminController) Post() {

}
