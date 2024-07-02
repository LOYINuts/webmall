package api

import (
	"mywebmall/pkg/util"
	"mywebmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	// 获取是谁在创建商品
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	var cpService service.ProductService
	if err := c.ShouldBind(&cpService); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("create product api error:", err)
	} else {
		res := cpService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(http.StatusOK, res)
	}
}

// 获取所有商品
func ListProduct(c *gin.Context) {
	var listProductService service.ProductService
	if err := c.ShouldBind(&listProductService); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("list product api error:", err)
	} else {
		res := listProductService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

// 搜索商品
func SearchProduct(c *gin.Context) {
	var searchProductService service.ProductService
	if err := c.ShouldBind(&searchProductService); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("search product api error:", err)
	} else {
		res := searchProductService.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

// 获取某一商品的信息
func ShowProduct(c *gin.Context) {
	var showProductService service.ProductService
	if err := c.ShouldBind(&showProductService); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("show product api error:", err)
	} else {
		res := showProductService.Show(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	}
}
