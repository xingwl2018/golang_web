package v2

import "time"

// Tag 标签模型
type Tag struct {
	//gorm.Model
	ID        uint       `xorm:"pk 'id'"`
	CreatedAt time.Time  `xorm:"created"`
	UpdatedAt time.Time  `xorm:"updated"`
	DeletedAt *time.Time `xorm:"deleted index"`
	Name      string     `xorm:"type:varchar" json:"name"`
}
