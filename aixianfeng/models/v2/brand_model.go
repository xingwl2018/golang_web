package v2

import "gorm.io/gorm"

// Brand 银行结构体
type Brand struct {
	gorm.Model
	EnName string `gorm:"type:varchar" json:"en_name"`
	ChName string `gorm:"type:varchar" json:"ch_name"`
}
