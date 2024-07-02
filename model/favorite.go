package model

import (
	"gorm.io/gorm"
)

// 收藏夹
type Favorite struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:UserId"`
	UserId    uint    `gorm:"not null"`
	Product   Product `gorm:"ForeignKey:ProductId"`
	ProductId uint    `gorm:"ForeignKey:UserId"`
	Boss      User    `gorm:"ForeignKey:BossId"`
	BossId    uint    `gorm:"not null"`
}
