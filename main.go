package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "qrcode/routers"
)

func main() {
	beego.Run("127.0.0.1:8080")
}
