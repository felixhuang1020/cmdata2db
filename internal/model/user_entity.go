package model

type User struct {
	Id          int64  `gorm:"primaryKey"`
	UserID      int64  `gorm:"size:50;not null"`
	Password    string `gorm:"size:50;not null"`
	UserName    string `gorm:"size:30"`
	Email       string `gorm:"size:50'"`
	PhoneNumber string `gorm:"size:20"`
	Sex         string `gorm:"size:10"`
	Remark      string `gorm:"size:500"`
}

type Product struct {
	ID            string  `json:"id"`             // 产品唯一标识（如SKU、UUID）
	Name          string  `json:"name"`           // 产品名称
	Description   string  `json:"description"`    // 产品描述
	Category      string  `json:"category"`       // 主分类（如"电子产品"）
	SubCategory   string  `json:"sub_category"`   // 子分类（如"智能手机"）
	Brand         string  `json:"brand"`          // 品牌
	Price         float64 `json:"price"`          // 售价
	OriginalPrice float64 `json:"original_price"` // 原价（用于折扣计算）
	Currency      string  `json:"currency"`       // 货币单位（默认"CNY"）
	Stock         int     `json:"stock"`          // 库存量
	IsInStock     bool    `json:"is_in_stock"`
}

// TableName 方法用于返回表名
func (u User) TableName() string {
	return "t_user"
}
