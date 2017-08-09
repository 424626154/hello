package controllers

import (
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
)

// 图床服务器
type ImageHostingController struct {
	beego.Controller
}

func (c *ImageHostingController) Get() {
	beego.Debug("ImageHostingController Get")
	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:])
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	f, err := os.Open(filePath)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	defer f.Close()

	_, err = io.Copy(c.Ctx.ResponseWriter, f)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
}
