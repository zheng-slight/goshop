package api

import (
	"encoding/json"
	"goshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type V1Controller struct{}

// 获取导航列表
func (con V1Controller) Navlist(c *gin.Context) {
	navList := []models.Nav{}
	models.DB.Find(&navList)
	c.JSON(http.StatusOK, gin.H{
		"navList": navList,
	})
}

//后台通过c.PostForm获取数据
//api 当前端发送请求类型为:Content-Type: application/json，时,c.PostForm没法获取,需要通过c.GetRawData() 获取
//Content-Type: application/json; 发过来的数据需要通过c.GetRawData() 获取

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

// 登录操作:获取token
func (con V1Controller) Login(c *gin.Context) {
	tokenStr, err := models.SetToken(11)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "生成token失败重试",
			"success": false,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "获取token成功",
		"token":   tokenStr,
		"success": true,
	})
}

func (con V1Controller) DoLogin(c *gin.Context) {
	var userInfo UserInfo
	b, _ := c.GetRawData() //从 c.Request.Body 读取请求数据
	err := json.Unmarshal(b, &userInfo)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"userInfo": userInfo,
		})
	}
}

func (con V1Controller) DoLogin1(c *gin.Context) {
	userInfo := models.User{}
	//b, _ := c.GetRawData() //从 c.Request.Body 读取请求数据
	username := c.PostForm("username")

	//err := json.Unmarshal(b, &userInfo)
	if username == "" {
		c.JSON(200, gin.H{
			"err": 1,
		})
	} else {
		c.JSON(200, gin.H{
			"userInfo": userInfo,
			"username": username,
		})
	}
}

type Article struct {
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
}

// 编辑
func (con V1Controller) EditArticle(c *gin.Context) {
	var article Article
	b, _ := c.GetRawData() //从 c.Request.Body 读取请求数据
	err := json.Unmarshal(b, &article)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"article": article,
		})
	}
}

// 删除
func (con V1Controller) DeleteNav(c *gin.Context) {
	id := c.Query("id")
	c.JSON(200, gin.H{
		"message": "删除数据成功",
		"id":      id,
	})
}

// 获取收货地址
func (con V1Controller) AddressList(c *gin.Context) {
	//获取token
	tokenData := c.Request.Header.Get("Authorization")
	if len(tokenData) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "token传入错误长度不合法",
			"success": false,
		})
	}

	uid, err := models.GetToken(tokenData, 11)
	if err != nil { //校验失败
		c.JSON(http.StatusOK, gin.H{
			"message": err,
			"success": false,
		})
	}

	//校验成功
	c.JSON(http.StatusOK, gin.H{
		"uid":     uid,
		"success": true,
	})
}
