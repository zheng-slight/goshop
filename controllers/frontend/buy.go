package frontend

//购物车商品选中结算页面

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goshop/models"
	"net/http"
)

type BuyController struct {
	BaseController
}

//确认订单页面
func (con BuyController) Checkout(c *gin.Context) {
	//获取购物车中选择的商品
	cartList := []models.Cart{}
	models.Cookie.Get(c, "cartList", &cartList)
	//选中的商品
	orderList := []models.Cart{}
	//总价格
	var allPrice float64
	//总数量
	var allNum int
	//循环购物车,并获取选中的商品信息
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			orderList = append(orderList, cartList[i])
			allNum += cartList[i].Num
		}
	}

	//获取当前用户的收货地址
	//获取用户
	user := models.User{}
	models.Cookie.Get(c, "user", &user)
	addressList := []models.Address{}
	fmt.Println(user)
	//通过用户id获取收货地址列表
	models.DB.Where("uid = ?", user.Id).Order("id desc").Find(&addressList)

	//生成签名: 进入'订单结算页面'时,防止重复提交订单
	orderSign := models.Md5(models.GetRandomNum())
	session := sessions.Default(c)
	session.Set("orderSign", orderSign)
	session.Save()

	//判断orderList数据是否存在
	if len(orderList) == 0 {  //不存在,则跳转到购物车页面
		c.Redirect(http.StatusFound, "/cart")
		return
	}

	con.Render(c, "frontend/buy/checkout.html", gin.H{
		"orderList": orderList,
		"allPrice":    allPrice,
		"allNum":      allNum,
		"addressList": addressList,
		"orderSign":   orderSign,
	})
}

/*
提交订单执行结算
   1.防止重复提交订单
   2.获取用户信息 ,获取用户的收货地址信息
   3.获取购买商品的信息
   4.把订单信息放在订单表，把商品信息放在订单商品表
   5.删除购物车里面的选中数据
   6.跳转到支付页面
*/
		func (con BuyController) DoCheckout(c *gin.Context) {
	//1.防止重复提交订单
	//获取签名
	orderSignClient := c.PostForm("orderSign")
	session := sessions.Default(c)
	//获取session中保存的签名
	orderSignSession := session.Get("orderSign")
	//进行类型断言,转换成string类型
	orderSignServer, ok := orderSignSession.(string)
	if !ok {  //转换失败或者没有签名,则跳转到 '购物车' 页面
		c.Redirect(http.StatusFound, "/cart")
		return
	}
	//判断传入的签名和session中保存的签名是否一致,一致:说明是上一步传过来的,不一致:说明不是上一步传过来的,跳转到购物车页面
	if orderSignClient != orderSignServer {
		c.Redirect(http.StatusFound, "/cart")
		return
	}
	//判断完毕后,删除签名
	session.Delete("orderSign")
	session.Save()

	// 2.获取用户信息,获取用户的收货地址信息
	user := models.User{}
	models.Cookie.Get(c, "user", &user)
    //定义用户默认收货地址结构体
	addressResult := []models.Address{}
	models.DB.Where("uid = ? AND default_address = 1", user.Id).Find(&addressResult)
	//判断是否存在默认收货地址,如果不存在,则跳转到'确认订单'页面
	if len(addressResult) == 0 {
		c.Redirect(http.StatusFound, "/buy/checkout")
		return
	}

	// 3.获取购买商品的信息:可以从cookie中获取,如果创建了购物车数据表,也可以从购物车数据表中获取
	//目前从cookie中获取选中的商品
	cartList := []models.Cart{}
	models.Cookie.Get(c, "cartList", &cartList)
	//定义选中商品结构体
	orderList := []models.Cart{}
	var allPrice float64
	//循环购物车中商品,把选中的商品放入orderList结构体中
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			orderList = append(orderList, cartList[i])
		}
	}
	// 4.把订单信息放在订单表，把商品信息放在商品表
	order := models.Order{
		OrderId:     models.GetOrderId(),
		Uid:         user.Id,
		AllPrice:    allPrice,
		Phone:       addressResult[0].Phone,
		Name:        addressResult[0].Name,
		Address:     addressResult[0].Address,
		PayStatus:   0,
		PayType:     0,
		OrderStatus: 0,
		AddTime:     int(models.GetUnix()),
	}

	err := models.DB.Create(&order).Error

	//增加数据成功以后可以通过order.Id
	if err == nil {
		// 把商品信息放在商品订单表
		for i := 0; i < len(orderList); i++ {
			orderItem := models.OrderItem{
				OrderId:      order.Id,
				Uid:          user.Id,
				ProductTitle: orderList[i].Title,
				ProductId:    orderList[i].Id,
				ProductImg:   orderList[i].GoodsImg,
				ProductPrice: orderList[i].Price,
				ProductNum:   orderList[i].Num,
				GoodsVersion: orderList[i].GoodsVersion,
				GoodsColor:   orderList[i].GoodsColor,
			}
			models.DB.Create(&orderItem)
		}
	}

	// 5.删除购物车里面的选中数据
	noSelectCartList := []models.Cart{}
	for i := 0; i < len(cartList); i++ {
		if !cartList[i].Checked {
			noSelectCartList = append(noSelectCartList, cartList[i])
		}
	}
	models.Cookie.Set(c, "cartList", noSelectCartList)

	//跳转到'去支付'页面
	c.Redirect(http.StatusFound, "/buy/pay?orderId="+models.String(order.Id))
}

//支付:去支付页面
func (con BuyController) Pay(c *gin.Context) {
	//获取订单id
	orderId, err := models.Int(c.Query("orderId"))
	if err != nil {  // 订单id类型错误
		c.Redirect(http.StatusFound, "/cart")
	}
	//获取用户信息
	user := models.User{}
	models.Cookie.Get(c, "user", &user)
	//获取订单信息
	order := models.Order{}
	models.DB.Where("id = ?", orderId).Find(&order)
	if order.Uid != user.Id { //订单信息错误
		c.Redirect(http.StatusFound, "/cart")
		return
	}
	//获取订单对应的商品
	orderItems := []models.OrderItem{}
	models.DB.Where("order_id = ?", orderId).Find(&orderItems)

	//渲染页面
	con.Render(c, "frontend/buy/pay.html", gin.H{
		"order":      order,
		"orderItems": orderItems,
	})
}

//查看订单支付状态
func (con BuyController) OrderPayStatus(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入参数错误",
		})
		return
	}
	//获取用户信息
	user := models.User{}
	models.Cookie.Get(c, "user", &user)

	//获取主订单信息
	order := models.Order{}
	models.DB.Where("id = ?", id).Find(&order)

	//判断当前数据是否合法
	if user.Id != order.Uid {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "非法请求",
		})
		return
	}

	//判断是否支付
	if order.PayStatus == 1 && order.OrderStatus == 1 {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "支付成功",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "支付成功",
		})
		return
	}
}