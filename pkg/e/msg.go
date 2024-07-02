package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "参数错误",
	ErrorDataBase: "数据库错误",
	// 用户模块
	ErrorExistUser:        "用户名已存在!",
	ErrorFailedEncryption: "加密错误!",
	ErrorUserNotFound:     "找不到用户名",
	ErrorPwNotMatch:       "密码错误",
	ErrorAuthToken:        "Token认证失败",
	ErrorTokenTimeout:     "token过期",
	ErrorUploadFail:       "文件上传失败",
	ErrorSendEmail:        "邮件发送失败",
	ErrorEmailOPType:      "邮件操作码错误",
	ErrorFileType:         "文件类型错误",
	// 轮播图模块
	ErrorGetCarousel: "获取轮播图错误",
	// 商品模块
	ErrorProductImgUpload: "商品图片上传错误",
}

// 获取状态码对应的信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
