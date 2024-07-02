package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string
	Money          string //密文所以用字符串
}

const (
	PasswordCost int    = 12 //密码加密难度
	Active       string = "active"
)

// 设置用户密码并加密
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	// 根据函数介绍可知如果相等则err为nil，否则err有错误
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}
