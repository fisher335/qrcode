package routers

import (
	"qrcode/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"*:Index")
    beego.Router("/qrcode/", &controllers.MainController{},"GET:Index")
    beego.Router("/qrcode/", &controllers.MainController{},"POST:Qrcode")
    beego.Router("/download/?:dir/?:name", &controllers.MainController{},"*:Download")
    beego.Router("/test/", &controllers.MainController{},"*:Header")
	beego.Router("/upload/", &controllers.MainController{},"GET:Upload")
	beego.Router("/upload/", &controllers.MainController{},"POST:UploadSave")
	beego.Router("/file/", &controllers.MainController{},"*:File")
	beego.Router("/wiki/", &controllers.MainController{},"*:RedirectGithub")

}
