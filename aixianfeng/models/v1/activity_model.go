package v1

import "time"

// Activity 活动表结构
type Activity struct {
	base    `xorm:"extends"`
	Name    string    `xorm:"varchar(32)" json:"name"`
	Title   string    `xorm:"varchar(32)" json:"title"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	Avatar  string    `xorm:"varchar(255)" json:"avatar"`
	ShopIds []int     `xorm:"blob" json:"shop_ids"`
	Status  int       `xorm:"varchar(10)"`
}

func (a Activity) TableName() string {
	return "xingwl_activity"
}

type ActivitySerializer struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	Avatar    string    `json:"avatar"`
	ShopIds   []int     `json:"shop_ids"`
	Status    string    `json:"status"`
}

func (a Activity) Serializer() ActivitySerializer {
	return ActivitySerializer{
		Id:        a.ID,
		CreatedAt: a.CreatedAt.Truncate(time.Second),
		UpdatedAt: a.UpdatedAt.Truncate(time.Second),
		Name:      a.Name,
		Title:     a.Title,
		Start:     a.Start,
		End:       a.End,
		Avatar:    a.Avatar,
		ShopIds:   a.ShopIds,
		Status:    ActivityStatus[a.Status],
	}
}

// 定义常量表示：活动状态
const (
	// DOING iota定义序号：从零开始，后续的变量一次递增
	DOING = iota
	PROGRESSING
	CANCEL
	FINISH
	ADVANCE
)

// ActivityStatus 定义活动状态的中英文映射
var ActivityStatus = make(map[int]string)
var ActivityStatusEn = make(map[int]string)

// 初始化活动状态
func init() {
	ActivityStatus[DOING] = "未开始"
	ActivityStatus[PROGRESSING] = "进行中"
	ActivityStatus[CANCEL] = "取消"
	ActivityStatus[FINISH] = "结束"
	ActivityStatus[ADVANCE] = "提前"

	ActivityStatusEn[DOING] = "DOING"
	ActivityStatusEn[PROGRESSING] = "PROGRESSING"
	ActivityStatusEn[CANCEL] = "CANCEL"
	ActivityStatusEn[FINISH] = "FINISH"
	ActivityStatusEn[ADVANCE] = "ADVANCE"

}

// Activity2Product 各地区活动映射表结构：多对多关系
type Activity2Product struct {
	ProductId  int64 `xorm:"index"`
	ActivityId int64 `xorm:"index"`
}

func (s Activity2Product) TableName() string {
	return "xingwl_activity2Product"
}

// Shop2Activity 店铺活动映射表结构：多对多关系
type Shop2Activity struct {
	ShopId     int64 `xorm:"index"`
	ActivityId int64 `xorm:"index"`
}

func (s Shop2Activity) TableName() string {
	return "xingwl_shop2Activity"
}
