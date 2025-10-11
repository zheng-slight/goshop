package models

//订单商品数据相关结构体

type OrderItem struct {
	Id           int
	OrderId      int	//订单号
	Uid          int  //用户id
	ProductTitle string  //商品名称
	ProductId    int  //商品id
	ProductImg   string //商品图片
	ProductPrice float64  //商品单价
	ProductNum   int  //购买数量
	GoodsVersion string  //商品版本
	GoodsColor   string  //商品颜色
	AddTime      int  //增加订单商品时间
}

func (OrderItem) TableName() string {
	return "order_item"
}
