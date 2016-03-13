package routers

import (
	"github.com/astaxie/beego"
	"github.com/jpitsch/mblextest/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/admin", &controllers.AdminController{})

	beego.Router("/admin/edituser", &controllers.AdminController{}, "get:EditUsers")

	beego.Router("/admin/testadmin", &controllers.AdminController{}, "get:GetTestAdminMenu")
	beego.Router("/admin/createtest", &controllers.AdminController{}, "get:GetCreateTest")
	beego.Router("/admin/edittest", &controllers.AdminController{}, "get:GetEditTest")
	beego.Router("/admin/viewtests", &controllers.AdminController{}, "get:GetViewTests")

	beego.Router("/admin/test/addnextquestion", &controllers.AdminController{}, "get:GetNextQuestion")
	beego.Router("/admin/test/previousquestion", &controllers.AdminController{}, "get:GetPreviousQuestion")
	beego.Router("/admin/test/finalizetest", &controllers.AdminController{}, "get:FinalizeTestCreation")

	beego.Router("/login", &controllers.UserController{})
	beego.Router("/userinfo", &controllers.UserController{})
}
