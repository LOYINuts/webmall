package api

import (
	"mywebmall/pkg/util"
	"mywebmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取某个商品的所有商品图片
func ListProductImg(c *gin.Context) {
	var listProductImgService service.ProductImgService
	if err := c.ShouldBind(&listProductImgService); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("list product img api error:", err)
	} else {
		res := listProductImgService.List(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	}
}
