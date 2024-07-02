package service

import (
	"context"
	"mywebmall/dao"
	"mywebmall/pkg/e"
	"mywebmall/pkg/util"
	"mywebmall/serializer"
)

type CarouselService struct {
}

func (service *CarouselService) List(ctx context.Context) serializer.Response {
	carouselDao := dao.NewCarouselDao(ctx)
	code := e.Success
	carousels, err := carouselDao.ListCarousel()
	if err != nil {
		code = e.ErrorGetCarousel
		util.LogrusObj.Infoln("CarouselService list func error", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}
