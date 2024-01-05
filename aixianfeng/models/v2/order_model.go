package v2

import "gorm.io/gorm"

// 订单结构体
type Order struct {
	gorm.Model
	OrderStatus    []OrderStatus
	Status         string `gorm:"type:varchar" json:"status"`
	ShoppingCartID uint
}

// OrderStatus 订单状态
type OrderStatus struct {
	gorm.Model
	Product Product
	Amount  int `gorm:"type:integer" json:"amount"`
}
