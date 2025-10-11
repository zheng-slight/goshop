package frontend

//首页

import (
	"github.com/gin-gonic/gin"
	"goshop/models"
)

type IndexController struct {
	BaseController
}

func (con IndexController) Index(c *gin.Context) {
	//实例化redisCache结构体
	redisCache := models.RedisCache{}

	//获取顶部导航列表, 使用baseController.go中的

	//获取网站轮播图数据
	focusList := []models.Focus{}
	if hasFocusList := redisCache.Get("focusList", &focusList); !hasFocusList {
		models.DB.Where("status = 1 AND focus_type = 1").Find(&focusList)
		redisCache.Set("focusList", focusList, 3600)
	}

	//获取分类数据, 使用baseController.go中的
	//获取中间导航, 使用baseController.go中的
	//获取手机分类下面的商品
	phoneList := []models.Goods{}
	if hasPhoneList := redisCache.Get("phoneList", &phoneList); !hasPhoneList {
		phoneList := models.GetGoodsByCategory(23, "best", 10)
		redisCache.Set("phoneList", phoneList, 3600)
	}
	con.Render(c, "frontend/index/index.html", gin.H{
		"focusList": focusList,
		"phoneList": phoneList,
	})
}
