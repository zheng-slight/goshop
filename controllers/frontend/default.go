package frontend

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	//引入模块的时候前面加个.表示可以直接使用模块里面的方法，无需加模块名称
	. "github.com/hunterhug/go_image"
	//引入模块的时候前面命一个名,表示模块的别名, 下面引入模块中的函数时,就可以直接以别名的方式进行调用
	qrcode "github.com/skip2/go-qrcode"
	"net/http"
)

type defaultController struct {
}

func (con defaultController) Index(c *gin.Context) {
	//设置session
	session := sessions.Default(c)
	session.Set("username", "张2三")
	session.Save() // 设置session的时候必须调用

	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg": "我是一个msg",
		"t":   1629788010,
	})
}

func (con defaultController) News(c *gin.Context) {
	//获取session
	//初始化session对象
	session := sessions.Default(c)
	//设置过期时间
	session.Options(sessions.Options{
		MaxAge: 3600 * 6, //6hrs
	})
	username := session.Get("username")
	c.String(http.StatusOK, "新闻--session.username=%v", username)
}

func (con defaultController) Shop(c *gin.Context) {
	//获取cookie
	username, _ := c.Cookie("username")

	c.String(http.StatusOK, "shop--cookie.username="+username)
}

//删除cookie
func (con defaultController) DeleteCookie(c *gin.Context) {
	//删除cookie
	c.SetCookie("username", "李四", -1, "/", "127.0.0.1", false, true)
}

//图片处理案例:按宽度进行比例缩放，输入输出都是文件
func (con defaultController) Thumbnail1(c *gin.Context) {
	//按宽度进行比例缩放，输入输出都是文件
	//func ScaleF2F(filename string, savepath string, width int) (err error)
	filename := "static/images/test.png"
	savepPath := "static/images/test_w_600.png"
	//ScaleF2F():go_image插件方法
	err := ScaleF2F(filename, savepPath, 600)
	if err != nil {
		c.JSON(http.StatusOK, "图片处理失败")
		return
	}
	c.JSON(http.StatusOK, "图片处理成功")

	//查看图像文件的真正名字
	//如 ./testdata/gopher500.jpg其实是png类型,但是命名错误,需要纠正!
	//RealImageName():go_image插件方法
	realfilename, err := RealImageName(filename)
	if err != nil {
		fmt.Printf("真正的文件名: %s->? err:%s\n", filename, err.Error())
	} else {
		fmt.Printf("真正的文件名:%s->%s\n", filename, realfilename)
	}

	//文件改名,强制性
	//ChangeImageName():go_image插件方法
	err = ChangeImageName(filename, realfilename, true)
	if err != nil {
		fmt.Printf("文件改名失败:%s->%s,%s\n", filename, realfilename, err.Error())
	} else {
		fmt.Println("改名成功")
	}

	//文件改名,不强制性
	err = ChangeImageName(filename, realfilename, false)
	if err != nil {
		fmt.Printf("文件改名失败:%s->%s,%s\n", filename, realfilename, err.Error())
	}
}

//图片处理案例:按宽度进行比例缩放，输入和输出都是图片字节数组
func (con defaultController) Thumbnail3(c *gin.Context) {
	//func ScaleB2B(InRaw []byte, width int) (OutRaw []byte, err error)
	//InRaw:内容, width:  宽度
	filename := "static/images/test.png"
	//获取图片的二进制流
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		c.JSON(http.StatusOK, "获取图片二进制流失败" + err.Error())
		return
	}
	//ScaleB2B():go_image插件方法
	img, err1 := ScaleB2B(bytes, 300)
	if err1 != nil {
		c.JSON(http.StatusOK, "图片处理失败" + err1.Error())
		return
	}
	c.String(http.StatusOK, string(img))
}

//图片处理案例:按宽度和高度进行比例缩放，图片宽高不满足的会进行图片的裁剪,输入和输出都是文件
func (con defaultController) Thumbnail2(c *gin.Context) {
	//按宽度进行比例缩放，输入输出都是文件
	//func ScaleF2F(filename string, savepath string, width int) (err error)
	filename := "static/images/test.png"
	savepPath := "static/images/test_w_400_h_300.png"
	//ThumbnailF2F():go_image插件方法
	err := ThumbnailF2F(filename, savepPath, 400, 300)
	if err != nil {
		c.JSON(http.StatusOK, "图片处理失败")
		return
	}
	c.JSON(http.StatusOK, "图片处理成功")
}

//二维码处理案例:直接生成一个byte的二进制
func (con defaultController) Qrcode1(c *gin.Context) {
	var png []byte
	//func Encode(content string, level RecoveryLevel, size int) ([]byte, error):
	//content:内容, level 图片质量, size" 大小
	png, err := qrcode.Encode("https://www.baidu.com", qrcode.Medium, 256)
	if err != nil {
		c.JSON(http.StatusOK, "二维生成失败")
		return
	}
	c.String(http.StatusOK, string(png))
}

//二维码处理案例:生成一个文件
func (con defaultController) Qrcode2(c *gin.Context) {
	savePath := "static/image/qr.png"
	//func WriteFile(content string, level RecoveryLevel, size int, filename string) error
	//content:内容, level 图片质量, size" 大小, filename: 文件名
	err := qrcode.WriteFile("https://www.baidu.com", qrcode.Medium, 256, savePath)
	if err != nil {
		c.JSON(http.StatusOK, "二维码生成失败")
		return
	}
	//读取文件
	file, _ := ioutil.ReadFile(savePath)
	c.String(http.StatusOK, string(file))
}
