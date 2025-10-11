package models

// 商品类型属性设置

type GoodsTypeAttribute struct {
	Id        int `json:"id"`  // HTML页面使用名称
	CateId    int `json:"cate_id"`   //商品类型id:商品类型表goods_type.id
	Title     string `json:"title"`   // 属性名称
	AttrType  int `json:"attr_type"`   //属性录入方式: 1 单行文本框, 2 多行文本框, 3 从下面列表中选择(一行代表一个可选值)
	AttrValue string `json:"attr_value"`   //可选值列表
	Status    int `json:"status"`   // 状态
	Sort      int `json:"sort"`   //排序
	AddTime   int `json:"add_time"`   //增加时间
}

func (GoodsTypeAttribute) TableName() string {
	return "goods_type_attribute"
}
