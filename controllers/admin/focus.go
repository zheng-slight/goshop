package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goshop/models"
	"net/http"
	"os"
	"strings"
)

type FocusController struct {
	BaseController
}

func (con FocusController) Index(c *gin.Context) {
	//获取轮播图列表
	focusList := []models.Focus{}
	models.DB.Find(&focusList)
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})
}

func (con FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}

func (con FocusController) DoAdd(c *gin.Context) {
	//获取请求的表单数据
	title := strings.Trim(c.PostForm("title"), " ")
	focusType, err1 := models.Int(c.PostForm("focus_type"))
	link := strings.Trim(c.PostForm("link"), " ")
	sort, err2 := models.Int(c.PostForm("sort"))
	status, err3 := models.Int(c.PostForm("status"))

	if err1 != nil || err3 != nil {
		con.Error(c, "非法请求", "/admin/focus/add")
		return
	}
	if err2 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/add")
		return
	}

	//文件上传操作
	focusImgSrc, err := models.UploadImg(c, "focus_img")
	if err != nil {
		con.Error(c, "图片上传失败", "/admin/focus/add")
		return
	}
	//实例化focus模型
	focus := models.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  focusImgSrc,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(models.GetUnix()),
	}
	err = models.DB.Create(&focus).Error
	if err != nil {
		con.Error(c, "增加轮播图失败", "/admin/focus/add")
		return
	}
	con.Success(c, "增加轮播图成功", "/admin/focus")
}

func (con FocusController) Edit(c *gin.Context) {
	//获取要修改的轮播图id
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/focus")
		return
	}
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)

	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus": focus,
	})
}

func (con FocusController) DoEdit(c *gin.Context) {
	//获取请求的表单数据
	id, err := models.Int(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	focusType, err1 := models.Int(c.PostForm("focus_type"))
	link := strings.Trim(c.PostForm("link"), " ")
	sort, err2 := models.Int(c.PostForm("sort"))
	status, err3 := models.Int(c.PostForm("status"))

	if err != nil || err1 != nil || err3 != nil {
		con.Error(c, "非法请求", "/admin/focus")
		return
	}
	if err2 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/edit?id="+models.String(id))
		return
	}
	//文件上传操作
	focusImgSrc, err11 := models.UploadImg(c, "focus_img")
	if err11 != nil {
		con.Error(c, "图片上传失败", "/admin/focus/edit?id="+models.String(id))
		return
	}

	focus := models.Focus{Id: id}
	models.DB.Find(&focus)

	focus.Title = title
	focus.FocusType = focusType
	focus.Status = status
	focus.Sort = sort
	focus.Link = link
	if focusImgSrc != "" {
		focus.FocusImg = focusImgSrc
	}
	err4 := models.DB.Save(&focus).Error
	if err4 != nil {
		con.Error(c, "修改轮播图失败", "/admin/focus/edit?id="+models.String(id))
		return
	}
	con.Success(c, "修改轮播图成功", "/admin/focus")

}

//删除
func (con FocusController) Delete(c *gin.Context) {
	//获取提交的表单数据
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/focus")
		return
	}

	//查询数据是否存在
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	//需要对文件进行删除吗?
	if focus.FocusImg != "" {
		fmt.Println(focus.FocusImg)
		os.Remove(focus.FocusImg)
	}
	err = models.DB.Delete(&focus).Error
	if err != nil {
		con.Error(c, "删除数据失败", "/admin/focus")
		return
	}

	con.Success(c, "删除数据成功", "/admin/focus")
}
