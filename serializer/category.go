package serializer

import "mywebmall/model"

type CategoryVO struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreateAt     int64  `json:"create_at"`
}

func BuildCategory(item *model.Category) CategoryVO {
	return CategoryVO{
		Id:           item.ID,
		CategoryName: item.CategoryName,
		CreateAt:     item.CreatedAt.Unix(),
	}
}

func BuildCategories(items []*model.Category) (categories []CategoryVO) {
	for _, item := range items {
		category := BuildCategory(item)
		categories = append(categories, category)
	}
	return categories
}
