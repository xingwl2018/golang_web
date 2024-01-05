package main

import (
	"aixianfeng/cmd"
	"aixianfeng/configs"
	"log"
	"strings"
)

var ENV string

func main() {
	//// 启动服务端口监听
	//http.ListenAndServe(":8082", nil)

	//
	if strings.EqualFold(configs.ENV, "") {
		configs.ENV = "dev"
	} else {
		configs.ENV = ENV
	}
	log.Println("Running Web Server")
	cmd.Execute()
}
