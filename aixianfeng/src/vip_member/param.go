package vip_member

import (
	v1 "aixianfeng/models/v1"
	"github.com/go-playground/validator/v10"
)

type PatchVipMemberParam struct {
	Level   string  `json:"level" validate:"contains=V"`
	Start   int     `json:"start" validate:"gte=0"`
	End     int     `json:"end" validate:"gte=0"`
	Period  int     `json:"period" validate:"gte=0"`
	ToValue int     `json:"to_value" validate:"gte=0"`
	Points  float64 `json:"points" validate:"gte=0"`
}

func (param PatchVipMemberParam) Valid() *validator.Validate {
	valid := validator.New()
	valid.RegisterStructValidation(param.validation, v1.VipMember{})
	return valid

}

func (param PatchVipMemberParam) validation(sl validator.StructLevel) {
	vipMember := sl.Current().Interface().(v1.VipMember)

	if vipMember.Start >= vipMember.End {
		sl.ReportError(vipMember.Start, "Start", "start", "start", "start should be less than end")
	}
}
