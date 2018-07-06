package controllers

import (
	"github.com/skip2/go-qrcode"
	"fmt"
	"strings"
)

func Qrcode(content string) string {
	//rand.Seed(time.Now().Unix())
	var name = content
	var pa = strings.NewReplacer("http","","//","","/","","&","","\\","",".","",":","","com","")
	name = pa.Replace(name)
	fmt.Println(name)
	var path = "static"+"/"+"qrcode"+"/"+name+".png"
	err:= qrcode.WriteFile(content,qrcode.Medium,300,path)
	if err!=nil{
		fmt.Println("write file  error")
	}
	return name
}
