package dao

import (
	"fmt"
	"mywebmall/model"
)

// 对model进行迁移即建表
func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.Address{},
		&model.Admin{},
		&model.Carousel{},
		&model.Cart{},
		&model.Category{},
		&model.Favorite{},
		&model.Notice{},
		&model.Order{},
		&model.ProductImg{},
		&model.Product{},
		&model.User{},
	)
	if err != nil {
		fmt.Printf("error:%v\n", err)
	}
}
