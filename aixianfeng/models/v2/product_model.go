package v2

import (
	"database/sql"
	"gorm.io/gorm"
)

// Product 商品模型
type Product struct {
	gorm.Model
	Name          string          `gorm:"type:varchar" json:"name"`
	Avatar        string          `gorm:"type:varchar" json:"avatar"`
	Price         sql.NullFloat64 `json:"price"`
	Amount        int             `gorm:"type:integer" json:"amount"`
	Specification string          `gorm:"type:varchar" json:"specification"`
	Period        int             `gorm:"type:integer" json:"period"`
	BrandID       uint
	UintID        uint
	TagID         uint
}
