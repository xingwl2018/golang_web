package order

import (
	"aixianfeng/src/make_param"
	"github.com/go-playground/validator/v10"
)

type GetOrderParam struct {
	make_param.ReturnAll
}

type PatchOrderParam struct {
	Status string `json:"status" validate:"eq=readiness|eq=balance|eq=paid"`
}

func (p PatchOrderParam) Valid() error {
	return validator.New().Struct(p)
}

type PostOrderParam struct {
	ProductIds []int `json:"product_ids" validate:"required"`
	AccountId  int   `json:"account_id" validate:"required"`
}

func (p PostOrderParam) Valid() error {
	return validator.New().Struct(p)
}
