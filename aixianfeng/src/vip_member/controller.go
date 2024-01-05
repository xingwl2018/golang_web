package vip_member

import (
	v1 "aixianfeng/models/v1"
	"aixianfeng/pkg"
	"aixianfeng/src/make_response"
	"github.com/kataras/iris/v12"
	"net/http"
)

// 核心控制逻辑

func getVipMemberProcessor() ([]v1.VipMember, error) {
	var (
		vipMembers []v1.VipMember
		err        error
	)
	if dbErr := pkg.MyDatabase.Find(&vipMembers); dbErr != nil {
		err = pkg.ErrorV1{
			Code:    http.StatusNotFound,
			Message: dbErr.Error(),
			Detail:  "记录未找到",
		}
		return vipMembers, err
	}

	return vipMembers, nil

}

func getVipMemberHandle(ctx iris.Context) {
	results, err := getVipMemberProcessor()
	if err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}

	var resultSerializer []v1.VipMemberSerializer
	for _, i := range results {
		resultSerializer = append(resultSerializer, i.Serializer())
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, resultSerializer, false))
}

func getVipMemberOneHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	var vip v1.VipMember
	if _, dbError := pkg.MyDatabase.Where("id = ?", id).Get(&vip); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbError, true))
		return
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, vip.Serializer(), false))
}

func patchVipMemberProcessor(id uint, param PatchVipMemberParam) (v1.VipMember, error) {
	var (
		result v1.VipMember
		err    error
	)
	if _, dbError := pkg.MyDatabase.Where("id = ?", id).Get(&result); dbError != nil {
		err = pkg.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "记录未找到",
			Message: dbError.Error(),
		}
		return result, err
	}

	if err := param.Valid().Struct(param); err != nil {
		err = pkg.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "参数校验失败",
			Message: err.Error(),
		}
		return result, err
	}

	result.LevelName = param.Level
	result.Start = param.Start
	result.End = param.End
	result.Points = param.Points
	result.Period = param.Period
	result.ToValue = param.ToValue

	if _, dbError := pkg.MyDatabase.ID(result.ID).Update(result); dbError != nil {
		err = pkg.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "更新数据失败",
			Message: dbError.Error(),
		}
	}
	return result, nil

}

func patchVipMemberHandle(ctx iris.Context) {
	var param PatchVipMemberParam
	if err := ctx.ReadJSON(&param); err != nil {
		return
	}

	id, _ := ctx.Params().GetUint("id")
	result, err := patchVipMemberProcessor(id, param)
	if err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err, true))
		return
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, result.Serializer(), false))
}
