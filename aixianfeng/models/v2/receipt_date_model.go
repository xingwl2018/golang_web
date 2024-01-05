package v2

import (
	"gorm.io/gorm"
	"time"
)

// ReceiptDate 订单接收日期
type ReceiptDate struct {
	gorm.Model
	// 收货时间戳
	ReceiveDateID uint
	ReceiveDate   ReceiveDate
	FormTime      time.Time `gorm:"type:timestamp with time zone" json:"form_time"`
	ToTime        time.Time `gorm:"type:timestamp with time zone" json:"to_time"`
}

type ReceiveDate struct {
	gorm.Model
	Item string `gorm:"type:varchar" json:"item"`
}
