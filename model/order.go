package model

import "gorm.io/gorm"

// 订单
type Order struct {
	gorm.Model
	UserId    uint
	ProductId uint
	BossId    uint
	AddressId uint
	Num       int
	OrderNum  uint64
	Type      uint //1未支付2已经支付
	Money     float64
}
