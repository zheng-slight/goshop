package admin

import (
	"github.com/gin-gonic/gin"
	"goshop/models"
	"net/http"
)

type SettingController struct {
	BaseController
}

func (con SettingController) Index(c *gin.Context) {
	setting := models.Setting{}
	models.DB.First(&setting)
	c.HTML(http.StatusOK, "admin/setting/index.html", gin.H{
		"setting": setting,
	})
}

func (con SettingController) DoEdit(c *gin.Context) {
	setting := models.Setting{Id: 1}
	models.DB.Find(&setting)
	//c.ShouldBind(&setting): 绑定表单中的数据到结构体
	if err := c.ShouldBind(&setting); err != nil {
		con.Error(c, "修改数据失败,请重试", "/admin/setting")
		return
	} else {
		// 上传图片 logo
		siteLogo, err1 := models.UploadImg(c, "site_logo")
		if len(siteLogo) > 0 && err1 == nil {
			setting.SiteLogo = siteLogo
		} else if err1 != nil {
			 con.Error(c, "上传错误", "/admin/setting")
			 return
		}
		//上传图片 no_picture
		noPicture, err2 := models.UploadImg(c, "no_picture")
		if len(noPicture) > 0 && err2 == nil {
			setting.NoPicture = noPicture
		}

		err3 := models.DB.Save(&setting).Error
		if err3 != nil {
			con.Error(c, "修改数据失败", "/admin/setting")
			return
		}

		con.Success(c, "修改数据成功", "/admin/setting")
	}
}
