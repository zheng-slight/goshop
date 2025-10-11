package models

//前端用户表相关

type User struct { // 结构体首字母大写, 和数据库表名对应, 默认访问数据表users, 可以设置访问数据表的方法
	Id       int
	Phone    string  `form:"phone" json:"phone"`  //手机号
	Password string  `form:"password" json:"password"`  // 密码
	AddTime  int  //增加时间
	LastIp   string  // 最后登录ip
	Email    string  // email
	Status   int  //状态
}

//配置数据库操作的表名称
func (User) TableName() string {
	return "user"
}
