package middlewares

//中间件: 作用: 在执行路由之前或者之后进行相关逻辑判断

import (
	"goshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUserAuthMiddleware(c *gin.Context) {
	//获取Cookie里面保存的用户信息
	user := models.User{}
	isLogin := models.Cookie.Get(c, "user", &user)
	if !isLogin || len(user.Phone) != 11 { // 用户没有登录或者账号不正确, 跳转到登录页面
		c.Redirect(http.StatusFound, "/pass/login")
		return
	}
}
