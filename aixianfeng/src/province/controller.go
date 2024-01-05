package province

import (
	v1 "aixianfeng/models/v1"
	"aixianfeng/pkg"
	"aixianfeng/src/make_response"
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
)

// 核心控制逻辑

func getProvinceHandler(ctx iris.Context) {
	var param GetProvinceParam
	returnAll := ctx.URLParamDefault("return", "all_list")
	param.Return = returnAll
	level := ctx.URLParam("level")
	if level != "" {
		param.Level = level
	}

	if err := param.Valid(); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusNotFound, err.Error(), true))
		return
	}

	query := pkg.MyDatabase.NewSession()
	if level != "" {
		query = query.Where("level = ?", level)
	}

	var provinces []v1.Province
	var count int64
	var dbError error
	if count, dbError = query.FindAndCount(&provinces); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusNotFound, dbError.Error(), true))
		return
	}
	if returnAll == "all_count" {
		var countMap = make(map[string]int64)
		countMap["count"] = count
		ctx.JSON(make_response.MakeResponse(http.StatusOK, countMap, false))
		return
	}
	var results []v1.ProvinceSerializer
	for _, i := range provinces {
		results = append(results, i.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
}

func getOneHandler(ctx iris.Context) {
	name := ctx.URLParam("name")

	var province v1.Province

	fmt.Println(name)
	if name == "" {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, fmt.Errorf("url param name should not be null"), true))
		return
	}
	if ok, _ := pkg.MyDatabase.Where("name like ?", "%"+name+"%").Get(&province); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}
	var child []v1.Province
	var results = make(map[string]interface{})
	if province.Level == "province" {
		pkg.MyDatabase.Where("ad_code like ? AND level = ?", province.AdCode[:2]+"%", "city").Find(&child)
	} else if province.Level == "city" {
		pkg.MyDatabase.Where("city_code = ? AND level = ?", province.CityCode, "district").Find(&child)
	}
	if len(child) == 0 {
		ctx.JSON(make_response.MakeResponse(http.StatusOK, province.Serializer(), false))
		return
	}
	results[province.Level] = province.Serializer()
	var childResults []v1.ProvinceSerializer
	for _, i := range child {
		childResults = append(childResults, i.Serializer())
	}
	results["child"] = childResults
	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
}
