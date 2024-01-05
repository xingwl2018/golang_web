package cmd

import (
	v1 "aixianfeng/models/v1"
	"aixianfeng/pkg"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// 数据结构同步

var syncCMD = &cobra.Command{
	Use: "sync2",
	//命令描述
	Short: "xorm.Syn2(model)",
	Run:   sync2,
}

func sync2(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Panic("You should add one argument at least")
		return
	}
	pkg.DataBaseInit()
	//同步db表结构
	if args[0] == "db" {
		//遍历表结构模型，进行数据结构同步
		for _, i := range tables() {
			if err := pkg.MyDatabase.Sync2(i); err != nil {
				fmt.Println(err)
			}
		}
	}
	//初始化vip会员信息
	if args[0] == "vip" {
		vipMember()
	}
	//删除表
	if args[0] == "drop" {
		pkg.MyDatabase.DropTables(new(v1.Order))
	}
}

func tables() []interface{} {
	return []interface{}{
		new(v1.Account),
		new(v1.VipMember),
		new(v1.ExchangeCoupon),
		new(v1.Account2ExchangeCoupon),
		new(v1.RuleForExchangeOrCoupon),
		new(v1.Shop),
		new(v1.Province),
		new(v1.Activity),
		new(v1.Activity2Product),
		new(v1.Shop2Activity),
		new(v1.Product),
		new(v1.Product2Tags),
		new(v1.Tags),
		new(v1.Shop2Tags),
		new(v1.Brands),
		new(v1.Units),
		new(v1.Order),
	}
}

// 初始化会员
func vipMember() bool {
	if _, err := pkg.MyDatabase.Insert(v1.DefaultVipMemberRecord()); err != nil {
		return false
	}
	return true
}
