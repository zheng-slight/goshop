package models

//导航管理

type Nav struct {
	Id        int  `json:"id"`
	Title     string `json:"title"` //标题
	Link      string	`json:"link"` //跳转地址
	Position  int	`json:"position"` //位置: 1 顶部, 2 中间, 3 底部, ...
	IsOpennew int	 `json:"is_opennew"` //是否新窗口打开: 1 否, 2 是
	Relation  string `json:"relation"`	//关联商品,填写商品id,以逗号隔开,如: 25,26
	Sort      int	 `json:"sort"` //排序
	Status    int	`json:"status"` //状态
	AddTime   int	`json:"add_time"` //增加时间
	GoodsItems []Goods `gorm:"-" json:"goods_items"`//忽略该字段,前端使用:获取导航相关商品数据
}

func (Nav) TableName() string {
	return "nav"
}
