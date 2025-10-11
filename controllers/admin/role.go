package admin

import (
	"github.com/gin-gonic/gin"
	"goshop/models"
	"net/http"
	"strings"
)

type RoleController struct {
	BaseController
}

//角色列表
func (con RoleController) Index(c *gin.Context) {
	//定义一个角色切片
	roleList := []models.Role{}
	//获取角色
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})
}

//新增角色
func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}

//新增角色:提交
func (con RoleController) DoAdd(c *gin.Context) {
	//获取表单的提交数据
	//strings.Trim(str, cutset), 去除字符串两边的cutset字符
	title := strings.Trim(c.PostForm("title"), " ") // 去除字符串两边的空格
	description := strings.Trim(c.PostForm("description"), " ")

	//判断角色名称是否为空
	if title == "" {
		con.Error(c, "角色名称不能为空", "/admin/role/add")
		return
	}
	//给角色模型赋值,并保存数据到数据库
	role := models.Role{}
	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = int(models.GetUnix())
	err := models.DB.Create(&role).Error
	if err != nil {
		con.Error(c, "增加角色失败,请重试", "/admin/role/add")
		return
	}
	con.Success(c, "增加角色成功", "/admin/role")
}

//编辑角色
func (con RoleController) Edit(c *gin.Context) {
	//获取角色id
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
	} else {
		role := models.Role{Id: id}
		models.DB.Find(&role)
		c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
			"role": role,
		})
	}
}

//编辑角色:提交
func (con RoleController) DoEdit(c *gin.Context) {
	//获取提交的表单数据
	id, err := models.Int(c.PostForm("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	//获取表单的提交数据
	//strings.Trim(str, cutset), 去除字符串两边的cutset字符
	title := strings.Trim(c.PostForm("title"), " ") // 去除字符串两边的空格
	description := strings.Trim(c.PostForm("description"), " ")
	//判断角色名称是否为空
	if title == "" {
		con.Error(c, "角色名称不能为空", "/admin/role/add")
		return
	}
	//查询角色是否存在
	role := models.Role{Id: id}
	models.DB.Find(&role)

	//修改角色属性
	role.Title = title
	role.Description = description
	err = models.DB.Save(&role).Error
	if err != nil {
		con.Error(c, "修改数据失败", "/admin/role/edit?id="+models.String(id))
		return
	}
	con.Success(c, "修改数据成功", "/admin/role")
}

//删除角色
func (con RoleController) Delete(c *gin.Context) {
	//获取提交的表单数据
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}

	//查询角色是否存在
	role := models.Role{Id: id}
	err = models.DB.Delete(&role).Error
	if err != nil {
		con.Error(c, "删除数据失败", "/admin/role")
		return
	}
	con.Success(c, "删除数据成功", "/admin/role")
}

//授权
func (con RoleController) Auth(c *gin.Context) {
	//获取id
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	role := models.Role{Id: id}
	models.DB.Find(&role)

	//获取所有权限列表
	accessList := []models.Access{}
	models.DB.Where("module_id = ?", 0).Preload("AccessItem").Find(&accessList)

	//获取当前角色拥有的权限,并把权限id放在一个map对象中
	roleAccess := []models.RoleAccess{}
	models.DB.Where("role_id = ?", id).Find(&roleAccess)
	roleAccessMap := make(map[int]int)
	for _, v := range roleAccess {
		roleAccessMap[v.AccessId] = v.AccessId
	}

	//循环遍历所有权限数据,判断当前权限的id是否在角色权限的map对象中,如果是的话给当前数据加入checked属性
	for i := 0; i < len(accessList); i++ { //循环权限列表
		if _, ok := roleAccessMap[accessList[i].Id]; ok { // 判断当前权限是否在角色权限的map对象中
			accessList[i].Checked = true
		}
		for j := 0; j < len(accessList[i].AccessItem); j++ { // 判断当前权限的子栏位是否在角色权限的map中
			if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok { // 判断当前权限是否在角色权限的map对象中
				accessList[i].AccessItem[j].Checked = true
			}
		}
	}
	c.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"roleId":     id,
		"accessList": accessList,
	})
}

//授权提交
func (con RoleController) DoAuth(c *gin.Context) {
	//获取提交的表单数据
	roleId, err := models.Int(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	//获取表单提交的权限id切片
	accessIds := c.PostFormArray("access_node[]")
	//先删除当前角色对应的权限
	roleAccess := models.RoleAccess{}
	models.DB.Where("role_id = ?", roleId).Delete(&roleAccess)

	//循环遍历accessIds,增加当前角色对应的权限
	for _, v := range accessIds {
		roleAccess.RoleId = roleId
		accessId, _ := models.Int(v)
		roleAccess.AccessId = accessId
		models.DB.Create(&roleAccess)
	}

	con.Success(c, "角色授权成功", "/admin/role")
}
