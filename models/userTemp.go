package models

//用户发送验证码相关

type UserTemp struct {
	Id        int
	Ip        string  //ip地址
	Phone     string  //手机号
	SendCount int  //发送次数
	AddDay    string  // 生成日期
	AddTime   int  // 生成时间
	Sign      string  //页面跳转标签
}

func (UserTemp) TableName() string {
	return "user_temp"
}
