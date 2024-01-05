package v2

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// Exchange 兑换券结构体
type Exchange struct {
	gorm.Model
	Name     string    `gorm:"type:varchar" json:"name"`
	ZeroTime time.Time `gorm:"type:timestamp with time zone" json:"zero_time"`
	EndTime  time.Time `gorm:"type:timestamp with time zone" json:"end_time"`
	Price    sql.NullFloat64
}
