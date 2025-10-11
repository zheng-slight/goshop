package frontend

//商品相关

import (
	"github.com/gin-gonic/gin"
	"goshop/models"
	"math"
	"net/http"
	"strings"
)

type ProductController struct {
	//extend 基础控制器
	BaseController
}

// 根据商品分类获取分类下面的所有商品数据
func (con ProductController) Category(c *gin.Context) {
	//获取分类id
	cateId, _ := models.Int(c.Param("id"))
	//当前页
	page, _ := models.Int(c.Query("page"))
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 2

	//获取当前分类
	curCate := models.GoodsCate{}
	models.DB.Where("id = ? ", cateId).Find(&curCate)

	//判断当前分类是否顶级分类,如果是,则获取对应的二级分类,如果不是,则获取对应的兄弟分类
	subCate := []models.GoodsCate{}
	var tempSlice []int
	if curCate.Pid == 0 { // 当前分类是顶级分类,获取对应的二级分类
		models.DB.Where("pid = ?", cateId).Find(&subCate)
		//把二级分类id放到切片中
		for i := 0; i < len(subCate); i++ {
			tempSlice = append(tempSlice, subCate[i].Id)
		}
	} else { // 当前分类是二级分类,获取对应的兄弟分类
		models.DB.Where("pid = ?", curCate.Pid).Find(&subCate)
	}
	//把请求的分类id放入切片
	tempSlice = append(tempSlice, cateId)
	//通过上面的分类id,获取商品相关数据
	goodsList := []models.Goods{}
	where := "cate_id in ?"
	models.DB.Where(where, tempSlice).Where("status = ?", 1).Offset((page - 1) * pageSize).Limit(pageSize).Find(&goodsList)

	//获取总数量
	var count int64
	models.DB.Where(where, tempSlice).Table("goods").Count(&count)

	//定义请求的模板
	tpl := "frontend/product/list.html"
	//判断分类模板,如果后台没有设置,则使用默认模板
	if curCate.Template != "" {
		tpl = curCate.Template
	}

	con.Render(c, tpl, gin.H{
		"goodsList":   goodsList,
		"subCate":     subCate,
		"currentCate": curCate,
		"page":        page,
		"totalPages":  math.Ceil(float64(count) / float64(pageSize)),
	})
}

//商品详情
func (con ProductController) Detail(c *gin.Context) {
	//获取商品id
	id, err := models.Int(c.Query("id"))
	//判断商品id是否符合要求
	if err != nil {
		c.Redirect(302, "/")
		c.Abort()
	}
	//1、获取商品信息
	goods := models.Goods{Id: id}
	models.DB.Find(&goods)

	//2、获取关联商品  RelationGoods
	relationGoods := []models.Goods{}
	//把数据库中保存的关联商品id转换成切片类型
	goods.RelationGoods = strings.ReplaceAll(goods.RelationGoods, "，", ",")
	relationIds := strings.Split(goods.RelationGoods, ",")
	//查询关联商品
	models.DB.Where("id in ?", relationIds).Select("id,title,price,goods_version").Find(&relationGoods)

	//3、获取关联赠品 GoodsGift
	goodsGift := []models.Goods{}
	goods.GoodsGift = strings.ReplaceAll(goods.GoodsGift, "，", ",")
	giftIds := strings.Split(goods.GoodsGift, ",")
	models.DB.Where("id in ?", giftIds).Select("id,title,price,goods_version").Find(&goodsGift)

	//4、获取关联颜色 GoodsColor
	goodsColor := []models.GoodsColor{}
	goods.GoodsColor = strings.ReplaceAll(goods.GoodsColor, "，", ",")
	colorIds := strings.Split(goods.GoodsColor, ",")
	models.DB.Where("id in ?", colorIds).Find(&goodsColor)

	//5、获取关联配件 GoodsFitting
	goodsFitting := []models.Goods{}
	goods.GoodsFitting = strings.ReplaceAll(goods.GoodsFitting, "，", ",")
	fittingIds := strings.Split(goods.GoodsFitting, ",")
	models.DB.Where("id in ?", fittingIds).Select("id,title,price,goods_version").Find(&goodsFitting)

	//6、获取商品关联的图片 GoodsImage
	goodsImage := []models.GoodsImage{}
	models.DB.Where("goods_id = ?", goods.Id).Limit(6).Find(&goodsImage)

	//7、获取规格参数信息 GoodsAttr
	goodsAttr := []models.GoodsAttr{}
	models.DB.Where("goods_id = ?", goods.Id).Find(&goodsAttr)

	//8、获取更多属性
	/*
			颜色:红色,白色,黄色 | 尺寸:41,42,43
			切片
			[
				{
					Cate:"颜色",
					List:[红色,白色,黄色]
				},
				{
					Cate:"尺寸",
					List:[41,42,43]
				}
			]

		goodsAttrStrSlice[0]	尺寸:41,42,43
				tempSlice[0]    尺寸
				tempSlice[1]	41,42,43

		goodsAttrStrSlice[1]	套餐:套餐1,套餐2

	*/

	// 更多属性: goodsAttrStr := "尺寸:41,42,43|套餐:套餐1,套餐2"
	goodsAttrStr := goods.GoodsAttr
	//字符串替换操作
	goodsAttrStr = strings.ReplaceAll(goodsAttrStr, "，", ",")
	goodsAttrStr = strings.ReplaceAll(goodsAttrStr, "：", ":")
	//实例化商品更多属性结构体
	var goodsItemAttrList []models.GoodsItemAttr
	//strings.Contains 判断字符串中有没有冒号(:)
	if strings.Contains(goodsAttrStr, ":") {
		//字符串替换操作:获取属性切片
		goodsAttrStrSlice := strings.Split(goodsAttrStr, "|")
		//创建切片的存储空间
		goodsItemAttrList = make([]models.GoodsItemAttr, len(goodsAttrStrSlice))
		for i := 0; i < len(goodsAttrStrSlice); i++ {
			//strings.Split(s, sep string) 把字符串s按照sep转换成切片
			//拆分 "尺寸:41,42,43
			tempSlice := strings.Split(goodsAttrStrSlice[i], ":")
			goodsItemAttrList[i].Cate = tempSlice[0]
			//拆分 41,42,43
			listSlice := strings.Split(tempSlice[1], ",")
			goodsItemAttrList[i].List = listSlice
		}
	}

	//定义请求的模板
	tpl := "frontend/product/detail.html"
	con.Render(c, tpl, gin.H{
		"goods":             goods,
		"relationGoods":     relationGoods,
		"goodsGift":         goodsGift,
		"goodsColor":        goodsColor,
		"goodsFitting":      goodsFitting,
		"goodsImage":        goodsImage,
		"goodsAttr":         goodsAttr,
		"goodsItemAttrList": goodsItemAttrList,
	})
}

//获取商品对应颜色的图库信息
func (con ProductController) GetImgList(c *gin.Context) {
	//获取商品id
	goodsId, err1 := models.Int(c.Query("goods_id"))
	//获取商品对应的颜色id
	colorId, err2 := models.Int(c.Query("color_id"))

	//查询商品图库信息
	goodsImageList := []models.GoodsImage{}
	err3 := models.DB.Where("goods_id = ? AND color_id = ?", goodsId, colorId).Find(&goodsImageList).Error
	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "",
			"message": "参数错误",
		})
		return
	}

	//判断 goodsImageList的长度 如果goodsImageList没有数据，那么我们需要返回当前商品所有的图库信息
	if len(goodsImageList) == 0 {
		models.DB.Where("goods_id = ?", goodsId).Find(&goodsImageList)
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  goodsImageList,
		"message": "获取数据成功",
	})
}
