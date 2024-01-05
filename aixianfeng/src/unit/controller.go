package unit

import (
	v1 "aixianfeng/models/v1"
	"aixianfeng/pkg"
	"aixianfeng/src/make_response"
	"github.com/kataras/iris/v12"
	"net/http"
)

// 核心控制逻辑

func createUintHandle(ctx iris.Context) {
	var (
		param         CreateUnitParam
		readJsonError error
		validError    error
	)
	readJsonError = ctx.ReadJSON(&param)
	validError = param.Valid()

	if readJsonError != nil && validError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, validError.Error(), true))
		return
	}

	tx := pkg.MyDatabase.NewSession()
	defer tx.Close()
	tx.Begin()

	var results []v1.UnitsSerializer
	var resultUnits []v1.Units
	for _, i := range param.Data {
		var u v1.Units
		u = v1.Units{
			Name:      i.Name,
			EnName:    i.EnName,
			ShortCode: i.Code,
		}
		if _, dbError := tx.InsertOne(&u); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
		resultUnits = append(resultUnits, u)
	}
	tx.Commit()
	for _, i := range resultUnits {
		var temp v1.Units
		tx.ID(i.ID).Get(&temp)
		results = append(results, temp.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))

}

func patchUintHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("unit_id")

	var param PatchUintParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	if err := param.Valid(); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	tx := pkg.MyDatabase.NewSession()
	defer tx.Close()
	tx.Begin()
	var u v1.Units
	if ok, _ := tx.ID(id).Get(&u); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}

	if param.Name != "" {
		u.Name = param.Name
		if _, dbError := tx.ID(u.ID).Cols("name").Update(&u); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
	}
	if param.Code != "" {
		u.ShortCode = param.Code
		if _, dbError := tx.ID(u.ID).Cols("short_code").Update(&u); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
	}
	if param.EnName != "" {
		u.EnName = param.EnName
		if _, dbError := tx.ID(u.ID).Cols("en_name").Update(&u); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}

	}

	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, u.Serializer(), false))
}

func getUintHandle(ctx iris.Context) {
	returnAll := ctx.URLParamDefault("return", "all_list")

	var param GetUintParam
	param.ReturnAll.ReturnAll = returnAll

	if error := param.Valid(); error != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, error.Error(), true))
		return
	}
	var (
		us      []v1.Units
		count   int64
		dbError error
	)

	if count, dbError = pkg.MyDatabase.Desc("id").FindAndCount(&us); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	if returnAll == "all_list" {
		var results []v1.UnitsSerializer
		for _, i := range us {
			results = append(results, i.Serializer())
		}
		ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
		return
	}
	if returnAll == "all_count" {
		var results = make(map[string]interface{})
		results["count"] = count
		ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
		return
	}

}
