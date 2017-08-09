package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
	"os"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.Router("/admin", &controllers.AdminController{}, "*:Admin")
	beego.Router("/admin/login", &controllers.AdminController{}, "*:Login")
	beego.Router("/admin/logout", &controllers.AdminController{}, "*:Logout")
	beego.Router("/admin/add", &controllers.AdminController{}, "*:Add")
	beego.Router("/admin/up", &controllers.AdminController{}, "*:Up")

	// 附件处理
	os.Mkdir("imagehosting", os.ModePerm)
	beego.Router("/imagehosting/:all", &controllers.ImageHostingController{})
	beego.Router("/imagehosting", &controllers.ImageHostingController{})
}
