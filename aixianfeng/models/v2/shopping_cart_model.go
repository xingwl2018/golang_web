package v2

import "gorm.io/gorm"

// ShoppingCart 购物清单
type ShoppingCart struct {
	gorm.Model
	AccountID     uint
	ReceiptDateID uint
	OrderID       uint
	Order         Order
}
