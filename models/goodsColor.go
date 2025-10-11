package models

//商品颜色

type GoodsColor struct {
	Id         int
	ColorName  string // 颜色名称
	ColorValue string // 颜色值
	Status     int    // 状态
	Checked      bool   `gorm:"-"` //忽略该字段
}

func (GoodsColor) TableName() string {
	return "goods_color"
}
