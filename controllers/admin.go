package controllers

import (
	"github.com/astaxie/beego"
	"hello/hutil"
	"hello/models"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Admin() {
	account := hutil.GetAdminAccount(c.Ctx)
	if len(account) == 0 {
		c.Redirect("/admin/login", 302)
	}
	c.Data["Account"] = account
	c.TplName = "admin.html"
}

func (c *AdminController) Login() {
	if c.Ctx.Input.IsPost() {
		account := c.Input().Get("account")
		password := c.Input().Get("password")
		autologin := c.Input().Get("autologin")
		beego.Debug(account, password, autologin)

		if len(account) > 0 && len(password) > 0 {
			admin, err := models.GetOneAdmin(account)
			if err != nil {
				beego.Debug(err)
				c.Data["Error"] = err.Error()
			} else {
				if admin.Id > 0 {
					if admin.Password == password {
						hutil.SaveAdminAccount(account, c.Ctx)
						c.Redirect("/admin", 302)
					} else {
						c.Data["Error"] = "密码错误"
					}
				} else {
					c.Data["Error"] = "账号不存在"
				}

			}
		} else {

		}
		c.TplName = "alogin.html"
	}
	if c.Ctx.Input.IsGet() {
		c.Data["Error"] = ""
		c.TplName = "alogin.html"
	}

}

func (c *AdminController) Logout() {
	hutil.SaveAdminAccount("", c.Ctx)
	c.Redirect("/admin/login", 302)
}

func (c *AdminController) Add() {
	account := hutil.GetAdminAccount(c.Ctx)
	if len(account) == 0 {
		c.Redirect("/admin/login", 302)
	}
	c.Data["Account"] = account
	c.TplName = "adminadd.html"
}
