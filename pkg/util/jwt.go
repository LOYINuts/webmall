package util

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("LOYInuts")

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"` //权限
	jwt.StandardClaims
}

// 签发token
func GenerateToken(id uint, userName string, authority int) (string, error) {
	nowTime := time.Now()                     // 签发时间
	expireTime := nowTime.Add(24 * time.Hour) //token过期时间
	claims := Claims{
		ID:        id,
		UserName:  userName,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "LOYInuts'Mall", //token发行人
		},
	}
	// 使用hash256加密
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 验证用户token
func ParseToken(token string) (*Claims, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"json"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

// 签发email token
func GenerateEmailToken(userId, op uint, email, password string) (string, error) {
	nowTime := time.Now()                     // 签发时间
	expireTime := nowTime.Add(24 * time.Hour) //token过期时间
	claims := EmailClaims{
		UserID:        userId,
		Email:         email,
		Password:      password,
		OperationType: op,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "LOYInuts'Mall", //token发行人
		},
	}
	// 使用hash256加密
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 验证邮箱token
func ParseEmailToken(token string) (*EmailClaims, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
