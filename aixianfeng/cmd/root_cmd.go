package cmd

import (
	"aixianfeng/pkg"
	"github.com/kataras/iris/v12"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// 服务启动，以命令行的形式启动
var rootCMD = &cobra.Command{
	Use:   "root command",
	Short: "root command",
	Long:  "run web server",
	Run:   runRootCMD,
}

func runRootCMD(cmd *cobra.Command, args []string) {
	//初始化数据库连接
	pkg.DataBaseInit()
	//CTRL+C/CMD+C或收到unix终止命令时调用。
	iris.RegisterOnInterrupt(func() {
		//关闭数据库连接
		err := pkg.MyDatabase.Close()
		if err != nil {
			log.Panicf("进程终止异常！%s \n", err)
		}
	})
	app := pkg.ApplyRouter()
	err := app.Run(iris.Addr(":8082"), iris.WithCharset("UTF-8"))
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Execute 执行
func Execute() {
	// 使syncCMD成为rootCMD的子命令
	rootCMD.AddCommand(syncCMD)
	rootCMD.AddCommand(provinceCMD)
	if err := rootCMD.Execute(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
