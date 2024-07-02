package dao

import (
	"context"
	"mywebmall/model"

	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

func (dao *ProductDao) CreateProduct(product *model.Product) error {
	return dao.Model(&model.Product{}).Create(&product).Error
}

func (dao *ProductDao) UpdateProduct(pid uint, product *model.Product) error {
	return dao.Model(&model.Product{}).Where("id=?", pid).Updates(&product).Error
}

func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return total, err
}

// 根据条件展示商品
func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page model.BasePage) (products []*model.Product, err error) {
	// 根据page实行分页，offset跳过那些非当前页的商品。一页只展示pagesize大小的商品
	err = dao.DB.Where(condition).Limit(page.PageSize).Offset((page.PageNum - 1) * (page.PageSize)).Find(&products).Error
	return
}

// 搜索商品,这里使用了LIKE所以喜提全表扫描(难绷)
func (dao *ProductDao) SearchProduct(info string, page model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("title LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").
		Limit(page.PageSize).Offset((page.PageNum - 1) * (page.PageSize)).
		Find(&products).Error
	return
}

// 根据ID获取商品
func (dao *ProductDao) GetProductById(pid uint) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("id = ?", pid).First(&product).Error
	return
}
