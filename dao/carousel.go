package dao

import (
	"context"
	"mywebmall/model"

	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

// 创建dao
func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

// 复用db创建dao
func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

// 获取所有轮播图
func (dao *CarouselDao) ListCarousel() (carousel []model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return
}
