package account

import (
	v1 "aixianfeng/models/v1"
	"aixianfeng/pkg"
	"aixianfeng/src/common"
	"aixianfeng/src/make_response"
	"github.com/kataras/iris/v12"
	"net/http"
	"strings"
	"time"
)

// 核心控制逻辑

func registerProcessor(param RegisterParam) (v1.AccountGroupVip, error) {
	var (
		account v1.AccountGroupVip
		errV1   pkg.ErrorV1
	)
	if err := param.Valid().Struct(param); err != nil {
		return account, pkg.ErrorV1{
			Code:    http.StatusBadRequest,
			Message: "param not valid",
			Detail:  "请求参数校验不通过，请检查参数",
		}
	}

	var vipMember v1.VipMember

	if _, dbErr := pkg.MyDatabase.Where("level_name = ?", strings.ToUpper("v0")).Get(&vipMember); dbErr != nil {
		return account, pkg.ErrorV1{
			Code:    http.StatusBadRequest,
			Message: dbErr.Error(),
			Detail:  "会员等级未存在",
		}
	}
	account.VipMember = vipMember

	hashPassword, _ := common.GenerateFromPassword(param.Password, 8)
	hashToken := common.GenerateToken(20)

	account.Account = v1.Account{
		Phone:       param.Phone,
		Password:    string(hashPassword),
		Token:       hashToken,
		Points:      0,
		VipMemberID: vipMember.ID,
		VipTime:     time.Now(),
	}

	if _, err := pkg.MyDatabase.InsertOne(&account.Account); err != nil {
		errV1 = pkg.ErrorV1{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Detail:  "用户注册发生错误",
		}
		return account, errV1
	}
	return account, nil
}

func registerHandle(ctx iris.Context) {
	var param RegisterParam
	err := ctx.ReadJSON(&param)
	if err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err, true))
		return
	}
	account, err := registerProcessor(param)
	if err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err, true))
		return
	}
	ctx.JSON(iris.Map(make_response.MakeResponse(http.StatusOK, account.SerializerForGroup(), false)))

}

func signProcessor(param RegisterParam) (v1.AccountGroupVip, error) {

	var (
		account v1.AccountGroupVip
		err     error
	)

	if err := param.Valid().Struct(param); err != nil {
		err = pkg.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "登录参数校验失败",
			Message: err.Error(),
		}
		return account, err
	}
	if _, err := pkg.MyDatabase.Join("INNER", "beeQuick_vip_member", "beeQuick_vip_member.id = beeQuick_account.vip_member_id").Get(&account); err != nil {
		err = pkg.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "账号未注册",
			Message: err.Error(),
		}
		return account, err
	}
	if !common.CompareHashAndPassword([]byte(account.Password), []byte(param.Password)) {
		err = pkg.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "密码错误",
			Message: "password not correct",
		}
		return account, err
	}

	return account, nil
}

func signHandle(ctx iris.Context) {
	var param RegisterParam
	err := ctx.ReadJSON(&param)
	if err != nil {
		ctx.JSON(iris.Map(make_response.MakeResponse(http.StatusBadRequest, err, true)))
		return
	}
	account, err := signProcessor(param)
	if err != nil {
		ctx.JSON(iris.Map(make_response.MakeResponse(http.StatusBadRequest, err, true)))
		return
	}

	ctx.JSON(iris.Map(make_response.MakeResponse(http.StatusOK, account.SerializerForGroup(), false)))

}

func logoutHandle(ctx iris.Context) {
	account := ctx.Values().Get("current_admin")
	ctx.JSON(iris.Map(make_response.MakeResponse(http.StatusOK, account.(v1.Account).Serializer(), false)))
}

func getAccountProcessor(id uint) (v1.Account, error) {
	var (
		account v1.Account
		err     error
	)
	has, dbError := pkg.MyDatabase.Where("id = ?", id).Exist(&account)
	if dbError != nil || !has {
		err = pkg.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "记录未存在",
			Message: "Record not found",
		}
		return account, err
	} else {
		pkg.MyDatabase.Where("id = ?", id).Get(&account)
	}
	var vipMember v1.VipMember
	if account.VipMemberID != 0 {
		pkg.MyDatabase.Where("id = ?", account.VipMemberID).Get(&vipMember)
		account.VipMember = vipMember
	}

	return account, nil
}

func getAccountHandle(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")

	account, err := getAccountProcessor(id)
	if err != nil {
		ctx.JSON(iris.Map(make_response.MakeResponse(http.StatusBadRequest, err, true)))
		return
	}

	ctx.JSON(iris.Map(make_response.MakeResponse(http.StatusOK, account.Serializer(), false)))
}
