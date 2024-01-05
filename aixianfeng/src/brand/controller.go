package brand

import (
	v1 "aixianfeng/models/v1"
	"aixianfeng/pkg"
	"aixianfeng/src/make_response"
	"github.com/kataras/iris/v12"
	"net/http"
)

// 核心控制逻辑

func createBrandHandle(ctx iris.Context) {
	var param CreateBrandParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	if err := param.Valid(); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	var brand v1.Brands
	brand = v1.Brands{
		ChName: param.Name,
		EnName: param.EnName,
	}
	tx := pkg.MyDatabase.NewSession()
	defer tx.Close()
	tx.Begin()
	if _, dbError := tx.Insert(&brand); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, brand.Serializer(), false))
}

func patchBrandHandle(ctx iris.Context) {
	var param PatchBrandParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	id, _ := ctx.Params().GetInt("brand_id")
	var brand v1.Brands
	tx := pkg.MyDatabase.NewSession()
	defer tx.Close()
	tx.Begin()
	if _, dbError := tx.ID(id).Get(&brand); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	if param.Name != "" {
		brand.ChName = param.Name
		if _, dbError := tx.ID(brand.ID).Cols("ch_name").Update(&brand); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
	}
	if param.EnName != "" {
		brand.EnName = param.EnName
		if _, dbError := tx.ID(brand.ID).Cols("en_name").Update(&brand); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
	}
	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, brand.Serializer(), false))

}

func getBrandHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("brand_id")
	var brand v1.Brands

	if ok, dbError := pkg.MyDatabase.ID(id).Get(&brand); dbError != nil || !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, brand.Serializer(), false))
}

func getBrandsHandle(ctx iris.Context) {
	returnAll := ctx.URLParamDefault("return", "all_list")

	var (
		brands  []v1.Brands
		count   int64
		dbError error
	)

	if count, dbError = pkg.MyDatabase.FindAndCount(&brands); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	if returnAll == "all_count" {
		var results = make(map[string]interface{})
		results["count"] = count
		ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
		return
	}
	var results []v1.BrandsSerializer
	for _, i := range brands {
		results = append(results, i.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))

}

func createBrandsHandle(ctx iris.Context) {
	var param CreateBrandsParam
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
	var brandIds []uint
	for _, i := range param.Data {
		var temp v1.Brands
		temp = v1.Brands{
			ChName: i.Name,
			EnName: i.EnName,
		}
		if _, dbError := tx.Insert(&temp); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
		brandIds = append(brandIds, temp.ID)
	}

	var brands []v1.Brands
	if dbError := tx.In("id", brandIds).Find(&brands); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	var results []v1.BrandsSerializer
	for _, i := range brands {
		results = append(results, i.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))

}
