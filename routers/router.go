package routers

import (
	"github.com/jpitsch/mblextest/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.UserController{})
    beego.Router("/userinfo", &controllers.UserController{})
}
