package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	webs, err := models.GetFTypeWebs(0)
	if err != nil {
		beego.Debug(err.Error())
	}
	c.Data["Webs"] = webs
	c.TplName = "home.html"
}

func (c *HomeController) Blogs() {
	webs, err := models.GetFTypeWebs(1)
	if err != nil {
		beego.Debug(err.Error())
	}
	c.Data["Webs"] = webs
	c.TplName = "blogs.html"
}
