package v2

// Unit 度量单位结构体
type Unit struct {
	Name string `gorm:"type:varchar" json:"name"`
}
