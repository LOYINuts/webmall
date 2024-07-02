package model

import (
	"mywebmall/cache"
	"strconv"

	"gorm.io/gorm"
)

// 商品
type Product struct {
	gorm.Model
	Name          string
	CategoryId    uint
	Title         string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
}

// 生成商品被浏览的次数
func (product *Product) View() uint64 {
	// 在redis里面对每一个商品生成了一个key然后获取这个key的数据即可
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (product *Product) AddView() {
	// 增加商品点击数
	cache.RedisClient.Incr(cache.ProductViewKey(product.ID))
	// 在有序集合中对该商品的分数+1，因为要让浏览数最多的获得更多关注
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}
