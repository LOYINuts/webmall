package serializer

import (
	"mywebmall/model"
)

type CarouselVO struct {
	Id        uint   `json:"id"`
	ImagePath string `json:"image_path"`
	ProductId uint   `json:"product_id"`
	CreateAt  int64  `json:"create_at"`
}

// 单个
func BuildCarousel(item *model.Carousel) CarouselVO {
	return CarouselVO{
		Id:        item.ID,
		ImagePath: item.ImgPath,
		ProductId: item.ProductId,
		CreateAt:  item.CreatedAt.Unix(),
	}
}

// 多个
func BuildCarousels(items []model.Carousel) (carousels []CarouselVO) {
	for _, item := range items {
		c := BuildCarousel(&item)
		carousels = append(carousels, c)
	}
	return carousels
}
