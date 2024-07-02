package dao

import (
	"context"
	"mywebmall/model"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

// 创建dao
func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// 复用db创建dao
func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// 判断用户名是否存在，只有在不是记录未找到的错误的时候才返回err
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).First(&user).Error
	// 没有找到，不能用user == nil来判断因为user作为返回值已经进行了初始化
	// 是未找到错误则返回nil错误
	if err == gorm.ErrRecordNotFound {
		return nil, false, nil
	}
	// 如果其他错误则返回
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}

// 用户注册
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

// 根据ID获取用户
func (dao *UserDao) GetUserById(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
}

// 根据ID更新用户信息

func (dao *UserDao) UpdateUserById(uid uint, user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", uid).Updates(&user).Error
	return err
}
