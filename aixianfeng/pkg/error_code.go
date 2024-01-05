package pkg

import "fmt"

// ErrorV1 统一错误结构响应体
type ErrorV1 struct {
	Detail  string `json:"detail"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ErrorV1s []ErrorV1

// 实现Error接口方法
func (e ErrorV1) Error() string {
	return fmt.Sprintf("Detail: %s, Message: %s, Code: %d", e.Detail, e.Message, e.Code)
}

// 自定义错误信息变量
var (
	// database
	ErrorDatabase       = ErrorV1{Code: 400, Detail: "数据库错误", Message: "database error"}
	ErrorRecordNotFound = ErrorV1{Code: 400, Detail: "记录不存在", Message: "record not found"}

	// body
	ErrorBodyJson   = ErrorV1{Code: 400, Detail: "请求消息体失败", Message: "read json body fail"}
	ErrorBodyIsNull = ErrorV1{Code: 400, Detail: "参数为空", Message: "body is null"}
)
