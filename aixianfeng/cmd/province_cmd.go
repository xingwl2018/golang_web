package cmd

import (
	v1 "aixianfeng/models/v1"
	"aixianfeng/pkg"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

// 省市区命令处理
// 使用高德API
var api = "https://restapi.amap.com/v3/config/district?subdistrict=3&key=71fcb8ce4a9910aad34e0de0d613cf44"

var provinceCMD = &cobra.Command{
	Use:        "province",
	ArgAliases: []string{"p", "-p", "P", "-P"},
	Run:        importData,
}

func importData(cmd *cobra.Command, args []string) {
	// 获取
	response, err := http.Get(api)
	if err != nil {
		return
	}
	pkg.DataBaseInit()
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	jsonByte := gjson.ParseBytes(content)
	for _, i := range jsonByte.Get("districts").Array()[0].Get("districts").Array() {
		if i.Get("level").String() == "province" {

			var province v1.Province
			province = v1.Province{
				Name:   i.Get("name").String(),
				AdCode: i.Get("adcode").String(),
				Center: i.Get("center").String(),
				Level:  i.Get("level").String(),
			}
			//fmt.Println("Province", province)
			pkg.MyDatabase.InsertOne(&province)
			if len(i.Get("districts").Array()) != 0 {
				for _, j := range i.Get("districts").Array() {
					if j.Get("level").String() == "city" {
						var city v1.Province
						city = v1.Province{
							Name:     j.Get("name").String(),
							AdCode:   j.Get("adcode").String(),
							Center:   j.Get("center").String(),
							Level:    j.Get("level").String(),
							CityCode: j.Get("citycode").String(),
						}
						//fmt.Println("city", city)
						pkg.MyDatabase.InsertOne(&city)
						if len(j.Get("districts").Array()) != 0 {
							for _, k := range j.Get("districts").Array() {
								if k.Get("level").String() == "district" {
									var district v1.Province
									district = v1.Province{
										Name:     k.Get("name").String(),
										AdCode:   k.Get("adcode").String(),
										Center:   k.Get("center").String(),
										Level:    k.Get("level").String(),
										CityCode: k.Get("citycode").String(),
									}
									//fmt.Println("district", district)
									pkg.MyDatabase.InsertOne(&district)
								}
							}
						}
					}

				}

			}
		}
	}

}
