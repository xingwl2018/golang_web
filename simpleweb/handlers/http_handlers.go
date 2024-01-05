package handlers

import (
	"golang_web/middleware"
	"net/http"
)

// RegisterHttpHandlers 用于注册web handler
func RegisterHttpHandlers() {
	//注册服务路由
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	_, err := writer.Write([]byte("welcome to simple web!"))
	//	if err != nil {
	//		panic("响应数据异常!")
	//	}
	//	fmt.Println("请求路径：", request.URL.Path)
	//})

	// 其他接口
	// 首页
	http.HandleFunc("/", home)
	http.HandleFunc("/home", home)
	http.HandleFunc("/home2", middleware.Logger(home))
	http.Handle("/home3", middleware.MiddleWareOfLog(http.HandlerFunc(home)))
	http.HandleFunc("/passage", middleware.Logger(passage))
	http.HandleFunc("/login", middleware.Logger(login))
	http.HandleFunc("/logout", middleware.Logger(logout))
	http.HandleFunc("/apis", middleware.Logger(song))
	http.HandleFunc("/progress", middleware.Logger(progress))
}
