package main

import (
	_ "qrcode/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run("127.0.0.1:8080")
}

