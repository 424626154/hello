package main

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	// "hello/models"
	_ "hello/routers"
)

func init() {
	// models.RegisterDB()
}

func main() {
	// orm.RunSyncdb("default", false, true)
	beego.Run()
}
