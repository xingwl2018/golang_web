package v1

import "time"

// Shop 商店表结构
type Shop struct {
	base `xorm:"extends"`
	// 商店位置信息
	Location   string `xorm:"varchar(255)" json:"location"`
	ProvinceId int64  `xorm:"index"`
	// 省市区表数据，`-`：表示不进行表结构字段映射
	Province Province `xorm:"-" json:"—"`
	// 商店名
	Name string `xorm:"varchar(64)"`
}

// TableName 设置表名
func (s Shop) TableName() string {
	return "xingwl_shop"
}

// ShopSerializer 用于前端或客户端调用接口展示的字段
type ShopSerializer struct {
	Id         int64              `json:"id"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	ProvinceId int64              `json:"province_id"`
	Province   ProvinceSerializer `json:"province"`
	Name       string             `json:"name"`
	Location   string             `json:"location"`
}

func (c Shop) Serializer() ShopSerializer {
	return ShopSerializer{
		Id:         int64(c.ID),
		CreatedAt:  c.CreatedAt.Truncate(time.Second),
		UpdatedAt:  c.UpdatedAt.Truncate(time.Second),
		Province:   c.Province.Serializer(),
		ProvinceId: c.ProvinceId,
		Name:       c.Name,
		Location:   c.Location,
	}
}
