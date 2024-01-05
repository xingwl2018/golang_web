package v2

import "gorm.io/gorm"

// Coupon 优惠券结构体
type Coupon struct {
	gorm.Model
	Exchange
	Token string `gorm:"type:varchar" json:"token"`
}
