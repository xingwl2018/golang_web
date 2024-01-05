package pkg

import (
	"aixianfeng/configs"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

// MyDatabase 所有操作均需要事先创建并配置 ORM 引擎才可以进行。
// 一个 Engine 引擎用于对单个数据库进行操作，一个 Engine Group 引擎用于对读写分离的数据库或者负载均衡的数据库进行操作。
var MyDatabase *xorm.Engine

// 定义数据库连接信息变量
var (
	drivenName string
	db         string
	port       string
	password   string
	user       string
	dsn        string
)

func GetConfigDb() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	//指定配置文件所在目录，可以多次调用此函数，指定多个目录
	viper.AddConfigPath("../configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("读取配置文件异常! %v", err)
	}
	// 读取指定环境的配置信息
	drivenName = viper.GetString(configs.ENV + ".d1.drivenName")
	db = viper.GetString(configs.ENV + ".d1.db")
	user = viper.GetString(configs.ENV + ".d1.user")
	password = viper.GetString(configs.ENV + ".d1.password")
	fmt.Printf("读取结果：drivenName：%s, db: %s, user:%s, password:%s \n", drivenName, db, user, password)
}

// 初始化连接信息
func init() {
	GetConfigDb()
	if drivenName == "mysql2" {
		config := mysql.NewConfig()
		config = &mysql.Config{
			User:   user,
			Passwd: password,
			DBName: db,
			Addr:   port,
		}
		dsn = config.FormatDSN()
	} else {
		dsn = fmt.Sprintf("%s:%s@/%s??timeout=5000ms&readTimeout=5000ms&writeTimeout=5000ms&charset=utf8mb4&parseTime=true&loc=Local", user, password, db)
		//dsn = fmt.Sprintf("root:admin123@/beequick_dev?charset=utf8&parseTime=true&loc=Local")
	}

}

// DataBaseInit 对外调用：初始化数据库连接
func DataBaseInit() {
	var err error
	MyDatabase, err = xorm.NewEngine(drivenName, dsn)
	if err != nil {
		panic(err)
		return
	}
	//在控制台打印出生成的SQL语句
	MyDatabase.ShowSQL(true)
	MyDatabase.Logger()
	MyDatabase.Charset("utf8mb4")
	//名称映射规则
	//SnakeMapper 支持struct为驼峰式命名，表结构为下划线命名之间的转换，这个是默认的Maper；
	//SameMapper 支持结构体名称和对应的表名称以及结构体field名称与对应的表字段名称相同的命名；
	//GonicMapper 和SnakeMapper很类似，但是对于特定词支持更好，比如ID会翻译成id而不是i_d。
	MyDatabase.SetMapper(names.GonicMapper{})
	//MyDatabase.SetTableMapper(names.SameMapper{})
	// 设置连接池
	// 设置连接池的空闲数大小
	MyDatabase.SetMaxIdleConns(10)
	//设置最大打开连接数
	MyDatabase.SetMaxOpenConns(200)
	//设置连接的最大生存时间
	MyDatabase.SetConnMaxLifetime(60 * time.Second)
}
