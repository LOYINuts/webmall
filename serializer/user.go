package serializer

import (
	"mywebmall/conf"
	"mywebmall/model"
)

// VO View Object 视图对象，传给前端看的
type UserVO struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

func BuildUser(user *model.User) *UserVO {
	return &UserVO{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}
