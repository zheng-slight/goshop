package admin

import (
	"github.com/gin-gonic/gin"
	"goshop/models"
	"net/http"
	"strings"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	//获取管理员列表,以及关联对应的角色
	managerList := []models.Manager{}
	models.DB.Preload("Role").Find(&managerList)

	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})
}

//添加管理员
func (con ManagerController) Add(c *gin.Context) {
	//获取角色
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}

//添加管理员:提交
func (con ManagerController) DoAdd(c *gin.Context) {
	//获取角色id,判断是否合法
	roleId, err := models.Int(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, "角色不合法", "/admin/manager/add")
		return
	}
	//获取提交的表单信息
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")

	//判断用户名和密码是否符合要求
	if len(username) < 2 || len(password) < 6 {
		con.Error(c, "用户名或密码长度不合法", "/admin/manager/add")
		return
	}

	//判断管理员是否存在
	managerList := []models.Manager{}
	models.DB.Where("username = ?", username).Find(&managerList)
	if len(managerList) > 0 {
		con.Error(c, "管理员已存在", "/admin/manager/add")
		return
	}

	//实例化Manager,执行增加管理员
	manager := models.Manager{
		Username: username,
		Password: models.Md5(password),
		Email:    email,
		Mobile:   mobile,
		AddTime:  int(models.GetUnix()),
		RoleId:   roleId,
		Status:   1,
	}
	err = models.DB.Create(&manager).Error
	if err != nil {
		con.Error(c, "添加管理员失败", "/admin/manager/add")
		return
	}
	con.Success(c, "添加管理员成功", "/admin/manager")
}

//编辑管理员
func (con ManagerController) Edit(c *gin.Context) {
	//获取管理员
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}
	manager := models.Manager{Id: id}
	models.DB.Find(&manager)

	if manager.Username == "" {
		con.Error(c, "管理员#" + models.String(id) + "不存在", "/admin/manager")
		return
	}
	//获取所有角色
	roleList := []models.Role{}
	models.DB.Find(&roleList)

	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,
		"roleList": roleList,
	})
}

//编辑管理员提交
func (con ManagerController) DoEdit(c *gin.Context) {
	//获取管理员id,并判断
	id, err := models.Int(c.PostForm("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}
	//获取角色id,并判断
	roleId, err2 := models.Int(c.PostForm("role_id"))
	if err2 != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}

	//获取提交的表单信息
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	//执行修改
	manager := models.Manager{Id: id}
	models.DB.Find(&manager)
	manager.Username = username
	manager.Email = email
	manager.RoleId = roleId
	manager.Mobile = mobile
	//判断密码, 为空 表示不修改密码
	if password != "" {
		//判断密码长度
		if len(password) < 6 {
			con.Error(c, "密码长度不合法", "/admin/manager/edit?id" + models.String(id))
			return
		}
		manager.Password = models.Md5(password)
	}
	//保存
	err = models.DB.Save(&manager).Error
	if err != nil {
		con.Error(c, "修改数据失败", "/admin/manager/edit?id="+models.String(id))
		return
	}
	con.Success(c, "修改数据成功", "/admin/manager")
}

//删除
func (con ManagerController) Delete(c *gin.Context) {
	//获取提交的表单数据
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}

	//查询管理员是否存在
	manager := models.Manager{Id: id}
	err = models.DB.Delete(&manager).Error
	if err != nil {
		con.Error(c, "删除数据失败", "/admin/manager")
		return
	}
	con.Success(c, "删除数据成功", "/admin/manager")
}
