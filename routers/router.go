package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"qrcode/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/qrcode/", &controllers.MainController{}, "GET:Index")
	beego.Router("/qrcode/", &controllers.MainController{}, "POST:Qrcode")
	beego.Router("/download/?:dir/?:name", &controllers.MainController{}, "*:Download")
	beego.Router("/test/", &controllers.MainController{}, "*:Header")
	beego.Router("/upload/", &controllers.MainController{}, "GET:Upload")
	beego.Router("/upload/", &controllers.MainController{}, "POST:UploadSave")
	beego.Router("/file/", &controllers.MainController{}, "*:File")
	beego.Router("/wiki/", &controllers.MainController{}, "GET:Wiki")
	beego.Router("/pic/", &controllers.MainController{}, "POST:AiPic")
	beego.Router("/zhuang/", &controllers.MainController{}, "*:Zhuang")

}
