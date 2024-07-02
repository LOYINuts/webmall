package model

import "gorm.io/gorm"

// 分类
type Category struct {
	gorm.Model
	CategoryName string
}
