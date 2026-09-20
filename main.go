package main

import (
	"fmt"
	"goshop/models"
	"goshop/routers"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func main() {
	fmt.Println("test \n")

	//初始化路由,会设置默认中间件:engine.Use(Logger(), Recovery())，可以使用gin.New()来设置路由
	r := gin.Default()
	//配置gin允许跨域请求
	r.Use(cors.New(cors.Config{ //自定义配置
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},             //允许的方法
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"}, //header允许山高月小
		AllowCredentials: false,
		MaxAge:           12 * time.Hour, //有效时间
		ExposeHeaders:    []string{"Content-Length"},
		AllowOriginFunc: func(origin string) bool { //允许的域
			return true //所有
		},
	}))
	//初始化基于redis的存储引擎: 需要启动redis服务,不然会报错
	//参数说明:
	//自第1个参数-redis最大的空闲连接数
	//第2个参数-数通信协议tcp或者udp
	//第3个参数-redis地址,格式，host:port 第4个参数-redis密码
	//第5个参数-session加密密钥
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	//自定义模板函数,必须在r.LoadHTMLGlob前面(只调用,不执行, 可以在html 中使用)
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime, //注册模板函数
		"Str2Html":   models.Str2Html,
		"FormatImg":  models.FormatImg,
		"Sub":        models.Sub,
		"SubStr":     models.SubStr,
		"FormatAttr": models.FormatAttr,
		"Mul":        models.Mul,
	})
	//加载templates中所有模板文件, 使用不同目录下名称相同的模板,注意:一定要放在配置路由之前才得行
	//如果模板在多级目录里面的话需要这样配置 r.LoadHTMLGlob("templates/**/**/*") /** 表示目录
	//LoadHTMLGlob只能加载同一层级的文件
	//比如说使用router.LoadHTMLFile("/templates/**/*")，就只能加载/templates/admin/或者/templates/order/下面的文件
	//解决办法就是通过filepath.Walk来搜索/templates下的以.html结尾的文件，把这些html文件都加载一个数组中，然后用LoadHTMLFiles加载
	//r.LoadHTMLGlob("templates/**/**/*")
	var files []string
	filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})
	r.LoadHTMLFiles(files...)

	//配置静态web目录 第一个参数表示路由,第二个参数表示映射的目录
	r.Static("/static", "./static")

	//分组路由文件
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.FrontendRoutersInit(r)

	//演示gopkg.in/ini.v1模块的使用
	_, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	} else {
		fmt.Println("app.ini加载成功")
	}
	// 典型读取操作，默认分区可以使用空字符串表示
	//fmt.Println("获取配置根数据:", cfg.Section("").Key("app_name").String())
	//fmt.Println("获取配置mysql数据:", cfg.Section("mysql").Key("password").String())
	//fmt.Println("获取配置redis数据:", cfg.Section("redis").Key("port").String())

	//给ini写入数据
	//修改某个值然后进行保存
	//cfg.Section("").Key("app_name").SetValue("app测试")
	//写入一个新的配置
	//cfg.Section("").Key("app_mode").SetValue("production")
	//cfg.Section("redis").Key("database").SetValue("1")

	//cfg.SaveTo("./conf/app.ini")

	r.Run() // 启动一个web服务
}
