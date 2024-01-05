package v1

import "time"

// Province 省市区的数据表结构
type Province struct {
	base `xorm:"extends"`
	//名称
	Name string `xorm:"varchar(10)" json:"name"`
	//地区编码
	AdCode string `xorm:"varchar(10)" json:"ad_code"`
	//城市编码
	CityCode string `xorm:"varchar(6)" json:"city_code"`
	//区中心点
	Center string `xorm:"varchar(32)" json:"center"`
	//行政区划级别
	Level string `xorm:"varchar(10)" json:"level"`
}

// 实现设置表名的方法
func (p Province) TableName() string {
	return "xingwl_province"
}

// ProvinceSerializer 序列化的结构体，从db转换到对外的结构体
type ProvinceSerializer struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	AdCode    string    `json:"ad_code"`
	CityCode  string    `json:"city_code"`
	Center    string    `json:"center"`
	Level     string    `json:"level"`
}

// Serializer 将数据表结构的数据序列化为
func (p Province) Serializer() ProvinceSerializer {
	return ProvinceSerializer{
		Id:        int(p.ID),
		CreatedAt: p.CreatedAt.Truncate(time.Second),
		UpdatedAt: p.UpdatedAt.Truncate(time.Second),
		Name:      p.Name,
		AdCode:    p.AdCode,
		Center:    p.Center,
		Level:     p.Level,
		CityCode:  p.CityCode,
	}
}
