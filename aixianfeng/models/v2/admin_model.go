package v2

import (
	"database/sql"
	"gorm.io/gorm"
)

// Admin 账户结构体
type Admin struct {
	gorm.Model
	AccountID       uint
	AccountBalance  sql.NullFloat64
	ExchangesNumber int `gorm:"type:integer" json:"exchanges_number"`
	CouponsNumber   int `gorm:"type:integer" json:"coupons_number"`
	Exchanges       []Exchange
	Coupons         []Coupon
}
