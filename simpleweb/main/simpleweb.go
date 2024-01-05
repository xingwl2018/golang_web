package main

import (
	"golang_web/handlers"
	"log"
	"net/http"
)

// 定义web服务mian方法入口
func main() {
	//http handler
	handlers.RegisterHttpHandlers()
	//注册web服务，监听http服务端口：8081
	log.Fatal(http.ListenAndServe(":8081", nil))
}
