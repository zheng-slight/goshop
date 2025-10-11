package models

//商品属性

type GoodsAttr struct {
	Id              int
	GoodsId         int		//商品id
	AttributeCateId int		//商品类型id,关联GoodsType.Id
	AttributeId     int		//商品类型属性id,关联GoodsTypeAttribute.Id
	AttributeTitle  string	//类型属性标题
	AttributeType   int		//类型属性录入方式:GoodsTypeAttribute.AttrType
	AttributeValue  string  //类型属性值
	Sort            int
	AddTime         int
	Status          int
}

func (GoodsAttr) TableName() string {
	return "goods_attr"
}
