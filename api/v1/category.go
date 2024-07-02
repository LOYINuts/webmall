package api

import (
	"mywebmall/pkg/util"
	"mywebmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建分类
func CreateCategory(c *gin.Context) {
	var createCategoryService service.CategoryService
	if err := c.ShouldBind(&createCategoryService); err != nil {
		util.LogrusObj.Infoln("Create Category api error:", err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	} else {
		res := createCategoryService.Create(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

// 获取所有分类
func ListCategory(c *gin.Context) {
	var listCategoryService service.CategoryService
	if err := c.ShouldBind(&listCategoryService); err != nil {
		util.LogrusObj.Infoln("List Category api error:", err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	} else {
		res := listCategoryService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}
