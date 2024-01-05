package handlers

import (
	"fmt"
	"time"
)

// 公共函数

// 处理时间格式化函数
func timeHandle(date time.Time) string {
	return date.Format("2006/01/02")
}

func timeFormat(date time.Time) string {
	return fmt.Sprintf(date.Format(time.Stamp) + " By ")
}
