package models

//商品更多属性结构体

type GoodsItemAttr struct {
	// 尺寸:41,42,43
	Cate string  //属性类型:  尺寸
	List []string  // 对应的值 [41,42,43]
}
