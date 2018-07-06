package controllers

import (
	"github.com/skip2/go-qrcode"
	"fmt"
	"strings"
	"io/ioutil"
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


func ListDir(dirPth string) (files []string, err error) {
	files = make([]string, 0, 5)

	dir, err := ioutil.ReadDir(dirPth)

	if err != nil {
		return nil, err
	}

	//PthSep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}

		//files = append(files, dirPth+PthSep+fi.Name())
		files = append(files,fi.Name())

	}
	return files, nil
}
