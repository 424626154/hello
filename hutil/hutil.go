package hutil

import (
	"github.com/astaxie/beego"
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
