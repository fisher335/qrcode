package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/config"
	"github.com/skip2/go-qrcode"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type DallE struct {
	apiKey  string
	userId  string
	timeOut time.Duration // 超时时间, 0表示不超时
}

type DallEReq struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

type DallEResp struct {
	Created int64       `json:"created"`
	Data    []DallEData `json:"data"`
	Error   DallError   `json:"error"`
}

type DallEData struct {
	Url string `json:"url"`
}

type DallError struct {
	Message string `json:"message"`
}

func Qrcode(content string) string {
	//rand.Seed(time.Now().Unix())
	var name = content
	var pa = strings.NewReplacer("http", "", "//", "", "/", "", "&", "", "\\", "", ".", "", ":", "", "com", "")
	name = pa.Replace(name)
	fmt.Println(name)
	var path = "static" + "/" + "qrcode" + "/" + name + ".png"
	err := qrcode.WriteFile(content, qrcode.Medium, 300, path)
	if err != nil {
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
		files = append(files, fi.Name())

	}
	return files, nil
}

func CreatePic(prompt string) ([]string, error) {

	size := "1024x1024"

	requestBody := DallEReq{
		Prompt: prompt,
		N:      1,
		Size:   size,
	}
	key, _ := config.String("aikey")
	postData, _ := json.Marshal(requestBody)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", bytes.NewReader(postData))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", key))
	resp, e := client.Do(req)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, e
	}

	var dallEResp DallEResp
	err = json.Unmarshal(body, &dallEResp)
	if err != nil {
		return nil, err
	}

	if len(dallEResp.Error.Message) > 0 {
		return nil, fmt.Errorf("%v", dallEResp.Error.Message)
	}

	var out []string
	for _, v := range dallEResp.Data {
		out = append(out, v.Url)
	}
	return out, nil
}
