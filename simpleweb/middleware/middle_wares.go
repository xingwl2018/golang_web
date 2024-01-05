package middleware

import (
	"log"
	"net/http"
	"time"
)

// 中间件包

// http.HandlerFunc 函数实现了接口 http.Handler的ServeHTTP

// Logger 日志中间件，添加handler处理日志
func Logger(handler http.HandlerFunc) http.HandlerFunc {
	now := time.Now()
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("[Web-Server]: %s | %s", request.RequestURI, now.Format("2006/01/02 15:04:05"))
		handler.ServeHTTP(writer, request)
	}
}

// MiddleWareOfLog 日志中间件的第二种写法。
func MiddleWareOfLog(handle http.Handler) http.Handler {
	now := time.Now()
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("[Web-Server]: %s | %s", request.RequestURI, now.Format("2006/01/02 15:04:05"))
		handle.ServeHTTP(writer, request)
	})
}
