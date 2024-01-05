package province

import (
	"github.com/go-playground/validator/v10"
)

type GetProvinceParam struct {
	Level  string `json:"level" validate:"eq=province|eq=city|eq=district"`
	Return string `json:"return" validate:"eq=all_list|eq=all_count"`
}

func (g GetProvinceParam) Valid() error {
	return validator.New().Struct(g)
}
