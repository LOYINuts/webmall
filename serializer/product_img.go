package serializer

import (
	"mywebmall/conf"
	"mywebmall/model"
)

type ProductImgVO struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(item *model.ProductImg) ProductImgVO {
	return ProductImgVO{
		ProductId: item.ProductId,
		ImgPath:   conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
	}
}

func BuildProductImgs(items []*model.ProductImg) (productImgs []ProductImgVO) {
	for _, item := range items {
		productImg := BuildProductImg(item)
		productImgs = append(productImgs, productImg)
	}
	return productImgs
}
