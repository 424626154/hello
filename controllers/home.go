package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	webs, err := models.GetWebs()
	if err != nil {
		beego.Debug(err.Error())
	}
	c.Data["Webs"] = webs
	c.TplName = "home.html"
}
