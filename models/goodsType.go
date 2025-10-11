package models

//商品类型

type GoodsType struct {
	Id          int
	Title       string  // 类型名称
	Description string  // 介绍
	Status      int  // 状态
	AddTime     int  // 添加时间
}

func (GoodsType) TableName() string {
	return "goods_type"
}
