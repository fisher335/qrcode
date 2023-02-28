package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os"
	"path"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {

	c.TplName = "index.tpl"
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
	var content = c.GetString("url", "")
	var path = Qrcode(content)
	c.Data["path"] = path
	c.TplName = "qrcode.tpl"

}

func (c *MainController) Download() {
	var dir = c.Ctx.Input.Param(":dir")
	var name = c.Ctx.Input.Param(":name")
	var sep = string(os.PathSeparator)
	var path = "static" + sep + dir + sep + name
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

func (c *MainController) Wiki() {
	c.TplName = "wiki.tpl"

}
func (c *MainController) AiPic() {
	var content = c.GetString("url", "")
	var path, _ = CreatePic(content)
	c.Data["path"] = path[0]
	c.TplName = "aipic.tpl"

}

func (c *MainController) Zhuang() {
	c.Data["name"] = "冯文韬"
	c.Data["tel"] = "15110202919"
	c.Data["addr"] = "北京市海淀区厢黄旗东路柳浪家园南里26号楼1单元701"
	c.TplName = "zhuang.tpl"
}
