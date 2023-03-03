package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "qrcode/routers"
)

func main() {
	beego.Run("0.0.0.0:80")
}
