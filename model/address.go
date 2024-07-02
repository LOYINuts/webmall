package model

import (
	"gorm.io/gorm"
)

// 地址结构体:收货地址这些...
type Address struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	// 收货人名
	Name uint `gorm:"type:varchar(20) not null"`
	// 电话
	Phone string `gorm:"type:varchar(11) not null"`
	// 地址
	Address string `gorm:"type:varchar(50) not null"`
}
