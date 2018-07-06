package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"path"
	"github.com/astaxie/beego/logs"
	"os"
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
	var dir = c.Ctx.Input.Param(":dir")
	var name = c.Ctx.Input.Param(":name")
	var sep =string(os.PathSeparator)
	logs.Info(name)
	var path = "static" + sep+dir+sep + name
	c.Ctx.Output.Download(path, name)
	return
}

func (c *MainController) File() {
	file, _ := ListDir("static/videos")
	c.Data["file"] = file
	c.TplName = "file.tpl"
}

func (c *MainController) Upload() {

	c.TplName = "upload.tpl"
}

func (c *MainController) UploadSave() {

	f, h, _ := c.GetFile("file")
	fileName := h.Filename
	fmt.Println(fileName)
	f.Close()
	c.SaveToFile("file", path.Join("static/videos", fileName))
	c.Redirect("/file/", 302)
}

func (c *MainController) RedirectGithub(){
	c.Redirect("https://github.com/fisher335/wiki/issues",302)
	c.StopRun()
}