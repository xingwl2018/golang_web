package exchange_coupons

import (
	v1 "aixianfeng/models/v1"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type CouponsParam struct {
	Return string `json:"return"` // all_list, all_count, default: all_list
}

func (cop *CouponsParam) Suitable() (bool, error) {
	if cop.Return == "" {
		return false, fmt.Errorf("return should be all_list or all_count")
	}
	if !(cop.Return == "all_list" || cop.Return == "all_count") {
		cop.Return = "all_list"
		return true, nil
	}
	return true, nil
}

type PostCouponParam struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required,gt=0"`
	Total float64 `json:"total" validate:"required,gte=0"`
	Start string  `json:"start" validate:"required"`
	End   string  `json:"end" validate:"required"`
	Token string  `json:"token"`
	Type  int     `json:"type" validate:"eq=0|eq=1"` // 0,1 : 兑换券，优惠券
}

func (pcp PostCouponParam) Valid() *validator.Validate {
	valid := validator.New()
	valid.RegisterStructValidation(pcp.validation, v1.ExchangeCoupon{})
	return valid
}

func (pcp PostCouponParam) validation(sl validator.StructLevel) {
	exchangeCoupons := sl.Current().Interface().(v1.ExchangeCoupon)

	if exchangeCoupons.End.Before(exchangeCoupons.Start) {
		sl.ReportError(exchangeCoupons.Start, "Start", "start", "start", "start")
		sl.ReportError(exchangeCoupons.End, "End", "end", "end", "end")
	}

	if exchangeCoupons.Type == 1 && exchangeCoupons.Token == "" {
		sl.ReportError(exchangeCoupons.Token, "Token", "token", "token", "token")
		sl.ReportError(exchangeCoupons.Type, "Type", "type", "type", "type")
	}
}

type PatchCouponParam struct {
	Name  string  `json:"name"`
	Start string  `json:"start"`
	End   string  `json:"end"`
	Price float64 `json:"price"`
	Total float64 `json:"total"`
}

type PostCouponToAccount struct {
	ExchangeCouponId int `json:"exchange_coupon_id"`
}
