package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"hello/hutil"
	"hello/models"
	"path"
	"strings"
	"time"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Admin() {
	account := hutil.GetAdminAccount(c.Ctx)
	if len(account) == 0 {
		c.Redirect("/admin/login", 302)
	}
	op := c.Input().Get("op")
	if op == "del" {
		id := c.Input().Get("id")
		err := models.DelWebFid(id)
		if err != nil {
			beego.Error(err.Error())
		}
		c.Redirect("/admin", 302)
	}
	webs, err := models.GetWebs()
	if err != nil {
		beego.Debug(err.Error())
	}
	c.Data["Webs"] = webs
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
	if c.Ctx.Input.IsGet() {
		account := hutil.GetAdminAccount(c.Ctx)
		if len(account) == 0 {
			c.Redirect("/admin/login", 302)
		}
		c.Data["Account"] = account
		c.TplName = "adminadd.html"
	}
	if c.Ctx.Input.IsPost() {
		title := c.Input().Get("title")
		weburl := c.Input().Get("weburl")
		desc := c.Input().Get("desc")
		ftype := c.Input().Get("type")
		atype := 0
		if ftype == "type1" {
			atype = 0
		} else if ftype == "type2" {
			atype = 1
		}
		if len(title) > 0 && len(weburl) > 0 {
			_, fh, err := c.GetFile("image")
			if err != nil {
				beego.Error(err)
			}
			showimg := ""
			if fh != nil {
				beego.Debug("fh:", fh.Filename)
				tempname := fh.Filename
				t := time.Now().Unix()
				time_str := fmt.Sprintf("%d", t)
				img_uuid := uuid.NewV4().String()
				s := []string{tempname, time_str, img_uuid}
				h := md5.New()
				h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
				showimg = hex.EncodeToString(h.Sum(nil))
				err = c.SaveToFile("image", path.Join("imagehosting", showimg))
				if err != nil {
					beego.Error(err)
				}
			}
			_, err = models.AddWeb(title, weburl, desc, showimg, int8(atype))
			if err != nil {
				beego.Error(err)
				c.Redirect("/admin/add", 302)
			}
			c.Redirect("/admin", 302)
		}
	}
}

func (c *AdminController) Up() {
	if c.Ctx.Input.IsGet() {
		account := hutil.GetAdminAccount(c.Ctx)
		if len(account) == 0 {
			c.Redirect("/admin/login", 302)
		}
		id := c.Input().Get("id")
		web, err := models.GetWebFid(id)
		if err != nil {
			beego.Error(err.Error())
		}
		beego.Debug("web:", web)
		c.Data["Web"] = web
		c.Data["Account"] = account
		c.TplName = "adminup.html"
	}
	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		title := c.Input().Get("title")
		weburl := c.Input().Get("weburl")
		desc := c.Input().Get("desc")
		ftype := c.Input().Get("type")
		atype := 0
		if ftype == "type1" {
			atype = 0
		} else if ftype == "type2" {
			atype = 1
		}
		if len(title) > 0 && len(weburl) > 0 {
			err := models.UpWebInfo(id, title, weburl, desc, int8(atype))
			if err != nil {
				beego.Error(err)
				c.Redirect("/admin/up?id="+id, 302)
			}
		}
		_, fh, err := c.GetFile("image")
		if err != nil {
			beego.Error(err)
		}
		showimg := ""
		if fh != nil {
			beego.Debug("fh:", fh.Filename)
			tempname := fh.Filename
			t := time.Now().Unix()
			time_str := fmt.Sprintf("%d", t)
			img_uuid := uuid.NewV4().String()
			s := []string{tempname, time_str, img_uuid}
			h := md5.New()
			h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
			showimg = hex.EncodeToString(h.Sum(nil))
			err = c.SaveToFile("image", path.Join("imagehosting", showimg))
			if err != nil {
				beego.Error(err)
			}
		}
		if len(showimg) > 0 {
			err := models.UpWebIcon(id, showimg)
			if err != nil {
				beego.Error(err)
				c.Redirect("/admin/up?id="+id, 302)
			}
		}
		c.Redirect("/admin", 302)
	}
}
