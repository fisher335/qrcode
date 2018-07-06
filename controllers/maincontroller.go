package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {

	c.TplName="index.tpl"
}

func (c *MainController) Header() {

	head := make(map[string]interface{})

	for k, v := range c.Ctx.Request.Header {
		head[k] = v
	}
	head["url"] = c.Ctx.Request.URL
	head["methond"] = c.Ctx.Request.Method
	head["host"] = c.Ctx.Request.Host
	head["protl"] = c.Ctx.Request.Proto
	head["remote"] = c.Ctx.Request.RemoteAddr
	head["RequestURI"] = c.Ctx.Request.RequestURI
	c.Data["head"] = head
	c.TplName = "test.tpl"
}

func (c *MainController) Qrcode() {
	var content = c.GetString("url","")
	var path = Qrcode(content)
	c.Data["path"] = path
	c.TplName = "qrcode.tpl"

}

func (c *MainController) Download() {
	var name = c.Ctx.Input.Param(":name")

	var path = "static" + "/" + "qrcode" + "/" + name
	c.Ctx.Output.Download(path, name)
	return
}
