package models

//系统设置

type Setting struct {
	//使用反射功能:`form:"xxx"`配置:可以批量获取表单数据,form后面的需要和html中表单字段一致
	Id              int    `form:"id"`
	SiteTitle       string `form:"site_title"`	//标题
	SiteLogo        string	//logo
	SiteKeywords    string `form:"site_keywords"`	//关键词
	SiteDescription string `form:"site_description"`	//描述
	NoPicture       string	//默认图片,不需要form设置, 因为该字段的数据需要用到上传功能,要校验
	SiteIcp         string `form:"site_icp"`  //备案信息
	SiteTel         string `form:"site_tel"`	//电话
	SearchKeywords  string `form:"search_keywords"` //搜索关键词
	TongjiCode      string `form:"tongji_code"`	//统计代码
	Appid           string `form:"appid"`	//oss appid
	AppSecret       string `form:"app_secret"`	//oss 密钥
	EndPoint        string `form:"end_point"`	//oss 访问结点域名
	BucketName      string `form:"bucket_name"` //oss 桶名称
	OssStatus       int    `form:"oss_status"` //oss 开启状态: 1 是, 0 否
	OssDomain       string `form:"oss_domain"` //oss domain
	ThumbnailSize   string `form:"thumbnail_size"` //缩略图大小设置
}

func (Setting) TableName() string {
	return "setting"
}
