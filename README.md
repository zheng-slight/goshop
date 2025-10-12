
## 项目所需环境
1. 操作系统: Linux/Windows
2. Golang版本: go1.19+
3. 数据库: Mysql,redis, elasticsearch
4. fresh热加载环境:  当对代码进行修改时，程序能够自动重新加载并执行，gin 中并没有官方提供的热加载工具，这个时候要实现热加载就可以借助第三方的工具,安装步骤如下:
    (1).命令行执行: go install github.com/pilu/fresh@latest
    (2).然后在main.go所在目录,运行命令fresh即可进行热加载项目,这样可以减少断点调试以及重新启动项目的开发时间    
5. 以上环境请自行安装

## 项目介绍
1. 项目概述​​
*** 本项目是一个基于 ​​Golang Gin 框架​​ 开发的 ​​B2C 电商平台​​，采用 ​​MVC（Model-View-Controller）架构​​ 进行模块化设计，能够扩展为实现前后端分离，支持后台商品管理、用户系统、订单交易、支付集成、数据分析等功能,系统地展示了现代Web应用的全貌。该项目描绘了一个功能完整、技术选型现代的​​全栈电商项目​​。它从前端交互到后端管理，从业务逻辑到技术架构，从开发到运维，都做了全面的考量。使用Go语言作为后端，预示着项目对高并发性能有较高的要求。同时支持RESTful，也体现了技术上的前瞻性和灵活性。*** 
* 对应项目从最初阶段到运维阶段，项目主要步骤见博客专栏:https://blog.csdn.net/zhoupenghui168/category_12196464.html 

2. 后台管理系统​​
*** 这是商城的运营核心，供管理员使用，包含了对商城所有内容和数据的管控功能：***
* ​​用户RBAC权限管理​​： 系统的安全基石。RBAC（基于角色的访问控制）可以精细地控制不同管理员（如超级管理员、商品编辑、订单客服）的操作权限。
*​ ​轮播图管理​​： 可视化地管理首页或活动页面的宣传横幅广告。
* ​​商品管理​​： 商品的增、删、改、查，包括设置价格、库存、规格、详情图等。
* ​​会员管理​​： 查看和管理注册用户信息、积分、成长值等。
* ​​订单管理​​： 处理用户订单，包括查询、发货、退款、退货等全流程操作。
* ​​导航/文章/友情链接管理​​： 管理网站的辅助内容，如帮助中心文章、网站导航菜单、与其他网站的友情链接等。
* ​​商店设置​​： 配置商城的基础信息，如商城名称、Logo、客服联系方式等

3. 前端（面向用户的部分）​
*** 这部分是用户直接访问和交互的界面，涵盖了完整的电商购物流程：***
* ​​核心页面流​​： 首页-> 商品列表（分类浏览）-> 商品详情-> 加入购物车-> 结算-> 支付。
* ​​用户体系​​： 登录、注册、用户中心（查看订单、管理地址等）。
* ​​特色功能​​：
    + 频道页面： 通常用于活动专题页或特定商品聚合页。
    + 分词搜索： 提供更智能、更精准的商品搜索体验，支持输入关键词进行联想和匹配

4. API接口​
*** 现代应用通常支持多种终端，因此后端需要提供一套强大的API。该项目同时支持两种主流API风格：***
* ​​RESTful API​​： 一种传统但非常流行和规范的接口设计风格，结构清晰，易于理解。
* 用于支持：微信小程序、App、微信公众号

5. 其他技术与Linux运维​
*** 这部分是项目的技术基础设施，保证了系统的性能、安全性和可维护性 ***
* 认证与安全​​： Jwt + OAuth2.0： JWT用于用户身份令牌的生成与验证；OAuth2.0是行业标准的授权协议，可用于第三方登录（如微信、QQ登录）
* ​​存储与负载均衡​​：
    + 云存储： 将图片、文件等静态资源存储在云端（如阿里云OSS、腾讯云COS），减轻服务器压力。
    + Nginx负载均衡： 通过Nginx将用户请求分发到多台后端服务器，提升系统处理高并发的能力和可用性。
    + SSL/https： 为网站部署安全证书，对传输数据进行加密，保障安全。
* ​​缓存与部署​​：
    + Redis缓存数据： 将热点数据（如商品信息、会话）存入内存数据库Redis，极大提高读取速度。
    + Docker：使用容器化技术将应用和环境一起打包，实现快速、一致的跨平台（Windows/Linux）部署，简化运维流程
*  前后端分离​： 清晰的架构分层
*  ​​多端支持​： 一套后端支持多个前端平台
*  ​​现代化技术栈​： 采用当前主流的技术方案
*  ​​可扩展性： 模块化设计便于功能扩展

## 拉取项目
1. 拉取项目代码:
***这里我是把项目放在D:\www\go\src目录下的,进入src目录,执行git clone命令拉取项目代码: [git@github.com:phoenix-zhou/goshop.git](git@github.com:phoenix-zhou/goshop.git) ***
2. 进入项目根目录: cd goshop
3. 使用go mod tidy命令初始化项目依赖,步骤如下:
(1).使用go mod init goshop 初始化项目, 生成go.mod文件
(2).使用go mod tidy 更新依赖包
如果在go mod tidy 更新依赖时出现类似下面错误:
    go: finding module for package gorm.io/driver/mysql
    go: downloading gorm.io/driver/mysql v1.4.6
    gindemo/models imports
        gorm.io/driver/mysql: gorm.io/driver/mysql@v1.4.6: verifying module: gorm.io/driver/mysql@v1.4.6: checking tree#15366953 against tree#15495139: reading https://goproxy.io/sumdb/sum.golang.org/tile/8/1/234: 404 Not Found
        server response: not found nst tree#15495139: reading https://goproxy.io/sumdb/sum.golang.org/tile/8/1/234: 404 Not Found
        server response: not found
该原因是国内被墙了,解决方案:
    终端依次输入:
        go env -w GOSUMDB=off
        go env -w GOPROXY=https://goproxy.cn
    重新go mod tidy即可    
4. 上述步骤完成后,执行编译即可

## 项目打包以及部署
1. 项目打包成linux平台可执行文件:
    (1).进入main.go文件位置
    (2).依次执行下面操作命令
        set GOOS=linux
        set GOARCH=amd64
        go build -o goshop main.go   // goshop名称为项目名称
    (3). 如果(2)报以下错误
        # runtime/cgo
            cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%
        说明该程序依赖gcc编译器,请安装gcc编译器,并配置环境变量,如果已安装gcc,请检查环境变量是否配置正确
        还可以运行set CGO_ENABLED=0命令解决,然后再次运行go build -o goshop main.go,生成可执行文件
    
2. 部署:
(1).将编译好的项目文件goshop上传到服务器
(2).将config下的配置文件上传到服务器
(3).设置权限:chmod -R 777 goshop
(4).查看服务:sudo ps -ef | grep goshop
    如果服务存在,则杀死进程: sudo killall -9 goshop
(5).启动服务: nohup ./goshop &

## go.mod以及引用其他包操作详解
1. 生成go.mod命令:
    在main.go所在的文件夹下运行命令:
        -> go mod init goshop
    出现:
        -> go: creating new go.mod: module goshop
        -> go: to add module requirements and sums:
        ->        go mod tidy
    然后运行命令:
        ->        go mod tidy
        即可

2. 安装热加载:
    ->go get github.com/pilu/fresh
    ->fresh
    或者
    ->go install github.com/pilu/fresh@latest
    ->fresh

3. 安装session包: 在main.go对应的目录下运行:
-> go get github.com/gin-contrib/sessions
安装redis存储引擎的包: 在main.go对应的目录下运行:
-> go get github.com/gin-contrib/sessions/redis

4. gorm文档
https://gorm.io/zh_CN/docs/index.html
安装gorm
下载gorm以及对应的mysql驱动
-> go get -u gorm.io/gorm
-> go get -u gorm.io/driver/mysql

或者直接
import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)
然后 go mod tidy


5. 引入captcha
import (
    "github.com/mojocn/base64Captcha"
)
然后 go mod tidy

6. GoLang 图像处理插件使用
import (
    . "github.com/hunterhug/go_image"
)
然后 go mod tidy

7. GoLang 生成二维码
import (
    "github.com/skip2/go-qrcode"
)
然后 go mod tidy

8. GoLang oss上传图片
import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)
然后 go mod tidy

9. GoLang cos上传图片
import (
	"github.com/tencentyun/cos-go-sdk-v5"
)
然后 go mod tidy

10. markdown语法
import (
    "github.com/gomarkdown/markdown"
)
然后 go mod tidy


11. 腾讯云短信
import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
)
然后 go mod tidy

12. 跨域请求开启
import (
  "github.com/gin-contrib/cors"
)
然后 go mod tidy