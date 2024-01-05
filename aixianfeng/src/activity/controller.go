package activity

import (
	v1 "aixianfeng/models/v1"
	"aixianfeng/pkg"
	"aixianfeng/src/common"
	"aixianfeng/src/make_response"
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
	"strings"
	"time"
)

// 核心控制逻辑

func createOneActivityHandle(ctx iris.Context) {
	var param CreateActivityParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(
			http.StatusBadRequest, err.Error(), true,
		))
		return
	}

	if err := param.Valid(); err != nil {
		ctx.JSON(make_response.MakeResponse(
			http.StatusBadRequest, err.Error(), true,
		))
		return
	}

	var (
		start, end time.Time
		err        error
	)
	start, end, err = param.timeHandle()
	if err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	tx := pkg.MyDatabase.NewSession()
	defer tx.Close()
	tx.Begin()
	var shops []v1.Shop
	if dbError := tx.In("id", param.ShopIds).Find(&shops); dbError != nil {
		ctx.JSON(make_response.MakeResponse(
			http.StatusBadRequest, pkg.ErrorRecordNotFound, true,
		))
		return
	}
	var activity v1.Activity
	activity = v1.Activity{
		Name:    param.Name,
		Title:   param.Title,
		Start:   start,
		End:     end,
		ShopIds: param.ShopIds,
		Status:  v1.DOING,
	}
	if param.Avatar != "" {
		activity.Avatar = param.Avatar
	}

	if _, dbError := tx.InsertOne(&activity); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(
			http.StatusBadRequest, dbError.Error(), true,
		))
		return
	}

	for _, i := range shops {
		var shop2Activity v1.Shop2Activity
		shop2Activity = v1.Shop2Activity{
			ShopId:     int64(i.ID),
			ActivityId: int64(activity.ID),
		}
		if _, dbError := tx.InsertOne(&shop2Activity); dbError != nil {
			tx.Rollback()
			ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
			return
		}
	}

	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, activity.Serializer(), false))

}

func patchOneActivityHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("activity_id")
	var param PatchActivityParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	tx := pkg.MyDatabase.NewSession()
	defer tx.Commit()
	tx.Begin()
	var activity v1.Activity
	if ok, _ := tx.ID(id).Get(&activity); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}

	activity = v1.Activity{
		Name:    param.Name,
		Title:   param.Title,
		ShopIds: param.ShopIds,
		Start:   common.ToTime2(param.Start),
		End:     common.ToTime2(param.End),
		Status:  param.Status,
		Avatar:  param.Avatar,
	}

	if _, dbError := tx.Update(&activity); dbError != nil {
		tx.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	tx.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, activity.Serializer(), false))
}

func getOneActivityHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("activity_id")
	var activity v1.Activity
	if ok, dbError := pkg.MyDatabase.ID(id).Get(&activity); dbError != nil || !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}

	ctx.JSON(make_response.MakeResponse(http.StatusOK, activity.Serializer(), false))

}

func getAllActivityHandle(ctx iris.Context) {

	var activities []v1.Activity
	status := ctx.URLParam("status")
	returnAll := ctx.URLParamDefault("return", "all_list")
	var param GetActivityParam
	param.Status = status
	param.ReturnAll = returnAll
	if err := param.Valid(); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	query := pkg.MyDatabase.NewSession()
	defer query.Close()
	query.Begin()
	if param.Status != "" {
		key := func(status string) int {
			for k, v := range v1.ActivityStatusEn {
				if strings.ToUpper(v) == strings.ToUpper(status) {
					return k
				}
			}
			return -1
		}
		query = query.Where("status = ?", key(param.Status))
	}
	var (
		total int64
		err   error
	)

	if total, err = query.FindAndCount(&activities); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	if param.ReturnAll == "all_list" {
		var results []v1.ActivitySerializer
		for _, i := range activities {
			results = append(results, i.Serializer())
		}
		query.Commit()
		ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
		return
	}

	if param.ReturnAll == "all_count" {
		var results = make(map[string]int64)
		results["count"] = total
		query.Commit()
		ctx.JSON(make_response.MakeResponse(
			http.StatusOK, results, false,
		))
	}

}

func getAllByShopIdActivityHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("shop_id")
	fmt.Println(id, "id")
	var (
		dbError error
	)
	var shop2Activity []v1.Shop2Activity
	if _, dbError = pkg.MyDatabase.Where("shop_id = ?", id).FindAndCount(&shop2Activity); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}
	var activityIds []int64
	for _, i := range shop2Activity {
		activityIds = append(activityIds, i.ActivityId)
	}
	var activities []v1.Activity
	if dbError := pkg.MyDatabase.In("id", activityIds).Find(&activities); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError.Error(), true))
		return
	}

	var results []v1.ActivitySerializer
	for _, i := range activities {
		results = append(results, i.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))

}
