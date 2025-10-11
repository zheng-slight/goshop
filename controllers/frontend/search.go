package frontend

//Elasticsearch 控制器

import (
	"context"
	"encoding/json"
	"fmt"
	"goshop/models"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type SearchController struct {
	BaseController
}

// 初始化的时候判断索引goods是否存在,创建索引配置映射
func (con SearchController) Index(c *gin.Context) {
	exists, err := models.EsClient.IndexExists("goods").Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	print(exists)
	if !exists {
		// 配置映射
		mapping := `
		{
			"settings": {  //设置
			  "number_of_shards": 1,  //分区数配置
			  "number_of_replicas": 0  //副本数配置
			},
			"mappings": {  //映射
			  "properties": {
				"Content": {  //映射属性
				  "type": "text",  //类型
				  "analyzer": "ik_max_word", // 检测粒度
				  "search_analyzer": "ik_max_word"  //搜索粒度
				},
				"Title": {
				  "type": "text",
				  "analyzer": "ik_max_word",
				  "search_analyzer": "ik_max_word"
				}
			  }
			}
		  }
		`
		//注意：增加的写法-创建索引配置映射
		_, err := models.EsClient.CreateIndex("goods").Body(mapping).Do(context.Background())
		if err != nil {
			// Handle error
			fmt.Println(err)
		}
	}

	c.String(200, "创建索引配置映射成功")
}

// 增加商品数据到es
func (con SearchController) AddGoods(c *gin.Context) {
	//获取数据库中的商品数据
	goods := []models.Goods{}
	models.DB.Find(&goods)
	//循环商品,把每个商品存入es
	for i := 0; i < len(goods); i++ {
		_, err := models.EsClient.Index().
			Index("goods").                //设置索引
			Type("_doc").                  //设置类型
			Id(strconv.Itoa(goods[i].Id)). //设置id
			BodyJson(goods[i]).            //设置商品数据(结构体格式)
			Do(context.Background())
		if err != nil {
			// Handle error
			fmt.Println(err)
		}
	}

	c.String(200, "AddGoods success")
}

// 更新数据
func (con SearchController) UpdateGoods(c *gin.Context) {
	goods := []models.Goods{}
	models.DB.Find(&goods)
	goods[0].Title = "我是修改后的数据"
	goods[0].GoodsContent = "我是修改后的数据GoodsContent"

	_, err := models.EsClient.Update().
		Index("goods").
		Type("_doc").
		Id("19").      //要修改的数据id
		Doc(goods[0]). //要修改的数据结构体
		Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	c.String(200, "修改数据 success")
}

// 删除
func (con SearchController) DeleteGoods(c *gin.Context) {
	_, err := models.EsClient.Delete().
		Index("goods").
		Type("_doc").
		Id("19").
		Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	c.String(200, "删除成功 success")
}

// 查询一条数据
func (con SearchController) GetOne(c *gin.Context) {
	//defer 操作:捕获panic数据
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.String(200, "GetOne Error")
		}
	}()

	result, err := models.EsClient.Get().
		Index("goods").
		Type("_doc").
		Id("19").
		Do(context.Background())
	if err != nil { //判断数据是否存在,不存在则panic
		// Some other kind of error
		panic(err)
	}

	goods := models.Goods{}               //实例化一个商品结构体
	json.Unmarshal(result.Source, &goods) //把result结果解析到goods中

	c.JSON(200, gin.H{
		"goods": goods,
	})
}

// 模糊查询数据
func (con SearchController) Query(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.String(200, "Query Error")
		}
	}()
	//模糊查询操作
	query := elastic.NewMatchQuery("Title", "手机") //Title中包含 手机 的数据
	searchResult, err := models.EsClient.Search().
		Index("goods").          // search in index "goods"
		Query(query).            // specify the query
		Do(context.Background()) // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	goods := models.Goods{}
	c.JSON(200, gin.H{
		"searchResult": searchResult.Each(reflect.TypeOf(goods)), //查询的结果:reflect.TypeOf(goods)类型断言,可以判断是否商品结构体
	})
}

// 分页查询
func (con SearchController) PagingQuery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.String(200, "Query Error")
		}
	}()

	page, _ := strconv.Atoi(c.Query("page")) //获取当前页码数
	if page == 0 {
		page = 1
	}
	pageSize := 2
	query := elastic.NewMatchQuery("Title", "手机")
	searchResult, err := models.EsClient.Search().
		Index("goods").                             // search in index "goods"
		Query(query).                               // specify the query
		Sort("Id", true).                           // true 表示升序   false 降序
		From((page - 1) * pageSize).Size(pageSize). // 分页查询
		Do(context.Background())                    // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	goods := models.Goods{}
	c.JSON(200, gin.H{
		"searchResult": searchResult.Each(reflect.TypeOf(goods)),
	})

}

// 条件筛选查询
func (con SearchController) FilterQuery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			c.String(200, "Query Error")
		}
	}()

	//筛选
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("Title", "小米"))
	boolQ.Filter(elastic.NewRangeQuery("Id").Gt(19)) //Id 大于19
	boolQ.Filter(elastic.NewRangeQuery("Id").Lt(42)) //Id 小于42
	searchResult, err := models.EsClient.Search().
		Index("goods").
		Type("_doc").
		Sort("Id", true).
		Query(boolQ).
		Do(context.Background())

	if err != nil {
		fmt.Println(err)
	}
	goodsList := []models.Goods{}
	var goods models.Goods
	for _, item := range searchResult.Each(reflect.TypeOf(goods)) { //循环搜索结果,并把结果类型断言goods,如果结果类型是商品数据类型,则循环处理
		t := item.(models.Goods)
		fmt.Printf("Id:%v 标题：%v\n", t.Id, t.Title)
		goodsList = append(goodsList, t)
	}

	c.JSON(200, gin.H{
		"goodsList": goodsList,
	})
}
