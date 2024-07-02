package e

import "net/http"

// 定义的自己的错误代码，其实可以直接用http给好的status但是自己定义更有扩展性

const (
	Success       = http.StatusOK                  //成功，200
	Error         = http.StatusInternalServerError //错误，500
	InvalidParams = http.StatusBadRequest          // 参数错误，400
	ErrorDataBase = 666                            // 数据库错误
	// user模块错误
	ErrorExistUser        = 9001 //用户名已经存在
	ErrorFailedEncryption = 9002 //加密错误
	ErrorUserNotFound     = 9003 //找不到用户
	ErrorPwNotMatch       = 9004 //密码错误
	ErrorAuthToken        = 9005 //Token错误
	ErrorTokenTimeout     = 9006 //token过期
	ErrorUploadFail       = 9007 //文件上传失败
	ErrorSendEmail        = 9008 //邮件发送失败
	ErrorEmailOPType      = 9009 //邮件操作码错误
	ErrorFileType         = 9010 //文件类型错误

	// Carousel模块错误
	ErrorGetCarousel = 8001 //获取轮播图错误

	// product模块错误
	ErrorProductImgUpload = 7001 //商品图片上传错误
)
