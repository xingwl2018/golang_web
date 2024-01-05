package activity

import (
	"aixianfeng/src/common"
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type CreateActivityParam struct {
	Name    string `json:"name" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Start   string `json:"start" validate:"required"`
	End     string `json:"end" validate:"required"`
	Avatar  string `json:"avatar"`
	ShopIds []int  `json:"shop_ids" validate:"required"`
}

func (c CreateActivityParam) Valid() error {
	return validator.New().Struct(c)
}

func (c CreateActivityParam) timeHandle() (time.Time, time.Time, error) {
	start, _ := common.ToTime(c.Start)
	end, _ := common.ToTime(c.End)
	if end.Before(start) {
		return time.Time{}, time.Time{}, fmt.Errorf("end before start")
	}
	return start, end, nil
}

type PatchActivityParam struct {
	Name    string `json:"name" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Start   string `json:"start" validate:"required"`
	End     string `json:"end" validate:"required"`
	Avatar  string `json:"avatar"`
	ShopIds []int  `json:"shop_ids" validate:"required"`
	Status  int    `json:"status" validate:"eq=0|eq=1|eq=2|eq=3|eq=4"`
}

func (p PatchActivityParam) Valid() error {
	return validator.New().Struct(p)
}

type GetActivityParam struct {
	Status    string `json:"status" validate:"eq=doing|eq=progressing|eq=cancel|eq=finish|eq=advance"`
	ReturnAll string `json:"return" validate:"eq=all_list|eq=all_count"`
}

func (g GetActivityParam) Valid() error {
	return validator.New().Struct(g)
}
