package routers

import (
	"github.com/gin-gonic/gin"
	"goshop/controllers/frontend"
	"goshop/middlewares"
)

//设置前端路由
func FrontendRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		//发送腾讯云短信
		defaultRouters.GET("/sms", frontend.BaseController{}.SmsTencent)

		//首页
		defaultRouters.GET("/", frontend.IndexController{}.Index)
		//商品分类对应的商品列表页面
		defaultRouters.GET("/category:id", frontend.ProductController{}.Category)

		//商品详情页
		defaultRouters.GET("/detail", frontend.ProductController{}.Detail)
		//获取商品图库信息
		defaultRouters.GET("/getImgList", frontend.ProductController{}.GetImgList)

		//获取购物车数据
		defaultRouters.GET("/cart", frontend.CartController{}.Get)
		//增加购物车数据
		defaultRouters.GET("/cart/addCart", frontend.CartController{}.AddCart)
		//购物车增加成功跳转页面
		defaultRouters.GET("/cart/successTip", frontend.CartController{}.AddCartSuccess)
		//减少购物车中商品数量
		defaultRouters.GET("/cart/decCart", frontend.CartController{}.DecCart)
		//增加购物车中商品数量
		defaultRouters.GET("/cart/incCart", frontend.CartController{}.IncCart)
		//改变一个商品数据的选中状态
		defaultRouters.GET("/cart/changeOneCart", frontend.CartController{}.ChangeOneCart)
		//全选反选
		defaultRouters.GET("/cart/changeAllCart", frontend.CartController{}.ChangeAllCart)
		//删除购物车商品数据
		defaultRouters.GET("/cart/delCart", frontend.CartController{}.DelCart)

		//购物车:确认选中商品页面:增加一个中间件:判断用户权限
		defaultRouters.GET("/buy/checkout", middlewares.InitUserAuthMiddleware, frontend.BuyController{}.Checkout)
		//提交订单执行结算
		defaultRouters.POST("/buy/doCheckout", middlewares.InitUserAuthMiddleware, frontend.BuyController{}.DoCheckout)
		//支付:去结支付页面
		defaultRouters.GET("/buy/pay", middlewares.InitUserAuthMiddleware, frontend.BuyController{}.Pay)

		//查看订单支付状态
		defaultRouters.GET("/buy/orderPayStatus", middlewares.InitUserAuthMiddleware, frontend.BuyController{}.OrderPayStatus)
		//支付宝支付
		defaultRouters.GET("/alipay", middlewares.InitUserAuthMiddleware, frontend.AlipayController{}.Alipay)
		//支付宝支付回调
		defaultRouters.POST("/alipayNotify", frontend.AlipayController{}.AlipayNotify)
		//支付宝支付完成后跳转
		defaultRouters.GET("/alipayReturn", middlewares.InitUserAuthMiddleware, frontend.AlipayController{}.AlipayReturn)
		//微信支付
		defaultRouters.GET("/wxpay", middlewares.InitUserAuthMiddleware, frontend.WxpayController{}.Wxpay)
		//微信支付回调
		defaultRouters.POST("/wxpay/notify", frontend.WxpayController{}.WxpayNotify)

		//收货地址:增加收货地址
		defaultRouters.POST("/address/addAddress", middlewares.InitUserAuthMiddleware, frontend.AddressController{}.AddAddress)
		//收货地址:编辑收货地址
		defaultRouters.POST("/address/editAddress", middlewares.InitUserAuthMiddleware, frontend.AddressController{}.EditAddress)
		//收货地址:修改默认收货地址
		defaultRouters.GET("/address/changeDefaultAddress", middlewares.InitUserAuthMiddleware, frontend.AddressController{}.ChangeDefaultAddress)
		//收货地址:获取一个指定收货地址id的收货地址
		defaultRouters.GET("/address/getOneAddressList", middlewares.InitUserAuthMiddleware, frontend.AddressController{}.GetOneAddressList)

		//用户登录页面
		defaultRouters.GET("/pass/login", frontend.PassController{}.Login)
		//获取图形验证码
		defaultRouters.GET("/pass/captcha", frontend.PassController{}.Captcha)
		//注册第一步
		defaultRouters.GET("/pass/registerStep1", frontend.PassController{}.RegisterStep1)
		//注册第二步
		defaultRouters.GET("/pass/registerStep2", frontend.PassController{}.RegisterStep2)
		//注册第三步
		defaultRouters.GET("/pass/registerStep3", frontend.PassController{}.RegisterStep3)
		//发送手机验证码
		defaultRouters.GET("/pass/sendCode", frontend.PassController{}.SendCode)
		//校验手机验证码
		defaultRouters.GET("/pass/validateSmsCode", frontend.PassController{}.ValidateSmsCode)
		//用户注册操作
		defaultRouters.POST("/pass/doRegister", frontend.PassController{}.DoRegister)
		//用户登录操作
		defaultRouters.POST("/pass/doLogin", frontend.PassController{}.DoLogin)
		//登出
		defaultRouters.GET("/pass/loginOut", frontend.PassController{}.LoginOut)
		//用户个人中心首页
		defaultRouters.GET("/user", middlewares.InitUserAuthMiddleware, frontend.UserController{}.Index)
		//用户订单列表
		defaultRouters.GET("/user/order", middlewares.InitUserAuthMiddleware, frontend.UserController{}.OrderList)
		//用户订单详情
		defaultRouters.GET("/user/orderinfo", middlewares.InitUserAuthMiddleware, frontend.UserController{}.OrderInfo)

		//设置es索引以及配置
		defaultRouters.GET("/search", frontend.SearchController{}.Index)
		//获取一条es数据
		defaultRouters.GET("/search/getOne", frontend.SearchController{}.GetOne)
		//增加数据到es中
		defaultRouters.GET("/search/addGoods", frontend.SearchController{}.AddGoods)
		//更新es中对应的数据
		defaultRouters.GET("/search/updateGoods", frontend.SearchController{}.UpdateGoods)
		//删除es中的数据
		defaultRouters.GET("/search/deleteGoods", frontend.SearchController{}.DeleteGoods)
		//模糊查询es数据
		defaultRouters.GET("/search/query", frontend.SearchController{}.Query)
		//条件筛选es查询
		defaultRouters.GET("/search/filterQuery", frontend.SearchController{}.FilterQuery)
		//分页查询es数据
		defaultRouters.GET("/search/pagingQuery", frontend.SearchController{}.PagingQuery)
	}
}
