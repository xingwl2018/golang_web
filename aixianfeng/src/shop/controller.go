package shop

import (
	v1 "aixianfeng/models/v1"
	"aixianfeng/pkg"
	"aixianfeng/src/make_response"
	"github.com/kataras/iris/v12"
	"net/http"
)

// 核心控制逻辑

func createOneShopHandler(ctx iris.Context) {
	var param PostShopParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusNotFound, err.Error(), true))
		return
	}

	var province v1.Province
	if ok, _ := pkg.MyDatabase.ID(param.ProvinceId).Get(&province); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}

	var shop v1.Shop
	shop = v1.Shop{
		Name:       param.Name,
		ProvinceId: int64(province.ID),
		Location:   param.Location,
		Province:   province,
	}
	pkg.MyDatabase.InsertOne(&shop)
	ctx.JSON(make_response.MakeResponse(http.StatusOK, shop.Serializer(), false))

}

func patchOneShopHandler(ctx iris.Context) {
	var param PostShopParam
	id, _ := ctx.Params().GetInt("shop_id")
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	var shop v1.Shop

	tx := pkg.MyDatabase.NewSession()
	if ok, _ := tx.ID(id).Get(&shop); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}

	var province v1.Province
	if ok, _ := tx.ID(param.ProvinceId).Get(&province); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}

	shop.ProvinceId = int64(province.ID)
	shop.Province = province
	shop.Location = param.Location
	shop.Name = param.Name

	if _, dbError := tx.Update(&shop); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, shop.Serializer(), false))

}

func getOneShopHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("shop_id")
	var shop v1.Shop

	if ok, _ := pkg.MyDatabase.ID(id).Get(&shop); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}

	var province v1.Province
	if ok, _ := pkg.MyDatabase.ID(shop.ProvinceId).Get(&province); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}
	shop.Province = province
	ctx.JSON(make_response.MakeResponse(http.StatusOK, shop.Serializer(), false))
}

func getAllShopHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("province_id")
	var province v1.Province

	if ok, _ := pkg.MyDatabase.ID(id).Get(&province); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}

	var newProvince []v1.Province
	if dbError := pkg.MyDatabase.Where("ad_code like ?", province.AdCode[:2]+"%").Find(&newProvince); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}
	var shops []v1.Shop
	var provinceIds []int
	if len(newProvince) != 0 {
		for _, i := range newProvince {
			provinceIds = append(provinceIds, int(i.ID))
		}
	}
	if dbError := pkg.MyDatabase.In("province_id", provinceIds).Find(&shops); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}
	var results []v1.ShopSerializer
	for _, i := range shops {
		var p v1.Province
		pkg.MyDatabase.ID(i.ProvinceId).Get(&p)
		i.Province = p
		results = append(results, i.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
}
