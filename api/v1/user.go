package api

import (
	"mywebmall/pkg/util"
	"mywebmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	// 一定要记得shouldbind接收的是地址
	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user register api error:", err)
	} else {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	// 一定要记得shouldbind接收的是地址
	if err := c.ShouldBind(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user login api error:", err)
	} else {
		res := userLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

func UserUpdate(c *gin.Context) {
	var userUpdate service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	// 一定要记得shouldbind接收的是地址
	if err := c.ShouldBind(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("user update api error:", err)
	} else {
		res := userUpdate.Update(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}
}

func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	var uploadAvatar service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	// 一定要记得shouldbind接收的是地址
	if err := c.ShouldBind(&uploadAvatar); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("avatar upload api error:", err)
	} else {
		res := uploadAvatar.PostAvatar(c.Request.Context(), claims.ID, file, fileSize, fileHeader)
		c.JSON(http.StatusOK, res)
	}
}

func SendEmail(c *gin.Context) {
	var sendEmail service.SendEmailService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	// 一定要记得shouldbind接收的是地址
	if err := c.ShouldBind(&sendEmail); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("send email api error:", err)

	} else {
		res := sendEmail.Send(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}
}

func VarifyEmail(c *gin.Context) {
	var varifyEmail service.VarifyEmailService
	// 一定要记得shouldbind接收的是地址
	if err := c.ShouldBind(&varifyEmail); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("varify email error:", err)
	} else {
		res := varifyEmail.Varify(c.Request.Context(), c.GetHeader("Authorization"))
		c.JSON(http.StatusOK, res)
	}
}

func ShowMoney(c *gin.Context) {
	var showMoney service.ShowMoneyService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showMoney); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("show money api error:", err)
	} else {
		res := showMoney.Show(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}
}
