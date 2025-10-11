package models

//订单表相关结构体

type Order struct {
	Id          int
	OrderId     string		//订单号
	Uid         int  //用户id
	AllPrice    float64  //订单总价
	Phone       string  //收货人手机号
	Name        string  //收货人姓名
	Address     string  //收货地址
	PayStatus   int // 支付状态： 0 表示未支付, 1 已经支付
	PayType     int // 支付类型： 0 alipay, 1 wechat
	OrderStatus int // 订单状态： 0 已下单 ,1 已付款,2 已配货,3、发货, 4 交易成功, 5 退货, 6 取消
	AddTime     int  //订单生成时间
	PayTime          int //支付时间
	DistributionTime int //配货时间
	ExwarehouseTime  int //出库时间
	SuccessfulTime   int //交易成功时间
	CancelTime       int //取消时间
	ReturnTime       int //退款时间
	LogisticsCompany int //物流公司id: 可以创建对应的物流公司数据表,然后在后台增加对应的物流信息
	WaybillNo        string //运单号
	//其他的字段

	//订单商品关联关系
	OrderItem []OrderItem `gorm:"foreignKey:OrderId;references:Id"`
}

func (Order) TableName() string {
	return "order"
}
