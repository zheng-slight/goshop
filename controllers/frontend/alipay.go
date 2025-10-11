package frontend

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
)

type AlipayController struct{}

//支付宝支付请求接口
func (con AlipayController) Alipay(c *gin.Context) {
	//1、获取订单号 判断此订单号是否值当前用户的
	//2、获取订单里面的支付信息
	var privateKey = "xxx" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New("2021xxx588", privateKey, true)
	client.LoadAppPublicCertFromFile("crt/appCertPublicKey_2021xxx588.crt") // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile("crt/alipayRootCert.crt")   // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile("crt/alipayCertPublicKey_RSA2.crt")  // 加载支付宝公钥证书

	// 将 key 的验证调整到初始化阶段
	if err != nil {
		fmt.Println(err)
		return
	}

	//支付宝PC扫描方式支付
	var p = alipay.TradePagePay{}
	p.NotifyURL = "http://xxx/v3/alipayNotify"  // 回调地址
	p.ReturnURL = "http://xxx5/v3/alipayReturn"  //支付后跳转地址
	p.Subject = "测试 公钥证书模式-这是一个gin订单"
	template := "2006-01-02 15:04:05"
	p.OutTradeNo = time.Now().Format(template)
	p.TotalAmount = "0.1"  //支付金额(元)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"  //根据支付方式调整

	var url, err4 = client.TradePagePay(p)
	if err4 != nil {
		fmt.Println(err4)
	}

	//支付url
	var payURL = url.String()

	//重定向到支付url
	c.Redirect(http.StatusFound, payURL)

}

//支付回调
func (con AlipayController) AlipayNotify(c *gin.Context) {

	var privateKey = "xxx" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New("202xxx588", privateKey, true)
	client.LoadAppPublicCertFromFile("crt/appCertPublicKey_202xxx588.crt") // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile("crt/alipayRootCert.crt")  // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile("crt/alipayCertPublicKey_RSA2.crt") // 加载支付宝公钥证书

	if err != nil {
		fmt.Println(err)
		return
	}

	req := c.Request
	req.ParseForm()
	//支付校验
	ok, _ := client.VerifySign(req.Form)

	fmt.Println(ok)

	fmt.Println(req.Form)
	//订单逻辑处理...
	c.String(200, "ok")
}

//支付后跳转页面
func (con AlipayController) AlipayReturn(c *gin.Context) {
	c.String(200, "支付成功")
}
