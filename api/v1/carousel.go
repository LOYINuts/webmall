package api

import (
	"mywebmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCarousel(c *gin.Context) {
	var listCarousel service.CarouselService
	if err := c.ShouldBind(&listCarousel); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		res := listCarousel.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}
