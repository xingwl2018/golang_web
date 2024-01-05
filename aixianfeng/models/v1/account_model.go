package v1

import "time"

// 账户等级常量
const (
	MEMBER      = "会员"
	ADMIN       = "管理员"
	SUPEARADMIN = "超级管理员"
)

// Account 账户表结构模型
type Account struct {
	base `xorm:"extends"`
	//手机号
	Phone    string `xorm:"varchar(11) notnull unique 'phone'" json:"phone"`
	Password string `xorm:"varchar(128)" json:"password"`
	Token    string `xorm:"varchar(128) 'token'" json:"token"`
	Avatar   string `xorm:"varchar(128) 'avatar'" json:"avatar"`
	// 0 男 1 女
	Gender      string    `xorm:"varchar(1) 'gender'" json:"gender"`
	Birthday    time.Time `json:"birthday"`
	Points      int       `json:"points"`
	VipMemberID uint      `xorm:"index"`
	VipMember   VipMember `xorm:"-"`
	VipTime     time.Time `json:"vip_time"`
}

func (Account) TableName() string {
	return "xingwl_account"
}

type AccountSerializer struct {
	ID        uint                `json:"id"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	Phone     string              `json:"phone"`
	Password  string              `json:"-"`
	Token     string              `json:"token"`
	Avatar    string              `json:"avatar"`
	Gender    string              `json:"gender"`
	Age       int                 `json:"age"`
	Points    int                 `json:"points"`
	VipMember VipMemberSerializer `json:"vip_member"`
	VipTime   time.Time           `json:"vip_time"`
}

func (a Account) Serializer() AccountSerializer {

	// 转换性别
	gender := func() string {
		if a.Gender == "0" {
			return "男"
		}
		if a.Gender == "1" {
			return "女"
		}
		return a.Gender
	}
	// 计算年龄
	age := func() int {
		if a.Birthday.IsZero() {
			return 0
		}
		nowYear, _, _ := time.Now().Date()
		year, _, _ := a.Birthday.Date()
		if a.Birthday.After(time.Now()) {
			return 0
		}
		return nowYear - year
	}

	return AccountSerializer{
		ID:        a.ID,
		CreatedAt: a.CreatedAt.Truncate(time.Minute),
		UpdatedAt: a.UpdatedAt.Truncate(time.Minute),
		Phone:     a.Phone,
		Password:  a.Password,
		Token:     a.Token,
		Avatar:    a.Avatar,
		Points:    a.Points,
		Age:       age(),
		Gender:    gender(),
		VipTime:   a.VipTime.Truncate(time.Minute),
		VipMember: a.VipMember.Serializer(),
	}
}

// AccountGroupVip 账号会员关系表
type AccountGroupVip struct {
	Account   `xorm:"extends"`
	VipMember `xorm:"extends"`
}

func (AccountGroupVip) TableName() string {
	return "xingwl_account"
}
func (a AccountGroupVip) SerializerForGroup() AccountSerializer {
	result := a.Account.Serializer()
	result.VipMember = a.VipMember.Serializer()
	return result
}
