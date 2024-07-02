package routers

import (
	"mywebmall/api/v1"
	"mywebmall/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
		// 用户操作
		v1.POST("user/register", api.UserRegister) //注册用户
		v1.POST("user/login", api.UserLogin)       //登录用户

		// 轮播图
		v1.GET("carousels", api.ListCarousel)

		// 商品操作
		// 获取商品列表,这里根据传入是否有category_id则实现了展示那个分类下的所有商品
		v1.GET("products", api.ListProduct)
		// 搜索商品
		v1.POST("products/search", api.SearchProduct)
		// 获取特定商品信息
		v1.GET("products/:id", api.ShowProduct)
		// 获取商品图片
		v1.GET("products/:id/img", api.ListProductImg)
		// 获取分类
		v1.GET("category", api.ListCategory)
		// 创建分类
		v1.POST("category", api.CreateCategory)

		authed := v1.Group("/") // 需要登录保护的组
		authed.Use(middleware.JWT())
		{
			// 用户操作
			// 信息更新(NickName)
			authed.PUT("user", api.UserUpdate)
			// 头像更新
			authed.PUT("avatar", api.UploadAvatar)
			// 邮件发送
			authed.POST("user/sending-email", api.SendEmail)
			// 邮件验证
			authed.POST("user/valid-email", api.VarifyEmail)

			// 显示金额
			authed.POST("money", api.ShowMoney)

			// 商品操作
			// 创建商品
			authed.POST("product", api.CreateProduct)
		}
	}
	return r
}
