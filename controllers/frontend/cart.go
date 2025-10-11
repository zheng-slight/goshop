package frontend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goshop/models"
	"net/http"
)

type CartController struct {
	BaseController
}

//获取购物车数据
func (con CartController) Get(c *gin.Context) {
	//获取购物车数据 显示购物车数据
	cartList := []models.Cart{}
	models.Cookie.Get(c, "cartList", &cartList)
	//定义一个总价
	var allPrice float64

	for i := 0; i < len(cartList); i++ {  //计算总价: 价格*数量
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}
	var tpl = "frontend/cart/cart.html"
	con.Render(c, tpl, gin.H{
		"cartList": cartList,
		"allPrice": allPrice,
	})
}

//添加购物车
func (con CartController) AddCart(c *gin.Context) {
	/*
	   购物车数据保持到哪里？
	           1、购物车数据保存在本地    （cookie）
	           2、购物车数据保存到服务器(mysql)   （必须登录）
	           3、没有登录 购物车数据保存到本地 ， 登录成功后购物车数据保存到服务器  （用的最多）
	   增加购物车的实现逻辑：
	           1、获取增加购物车的数据  （把哪一个商品加入到购物车）
	           2、判断购物车有没有数据   （cookie）
	           3、如果购物车没有任何数据  直接把当前数据写入cookie
	           4、如果购物车有数据
	              1、判断购物车有没有当前数据
	                       有当前数据让当前数据的数量加1，然后写入到cookie
	              2、如果没有当前数据直接写入cookie
	*/
	// 1、获取增加购物车的数据,放在结构体里面  （把一个商品加入到购物车）
	colorId, _ := models.Int(c.Query("color_id"))
	goodsId, err := models.Int(c.Query("goods_id"))
	if err != nil {
		c.Redirect(http.StatusFound, "/")
	}
	//实例化商品结构体
	goods := models.Goods{}
	//实例化商品颜色结构体
	goodsColor := models.GoodsColor{}
	//获取商品信息
	models.DB.Where("id = ?", goodsId).Find(&goods)
	//获取商品颜色信息
	models.DB.Where("id = ?", colorId).Find(&goodsColor)

	currentData := models.Cart{
		Id:           goodsId,
		Title:        goods.Title,
		Price:        goods.Price,
		GoodsVersion: goods.GoodsVersion,
		Num:          1,
		GoodsColor:   goodsColor.ColorName,
		GoodsImg:     goods.GoodsImg,
		GoodsGift:    goods.GoodsGift, /*赠品*/
		GoodsAttr:    "",              //根据自己的需求拓展
		Checked:      true,            /*默认选中*/
	}

	// 2、判断购物车有没有数据（cookie）
	cartList := []models.Cart{}
	models.Cookie.Get(c, "cartList", &cartList)
	fmt.Println(cartList)
	if len(cartList) > 0 { // 说明有数据
		//4、购物车有数据:判断购物车有没有当前数据
		if models.HasCartData(cartList, currentData) { //有当前商品数据,数量加1
			for i := 0; i < len(cartList); i++ {
				if cartList[i].Id == currentData.Id && cartList[i].GoodsColor == currentData.GoodsColor && cartList[i].GoodsAttr == currentData.GoodsAttr {
					cartList[i].Num = cartList[i].Num + 1
				}
			}
		} else { // 购物车中没有当前商品, 则加入购物车
			cartList = append(cartList, currentData)
		}
		//更新购物车
		models.Cookie.Set(c, "cartList", cartList)
	} else {
		// 3、如果购物车没有任何数据  直接把当前数据写入cookie
		cartList = append(cartList, currentData)
		fmt.Println(cartList)
		models.Cookie.Set(c, "cartList", cartList)
	}

	//添加购物车成功后,跳转到添加成功页面
	c.Redirect(http.StatusFound, "/cart/successTip?goods_id=" + models.String(goodsId))
}

//购物车增加成功跳转页面
func (con CartController) AddCartSuccess(c *gin.Context) {
	goodsId, err := models.Int(c.Query("goods_id"))
	if err != nil {
		c.Redirect(http.StatusFound, "/")
	}
	//实例化商品
	goods := models.Goods{}
	models.DB.Where("id = ?", goodsId).Find(&goods)

	//跳转到添加购物车成功的模板页面
	var tpl = "frontend/cart/addcart_success.html"
	con.Render(c, tpl, gin.H{
		"goods": goods,
	})
}

//增加购物车数量
func (con CartController) IncCart(c *gin.Context) {
	//1、获取客户端穿过来的数据
	goodsId, err := models.Int(c.Query("goods_id"))
	goodsColor := c.Query("goods_color")
	GoodsAttr := ""

	//定义返回的数据
	//总的选中商品价格
	var allPrice float64
	//当前商品总价
	var currentPrice float64
	var num int

	var response gin.H

	//2、判断数据是否合法
	if err != nil {
		response = gin.H{
			"success": false,
			"message": "传入参数错误",
		}
	} else {
		//获取购物车数据
		cartList := []models.Cart{}
		models.Cookie.Get(c, "cartList", &cartList)
		if len(cartList) > 0 {
			for i := 0; i < len(cartList); i++ { //循环购物车,找到要增加的商品,给其数量+1
				if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == GoodsAttr {
					cartList[i].Num = cartList[i].Num + 1
					currentPrice = float64(cartList[i].Num) * cartList[i].Price
					num = cartList[i].Num
				}
				//判断该商品是否选中,并计算总价格
				if cartList[i].Checked {
					allPrice += cartList[i].Price * float64(cartList[i].Num)
				}
			}
			//重新写入数据
			models.Cookie.Set(c, "cartList", cartList)

			response = gin.H{
				"success":      true,
				"message":      "更新数据成功",
				"allPrice":     allPrice,
				"num":          num,
				"currentPrice": currentPrice,
			}
		} else {
			response = gin.H{
				"success": false,
				"message": "传入参数错误",
			}
		}
	}
	c.JSON(http.StatusOK, response)
}

//减少购物车数量
func (con CartController) DecCart(c *gin.Context) {
	//1、获取客户端穿过来的数据
	goodsId, err := models.Int(c.Query("goods_id"))
	goodsColor := c.Query("goods_color")
	GoodsAttr := ""

	//定义返回的数据
	//总价
	var allPrice float64
	//当前商品的价格小计
	var currentPrice float64
	var num int

	var response gin.H
	//2、判断数据是否合法
	if err != nil {
		response = gin.H{
			"success": false,
			"message": "传入参数错误",
		}
	} else {
		cartList := []models.Cart{}
		models.Cookie.Get(c, "cartList", &cartList)
		if len(cartList) > 0 {
			for i := 0; i < len(cartList); i++ {
				if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == GoodsAttr {
					if cartList[i].Num > 1 {
						cartList[i].Num = cartList[i].Num - 1
					}
					currentPrice = float64(cartList[i].Num) * cartList[i].Price
					num = cartList[i].Num
				}

				if cartList[i].Checked {
					allPrice += cartList[i].Price * float64(cartList[i].Num)
				}
			}
			//重新写入数据
			models.Cookie.Set(c, "cartList", cartList)

			response = gin.H{
				"success":      true,
				"message":      "更新数据成功",
				"allPrice":     allPrice,
				"num":          num,
				"currentPrice": currentPrice,
			}
		} else {
			response = gin.H{
				"success": false,
				"message": "传入参数错误",
			}
		}
	}

	c.JSON(http.StatusOK, response)
}

//改变一个商品数据的选中状态
func (con CartController) ChangeOneCart(c *gin.Context) {
	//1、获取客户端传过来的数据
	goodsId, err := models.Int(c.Query("goods_id"))
	goodsColor := c.Query("goods_color")
	GoodsAttr := ""

	//定义返回的数据
	var allPrice float64

	var response gin.H
	//2、判断数据是否合法
	if err != nil {
		response = gin.H{
			"success": false,
			"message": "传入参数错误",
		}
	} else {
		//获取购物车中的商品数据
		cartList := []models.Cart{}
		models.Cookie.Get(c, "cartList", &cartList)
		if len(cartList) > 0 {
			for i := 0; i < len(cartList); i++ { // 修改其选中状态
				if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == GoodsAttr {
					cartList[i].Checked = !cartList[i].Checked
				}

				if cartList[i].Checked { // 计算购物车商品总价格
					allPrice += cartList[i].Price * float64(cartList[i].Num)
				}

			}
			//重新写入数据
			models.Cookie.Set(c, "cartList", cartList)

			response = gin.H{
				"success":  true,
				"message":  "更新数据成功",
				"allPrice": allPrice,
			}
		} else {
			response = gin.H{
				"success": false,
				"message": "传入参数错误",
			}
		}
	}

	c.JSON(http.StatusOK, response)
}

//全选反选
func (con CartController) ChangeAllCart(c *gin.Context) {
	//获取flag标签: 1 全选, 0 反选
	flag, _ := models.Int(c.Query("flag"))

	//定义返回的数据
	var allPrice float64

	var response gin.H
	//获取购物车数据
	cartList := []models.Cart{}
	models.Cookie.Get(c, "cartList", &cartList)

	if len(cartList) > 0 {  // 判断购物车是否为空
		for i := 0; i < len(cartList); i++ {
			if flag == 1 {  // 全选
				cartList[i].Checked = true
			} else {  // 反选
				cartList[i].Checked = false
			}
			if cartList[i].Checked { //计算总价格
				allPrice += cartList[i].Price * float64(cartList[i].Num)
			}

		}
		//重新写入数据
		models.Cookie.Set(c, "cartList", cartList)

		response = gin.H{
			"success":  true,
			"message":  "更新数据成功",
			"allPrice": allPrice,
		}
	} else {
		response = gin.H{
			"success": false,
			"message": "传入参数错误",
		}
	}

	c.JSON(http.StatusOK, response)
}

//删除购物车数据
func (con CartController) DelCart(c *gin.Context) {
	//获取请求的商品id以及颜色
	goodsId, _ := models.Int(c.Query("goods_id"))
	goodsColor := c.Query("goods_color")
	GoodsAttr := ""

	//获取购物车数据
	cartList := []models.Cart{}
	models.Cookie.Get(c, "cartList", &cartList)

	for i := 0; i < len(cartList); i++ {
		//找到要删除的商品
		if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == GoodsAttr {
			//cartList[:i]: 目标商品之前的商品数据, cartList[(i+1):]:目标商品之后的商品数据
			//重构购物车数据
			cartList = append(cartList[:i], cartList[(i+1):]...)
		}
	}
	//保存新的购物车数据到cookie
	models.Cookie.Set(c, "cartList", cartList)
	//重定向跳转
	c.Redirect(http.StatusFound, "/cart")
}