package frontend

//基础控制器

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"goshop/models"
	"net/http"
	"os"
	"strings"
)

type BaseController struct{}

/*
加载公共模板方法
tpl string 模板
data map 请求的数据
*/
func (con BaseController) Render(c *gin.Context, tpl string, data map[string]interface{}) {
	//实例化redisCache结构体
	redisCache := models.RedisCache{}

	//获取顶部导航列表
	topNavList := []models.Nav{}
	//判断redis中是否存在数据
	if hasTopNavList := redisCache.Get("topNavList", &topNavList); !hasTopNavList { //不存在数据,则从数据中获取数据,并把数据保存到redis
		models.DB.Where("status = 1 AND position = 1").Find(&topNavList)
		redisCache.Set("topNavList", topNavList, 3600)
	}
	//获取分类数据
	goodsCateList := []models.GoodsCate{}
	if hasGoodsCateList := redisCache.Get("goodsCateList", &goodsCateList); !hasGoodsCateList {
		//获取分类列表以及下级分类,并进行排序
		models.DB.Where("pid = ? AND status = ?", 0, 1).Order("sort DESC").Preload("GoodsCateItems", func(db *gorm.DB) *gorm.DB {
			return db.Where("goods_cate.status = 1").Order("goods_cate.sort DESC")
		}).Find(&goodsCateList)
		redisCache.Set("goodsCateList", goodsCateList, 3600)
	}

	//获取中间导航
	middleNavList := []models.Nav{}
	if hasMiddleNavList := redisCache.Get("middleNavList", &middleNavList); !hasMiddleNavList {
		models.DB.Where("status = ? AND position = ? ", 1, 2).Find(&middleNavList)
		//循环,获取中间导航对应的商品数据
		for i := 0; i < len(middleNavList); i++ {
			//获取管理商品
			//替换字符串中的中文逗号strings.ReplaceAll()
			relation := strings.ReplaceAll(middleNavList[i].Relation, "，", ",")
			//把字符串转换成切片
			relationIds := strings.Split(relation, ",")
			//获取对应的商品信息
			goodsList := []models.Goods{}
			models.DB.Where("status = ?", 1).Where("id in ?", relationIds).Select("id, title, goods_img, price").Find(&goodsList)
			middleNavList[i].GoodsItems = goodsList
		}
		redisCache.Set("middleNavList", middleNavList, 3600)
	}

	//获取Cookie里面保存的用户信息
	user := models.User{}
	isLogin := models.Cookie.Get(c, "user", &user)
	var userinfo string
	if isLogin && len(user.Phone) == 11 {
		userinfo = fmt.Sprintf(`<li class="userinfo">
			<a href="#">%v</a>		
			<i class="i"></i>
			<ol>
				<li><a href="/user">个人中心</a></li>
				<li><a href="#">喜欢</a></li>
				<li><a href="/pass/loginOut">退出登录</a></li>
			</ol>								
		</li> `, user.Phone)
	} else {
		userinfo = fmt.Sprintf(`<li><a href="/pass/login">登录</a></li>
		<li>|</li>
		<li><a href="/pass/registerStep1" target="_blank" >注册</a></li>
		<li>|</li>`)
	}

	renderData := gin.H{
		"topNavList":    topNavList,
		"goodsCateList": goodsCateList,
		"middleNavList": middleNavList,
		"userinfo":      userinfo,
	}

	for key, v := range data {
		renderData[key] = v
	}

	c.HTML(http.StatusOK, tpl, renderData)
}

//发送短信功能
func (con BaseController) SmsTencent(c *gin.Context) {
	//获取电话号码
	phone := "19950326585"
	//获取随机数
	code := models.GetRandomNum()
	//发送短信
	_, err := models.SmsTencent(phone, code)
	if err != nil {
		fmt.Println(err)
	}
	//发送成功后处理逻辑
	//保存短信验证码到redis,以及记录发送记录等
	//保存短信随机数到redis,以便后续校验短信验证码
	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	//保存的key
	session_key := cfg.Section("sms_tencent").Key("session_key").String()
	//过期时间
	captcha_expiring := cfg.Section("sms_tencent").Key("captcha_expiring").String()
	expiration, _ := models.Int(captcha_expiring)
	//保存的key
	key := session_key + ":" + phone
	//保存的值
	value := map[string]string{
		"code": code,
	}
	redisCache := models.RedisCache{}
	redisCache.Set(key, value, expiration)

	//获取保存的数据
	var obj map[string]string
	if hasCode := redisCache.Get(key, &obj); !hasCode {
		fmt.Println("错误")
	}
	fmt.Println(obj["code"])

	//保存发送记录到数据库(略)
}
