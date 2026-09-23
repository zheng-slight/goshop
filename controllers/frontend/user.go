package frontend

//用户中心

import (
	"github.com/gin-gonic/gin"
	"goshop/models"
	"math"
	"net/http"
)

type UserController struct {
	BaseController
}

// 个人中心首页
func (con UserController) Index(c *gin.Context) {
	var tpl = "frontend/user/welcome.html"
	con.Render(c, tpl, gin.H{})
}

// 我的订单列表
func (con UserController) OrderList(c *gin.Context) {
	// 页码数: 当前页
	page, _ := models.Int(c.Query("page"))
	if page == 0 {
		page = 1
	}
	//页大小
	pageSize := 2

	//获取当前用户
	user, _ := models.GetUserBySession(c)
	//模糊查询
	where := "uid =" + models.String(user.Id)
	keywords := c.Query("keywords")

	if keywords != "" {
		//查询订单商品title
		orderItemList := []models.OrderItem{}
		models.DB.Where("product_title like ?", "%"+keywords+"%").Find(&orderItemList)
		//拼接订单id
		var str string
		// 字符串：   12,12,22
		for i := 0; i < len(orderItemList); i++ {
			if i == 0 {
				str += models.String(orderItemList[i].OrderId)
			} else {
				str += "," + models.String(orderItemList[i].OrderId)
			}
		}
		//拼接where条件
		where += " AND id in (" + str + ")"
	}

	//获取订单状态: 按照状态筛选订单
	orderStatus, statusErr := models.Int(c.Query("orderStatus"))
	if statusErr == nil && orderStatus >= 0 { // 判断订单状态
		where += " AND order_status=" + models.String(orderStatus)
	} else {
		orderStatus = -1
	}

	//获取当前用户下面订单信息(条件查询,关联查询)
	orderList := []models.Order{}
	models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Preload("OrderItem").Order("add_time desc").Find(&orderList)

	//获取总数量
	var count int64
	models.DB.Where(where).Table("order").Count(&count)

	var tpl = "frontend/user/order.html"
	con.Render(c, tpl, gin.H{
		"order":       orderList,
		"page":        page,
		"keywords":    keywords,
		"orderStatus": orderStatus,
		"totalPages":  math.Ceil(float64(count) / float64(pageSize)), //计算总页数
	})
}

// 订单详情
func (con UserController) OrderInfo(c *gin.Context) {
	//获取有效的订单id
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.Redirect(http.StatusFound, "/user/order")
	}

	//获取用户数据
	user, _ := models.GetUserBySession(c)
	//获取订单数据
	order := []models.Order{}
	models.DB.Where("id = ? And uid = ?", id, user.Id).Preload("OrderItem").Find(&order)

	//判断订单是否存在:不存在则跳转到用户订单列表页面
	if len(order) == 0 {
		c.Redirect(http.StatusFound, "/user/order")
		return
	}

	//跳转到订单详情页面
	var tpl = "frontend/user/order_info.html"
	con.Render(c, tpl, gin.H{
		"order": order[0],
	})
}
