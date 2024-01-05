package pkg

import (
	"aixianfeng/configs"
	"fmt"
	"strings"
	"testing"
)

func TestGetConfigDb(t *testing.T) {
	//configs.ENV = "dev"
	//GetConfigDb()
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	////指定配置文件所在目录，可以多次调用此函数，指定多个目录
	//viper.AddConfigPath("aixianfeng/configs")
	//err := viper.ReadInConfig()
	//if err != nil {
	//	log.Panicf("读取配置文件异常!")
	//}
	//fmt.Printf("\n dev:%v", viper.Get("dev"))
	fmt.Printf("判断字符串是否相等：%v \n", strings.EqualFold(configs.ENV, ""))
}
