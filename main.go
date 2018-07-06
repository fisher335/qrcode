package main

import (
	_ "beegolearn01/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run("127.0.0.1:8080")
}

