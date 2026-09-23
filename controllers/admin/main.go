package admin

// 进入后台系统首页
import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"goshop/models"
	"net/http"
)

type MainController struct {
	BaseController
}

func (con MainController) Index(c *gin.Context) {
	//获取Session里面保存的用户信息
	session := sessions.Default(c)
	userinfo := session.Get("userinfo_admin")
	//判断Session中的用户信息是否存在，如果不存在跳转到登录页面（注意需要判断） 如果存在继续向下执行
	//session.Get获取返回的结果是一个空接口类型,所以需要进行类型断言: 判断userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)
	if ok { // 说明是一个string
		//1.获取用户信息
		var adminSession models.AdminSession
		//把获取到的用户信息转换结构体
		json.Unmarshal([]byte(userinfoStr), &adminSession)

		//获取所有权限列表
		accessList := []models.Access{}
		models.DB.Where("module_id = ?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
			return db.Order("access.sort DESC")
		}).Order("sort DESC").Find(&accessList)

		//获取当前角色拥有的权限,并把权限id放在一个map对象中
		roleAccess := []models.RoleAccess{}
		models.DB.Where("role_id = ?", adminSession.RoleId).Find(&roleAccess)
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

		c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
			"username":   adminSession.Username,
			"isSuper":    adminSession.IsSuper, // 是否超级管理员, 超级管理员显示全部
			"accessList": accessList,
		})
	} else {
		c.Redirect(http.StatusFound, "/admin/login")
	}
}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}

// 公共方法:改变表状态的
func (con MainController) ChangeStatus(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "id参数错误",
		})
		return
	}
	table := c.Query("table")
	field := c.Query("field")

	//修改
	err1 := models.DB.Exec("update "+table+" set "+field+" = ABS("+field+" - 1) where id = ?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败,请重试",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "操作成功",
	})
}

// 公共方法:改变排序数字
func (con MainController) ChangeNum(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "id参数错误",
		})
		return
	}
	table := c.Query("table")
	field := c.Query("field")
	num := c.Query("num")

	//修改
	err1 := models.DB.Exec("update "+table+" set "+field+" = "+num+" where id = ?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败,请重试",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "操作成功",
	})
}

// 清除缓存
func (con MainController) FlushAll(c *gin.Context) {
	redisCache := models.RedisCache{}
	redisCache.FlushAll()
	con.Success(c, "清除缓存成功", "/admin")
}
