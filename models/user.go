package models

//前端用户表相关

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type User struct { // 结构体首字母大写, 和数据库表名对应, 默认访问数据表users, 可以设置访问数据表的方法
	Id       int
	Phone    string `form:"phone" json:"phone"`       //手机号
	Password string `form:"password" json:"password"` // 密码
	AddTime  int    //增加时间
	LastIp   string // 最后登录ip
	Email    string // email
	Status   int    //状态
}

// 配置数据库操作的表名称
func (User) TableName() string {
	return "user"
}

const UserSessionKey = "userinfo_user"

// SetUserSession stores only the user id in the session and lets the session
// cookie carry the random session identifier.
func SetUserSession(c *gin.Context, id int) error {
	s := sessions.Default(c)
	s.Set(UserSessionKey, String(id))
	return s.Save()
}

// GetUserBySession loads the current user from the database by the id kept in
// the session, so no sensitive user fields are exposed to the client.
func GetUserBySession(c *gin.Context) (User, bool) {
	s := sessions.Default(c)
	v, ok := s.Get(UserSessionKey).(string)
	if !ok || v == "" {
		return User{}, false
	}

	id, err := Int(v)
	if err != nil {
		return User{}, false
	}

	var user User
	DB.Where("id = ?", id).First(&user)
	if user.Id == 0 {
		return User{}, false
	}
	return user, true
}

func ClearUserSession(c *gin.Context) {
	s := sessions.Default(c)
	s.Delete(UserSessionKey)
	_ = s.Save()
}
