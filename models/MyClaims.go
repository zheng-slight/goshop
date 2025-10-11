package models

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

//定义key和过期时间
var jwtKey = []byte("www.xxx.comx") //byte类型的切片
var expireTime = time.Now().Add(24 * time.Hour).Unix()

//自定义一个结构体，这个结构体需要继承 jwt.StandardClaims 结构体
type MyClaims struct {
	Uid int //自定义的属性 用于不同接口传值
	jwt.StandardClaims
}

//设置token
func SetToken(uid int) (string, error) {
	//实例化 存储token的结构体
	myClaimsObj := MyClaims{
		uid, //自定义参数: 可自行传值
		jwt.StandardClaims{
			ExpiresAt: expireTime, //过期时间
			Issuer:    "www.xxx.com",
		},
	}

	// 使用指定的签名方法创建签名对象
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaimsObj)
	// 使用指定的 secret 签名并获得完整的编码后的字符串 token
	tokenStr, err := tokenObj.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func GetToken(tokenData string, uid int) (int, error) {
	//获取token
	tokenStr := strings.Split(tokenData, " ")[1]
	//校验token
	token, myClaims, err := ParseToken(tokenStr)

	if err != nil || !token.Valid { //校验失败
		return 0, err
	} else {
		return myClaims.Uid, nil
	}
}

//验证token是否合法
func ParseToken(tokenStr string) (*jwt.Token, *MyClaims, error) {
	myClaims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, myClaims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, myClaims, err
}
