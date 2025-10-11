package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goshop/models"
	"math"
	"net/http"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup //可以实现主线程等待协程执行完毕

type GoodsController struct {
	BaseController
}

func (con GoodsController) Index(c *gin.Context) {
	//分页查询
	page, _ := models.Int(c.Query("page")) // 获取分页,当分页page数据格式不正确时, page == 0
	if page == 0 {
		page = 1
	}
	//每页查询的数量
	pageSize := 2

	//定义商品结构体
	goodsList := []models.Goods{}
	//条件
	where := "is_delete=0"
	//关键字查询
	keyword := c.Query("keyword")
	if len(keyword) > 0 {
		where += " and title like \"%" + keyword + "%\""
		//也可以使用下面的方式
		//where += ` and title like "%` + keyword +`%"`
	}
	//分页查询
	models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&goodsList)

	//获取总数量
	var count int64
	models.DB.Table("goods").Where(where).Count(&count)
	//计算总页数:math.Ceil()向上取整,注意float64类型
	totalPages := math.Ceil(float64(count) / float64(pageSize))

	//判断最后一页有没有数据,如果没有则跳转到第一页
	if len(goodsList) > 0 {
		c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{
			"goodsList":  goodsList,
			"totalPages": totalPages,
			"page":       page,
			"keyword":    keyword,
		})
		return
	}

	//最后一页没有数据,判断页码数是否等于1,如果不等于,则重定向到列表页
	if page != 1 {
		c.Redirect(http.StatusFound, "/admin/goods")
		return
	}

	c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{
		"goodsList":  goodsList,
		"totalPages": totalPages,
		"page":       page,
		"keyword":    keyword,
	})
}

func (con GoodsController) Add(c *gin.Context) {
	//获取商品分类
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid = 0").Preload("GoodsCateItems").Find(&goodsCateList)
	//获取商品颜色
	goodsColorList := []models.GoodsColor{}
	models.DB.Where("status = 1").Find(&goodsColorList)
	//获取商品类型
	goodsTypeList := []models.GoodsType{}
	models.DB.Where("status = 1").Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goods/add.html", gin.H{
		"goodsCateList":  goodsCateList,
		"goodsColorList": goodsColorList,
		"goodsTypeList":  goodsTypeList,
	})
}

//增加商品信息表单提交
func (con GoodsController) DoAdd(c *gin.Context) {
	//获取表单提交过来的数据,并进行判断是否合法
	title := c.PostForm("title")
	subTitle := c.PostForm("sub_title")
	goodsSn := c.PostForm("goods_sn")
	cateId, _ := models.Int(c.PostForm("cate_id"))
	goodsNumber, _ := models.Int(c.PostForm("goods_number"))
	//价格,注意小数点
	marketPrice, _ := models.Float(c.PostForm("market_price"))
	price, _ := models.Float(c.PostForm("price"))

	relationGoods := c.PostForm("relation_goods")
	goodsAttr := c.PostForm("goods_attr")
	goodsVersion := c.PostForm("goods_version")
	goodsGift := c.PostForm("goods_gift")
	goodsFitting := c.PostForm("goods_fitting")
	//颜色:获取的是切片
	goodsColorArr := c.PostFormArray("goods_color")

	goodsKeywords := c.PostForm("goods_keywords")
	goodsDesc := c.PostForm("goods_desc")
	goodsContent := c.PostForm("goods_content")
	isDelete, _ := models.Int(c.PostForm("is_delete"))
	isHot, _ := models.Int(c.PostForm("is_hot"))
	isBest, _ := models.Int(c.PostForm("is_best"))
	isNew, _ := models.Int(c.PostForm("is_new"))
	goodsTypeId, _ := models.Int(c.PostForm("goods_type_id"))
	sort, _ := models.Int(c.PostForm("sort"))
	status, _ := models.Int(c.PostForm("status"))
	addTime := int(models.GetUnix())

	//获取颜色信息,把颜色转换成字符串
	goodsColorStr := strings.Join(goodsColorArr, ",")
	//上传图片,生成缩略图
	goodsImg, _ := models.UploadImg(c, "goods_img")
	if len(goodsImg) > 0 {
		//判断本地图片才需要处理缩略图
		if models.GetOssStatus() != 1 {
			//开启协程
			wg.Add(1)
			go func() {
				models.ResizeGoodsImage(goodsImg)
				wg.Done()
			}()
		}
	}
	//增加商品数据
	//实例化商品结构体
	goods := models.Goods{
		Title:         title,
		SubTitle:      subTitle,
		GoodsSn:       goodsSn,
		CateId:        cateId,
		ClickCount:    100,
		GoodsNumber:   goodsNumber,
		MarketPrice:   marketPrice,
		Price:         price,
		RelationGoods: relationGoods,
		GoodsAttr:     goodsAttr,
		GoodsVersion:  goodsVersion,
		GoodsGift:     goodsGift,
		GoodsFitting:  goodsFitting,
		GoodsKeywords: goodsKeywords,
		GoodsDesc:     goodsDesc,
		GoodsContent:  goodsContent,
		IsDelete:      isDelete,
		IsHot:         isHot,
		IsBest:        isBest,
		IsNew:         isNew,
		GoodsTypeId:   goodsTypeId,
		Sort:          sort,
		Status:        status,
		AddTime:       addTime,
		GoodsColor:    goodsColorStr,
		GoodsImg:      goodsImg,
	}
	err := models.DB.Create(&goods).Error
	if err != nil {
		con.Error(c, "增加失败", "/admin/goods/add")
		return
	}
	//增加图库信息
	//开启协程
	wg.Add(1)
	go func() {
		goodsImageList := c.PostFormArray("goods_image_list") //获取图片切片
		for _, v := range goodsImageList {
			goodsImgObj := models.GoodsImage{}
			goodsImgObj.GoodsId = goods.Id
			goodsImgObj.ImgUrl = v
			goodsImgObj.Sort = 10
			goodsImgObj.Status = 1
			goodsImgObj.AddTime = int(models.GetUnix())
			models.DB.Create(&goodsImgObj)
		}
		wg.Done()
	}()

	//增加规格包装
	wg.Add(1) //启动一个 goroutine 就登记+1
	//商品类型属性id和商品类型属性值一一对应
	go func() {
		attrIdList := c.PostFormArray("attr_id_list")
		attrValueList := c.PostFormArray("attr_value_list")
		for i := 0; i < len(attrIdList); i++ {
			//获取商品类型属性id
			goodsTypeAttributeId, attributeIdErr := models.Int(attrIdList[i])
			if attributeIdErr == nil {
				//获取商品类型属性的数据
				goodsTypeAttributeObj := models.GoodsTypeAttribute{Id: goodsTypeAttributeId}
				models.DB.Find(&goodsTypeAttributeObj)
				//给商品属性里面增加数据  规格包装
				goodsAttrObj := models.GoodsAttr{}
				goodsAttrObj.GoodsId = goods.Id
				goodsAttrObj.AttributeTitle = goodsTypeAttributeObj.Title
				goodsAttrObj.AttributeType = goodsTypeAttributeObj.AttrType
				goodsAttrObj.AttributeId = goodsTypeAttributeObj.Id
				goodsAttrObj.AttributeCateId = goodsTypeAttributeObj.CateId
				goodsAttrObj.AttributeValue = attrValueList[i] //值从attrValueList中获取
				goodsAttrObj.Status = 1
				goodsAttrObj.Sort = 10
				goodsAttrObj.AddTime = int(models.GetUnix())
				models.DB.Create(&goodsAttrObj)
			}
		}
		wg.Done() //goroutine 结束就登记-1
	}()
	//等待所有登记的 goroutine 都结束
	wg.Wait()

	con.Success(c, "增加数据成功", "/admin/goods")
}

func (con GoodsController) Edit(c *gin.Context) {
	//获取要修改的商品信息
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入参数错误", "/admin/goods")
		return
	}
	goods := models.Goods{Id: id}
	models.DB.Find(&goods)

	//获取商品分类
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid = 0").Preload("GoodsCateItems").Find(&goodsCateList)

	//获取商品颜色,以及选择的颜色
	//把商品颜色字符串转换成切片数组
	goodsColorSlice := strings.Split(goods.GoodsColor, ",")
	//定义一个商品颜色Map
	goodsColorMap := make(map[string]string)
	//循环颜色切片,把数据放入Map中
	for _, v := range goodsColorSlice {
		goodsColorMap[v] = v
	}
	//获取商品颜色列表
	goodsColorList := []models.GoodsColor{}
	models.DB.Where("status = 1").Find(&goodsColorList)
	//循环颜色列表,并与goodsColorMap比较,判断该商品是否有该颜色,并设置check
	for i := 0; i < len(goodsCateList); i++ {
		//断该商品是否有该颜色
		_, ok := goodsColorMap[models.String(goodsColorList[i].Id)]
		if ok { //该商品存在该颜色,设置Check=true
			goodsColorList[i].Checked = true
		}
	}

	//获取商品图库信息
	goodsImageList := []models.GoodsImage{}
	models.DB.Where("goods_id = ?", goods.Id).Find(&goodsImageList)

	//获取商品类型
	goodsTypeList := []models.GoodsType{}
	models.DB.Where("status = 1").Find(&goodsTypeList)

	//获取规格信息
	goodsAttr := []models.GoodsAttr{}
	models.DB.Where("goods_id = ?", goods.Id).Find(&goodsAttr)
	//拼接规格信息表单html
	goodsAttrStr := ""
	//循环规格信息数
	for _, v := range goodsAttr {
		if v.AttributeType == 1 { //当属性类型=1(单行文本框)时,显示的input html
			//fmt.Sprintf(): 拼接字符串
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: </span> <input type="hidden" name="attr_id_list" value="%v" />   <input type="text" name="attr_value_list" value="%v" /></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		} else if v.AttributeType == 2 { //当属性类型=2(多行文本框)时,显示的textareahtml
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span><input type="hidden" name="attr_id_list" value="%v" />  <textarea cols="50" rows="3" name="attr_value_list">%v</textarea></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		} else if v.AttributeType == 3 { //当属性类型=3(下拉框选择)时,显示的select html
			//获取当前类型对应的值(下拉框应该有多个选择的值)
			goodsTypeAttribute := models.GoodsTypeAttribute{Id: v.AttributeId}
			models.DB.Find(&goodsTypeAttribute)
			//把下拉框中的值转换成切片
			attrValueSlice := strings.Split(goodsTypeAttribute.AttrValue, "\n")
			//属性id input hidden
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span>  <input type="hidden" name="attr_id_list" value="%v" /> `, v.AttributeTitle, v.AttributeId)
			goodsAttrStr += fmt.Sprintf(`<select name="attr_value_list">`)
			//循环切片, 生成下拉框option
			for i := 0; i < len(attrValueSlice); i++ {
				if attrValueSlice[i] == v.AttributeValue { // 当前商品下拉属性 == 对应的商品下拉属性时, selected
					goodsAttrStr += fmt.Sprintf(`<option value="%v" selected >%v</option>`, attrValueSlice[i], attrValueSlice[i])
				} else {
					goodsAttrStr += fmt.Sprintf(`<option value="%v">%v</option>`, attrValueSlice[i], attrValueSlice[i])
				}
			}
			goodsAttrStr += fmt.Sprintf(`</select>`)
			goodsAttrStr += fmt.Sprintf(`</li>`)
		}
	}

	c.HTML(http.StatusOK, "admin/goods/edit.html", gin.H{
		"goods":          goods,
		"goodsCateList":  goodsCateList,
		"goodsColorList": goodsColorList,
		"goodsTypeList":  goodsTypeList,
		"goodsAttrStr":   goodsAttrStr,
		"goodsImageList": goodsImageList,
		"prePage":        c.Request.Referer(), // 获取上一页的地址
	})
}

//修改商品信息表单提交
func (con GoodsController) DoEdit(c *gin.Context) {
	//获取表单提交过来的数据,并进行判断是否合法
	id, err1 := models.Int(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "传入参数错误", "/admin/goods")
		return
	}
	prePage := c.PostForm("prePage") // 获取上一页地址
	title := c.PostForm("title")
	subTitle := c.PostForm("sub_title")
	goodsSn := c.PostForm("goods_sn")
	cateId, _ := models.Int(c.PostForm("cate_id"))
	goodsNumber, _ := models.Int(c.PostForm("goods_number"))
	//价格,注意小数点
	marketPrice, _ := models.Float(c.PostForm("market_price"))
	price, _ := models.Float(c.PostForm("price"))

	relationGoods := c.PostForm("relation_goods")
	goodsAttr := c.PostForm("goods_attr")
	goodsVersion := c.PostForm("goods_version")
	goodsGift := c.PostForm("goods_gift")
	goodsFitting := c.PostForm("goods_fitting")
	//颜色:获取的是切片
	goodsColorArr := c.PostFormArray("goods_color")

	goodsKeywords := c.PostForm("goods_keywords")
	goodsDesc := c.PostForm("goods_desc")
	goodsContent := c.PostForm("goods_content")
	isDelete, _ := models.Int(c.PostForm("is_delete"))
	isHot, _ := models.Int(c.PostForm("is_hot"))
	isBest, _ := models.Int(c.PostForm("is_best"))
	isNew, _ := models.Int(c.PostForm("is_new"))
	goodsTypeId, _ := models.Int(c.PostForm("goods_type_id"))
	sort, _ := models.Int(c.PostForm("sort"))
	status, _ := models.Int(c.PostForm("status"))

	//获取颜色信息,把颜色转换成字符串
	goodsColorStr := strings.Join(goodsColorArr, ",")

	//修改数据
	goods := models.Goods{Id: id}
	models.DB.Find(&goods)
	goods.Title = title
	goods.SubTitle = subTitle
	goods.GoodsSn = goodsSn
	goods.CateId = cateId
	goods.GoodsNumber = goodsNumber
	goods.MarketPrice = marketPrice
	goods.Price = price
	goods.RelationGoods = relationGoods
	goods.GoodsAttr = goodsAttr
	goods.GoodsVersion = goodsVersion
	goods.GoodsGift = goodsGift
	goods.GoodsFitting = goodsFitting
	goods.GoodsKeywords = goodsKeywords
	goods.GoodsDesc = goodsDesc
	goods.GoodsContent = goodsContent
	goods.IsDelete = isDelete
	goods.IsHot = isHot
	goods.IsBest = isBest
	goods.IsNew = isNew
	goods.GoodsTypeId = goodsTypeId
	goods.Sort = sort
	goods.Status = status
	goods.GoodsColor = goodsColorStr

	//上传图片,生成缩略图
	goodsImg, err2 := models.UploadImg(c, "goods_img")
	if err2 == nil && len(goodsImg) > 0 { // 说明修改了图片,那么就要设置图片属性
		goods.GoodsImg = goodsImg
		//判断本地图片才需要处理缩略图
		if models.GetOssStatus() != 1 {
			//开启协程
			wg.Add(1)
			go func() {
				models.ResizeGoodsImage(goodsImg)
				wg.Done()
			}()
		}
	}

	err := models.DB.Save(&goods).Error
	if err != nil {
		con.Error(c, "修改失败", "/admin/goods/edit?id="+models.String(id))
		return
	}
	//增加图库信息
	//开启协程
	wg.Add(1)
	go func() {
		goodsImageList := c.PostFormArray("goods_image_list") //获取图片切片
		for _, v := range goodsImageList {
			goodsImgObj := models.GoodsImage{}
			goodsImgObj.GoodsId = goods.Id
			goodsImgObj.ImgUrl = v
			goodsImgObj.Sort = 10
			goodsImgObj.Status = 1
			goodsImgObj.AddTime = int(models.GetUnix())
			models.DB.Create(&goodsImgObj)
		}
		wg.Done()
	}()

	//修改规格包装:1.删除当前商品下面的规格包装,2.重新执行增加
	//1.删除当前商品下面的规格包装
	goodsAttrObj := models.GoodsAttr{}
	models.DB.Where("goods_id = ?", goods.Id).Delete(&goodsAttrObj)

	//2.重新执行增加
	wg.Add(1) //启动一个 goroutine 就登记+1
	//商品类型属性id和商品类型属性值一一对应
	go func() {
		attrIdList := c.PostFormArray("attr_id_list")
		attrValueList := c.PostFormArray("attr_value_list")
		for i := 0; i < len(attrIdList); i++ {
			//获取商品类型属性id
			goodsTypeAttributeId, attributeIdErr := models.Int(attrIdList[i])
			if attributeIdErr == nil {
				//获取商品类型属性的数据
				goodsTypeAttributeObj := models.GoodsTypeAttribute{Id: goodsTypeAttributeId}
				models.DB.Find(&goodsTypeAttributeObj)
				//给商品属性里面增加数据  规格包装
				goodsAttrObj := models.GoodsAttr{}
				goodsAttrObj.GoodsId = goods.Id
				goodsAttrObj.AttributeTitle = goodsTypeAttributeObj.Title
				goodsAttrObj.AttributeType = goodsTypeAttributeObj.AttrType
				goodsAttrObj.AttributeId = goodsTypeAttributeObj.Id
				goodsAttrObj.AttributeCateId = goodsTypeAttributeObj.CateId
				goodsAttrObj.AttributeValue = attrValueList[i] //值从attrValueList中获取
				goodsAttrObj.Status = 1
				goodsAttrObj.Sort = 10
				goodsAttrObj.AddTime = int(models.GetUnix())
				models.DB.Create(&goodsAttrObj)
			}
		}
		wg.Done() //goroutine 结束就登记-1
	}()
	//等待所有登记的 goroutine 都结束
	wg.Wait()

	if len(prePage) > 0 { //跳转到上一页
		con.Success(c, "修改数据成功", prePage)
		return
	}
	con.Success(c, "修改数据成功", "/admin/goods")
}

//获取商品类型对应的属性
func (con GoodsController) GoodsTypeAttribute(c *gin.Context) {
	cateId, err1 := models.Int(c.Query("cateId"))
	goodsTypeAttributeList := []models.GoodsTypeAttribute{}
	err2 := models.DB.Where("cate_id = ?", cateId).Find(&goodsTypeAttributeList).Error
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  goodsTypeAttributeList,
	})
}

//富文本编辑器上传图片方法
func (con GoodsController) EditorImageUpload(c *gin.Context) {
	imgDir, err := models.UploadImg(c, "file") //传递的参数默认是file
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"link": "", //富文本要求返回的格式: {link: 'path/to/image.jpg'}
		})
		return
	}

	//判断本地图片才需要处理缩略图
	if models.GetOssStatus() != 1 {
		//开启协程
		wg.Add(1)
		go func() {
			models.ResizeGoodsImage(imgDir)
			wg.Done()
		}()
		//本地图片,返回本地图片地址
		c.JSON(http.StatusOK, gin.H{
			"link": "/" + imgDir,
		})
	} else {
		//云服务器对象存储图片,返回云服务器图片地址
		c.JSON(http.StatusOK, gin.H{
			"link": models.GetSettingFromColumn("OssDomain") + imgDir,
		})
	}
}

//商品上传图片方法
func (con GoodsController) GoodsImageUpload(c *gin.Context) {
	imgDir, err := models.UploadImg(c, "file") //传递的参数默认是file
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"link": "", //富文本要求返回的格式: {link: 'path/to/image.jpg'}
		})
		return
	}

	//判断本地图片才需要处理缩略图
	if models.GetOssStatus() != 1 {
		//开启协程
		wg.Add(1)
		go func() {
			models.ResizeGoodsImage(imgDir)
			wg.Done()
		}()

	}
	//返回图片地址
	c.JSON(http.StatusOK, gin.H{
		"link": imgDir,
	})
}

//修改商品图库关联的颜色
func (con GoodsController) ChangeGoodsImageColor(c *gin.Context) {
	//获取图片id 获取颜色id
	goodsImageId, err1 := models.Int(c.Query("goods_image_id"))
	colorId, err2 := models.Int(c.Query("color_id"))
	goodsImage := models.GoodsImage{Id: goodsImageId}
	models.DB.Find(&goodsImage)
	goodsImage.ColorId = colorId
	err3 := models.DB.Save(&goodsImage).Error
	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  "更新失败",
			"success": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":  "更新成功",
		"success": true,
	})
}

//删除图库
func (con GoodsController) RemoveGoodsImage(c *gin.Context) {
	//获取图片id
	goodsImageId, err1 := models.Int(c.Query("goods_image_id"))
	goodsImage := models.GoodsImage{Id: goodsImageId}
	//获取图片
	models.DB.Find(&goodsImage)
	fileName := goodsImage.ImgUrl

	//todo 是否删除服务器保存的图片? 判断是否本地 or 云服务器图片

	//判断是否开启oss
	if models.GetOssStatus() == 1 {
		//return CosDeleteImg(fileName)
	} else {
		err3 := os.Remove(strings.TrimLeft(fileName, "/"))
		if err3 != nil {
			c.JSON(http.StatusOK, gin.H{
				"result":  err3.Error(),
				"success": false,
			})
			return
		}
	}

	fmt.Println(goodsImage.ImgUrl)
	//删除数据库中的数据
	err2 := models.DB.Delete(&goodsImage).Error

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  "删除失败",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":  "删除成功",
		"success": true,
	})
}

//删除
func (con GoodsController) Delete(c *gin.Context) {
	//获取提交的表单数据
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goods")
		return
	}
	//查询商品
	goods := models.Goods{Id: id}
	models.DB.Find(&goods)
	//软删除
	goods.IsDelete = 1
	goods.Status = 0
	models.DB.Save(&goods)

	//获取上一页地址,判断是否存在,如果存在则跳转,不存在则跳转到列表首页
	prePage := c.Request.Referer()
	if len(prePage) > 0 {
		con.Success(c, "删除数据成功", prePage)
		return
	}
	con.Success(c, "删除数据成功", "/admin/goods")
}
