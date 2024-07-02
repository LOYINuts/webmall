package service

import (
	"context"
	"mywebmall/dao"
	"mywebmall/model"
	"mywebmall/pkg/e"
	"mywebmall/pkg/util"
	"mywebmall/serializer"
)

type CategoryService struct {
	CategoryName string `json:"category_name" form:"category_name"`
}

func (service *CategoryService) List(ctx context.Context) serializer.Response {
	code := e.Success
	categoryDao := dao.NewCategoryDao(ctx)
	categories, err := categoryDao.ListCategories()
	if err != nil {
		code = e.ErrorDataBase
		util.LogrusObj.Infoln("CategoryService list func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategories(categories), uint(len(categories)))
}

func (service *CategoryService) Create(ctx context.Context) serializer.Response {
	code := e.Success
	categoryDao := dao.NewCategoryDao(ctx)
	category := &model.Category{
		CategoryName: service.CategoryName,
	}
	if err := categoryDao.CreateCategory(category); err != nil {
		code = e.ErrorDataBase
		util.LogrusObj.Infoln("CategoryService Create func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCategory(category),
	}
}
