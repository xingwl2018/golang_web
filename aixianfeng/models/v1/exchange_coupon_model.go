package v1

import "time"

// 券类别常量
const (
	// 兑换券，优惠券
	EXCHANGE = iota
	COUPON
)

// 券状态常量
const (
	// 未使用、已使用、已过期
	NEW = iota
	USED
	EXPIRE
)

var CouponType = make(map[int]string)
var StatusMap = make(map[int]string)

func init() {
	StatusMap[NEW] = "未使用"
	StatusMap[USED] = "已使用"
	StatusMap[EXPIRE] = "已过期"

	CouponType[EXCHANGE] = "兑换券"
	CouponType[COUPON] = "优惠券"
}

// ExchangeCoupon 券表结构模型
type ExchangeCoupon struct {
	base  `xorm:"extends"`
	Name  string    `xorm:"varchar(32) unique" json:"name"`
	Price float64   `json:"price"`
	Total float64   `json:"total"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Token string    `json:"token"`
	Type  int       `json:"type"` // 0,1 : 兑换券 抵消价格，优惠券 类似几折
}

func (exchange ExchangeCoupon) TableName() string {
	return "xingwl_exchange_coupons"
}

type ExchangeCouponSerializer struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Total     float64   `json:"total"`
	Start     string    `json:"start"` //  格式：2006/01/02
	End       string    `json:"end"`   // 格式：2006/01/02
	Status    string    `json:"status"`
	// 0，1:兑换券抵消价格，优惠券类似几折
	Type string `json:"type"`
}

func (exchange ExchangeCoupon) Serializer(status string) ExchangeCouponSerializer {

	return ExchangeCouponSerializer{
		ID:        exchange.ID,
		CreatedAt: exchange.CreatedAt.Truncate(time.Second),
		UpdatedAt: exchange.UpdatedAt.Truncate(time.Second),
		Name:      exchange.Name,
		Price:     exchange.Price,
		Total:     exchange.Total,
		Start:     exchange.Start.Format("2006-01-02 15:04:05"),
		End:       exchange.End.Format("2006-01-02 15:04:05"),
		Status:    status,
		Type:      CouponType[exchange.Type],
	}
}

// Account2ExchangeCoupon 账号优惠券关联关系表
type Account2ExchangeCoupon struct {
	AccountId        int64 `xorm:"index"`
	ExchangeCouponId int64 `xorm:"index"`
	Status           int   `json:"status"` // 0,1,2:未使用，已使用，已过期
}

func (a2e Account2ExchangeCoupon) TableName() string {
	return "xingwl_account2exchange_coupon"
}

// RuleForExchangeOrCoupon 会员规则表模型
type RuleForExchangeOrCoupon struct {
	base     `xorm:"extends"`
	Question string `xorm:"unique"`
	Answer   string
	Type     int
}

func (RuleForExchangeOrCoupon) TableName() string {
	return "xingwl_rule_coupon"
}

type RuleForExchangeOrCouponSerializer struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	Type      string    `json:"type"`
}

func (r RuleForExchangeOrCoupon) Serializer() RuleForExchangeOrCouponSerializer {
	return RuleForExchangeOrCouponSerializer{
		Id:        r.ID,
		CreatedAt: r.CreatedAt.Truncate(time.Second),
		UpdatedAt: r.UpdatedAt.Truncate(time.Second),
		Question:  r.Question,
		Answer:    r.Answer,
		Type:      CouponType[r.Type],
	}
}
