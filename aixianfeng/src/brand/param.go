package brand

import (
	"aixianfeng/src/make_param"
	"github.com/go-playground/validator/v10"
)

type CreateBrandParam struct {
	Name   string `json:"name" validate:"required_with_all"`
	EnName string `json:"en_name" validate:"required_with_all"`
}

func (c CreateBrandParam) Valid() error {
	return validator.New().Struct(c)
}

type PatchBrandParam struct {
	Name   string `json:"name"`
	EnName string `json:"en_name"`
}

type GetBrandParam struct {
	make_param.ReturnAll
}

type CreateBrandsParam struct {
	Data []CreateBrandParam `json:"data" validate:"required,dive,required"`
}

func (c CreateBrandsParam) Valid() error {
	return validator.New().Struct(c)
}
