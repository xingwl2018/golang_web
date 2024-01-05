package common

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// 辅助函数

// GenerateFromPassword 加密密码
func GenerateFromPassword(password string, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), cost)
}

// CompareHashAndPassword 比较密码
func CompareHashAndPassword(hashed []byte, password []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hashed, password); err != nil {
		return false
	}
	return true
}

// GenerateToken 生成token
func GenerateToken(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// ToTime 转换时间
func ToTime(value string) (time.Time, error) {
	var (
		timeValue time.Time
		err       error
	)
	if len(value) == 10 {
		value = fmt.Sprintf("%s 00:00:00", value)
	}
	v, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
	if err != nil {
		return timeValue, err
	}
	return v, nil
}

func ToTime2(value string) time.Time {
	var (
		timeValue time.Time
		err       error
	)
	if len(value) == 10 {
		value = fmt.Sprintf("%s 00:00:00", value)
	}
	v, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
	if err != nil {
		return timeValue
	}
	return v
}
