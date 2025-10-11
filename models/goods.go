package models

//商品表

type Goods struct {
	Id            int
	Title         string  //商品标题
	SubTitle      string  //附属标题
	GoodsSn       string  //商品编号
	CateId        int     //商品分类id: goods_cate.id
	ClickCount    int     //商品点击数量
	GoodsNumber   int     //商品库存
	Price         float64 //价格
	MarketPrice   float64 //商品市场价(原价)
	RelationGoods string  //关联商品id,如: 1, 23,55 ,商品id以逗号隔开
	GoodsAttr     string  //商品更多属性
	GoodsVersion  string  //商品版本
	GoodsImg      string  //图片
	GoodsGift     string  //商品赠品
	GoodsFitting  string  //商品配件
	GoodsColor    string  //颜色
	GoodsKeywords string  //SEO关键字
	GoodsDesc     string  //SEO商品描述
	GoodsContent  string  //商品详情
	IsDelete      int     //是否删除
	IsHot         int     //是否热销
	IsBest        int     //是否精品
	IsNew         int     //是否新品
	GoodsTypeId   int     //商品类型id,关联GoodsType.Id
	Sort          int     //排序
	Status        int     //状态
	AddTime       int     //添加时间
}

func (Goods) TableName() string {
	return "goods"
}

/**
根据分类id,商品类型获取分类下面的数据
*/
func GetGoodsByCategory(cateId int, goodsType string, limitNum int) []Goods {
	//判断是否顶级分类
	goodsCate := GoodsCate{Id: cateId}
	DB.Find(&goodsCate)
	var tempSlice []int
	if goodsCate.Pid == 0 { // 说明是顶级分类,则需要获取其下面的二级分类
		goodsCateList := []GoodsCate{}
		DB.Where("pid = ?", goodsCate.Id).Find(&goodsCate)
		//把二级分类id存入切片
		for i := 0; i < len(goodsCateList); i++ {
			tempSlice = append(tempSlice, goodsCateList[i].Id)
		}
	}
	tempSlice = append(tempSlice, goodsCate.Id)
	where := "cate_id in ?"

	//通过商品类型,拼接条件
	switch goodsType {
		case "is_best":
			where += " AND is_best = 1"
		case "is_hot":
			where += " AND is_hot = 1"
		case "is_new":
			where += " AND is_new = 1"
		default:
			break
	}
	goodsList := []Goods{}
	DB.Where(where, tempSlice).Order("sort DESC").Select("id, title, price, goods_img, sub_title").Limit(limitNum).Find(&goodsList)
	return goodsList
}
