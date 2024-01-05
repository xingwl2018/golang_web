package make_param

import "github.com/go-playground/validator/v10"

type ReturnAll struct {
	ReturnAll string `json:"return_all" validate:"eq=all_count|eq=all_list"`
}

func (r ReturnAll) Valid() error {
	return validator.New().Struct(r)
}
