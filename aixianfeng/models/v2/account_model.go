package v2

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// Account 用户结构体
type Account struct {
	gorm.Model
	LevelID uint
	Phone   string `gorm:"type:varchar" json:"phone"`
	Avatar  string `gorm:"type:varchar" json:"avatar"`
	Name    string `gorm:"type:varchar" json:"name"`
	// 0 男 1 女
	Gender   int       `gorm:"type:integer" json:"gender"`
	Birthday time.Time `gorm:"type:timestamp with time zone" json:"birthday"`
	Points   sql.NullFloat64
}
