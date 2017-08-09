package hutil

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/context"
)

func SaveAdminAccount(account string, ctx *context.Context) {
	maxAge := 1<<31 - 1
	ctx.SetCookie("admin_account", account, maxAge, "/")
}

func GetAdminAccount(ctx *context.Context) string {
	ck, err := ctx.Request.Cookie("admin_account")
	if err != nil {
		beego.Error(err)
		return ""
	}
	return ck.Value
}

//图床服务器路径
func IsImgPath(imgid string) (imgpath string) {
	url := "/imagehosting/"
	return fmt.Sprintf("%s%s", url, imgid)
}

func IsVersion() (version string) {
	cversion := "1.0.0Bate"
	cnf, err := config.NewConfig("ini", "conf/hello.conf")
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("-------------cnf:", cnf)
	cversion = cnf.String("hello::version")
	return cversion
}
