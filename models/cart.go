package models

//购物车结构体

type Cart struct {
	Id           int  //商品id
	Title        string  //商品标题
	Price        float64  //价格
	GoodsVersion string  // 版本
	Uid          int  //用户id
	Num          int  // 数量
	GoodsGift    string  // 赠品
	GoodsFitting string  // 商品配件
	GoodsColor   string  // 颜色
	GoodsImg     string  //图片
	GoodsAttr    string  // 其他属性
	Checked      bool  //商品是否选中
}

//判断购物车里面有没有当前数据
func HasCartData(cartList []Cart, currentData Cart) bool {
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == currentData.Id && cartList[i].GoodsColor == currentData.GoodsColor && cartList[i].GoodsAttr == currentData.GoodsAttr {
			return true
		}
	}
	return false
}
