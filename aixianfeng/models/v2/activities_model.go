package v2

import (
	"gorm.io/gorm"
	"time"
)

// Activities 活动模型
type Activities struct {
	gorm.Model
	Title    string    `gorm:"type:varchar" json:"title"`
	FromDate time.Time `gorm:"type:timestamp with time zone" json:"from_date"`
	ToDate   time.Time `gorm:"type:timestamp with time zone" json:"to_date"`
	// 商品模型
	Products []Product `gorm:"type:many2many: activity2products" json:"products"`
}

// TableName 自定义表命
func (a Activities) TableName() string {
	return "activities"
}
