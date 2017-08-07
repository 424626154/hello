package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "*:Index")
	// beego.Router("/admin", &controllers.AdminController{}, "*:Admin")
	// beego.Router("/admin/login", &controllers.AdminController{}, "*:Login")
	// beego.Router("/admin/logout", &controllers.AdminController{}, "*:Logout")
	// beego.Router("/admin/add", &controllers.AdminController{}, "*:Add")
}
