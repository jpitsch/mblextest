package routers

import (
	"github.com/jpitsch/mblextest/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
