package handlers

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"testing"
)

func TestLocalPwd(t *testing.T) {
	pwd, _ := os.Getwd()
	lastIndex := strings.LastIndex(pwd, "/")
	fmt.Println(pwd)
	fmt.Println(lastIndex, pwd[:lastIndex]+"/templates")
	strs := strings.Split(os.Getenv("GOPATH"), ":")
	fmt.Println(strs[0])
	template.New("index.html")
}
