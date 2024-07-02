package dao

import (
	"context"
	"mywebmall/model"

	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

// 创建dao
func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

// 复用db创建dao
func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// 根据id获取notice
func (dao *NoticeDao) GetNoticeById(id uint) (notice *model.Notice, err error) {
	err = dao.DB.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	return
}
