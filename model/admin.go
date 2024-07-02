package model

import "gorm.io/gorm"

// 管理员
type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Avatar         string
}
