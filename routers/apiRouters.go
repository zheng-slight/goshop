package routers

import (
	"goshop/controllers/api"
	"github.com/gin-gonic/gin"
)

//设置api路由
func ApiRoutersInit(r *gin.Engine) {
	//多版本api
	apiRouters := r.Group("/v1")
	{
		//获取导航列表
		apiRouters.GET("/navList", api.V1Controller{}.Navlist)
		//登录校验操作操作
		apiRouters.POST("/doLogin", api.V1Controller{}.DoLogin)
		//登录操作(生成token)
		apiRouters.GET("/login", api.V1Controller{}.Login)
		//编辑操作
		apiRouters.PUT("/editArticle", api.V1Controller{}.EditArticle)
		//删除操作
		apiRouters.DELETE("/deleteNav", api.V1Controller{}.DeleteNav)
		//获取收货地址(校验token)
		apiRouters.GET("/addressList", api.V1Controller{}.AddressList)
	}

	api2Routers := r.Group("/v2")
	{
		//获取导航列表
		api2Routers.GET("/navList", api.V2Controller{}.Navlist)
		//登录操作
		api2Routers.POST("/doLogin", api.V2Controller{}.DoLogin)
		//编辑操作
		api2Routers.PUT("/editArticle", api.V2Controller{}.EditArticle)
		//删除操作
		api2Routers.DELETE("/deleteNav", api.V2Controller{}.DeleteNav)
	}
}
