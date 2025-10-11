package admin

//商品分类

import (
	"github.com/gin-gonic/gin"
	"goshop/models"
	"net/http"
	"strings"
)

type GoodsCateController struct {
	BaseController
}

func (con GoodsCateController) Index(c *gin.Context) {
	//定义一个切片
	goodsCateList := []models.GoodsCate{}
	//获取分类列表以及下级分类
	models.DB.Where("pid = ?", 0).Preload("GoodsCateItems").Find(&goodsCateList)
	c.HTML(http.StatusOK, "admin/goodsCate/index.html", gin.H{
		"goodsCateList": goodsCateList,
	})
}

//新增
func (con GoodsCateController) Add(c *gin.Context) {
	//获取商品分类顶级分类
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid = ?", 0).Find(&goodsCateList)

	c.HTML(http.StatusOK, "admin/goodsCate/add.html", gin.H{
		"goodsCateList": goodsCateList,
	})
}

//新增:提交
func (con GoodsCateController) DoAdd(c *gin.Context) {
	//获取请求的表单数据
	title := strings.Trim(c.PostForm("title"), " ")
	link := strings.Trim(c.PostForm("link"), " ")
	template := strings.Trim(c.PostForm("template"), " ")
	pid, err1 := models.Int(c.PostForm("pid"))
	subTitle := strings.Trim(c.PostForm("subTitle"), " ")
	keywords := strings.Trim(c.PostForm("keywords"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	sort, err2 := models.Int(c.PostForm("sort"))
	status, err3 := models.Int(c.PostForm("status"))

	if err1 != nil || err3 != nil {
		con.Error(c, "非法请求", "/admin/goodsCate/add")
		return
	}
	if err2 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/goodsCate/add")
		return
	}

	//文件上传操作
	imgSrc, err := models.UploadImg(c, "cate_img")
	if err != nil {
		con.Error(c, "图片上传失败", "/admin/goodsCate/add")
		return
	}
	//实例化GoodsCate模型
	goodsCate := models.GoodsCate{
		Title:       title,
		Link:        link,
		Sort:        sort,
		Status:      status,
		CateImg:     imgSrc,
		Template:    template,
		Pid:         pid,
		SubTitle:    subTitle,
		Keywords:    keywords,
		Description: description,
		AddTime:     int(models.GetUnix()),
	}
	err = models.DB.Create(&goodsCate).Error
	if err != nil {
		con.Error(c, "增加商品失败", "/admin/goodsCate/add")
		return
	}
	con.Success(c, "增加商品成功", "/admin/goodsCate")
}

//编辑
func (con GoodsCateController) Edit(c *gin.Context) {
	//获取角色id
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goodsCate")
		return
	}
	//获取商品分类顶级分类
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid = ?", 0).Find(&goodsCateList)

	//获取商品
	goodsCate := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate)
	c.HTML(http.StatusOK, "admin/goodsCate/edit.html", gin.H{
		"goodsCate":     goodsCate,
		"goodsCateList": goodsCateList,
	})
}

//编辑:提交
func (con GoodsCateController) DoEdit(c *gin.Context) {
	//获取提交的表单数据
	id, err := models.Int(c.PostForm("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goodsCate")
		return
	}
	//获取请求的表单数据
	title := strings.Trim(c.PostForm("title"), " ")
	link := strings.Trim(c.PostForm("link"), " ")
	template := strings.Trim(c.PostForm("template"), " ")
	pid, err1 := models.Int(c.PostForm("pid"))
	subTitle := strings.Trim(c.PostForm("subTitle"), " ")
	keywords := strings.Trim(c.PostForm("keywords"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	sort, err2 := models.Int(c.PostForm("sort"))
	status, err3 := models.Int(c.PostForm("status"))

	if err1 != nil || err3 != nil {
		con.Error(c, "非法请求", "/admin/goodsCate/add")
		return
	}
	if err2 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/goodsCate/add")
		return
	}

	//文件上传操作
	imgSrc, err := models.UploadImg(c, "cate_img")
	if err != nil {
		con.Error(c, "图片上传失败", "/admin/goodsCate/add")
		return
	}
	//查询分类是否存在
	goodsCate := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate)

	if imgSrc != "" {
		goodsCate.CateImg = imgSrc
	}
	goodsCate.Title = title
	goodsCate.Link = link
	goodsCate.Sort = sort
	goodsCate.Status = status
	goodsCate.Template = template
	goodsCate.Pid = pid
	goodsCate.SubTitle = subTitle
	goodsCate.Keywords = keywords
	goodsCate.Description = description
	err = models.DB.Save(&goodsCate).Error

	if err != nil {
		con.Error(c, "修改数据失败", "/admin/goodsCate/edit?id="+models.String(id))
		return
	}
	con.Success(c, "修改数据成功", "/admin/goodsCate")
}

//删除
func (con GoodsCateController) Delete(c *gin.Context) {
	//获取提交的表单数据
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goodsCate")
		return
	}

	//查询数据是否存在
	goodsCate := models.GoodsCate{Id: id}
	if goodsCate.Pid == 0 { // 顶级分类
		goodsCateList := []models.GoodsCate{}
		models.DB.Where("pid = ? ", goodsCate.Id).Find(&goodsCateList)
		if len(goodsCateList) > 0 {
			con.Error(c, "当前分类下存在子分类,请先删除子分类后再来删除这个数据", "/admin/goodsCate")
			return
		}
	}

	err = models.DB.Delete(&goodsCate).Error
	if err != nil {
		con.Error(c, "删除数据失败", "/admin/goodsCate")
		return
	}
	con.Success(c, "删除数据成功", "/admin/goodsCate")
}
