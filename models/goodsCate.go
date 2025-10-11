package models

//商品分类

type GoodsCate struct {
	Id             int
	Title          string  // 标题
	CateImg        string  // 分类图片
	Link           string  // 跳转地址
	Template       string  // 加载的模板: 为空的话加载默认模板, 不为空的话加载自定义模板
	Pid            int		// 上级id: 为0的话则是顶级分类
	SubTitle       string	// SEO标题
	Keywords       string	// SEO关键字
	Description    string	// SEO描述
	Sort           int	// 排序
	Status         int	// 状态: 1 显示, 0 隐藏
	AddTime        int	// 添加时间
	GoodsCateItems []GoodsCate `gorm:"foreignKey:pid;references:Id"` // 关联自身,下级分类
}

func (GoodsCate) TableName() string {
	return "goods_cate"
}
