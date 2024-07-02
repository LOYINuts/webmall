package dao

import (
	"context"
	"mywebmall/model"

	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) error {
	return dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
}

func (dao *ProductImgDao) ListProductImg(pid uint) (productImgs []*model.ProductImg, err error) {
	err = dao.DB.Model(&model.ProductImg{}).Where("product_id = ?", pid).Find(&productImgs).Error
	return
}
