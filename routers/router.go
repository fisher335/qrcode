package routers

import (
	"beegolearn01/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"*:Index")
    beego.Router("/qrcode", &controllers.MainController{},"GET:Index")
    beego.Router("/qrcode", &controllers.MainController{},"POST:Qrcode")
    beego.Router("/download/?:name", &controllers.MainController{},"*:Download")
    beego.Router("/test", &controllers.MainController{},"*:Header")
}
