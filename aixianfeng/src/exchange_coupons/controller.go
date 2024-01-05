package exchange_coupons

import (
	v1 "aixianfeng/models/v1"
	"aixianfeng/pkg"
	"aixianfeng/src/common"
	"aixianfeng/src/make_response"
	"fmt"
	"github.com/kataras/iris/v12"
	"log"
	"net/http"
	"time"
)

// 核心控制逻辑

func getCouponsProcessor(param CouponsParam) ([]v1.ExchangeCoupon, error) {
	var (
		result []v1.ExchangeCoupon
	)
	if ok, err := param.Suitable(); !ok || err != nil {
		return result, pkg.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "请求参数不合法",
			Message: err.Error(),
		}
	}
	if dbError := pkg.MyDatabase.OrderBy("id desc").Find(&result); dbError != nil {
		return result, pkg.ErrorV1{
			Code:    http.StatusBadRequest,
			Detail:  "无记录",
			Message: dbError.Error(),
		}
	}
	return result, nil

}

func getCouponsHandler(ctx iris.Context) {
	var param CouponsParam
	log.Println("param", param, ctx.URLParam("return"))
	param.Return = ctx.URLParam("return")
	results, err := getCouponsProcessor(param)
	if err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err, true))
		return
	}
	var data []v1.ExchangeCouponSerializer
	if len(results) == 0 {
		ctx.JSON(make_response.MakeResponse(http.StatusOK, make([]v1.ExchangeCouponSerializer, 0), false))
		return
	}
	for _, i := range results {
		data = append(data, i.Serializer(""))
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, data, false))
}

func getCouponWithAccountHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("account_id")

	var account v1.Account
	if ok, _ := pkg.MyDatabase.Where("id = ?", id).Get(&account); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}

	var account2ExchangeCoupons []v1.Account2ExchangeCoupon

	if err := pkg.MyDatabase.Where("account_id = ?", account.ID).Find(&account2ExchangeCoupons); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}

	var ids []int
	for _, i := range account2ExchangeCoupons {
		ids = append(ids, int(i.ExchangeCouponId))
	}

	var exchangeCoupons []v1.ExchangeCoupon
	if dbError := pkg.MyDatabase.In("id", ids).Find(&exchangeCoupons); dbError != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, pkg.ErrorRecordNotFound, true))
		return
	}
	var results []v1.ExchangeCouponSerializer
	for _, i := range exchangeCoupons {
		results = append(results, i.Serializer(""))
	}
	ctx.JSON(make_response.MakeResponse(http.StatusOK, results, false))
}

func postCouponHandler(ctx iris.Context) {
	var param PostCouponParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err, true))
		return
	}

	valid := param.Valid()
	if err := valid.Struct(param); err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	var exchange v1.ExchangeCoupon
	var (
		start time.Time
		end   time.Time
	)
	if param.Start != "" {
		start, _ = common.ToTime(param.Start)
	}
	if param.End != "" {
		end, _ = common.ToTime(param.End)
	}
	exchange = v1.ExchangeCoupon{
		Name:  param.Name,
		Price: param.Price,
		Total: param.Total,
		Start: start,
		End:   end,
		Type:  param.Type,
		Token: param.Token,
	}
	session := pkg.MyDatabase.NewSession()
	session.Begin()

	if _, dbErr := session.Insert(&exchange); dbErr != nil {
		session.Rollback()
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, dbErr, true))
		return
	}
	session.Commit()
	ctx.JSON(make_response.MakeResponse(http.StatusOK, exchange.Serializer(""), false))
}

func patchCouponHandler(ctx iris.Context) {
	var param PatchCouponParam
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.ReadJSON(make_response.MakeResponse(http.StatusBadRequest, err, true))
		return
	}
	id, _ := ctx.Params().GetInt("coupon_id")
	fmt.Println(id)
	var exchangeCoupon v1.ExchangeCoupon
	if ok, _ := pkg.MyDatabase.Where("id = ?", id).Get(&exchangeCoupon); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, fmt.Errorf("record not found").Error(), true))
		return
	}
	if exchangeCoupon.ID == 0 {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, fmt.Errorf("record not found"), true))
		return
	}
	if param.Name != "" {
		exchangeCoupon.Name = param.Name
	}
	if param.Price != 0 && param.Total != 0 {
		exchangeCoupon.Price = param.Price
		exchangeCoupon.Total = param.Total
	}
	if param.Start != "" && param.End != "" {
		exchangeCoupon.Start, _ = common.ToTime(param.Start)
		exchangeCoupon.End, _ = common.ToTime(param.End)
	}
	pkg.MyDatabase.ID(exchangeCoupon.ID).Update(&exchangeCoupon)
	ctx.JSON(make_response.MakeResponse(http.StatusOK, exchangeCoupon.Serializer(""), false))
}

func postCouponToAccountHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("account_id")
	var param PostCouponToAccount
	err := ctx.ReadJSON(&param)
	if err != nil {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, err.Error(), true))
		return
	}
	var exchangeCoupon v1.ExchangeCoupon
	var account v1.Account

	if ok, _ := pkg.MyDatabase.Where("id = ?", param.ExchangeCouponId).Get(&exchangeCoupon); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, fmt.Errorf("record not found"), true))
		return
	}
	if ok, _ := pkg.MyDatabase.Where("id = ?", id).Get(&account); !ok {
		ctx.JSON(make_response.MakeResponse(http.StatusBadRequest, fmt.Errorf("record not found"), true))
		return
	}

	var account2ExchangeCoupon v1.Account2ExchangeCoupon
	account2ExchangeCoupon = v1.Account2ExchangeCoupon{
		ExchangeCouponId: int64(exchangeCoupon.ID),
		AccountId:        int64(account.ID),
		Status:           v1.NEW,
	}
	pkg.MyDatabase.InsertOne(&account2ExchangeCoupon)
	ctx.JSON(make_response.MakeResponse(http.StatusOK, exchangeCoupon.Serializer(""), false))
}
