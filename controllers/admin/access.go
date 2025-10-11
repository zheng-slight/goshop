package admin

import (
	"github.com/gin-gonic/gin"
	"goshop/models"
	"net/http"
	"strings"
)

type AccessController struct {
	BaseController
}

func (con AccessController) Index(c *gin.Context) {
	//获取权限列表
	accessList := []models.Access{}
	models.DB.Where("module_id = ?", 0).Preload("AccessItem").Find(&accessList)
	c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
		"accessList": accessList,
	})
}

func (con AccessController) Add(c *gin.Context) {
	//获取顶级模块
	accessList := []models.Access{}
	models.DB.Where("module_id = ?", 0).Find(&accessList)

	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": accessList,
	})
}

func (con AccessController) DoAdd(c *gin.Context) {
	//获取表单数据
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	actionName := strings.Trim(c.PostForm("action_name"), " ")
	accessType, err1 := models.Int(c.PostForm("type"))
	url := c.PostForm("url")
	moduleId, err2 := models.Int(c.PostForm("module_id"))
	sort, err3 := models.Int(c.PostForm("sort"))
	status, err4 := models.Int(c.PostForm("status"))
	description := strings.Trim(c.PostForm("description"), " ")
	//判断err
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		con.Error(c, "传入参数错误", "/admin/access/add")
		return
	}
	//判断moduleName
	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/add")
		return
	}
	//实例化access
	access := models.Access{
		ModuleName:  moduleName,
		ActionName:  actionName,
		Type:        accessType,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Status:      status,
		Description: description,
	}
	err5 := models.DB.Create(&access).Error
	if err5 != nil {
		con.Error(c, "添加权限失败", "/admin/access/add")
		return
	}
	con.Success(c, "添加权限成功", "/admin/access")
}

//编辑
func (con AccessController) Edit(c *gin.Context) {
	//获取id
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/access")
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)

	//获取顶级模块
	accessList := []models.Access{}
	models.DB.Where("module_id = ?", 0).Find(&accessList)

	c.HTML(http.StatusOK, "admin/access/edit.html", gin.H{
		"access":     access,
		"accessList": accessList,
	})
}

//编辑:提交
func (con AccessController) DoEdit(c *gin.Context) {
	//获取提交的表单数据
	id, err := models.Int(c.PostForm("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/access/edit?id"+models.String(id))
		return
	}
	//获取表单数据
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	actionName := strings.Trim(c.PostForm("action_name"), " ")
	accessType, err1 := models.Int(c.PostForm("type"))
	url := c.PostForm("url")
	moduleId, err2 := models.Int(c.PostForm("module_id"))
	sort, err3 := models.Int(c.PostForm("sort"))
	status, err4 := models.Int(c.PostForm("status"))
	description := strings.Trim(c.PostForm("description"), " ")
	//判断err
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		con.Error(c, "传入数据错误", "/admin/access/edit?id"+models.String(id))
		return
	}
	//判断moduleName
	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/edit?id"+models.String(id))
		return
	}

	//获取要修改的数据
	access := models.Access{Id: id}
	models.DB.Find(&access)
	access.ModuleName = moduleName
	access.ActionName = actionName
	access.Type = accessType
	access.Url = url
	access.ModuleId = moduleId
	access.Sort = sort
	access.Status = status
	access.Description = description
	//保存
	err5 := models.DB.Save(&access).Error
	if err5 != nil {
		con.Error(c, "编辑权限失败", "/admin/access/edit?id"+models.String(id))
		return
	}
	con.Success(c, "编辑权限成功", "/admin/access")
}

//删除
func (con AccessController) Delete(c *gin.Context) {
	//获取提交的表单数据
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/access")
		return
	}

	//获取要删除的数据
	access := models.Access{Id: id}
	models.DB.Find(&access)
	if access.ModuleId == 0 { // 顶级模块
		accessList := []models.Access{}
		models.DB.Where("module_id = ? ", access.Id).Find(&accessList)
		if len(accessList) > 0 {
			con.Error(c, "当前模块下子菜单,请先删除子菜单后再来删除这个数据", "/admin/access")
			return
		}
	}
	// 操作 或者 菜单, 或者顶级模块下面没有子菜单, 可以直接删除
	err = models.DB.Delete(&access).Error
	if err != nil {
		con.Error(c, "删除数据失败", "/admin/access")
		return
	}
	con.Success(c, "删除数据成功", "/admin/access")
}
