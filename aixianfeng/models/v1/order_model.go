package v1

import "time"

// 定义常量：订单状态
const (
	// 准备状态、未付款状态、已付款状态
	READINESS = iota
	BALANCE
	PAID
)

// 订单状态名称映射
var (
	STATUS_MAP    = make(map[int]string)
	STATUS_MAP_EN = make(map[int]string)
)

// 初始化订单状态：映射信息
func init() {
	STATUS_MAP[READINESS] = "准备状态"
	STATUS_MAP[BALANCE] = "未付款状态"
	STATUS_MAP[PAID] = "已付款状态"
	STATUS_MAP_EN[READINESS] = "readiness"
	STATUS_MAP_EN[BALANCE] = "balance"
	STATUS_MAP_EN[PAID] = "paid"
}

// Order 订单表结构模型
type Order struct {
	base `xorm:"extends"`
	// 一个订单和多个商品进行关联
	ProductIds []int `xorm:"blob"`
	Status     int
	AccountId  int64
	Account    Account `xorm:"-"`
	Total      float64
}

func (o Order) TableName() string {
	return "xingwl_order"
}

type OrderSerializer struct {
	Id         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Status     string    `json:"status"`
	Phone      string    `json:"phone"`
	AccountId  uint      `json:"account_id"`
	Total      float64   `json:"total"`
	ProductIds []int     `json:"product_ids"`
}

func (o Order) Serializer() OrderSerializer {
	return OrderSerializer{
		Id:         o.ID,
		CreatedAt:  o.CreatedAt.Truncate(time.Second),
		UpdatedAt:  o.UpdatedAt.Truncate(time.Second),
		Status:     STATUS_MAP[o.Status],
		AccountId:  o.Account.ID,
		Phone:      o.Account.Phone,
		Total:      o.Total,
		ProductIds: o.ProductIds,
	}
}
