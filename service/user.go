package service

import (
	"context"
	"mime/multipart"
	"mywebmall/conf"
	"mywebmall/dao"
	"mywebmall/model"
	"mywebmall/pkg/e"
	"mywebmall/pkg/util"
	"mywebmall/serializer"
	"time"

	"gopkg.in/gomail.v2"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` //密钥，前端验证
}

type SendEmailService struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	// 1:绑定邮箱 2:解绑邮箱 3:改密
	OperationType uint `json:"operation_type" form:"operation_type"`
}

type VarifyEmailService struct {
}

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

// 用户注册逻辑
func (service *UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success
	// 密钥为空或者密钥长度不对
	if service.Key == "" || len(service.Key) != 16 {
		code = e.Error
		util.LogrusObj.Infoln("UserService Register func error. key not long enough")
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度不足",
		}
	}
	// 密文存储，对称加密
	util.Encrypt.SetKey(service.Key)

	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	// 报错
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("UserService Register func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 用户名已经存在
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = model.User{
		UserName: service.UserName,
		NickName: service.NickName,
		Avatar:   "avatar.png",
		Status:   model.Active,
		Money:    util.Encrypt.AesEncoding("10000"), // 初始金额10000
	}
	// 密码加密
	if err := user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailedEncryption
		util.LogrusObj.Infoln("UserService Register func error.Password encryption error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 创建用户
	if err := userDao.CreateUser(&user); err != nil {
		code = e.Error
		util.LogrusObj.Infoln("UserService Register func error. In Sql create user error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	util.LogrusObj.Infoln("New User:", serializer.BuildUser(&user))
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 用户登录逻辑
func (service *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	// 检查用户是否存在
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("UserService Login func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 用户不存在
	if !exist {
		code = e.ErrorUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在，请先注册",
		}
	}
	// 密码校验
	ok := user.CheckPassword(service.Password)
	if !ok {
		code = e.ErrorPwNotMatch
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// token签发
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		util.LogrusObj.Infoln("UserService login func error. token error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
	}
}

// 用户信息更新
func (service *UserService) Update(ctx context.Context, uid uint) serializer.Response {
	var user *model.User
	var err error
	code := e.Success
	// 找到这个用户
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uid)
	// 修改昵称
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("UserService Update func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = userDao.UpdateUserById(uid, user)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("UserService Update func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// 用户头像上传
func (service *UserService) PostAvatar(ctx context.Context, uid uint, file multipart.File, filesize int64, fileHeader *multipart.FileHeader) serializer.Response {
	code := e.Success
	ok, fileType := CheckPhotoType(fileHeader)
	if !ok {
		code = e.ErrorFileType
		return serializer.Response{
			Status: e.ErrorFileType,
			Msg:    e.GetMsg(e.ErrorFileType),
			Data:   "上传文件类型错误!请只上传jpeg,jpg,png的图片!",
		}
	}
	var user *model.User
	var err error
	userDao := dao.NewUserDao(ctx)
	// 获取用户
	user, err = userDao.GetUserById(uid)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("UserService PostAvatar func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 保存图片到本地
	path, err := UploadAvatarToLocalStatic(file, uid, user.UserName, fileType)
	if err != nil {
		code = e.ErrorUploadFail
		util.LogrusObj.Infoln("UserService PostAvatar func error. Upload file failed:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Avatar = path
	// 数据库更新
	err = userDao.UpdateUserById(uid, user)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("UserService PostAvatar func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// 邮件发送
func (service *SendEmailService) Send(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	var address string
	var notice *model.Notice //绑定邮箱，修改密码使用模板通知
	// 使用用户名以及该邮件的操作类型等等来签发一个token
	token, err := util.GenerateEmailToken(uid, service.OperationType, service.Email, service.Password)
	if err != nil {
		code = e.ErrorAuthToken
		util.LogrusObj.Infoln("EmailService Send func error,token generate error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(service.OperationType)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("EmailService Send func error", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address = conf.ValidEmail + token //发送方
	mailStr := notice.Text
	mailText := mailStr + address
	m := gomail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "Web Mall")
	m.SetBody("text/html", mailText)
	// qq邮箱的规定端口为465
	d := gomail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	// 发送邮件
	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		util.LogrusObj.Infoln("EmailService Send func error,sending email error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 验证邮箱
func (service *VarifyEmailService) Varify(ctx context.Context, token string) serializer.Response {
	code := e.Success
	var userId uint
	var email string
	var password string
	var opType uint
	if token == "" {
		code = e.InvalidParams
	} else {
		claims, err := util.ParseEmailToken(token)
		if err != nil {
			code = e.ErrorAuthToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorTokenTimeout
		} else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			opType = claims.OperationType
		}
	}
	// 如果前面token解析等过程有错误这一步直接返回
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 更新到数据库
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("EmailService Varify func error", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if opType == 1 {
		// 绑定邮箱
		user.Email = email
	} else if opType == 2 {
		// 解除绑定邮箱
		user.Email = ""
	} else if opType == 3 {
		if err = user.SetPassword(password); err != nil {
			code = e.ErrorFailedEncryption
			util.LogrusObj.Infoln("EmailService Varify func error,password encryption error:", err)
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		code = e.ErrorEmailOPType
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("EmailService Varify func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// 显示用户金额
func (service *ShowMoneyService) Show(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildMoney(user, service.Key),
		Msg:    e.GetMsg(code),
	}
}
