package service

import (
	"context"
	"mywebmall/dao"
	"mywebmall/pkg/e"
	"mywebmall/pkg/util"
	"mywebmall/serializer"
	"strconv"
)

type ProductImgService struct {
}

func (service *ProductImgService) List(ctx context.Context, id string) serializer.Response {
	pid, _ := strconv.Atoi(id)
	code := e.Success
	productImgDao := dao.NewProductImgDao(ctx)
	productImgs, err := productImgDao.ListProductImg(uint(pid))
	if err != nil {
		code = e.ErrorDataBase
		util.LogrusObj.Infoln("ProductImgService List func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildProductImgs(productImgs), uint(len(productImgs)))
}
