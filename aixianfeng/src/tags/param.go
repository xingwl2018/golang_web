package tags

import (
	"aixianfeng/src/make_param"
	"github.com/go-playground/validator/v10"
)

type CreateTagParam struct {
	Name string `json:"name" validate:"required"`
}

func (c CreateTagParam) Valid() error {
	return validator.New().Struct(c)
}

type CreateTagsParam struct {
	Data []CreateTagParam `json:"data" validate:"required,dive,required"`
}

func (c CreateTagsParam) Valid() error {
	return validator.New().Struct(c)
}

type GetTagsParam struct {
	make_param.ReturnAll
}
