package models

//cookie结构体

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

//定义结构体  缓存结构体 私有
type ginCookie struct{}

//写入数据的方法
func (cookie ginCookie) Set(c *gin.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	//保存数据时,需要对数据进行des加密操作
	desKey := []byte("123abcQS")  // 注意:key必须是8位
	encData, _ := DesEncrypt(bytes, desKey)
	c.SetCookie(key, string(encData), 3600, "/", "127.0.0.1", false, true)
}

//获取数据的方法
func (cookie ginCookie) Get(c *gin.Context, key string, obj interface{}) bool {
	valueStr, err1 := c.Cookie(key)
	if err1 == nil && valueStr != "" && valueStr != "[]" {
		//获取数据时,需要对数据进行des解密
		desKey := []byte("123abcQS")  // 注意:key必须是8位
		decData, e := DesDecrypt([]byte(valueStr), desKey)
		if e != nil {
			return false
		} else {
			err2 := json.Unmarshal([]byte(decData), obj)
			return err2 == nil
		}
	}
	return false
}

//删除数据
func (cookie ginCookie) Remove(c *gin.Context, key string) bool {
	c.SetCookie(key, "", -1, "/", c.Request.Host, false, true)
	return true
}

//实例化结构体
var Cookie = &ginCookie{}
