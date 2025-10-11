package api

import (
	"encoding/json"
	"goshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type V2Controller struct{}

// 获取导航列表
func (con V2Controller) Navlist(c *gin.Context) {
	navList := []models.Nav{}
	models.DB.Find(&navList)
	c.JSON(http.StatusOK, gin.H{
		"navList": navList,
	})
}

//后台通过c.PostForm获取数据
//api 当前端发送请求类型为:Content-Type: application/json，时,c.PostForm没法获取,需要通过c.GetRawData() 获取
//Content-Type: application/json; 发过来的数据需要通过c.GetRawData() 获取

func (con V2Controller) DoLogin(c *gin.Context) {
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

// 编辑
func (con V2Controller) EditArticle(c *gin.Context) {
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
func (con V2Controller) DeleteNav(c *gin.Context) {
	id := c.Query("id")
	c.JSON(200, gin.H{
		"message": "删除数据成功",
		"id":      id,
	})
}
