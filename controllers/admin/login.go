package admin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goshop/models"
	"net/http"
)

type LoginController struct {
	BaseController
}

// 进入登录页面
func (con LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

// 执行登录操作
func (con LoginController) DoIndex(c *gin.Context) {
	//获取表单中的数据
	captchaId := c.PostForm("captchaId")     // 验证码id
	verifyValue := c.PostForm("verifyValue") //验证码的值
	//获取用户名以及密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 1.判断验证码是否验证成功
	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag {
		//2.查询数据库,判断用户以及密码是否正确
		manager := models.Manager{}
		models.DB.Where("username = ?", username).First(&manager)
		if manager.Username != "" && models.VerifyPassword(password, manager.Password) {
			//3.执行登录,只保存非敏感字段,执行跳转操作
			session := sessions.Default(c)
			adminSession, _ := json.Marshal(models.AdminSession{
				Id:       manager.Id,
				Username: manager.Username,
				RoleId:   manager.RoleId,
				IsSuper:  manager.IsSuper,
			})
			session.Set("userinfo_admin", string(adminSession))
			session.Save()
			con.Success(c, "登录成功", "/admin")
		} else {
			con.Error(c, "用户名或密码错误", "/admin/login")
		}
	} else {
		con.Error(c, "验证码验证失败", "/admin/login")
	}
}

// 获取验证码,验证验证码
func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := models.MakeCaptcha(50, 100, 1)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}

func (con LoginController) LoginOut(c *gin.Context) {
	//1.销毁session中用户信息
	session := sessions.Default(c)
	session.Delete("userinfo_admin")
	session.Save()
	con.Success(c, "退出登录成功", "/admin/login")
}
