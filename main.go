package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"hello/hutil"
	"hello/models"
	_ "hello/routers"
)

func init() {
	models.RegisterDB()
	beego.AddFuncMap("isimgpath", hutil.IsImgPath)
	beego.AddFuncMap("isversion", hutil.IsVersion)
}

func main() {
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
